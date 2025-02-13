// Code generated by mockery v2.52.2. DO NOT EDIT.

package apiv2connect

import (
	connect "connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IPServiceClient is an autogenerated mock type for the IPServiceClient type
type IPServiceClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *IPServiceClient) Create(_a0 context.Context, _a1 *connect.Request[apiv2.IPServiceCreateRequest]) (*connect.Response[apiv2.IPServiceCreateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *connect.Response[apiv2.IPServiceCreateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceCreateRequest]) (*connect.Response[apiv2.IPServiceCreateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceCreateRequest]) *connect.Response[apiv2.IPServiceCreateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.IPServiceCreateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.IPServiceCreateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *IPServiceClient) Delete(_a0 context.Context, _a1 *connect.Request[apiv2.IPServiceDeleteRequest]) (*connect.Response[apiv2.IPServiceDeleteResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *connect.Response[apiv2.IPServiceDeleteResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceDeleteRequest]) (*connect.Response[apiv2.IPServiceDeleteResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceDeleteRequest]) *connect.Response[apiv2.IPServiceDeleteResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.IPServiceDeleteResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.IPServiceDeleteRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *IPServiceClient) Get(_a0 context.Context, _a1 *connect.Request[apiv2.IPServiceGetRequest]) (*connect.Response[apiv2.IPServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv2.IPServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceGetRequest]) (*connect.Response[apiv2.IPServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceGetRequest]) *connect.Response[apiv2.IPServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.IPServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.IPServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *IPServiceClient) List(_a0 context.Context, _a1 *connect.Request[apiv2.IPServiceListRequest]) (*connect.Response[apiv2.IPServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[apiv2.IPServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceListRequest]) (*connect.Response[apiv2.IPServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceListRequest]) *connect.Response[apiv2.IPServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.IPServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.IPServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *IPServiceClient) Update(_a0 context.Context, _a1 *connect.Request[apiv2.IPServiceUpdateRequest]) (*connect.Response[apiv2.IPServiceUpdateResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *connect.Response[apiv2.IPServiceUpdateResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceUpdateRequest]) (*connect.Response[apiv2.IPServiceUpdateResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.IPServiceUpdateRequest]) *connect.Response[apiv2.IPServiceUpdateResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.IPServiceUpdateResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.IPServiceUpdateRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIPServiceClient creates a new instance of IPServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPServiceClient {
	mock := &IPServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
