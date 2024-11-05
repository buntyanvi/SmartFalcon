// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go-apiv2/common"
	"github.com/hyperledger/fabric-protos-go-apiv2/orderer"
	"github.com/hyperledger/fabric/common/deliverclient/blocksprovider"
	"github.com/hyperledger/fabric/common/deliverclient/orderers"
)

type DeliverClientRequester struct {
	ConnectStub        func(*common.Envelope, *orderers.Endpoint) (orderer.AtomicBroadcast_DeliverClient, func(), error)
	connectMutex       sync.RWMutex
	connectArgsForCall []struct {
		arg1 *common.Envelope
		arg2 *orderers.Endpoint
	}
	connectReturns struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 func()
		result3 error
	}
	connectReturnsOnCall map[int]struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 func()
		result3 error
	}
	SeekInfoHeadersFromStub        func(uint64) (*common.Envelope, error)
	seekInfoHeadersFromMutex       sync.RWMutex
	seekInfoHeadersFromArgsForCall []struct {
		arg1 uint64
	}
	seekInfoHeadersFromReturns struct {
		result1 *common.Envelope
		result2 error
	}
	seekInfoHeadersFromReturnsOnCall map[int]struct {
		result1 *common.Envelope
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeliverClientRequester) Connect(arg1 *common.Envelope, arg2 *orderers.Endpoint) (orderer.AtomicBroadcast_DeliverClient, func(), error) {
	fake.connectMutex.Lock()
	ret, specificReturn := fake.connectReturnsOnCall[len(fake.connectArgsForCall)]
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct {
		arg1 *common.Envelope
		arg2 *orderers.Endpoint
	}{arg1, arg2})
	stub := fake.ConnectStub
	fakeReturns := fake.connectReturns
	fake.recordInvocation("Connect", []interface{}{arg1, arg2})
	fake.connectMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *DeliverClientRequester) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *DeliverClientRequester) ConnectCalls(stub func(*common.Envelope, *orderers.Endpoint) (orderer.AtomicBroadcast_DeliverClient, func(), error)) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = stub
}

func (fake *DeliverClientRequester) ConnectArgsForCall(i int) (*common.Envelope, *orderers.Endpoint) {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	argsForCall := fake.connectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *DeliverClientRequester) ConnectReturns(result1 orderer.AtomicBroadcast_DeliverClient, result2 func(), result3 error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 func()
		result3 error
	}{result1, result2, result3}
}

func (fake *DeliverClientRequester) ConnectReturnsOnCall(i int, result1 orderer.AtomicBroadcast_DeliverClient, result2 func(), result3 error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = nil
	if fake.connectReturnsOnCall == nil {
		fake.connectReturnsOnCall = make(map[int]struct {
			result1 orderer.AtomicBroadcast_DeliverClient
			result2 func()
			result3 error
		})
	}
	fake.connectReturnsOnCall[i] = struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 func()
		result3 error
	}{result1, result2, result3}
}

func (fake *DeliverClientRequester) SeekInfoHeadersFrom(arg1 uint64) (*common.Envelope, error) {
	fake.seekInfoHeadersFromMutex.Lock()
	ret, specificReturn := fake.seekInfoHeadersFromReturnsOnCall[len(fake.seekInfoHeadersFromArgsForCall)]
	fake.seekInfoHeadersFromArgsForCall = append(fake.seekInfoHeadersFromArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.SeekInfoHeadersFromStub
	fakeReturns := fake.seekInfoHeadersFromReturns
	fake.recordInvocation("SeekInfoHeadersFrom", []interface{}{arg1})
	fake.seekInfoHeadersFromMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeliverClientRequester) SeekInfoHeadersFromCallCount() int {
	fake.seekInfoHeadersFromMutex.RLock()
	defer fake.seekInfoHeadersFromMutex.RUnlock()
	return len(fake.seekInfoHeadersFromArgsForCall)
}

func (fake *DeliverClientRequester) SeekInfoHeadersFromCalls(stub func(uint64) (*common.Envelope, error)) {
	fake.seekInfoHeadersFromMutex.Lock()
	defer fake.seekInfoHeadersFromMutex.Unlock()
	fake.SeekInfoHeadersFromStub = stub
}

func (fake *DeliverClientRequester) SeekInfoHeadersFromArgsForCall(i int) uint64 {
	fake.seekInfoHeadersFromMutex.RLock()
	defer fake.seekInfoHeadersFromMutex.RUnlock()
	argsForCall := fake.seekInfoHeadersFromArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeliverClientRequester) SeekInfoHeadersFromReturns(result1 *common.Envelope, result2 error) {
	fake.seekInfoHeadersFromMutex.Lock()
	defer fake.seekInfoHeadersFromMutex.Unlock()
	fake.SeekInfoHeadersFromStub = nil
	fake.seekInfoHeadersFromReturns = struct {
		result1 *common.Envelope
		result2 error
	}{result1, result2}
}

func (fake *DeliverClientRequester) SeekInfoHeadersFromReturnsOnCall(i int, result1 *common.Envelope, result2 error) {
	fake.seekInfoHeadersFromMutex.Lock()
	defer fake.seekInfoHeadersFromMutex.Unlock()
	fake.SeekInfoHeadersFromStub = nil
	if fake.seekInfoHeadersFromReturnsOnCall == nil {
		fake.seekInfoHeadersFromReturnsOnCall = make(map[int]struct {
			result1 *common.Envelope
			result2 error
		})
	}
	fake.seekInfoHeadersFromReturnsOnCall[i] = struct {
		result1 *common.Envelope
		result2 error
	}{result1, result2}
}

func (fake *DeliverClientRequester) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	fake.seekInfoHeadersFromMutex.RLock()
	defer fake.seekInfoHeadersFromMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeliverClientRequester) recordInvocation(key string, args []interface{}) {
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

var _ blocksprovider.DeliverClientRequester = new(DeliverClientRequester)