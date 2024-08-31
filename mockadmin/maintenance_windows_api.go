// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240805003/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MaintenanceWindowsApi is an autogenerated mock type for the MaintenanceWindowsApi type
type MaintenanceWindowsApi struct {
	mock.Mock
}

type MaintenanceWindowsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *MaintenanceWindowsApi) EXPECT() *MaintenanceWindowsApi_Expecter {
	return &MaintenanceWindowsApi_Expecter{mock: &_m.Mock}
}

// DeferMaintenanceWindow provides a mock function with given fields: ctx, groupId
func (_m *MaintenanceWindowsApi) DeferMaintenanceWindow(ctx context.Context, groupId string) admin.DeferMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for DeferMaintenanceWindow")
	}

	var r0 admin.DeferMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.DeferMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.DeferMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_DeferMaintenanceWindow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeferMaintenanceWindow'
type MaintenanceWindowsApi_DeferMaintenanceWindow_Call struct {
	*mock.Call
}

// DeferMaintenanceWindow is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *MaintenanceWindowsApi_Expecter) DeferMaintenanceWindow(ctx interface{}, groupId interface{}) *MaintenanceWindowsApi_DeferMaintenanceWindow_Call {
	return &MaintenanceWindowsApi_DeferMaintenanceWindow_Call{Call: _e.mock.On("DeferMaintenanceWindow", ctx, groupId)}
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindow_Call) Run(run func(ctx context.Context, groupId string)) *MaintenanceWindowsApi_DeferMaintenanceWindow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindow_Call) Return(_a0 admin.DeferMaintenanceWindowApiRequest) *MaintenanceWindowsApi_DeferMaintenanceWindow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindow_Call) RunAndReturn(run func(context.Context, string) admin.DeferMaintenanceWindowApiRequest) *MaintenanceWindowsApi_DeferMaintenanceWindow_Call {
	_c.Call.Return(run)
	return _c
}

// DeferMaintenanceWindowExecute provides a mock function with given fields: r
func (_m *MaintenanceWindowsApi) DeferMaintenanceWindowExecute(r admin.DeferMaintenanceWindowApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeferMaintenanceWindowExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.DeferMaintenanceWindowApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeferMaintenanceWindowApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeferMaintenanceWindowApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeferMaintenanceWindowExecute'
type MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call struct {
	*mock.Call
}

// DeferMaintenanceWindowExecute is a helper method to define mock.On call
//   - r admin.DeferMaintenanceWindowApiRequest
func (_e *MaintenanceWindowsApi_Expecter) DeferMaintenanceWindowExecute(r interface{}) *MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call {
	return &MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call{Call: _e.mock.On("DeferMaintenanceWindowExecute", r)}
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call) Run(run func(r admin.DeferMaintenanceWindowApiRequest)) *MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeferMaintenanceWindowApiRequest))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call) Return(_a0 *http.Response, _a1 error) *MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call) RunAndReturn(run func(admin.DeferMaintenanceWindowApiRequest) (*http.Response, error)) *MaintenanceWindowsApi_DeferMaintenanceWindowExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeferMaintenanceWindowWithParams provides a mock function with given fields: ctx, args
func (_m *MaintenanceWindowsApi) DeferMaintenanceWindowWithParams(ctx context.Context, args *admin.DeferMaintenanceWindowApiParams) admin.DeferMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeferMaintenanceWindowWithParams")
	}

	var r0 admin.DeferMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeferMaintenanceWindowApiParams) admin.DeferMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeferMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeferMaintenanceWindowWithParams'
type MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call struct {
	*mock.Call
}

// DeferMaintenanceWindowWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeferMaintenanceWindowApiParams
func (_e *MaintenanceWindowsApi_Expecter) DeferMaintenanceWindowWithParams(ctx interface{}, args interface{}) *MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call {
	return &MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call{Call: _e.mock.On("DeferMaintenanceWindowWithParams", ctx, args)}
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call) Run(run func(ctx context.Context, args *admin.DeferMaintenanceWindowApiParams)) *MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeferMaintenanceWindowApiParams))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call) Return(_a0 admin.DeferMaintenanceWindowApiRequest) *MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeferMaintenanceWindowApiParams) admin.DeferMaintenanceWindowApiRequest) *MaintenanceWindowsApi_DeferMaintenanceWindowWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetMaintenanceWindow provides a mock function with given fields: ctx, groupId
func (_m *MaintenanceWindowsApi) GetMaintenanceWindow(ctx context.Context, groupId string) admin.GetMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetMaintenanceWindow")
	}

	var r0 admin.GetMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.GetMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_GetMaintenanceWindow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMaintenanceWindow'
type MaintenanceWindowsApi_GetMaintenanceWindow_Call struct {
	*mock.Call
}

// GetMaintenanceWindow is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *MaintenanceWindowsApi_Expecter) GetMaintenanceWindow(ctx interface{}, groupId interface{}) *MaintenanceWindowsApi_GetMaintenanceWindow_Call {
	return &MaintenanceWindowsApi_GetMaintenanceWindow_Call{Call: _e.mock.On("GetMaintenanceWindow", ctx, groupId)}
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindow_Call) Run(run func(ctx context.Context, groupId string)) *MaintenanceWindowsApi_GetMaintenanceWindow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindow_Call) Return(_a0 admin.GetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_GetMaintenanceWindow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindow_Call) RunAndReturn(run func(context.Context, string) admin.GetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_GetMaintenanceWindow_Call {
	_c.Call.Return(run)
	return _c
}

// GetMaintenanceWindowExecute provides a mock function with given fields: r
func (_m *MaintenanceWindowsApi) GetMaintenanceWindowExecute(r admin.GetMaintenanceWindowApiRequest) (*admin.GroupMaintenanceWindow, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetMaintenanceWindowExecute")
	}

	var r0 *admin.GroupMaintenanceWindow
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetMaintenanceWindowApiRequest) (*admin.GroupMaintenanceWindow, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetMaintenanceWindowApiRequest) *admin.GroupMaintenanceWindow); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GroupMaintenanceWindow)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetMaintenanceWindowApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetMaintenanceWindowApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMaintenanceWindowExecute'
type MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call struct {
	*mock.Call
}

// GetMaintenanceWindowExecute is a helper method to define mock.On call
//   - r admin.GetMaintenanceWindowApiRequest
func (_e *MaintenanceWindowsApi_Expecter) GetMaintenanceWindowExecute(r interface{}) *MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call {
	return &MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call{Call: _e.mock.On("GetMaintenanceWindowExecute", r)}
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call) Run(run func(r admin.GetMaintenanceWindowApiRequest)) *MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetMaintenanceWindowApiRequest))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call) Return(_a0 *admin.GroupMaintenanceWindow, _a1 *http.Response, _a2 error) *MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call) RunAndReturn(run func(admin.GetMaintenanceWindowApiRequest) (*admin.GroupMaintenanceWindow, *http.Response, error)) *MaintenanceWindowsApi_GetMaintenanceWindowExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetMaintenanceWindowWithParams provides a mock function with given fields: ctx, args
func (_m *MaintenanceWindowsApi) GetMaintenanceWindowWithParams(ctx context.Context, args *admin.GetMaintenanceWindowApiParams) admin.GetMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetMaintenanceWindowWithParams")
	}

	var r0 admin.GetMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetMaintenanceWindowApiParams) admin.GetMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMaintenanceWindowWithParams'
type MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call struct {
	*mock.Call
}

// GetMaintenanceWindowWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetMaintenanceWindowApiParams
func (_e *MaintenanceWindowsApi_Expecter) GetMaintenanceWindowWithParams(ctx interface{}, args interface{}) *MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call {
	return &MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call{Call: _e.mock.On("GetMaintenanceWindowWithParams", ctx, args)}
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call) Run(run func(ctx context.Context, args *admin.GetMaintenanceWindowApiParams)) *MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetMaintenanceWindowApiParams))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call) Return(_a0 admin.GetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetMaintenanceWindowApiParams) admin.GetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_GetMaintenanceWindowWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ResetMaintenanceWindow provides a mock function with given fields: ctx, groupId
func (_m *MaintenanceWindowsApi) ResetMaintenanceWindow(ctx context.Context, groupId string) admin.ResetMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ResetMaintenanceWindow")
	}

	var r0 admin.ResetMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ResetMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ResetMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_ResetMaintenanceWindow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetMaintenanceWindow'
type MaintenanceWindowsApi_ResetMaintenanceWindow_Call struct {
	*mock.Call
}

// ResetMaintenanceWindow is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *MaintenanceWindowsApi_Expecter) ResetMaintenanceWindow(ctx interface{}, groupId interface{}) *MaintenanceWindowsApi_ResetMaintenanceWindow_Call {
	return &MaintenanceWindowsApi_ResetMaintenanceWindow_Call{Call: _e.mock.On("ResetMaintenanceWindow", ctx, groupId)}
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindow_Call) Run(run func(ctx context.Context, groupId string)) *MaintenanceWindowsApi_ResetMaintenanceWindow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindow_Call) Return(_a0 admin.ResetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_ResetMaintenanceWindow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindow_Call) RunAndReturn(run func(context.Context, string) admin.ResetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_ResetMaintenanceWindow_Call {
	_c.Call.Return(run)
	return _c
}

// ResetMaintenanceWindowExecute provides a mock function with given fields: r
func (_m *MaintenanceWindowsApi) ResetMaintenanceWindowExecute(r admin.ResetMaintenanceWindowApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ResetMaintenanceWindowExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.ResetMaintenanceWindowApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ResetMaintenanceWindowApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ResetMaintenanceWindowApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetMaintenanceWindowExecute'
type MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call struct {
	*mock.Call
}

// ResetMaintenanceWindowExecute is a helper method to define mock.On call
//   - r admin.ResetMaintenanceWindowApiRequest
func (_e *MaintenanceWindowsApi_Expecter) ResetMaintenanceWindowExecute(r interface{}) *MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call {
	return &MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call{Call: _e.mock.On("ResetMaintenanceWindowExecute", r)}
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call) Run(run func(r admin.ResetMaintenanceWindowApiRequest)) *MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ResetMaintenanceWindowApiRequest))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call) Return(_a0 *http.Response, _a1 error) *MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call) RunAndReturn(run func(admin.ResetMaintenanceWindowApiRequest) (*http.Response, error)) *MaintenanceWindowsApi_ResetMaintenanceWindowExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ResetMaintenanceWindowWithParams provides a mock function with given fields: ctx, args
func (_m *MaintenanceWindowsApi) ResetMaintenanceWindowWithParams(ctx context.Context, args *admin.ResetMaintenanceWindowApiParams) admin.ResetMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ResetMaintenanceWindowWithParams")
	}

	var r0 admin.ResetMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ResetMaintenanceWindowApiParams) admin.ResetMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ResetMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetMaintenanceWindowWithParams'
type MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call struct {
	*mock.Call
}

// ResetMaintenanceWindowWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ResetMaintenanceWindowApiParams
func (_e *MaintenanceWindowsApi_Expecter) ResetMaintenanceWindowWithParams(ctx interface{}, args interface{}) *MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call {
	return &MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call{Call: _e.mock.On("ResetMaintenanceWindowWithParams", ctx, args)}
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call) Run(run func(ctx context.Context, args *admin.ResetMaintenanceWindowApiParams)) *MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ResetMaintenanceWindowApiParams))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call) Return(_a0 admin.ResetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call) RunAndReturn(run func(context.Context, *admin.ResetMaintenanceWindowApiParams) admin.ResetMaintenanceWindowApiRequest) *MaintenanceWindowsApi_ResetMaintenanceWindowWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleMaintenanceAutoDefer provides a mock function with given fields: ctx, groupId
func (_m *MaintenanceWindowsApi) ToggleMaintenanceAutoDefer(ctx context.Context, groupId string) admin.ToggleMaintenanceAutoDeferApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ToggleMaintenanceAutoDefer")
	}

	var r0 admin.ToggleMaintenanceAutoDeferApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ToggleMaintenanceAutoDeferApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ToggleMaintenanceAutoDeferApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleMaintenanceAutoDefer'
type MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call struct {
	*mock.Call
}

// ToggleMaintenanceAutoDefer is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *MaintenanceWindowsApi_Expecter) ToggleMaintenanceAutoDefer(ctx interface{}, groupId interface{}) *MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call {
	return &MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call{Call: _e.mock.On("ToggleMaintenanceAutoDefer", ctx, groupId)}
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call) Run(run func(ctx context.Context, groupId string)) *MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call) Return(_a0 admin.ToggleMaintenanceAutoDeferApiRequest) *MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call) RunAndReturn(run func(context.Context, string) admin.ToggleMaintenanceAutoDeferApiRequest) *MaintenanceWindowsApi_ToggleMaintenanceAutoDefer_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleMaintenanceAutoDeferExecute provides a mock function with given fields: r
func (_m *MaintenanceWindowsApi) ToggleMaintenanceAutoDeferExecute(r admin.ToggleMaintenanceAutoDeferApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ToggleMaintenanceAutoDeferExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.ToggleMaintenanceAutoDeferApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ToggleMaintenanceAutoDeferApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ToggleMaintenanceAutoDeferApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleMaintenanceAutoDeferExecute'
type MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call struct {
	*mock.Call
}

// ToggleMaintenanceAutoDeferExecute is a helper method to define mock.On call
//   - r admin.ToggleMaintenanceAutoDeferApiRequest
func (_e *MaintenanceWindowsApi_Expecter) ToggleMaintenanceAutoDeferExecute(r interface{}) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call {
	return &MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call{Call: _e.mock.On("ToggleMaintenanceAutoDeferExecute", r)}
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call) Run(run func(r admin.ToggleMaintenanceAutoDeferApiRequest)) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ToggleMaintenanceAutoDeferApiRequest))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call) Return(_a0 *http.Response, _a1 error) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call) RunAndReturn(run func(admin.ToggleMaintenanceAutoDeferApiRequest) (*http.Response, error)) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleMaintenanceAutoDeferWithParams provides a mock function with given fields: ctx, args
func (_m *MaintenanceWindowsApi) ToggleMaintenanceAutoDeferWithParams(ctx context.Context, args *admin.ToggleMaintenanceAutoDeferApiParams) admin.ToggleMaintenanceAutoDeferApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ToggleMaintenanceAutoDeferWithParams")
	}

	var r0 admin.ToggleMaintenanceAutoDeferApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ToggleMaintenanceAutoDeferApiParams) admin.ToggleMaintenanceAutoDeferApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ToggleMaintenanceAutoDeferApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleMaintenanceAutoDeferWithParams'
type MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call struct {
	*mock.Call
}

// ToggleMaintenanceAutoDeferWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ToggleMaintenanceAutoDeferApiParams
func (_e *MaintenanceWindowsApi_Expecter) ToggleMaintenanceAutoDeferWithParams(ctx interface{}, args interface{}) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call {
	return &MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call{Call: _e.mock.On("ToggleMaintenanceAutoDeferWithParams", ctx, args)}
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call) Run(run func(ctx context.Context, args *admin.ToggleMaintenanceAutoDeferApiParams)) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ToggleMaintenanceAutoDeferApiParams))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call) Return(_a0 admin.ToggleMaintenanceAutoDeferApiRequest) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call) RunAndReturn(run func(context.Context, *admin.ToggleMaintenanceAutoDeferApiParams) admin.ToggleMaintenanceAutoDeferApiRequest) *MaintenanceWindowsApi_ToggleMaintenanceAutoDeferWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMaintenanceWindow provides a mock function with given fields: ctx, groupId, groupMaintenanceWindow
func (_m *MaintenanceWindowsApi) UpdateMaintenanceWindow(ctx context.Context, groupId string, groupMaintenanceWindow *admin.GroupMaintenanceWindow) admin.UpdateMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, groupId, groupMaintenanceWindow)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMaintenanceWindow")
	}

	var r0 admin.UpdateMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.GroupMaintenanceWindow) admin.UpdateMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, groupId, groupMaintenanceWindow)
	} else {
		r0 = ret.Get(0).(admin.UpdateMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_UpdateMaintenanceWindow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMaintenanceWindow'
type MaintenanceWindowsApi_UpdateMaintenanceWindow_Call struct {
	*mock.Call
}

// UpdateMaintenanceWindow is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - groupMaintenanceWindow *admin.GroupMaintenanceWindow
func (_e *MaintenanceWindowsApi_Expecter) UpdateMaintenanceWindow(ctx interface{}, groupId interface{}, groupMaintenanceWindow interface{}) *MaintenanceWindowsApi_UpdateMaintenanceWindow_Call {
	return &MaintenanceWindowsApi_UpdateMaintenanceWindow_Call{Call: _e.mock.On("UpdateMaintenanceWindow", ctx, groupId, groupMaintenanceWindow)}
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindow_Call) Run(run func(ctx context.Context, groupId string, groupMaintenanceWindow *admin.GroupMaintenanceWindow)) *MaintenanceWindowsApi_UpdateMaintenanceWindow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.GroupMaintenanceWindow))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindow_Call) Return(_a0 admin.UpdateMaintenanceWindowApiRequest) *MaintenanceWindowsApi_UpdateMaintenanceWindow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindow_Call) RunAndReturn(run func(context.Context, string, *admin.GroupMaintenanceWindow) admin.UpdateMaintenanceWindowApiRequest) *MaintenanceWindowsApi_UpdateMaintenanceWindow_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMaintenanceWindowExecute provides a mock function with given fields: r
func (_m *MaintenanceWindowsApi) UpdateMaintenanceWindowExecute(r admin.UpdateMaintenanceWindowApiRequest) (interface{}, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMaintenanceWindowExecute")
	}

	var r0 interface{}
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.UpdateMaintenanceWindowApiRequest) (interface{}, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.UpdateMaintenanceWindowApiRequest) interface{}); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(admin.UpdateMaintenanceWindowApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.UpdateMaintenanceWindowApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMaintenanceWindowExecute'
type MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call struct {
	*mock.Call
}

// UpdateMaintenanceWindowExecute is a helper method to define mock.On call
//   - r admin.UpdateMaintenanceWindowApiRequest
func (_e *MaintenanceWindowsApi_Expecter) UpdateMaintenanceWindowExecute(r interface{}) *MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call {
	return &MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call{Call: _e.mock.On("UpdateMaintenanceWindowExecute", r)}
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call) Run(run func(r admin.UpdateMaintenanceWindowApiRequest)) *MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.UpdateMaintenanceWindowApiRequest))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call) Return(_a0 interface{}, _a1 *http.Response, _a2 error) *MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call) RunAndReturn(run func(admin.UpdateMaintenanceWindowApiRequest) (interface{}, *http.Response, error)) *MaintenanceWindowsApi_UpdateMaintenanceWindowExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMaintenanceWindowWithParams provides a mock function with given fields: ctx, args
func (_m *MaintenanceWindowsApi) UpdateMaintenanceWindowWithParams(ctx context.Context, args *admin.UpdateMaintenanceWindowApiParams) admin.UpdateMaintenanceWindowApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMaintenanceWindowWithParams")
	}

	var r0 admin.UpdateMaintenanceWindowApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.UpdateMaintenanceWindowApiParams) admin.UpdateMaintenanceWindowApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.UpdateMaintenanceWindowApiRequest)
	}

	return r0
}

// MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMaintenanceWindowWithParams'
type MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call struct {
	*mock.Call
}

// UpdateMaintenanceWindowWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.UpdateMaintenanceWindowApiParams
func (_e *MaintenanceWindowsApi_Expecter) UpdateMaintenanceWindowWithParams(ctx interface{}, args interface{}) *MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call {
	return &MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call{Call: _e.mock.On("UpdateMaintenanceWindowWithParams", ctx, args)}
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call) Run(run func(ctx context.Context, args *admin.UpdateMaintenanceWindowApiParams)) *MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.UpdateMaintenanceWindowApiParams))
	})
	return _c
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call) Return(_a0 admin.UpdateMaintenanceWindowApiRequest) *MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call) RunAndReturn(run func(context.Context, *admin.UpdateMaintenanceWindowApiParams) admin.UpdateMaintenanceWindowApiRequest) *MaintenanceWindowsApi_UpdateMaintenanceWindowWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewMaintenanceWindowsApi creates a new instance of MaintenanceWindowsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMaintenanceWindowsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *MaintenanceWindowsApi {
	mock := &MaintenanceWindowsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
