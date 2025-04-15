// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115014/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RootApi is an autogenerated mock type for the RootApi type
type RootApi struct {
	mock.Mock
}

type RootApi_Expecter struct {
	mock *mock.Mock
}

func (_m *RootApi) EXPECT() *RootApi_Expecter {
	return &RootApi_Expecter{mock: &_m.Mock}
}

// GetSystemStatus provides a mock function with given fields: ctx
func (_m *RootApi) GetSystemStatus(ctx context.Context) admin.GetSystemStatusApiRequest {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSystemStatus")
	}

	var r0 admin.GetSystemStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context) admin.GetSystemStatusApiRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(admin.GetSystemStatusApiRequest)
	}

	return r0
}

// RootApi_GetSystemStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSystemStatus'
type RootApi_GetSystemStatus_Call struct {
	*mock.Call
}

// GetSystemStatus is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RootApi_Expecter) GetSystemStatus(ctx any) *RootApi_GetSystemStatus_Call {
	return &RootApi_GetSystemStatus_Call{Call: _e.mock.On("GetSystemStatus", ctx)}
}

func (_c *RootApi_GetSystemStatus_Call) Run(run func(ctx context.Context)) *RootApi_GetSystemStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RootApi_GetSystemStatus_Call) Return(_a0 admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_GetSystemStatus_Call) RunAndReturn(run func(context.Context) admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetSystemStatusExecute provides a mock function with given fields: r
func (_m *RootApi) GetSystemStatusExecute(r admin.GetSystemStatusApiRequest) (*admin.SystemStatus, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetSystemStatusExecute")
	}

	var r0 *admin.SystemStatus
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetSystemStatusApiRequest) (*admin.SystemStatus, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetSystemStatusApiRequest) *admin.SystemStatus); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.SystemStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetSystemStatusApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetSystemStatusApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RootApi_GetSystemStatusExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSystemStatusExecute'
type RootApi_GetSystemStatusExecute_Call struct {
	*mock.Call
}

// GetSystemStatusExecute is a helper method to define mock.On call
//   - r admin.GetSystemStatusApiRequest
func (_e *RootApi_Expecter) GetSystemStatusExecute(r any) *RootApi_GetSystemStatusExecute_Call {
	return &RootApi_GetSystemStatusExecute_Call{Call: _e.mock.On("GetSystemStatusExecute", r)}
}

func (_c *RootApi_GetSystemStatusExecute_Call) Run(run func(r admin.GetSystemStatusApiRequest)) *RootApi_GetSystemStatusExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetSystemStatusApiRequest))
	})
	return _c
}

func (_c *RootApi_GetSystemStatusExecute_Call) Return(_a0 *admin.SystemStatus, _a1 *http.Response, _a2 error) *RootApi_GetSystemStatusExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RootApi_GetSystemStatusExecute_Call) RunAndReturn(run func(admin.GetSystemStatusApiRequest) (*admin.SystemStatus, *http.Response, error)) *RootApi_GetSystemStatusExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetSystemStatusWithParams provides a mock function with given fields: ctx, args
func (_m *RootApi) GetSystemStatusWithParams(ctx context.Context, args *admin.GetSystemStatusApiParams) admin.GetSystemStatusApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetSystemStatusWithParams")
	}

	var r0 admin.GetSystemStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetSystemStatusApiParams) admin.GetSystemStatusApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetSystemStatusApiRequest)
	}

	return r0
}

// RootApi_GetSystemStatusWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSystemStatusWithParams'
type RootApi_GetSystemStatusWithParams_Call struct {
	*mock.Call
}

// GetSystemStatusWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetSystemStatusApiParams
func (_e *RootApi_Expecter) GetSystemStatusWithParams(ctx any, args any) *RootApi_GetSystemStatusWithParams_Call {
	return &RootApi_GetSystemStatusWithParams_Call{Call: _e.mock.On("GetSystemStatusWithParams", ctx, args)}
}

func (_c *RootApi_GetSystemStatusWithParams_Call) Run(run func(ctx context.Context, args *admin.GetSystemStatusApiParams)) *RootApi_GetSystemStatusWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetSystemStatusApiParams))
	})
	return _c
}

func (_c *RootApi_GetSystemStatusWithParams_Call) Return(_a0 admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatusWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_GetSystemStatusWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetSystemStatusApiParams) admin.GetSystemStatusApiRequest) *RootApi_GetSystemStatusWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllControlPlaneIpAddresses provides a mock function with given fields: ctx
func (_m *RootApi) ReturnAllControlPlaneIpAddresses(ctx context.Context) admin.ReturnAllControlPlaneIpAddressesApiRequest {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllControlPlaneIpAddresses")
	}

	var r0 admin.ReturnAllControlPlaneIpAddressesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context) admin.ReturnAllControlPlaneIpAddressesApiRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(admin.ReturnAllControlPlaneIpAddressesApiRequest)
	}

	return r0
}

// RootApi_ReturnAllControlPlaneIpAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllControlPlaneIpAddresses'
type RootApi_ReturnAllControlPlaneIpAddresses_Call struct {
	*mock.Call
}

// ReturnAllControlPlaneIpAddresses is a helper method to define mock.On call
//   - ctx context.Context
func (_e *RootApi_Expecter) ReturnAllControlPlaneIpAddresses(ctx any) *RootApi_ReturnAllControlPlaneIpAddresses_Call {
	return &RootApi_ReturnAllControlPlaneIpAddresses_Call{Call: _e.mock.On("ReturnAllControlPlaneIpAddresses", ctx)}
}

func (_c *RootApi_ReturnAllControlPlaneIpAddresses_Call) Run(run func(ctx context.Context)) *RootApi_ReturnAllControlPlaneIpAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIpAddresses_Call) Return(_a0 admin.ReturnAllControlPlaneIpAddressesApiRequest) *RootApi_ReturnAllControlPlaneIpAddresses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIpAddresses_Call) RunAndReturn(run func(context.Context) admin.ReturnAllControlPlaneIpAddressesApiRequest) *RootApi_ReturnAllControlPlaneIpAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllControlPlaneIpAddressesExecute provides a mock function with given fields: r
func (_m *RootApi) ReturnAllControlPlaneIpAddressesExecute(r admin.ReturnAllControlPlaneIpAddressesApiRequest) (*admin.ControlPlaneIPAddresses, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllControlPlaneIpAddressesExecute")
	}

	var r0 *admin.ControlPlaneIPAddresses
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ReturnAllControlPlaneIpAddressesApiRequest) (*admin.ControlPlaneIPAddresses, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ReturnAllControlPlaneIpAddressesApiRequest) *admin.ControlPlaneIPAddresses); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ControlPlaneIPAddresses)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ReturnAllControlPlaneIpAddressesApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ReturnAllControlPlaneIpAddressesApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RootApi_ReturnAllControlPlaneIpAddressesExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllControlPlaneIpAddressesExecute'
type RootApi_ReturnAllControlPlaneIpAddressesExecute_Call struct {
	*mock.Call
}

// ReturnAllControlPlaneIpAddressesExecute is a helper method to define mock.On call
//   - r admin.ReturnAllControlPlaneIpAddressesApiRequest
func (_e *RootApi_Expecter) ReturnAllControlPlaneIpAddressesExecute(r any) *RootApi_ReturnAllControlPlaneIpAddressesExecute_Call {
	return &RootApi_ReturnAllControlPlaneIpAddressesExecute_Call{Call: _e.mock.On("ReturnAllControlPlaneIpAddressesExecute", r)}
}

func (_c *RootApi_ReturnAllControlPlaneIpAddressesExecute_Call) Run(run func(r admin.ReturnAllControlPlaneIpAddressesApiRequest)) *RootApi_ReturnAllControlPlaneIpAddressesExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ReturnAllControlPlaneIpAddressesApiRequest))
	})
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIpAddressesExecute_Call) Return(_a0 *admin.ControlPlaneIPAddresses, _a1 *http.Response, _a2 error) *RootApi_ReturnAllControlPlaneIpAddressesExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIpAddressesExecute_Call) RunAndReturn(run func(admin.ReturnAllControlPlaneIpAddressesApiRequest) (*admin.ControlPlaneIPAddresses, *http.Response, error)) *RootApi_ReturnAllControlPlaneIpAddressesExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllControlPlaneIpAddressesWithParams provides a mock function with given fields: ctx, args
func (_m *RootApi) ReturnAllControlPlaneIpAddressesWithParams(ctx context.Context, args *admin.ReturnAllControlPlaneIpAddressesApiParams) admin.ReturnAllControlPlaneIpAddressesApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllControlPlaneIpAddressesWithParams")
	}

	var r0 admin.ReturnAllControlPlaneIpAddressesApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ReturnAllControlPlaneIpAddressesApiParams) admin.ReturnAllControlPlaneIpAddressesApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ReturnAllControlPlaneIpAddressesApiRequest)
	}

	return r0
}

// RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllControlPlaneIpAddressesWithParams'
type RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call struct {
	*mock.Call
}

// ReturnAllControlPlaneIpAddressesWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ReturnAllControlPlaneIpAddressesApiParams
func (_e *RootApi_Expecter) ReturnAllControlPlaneIpAddressesWithParams(ctx any, args any) *RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call {
	return &RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call{Call: _e.mock.On("ReturnAllControlPlaneIpAddressesWithParams", ctx, args)}
}

func (_c *RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call) Run(run func(ctx context.Context, args *admin.ReturnAllControlPlaneIpAddressesApiParams)) *RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ReturnAllControlPlaneIpAddressesApiParams))
	})
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call) Return(_a0 admin.ReturnAllControlPlaneIpAddressesApiRequest) *RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call) RunAndReturn(run func(context.Context, *admin.ReturnAllControlPlaneIpAddressesApiParams) admin.ReturnAllControlPlaneIpAddressesApiRequest) *RootApi_ReturnAllControlPlaneIpAddressesWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewRootApi creates a new instance of RootApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRootApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *RootApi {
	mock := &RootApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
