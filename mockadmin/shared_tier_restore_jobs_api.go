// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20250312005/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// SharedTierRestoreJobsApi is an autogenerated mock type for the SharedTierRestoreJobsApi type
type SharedTierRestoreJobsApi struct {
	mock.Mock
}

type SharedTierRestoreJobsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *SharedTierRestoreJobsApi) EXPECT() *SharedTierRestoreJobsApi_Expecter {
	return &SharedTierRestoreJobsApi_Expecter{mock: &_m.Mock}
}

// CreateSharedClusterBackupRestoreJob provides a mock function with given fields: ctx, clusterName, groupId, tenantRestore
func (_m *SharedTierRestoreJobsApi) CreateSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, tenantRestore *admin.TenantRestore) admin.CreateSharedClusterBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, clusterName, groupId, tenantRestore)

	if len(ret) == 0 {
		panic("no return value specified for CreateSharedClusterBackupRestoreJob")
	}

	var r0 admin.CreateSharedClusterBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.TenantRestore) admin.CreateSharedClusterBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, clusterName, groupId, tenantRestore)
	} else {
		r0 = ret.Get(0).(admin.CreateSharedClusterBackupRestoreJobApiRequest)
	}

	return r0
}

// SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSharedClusterBackupRestoreJob'
type SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call struct {
	*mock.Call
}

// CreateSharedClusterBackupRestoreJob is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - groupId string
//   - tenantRestore *admin.TenantRestore
func (_e *SharedTierRestoreJobsApi_Expecter) CreateSharedClusterBackupRestoreJob(ctx any, clusterName any, groupId any, tenantRestore any) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call {
	return &SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call{Call: _e.mock.On("CreateSharedClusterBackupRestoreJob", ctx, clusterName, groupId, tenantRestore)}
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call) Run(run func(ctx context.Context, clusterName string, groupId string, tenantRestore *admin.TenantRestore)) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.TenantRestore))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call) Return(_a0 admin.CreateSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call) RunAndReturn(run func(context.Context, string, string, *admin.TenantRestore) admin.CreateSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJob_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSharedClusterBackupRestoreJobExecute provides a mock function with given fields: r
func (_m *SharedTierRestoreJobsApi) CreateSharedClusterBackupRestoreJobExecute(r admin.CreateSharedClusterBackupRestoreJobApiRequest) (*admin.TenantRestore, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateSharedClusterBackupRestoreJobExecute")
	}

	var r0 *admin.TenantRestore
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateSharedClusterBackupRestoreJobApiRequest) (*admin.TenantRestore, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateSharedClusterBackupRestoreJobApiRequest) *admin.TenantRestore); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.TenantRestore)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateSharedClusterBackupRestoreJobApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateSharedClusterBackupRestoreJobApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSharedClusterBackupRestoreJobExecute'
type SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call struct {
	*mock.Call
}

// CreateSharedClusterBackupRestoreJobExecute is a helper method to define mock.On call
//   - r admin.CreateSharedClusterBackupRestoreJobApiRequest
func (_e *SharedTierRestoreJobsApi_Expecter) CreateSharedClusterBackupRestoreJobExecute(r any) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call {
	return &SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call{Call: _e.mock.On("CreateSharedClusterBackupRestoreJobExecute", r)}
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call) Run(run func(r admin.CreateSharedClusterBackupRestoreJobApiRequest)) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateSharedClusterBackupRestoreJobApiRequest))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call) Return(_a0 *admin.TenantRestore, _a1 *http.Response, _a2 error) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call) RunAndReturn(run func(admin.CreateSharedClusterBackupRestoreJobApiRequest) (*admin.TenantRestore, *http.Response, error)) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSharedClusterBackupRestoreJobWithParams provides a mock function with given fields: ctx, args
func (_m *SharedTierRestoreJobsApi) CreateSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *admin.CreateSharedClusterBackupRestoreJobApiParams) admin.CreateSharedClusterBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateSharedClusterBackupRestoreJobWithParams")
	}

	var r0 admin.CreateSharedClusterBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateSharedClusterBackupRestoreJobApiParams) admin.CreateSharedClusterBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateSharedClusterBackupRestoreJobApiRequest)
	}

	return r0
}

// SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSharedClusterBackupRestoreJobWithParams'
type SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call struct {
	*mock.Call
}

// CreateSharedClusterBackupRestoreJobWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateSharedClusterBackupRestoreJobApiParams
func (_e *SharedTierRestoreJobsApi_Expecter) CreateSharedClusterBackupRestoreJobWithParams(ctx any, args any) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call {
	return &SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call{Call: _e.mock.On("CreateSharedClusterBackupRestoreJobWithParams", ctx, args)}
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateSharedClusterBackupRestoreJobApiParams)) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateSharedClusterBackupRestoreJobApiParams))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call) Return(_a0 admin.CreateSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateSharedClusterBackupRestoreJobApiParams) admin.CreateSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_CreateSharedClusterBackupRestoreJobWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetSharedClusterBackupRestoreJob provides a mock function with given fields: ctx, clusterName, groupId, restoreId
func (_m *SharedTierRestoreJobsApi) GetSharedClusterBackupRestoreJob(ctx context.Context, clusterName string, groupId string, restoreId string) admin.GetSharedClusterBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, clusterName, groupId, restoreId)

	if len(ret) == 0 {
		panic("no return value specified for GetSharedClusterBackupRestoreJob")
	}

	var r0 admin.GetSharedClusterBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) admin.GetSharedClusterBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, clusterName, groupId, restoreId)
	} else {
		r0 = ret.Get(0).(admin.GetSharedClusterBackupRestoreJobApiRequest)
	}

	return r0
}

// SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSharedClusterBackupRestoreJob'
type SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call struct {
	*mock.Call
}

// GetSharedClusterBackupRestoreJob is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - groupId string
//   - restoreId string
func (_e *SharedTierRestoreJobsApi_Expecter) GetSharedClusterBackupRestoreJob(ctx any, clusterName any, groupId any, restoreId any) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call {
	return &SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call{Call: _e.mock.On("GetSharedClusterBackupRestoreJob", ctx, clusterName, groupId, restoreId)}
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call) Run(run func(ctx context.Context, clusterName string, groupId string, restoreId string)) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call) Return(_a0 admin.GetSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call) RunAndReturn(run func(context.Context, string, string, string) admin.GetSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJob_Call {
	_c.Call.Return(run)
	return _c
}

// GetSharedClusterBackupRestoreJobExecute provides a mock function with given fields: r
func (_m *SharedTierRestoreJobsApi) GetSharedClusterBackupRestoreJobExecute(r admin.GetSharedClusterBackupRestoreJobApiRequest) (*admin.TenantRestore, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetSharedClusterBackupRestoreJobExecute")
	}

	var r0 *admin.TenantRestore
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetSharedClusterBackupRestoreJobApiRequest) (*admin.TenantRestore, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetSharedClusterBackupRestoreJobApiRequest) *admin.TenantRestore); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.TenantRestore)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetSharedClusterBackupRestoreJobApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetSharedClusterBackupRestoreJobApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSharedClusterBackupRestoreJobExecute'
type SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call struct {
	*mock.Call
}

// GetSharedClusterBackupRestoreJobExecute is a helper method to define mock.On call
//   - r admin.GetSharedClusterBackupRestoreJobApiRequest
func (_e *SharedTierRestoreJobsApi_Expecter) GetSharedClusterBackupRestoreJobExecute(r any) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call {
	return &SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call{Call: _e.mock.On("GetSharedClusterBackupRestoreJobExecute", r)}
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call) Run(run func(r admin.GetSharedClusterBackupRestoreJobApiRequest)) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetSharedClusterBackupRestoreJobApiRequest))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call) Return(_a0 *admin.TenantRestore, _a1 *http.Response, _a2 error) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call) RunAndReturn(run func(admin.GetSharedClusterBackupRestoreJobApiRequest) (*admin.TenantRestore, *http.Response, error)) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetSharedClusterBackupRestoreJobWithParams provides a mock function with given fields: ctx, args
func (_m *SharedTierRestoreJobsApi) GetSharedClusterBackupRestoreJobWithParams(ctx context.Context, args *admin.GetSharedClusterBackupRestoreJobApiParams) admin.GetSharedClusterBackupRestoreJobApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetSharedClusterBackupRestoreJobWithParams")
	}

	var r0 admin.GetSharedClusterBackupRestoreJobApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetSharedClusterBackupRestoreJobApiParams) admin.GetSharedClusterBackupRestoreJobApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetSharedClusterBackupRestoreJobApiRequest)
	}

	return r0
}

// SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSharedClusterBackupRestoreJobWithParams'
type SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call struct {
	*mock.Call
}

// GetSharedClusterBackupRestoreJobWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetSharedClusterBackupRestoreJobApiParams
func (_e *SharedTierRestoreJobsApi_Expecter) GetSharedClusterBackupRestoreJobWithParams(ctx any, args any) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call {
	return &SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call{Call: _e.mock.On("GetSharedClusterBackupRestoreJobWithParams", ctx, args)}
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call) Run(run func(ctx context.Context, args *admin.GetSharedClusterBackupRestoreJobApiParams)) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetSharedClusterBackupRestoreJobApiParams))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call) Return(_a0 admin.GetSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetSharedClusterBackupRestoreJobApiParams) admin.GetSharedClusterBackupRestoreJobApiRequest) *SharedTierRestoreJobsApi_GetSharedClusterBackupRestoreJobWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListSharedClusterBackupRestoreJobs provides a mock function with given fields: ctx, clusterName, groupId
func (_m *SharedTierRestoreJobsApi) ListSharedClusterBackupRestoreJobs(ctx context.Context, clusterName string, groupId string) admin.ListSharedClusterBackupRestoreJobsApiRequest {
	ret := _m.Called(ctx, clusterName, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListSharedClusterBackupRestoreJobs")
	}

	var r0 admin.ListSharedClusterBackupRestoreJobsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.ListSharedClusterBackupRestoreJobsApiRequest); ok {
		r0 = rf(ctx, clusterName, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListSharedClusterBackupRestoreJobsApiRequest)
	}

	return r0
}

// SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSharedClusterBackupRestoreJobs'
type SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call struct {
	*mock.Call
}

// ListSharedClusterBackupRestoreJobs is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - groupId string
func (_e *SharedTierRestoreJobsApi_Expecter) ListSharedClusterBackupRestoreJobs(ctx any, clusterName any, groupId any) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call {
	return &SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call{Call: _e.mock.On("ListSharedClusterBackupRestoreJobs", ctx, clusterName, groupId)}
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call) Run(run func(ctx context.Context, clusterName string, groupId string)) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call) Return(_a0 admin.ListSharedClusterBackupRestoreJobsApiRequest) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call) RunAndReturn(run func(context.Context, string, string) admin.ListSharedClusterBackupRestoreJobsApiRequest) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobs_Call {
	_c.Call.Return(run)
	return _c
}

// ListSharedClusterBackupRestoreJobsExecute provides a mock function with given fields: r
func (_m *SharedTierRestoreJobsApi) ListSharedClusterBackupRestoreJobsExecute(r admin.ListSharedClusterBackupRestoreJobsApiRequest) (*admin.PaginatedTenantRestore, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListSharedClusterBackupRestoreJobsExecute")
	}

	var r0 *admin.PaginatedTenantRestore
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListSharedClusterBackupRestoreJobsApiRequest) (*admin.PaginatedTenantRestore, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListSharedClusterBackupRestoreJobsApiRequest) *admin.PaginatedTenantRestore); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedTenantRestore)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListSharedClusterBackupRestoreJobsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListSharedClusterBackupRestoreJobsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSharedClusterBackupRestoreJobsExecute'
type SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call struct {
	*mock.Call
}

// ListSharedClusterBackupRestoreJobsExecute is a helper method to define mock.On call
//   - r admin.ListSharedClusterBackupRestoreJobsApiRequest
func (_e *SharedTierRestoreJobsApi_Expecter) ListSharedClusterBackupRestoreJobsExecute(r any) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call {
	return &SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call{Call: _e.mock.On("ListSharedClusterBackupRestoreJobsExecute", r)}
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call) Run(run func(r admin.ListSharedClusterBackupRestoreJobsApiRequest)) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListSharedClusterBackupRestoreJobsApiRequest))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call) Return(_a0 *admin.PaginatedTenantRestore, _a1 *http.Response, _a2 error) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call) RunAndReturn(run func(admin.ListSharedClusterBackupRestoreJobsApiRequest) (*admin.PaginatedTenantRestore, *http.Response, error)) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListSharedClusterBackupRestoreJobsWithParams provides a mock function with given fields: ctx, args
func (_m *SharedTierRestoreJobsApi) ListSharedClusterBackupRestoreJobsWithParams(ctx context.Context, args *admin.ListSharedClusterBackupRestoreJobsApiParams) admin.ListSharedClusterBackupRestoreJobsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListSharedClusterBackupRestoreJobsWithParams")
	}

	var r0 admin.ListSharedClusterBackupRestoreJobsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListSharedClusterBackupRestoreJobsApiParams) admin.ListSharedClusterBackupRestoreJobsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListSharedClusterBackupRestoreJobsApiRequest)
	}

	return r0
}

// SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSharedClusterBackupRestoreJobsWithParams'
type SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call struct {
	*mock.Call
}

// ListSharedClusterBackupRestoreJobsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListSharedClusterBackupRestoreJobsApiParams
func (_e *SharedTierRestoreJobsApi_Expecter) ListSharedClusterBackupRestoreJobsWithParams(ctx any, args any) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call {
	return &SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call{Call: _e.mock.On("ListSharedClusterBackupRestoreJobsWithParams", ctx, args)}
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListSharedClusterBackupRestoreJobsApiParams)) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListSharedClusterBackupRestoreJobsApiParams))
	})
	return _c
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call) Return(_a0 admin.ListSharedClusterBackupRestoreJobsApiRequest) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListSharedClusterBackupRestoreJobsApiParams) admin.ListSharedClusterBackupRestoreJobsApiRequest) *SharedTierRestoreJobsApi_ListSharedClusterBackupRestoreJobsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewSharedTierRestoreJobsApi creates a new instance of SharedTierRestoreJobsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSharedTierRestoreJobsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *SharedTierRestoreJobsApi {
	mock := &SharedTierRestoreJobsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
