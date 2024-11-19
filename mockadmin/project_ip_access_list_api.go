// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20241113001/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ProjectIPAccessListApi is an autogenerated mock type for the ProjectIPAccessListApi type
type ProjectIPAccessListApi struct {
	mock.Mock
}

type ProjectIPAccessListApi_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectIPAccessListApi) EXPECT() *ProjectIPAccessListApi_Expecter {
	return &ProjectIPAccessListApi_Expecter{mock: &_m.Mock}
}

// CreateProjectIpAccessList provides a mock function with given fields: ctx, groupId, networkPermissionEntry
func (_m *ProjectIPAccessListApi) CreateProjectIpAccessList(ctx context.Context, groupId string, networkPermissionEntry *[]admin.NetworkPermissionEntry) admin.CreateProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, groupId, networkPermissionEntry)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectIpAccessList")
	}

	var r0 admin.CreateProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, *[]admin.NetworkPermissionEntry) admin.CreateProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, groupId, networkPermissionEntry)
	} else {
		r0 = ret.Get(0).(admin.CreateProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_CreateProjectIpAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectIpAccessList'
type ProjectIPAccessListApi_CreateProjectIpAccessList_Call struct {
	*mock.Call
}

// CreateProjectIpAccessList is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - networkPermissionEntry *[]admin.NetworkPermissionEntry
func (_e *ProjectIPAccessListApi_Expecter) CreateProjectIpAccessList(ctx any, groupId any, networkPermissionEntry any) *ProjectIPAccessListApi_CreateProjectIpAccessList_Call {
	return &ProjectIPAccessListApi_CreateProjectIpAccessList_Call{Call: _e.mock.On("CreateProjectIpAccessList", ctx, groupId, networkPermissionEntry)}
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessList_Call) Run(run func(ctx context.Context, groupId string, networkPermissionEntry *[]admin.NetworkPermissionEntry)) *ProjectIPAccessListApi_CreateProjectIpAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*[]admin.NetworkPermissionEntry))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessList_Call) Return(_a0 admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApi_CreateProjectIpAccessList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessList_Call) RunAndReturn(run func(context.Context, string, *[]admin.NetworkPermissionEntry) admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApi_CreateProjectIpAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProjectIpAccessListExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApi) CreateProjectIpAccessListExecute(r admin.CreateProjectIpAccessListApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectIpAccessListExecute")
	}

	var r0 *admin.PaginatedNetworkAccess
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateProjectIpAccessListApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateProjectIpAccessListApiRequest) *admin.PaginatedNetworkAccess); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedNetworkAccess)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateProjectIpAccessListApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateProjectIpAccessListApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectIpAccessListExecute'
type ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call struct {
	*mock.Call
}

// CreateProjectIpAccessListExecute is a helper method to define mock.On call
//   - r admin.CreateProjectIpAccessListApiRequest
func (_e *ProjectIPAccessListApi_Expecter) CreateProjectIpAccessListExecute(r any) *ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call {
	return &ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call{Call: _e.mock.On("CreateProjectIpAccessListExecute", r)}
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call) Run(run func(r admin.CreateProjectIpAccessListApiRequest)) *ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateProjectIpAccessListApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call) Return(_a0 *admin.PaginatedNetworkAccess, _a1 *http.Response, _a2 error) *ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call) RunAndReturn(run func(admin.CreateProjectIpAccessListApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)) *ProjectIPAccessListApi_CreateProjectIpAccessListExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProjectIpAccessListWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApi) CreateProjectIpAccessListWithParams(ctx context.Context, args *admin.CreateProjectIpAccessListApiParams) admin.CreateProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectIpAccessListWithParams")
	}

	var r0 admin.CreateProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateProjectIpAccessListApiParams) admin.CreateProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectIpAccessListWithParams'
type ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call struct {
	*mock.Call
}

// CreateProjectIpAccessListWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateProjectIpAccessListApiParams
func (_e *ProjectIPAccessListApi_Expecter) CreateProjectIpAccessListWithParams(ctx any, args any) *ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call {
	return &ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call{Call: _e.mock.On("CreateProjectIpAccessListWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateProjectIpAccessListApiParams)) *ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateProjectIpAccessListApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call) Return(_a0 admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateProjectIpAccessListApiParams) admin.CreateProjectIpAccessListApiRequest) *ProjectIPAccessListApi_CreateProjectIpAccessListWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectIpAccessList provides a mock function with given fields: ctx, groupId, entryValue
func (_m *ProjectIPAccessListApi) DeleteProjectIpAccessList(ctx context.Context, groupId string, entryValue string) admin.DeleteProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, groupId, entryValue)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectIpAccessList")
	}

	var r0 admin.DeleteProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.DeleteProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, groupId, entryValue)
	} else {
		r0 = ret.Get(0).(admin.DeleteProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_DeleteProjectIpAccessList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectIpAccessList'
type ProjectIPAccessListApi_DeleteProjectIpAccessList_Call struct {
	*mock.Call
}

// DeleteProjectIpAccessList is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - entryValue string
func (_e *ProjectIPAccessListApi_Expecter) DeleteProjectIpAccessList(ctx any, groupId any, entryValue any) *ProjectIPAccessListApi_DeleteProjectIpAccessList_Call {
	return &ProjectIPAccessListApi_DeleteProjectIpAccessList_Call{Call: _e.mock.On("DeleteProjectIpAccessList", ctx, groupId, entryValue)}
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessList_Call) Run(run func(ctx context.Context, groupId string, entryValue string)) *ProjectIPAccessListApi_DeleteProjectIpAccessList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessList_Call) Return(_a0 admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApi_DeleteProjectIpAccessList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessList_Call) RunAndReturn(run func(context.Context, string, string) admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApi_DeleteProjectIpAccessList_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectIpAccessListExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApi) DeleteProjectIpAccessListExecute(r admin.DeleteProjectIpAccessListApiRequest) (any, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectIpAccessListExecute")
	}

	var r0 any
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DeleteProjectIpAccessListApiRequest) (any, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteProjectIpAccessListApiRequest) any); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteProjectIpAccessListApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DeleteProjectIpAccessListApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectIpAccessListExecute'
type ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call struct {
	*mock.Call
}

// DeleteProjectIpAccessListExecute is a helper method to define mock.On call
//   - r admin.DeleteProjectIpAccessListApiRequest
func (_e *ProjectIPAccessListApi_Expecter) DeleteProjectIpAccessListExecute(r any) *ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call {
	return &ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call{Call: _e.mock.On("DeleteProjectIpAccessListExecute", r)}
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call) Run(run func(r admin.DeleteProjectIpAccessListApiRequest)) *ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteProjectIpAccessListApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call) Return(_a0 any, _a1 *http.Response, _a2 error) *ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call) RunAndReturn(run func(admin.DeleteProjectIpAccessListApiRequest) (any, *http.Response, error)) *ProjectIPAccessListApi_DeleteProjectIpAccessListExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectIpAccessListWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApi) DeleteProjectIpAccessListWithParams(ctx context.Context, args *admin.DeleteProjectIpAccessListApiParams) admin.DeleteProjectIpAccessListApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectIpAccessListWithParams")
	}

	var r0 admin.DeleteProjectIpAccessListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteProjectIpAccessListApiParams) admin.DeleteProjectIpAccessListApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteProjectIpAccessListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectIpAccessListWithParams'
type ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call struct {
	*mock.Call
}

// DeleteProjectIpAccessListWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteProjectIpAccessListApiParams
func (_e *ProjectIPAccessListApi_Expecter) DeleteProjectIpAccessListWithParams(ctx any, args any) *ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call {
	return &ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call{Call: _e.mock.On("DeleteProjectIpAccessListWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteProjectIpAccessListApiParams)) *ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteProjectIpAccessListApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call) Return(_a0 admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteProjectIpAccessListApiParams) admin.DeleteProjectIpAccessListApiRequest) *ProjectIPAccessListApi_DeleteProjectIpAccessListWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpAccessListStatus provides a mock function with given fields: ctx, groupId, entryValue
func (_m *ProjectIPAccessListApi) GetProjectIpAccessListStatus(ctx context.Context, groupId string, entryValue string) admin.GetProjectIpAccessListStatusApiRequest {
	ret := _m.Called(ctx, groupId, entryValue)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpAccessListStatus")
	}

	var r0 admin.GetProjectIpAccessListStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetProjectIpAccessListStatusApiRequest); ok {
		r0 = rf(ctx, groupId, entryValue)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpAccessListStatusApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpAccessListStatus'
type ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call struct {
	*mock.Call
}

// GetProjectIpAccessListStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - entryValue string
func (_e *ProjectIPAccessListApi_Expecter) GetProjectIpAccessListStatus(ctx any, groupId any, entryValue any) *ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call {
	return &ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call{Call: _e.mock.On("GetProjectIpAccessListStatus", ctx, groupId, entryValue)}
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call) Run(run func(ctx context.Context, groupId string, entryValue string)) *ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call) Return(_a0 admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call) RunAndReturn(run func(context.Context, string, string) admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApi_GetProjectIpAccessListStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpAccessListStatusExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApi) GetProjectIpAccessListStatusExecute(r admin.GetProjectIpAccessListStatusApiRequest) (*admin.NetworkPermissionEntryStatus, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpAccessListStatusExecute")
	}

	var r0 *admin.NetworkPermissionEntryStatus
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpAccessListStatusApiRequest) (*admin.NetworkPermissionEntryStatus, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpAccessListStatusApiRequest) *admin.NetworkPermissionEntryStatus); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NetworkPermissionEntryStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetProjectIpAccessListStatusApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetProjectIpAccessListStatusApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpAccessListStatusExecute'
type ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call struct {
	*mock.Call
}

// GetProjectIpAccessListStatusExecute is a helper method to define mock.On call
//   - r admin.GetProjectIpAccessListStatusApiRequest
func (_e *ProjectIPAccessListApi_Expecter) GetProjectIpAccessListStatusExecute(r any) *ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call {
	return &ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call{Call: _e.mock.On("GetProjectIpAccessListStatusExecute", r)}
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call) Run(run func(r admin.GetProjectIpAccessListStatusApiRequest)) *ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetProjectIpAccessListStatusApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call) Return(_a0 *admin.NetworkPermissionEntryStatus, _a1 *http.Response, _a2 error) *ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call) RunAndReturn(run func(admin.GetProjectIpAccessListStatusApiRequest) (*admin.NetworkPermissionEntryStatus, *http.Response, error)) *ProjectIPAccessListApi_GetProjectIpAccessListStatusExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpAccessListStatusWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApi) GetProjectIpAccessListStatusWithParams(ctx context.Context, args *admin.GetProjectIpAccessListStatusApiParams) admin.GetProjectIpAccessListStatusApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpAccessListStatusWithParams")
	}

	var r0 admin.GetProjectIpAccessListStatusApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetProjectIpAccessListStatusApiParams) admin.GetProjectIpAccessListStatusApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpAccessListStatusApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpAccessListStatusWithParams'
type ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call struct {
	*mock.Call
}

// GetProjectIpAccessListStatusWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetProjectIpAccessListStatusApiParams
func (_e *ProjectIPAccessListApi_Expecter) GetProjectIpAccessListStatusWithParams(ctx any, args any) *ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call {
	return &ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call{Call: _e.mock.On("GetProjectIpAccessListStatusWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call) Run(run func(ctx context.Context, args *admin.GetProjectIpAccessListStatusApiParams)) *ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetProjectIpAccessListStatusApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call) Return(_a0 admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetProjectIpAccessListStatusApiParams) admin.GetProjectIpAccessListStatusApiRequest) *ProjectIPAccessListApi_GetProjectIpAccessListStatusWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpList provides a mock function with given fields: ctx, groupId, entryValue
func (_m *ProjectIPAccessListApi) GetProjectIpList(ctx context.Context, groupId string, entryValue string) admin.GetProjectIpListApiRequest {
	ret := _m.Called(ctx, groupId, entryValue)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpList")
	}

	var r0 admin.GetProjectIpListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetProjectIpListApiRequest); ok {
		r0 = rf(ctx, groupId, entryValue)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_GetProjectIpList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpList'
type ProjectIPAccessListApi_GetProjectIpList_Call struct {
	*mock.Call
}

// GetProjectIpList is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - entryValue string
func (_e *ProjectIPAccessListApi_Expecter) GetProjectIpList(ctx any, groupId any, entryValue any) *ProjectIPAccessListApi_GetProjectIpList_Call {
	return &ProjectIPAccessListApi_GetProjectIpList_Call{Call: _e.mock.On("GetProjectIpList", ctx, groupId, entryValue)}
}

func (_c *ProjectIPAccessListApi_GetProjectIpList_Call) Run(run func(ctx context.Context, groupId string, entryValue string)) *ProjectIPAccessListApi_GetProjectIpList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpList_Call) Return(_a0 admin.GetProjectIpListApiRequest) *ProjectIPAccessListApi_GetProjectIpList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpList_Call) RunAndReturn(run func(context.Context, string, string) admin.GetProjectIpListApiRequest) *ProjectIPAccessListApi_GetProjectIpList_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpListExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApi) GetProjectIpListExecute(r admin.GetProjectIpListApiRequest) (*admin.NetworkPermissionEntry, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpListExecute")
	}

	var r0 *admin.NetworkPermissionEntry
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpListApiRequest) (*admin.NetworkPermissionEntry, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetProjectIpListApiRequest) *admin.NetworkPermissionEntry); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.NetworkPermissionEntry)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetProjectIpListApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetProjectIpListApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApi_GetProjectIpListExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpListExecute'
type ProjectIPAccessListApi_GetProjectIpListExecute_Call struct {
	*mock.Call
}

// GetProjectIpListExecute is a helper method to define mock.On call
//   - r admin.GetProjectIpListApiRequest
func (_e *ProjectIPAccessListApi_Expecter) GetProjectIpListExecute(r any) *ProjectIPAccessListApi_GetProjectIpListExecute_Call {
	return &ProjectIPAccessListApi_GetProjectIpListExecute_Call{Call: _e.mock.On("GetProjectIpListExecute", r)}
}

func (_c *ProjectIPAccessListApi_GetProjectIpListExecute_Call) Run(run func(r admin.GetProjectIpListApiRequest)) *ProjectIPAccessListApi_GetProjectIpListExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetProjectIpListApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpListExecute_Call) Return(_a0 *admin.NetworkPermissionEntry, _a1 *http.Response, _a2 error) *ProjectIPAccessListApi_GetProjectIpListExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpListExecute_Call) RunAndReturn(run func(admin.GetProjectIpListApiRequest) (*admin.NetworkPermissionEntry, *http.Response, error)) *ProjectIPAccessListApi_GetProjectIpListExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectIpListWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApi) GetProjectIpListWithParams(ctx context.Context, args *admin.GetProjectIpListApiParams) admin.GetProjectIpListApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectIpListWithParams")
	}

	var r0 admin.GetProjectIpListApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetProjectIpListApiParams) admin.GetProjectIpListApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetProjectIpListApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_GetProjectIpListWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectIpListWithParams'
type ProjectIPAccessListApi_GetProjectIpListWithParams_Call struct {
	*mock.Call
}

// GetProjectIpListWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetProjectIpListApiParams
func (_e *ProjectIPAccessListApi_Expecter) GetProjectIpListWithParams(ctx any, args any) *ProjectIPAccessListApi_GetProjectIpListWithParams_Call {
	return &ProjectIPAccessListApi_GetProjectIpListWithParams_Call{Call: _e.mock.On("GetProjectIpListWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApi_GetProjectIpListWithParams_Call) Run(run func(ctx context.Context, args *admin.GetProjectIpListApiParams)) *ProjectIPAccessListApi_GetProjectIpListWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetProjectIpListApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpListWithParams_Call) Return(_a0 admin.GetProjectIpListApiRequest) *ProjectIPAccessListApi_GetProjectIpListWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_GetProjectIpListWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetProjectIpListApiParams) admin.GetProjectIpListApiRequest) *ProjectIPAccessListApi_GetProjectIpListWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectIpAccessLists provides a mock function with given fields: ctx, groupId
func (_m *ProjectIPAccessListApi) ListProjectIpAccessLists(ctx context.Context, groupId string) admin.ListProjectIpAccessListsApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectIpAccessLists")
	}

	var r0 admin.ListProjectIpAccessListsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListProjectIpAccessListsApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListProjectIpAccessListsApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_ListProjectIpAccessLists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectIpAccessLists'
type ProjectIPAccessListApi_ListProjectIpAccessLists_Call struct {
	*mock.Call
}

// ListProjectIpAccessLists is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *ProjectIPAccessListApi_Expecter) ListProjectIpAccessLists(ctx any, groupId any) *ProjectIPAccessListApi_ListProjectIpAccessLists_Call {
	return &ProjectIPAccessListApi_ListProjectIpAccessLists_Call{Call: _e.mock.On("ListProjectIpAccessLists", ctx, groupId)}
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessLists_Call) Run(run func(ctx context.Context, groupId string)) *ProjectIPAccessListApi_ListProjectIpAccessLists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessLists_Call) Return(_a0 admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApi_ListProjectIpAccessLists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessLists_Call) RunAndReturn(run func(context.Context, string) admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApi_ListProjectIpAccessLists_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectIpAccessListsExecute provides a mock function with given fields: r
func (_m *ProjectIPAccessListApi) ListProjectIpAccessListsExecute(r admin.ListProjectIpAccessListsApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectIpAccessListsExecute")
	}

	var r0 *admin.PaginatedNetworkAccess
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListProjectIpAccessListsApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListProjectIpAccessListsApiRequest) *admin.PaginatedNetworkAccess); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedNetworkAccess)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListProjectIpAccessListsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListProjectIpAccessListsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectIpAccessListsExecute'
type ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call struct {
	*mock.Call
}

// ListProjectIpAccessListsExecute is a helper method to define mock.On call
//   - r admin.ListProjectIpAccessListsApiRequest
func (_e *ProjectIPAccessListApi_Expecter) ListProjectIpAccessListsExecute(r any) *ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call {
	return &ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call{Call: _e.mock.On("ListProjectIpAccessListsExecute", r)}
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call) Run(run func(r admin.ListProjectIpAccessListsApiRequest)) *ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListProjectIpAccessListsApiRequest))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call) Return(_a0 *admin.PaginatedNetworkAccess, _a1 *http.Response, _a2 error) *ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call) RunAndReturn(run func(admin.ListProjectIpAccessListsApiRequest) (*admin.PaginatedNetworkAccess, *http.Response, error)) *ProjectIPAccessListApi_ListProjectIpAccessListsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectIpAccessListsWithParams provides a mock function with given fields: ctx, args
func (_m *ProjectIPAccessListApi) ListProjectIpAccessListsWithParams(ctx context.Context, args *admin.ListProjectIpAccessListsApiParams) admin.ListProjectIpAccessListsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectIpAccessListsWithParams")
	}

	var r0 admin.ListProjectIpAccessListsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListProjectIpAccessListsApiParams) admin.ListProjectIpAccessListsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListProjectIpAccessListsApiRequest)
	}

	return r0
}

// ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectIpAccessListsWithParams'
type ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call struct {
	*mock.Call
}

// ListProjectIpAccessListsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListProjectIpAccessListsApiParams
func (_e *ProjectIPAccessListApi_Expecter) ListProjectIpAccessListsWithParams(ctx any, args any) *ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call {
	return &ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call{Call: _e.mock.On("ListProjectIpAccessListsWithParams", ctx, args)}
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListProjectIpAccessListsApiParams)) *ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListProjectIpAccessListsApiParams))
	})
	return _c
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call) Return(_a0 admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListProjectIpAccessListsApiParams) admin.ListProjectIpAccessListsApiRequest) *ProjectIPAccessListApi_ListProjectIpAccessListsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectIPAccessListApi creates a new instance of ProjectIPAccessListApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectIPAccessListApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectIPAccessListApi {
	mock := &ProjectIPAccessListApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
