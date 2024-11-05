// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	common "github.com/hyperledger/fabric-protos-go-apiv2/common"
	mock "github.com/stretchr/testify/mock"
)

// BlockRetriever is an autogenerated mock type for the BlockRetriever type
type BlockRetriever struct {
	mock.Mock
}

// Block provides a mock function with given fields: number
func (_m *BlockRetriever) Block(number uint64) *common.Block {
	ret := _m.Called(number)

	var r0 *common.Block
	if rf, ok := ret.Get(0).(func(uint64) *common.Block); ok {
		r0 = rf(number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Block)
		}
	}

	return r0
}

type mockConstructorTestingTNewBlockRetriever interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlockRetriever creates a new instance of BlockRetriever. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlockRetriever(t mockConstructorTestingTNewBlockRetriever) *BlockRetriever {
	mock := &BlockRetriever{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}