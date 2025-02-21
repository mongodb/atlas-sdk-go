// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20250219001/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// FlexRestoreJobsApi is an autogenerated mock type for the FlexRestoreJobsApi type
type FlexRestoreJobsApi struct {
	mock.Mock
}

type FlexRestoreJobsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *FlexRestoreJobsApi) EXPECT() *FlexRestoreJobsApi_Expecter {
	return &FlexRestoreJobsApi_Expecter{mock: &_m.Mock}
}

// CreateFlexBackupRestoreJob provides a mock function with given fields: ctx, groupId, name, flexBackupRestoreJobCreate20241113
func (_m *FlexRestoreJobsApi) CreateFlexBackupRestoreJob(ctx context.Context, groupId string, name string, flexBackupRestoreJobCreate20241113 *admin.FlexBackupRestoreJobCreate20241113) admin.CreateFlexBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, groupId, name, flexBackupRestoreJobCreate20241113)

	if len(ret) == 0 {
		panic("no return value specified for CreateFlexBackupRestoreJob")
	}

	var r0 admin.CreateFlexBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.FlexBackupRestoreJobCreate20241113) admin.CreateFlexBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, groupId, name, flexBackupRestoreJobCreate20241113)
	} else {
		r0 = ret.Get(0).(admin.CreateFlexBackupRestoreJobApiRequest)
	}

	return r0
}

// FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateFlexBackupRestoreJob'
type FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call struct {
	*mock.Call
}

// CreateFlexBackupRestoreJob is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - name string
//   - flexBackupRestoreJobCreate20241113 *admin.FlexBackupRestoreJobCreate20241113
func (_e *FlexRestoreJobsApi_Expecter) CreateFlexBackupRestoreJob(ctx any, groupId any, name any, flexBackupRestoreJobCreate20241113 any) *FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call {
	return &FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call{Call: _e.mock.On("CreateFlexBackupRestoreJob", ctx, groupId, name, flexBackupRestoreJobCreate20241113)}
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call) Run(run func(ctx context.Context, groupId string, name string, flexBackupRestoreJobCreate20241113 *admin.FlexBackupRestoreJobCreate20241113)) *FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.FlexBackupRestoreJobCreate20241113))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call) Return(_a0 admin.CreateFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call) RunAndReturn(run func(context.Context, string, string, *admin.FlexBackupRestoreJobCreate20241113) admin.CreateFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_CreateFlexBackupRestoreJob_Call {
	_c.Call.Return(run)
	return _c
}

// CreateFlexBackupRestoreJobExecute provides a mock function with given fields: r
func (_m *FlexRestoreJobsApi) CreateFlexBackupRestoreJobExecute(r admin.CreateFlexBackupRestoreJobApiRequest) (*admin.FlexBackupRestoreJob20241113, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateFlexBackupRestoreJobExecute")
	}

	var r0 *admin.FlexBackupRestoreJob20241113
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateFlexBackupRestoreJobApiRequest) (*admin.FlexBackupRestoreJob20241113, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateFlexBackupRestoreJobApiRequest) *admin.FlexBackupRestoreJob20241113); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.FlexBackupRestoreJob20241113)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateFlexBackupRestoreJobApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateFlexBackupRestoreJobApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateFlexBackupRestoreJobExecute'
type FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call struct {
	*mock.Call
}

// CreateFlexBackupRestoreJobExecute is a helper method to define mock.On call
//   - r admin.CreateFlexBackupRestoreJobApiRequest
func (_e *FlexRestoreJobsApi_Expecter) CreateFlexBackupRestoreJobExecute(r any) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call {
	return &FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call{Call: _e.mock.On("CreateFlexBackupRestoreJobExecute", r)}
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call) Run(run func(r admin.CreateFlexBackupRestoreJobApiRequest)) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateFlexBackupRestoreJobApiRequest))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call) Return(_a0 *admin.FlexBackupRestoreJob20241113, _a1 *http.Response, _a2 error) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call) RunAndReturn(run func(admin.CreateFlexBackupRestoreJobApiRequest) (*admin.FlexBackupRestoreJob20241113, *http.Response, error)) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateFlexBackupRestoreJobWithParams provides a mock function with given fields: ctx, args
func (_m *FlexRestoreJobsApi) CreateFlexBackupRestoreJobWithParams(ctx context.Context, args *admin.CreateFlexBackupRestoreJobApiParams) admin.CreateFlexBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateFlexBackupRestoreJobWithParams")
	}

	var r0 admin.CreateFlexBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateFlexBackupRestoreJobApiParams) admin.CreateFlexBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateFlexBackupRestoreJobApiRequest)
	}

	return r0
}

// FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateFlexBackupRestoreJobWithParams'
type FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call struct {
	*mock.Call
}

// CreateFlexBackupRestoreJobWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateFlexBackupRestoreJobApiParams
func (_e *FlexRestoreJobsApi_Expecter) CreateFlexBackupRestoreJobWithParams(ctx any, args any) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call {
	return &FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call{Call: _e.mock.On("CreateFlexBackupRestoreJobWithParams", ctx, args)}
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateFlexBackupRestoreJobApiParams)) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateFlexBackupRestoreJobApiParams))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call) Return(_a0 admin.CreateFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateFlexBackupRestoreJobApiParams) admin.CreateFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_CreateFlexBackupRestoreJobWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetFlexBackupRestoreJob provides a mock function with given fields: ctx, groupId, name, restoreJobId
func (_m *FlexRestoreJobsApi) GetFlexBackupRestoreJob(ctx context.Context, groupId string, name string, restoreJobId string) admin.GetFlexBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, groupId, name, restoreJobId)

	if len(ret) == 0 {
		panic("no return value specified for GetFlexBackupRestoreJob")
	}

	var r0 admin.GetFlexBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) admin.GetFlexBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, groupId, name, restoreJobId)
	} else {
		r0 = ret.Get(0).(admin.GetFlexBackupRestoreJobApiRequest)
	}

	return r0
}

// FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlexBackupRestoreJob'
type FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call struct {
	*mock.Call
}

// GetFlexBackupRestoreJob is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - name string
//   - restoreJobId string
func (_e *FlexRestoreJobsApi_Expecter) GetFlexBackupRestoreJob(ctx any, groupId any, name any, restoreJobId any) *FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call {
	return &FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call{Call: _e.mock.On("GetFlexBackupRestoreJob", ctx, groupId, name, restoreJobId)}
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call) Run(run func(ctx context.Context, groupId string, name string, restoreJobId string)) *FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call) Return(_a0 admin.GetFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call) RunAndReturn(run func(context.Context, string, string, string) admin.GetFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_GetFlexBackupRestoreJob_Call {
	_c.Call.Return(run)
	return _c
}

// GetFlexBackupRestoreJobExecute provides a mock function with given fields: r
func (_m *FlexRestoreJobsApi) GetFlexBackupRestoreJobExecute(r admin.GetFlexBackupRestoreJobApiRequest) (*admin.FlexBackupRestoreJob20241113, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetFlexBackupRestoreJobExecute")
	}

	var r0 *admin.FlexBackupRestoreJob20241113
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetFlexBackupRestoreJobApiRequest) (*admin.FlexBackupRestoreJob20241113, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetFlexBackupRestoreJobApiRequest) *admin.FlexBackupRestoreJob20241113); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.FlexBackupRestoreJob20241113)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetFlexBackupRestoreJobApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetFlexBackupRestoreJobApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlexBackupRestoreJobExecute'
type FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call struct {
	*mock.Call
}

// GetFlexBackupRestoreJobExecute is a helper method to define mock.On call
//   - r admin.GetFlexBackupRestoreJobApiRequest
func (_e *FlexRestoreJobsApi_Expecter) GetFlexBackupRestoreJobExecute(r any) *FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call {
	return &FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call{Call: _e.mock.On("GetFlexBackupRestoreJobExecute", r)}
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call) Run(run func(r admin.GetFlexBackupRestoreJobApiRequest)) *FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetFlexBackupRestoreJobApiRequest))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call) Return(_a0 *admin.FlexBackupRestoreJob20241113, _a1 *http.Response, _a2 error) *FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call) RunAndReturn(run func(admin.GetFlexBackupRestoreJobApiRequest) (*admin.FlexBackupRestoreJob20241113, *http.Response, error)) *FlexRestoreJobsApi_GetFlexBackupRestoreJobExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetFlexBackupRestoreJobWithParams provides a mock function with given fields: ctx, args
func (_m *FlexRestoreJobsApi) GetFlexBackupRestoreJobWithParams(ctx context.Context, args *admin.GetFlexBackupRestoreJobApiParams) admin.GetFlexBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetFlexBackupRestoreJobWithParams")
	}

	var r0 admin.GetFlexBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetFlexBackupRestoreJobApiParams) admin.GetFlexBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetFlexBackupRestoreJobApiRequest)
	}

	return r0
}

// FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlexBackupRestoreJobWithParams'
type FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call struct {
	*mock.Call
}

// GetFlexBackupRestoreJobWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetFlexBackupRestoreJobApiParams
func (_e *FlexRestoreJobsApi_Expecter) GetFlexBackupRestoreJobWithParams(ctx any, args any) *FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call {
	return &FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call{Call: _e.mock.On("GetFlexBackupRestoreJobWithParams", ctx, args)}
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call) Run(run func(ctx context.Context, args *admin.GetFlexBackupRestoreJobApiParams)) *FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetFlexBackupRestoreJobApiParams))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call) Return(_a0 admin.GetFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetFlexBackupRestoreJobApiParams) admin.GetFlexBackupRestoreJobApiRequest) *FlexRestoreJobsApi_GetFlexBackupRestoreJobWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListFlexBackupRestoreJobs provides a mock function with given fields: ctx, groupId, name
func (_m *FlexRestoreJobsApi) ListFlexBackupRestoreJobs(ctx context.Context, groupId string, name string) admin.ListFlexBackupRestoreJobsApiRequest {
	ret := _m.Called(ctx, groupId, name)

	if len(ret) == 0 {
		panic("no return value specified for ListFlexBackupRestoreJobs")
	}

	var r0 admin.ListFlexBackupRestoreJobsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.ListFlexBackupRestoreJobsApiRequest); ok {
		r0 = rf(ctx, groupId, name)
	} else {
		r0 = ret.Get(0).(admin.ListFlexBackupRestoreJobsApiRequest)
	}

	return r0
}

// FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFlexBackupRestoreJobs'
type FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call struct {
	*mock.Call
}

// ListFlexBackupRestoreJobs is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - name string
func (_e *FlexRestoreJobsApi_Expecter) ListFlexBackupRestoreJobs(ctx any, groupId any, name any) *FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call {
	return &FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call{Call: _e.mock.On("ListFlexBackupRestoreJobs", ctx, groupId, name)}
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call) Run(run func(ctx context.Context, groupId string, name string)) *FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call) Return(_a0 admin.ListFlexBackupRestoreJobsApiRequest) *FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call) RunAndReturn(run func(context.Context, string, string) admin.ListFlexBackupRestoreJobsApiRequest) *FlexRestoreJobsApi_ListFlexBackupRestoreJobs_Call {
	_c.Call.Return(run)
	return _c
}

// ListFlexBackupRestoreJobsExecute provides a mock function with given fields: r
func (_m *FlexRestoreJobsApi) ListFlexBackupRestoreJobsExecute(r admin.ListFlexBackupRestoreJobsApiRequest) (*admin.PaginatedApiAtlasFlexBackupRestoreJob20241113, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListFlexBackupRestoreJobsExecute")
	}

	var r0 *admin.PaginatedApiAtlasFlexBackupRestoreJob20241113
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListFlexBackupRestoreJobsApiRequest) (*admin.PaginatedApiAtlasFlexBackupRestoreJob20241113, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListFlexBackupRestoreJobsApiRequest) *admin.PaginatedApiAtlasFlexBackupRestoreJob20241113); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedApiAtlasFlexBackupRestoreJob20241113)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListFlexBackupRestoreJobsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListFlexBackupRestoreJobsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFlexBackupRestoreJobsExecute'
type FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call struct {
	*mock.Call
}

// ListFlexBackupRestoreJobsExecute is a helper method to define mock.On call
//   - r admin.ListFlexBackupRestoreJobsApiRequest
func (_e *FlexRestoreJobsApi_Expecter) ListFlexBackupRestoreJobsExecute(r any) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call {
	return &FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call{Call: _e.mock.On("ListFlexBackupRestoreJobsExecute", r)}
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call) Run(run func(r admin.ListFlexBackupRestoreJobsApiRequest)) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListFlexBackupRestoreJobsApiRequest))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call) Return(_a0 *admin.PaginatedApiAtlasFlexBackupRestoreJob20241113, _a1 *http.Response, _a2 error) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call) RunAndReturn(run func(admin.ListFlexBackupRestoreJobsApiRequest) (*admin.PaginatedApiAtlasFlexBackupRestoreJob20241113, *http.Response, error)) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListFlexBackupRestoreJobsWithParams provides a mock function with given fields: ctx, args
func (_m *FlexRestoreJobsApi) ListFlexBackupRestoreJobsWithParams(ctx context.Context, args *admin.ListFlexBackupRestoreJobsApiParams) admin.ListFlexBackupRestoreJobsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListFlexBackupRestoreJobsWithParams")
	}

	var r0 admin.ListFlexBackupRestoreJobsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListFlexBackupRestoreJobsApiParams) admin.ListFlexBackupRestoreJobsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListFlexBackupRestoreJobsApiRequest)
	}

	return r0
}

// FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListFlexBackupRestoreJobsWithParams'
type FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call struct {
	*mock.Call
}

// ListFlexBackupRestoreJobsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListFlexBackupRestoreJobsApiParams
func (_e *FlexRestoreJobsApi_Expecter) ListFlexBackupRestoreJobsWithParams(ctx any, args any) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call {
	return &FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call{Call: _e.mock.On("ListFlexBackupRestoreJobsWithParams", ctx, args)}
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListFlexBackupRestoreJobsApiParams)) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListFlexBackupRestoreJobsApiParams))
	})
	return _c
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call) Return(_a0 admin.ListFlexBackupRestoreJobsApiRequest) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListFlexBackupRestoreJobsApiParams) admin.ListFlexBackupRestoreJobsApiRequest) *FlexRestoreJobsApi_ListFlexBackupRestoreJobsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewFlexRestoreJobsApi creates a new instance of FlexRestoreJobsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlexRestoreJobsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *FlexRestoreJobsApi {
	mock := &FlexRestoreJobsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
