// This file was generated by counterfeiter
package netfakes

import (
	"net"
	"sync"
	"time"
)

type FakeConn struct {
	ReadStub        func(b []byte) (n int, err error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		b []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	WriteStub        func(b []byte) (n int, err error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		b []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	LocalAddrStub        func() net.Addr
	localAddrMutex       sync.RWMutex
	localAddrArgsForCall []struct{}
	localAddrReturns     struct {
		result1 net.Addr
	}
	RemoteAddrStub        func() net.Addr
	remoteAddrMutex       sync.RWMutex
	remoteAddrArgsForCall []struct{}
	remoteAddrReturns     struct {
		result1 net.Addr
	}
	SetDeadlineStub        func(t time.Time) error
	setDeadlineMutex       sync.RWMutex
	setDeadlineArgsForCall []struct {
		t time.Time
	}
	setDeadlineReturns struct {
		result1 error
	}
	SetReadDeadlineStub        func(t time.Time) error
	setReadDeadlineMutex       sync.RWMutex
	setReadDeadlineArgsForCall []struct {
		t time.Time
	}
	setReadDeadlineReturns struct {
		result1 error
	}
	SetWriteDeadlineStub        func(t time.Time) error
	setWriteDeadlineMutex       sync.RWMutex
	setWriteDeadlineArgsForCall []struct {
		t time.Time
	}
	setWriteDeadlineReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConn) Read(b []byte) (n int, err error) {
	var bCopy []byte
	if b != nil {
		bCopy = make([]byte, len(b))
		copy(bCopy, b)
	}
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		b []byte
	}{bCopy})
	fake.recordInvocation("Read", []interface{}{bCopy})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(b)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeConn) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeConn) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].b
}

func (fake *FakeConn) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Write(b []byte) (n int, err error) {
	var bCopy []byte
	if b != nil {
		bCopy = make([]byte, len(b))
		copy(bCopy, b)
	}
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		b []byte
	}{bCopy})
	fake.recordInvocation("Write", []interface{}{bCopy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(b)
	} else {
		return fake.writeReturns.result1, fake.writeReturns.result2
	}
}

func (fake *FakeConn) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeConn) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].b
}

func (fake *FakeConn) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeConn) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeConn) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConn) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) LocalAddr() net.Addr {
	fake.localAddrMutex.Lock()
	fake.localAddrArgsForCall = append(fake.localAddrArgsForCall, struct{}{})
	fake.recordInvocation("LocalAddr", []interface{}{})
	fake.localAddrMutex.Unlock()
	if fake.LocalAddrStub != nil {
		return fake.LocalAddrStub()
	} else {
		return fake.localAddrReturns.result1
	}
}

func (fake *FakeConn) LocalAddrCallCount() int {
	fake.localAddrMutex.RLock()
	defer fake.localAddrMutex.RUnlock()
	return len(fake.localAddrArgsForCall)
}

func (fake *FakeConn) LocalAddrReturns(result1 net.Addr) {
	fake.LocalAddrStub = nil
	fake.localAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeConn) RemoteAddr() net.Addr {
	fake.remoteAddrMutex.Lock()
	fake.remoteAddrArgsForCall = append(fake.remoteAddrArgsForCall, struct{}{})
	fake.recordInvocation("RemoteAddr", []interface{}{})
	fake.remoteAddrMutex.Unlock()
	if fake.RemoteAddrStub != nil {
		return fake.RemoteAddrStub()
	} else {
		return fake.remoteAddrReturns.result1
	}
}

func (fake *FakeConn) RemoteAddrCallCount() int {
	fake.remoteAddrMutex.RLock()
	defer fake.remoteAddrMutex.RUnlock()
	return len(fake.remoteAddrArgsForCall)
}

func (fake *FakeConn) RemoteAddrReturns(result1 net.Addr) {
	fake.RemoteAddrStub = nil
	fake.remoteAddrReturns = struct {
		result1 net.Addr
	}{result1}
}

func (fake *FakeConn) SetDeadline(t time.Time) error {
	fake.setDeadlineMutex.Lock()
	fake.setDeadlineArgsForCall = append(fake.setDeadlineArgsForCall, struct {
		t time.Time
	}{t})
	fake.recordInvocation("SetDeadline", []interface{}{t})
	fake.setDeadlineMutex.Unlock()
	if fake.SetDeadlineStub != nil {
		return fake.SetDeadlineStub(t)
	} else {
		return fake.setDeadlineReturns.result1
	}
}

func (fake *FakeConn) SetDeadlineCallCount() int {
	fake.setDeadlineMutex.RLock()
	defer fake.setDeadlineMutex.RUnlock()
	return len(fake.setDeadlineArgsForCall)
}

func (fake *FakeConn) SetDeadlineArgsForCall(i int) time.Time {
	fake.setDeadlineMutex.RLock()
	defer fake.setDeadlineMutex.RUnlock()
	return fake.setDeadlineArgsForCall[i].t
}

func (fake *FakeConn) SetDeadlineReturns(result1 error) {
	fake.SetDeadlineStub = nil
	fake.setDeadlineReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) SetReadDeadline(t time.Time) error {
	fake.setReadDeadlineMutex.Lock()
	fake.setReadDeadlineArgsForCall = append(fake.setReadDeadlineArgsForCall, struct {
		t time.Time
	}{t})
	fake.recordInvocation("SetReadDeadline", []interface{}{t})
	fake.setReadDeadlineMutex.Unlock()
	if fake.SetReadDeadlineStub != nil {
		return fake.SetReadDeadlineStub(t)
	} else {
		return fake.setReadDeadlineReturns.result1
	}
}

func (fake *FakeConn) SetReadDeadlineCallCount() int {
	fake.setReadDeadlineMutex.RLock()
	defer fake.setReadDeadlineMutex.RUnlock()
	return len(fake.setReadDeadlineArgsForCall)
}

func (fake *FakeConn) SetReadDeadlineArgsForCall(i int) time.Time {
	fake.setReadDeadlineMutex.RLock()
	defer fake.setReadDeadlineMutex.RUnlock()
	return fake.setReadDeadlineArgsForCall[i].t
}

func (fake *FakeConn) SetReadDeadlineReturns(result1 error) {
	fake.SetReadDeadlineStub = nil
	fake.setReadDeadlineReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) SetWriteDeadline(t time.Time) error {
	fake.setWriteDeadlineMutex.Lock()
	fake.setWriteDeadlineArgsForCall = append(fake.setWriteDeadlineArgsForCall, struct {
		t time.Time
	}{t})
	fake.recordInvocation("SetWriteDeadline", []interface{}{t})
	fake.setWriteDeadlineMutex.Unlock()
	if fake.SetWriteDeadlineStub != nil {
		return fake.SetWriteDeadlineStub(t)
	} else {
		return fake.setWriteDeadlineReturns.result1
	}
}

func (fake *FakeConn) SetWriteDeadlineCallCount() int {
	fake.setWriteDeadlineMutex.RLock()
	defer fake.setWriteDeadlineMutex.RUnlock()
	return len(fake.setWriteDeadlineArgsForCall)
}

func (fake *FakeConn) SetWriteDeadlineArgsForCall(i int) time.Time {
	fake.setWriteDeadlineMutex.RLock()
	defer fake.setWriteDeadlineMutex.RUnlock()
	return fake.setWriteDeadlineArgsForCall[i].t
}

func (fake *FakeConn) SetWriteDeadlineReturns(result1 error) {
	fake.SetWriteDeadlineStub = nil
	fake.setWriteDeadlineReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConn) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.localAddrMutex.RLock()
	defer fake.localAddrMutex.RUnlock()
	fake.remoteAddrMutex.RLock()
	defer fake.remoteAddrMutex.RUnlock()
	fake.setDeadlineMutex.RLock()
	defer fake.setDeadlineMutex.RUnlock()
	fake.setReadDeadlineMutex.RLock()
	defer fake.setReadDeadlineMutex.RUnlock()
	fake.setWriteDeadlineMutex.RLock()
	defer fake.setWriteDeadlineMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConn) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ net.Conn = new(FakeConn)
