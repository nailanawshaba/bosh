module Bosh::Director
  module DeploymentPlan
    module Steps
      class CreateVmStep
        include PasswordHelper

        def initialize(instance_plan, agent_broadcaster, vm_deleter, disks, tags = {}, use_existing = false)
          @instance_plan = instance_plan
          @agent_broadcaster = agent_broadcaster
          @vm_deleter = vm_deleter
          @disks = disks
          @use_existing = use_existing
          @tags = tags
          @logger = Config.logger
        end

        def perform
          instance = @instance_plan.instance

          cpi_factory, stemcell_cid = choose_factory_and_stemcell_cid(@instance_plan, @use_existing)

          instance_model = instance.model
          @logger.info('Creating VM')

          # create the vm
          vm = create(
            instance,
            stemcell_cid,
            instance.cloud_properties,
            @instance_plan.network_settings_hash,
            @disks,
            instance.env,
            cpi_factory,
          )

          begin
            update_metadata(@instance_plan, vm, cpi_factory)
          rescue Exception => e
            # cleanup in case of failure
            @logger.error("Failed to create/contact VM #{instance_model.vm_cid}: #{e.inspect}")
            if Config.keep_unreachable_vms
              @logger.info('Keeping the VM for debugging')
            else
              @vm_deleter.delete_for_instance(instance_model)
            end
            raise e
          end
        end

        private

        # def update_metadata(vm, factory)
        def update_metadata(instance_plan, vm, factory)
          instance_model = instance_plan.instance.model
          MetadataUpdater.build.update_vm_metadata(instance_model, vm, @tags, factory)
          agent_client = AgentClient.with_agent_id(instance_model.agent_id)
          agent_client.wait_until_ready

          # delete arp entries from all other agents
          if Config.flush_arp
            ip_addresses = instance_plan.network_settings_hash.map do |index, network|
              network['ip']
            end.compact

            # TODO:  we may want to extract this
            @agent_broadcaster.delete_arp_entries(vm.cid, ip_addresses)
          end
        end

        def choose_factory_and_stemcell_cid(instance_plan, use_existing)
          if use_existing && !!instance_plan.existing_instance.availability_zone
            factory = CloudFactory.create_from_deployment(instance_plan.existing_instance.deployment)

            stemcell = instance_plan.instance.stemcell
            cpi = factory.get_name_for_az(instance_plan.existing_instance.availability_zone)
            stemcell_cid = stemcell.models.find { |model| model.cpi == cpi }.cid
            return factory, stemcell_cid
          else
            return CloudFactory.create_with_latest_configs, instance_plan.instance.stemcell_cid
          end
        end

        def create(instance, stemcell_cid, cloud_properties, network_settings, disks, env, cloud_factory)
          instance_model = instance.model
          deployment_name = instance_model.deployment.name
          parent_id = add_event(deployment_name, instance_model.name, 'create')
          agent_id = SecureRandom.uuid

          config_server_client = ConfigServer::ClientFactory.create(@logger) .create_client
          env = config_server_client.interpolate_with_versioning(env, instance.desired_variable_set)
          cloud_properties = config_server_client.interpolate_with_versioning(cloud_properties, instance.desired_variable_set)
          network_settings = config_server_client.interpolate_with_versioning(network_settings, instance.desired_variable_set)

          cpi = cloud_factory.get_name_for_az(instance_model.availability_zone)

          vm_options = {instance: instance_model, agent_id: agent_id, cpi: cpi}

          env['bosh'] ||= {}
          env['bosh'] = Config.agent_env.merge(env['bosh'])

          if Config.nats_server_ca
            env['bosh'] ||= {}
            env['bosh']['mbus'] ||= {}
            env['bosh']['mbus']['cert'] ||= {}
            env['bosh']['mbus']['cert']['ca'] = Config.nats_server_ca
            cert_generator = NatsClientCertGenerator.new(@logger)
            agent_cert_key_result = cert_generator.generate_nats_client_certificate "#{agent_id}.agent.bosh-internal"
            env['bosh']['mbus']['cert']['certificate'] = agent_cert_key_result[:cert].to_pem
            env['bosh']['mbus']['cert']['private_key'] = agent_cert_key_result[:key].to_pem
          end

          password = env.fetch('bosh', {}).fetch('password', "")
          if Config.generate_vm_passwords && password == ""
            env['bosh'] ||= {}
            env['bosh']['password'] = sha512_hashed_password
          end

          if instance_model.job
            env['bosh'] ||= {}
            env['bosh']['group'] = Canonicalizer.canonicalize("#{Bosh::Director::Config.name}-#{deployment_name}-#{instance_model.job}")
            env['bosh']['groups'] = [
              Bosh::Director::Config.name,
              deployment_name,
              instance_model.job,
              "#{Bosh::Director::Config.name}-#{deployment_name}",
              "#{deployment_name}-#{instance_model.job}",
              "#{Bosh::Director::Config.name}-#{deployment_name}-#{instance_model.job}",
            ]
            env['bosh']['groups'].map! { |name| Canonicalizer.canonicalize(name) }
          end

          count = 0
          begin
            cloud = cloud_factory.get(vm_options[:cpi])
            vm_cid = cloud.create_vm(agent_id, stemcell_cid, cloud_properties, network_settings, disks, env)
          rescue Bosh::Clouds::VMCreationFailed => e
            count += 1
            @logger.error("failed to create VM, retrying (#{count})")
            retry if e.ok_to_retry && count < Config.max_vm_create_tries
            raise e
          end

          vm_options[:cid] = vm_cid
          vm_options[:created_at] = Time.now
          Models::Vm.create(vm_options)
        rescue => e
          @logger.error("error creating vm: #{e.message}")
          if vm_cid
            parent_id = add_event(deployment_name, instance_model.name, 'delete', vm_cid)
            @vm_deleter.delete_vm_by_cid(vm_cid)
            add_event(deployment_name, instance_model.name, 'delete', vm_cid, parent_id)
          end
          raise e
        ensure
          add_event(deployment_name, instance_model.name, 'create', vm_cid, parent_id, e)
        end

        def add_event(deployment_name, instance_name, action, object_name = nil, parent_id = nil, error = nil)
          event = Config.current_job.event_manager.create_event(
            {
              parent_id: parent_id,
              user: Config.current_job.username,
              action: action,
              object_type: 'vm',
              object_name: object_name,
              task: Config.current_job.task_id,
              deployment: deployment_name,
              instance: instance_name,
              error: error
            })
          event.id
        end
      end
    end
  end
end
