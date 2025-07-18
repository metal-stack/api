// Code generated by mockery v2.53.4. DO NOT EDIT.

package client

import mock "github.com/stretchr/testify/mock"

// PersistTokenFn is an autogenerated mock type for the PersistTokenFn type
type PersistTokenFn struct {
	mock.Mock
}

// Execute provides a mock function with given fields: token
func (_m *PersistTokenFn) Execute(token string) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPersistTokenFn creates a new instance of PersistTokenFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPersistTokenFn(t interface {
	mock.TestingT
	Cleanup(func())
}) *PersistTokenFn {
	mock := &PersistTokenFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
