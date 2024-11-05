// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	orderer "github.com/hyperledger/fabric-protos-go-apiv2/orderer"
	mock "github.com/stretchr/testify/mock"
)

// RPC is an autogenerated mock type for the RPC type
type RPC struct {
	mock.Mock
}

// SendConsensus provides a mock function with given fields: dest, msg
func (_m *RPC) SendConsensus(dest uint64, msg *orderer.ConsensusRequest) error {
	ret := _m.Called(dest, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, *orderer.ConsensusRequest) error); ok {
		r0 = rf(dest, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendSubmit provides a mock function with given fields: destination, request, report
func (_m *RPC) SendSubmit(destination uint64, request *orderer.SubmitRequest, report func(error)) error {
	ret := _m.Called(destination, request, report)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, *orderer.SubmitRequest, func(error)) error); ok {
		r0 = rf(destination, request, report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRPC interface {
	mock.TestingT
	Cleanup(func())
}

// NewRPC creates a new instance of RPC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRPC(t mockConstructorTestingTNewRPC) *RPC {
	mock := &RPC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}