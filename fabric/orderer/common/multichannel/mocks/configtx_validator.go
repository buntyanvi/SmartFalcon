// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go-apiv2/common"
)

type ConfigTXValidator struct {
	ChannelIDStub        func() string
	channelIDMutex       sync.RWMutex
	channelIDArgsForCall []struct {
	}
	channelIDReturns struct {
		result1 string
	}
	channelIDReturnsOnCall map[int]struct {
		result1 string
	}
	ConfigProtoStub        func() *common.Config
	configProtoMutex       sync.RWMutex
	configProtoArgsForCall []struct {
	}
	configProtoReturns struct {
		result1 *common.Config
	}
	configProtoReturnsOnCall map[int]struct {
		result1 *common.Config
	}
	ProposeConfigUpdateStub        func(*common.Envelope) (*common.ConfigEnvelope, error)
	proposeConfigUpdateMutex       sync.RWMutex
	proposeConfigUpdateArgsForCall []struct {
		arg1 *common.Envelope
	}
	proposeConfigUpdateReturns struct {
		result1 *common.ConfigEnvelope
		result2 error
	}
	proposeConfigUpdateReturnsOnCall map[int]struct {
		result1 *common.ConfigEnvelope
		result2 error
	}
	SequenceStub        func() uint64
	sequenceMutex       sync.RWMutex
	sequenceArgsForCall []struct {
	}
	sequenceReturns struct {
		result1 uint64
	}
	sequenceReturnsOnCall map[int]struct {
		result1 uint64
	}
	ValidateStub        func(*common.ConfigEnvelope) error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		arg1 *common.ConfigEnvelope
	}
	validateReturns struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigTXValidator) ChannelID() string {
	fake.channelIDMutex.Lock()
	ret, specificReturn := fake.channelIDReturnsOnCall[len(fake.channelIDArgsForCall)]
	fake.channelIDArgsForCall = append(fake.channelIDArgsForCall, struct {
	}{})
	stub := fake.ChannelIDStub
	fakeReturns := fake.channelIDReturns
	fake.recordInvocation("ChannelID", []interface{}{})
	fake.channelIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigTXValidator) ChannelIDCallCount() int {
	fake.channelIDMutex.RLock()
	defer fake.channelIDMutex.RUnlock()
	return len(fake.channelIDArgsForCall)
}

func (fake *ConfigTXValidator) ChannelIDCalls(stub func() string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = stub
}

func (fake *ConfigTXValidator) ChannelIDReturns(result1 string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = nil
	fake.channelIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *ConfigTXValidator) ChannelIDReturnsOnCall(i int, result1 string) {
	fake.channelIDMutex.Lock()
	defer fake.channelIDMutex.Unlock()
	fake.ChannelIDStub = nil
	if fake.channelIDReturnsOnCall == nil {
		fake.channelIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.channelIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ConfigTXValidator) ConfigProto() *common.Config {
	fake.configProtoMutex.Lock()
	ret, specificReturn := fake.configProtoReturnsOnCall[len(fake.configProtoArgsForCall)]
	fake.configProtoArgsForCall = append(fake.configProtoArgsForCall, struct {
	}{})
	stub := fake.ConfigProtoStub
	fakeReturns := fake.configProtoReturns
	fake.recordInvocation("ConfigProto", []interface{}{})
	fake.configProtoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigTXValidator) ConfigProtoCallCount() int {
	fake.configProtoMutex.RLock()
	defer fake.configProtoMutex.RUnlock()
	return len(fake.configProtoArgsForCall)
}

func (fake *ConfigTXValidator) ConfigProtoCalls(stub func() *common.Config) {
	fake.configProtoMutex.Lock()
	defer fake.configProtoMutex.Unlock()
	fake.ConfigProtoStub = stub
}

func (fake *ConfigTXValidator) ConfigProtoReturns(result1 *common.Config) {
	fake.configProtoMutex.Lock()
	defer fake.configProtoMutex.Unlock()
	fake.ConfigProtoStub = nil
	fake.configProtoReturns = struct {
		result1 *common.Config
	}{result1}
}

func (fake *ConfigTXValidator) ConfigProtoReturnsOnCall(i int, result1 *common.Config) {
	fake.configProtoMutex.Lock()
	defer fake.configProtoMutex.Unlock()
	fake.ConfigProtoStub = nil
	if fake.configProtoReturnsOnCall == nil {
		fake.configProtoReturnsOnCall = make(map[int]struct {
			result1 *common.Config
		})
	}
	fake.configProtoReturnsOnCall[i] = struct {
		result1 *common.Config
	}{result1}
}

func (fake *ConfigTXValidator) ProposeConfigUpdate(arg1 *common.Envelope) (*common.ConfigEnvelope, error) {
	fake.proposeConfigUpdateMutex.Lock()
	ret, specificReturn := fake.proposeConfigUpdateReturnsOnCall[len(fake.proposeConfigUpdateArgsForCall)]
	fake.proposeConfigUpdateArgsForCall = append(fake.proposeConfigUpdateArgsForCall, struct {
		arg1 *common.Envelope
	}{arg1})
	stub := fake.ProposeConfigUpdateStub
	fakeReturns := fake.proposeConfigUpdateReturns
	fake.recordInvocation("ProposeConfigUpdate", []interface{}{arg1})
	fake.proposeConfigUpdateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ConfigTXValidator) ProposeConfigUpdateCallCount() int {
	fake.proposeConfigUpdateMutex.RLock()
	defer fake.proposeConfigUpdateMutex.RUnlock()
	return len(fake.proposeConfigUpdateArgsForCall)
}

func (fake *ConfigTXValidator) ProposeConfigUpdateCalls(stub func(*common.Envelope) (*common.ConfigEnvelope, error)) {
	fake.proposeConfigUpdateMutex.Lock()
	defer fake.proposeConfigUpdateMutex.Unlock()
	fake.ProposeConfigUpdateStub = stub
}

func (fake *ConfigTXValidator) ProposeConfigUpdateArgsForCall(i int) *common.Envelope {
	fake.proposeConfigUpdateMutex.RLock()
	defer fake.proposeConfigUpdateMutex.RUnlock()
	argsForCall := fake.proposeConfigUpdateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigTXValidator) ProposeConfigUpdateReturns(result1 *common.ConfigEnvelope, result2 error) {
	fake.proposeConfigUpdateMutex.Lock()
	defer fake.proposeConfigUpdateMutex.Unlock()
	fake.ProposeConfigUpdateStub = nil
	fake.proposeConfigUpdateReturns = struct {
		result1 *common.ConfigEnvelope
		result2 error
	}{result1, result2}
}

func (fake *ConfigTXValidator) ProposeConfigUpdateReturnsOnCall(i int, result1 *common.ConfigEnvelope, result2 error) {
	fake.proposeConfigUpdateMutex.Lock()
	defer fake.proposeConfigUpdateMutex.Unlock()
	fake.ProposeConfigUpdateStub = nil
	if fake.proposeConfigUpdateReturnsOnCall == nil {
		fake.proposeConfigUpdateReturnsOnCall = make(map[int]struct {
			result1 *common.ConfigEnvelope
			result2 error
		})
	}
	fake.proposeConfigUpdateReturnsOnCall[i] = struct {
		result1 *common.ConfigEnvelope
		result2 error
	}{result1, result2}
}

func (fake *ConfigTXValidator) Sequence() uint64 {
	fake.sequenceMutex.Lock()
	ret, specificReturn := fake.sequenceReturnsOnCall[len(fake.sequenceArgsForCall)]
	fake.sequenceArgsForCall = append(fake.sequenceArgsForCall, struct {
	}{})
	stub := fake.SequenceStub
	fakeReturns := fake.sequenceReturns
	fake.recordInvocation("Sequence", []interface{}{})
	fake.sequenceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigTXValidator) SequenceCallCount() int {
	fake.sequenceMutex.RLock()
	defer fake.sequenceMutex.RUnlock()
	return len(fake.sequenceArgsForCall)
}

func (fake *ConfigTXValidator) SequenceCalls(stub func() uint64) {
	fake.sequenceMutex.Lock()
	defer fake.sequenceMutex.Unlock()
	fake.SequenceStub = stub
}

func (fake *ConfigTXValidator) SequenceReturns(result1 uint64) {
	fake.sequenceMutex.Lock()
	defer fake.sequenceMutex.Unlock()
	fake.SequenceStub = nil
	fake.sequenceReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *ConfigTXValidator) SequenceReturnsOnCall(i int, result1 uint64) {
	fake.sequenceMutex.Lock()
	defer fake.sequenceMutex.Unlock()
	fake.SequenceStub = nil
	if fake.sequenceReturnsOnCall == nil {
		fake.sequenceReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.sequenceReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *ConfigTXValidator) Validate(arg1 *common.ConfigEnvelope) error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		arg1 *common.ConfigEnvelope
	}{arg1})
	stub := fake.ValidateStub
	fakeReturns := fake.validateReturns
	fake.recordInvocation("Validate", []interface{}{arg1})
	fake.validateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ConfigTXValidator) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *ConfigTXValidator) ValidateCalls(stub func(*common.ConfigEnvelope) error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *ConfigTXValidator) ValidateArgsForCall(i int) *common.ConfigEnvelope {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	argsForCall := fake.validateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigTXValidator) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigTXValidator) ValidateReturnsOnCall(i int, result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigTXValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.channelIDMutex.RLock()
	defer fake.channelIDMutex.RUnlock()
	fake.configProtoMutex.RLock()
	defer fake.configProtoMutex.RUnlock()
	fake.proposeConfigUpdateMutex.RLock()
	defer fake.proposeConfigUpdateMutex.RUnlock()
	fake.sequenceMutex.RLock()
	defer fake.sequenceMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigTXValidator) recordInvocation(key string, args []interface{}) {
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