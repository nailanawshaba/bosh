require 'spec_helper'

module Bosh::Director
  module DeploymentPlan
    module Steps
      # describe UpdateInstanceSettingsStep do
      #   subject(:step) { UpdateInstanceSettingsStep.new(instance) }

      #   let(:instance) { Models::Instance.make }
      #   let!(:vm) { Models::Vm.make(instance: instance, active: true) }
      #   let(:disk1) { Models::PersistentDisk.make(instance: instance, name: '') }
      #   let(:disk2) { Models::PersistentDisk.make(instance: instance, name: 'unmanaged') }

      #   describe '#perform' do
      #     context 'when there are unmanaged persistent disks' do
      #       it 'updates agent disk associations' do
      #         step.perform
      #       end
      #     end
          
      #     it 'updates the agent settings and VM table with configured trusted certs' do
      #       expect(agent_client).to receive(:update_settings).with(trusted_certs, {})
      #       expect { step.perform }.to change(vm.trusted_certs).from('').to('fake_cert')
      #     end
      #   end
      # end
    end
  end
end
