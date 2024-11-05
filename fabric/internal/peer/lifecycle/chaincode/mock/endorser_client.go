// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"context"
	"sync"

	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
	"google.golang.org/grpc"
)

type EndorserClient struct {
	ProcessProposalStub        func(context.Context, *peer.SignedProposal, ...grpc.CallOption) (*peer.ProposalResponse, error)
	processProposalMutex       sync.RWMutex
	processProposalArgsForCall []struct {
		arg1 context.Context
		arg2 *peer.SignedProposal
		arg3 []grpc.CallOption
	}
	processProposalReturns struct {
		result1 *peer.ProposalResponse
		result2 error
	}
	processProposalReturnsOnCall map[int]struct {
		result1 *peer.ProposalResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EndorserClient) ProcessProposal(arg1 context.Context, arg2 *peer.SignedProposal, arg3 ...grpc.CallOption) (*peer.ProposalResponse, error) {
	fake.processProposalMutex.Lock()
	ret, specificReturn := fake.processProposalReturnsOnCall[len(fake.processProposalArgsForCall)]
	fake.processProposalArgsForCall = append(fake.processProposalArgsForCall, struct {
		arg1 context.Context
		arg2 *peer.SignedProposal
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	fake.recordInvocation("ProcessProposal", []interface{}{arg1, arg2, arg3})
	fake.processProposalMutex.Unlock()
	if fake.ProcessProposalStub != nil {
		return fake.ProcessProposalStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.processProposalReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EndorserClient) ProcessProposalCallCount() int {
	fake.processProposalMutex.RLock()
	defer fake.processProposalMutex.RUnlock()
	return len(fake.processProposalArgsForCall)
}

func (fake *EndorserClient) ProcessProposalCalls(stub func(context.Context, *peer.SignedProposal, ...grpc.CallOption) (*peer.ProposalResponse, error)) {
	fake.processProposalMutex.Lock()
	defer fake.processProposalMutex.Unlock()
	fake.ProcessProposalStub = stub
}

func (fake *EndorserClient) ProcessProposalArgsForCall(i int) (context.Context, *peer.SignedProposal, []grpc.CallOption) {
	fake.processProposalMutex.RLock()
	defer fake.processProposalMutex.RUnlock()
	argsForCall := fake.processProposalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *EndorserClient) ProcessProposalReturns(result1 *peer.ProposalResponse, result2 error) {
	fake.processProposalMutex.Lock()
	defer fake.processProposalMutex.Unlock()
	fake.ProcessProposalStub = nil
	fake.processProposalReturns = struct {
		result1 *peer.ProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *EndorserClient) ProcessProposalReturnsOnCall(i int, result1 *peer.ProposalResponse, result2 error) {
	fake.processProposalMutex.Lock()
	defer fake.processProposalMutex.Unlock()
	fake.ProcessProposalStub = nil
	if fake.processProposalReturnsOnCall == nil {
		fake.processProposalReturnsOnCall = make(map[int]struct {
			result1 *peer.ProposalResponse
			result2 error
		})
	}
	fake.processProposalReturnsOnCall[i] = struct {
		result1 *peer.ProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *EndorserClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.processProposalMutex.RLock()
	defer fake.processProposalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EndorserClient) recordInvocation(key string, args []interface{}) {
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
