// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20241023002/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// GlobalClustersApi is an autogenerated mock type for the GlobalClustersApi type
type GlobalClustersApi struct {
	mock.Mock
}

type GlobalClustersApi_Expecter struct {
	mock *mock.Mock
}

func (_m *GlobalClustersApi) EXPECT() *GlobalClustersApi_Expecter {
	return &GlobalClustersApi_Expecter{mock: &_m.Mock}
}

// CreateCustomZoneMapping provides a mock function with given fields: ctx, groupId, clusterName, customZoneMappings
func (_m *GlobalClustersApi) CreateCustomZoneMapping(ctx context.Context, groupId string, clusterName string, customZoneMappings *admin.CustomZoneMappings) admin.CreateCustomZoneMappingApiRequest {
	ret := _m.Called(ctx, groupId, clusterName, customZoneMappings)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomZoneMapping")
	}

	var r0 admin.CreateCustomZoneMappingApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.CustomZoneMappings) admin.CreateCustomZoneMappingApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName, customZoneMappings)
	} else {
		r0 = ret.Get(0).(admin.CreateCustomZoneMappingApiRequest)
	}

	return r0
}

// GlobalClustersApi_CreateCustomZoneMapping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomZoneMapping'
type GlobalClustersApi_CreateCustomZoneMapping_Call struct {
	*mock.Call
}

// CreateCustomZoneMapping is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
//   - customZoneMappings *admin.CustomZoneMappings
func (_e *GlobalClustersApi_Expecter) CreateCustomZoneMapping(ctx any, groupId any, clusterName any, customZoneMappings any) *GlobalClustersApi_CreateCustomZoneMapping_Call {
	return &GlobalClustersApi_CreateCustomZoneMapping_Call{Call: _e.mock.On("CreateCustomZoneMapping", ctx, groupId, clusterName, customZoneMappings)}
}

func (_c *GlobalClustersApi_CreateCustomZoneMapping_Call) Run(run func(ctx context.Context, groupId string, clusterName string, customZoneMappings *admin.CustomZoneMappings)) *GlobalClustersApi_CreateCustomZoneMapping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.CustomZoneMappings))
	})
	return _c
}

func (_c *GlobalClustersApi_CreateCustomZoneMapping_Call) Return(_a0 admin.CreateCustomZoneMappingApiRequest) *GlobalClustersApi_CreateCustomZoneMapping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_CreateCustomZoneMapping_Call) RunAndReturn(run func(context.Context, string, string, *admin.CustomZoneMappings) admin.CreateCustomZoneMappingApiRequest) *GlobalClustersApi_CreateCustomZoneMapping_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCustomZoneMappingExecute provides a mock function with given fields: r
func (_m *GlobalClustersApi) CreateCustomZoneMappingExecute(r admin.CreateCustomZoneMappingApiRequest) (*admin.GeoSharding20240805, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomZoneMappingExecute")
	}

	var r0 *admin.GeoSharding20240805
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateCustomZoneMappingApiRequest) (*admin.GeoSharding20240805, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateCustomZoneMappingApiRequest) *admin.GeoSharding20240805); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GeoSharding20240805)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateCustomZoneMappingApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateCustomZoneMappingApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GlobalClustersApi_CreateCustomZoneMappingExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomZoneMappingExecute'
type GlobalClustersApi_CreateCustomZoneMappingExecute_Call struct {
	*mock.Call
}

// CreateCustomZoneMappingExecute is a helper method to define mock.On call
//   - r admin.CreateCustomZoneMappingApiRequest
func (_e *GlobalClustersApi_Expecter) CreateCustomZoneMappingExecute(r any) *GlobalClustersApi_CreateCustomZoneMappingExecute_Call {
	return &GlobalClustersApi_CreateCustomZoneMappingExecute_Call{Call: _e.mock.On("CreateCustomZoneMappingExecute", r)}
}

func (_c *GlobalClustersApi_CreateCustomZoneMappingExecute_Call) Run(run func(r admin.CreateCustomZoneMappingApiRequest)) *GlobalClustersApi_CreateCustomZoneMappingExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateCustomZoneMappingApiRequest))
	})
	return _c
}

func (_c *GlobalClustersApi_CreateCustomZoneMappingExecute_Call) Return(_a0 *admin.GeoSharding20240805, _a1 *http.Response, _a2 error) *GlobalClustersApi_CreateCustomZoneMappingExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GlobalClustersApi_CreateCustomZoneMappingExecute_Call) RunAndReturn(run func(admin.CreateCustomZoneMappingApiRequest) (*admin.GeoSharding20240805, *http.Response, error)) *GlobalClustersApi_CreateCustomZoneMappingExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCustomZoneMappingWithParams provides a mock function with given fields: ctx, args
func (_m *GlobalClustersApi) CreateCustomZoneMappingWithParams(ctx context.Context, args *admin.CreateCustomZoneMappingApiParams) admin.CreateCustomZoneMappingApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomZoneMappingWithParams")
	}

	var r0 admin.CreateCustomZoneMappingApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateCustomZoneMappingApiParams) admin.CreateCustomZoneMappingApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateCustomZoneMappingApiRequest)
	}

	return r0
}

// GlobalClustersApi_CreateCustomZoneMappingWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomZoneMappingWithParams'
type GlobalClustersApi_CreateCustomZoneMappingWithParams_Call struct {
	*mock.Call
}

// CreateCustomZoneMappingWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateCustomZoneMappingApiParams
func (_e *GlobalClustersApi_Expecter) CreateCustomZoneMappingWithParams(ctx any, args any) *GlobalClustersApi_CreateCustomZoneMappingWithParams_Call {
	return &GlobalClustersApi_CreateCustomZoneMappingWithParams_Call{Call: _e.mock.On("CreateCustomZoneMappingWithParams", ctx, args)}
}

func (_c *GlobalClustersApi_CreateCustomZoneMappingWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateCustomZoneMappingApiParams)) *GlobalClustersApi_CreateCustomZoneMappingWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateCustomZoneMappingApiParams))
	})
	return _c
}

func (_c *GlobalClustersApi_CreateCustomZoneMappingWithParams_Call) Return(_a0 admin.CreateCustomZoneMappingApiRequest) *GlobalClustersApi_CreateCustomZoneMappingWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_CreateCustomZoneMappingWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateCustomZoneMappingApiParams) admin.CreateCustomZoneMappingApiRequest) *GlobalClustersApi_CreateCustomZoneMappingWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// CreateManagedNamespace provides a mock function with given fields: ctx, groupId, clusterName, managedNamespaces
func (_m *GlobalClustersApi) CreateManagedNamespace(ctx context.Context, groupId string, clusterName string, managedNamespaces *admin.ManagedNamespaces) admin.CreateManagedNamespaceApiRequest {
	ret := _m.Called(ctx, groupId, clusterName, managedNamespaces)

	if len(ret) == 0 {
		panic("no return value specified for CreateManagedNamespace")
	}

	var r0 admin.CreateManagedNamespaceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.ManagedNamespaces) admin.CreateManagedNamespaceApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName, managedNamespaces)
	} else {
		r0 = ret.Get(0).(admin.CreateManagedNamespaceApiRequest)
	}

	return r0
}

// GlobalClustersApi_CreateManagedNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateManagedNamespace'
type GlobalClustersApi_CreateManagedNamespace_Call struct {
	*mock.Call
}

// CreateManagedNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
//   - managedNamespaces *admin.ManagedNamespaces
func (_e *GlobalClustersApi_Expecter) CreateManagedNamespace(ctx any, groupId any, clusterName any, managedNamespaces any) *GlobalClustersApi_CreateManagedNamespace_Call {
	return &GlobalClustersApi_CreateManagedNamespace_Call{Call: _e.mock.On("CreateManagedNamespace", ctx, groupId, clusterName, managedNamespaces)}
}

func (_c *GlobalClustersApi_CreateManagedNamespace_Call) Run(run func(ctx context.Context, groupId string, clusterName string, managedNamespaces *admin.ManagedNamespaces)) *GlobalClustersApi_CreateManagedNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.ManagedNamespaces))
	})
	return _c
}

func (_c *GlobalClustersApi_CreateManagedNamespace_Call) Return(_a0 admin.CreateManagedNamespaceApiRequest) *GlobalClustersApi_CreateManagedNamespace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_CreateManagedNamespace_Call) RunAndReturn(run func(context.Context, string, string, *admin.ManagedNamespaces) admin.CreateManagedNamespaceApiRequest) *GlobalClustersApi_CreateManagedNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// CreateManagedNamespaceExecute provides a mock function with given fields: r
func (_m *GlobalClustersApi) CreateManagedNamespaceExecute(r admin.CreateManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateManagedNamespaceExecute")
	}

	var r0 *admin.GeoSharding20240805
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.CreateManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateManagedNamespaceApiRequest) *admin.GeoSharding20240805); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GeoSharding20240805)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateManagedNamespaceApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.CreateManagedNamespaceApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GlobalClustersApi_CreateManagedNamespaceExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateManagedNamespaceExecute'
type GlobalClustersApi_CreateManagedNamespaceExecute_Call struct {
	*mock.Call
}

// CreateManagedNamespaceExecute is a helper method to define mock.On call
//   - r admin.CreateManagedNamespaceApiRequest
func (_e *GlobalClustersApi_Expecter) CreateManagedNamespaceExecute(r any) *GlobalClustersApi_CreateManagedNamespaceExecute_Call {
	return &GlobalClustersApi_CreateManagedNamespaceExecute_Call{Call: _e.mock.On("CreateManagedNamespaceExecute", r)}
}

func (_c *GlobalClustersApi_CreateManagedNamespaceExecute_Call) Run(run func(r admin.CreateManagedNamespaceApiRequest)) *GlobalClustersApi_CreateManagedNamespaceExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateManagedNamespaceApiRequest))
	})
	return _c
}

func (_c *GlobalClustersApi_CreateManagedNamespaceExecute_Call) Return(_a0 *admin.GeoSharding20240805, _a1 *http.Response, _a2 error) *GlobalClustersApi_CreateManagedNamespaceExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GlobalClustersApi_CreateManagedNamespaceExecute_Call) RunAndReturn(run func(admin.CreateManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error)) *GlobalClustersApi_CreateManagedNamespaceExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateManagedNamespaceWithParams provides a mock function with given fields: ctx, args
func (_m *GlobalClustersApi) CreateManagedNamespaceWithParams(ctx context.Context, args *admin.CreateManagedNamespaceApiParams) admin.CreateManagedNamespaceApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateManagedNamespaceWithParams")
	}

	var r0 admin.CreateManagedNamespaceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateManagedNamespaceApiParams) admin.CreateManagedNamespaceApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateManagedNamespaceApiRequest)
	}

	return r0
}

// GlobalClustersApi_CreateManagedNamespaceWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateManagedNamespaceWithParams'
type GlobalClustersApi_CreateManagedNamespaceWithParams_Call struct {
	*mock.Call
}

// CreateManagedNamespaceWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateManagedNamespaceApiParams
func (_e *GlobalClustersApi_Expecter) CreateManagedNamespaceWithParams(ctx any, args any) *GlobalClustersApi_CreateManagedNamespaceWithParams_Call {
	return &GlobalClustersApi_CreateManagedNamespaceWithParams_Call{Call: _e.mock.On("CreateManagedNamespaceWithParams", ctx, args)}
}

func (_c *GlobalClustersApi_CreateManagedNamespaceWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateManagedNamespaceApiParams)) *GlobalClustersApi_CreateManagedNamespaceWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateManagedNamespaceApiParams))
	})
	return _c
}

func (_c *GlobalClustersApi_CreateManagedNamespaceWithParams_Call) Return(_a0 admin.CreateManagedNamespaceApiRequest) *GlobalClustersApi_CreateManagedNamespaceWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_CreateManagedNamespaceWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateManagedNamespaceApiParams) admin.CreateManagedNamespaceApiRequest) *GlobalClustersApi_CreateManagedNamespaceWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAllCustomZoneMappings provides a mock function with given fields: ctx, groupId, clusterName
func (_m *GlobalClustersApi) DeleteAllCustomZoneMappings(ctx context.Context, groupId string, clusterName string) admin.DeleteAllCustomZoneMappingsApiRequest {
	ret := _m.Called(ctx, groupId, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllCustomZoneMappings")
	}

	var r0 admin.DeleteAllCustomZoneMappingsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.DeleteAllCustomZoneMappingsApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName)
	} else {
		r0 = ret.Get(0).(admin.DeleteAllCustomZoneMappingsApiRequest)
	}

	return r0
}

// GlobalClustersApi_DeleteAllCustomZoneMappings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAllCustomZoneMappings'
type GlobalClustersApi_DeleteAllCustomZoneMappings_Call struct {
	*mock.Call
}

// DeleteAllCustomZoneMappings is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
func (_e *GlobalClustersApi_Expecter) DeleteAllCustomZoneMappings(ctx any, groupId any, clusterName any) *GlobalClustersApi_DeleteAllCustomZoneMappings_Call {
	return &GlobalClustersApi_DeleteAllCustomZoneMappings_Call{Call: _e.mock.On("DeleteAllCustomZoneMappings", ctx, groupId, clusterName)}
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappings_Call) Run(run func(ctx context.Context, groupId string, clusterName string)) *GlobalClustersApi_DeleteAllCustomZoneMappings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappings_Call) Return(_a0 admin.DeleteAllCustomZoneMappingsApiRequest) *GlobalClustersApi_DeleteAllCustomZoneMappings_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappings_Call) RunAndReturn(run func(context.Context, string, string) admin.DeleteAllCustomZoneMappingsApiRequest) *GlobalClustersApi_DeleteAllCustomZoneMappings_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAllCustomZoneMappingsExecute provides a mock function with given fields: r
func (_m *GlobalClustersApi) DeleteAllCustomZoneMappingsExecute(r admin.DeleteAllCustomZoneMappingsApiRequest) (*admin.GeoSharding20240805, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllCustomZoneMappingsExecute")
	}

	var r0 *admin.GeoSharding20240805
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DeleteAllCustomZoneMappingsApiRequest) (*admin.GeoSharding20240805, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteAllCustomZoneMappingsApiRequest) *admin.GeoSharding20240805); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GeoSharding20240805)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteAllCustomZoneMappingsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DeleteAllCustomZoneMappingsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAllCustomZoneMappingsExecute'
type GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call struct {
	*mock.Call
}

// DeleteAllCustomZoneMappingsExecute is a helper method to define mock.On call
//   - r admin.DeleteAllCustomZoneMappingsApiRequest
func (_e *GlobalClustersApi_Expecter) DeleteAllCustomZoneMappingsExecute(r any) *GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call {
	return &GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call{Call: _e.mock.On("DeleteAllCustomZoneMappingsExecute", r)}
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call) Run(run func(r admin.DeleteAllCustomZoneMappingsApiRequest)) *GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteAllCustomZoneMappingsApiRequest))
	})
	return _c
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call) Return(_a0 *admin.GeoSharding20240805, _a1 *http.Response, _a2 error) *GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call) RunAndReturn(run func(admin.DeleteAllCustomZoneMappingsApiRequest) (*admin.GeoSharding20240805, *http.Response, error)) *GlobalClustersApi_DeleteAllCustomZoneMappingsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAllCustomZoneMappingsWithParams provides a mock function with given fields: ctx, args
func (_m *GlobalClustersApi) DeleteAllCustomZoneMappingsWithParams(ctx context.Context, args *admin.DeleteAllCustomZoneMappingsApiParams) admin.DeleteAllCustomZoneMappingsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllCustomZoneMappingsWithParams")
	}

	var r0 admin.DeleteAllCustomZoneMappingsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteAllCustomZoneMappingsApiParams) admin.DeleteAllCustomZoneMappingsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteAllCustomZoneMappingsApiRequest)
	}

	return r0
}

// GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAllCustomZoneMappingsWithParams'
type GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call struct {
	*mock.Call
}

// DeleteAllCustomZoneMappingsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteAllCustomZoneMappingsApiParams
func (_e *GlobalClustersApi_Expecter) DeleteAllCustomZoneMappingsWithParams(ctx any, args any) *GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call {
	return &GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call{Call: _e.mock.On("DeleteAllCustomZoneMappingsWithParams", ctx, args)}
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteAllCustomZoneMappingsApiParams)) *GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteAllCustomZoneMappingsApiParams))
	})
	return _c
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call) Return(_a0 admin.DeleteAllCustomZoneMappingsApiRequest) *GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteAllCustomZoneMappingsApiParams) admin.DeleteAllCustomZoneMappingsApiRequest) *GlobalClustersApi_DeleteAllCustomZoneMappingsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteManagedNamespace provides a mock function with given fields: ctx, clusterName, groupId
func (_m *GlobalClustersApi) DeleteManagedNamespace(ctx context.Context, clusterName string, groupId string) admin.DeleteManagedNamespaceApiRequest {
	ret := _m.Called(ctx, clusterName, groupId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteManagedNamespace")
	}

	var r0 admin.DeleteManagedNamespaceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.DeleteManagedNamespaceApiRequest); ok {
		r0 = rf(ctx, clusterName, groupId)
	} else {
		r0 = ret.Get(0).(admin.DeleteManagedNamespaceApiRequest)
	}

	return r0
}

// GlobalClustersApi_DeleteManagedNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteManagedNamespace'
type GlobalClustersApi_DeleteManagedNamespace_Call struct {
	*mock.Call
}

// DeleteManagedNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - groupId string
func (_e *GlobalClustersApi_Expecter) DeleteManagedNamespace(ctx any, clusterName any, groupId any) *GlobalClustersApi_DeleteManagedNamespace_Call {
	return &GlobalClustersApi_DeleteManagedNamespace_Call{Call: _e.mock.On("DeleteManagedNamespace", ctx, clusterName, groupId)}
}

func (_c *GlobalClustersApi_DeleteManagedNamespace_Call) Run(run func(ctx context.Context, clusterName string, groupId string)) *GlobalClustersApi_DeleteManagedNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *GlobalClustersApi_DeleteManagedNamespace_Call) Return(_a0 admin.DeleteManagedNamespaceApiRequest) *GlobalClustersApi_DeleteManagedNamespace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_DeleteManagedNamespace_Call) RunAndReturn(run func(context.Context, string, string) admin.DeleteManagedNamespaceApiRequest) *GlobalClustersApi_DeleteManagedNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteManagedNamespaceExecute provides a mock function with given fields: r
func (_m *GlobalClustersApi) DeleteManagedNamespaceExecute(r admin.DeleteManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for DeleteManagedNamespaceExecute")
	}

	var r0 *admin.GeoSharding20240805
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.DeleteManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.DeleteManagedNamespaceApiRequest) *admin.GeoSharding20240805); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GeoSharding20240805)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.DeleteManagedNamespaceApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.DeleteManagedNamespaceApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GlobalClustersApi_DeleteManagedNamespaceExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteManagedNamespaceExecute'
type GlobalClustersApi_DeleteManagedNamespaceExecute_Call struct {
	*mock.Call
}

// DeleteManagedNamespaceExecute is a helper method to define mock.On call
//   - r admin.DeleteManagedNamespaceApiRequest
func (_e *GlobalClustersApi_Expecter) DeleteManagedNamespaceExecute(r any) *GlobalClustersApi_DeleteManagedNamespaceExecute_Call {
	return &GlobalClustersApi_DeleteManagedNamespaceExecute_Call{Call: _e.mock.On("DeleteManagedNamespaceExecute", r)}
}

func (_c *GlobalClustersApi_DeleteManagedNamespaceExecute_Call) Run(run func(r admin.DeleteManagedNamespaceApiRequest)) *GlobalClustersApi_DeleteManagedNamespaceExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.DeleteManagedNamespaceApiRequest))
	})
	return _c
}

func (_c *GlobalClustersApi_DeleteManagedNamespaceExecute_Call) Return(_a0 *admin.GeoSharding20240805, _a1 *http.Response, _a2 error) *GlobalClustersApi_DeleteManagedNamespaceExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GlobalClustersApi_DeleteManagedNamespaceExecute_Call) RunAndReturn(run func(admin.DeleteManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error)) *GlobalClustersApi_DeleteManagedNamespaceExecute_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteManagedNamespaceWithParams provides a mock function with given fields: ctx, args
func (_m *GlobalClustersApi) DeleteManagedNamespaceWithParams(ctx context.Context, args *admin.DeleteManagedNamespaceApiParams) admin.DeleteManagedNamespaceApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for DeleteManagedNamespaceWithParams")
	}

	var r0 admin.DeleteManagedNamespaceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.DeleteManagedNamespaceApiParams) admin.DeleteManagedNamespaceApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.DeleteManagedNamespaceApiRequest)
	}

	return r0
}

// GlobalClustersApi_DeleteManagedNamespaceWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteManagedNamespaceWithParams'
type GlobalClustersApi_DeleteManagedNamespaceWithParams_Call struct {
	*mock.Call
}

// DeleteManagedNamespaceWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.DeleteManagedNamespaceApiParams
func (_e *GlobalClustersApi_Expecter) DeleteManagedNamespaceWithParams(ctx any, args any) *GlobalClustersApi_DeleteManagedNamespaceWithParams_Call {
	return &GlobalClustersApi_DeleteManagedNamespaceWithParams_Call{Call: _e.mock.On("DeleteManagedNamespaceWithParams", ctx, args)}
}

func (_c *GlobalClustersApi_DeleteManagedNamespaceWithParams_Call) Run(run func(ctx context.Context, args *admin.DeleteManagedNamespaceApiParams)) *GlobalClustersApi_DeleteManagedNamespaceWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.DeleteManagedNamespaceApiParams))
	})
	return _c
}

func (_c *GlobalClustersApi_DeleteManagedNamespaceWithParams_Call) Return(_a0 admin.DeleteManagedNamespaceApiRequest) *GlobalClustersApi_DeleteManagedNamespaceWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_DeleteManagedNamespaceWithParams_Call) RunAndReturn(run func(context.Context, *admin.DeleteManagedNamespaceApiParams) admin.DeleteManagedNamespaceApiRequest) *GlobalClustersApi_DeleteManagedNamespaceWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetManagedNamespace provides a mock function with given fields: ctx, groupId, clusterName
func (_m *GlobalClustersApi) GetManagedNamespace(ctx context.Context, groupId string, clusterName string) admin.GetManagedNamespaceApiRequest {
	ret := _m.Called(ctx, groupId, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetManagedNamespace")
	}

	var r0 admin.GetManagedNamespaceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetManagedNamespaceApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName)
	} else {
		r0 = ret.Get(0).(admin.GetManagedNamespaceApiRequest)
	}

	return r0
}

// GlobalClustersApi_GetManagedNamespace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManagedNamespace'
type GlobalClustersApi_GetManagedNamespace_Call struct {
	*mock.Call
}

// GetManagedNamespace is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
func (_e *GlobalClustersApi_Expecter) GetManagedNamespace(ctx any, groupId any, clusterName any) *GlobalClustersApi_GetManagedNamespace_Call {
	return &GlobalClustersApi_GetManagedNamespace_Call{Call: _e.mock.On("GetManagedNamespace", ctx, groupId, clusterName)}
}

func (_c *GlobalClustersApi_GetManagedNamespace_Call) Run(run func(ctx context.Context, groupId string, clusterName string)) *GlobalClustersApi_GetManagedNamespace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *GlobalClustersApi_GetManagedNamespace_Call) Return(_a0 admin.GetManagedNamespaceApiRequest) *GlobalClustersApi_GetManagedNamespace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_GetManagedNamespace_Call) RunAndReturn(run func(context.Context, string, string) admin.GetManagedNamespaceApiRequest) *GlobalClustersApi_GetManagedNamespace_Call {
	_c.Call.Return(run)
	return _c
}

// GetManagedNamespaceExecute provides a mock function with given fields: r
func (_m *GlobalClustersApi) GetManagedNamespaceExecute(r admin.GetManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetManagedNamespaceExecute")
	}

	var r0 *admin.GeoSharding20240805
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetManagedNamespaceApiRequest) *admin.GeoSharding20240805); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GeoSharding20240805)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetManagedNamespaceApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetManagedNamespaceApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GlobalClustersApi_GetManagedNamespaceExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManagedNamespaceExecute'
type GlobalClustersApi_GetManagedNamespaceExecute_Call struct {
	*mock.Call
}

// GetManagedNamespaceExecute is a helper method to define mock.On call
//   - r admin.GetManagedNamespaceApiRequest
func (_e *GlobalClustersApi_Expecter) GetManagedNamespaceExecute(r any) *GlobalClustersApi_GetManagedNamespaceExecute_Call {
	return &GlobalClustersApi_GetManagedNamespaceExecute_Call{Call: _e.mock.On("GetManagedNamespaceExecute", r)}
}

func (_c *GlobalClustersApi_GetManagedNamespaceExecute_Call) Run(run func(r admin.GetManagedNamespaceApiRequest)) *GlobalClustersApi_GetManagedNamespaceExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetManagedNamespaceApiRequest))
	})
	return _c
}

func (_c *GlobalClustersApi_GetManagedNamespaceExecute_Call) Return(_a0 *admin.GeoSharding20240805, _a1 *http.Response, _a2 error) *GlobalClustersApi_GetManagedNamespaceExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GlobalClustersApi_GetManagedNamespaceExecute_Call) RunAndReturn(run func(admin.GetManagedNamespaceApiRequest) (*admin.GeoSharding20240805, *http.Response, error)) *GlobalClustersApi_GetManagedNamespaceExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetManagedNamespaceWithParams provides a mock function with given fields: ctx, args
func (_m *GlobalClustersApi) GetManagedNamespaceWithParams(ctx context.Context, args *admin.GetManagedNamespaceApiParams) admin.GetManagedNamespaceApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetManagedNamespaceWithParams")
	}

	var r0 admin.GetManagedNamespaceApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetManagedNamespaceApiParams) admin.GetManagedNamespaceApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetManagedNamespaceApiRequest)
	}

	return r0
}

// GlobalClustersApi_GetManagedNamespaceWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManagedNamespaceWithParams'
type GlobalClustersApi_GetManagedNamespaceWithParams_Call struct {
	*mock.Call
}

// GetManagedNamespaceWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetManagedNamespaceApiParams
func (_e *GlobalClustersApi_Expecter) GetManagedNamespaceWithParams(ctx any, args any) *GlobalClustersApi_GetManagedNamespaceWithParams_Call {
	return &GlobalClustersApi_GetManagedNamespaceWithParams_Call{Call: _e.mock.On("GetManagedNamespaceWithParams", ctx, args)}
}

func (_c *GlobalClustersApi_GetManagedNamespaceWithParams_Call) Run(run func(ctx context.Context, args *admin.GetManagedNamespaceApiParams)) *GlobalClustersApi_GetManagedNamespaceWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetManagedNamespaceApiParams))
	})
	return _c
}

func (_c *GlobalClustersApi_GetManagedNamespaceWithParams_Call) Return(_a0 admin.GetManagedNamespaceApiRequest) *GlobalClustersApi_GetManagedNamespaceWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GlobalClustersApi_GetManagedNamespaceWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetManagedNamespaceApiParams) admin.GetManagedNamespaceApiRequest) *GlobalClustersApi_GetManagedNamespaceWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewGlobalClustersApi creates a new instance of GlobalClustersApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGlobalClustersApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *GlobalClustersApi {
	mock := &GlobalClustersApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
