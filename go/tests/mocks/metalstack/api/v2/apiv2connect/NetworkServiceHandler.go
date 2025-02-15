// Code generated by mockery v2.52.2. DO NOT EDIT.

package apiv2connect

import (
	connect "connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// NetworkServiceHandler is an autogenerated mock type for the NetworkServiceHandler type
type NetworkServiceHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *NetworkServiceHandler) Create(_a0 context.Context, _a1 *connect.Request[apiv2.NetworkServiceCreateRequest]) (*connect.Response[apiv2.NetworkServiceCreateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *connect.Response[apiv2.NetworkServiceCreateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceCreateRequest]) (*connect.Response[apiv2.NetworkServiceCreateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceCreateRequest]) *connect.Response[apiv2.NetworkServiceCreateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.NetworkServiceCreateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.NetworkServiceCreateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *NetworkServiceHandler) Delete(_a0 context.Context, _a1 *connect.Request[apiv2.NetworkServiceDeleteRequest]) (*connect.Response[apiv2.NetworkServiceDeleteResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *connect.Response[apiv2.NetworkServiceDeleteResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceDeleteRequest]) (*connect.Response[apiv2.NetworkServiceDeleteResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceDeleteRequest]) *connect.Response[apiv2.NetworkServiceDeleteResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.NetworkServiceDeleteResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.NetworkServiceDeleteRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *NetworkServiceHandler) Get(_a0 context.Context, _a1 *connect.Request[apiv2.NetworkServiceGetRequest]) (*connect.Response[apiv2.NetworkServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv2.NetworkServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceGetRequest]) (*connect.Response[apiv2.NetworkServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceGetRequest]) *connect.Response[apiv2.NetworkServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.NetworkServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.NetworkServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *NetworkServiceHandler) List(_a0 context.Context, _a1 *connect.Request[apiv2.NetworkServiceListRequest]) (*connect.Response[apiv2.NetworkServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[apiv2.NetworkServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceListRequest]) (*connect.Response[apiv2.NetworkServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceListRequest]) *connect.Response[apiv2.NetworkServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.NetworkServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.NetworkServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *NetworkServiceHandler) Update(_a0 context.Context, _a1 *connect.Request[apiv2.NetworkServiceUpdateRequest]) (*connect.Response[apiv2.NetworkServiceUpdateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *connect.Response[apiv2.NetworkServiceUpdateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceUpdateRequest]) (*connect.Response[apiv2.NetworkServiceUpdateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.NetworkServiceUpdateRequest]) *connect.Response[apiv2.NetworkServiceUpdateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.NetworkServiceUpdateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.NetworkServiceUpdateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNetworkServiceHandler creates a new instance of NetworkServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNetworkServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *NetworkServiceHandler {
	mock := &NetworkServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
