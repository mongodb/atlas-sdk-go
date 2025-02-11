// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20241113005/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// SharedTierSnapshotsApi is an autogenerated mock type for the SharedTierSnapshotsApi type
type SharedTierSnapshotsApi struct {
	mock.Mock
}

type SharedTierSnapshotsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *SharedTierSnapshotsApi) EXPECT() *SharedTierSnapshotsApi_Expecter {
	return &SharedTierSnapshotsApi_Expecter{mock: &_m.Mock}
}

// DownloadSharedClusterBackup provides a mock function with given fields: ctx, clusterName, groupId, tenantRestore
func (_m *SharedTierSnapshotsApi) DownloadSharedClusterBackup(ctx context.Context, clusterName string, groupId string, tenantRestore *admin.TenantRestore) admin.DownloadSharedClusterBackupApiRequest {
	ret := _m.Called(ctx, clusterName, groupId, tenantRestore)

	if len(ret) == 0 {
		panic("no return value specified for DownloadSharedClusterBackup")
	}

	var r0 admin.DownloadSharedClusterBackupApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.TenantRestore) admin.DownloadSharedClusterBackupApiRequest); ok {
		r0 = rf(ctx, clusterName, groupId, tenantRestore)
	} else {
		r0 = ret.Get(0).(admin.DownloadSharedClusterBackupApiRequest)
	}

	return r0
}

// SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadSharedClusterBackup'
type SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call struct {
	*mock.Call
}

// DownloadSharedClusterBackup is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - groupId string
//   - tenantRestore *admin.TenantRestore
func (_e *SharedTierSnapshotsApi_Expecter) DownloadSharedClusterBackup(ctx any, clusterName any, groupId any, tenantRestore any) *SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call {
	return &SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call{Call: _e.mock.On("DownloadSharedClusterBackup", ctx, clusterName, groupId, tenantRestore)}
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call) Run(run func(ctx context.Context, clusterName string, groupId string, tenantRestore *admin.TenantRestore)) *SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.TenantRestore))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call) Return(_a0 admin.DownloadSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call) RunAndReturn(run func(context.Context, string, string, *admin.TenantRestore) admin.DownloadSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_DownloadSharedClusterBackup_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadSharedClusterBackupExecute provides a mock function with given fields: r
func (_m *SharedTierSnapshotsApi) DownloadSharedClusterBackupExecute(r admin.DownloadSharedClusterBackupApiRequest) (*admin.TenantRestore, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DownloadSharedClusterBackupExecute")
	}

	var r0 *admin.TenantRestore
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DownloadSharedClusterBackupApiRequest) (*admin.TenantRestore, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DownloadSharedClusterBackupApiRequest) *admin.TenantRestore); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.TenantRestore)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DownloadSharedClusterBackupApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DownloadSharedClusterBackupApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadSharedClusterBackupExecute'
type SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call struct {
	*mock.Call
}

// DownloadSharedClusterBackupExecute is a helper method to define mock.On call
//   - r admin.DownloadSharedClusterBackupApiRequest
func (_e *SharedTierSnapshotsApi_Expecter) DownloadSharedClusterBackupExecute(r any) *SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call {
	return &SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call{Call: _e.mock.On("DownloadSharedClusterBackupExecute", r)}
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call) Run(run func(r admin.DownloadSharedClusterBackupApiRequest)) *SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DownloadSharedClusterBackupApiRequest))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call) Return(_a0 *admin.TenantRestore, _a1 *http.Response, _a2 error) *SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call) RunAndReturn(run func(admin.DownloadSharedClusterBackupApiRequest) (*admin.TenantRestore, *http.Response, error)) *SharedTierSnapshotsApi_DownloadSharedClusterBackupExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DownloadSharedClusterBackupWithParams provides a mock function with given fields: ctx, args
func (_m *SharedTierSnapshotsApi) DownloadSharedClusterBackupWithParams(ctx context.Context, args *admin.DownloadSharedClusterBackupApiParams) admin.DownloadSharedClusterBackupApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DownloadSharedClusterBackupWithParams")
	}

	var r0 admin.DownloadSharedClusterBackupApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DownloadSharedClusterBackupApiParams) admin.DownloadSharedClusterBackupApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DownloadSharedClusterBackupApiRequest)
	}

	return r0
}

// SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadSharedClusterBackupWithParams'
type SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call struct {
	*mock.Call
}

// DownloadSharedClusterBackupWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DownloadSharedClusterBackupApiParams
func (_e *SharedTierSnapshotsApi_Expecter) DownloadSharedClusterBackupWithParams(ctx any, args any) *SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call {
	return &SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call{Call: _e.mock.On("DownloadSharedClusterBackupWithParams", ctx, args)}
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call) Run(run func(ctx context.Context, args *admin.DownloadSharedClusterBackupApiParams)) *SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DownloadSharedClusterBackupApiParams))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call) Return(_a0 admin.DownloadSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call) RunAndReturn(run func(context.Context, *admin.DownloadSharedClusterBackupApiParams) admin.DownloadSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_DownloadSharedClusterBackupWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetSharedClusterBackup provides a mock function with given fields: ctx, groupId, clusterName, snapshotId
func (_m *SharedTierSnapshotsApi) GetSharedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) admin.GetSharedClusterBackupApiRequest {
	ret := _m.Called(ctx, groupId, clusterName, snapshotId)

	if len(ret) == 0 {
		panic("no return value specified for GetSharedClusterBackup")
	}

	var r0 admin.GetSharedClusterBackupApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) admin.GetSharedClusterBackupApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName, snapshotId)
	} else {
		r0 = ret.Get(0).(admin.GetSharedClusterBackupApiRequest)
	}

	return r0
}

// SharedTierSnapshotsApi_GetSharedClusterBackup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSharedClusterBackup'
type SharedTierSnapshotsApi_GetSharedClusterBackup_Call struct {
	*mock.Call
}

// GetSharedClusterBackup is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
//   - snapshotId string
func (_e *SharedTierSnapshotsApi_Expecter) GetSharedClusterBackup(ctx any, groupId any, clusterName any, snapshotId any) *SharedTierSnapshotsApi_GetSharedClusterBackup_Call {
	return &SharedTierSnapshotsApi_GetSharedClusterBackup_Call{Call: _e.mock.On("GetSharedClusterBackup", ctx, groupId, clusterName, snapshotId)}
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackup_Call) Run(run func(ctx context.Context, groupId string, clusterName string, snapshotId string)) *SharedTierSnapshotsApi_GetSharedClusterBackup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackup_Call) Return(_a0 admin.GetSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_GetSharedClusterBackup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackup_Call) RunAndReturn(run func(context.Context, string, string, string) admin.GetSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_GetSharedClusterBackup_Call {
	_c.Call.Return(run)
	return _c
}

// GetSharedClusterBackupExecute provides a mock function with given fields: r
func (_m *SharedTierSnapshotsApi) GetSharedClusterBackupExecute(r admin.GetSharedClusterBackupApiRequest) (*admin.BackupTenantSnapshot, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetSharedClusterBackupExecute")
	}

	var r0 *admin.BackupTenantSnapshot
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetSharedClusterBackupApiRequest) (*admin.BackupTenantSnapshot, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetSharedClusterBackupApiRequest) *admin.BackupTenantSnapshot); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.BackupTenantSnapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetSharedClusterBackupApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetSharedClusterBackupApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSharedClusterBackupExecute'
type SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call struct {
	*mock.Call
}

// GetSharedClusterBackupExecute is a helper method to define mock.On call
//   - r admin.GetSharedClusterBackupApiRequest
func (_e *SharedTierSnapshotsApi_Expecter) GetSharedClusterBackupExecute(r any) *SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call {
	return &SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call{Call: _e.mock.On("GetSharedClusterBackupExecute", r)}
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call) Run(run func(r admin.GetSharedClusterBackupApiRequest)) *SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetSharedClusterBackupApiRequest))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call) Return(_a0 *admin.BackupTenantSnapshot, _a1 *http.Response, _a2 error) *SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call) RunAndReturn(run func(admin.GetSharedClusterBackupApiRequest) (*admin.BackupTenantSnapshot, *http.Response, error)) *SharedTierSnapshotsApi_GetSharedClusterBackupExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetSharedClusterBackupWithParams provides a mock function with given fields: ctx, args
func (_m *SharedTierSnapshotsApi) GetSharedClusterBackupWithParams(ctx context.Context, args *admin.GetSharedClusterBackupApiParams) admin.GetSharedClusterBackupApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetSharedClusterBackupWithParams")
	}

	var r0 admin.GetSharedClusterBackupApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetSharedClusterBackupApiParams) admin.GetSharedClusterBackupApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetSharedClusterBackupApiRequest)
	}

	return r0
}

// SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSharedClusterBackupWithParams'
type SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call struct {
	*mock.Call
}

// GetSharedClusterBackupWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetSharedClusterBackupApiParams
func (_e *SharedTierSnapshotsApi_Expecter) GetSharedClusterBackupWithParams(ctx any, args any) *SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call {
	return &SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call{Call: _e.mock.On("GetSharedClusterBackupWithParams", ctx, args)}
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call) Run(run func(ctx context.Context, args *admin.GetSharedClusterBackupApiParams)) *SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetSharedClusterBackupApiParams))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call) Return(_a0 admin.GetSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetSharedClusterBackupApiParams) admin.GetSharedClusterBackupApiRequest) *SharedTierSnapshotsApi_GetSharedClusterBackupWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListSharedClusterBackups provides a mock function with given fields: ctx, groupId, clusterName
func (_m *SharedTierSnapshotsApi) ListSharedClusterBackups(ctx context.Context, groupId string, clusterName string) admin.ListSharedClusterBackupsApiRequest {
	ret := _m.Called(ctx, groupId, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for ListSharedClusterBackups")
	}

	var r0 admin.ListSharedClusterBackupsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.ListSharedClusterBackupsApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName)
	} else {
		r0 = ret.Get(0).(admin.ListSharedClusterBackupsApiRequest)
	}

	return r0
}

// SharedTierSnapshotsApi_ListSharedClusterBackups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSharedClusterBackups'
type SharedTierSnapshotsApi_ListSharedClusterBackups_Call struct {
	*mock.Call
}

// ListSharedClusterBackups is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
func (_e *SharedTierSnapshotsApi_Expecter) ListSharedClusterBackups(ctx any, groupId any, clusterName any) *SharedTierSnapshotsApi_ListSharedClusterBackups_Call {
	return &SharedTierSnapshotsApi_ListSharedClusterBackups_Call{Call: _e.mock.On("ListSharedClusterBackups", ctx, groupId, clusterName)}
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackups_Call) Run(run func(ctx context.Context, groupId string, clusterName string)) *SharedTierSnapshotsApi_ListSharedClusterBackups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackups_Call) Return(_a0 admin.ListSharedClusterBackupsApiRequest) *SharedTierSnapshotsApi_ListSharedClusterBackups_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackups_Call) RunAndReturn(run func(context.Context, string, string) admin.ListSharedClusterBackupsApiRequest) *SharedTierSnapshotsApi_ListSharedClusterBackups_Call {
	_c.Call.Return(run)
	return _c
}

// ListSharedClusterBackupsExecute provides a mock function with given fields: r
func (_m *SharedTierSnapshotsApi) ListSharedClusterBackupsExecute(r admin.ListSharedClusterBackupsApiRequest) (*admin.PaginatedTenantSnapshot, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListSharedClusterBackupsExecute")
	}

	var r0 *admin.PaginatedTenantSnapshot
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListSharedClusterBackupsApiRequest) (*admin.PaginatedTenantSnapshot, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListSharedClusterBackupsApiRequest) *admin.PaginatedTenantSnapshot); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedTenantSnapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListSharedClusterBackupsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListSharedClusterBackupsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSharedClusterBackupsExecute'
type SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call struct {
	*mock.Call
}

// ListSharedClusterBackupsExecute is a helper method to define mock.On call
//   - r admin.ListSharedClusterBackupsApiRequest
func (_e *SharedTierSnapshotsApi_Expecter) ListSharedClusterBackupsExecute(r any) *SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call {
	return &SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call{Call: _e.mock.On("ListSharedClusterBackupsExecute", r)}
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call) Run(run func(r admin.ListSharedClusterBackupsApiRequest)) *SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListSharedClusterBackupsApiRequest))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call) Return(_a0 *admin.PaginatedTenantSnapshot, _a1 *http.Response, _a2 error) *SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call) RunAndReturn(run func(admin.ListSharedClusterBackupsApiRequest) (*admin.PaginatedTenantSnapshot, *http.Response, error)) *SharedTierSnapshotsApi_ListSharedClusterBackupsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListSharedClusterBackupsWithParams provides a mock function with given fields: ctx, args
func (_m *SharedTierSnapshotsApi) ListSharedClusterBackupsWithParams(ctx context.Context, args *admin.ListSharedClusterBackupsApiParams) admin.ListSharedClusterBackupsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListSharedClusterBackupsWithParams")
	}

	var r0 admin.ListSharedClusterBackupsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListSharedClusterBackupsApiParams) admin.ListSharedClusterBackupsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListSharedClusterBackupsApiRequest)
	}

	return r0
}

// SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSharedClusterBackupsWithParams'
type SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call struct {
	*mock.Call
}

// ListSharedClusterBackupsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListSharedClusterBackupsApiParams
func (_e *SharedTierSnapshotsApi_Expecter) ListSharedClusterBackupsWithParams(ctx any, args any) *SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call {
	return &SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call{Call: _e.mock.On("ListSharedClusterBackupsWithParams", ctx, args)}
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListSharedClusterBackupsApiParams)) *SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListSharedClusterBackupsApiParams))
	})
	return _c
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call) Return(_a0 admin.ListSharedClusterBackupsApiRequest) *SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListSharedClusterBackupsApiParams) admin.ListSharedClusterBackupsApiRequest) *SharedTierSnapshotsApi_ListSharedClusterBackupsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewSharedTierSnapshotsApi creates a new instance of SharedTierSnapshotsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSharedTierSnapshotsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *SharedTierSnapshotsApi {
	mock := &SharedTierSnapshotsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
