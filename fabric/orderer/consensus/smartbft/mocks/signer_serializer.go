// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SignerSerializer is an autogenerated mock type for the SignerSerializer type
type SignerSerializer struct {
	mock.Mock
}

type SignerSerializer_Expecter struct {
	mock *mock.Mock
}

func (_m *SignerSerializer) EXPECT() *SignerSerializer_Expecter {
	return &SignerSerializer_Expecter{mock: &_m.Mock}
}

// Serialize provides a mock function with given fields:
func (_m *SignerSerializer) Serialize() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Serialize")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignerSerializer_Serialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Serialize'
type SignerSerializer_Serialize_Call struct {
	*mock.Call
}

// Serialize is a helper method to define mock.On call
func (_e *SignerSerializer_Expecter) Serialize() *SignerSerializer_Serialize_Call {
	return &SignerSerializer_Serialize_Call{Call: _e.mock.On("Serialize")}
}

func (_c *SignerSerializer_Serialize_Call) Run(run func()) *SignerSerializer_Serialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SignerSerializer_Serialize_Call) Return(_a0 []byte, _a1 error) *SignerSerializer_Serialize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SignerSerializer_Serialize_Call) RunAndReturn(run func() ([]byte, error)) *SignerSerializer_Serialize_Call {
	_c.Call.Return(run)
	return _c
}

// Sign provides a mock function with given fields: message
func (_m *SignerSerializer) Sign(message []byte) ([]byte, error) {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for Sign")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(message)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignerSerializer_Sign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sign'
type SignerSerializer_Sign_Call struct {
	*mock.Call
}

// Sign is a helper method to define mock.On call
//   - message []byte
func (_e *SignerSerializer_Expecter) Sign(message interface{}) *SignerSerializer_Sign_Call {
	return &SignerSerializer_Sign_Call{Call: _e.mock.On("Sign", message)}
}

func (_c *SignerSerializer_Sign_Call) Run(run func(message []byte)) *SignerSerializer_Sign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *SignerSerializer_Sign_Call) Return(_a0 []byte, _a1 error) *SignerSerializer_Sign_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SignerSerializer_Sign_Call) RunAndReturn(run func([]byte) ([]byte, error)) *SignerSerializer_Sign_Call {
	_c.Call.Return(run)
	return _c
}

// NewSignerSerializer creates a new instance of SignerSerializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSignerSerializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *SignerSerializer {
	mock := &SignerSerializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
