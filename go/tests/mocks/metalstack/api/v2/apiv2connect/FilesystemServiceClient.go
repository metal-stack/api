// Code generated by mockery v2.52.3. DO NOT EDIT.

package apiv2connect

import (
	connect "connectrpc.com/connect"
	apiv2 "github.com/metal-stack/api/go/metalstack/api/v2"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FilesystemServiceClient is an autogenerated mock type for the FilesystemServiceClient type
type FilesystemServiceClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *FilesystemServiceClient) Get(_a0 context.Context, _a1 *connect.Request[apiv2.FilesystemServiceGetRequest]) (*connect.Response[apiv2.FilesystemServiceGetResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *connect.Response[apiv2.FilesystemServiceGetResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceGetRequest]) (*connect.Response[apiv2.FilesystemServiceGetResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceGetRequest]) *connect.Response[apiv2.FilesystemServiceGetResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.FilesystemServiceGetResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.FilesystemServiceGetRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: _a0, _a1
func (_m *FilesystemServiceClient) List(_a0 context.Context, _a1 *connect.Request[apiv2.FilesystemServiceListRequest]) (*connect.Response[apiv2.FilesystemServiceListResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 *connect.Response[apiv2.FilesystemServiceListResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceListRequest]) (*connect.Response[apiv2.FilesystemServiceListResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceListRequest]) *connect.Response[apiv2.FilesystemServiceListResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.FilesystemServiceListResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.FilesystemServiceListRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Match provides a mock function with given fields: _a0, _a1
func (_m *FilesystemServiceClient) Match(_a0 context.Context, _a1 *connect.Request[apiv2.FilesystemServiceMatchRequest]) (*connect.Response[apiv2.FilesystemServiceMatchResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Match")
	}

	var r0 *connect.Response[apiv2.FilesystemServiceMatchResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceMatchRequest]) (*connect.Response[apiv2.FilesystemServiceMatchResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceMatchRequest]) *connect.Response[apiv2.FilesystemServiceMatchResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.FilesystemServiceMatchResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.FilesystemServiceMatchRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Try provides a mock function with given fields: _a0, _a1
func (_m *FilesystemServiceClient) Try(_a0 context.Context, _a1 *connect.Request[apiv2.FilesystemServiceTryRequest]) (*connect.Response[apiv2.FilesystemServiceTryResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Try")
	}

	var r0 *connect.Response[apiv2.FilesystemServiceTryResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceTryRequest]) (*connect.Response[apiv2.FilesystemServiceTryResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[apiv2.FilesystemServiceTryRequest]) *connect.Response[apiv2.FilesystemServiceTryResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[apiv2.FilesystemServiceTryResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[apiv2.FilesystemServiceTryRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFilesystemServiceClient creates a new instance of FilesystemServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFilesystemServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *FilesystemServiceClient {
	mock := &FilesystemServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
