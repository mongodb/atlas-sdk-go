// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20241113002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// PushBasedLogExportApi is an autogenerated mock type for the PushBasedLogExportApi type
type PushBasedLogExportApi struct {
	mock.Mock
}

type PushBasedLogExportApi_Expecter struct {
	mock *mock.Mock
}

func (_m *PushBasedLogExportApi) EXPECT() *PushBasedLogExportApi_Expecter {
	return &PushBasedLogExportApi_Expecter{mock: &_m.Mock}
}

// CreatePushBasedLogConfiguration provides a mock function with given fields: ctx, groupId, createPushBasedLogExportProjectRequest
func (_m *PushBasedLogExportApi) CreatePushBasedLogConfiguration(ctx context.Context, groupId string, createPushBasedLogExportProjectRequest *admin.CreatePushBasedLogExportProjectRequest) admin.CreatePushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, groupId, createPushBasedLogExportProjectRequest)

	if len(ret) == 0 {
		panic("no return value specified for CreatePushBasedLogConfiguration")
	}

	var r0 admin.CreatePushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.CreatePushBasedLogExportProjectRequest) admin.CreatePushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId, createPushBasedLogExportProjectRequest)
	} else {
		r0 = ret.Get(0).(admin.CreatePushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePushBasedLogConfiguration'
type PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call struct {
	*mock.Call
}

// CreatePushBasedLogConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - createPushBasedLogExportProjectRequest *admin.CreatePushBasedLogExportProjectRequest
func (_e *PushBasedLogExportApi_Expecter) CreatePushBasedLogConfiguration(ctx any, groupId any, createPushBasedLogExportProjectRequest any) *PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call {
	return &PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call{Call: _e.mock.On("CreatePushBasedLogConfiguration", ctx, groupId, createPushBasedLogExportProjectRequest)}
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call) Run(run func(ctx context.Context, groupId string, createPushBasedLogExportProjectRequest *admin.CreatePushBasedLogExportProjectRequest)) *PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.CreatePushBasedLogExportProjectRequest))
	})
	return _c
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call) Return(_a0 admin.CreatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call) RunAndReturn(run func(context.Context, string, *admin.CreatePushBasedLogExportProjectRequest) admin.CreatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_CreatePushBasedLogConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePushBasedLogConfigurationExecute provides a mock function with given fields: r
func (_m *PushBasedLogExportApi) CreatePushBasedLogConfigurationExecute(r admin.CreatePushBasedLogConfigurationApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreatePushBasedLogConfigurationExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.CreatePushBasedLogConfigurationApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreatePushBasedLogConfigurationApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreatePushBasedLogConfigurationApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePushBasedLogConfigurationExecute'
type PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call struct {
	*mock.Call
}

// CreatePushBasedLogConfigurationExecute is a helper method to define mock.On call
//   - r admin.CreatePushBasedLogConfigurationApiRequest
func (_e *PushBasedLogExportApi_Expecter) CreatePushBasedLogConfigurationExecute(r any) *PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call {
	return &PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call{Call: _e.mock.On("CreatePushBasedLogConfigurationExecute", r)}
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call) Run(run func(r admin.CreatePushBasedLogConfigurationApiRequest)) *PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreatePushBasedLogConfigurationApiRequest))
	})
	return _c
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call) Return(_a0 *http.Response, _a1 error) *PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call) RunAndReturn(run func(admin.CreatePushBasedLogConfigurationApiRequest) (*http.Response, error)) *PushBasedLogExportApi_CreatePushBasedLogConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePushBasedLogConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *PushBasedLogExportApi) CreatePushBasedLogConfigurationWithParams(ctx context.Context, args *admin.CreatePushBasedLogConfigurationApiParams) admin.CreatePushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreatePushBasedLogConfigurationWithParams")
	}

	var r0 admin.CreatePushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreatePushBasedLogConfigurationApiParams) admin.CreatePushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreatePushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePushBasedLogConfigurationWithParams'
type PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call struct {
	*mock.Call
}

// CreatePushBasedLogConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreatePushBasedLogConfigurationApiParams
func (_e *PushBasedLogExportApi_Expecter) CreatePushBasedLogConfigurationWithParams(ctx any, args any) *PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call {
	return &PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call{Call: _e.mock.On("CreatePushBasedLogConfigurationWithParams", ctx, args)}
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.CreatePushBasedLogConfigurationApiParams)) *PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreatePushBasedLogConfigurationApiParams))
	})
	return _c
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call) Return(_a0 admin.CreatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreatePushBasedLogConfigurationApiParams) admin.CreatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_CreatePushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePushBasedLogConfiguration provides a mock function with given fields: ctx, groupId
func (_m *PushBasedLogExportApi) DeletePushBasedLogConfiguration(ctx context.Context, groupId string) admin.DeletePushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for DeletePushBasedLogConfiguration")
	}

	var r0 admin.DeletePushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.DeletePushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.DeletePushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePushBasedLogConfiguration'
type PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call struct {
	*mock.Call
}

// DeletePushBasedLogConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *PushBasedLogExportApi_Expecter) DeletePushBasedLogConfiguration(ctx any, groupId any) *PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call {
	return &PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call{Call: _e.mock.On("DeletePushBasedLogConfiguration", ctx, groupId)}
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call) Run(run func(ctx context.Context, groupId string)) *PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call) Return(_a0 admin.DeletePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call) RunAndReturn(run func(context.Context, string) admin.DeletePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_DeletePushBasedLogConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePushBasedLogConfigurationExecute provides a mock function with given fields: r
func (_m *PushBasedLogExportApi) DeletePushBasedLogConfigurationExecute(r admin.DeletePushBasedLogConfigurationApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeletePushBasedLogConfigurationExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.DeletePushBasedLogConfigurationApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeletePushBasedLogConfigurationApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeletePushBasedLogConfigurationApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePushBasedLogConfigurationExecute'
type PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call struct {
	*mock.Call
}

// DeletePushBasedLogConfigurationExecute is a helper method to define mock.On call
//   - r admin.DeletePushBasedLogConfigurationApiRequest
func (_e *PushBasedLogExportApi_Expecter) DeletePushBasedLogConfigurationExecute(r any) *PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call {
	return &PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call{Call: _e.mock.On("DeletePushBasedLogConfigurationExecute", r)}
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call) Run(run func(r admin.DeletePushBasedLogConfigurationApiRequest)) *PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeletePushBasedLogConfigurationApiRequest))
	})
	return _c
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call) Return(_a0 *http.Response, _a1 error) *PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call) RunAndReturn(run func(admin.DeletePushBasedLogConfigurationApiRequest) (*http.Response, error)) *PushBasedLogExportApi_DeletePushBasedLogConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePushBasedLogConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *PushBasedLogExportApi) DeletePushBasedLogConfigurationWithParams(ctx context.Context, args *admin.DeletePushBasedLogConfigurationApiParams) admin.DeletePushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeletePushBasedLogConfigurationWithParams")
	}

	var r0 admin.DeletePushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeletePushBasedLogConfigurationApiParams) admin.DeletePushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeletePushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePushBasedLogConfigurationWithParams'
type PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call struct {
	*mock.Call
}

// DeletePushBasedLogConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeletePushBasedLogConfigurationApiParams
func (_e *PushBasedLogExportApi_Expecter) DeletePushBasedLogConfigurationWithParams(ctx any, args any) *PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call {
	return &PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call{Call: _e.mock.On("DeletePushBasedLogConfigurationWithParams", ctx, args)}
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.DeletePushBasedLogConfigurationApiParams)) *PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeletePushBasedLogConfigurationApiParams))
	})
	return _c
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call) Return(_a0 admin.DeletePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeletePushBasedLogConfigurationApiParams) admin.DeletePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_DeletePushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetPushBasedLogConfiguration provides a mock function with given fields: ctx, groupId
func (_m *PushBasedLogExportApi) GetPushBasedLogConfiguration(ctx context.Context, groupId string) admin.GetPushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for GetPushBasedLogConfiguration")
	}

	var r0 admin.GetPushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.GetPushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.GetPushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_GetPushBasedLogConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPushBasedLogConfiguration'
type PushBasedLogExportApi_GetPushBasedLogConfiguration_Call struct {
	*mock.Call
}

// GetPushBasedLogConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *PushBasedLogExportApi_Expecter) GetPushBasedLogConfiguration(ctx any, groupId any) *PushBasedLogExportApi_GetPushBasedLogConfiguration_Call {
	return &PushBasedLogExportApi_GetPushBasedLogConfiguration_Call{Call: _e.mock.On("GetPushBasedLogConfiguration", ctx, groupId)}
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfiguration_Call) Run(run func(ctx context.Context, groupId string)) *PushBasedLogExportApi_GetPushBasedLogConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfiguration_Call) Return(_a0 admin.GetPushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_GetPushBasedLogConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfiguration_Call) RunAndReturn(run func(context.Context, string) admin.GetPushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_GetPushBasedLogConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// GetPushBasedLogConfigurationExecute provides a mock function with given fields: r
func (_m *PushBasedLogExportApi) GetPushBasedLogConfigurationExecute(r admin.GetPushBasedLogConfigurationApiRequest) (*admin.PushBasedLogExportProject, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetPushBasedLogConfigurationExecute")
	}

	var r0 *admin.PushBasedLogExportProject
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetPushBasedLogConfigurationApiRequest) (*admin.PushBasedLogExportProject, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetPushBasedLogConfigurationApiRequest) *admin.PushBasedLogExportProject); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PushBasedLogExportProject)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetPushBasedLogConfigurationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetPushBasedLogConfigurationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPushBasedLogConfigurationExecute'
type PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call struct {
	*mock.Call
}

// GetPushBasedLogConfigurationExecute is a helper method to define mock.On call
//   - r admin.GetPushBasedLogConfigurationApiRequest
func (_e *PushBasedLogExportApi_Expecter) GetPushBasedLogConfigurationExecute(r any) *PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call {
	return &PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call{Call: _e.mock.On("GetPushBasedLogConfigurationExecute", r)}
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call) Run(run func(r admin.GetPushBasedLogConfigurationApiRequest)) *PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetPushBasedLogConfigurationApiRequest))
	})
	return _c
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call) Return(_a0 *admin.PushBasedLogExportProject, _a1 *http.Response, _a2 error) *PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call) RunAndReturn(run func(admin.GetPushBasedLogConfigurationApiRequest) (*admin.PushBasedLogExportProject, *http.Response, error)) *PushBasedLogExportApi_GetPushBasedLogConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetPushBasedLogConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *PushBasedLogExportApi) GetPushBasedLogConfigurationWithParams(ctx context.Context, args *admin.GetPushBasedLogConfigurationApiParams) admin.GetPushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetPushBasedLogConfigurationWithParams")
	}

	var r0 admin.GetPushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetPushBasedLogConfigurationApiParams) admin.GetPushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetPushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPushBasedLogConfigurationWithParams'
type PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call struct {
	*mock.Call
}

// GetPushBasedLogConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetPushBasedLogConfigurationApiParams
func (_e *PushBasedLogExportApi_Expecter) GetPushBasedLogConfigurationWithParams(ctx any, args any) *PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call {
	return &PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call{Call: _e.mock.On("GetPushBasedLogConfigurationWithParams", ctx, args)}
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.GetPushBasedLogConfigurationApiParams)) *PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetPushBasedLogConfigurationApiParams))
	})
	return _c
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call) Return(_a0 admin.GetPushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetPushBasedLogConfigurationApiParams) admin.GetPushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_GetPushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePushBasedLogConfiguration provides a mock function with given fields: ctx, groupId, pushBasedLogExportProject
func (_m *PushBasedLogExportApi) UpdatePushBasedLogConfiguration(ctx context.Context, groupId string, pushBasedLogExportProject *admin.PushBasedLogExportProject) admin.UpdatePushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, groupId, pushBasedLogExportProject)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePushBasedLogConfiguration")
	}

	var r0 admin.UpdatePushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *admin.PushBasedLogExportProject) admin.UpdatePushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, groupId, pushBasedLogExportProject)
	} else {
		r0 = ret.Get(0).(admin.UpdatePushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePushBasedLogConfiguration'
type PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call struct {
	*mock.Call
}

// UpdatePushBasedLogConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - pushBasedLogExportProject *admin.PushBasedLogExportProject
func (_e *PushBasedLogExportApi_Expecter) UpdatePushBasedLogConfiguration(ctx any, groupId any, pushBasedLogExportProject any) *PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call {
	return &PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call{Call: _e.mock.On("UpdatePushBasedLogConfiguration", ctx, groupId, pushBasedLogExportProject)}
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call) Run(run func(ctx context.Context, groupId string, pushBasedLogExportProject *admin.PushBasedLogExportProject)) *PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*admin.PushBasedLogExportProject))
	})
	return _c
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call) Return(_a0 admin.UpdatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call) RunAndReturn(run func(context.Context, string, *admin.PushBasedLogExportProject) admin.UpdatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_UpdatePushBasedLogConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePushBasedLogConfigurationExecute provides a mock function with given fields: r
func (_m *PushBasedLogExportApi) UpdatePushBasedLogConfigurationExecute(r admin.UpdatePushBasedLogConfigurationApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePushBasedLogConfigurationExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.UpdatePushBasedLogConfigurationApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.UpdatePushBasedLogConfigurationApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.UpdatePushBasedLogConfigurationApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePushBasedLogConfigurationExecute'
type PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call struct {
	*mock.Call
}

// UpdatePushBasedLogConfigurationExecute is a helper method to define mock.On call
//   - r admin.UpdatePushBasedLogConfigurationApiRequest
func (_e *PushBasedLogExportApi_Expecter) UpdatePushBasedLogConfigurationExecute(r any) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call {
	return &PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call{Call: _e.mock.On("UpdatePushBasedLogConfigurationExecute", r)}
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call) Run(run func(r admin.UpdatePushBasedLogConfigurationApiRequest)) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.UpdatePushBasedLogConfigurationApiRequest))
	})
	return _c
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call) Return(_a0 *http.Response, _a1 error) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call) RunAndReturn(run func(admin.UpdatePushBasedLogConfigurationApiRequest) (*http.Response, error)) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePushBasedLogConfigurationWithParams provides a mock function with given fields: ctx, args
func (_m *PushBasedLogExportApi) UpdatePushBasedLogConfigurationWithParams(ctx context.Context, args *admin.UpdatePushBasedLogConfigurationApiParams) admin.UpdatePushBasedLogConfigurationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePushBasedLogConfigurationWithParams")
	}

	var r0 admin.UpdatePushBasedLogConfigurationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.UpdatePushBasedLogConfigurationApiParams) admin.UpdatePushBasedLogConfigurationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.UpdatePushBasedLogConfigurationApiRequest)
	}

	return r0
}

// PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePushBasedLogConfigurationWithParams'
type PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call struct {
	*mock.Call
}

// UpdatePushBasedLogConfigurationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.UpdatePushBasedLogConfigurationApiParams
func (_e *PushBasedLogExportApi_Expecter) UpdatePushBasedLogConfigurationWithParams(ctx any, args any) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call {
	return &PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call{Call: _e.mock.On("UpdatePushBasedLogConfigurationWithParams", ctx, args)}
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call) Run(run func(ctx context.Context, args *admin.UpdatePushBasedLogConfigurationApiParams)) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.UpdatePushBasedLogConfigurationApiParams))
	})
	return _c
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call) Return(_a0 admin.UpdatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call) RunAndReturn(run func(context.Context, *admin.UpdatePushBasedLogConfigurationApiParams) admin.UpdatePushBasedLogConfigurationApiRequest) *PushBasedLogExportApi_UpdatePushBasedLogConfigurationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewPushBasedLogExportApi creates a new instance of PushBasedLogExportApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPushBasedLogExportApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *PushBasedLogExportApi {
	mock := &PushBasedLogExportApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
