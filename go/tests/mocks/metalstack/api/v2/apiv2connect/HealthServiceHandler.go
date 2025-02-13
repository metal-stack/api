// Code generated by mockery v2.52.2. DO NOT EDIT.

package apiv2connect

import (
	connect "connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// HealthServiceHandler is an autogenerated mock type for the HealthServiceHandler type
type HealthServiceHandler struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *HealthServiceHandler) Get(_a0 context.Context, _a1 *connect.Request[apiv2.HealthServiceGetRequest]) (*connect.Response[apiv2.HealthServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv2.HealthServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.HealthServiceGetRequest]) (*connect.Response[apiv2.HealthServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.HealthServiceGetRequest]) *connect.Response[apiv2.HealthServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.HealthServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.HealthServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHealthServiceHandler creates a new instance of HealthServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHealthServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *HealthServiceHandler {
	mock := &HealthServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
