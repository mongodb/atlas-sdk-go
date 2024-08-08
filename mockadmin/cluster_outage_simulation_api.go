// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240805001/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// ClusterOutageSimulationApi is an autogenerated mock type for the ClusterOutageSimulationApi type
type ClusterOutageSimulationApi struct {
	mock.Mock
}

type ClusterOutageSimulationApi_Expecter struct {
	mock *mock.Mock
}

func (_m *ClusterOutageSimulationApi) EXPECT() *ClusterOutageSimulationApi_Expecter {
	return &ClusterOutageSimulationApi_Expecter{mock: &_m.Mock}
}

// EndOutageSimulation provides a mock function with given fields: ctx, groupId, clusterName
func (_m *ClusterOutageSimulationApi) EndOutageSimulation(ctx context.Context, groupId string, clusterName string) admin.EndOutageSimulationApiRequest {
	ret := _m.Called(ctx, groupId, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for EndOutageSimulation")
	}

	var r0 admin.EndOutageSimulationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.EndOutageSimulationApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName)
	} else {
		r0 = ret.Get(0).(admin.EndOutageSimulationApiRequest)
	}

	return r0
}

// ClusterOutageSimulationApi_EndOutageSimulation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndOutageSimulation'
type ClusterOutageSimulationApi_EndOutageSimulation_Call struct {
	*mock.Call
}

// EndOutageSimulation is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
func (_e *ClusterOutageSimulationApi_Expecter) EndOutageSimulation(ctx interface{}, groupId interface{}, clusterName interface{}) *ClusterOutageSimulationApi_EndOutageSimulation_Call {
	return &ClusterOutageSimulationApi_EndOutageSimulation_Call{Call: _e.mock.On("EndOutageSimulation", ctx, groupId, clusterName)}
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulation_Call) Run(run func(ctx context.Context, groupId string, clusterName string)) *ClusterOutageSimulationApi_EndOutageSimulation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulation_Call) Return(_a0 admin.EndOutageSimulationApiRequest) *ClusterOutageSimulationApi_EndOutageSimulation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulation_Call) RunAndReturn(run func(context.Context, string, string) admin.EndOutageSimulationApiRequest) *ClusterOutageSimulationApi_EndOutageSimulation_Call {
	_c.Call.Return(run)
	return _c
}

// EndOutageSimulationExecute provides a mock function with given fields: r
func (_m *ClusterOutageSimulationApi) EndOutageSimulationExecute(r admin.EndOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for EndOutageSimulationExecute")
	}

	var r0 *admin.ClusterOutageSimulation
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.EndOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.EndOutageSimulationApiRequest) *admin.ClusterOutageSimulation); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ClusterOutageSimulation)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.EndOutageSimulationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.EndOutageSimulationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterOutageSimulationApi_EndOutageSimulationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndOutageSimulationExecute'
type ClusterOutageSimulationApi_EndOutageSimulationExecute_Call struct {
	*mock.Call
}

// EndOutageSimulationExecute is a helper method to define mock.On call
//   - r admin.EndOutageSimulationApiRequest
func (_e *ClusterOutageSimulationApi_Expecter) EndOutageSimulationExecute(r interface{}) *ClusterOutageSimulationApi_EndOutageSimulationExecute_Call {
	return &ClusterOutageSimulationApi_EndOutageSimulationExecute_Call{Call: _e.mock.On("EndOutageSimulationExecute", r)}
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulationExecute_Call) Run(run func(r admin.EndOutageSimulationApiRequest)) *ClusterOutageSimulationApi_EndOutageSimulationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.EndOutageSimulationApiRequest))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulationExecute_Call) Return(_a0 *admin.ClusterOutageSimulation, _a1 *http.Response, _a2 error) *ClusterOutageSimulationApi_EndOutageSimulationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulationExecute_Call) RunAndReturn(run func(admin.EndOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error)) *ClusterOutageSimulationApi_EndOutageSimulationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// EndOutageSimulationWithParams provides a mock function with given fields: ctx, args
func (_m *ClusterOutageSimulationApi) EndOutageSimulationWithParams(ctx context.Context, args *admin.EndOutageSimulationApiParams) admin.EndOutageSimulationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for EndOutageSimulationWithParams")
	}

	var r0 admin.EndOutageSimulationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.EndOutageSimulationApiParams) admin.EndOutageSimulationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.EndOutageSimulationApiRequest)
	}

	return r0
}

// ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EndOutageSimulationWithParams'
type ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call struct {
	*mock.Call
}

// EndOutageSimulationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.EndOutageSimulationApiParams
func (_e *ClusterOutageSimulationApi_Expecter) EndOutageSimulationWithParams(ctx interface{}, args interface{}) *ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call {
	return &ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call{Call: _e.mock.On("EndOutageSimulationWithParams", ctx, args)}
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call) Run(run func(ctx context.Context, args *admin.EndOutageSimulationApiParams)) *ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.EndOutageSimulationApiParams))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call) Return(_a0 admin.EndOutageSimulationApiRequest) *ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call) RunAndReturn(run func(context.Context, *admin.EndOutageSimulationApiParams) admin.EndOutageSimulationApiRequest) *ClusterOutageSimulationApi_EndOutageSimulationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetOutageSimulation provides a mock function with given fields: ctx, groupId, clusterName
func (_m *ClusterOutageSimulationApi) GetOutageSimulation(ctx context.Context, groupId string, clusterName string) admin.GetOutageSimulationApiRequest {
	ret := _m.Called(ctx, groupId, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetOutageSimulation")
	}

	var r0 admin.GetOutageSimulationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetOutageSimulationApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName)
	} else {
		r0 = ret.Get(0).(admin.GetOutageSimulationApiRequest)
	}

	return r0
}

// ClusterOutageSimulationApi_GetOutageSimulation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOutageSimulation'
type ClusterOutageSimulationApi_GetOutageSimulation_Call struct {
	*mock.Call
}

// GetOutageSimulation is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
func (_e *ClusterOutageSimulationApi_Expecter) GetOutageSimulation(ctx interface{}, groupId interface{}, clusterName interface{}) *ClusterOutageSimulationApi_GetOutageSimulation_Call {
	return &ClusterOutageSimulationApi_GetOutageSimulation_Call{Call: _e.mock.On("GetOutageSimulation", ctx, groupId, clusterName)}
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulation_Call) Run(run func(ctx context.Context, groupId string, clusterName string)) *ClusterOutageSimulationApi_GetOutageSimulation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulation_Call) Return(_a0 admin.GetOutageSimulationApiRequest) *ClusterOutageSimulationApi_GetOutageSimulation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulation_Call) RunAndReturn(run func(context.Context, string, string) admin.GetOutageSimulationApiRequest) *ClusterOutageSimulationApi_GetOutageSimulation_Call {
	_c.Call.Return(run)
	return _c
}

// GetOutageSimulationExecute provides a mock function with given fields: r
func (_m *ClusterOutageSimulationApi) GetOutageSimulationExecute(r admin.GetOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetOutageSimulationExecute")
	}

	var r0 *admin.ClusterOutageSimulation
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetOutageSimulationApiRequest) *admin.ClusterOutageSimulation); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ClusterOutageSimulation)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetOutageSimulationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetOutageSimulationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterOutageSimulationApi_GetOutageSimulationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOutageSimulationExecute'
type ClusterOutageSimulationApi_GetOutageSimulationExecute_Call struct {
	*mock.Call
}

// GetOutageSimulationExecute is a helper method to define mock.On call
//   - r admin.GetOutageSimulationApiRequest
func (_e *ClusterOutageSimulationApi_Expecter) GetOutageSimulationExecute(r interface{}) *ClusterOutageSimulationApi_GetOutageSimulationExecute_Call {
	return &ClusterOutageSimulationApi_GetOutageSimulationExecute_Call{Call: _e.mock.On("GetOutageSimulationExecute", r)}
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulationExecute_Call) Run(run func(r admin.GetOutageSimulationApiRequest)) *ClusterOutageSimulationApi_GetOutageSimulationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetOutageSimulationApiRequest))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulationExecute_Call) Return(_a0 *admin.ClusterOutageSimulation, _a1 *http.Response, _a2 error) *ClusterOutageSimulationApi_GetOutageSimulationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulationExecute_Call) RunAndReturn(run func(admin.GetOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error)) *ClusterOutageSimulationApi_GetOutageSimulationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetOutageSimulationWithParams provides a mock function with given fields: ctx, args
func (_m *ClusterOutageSimulationApi) GetOutageSimulationWithParams(ctx context.Context, args *admin.GetOutageSimulationApiParams) admin.GetOutageSimulationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetOutageSimulationWithParams")
	}

	var r0 admin.GetOutageSimulationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetOutageSimulationApiParams) admin.GetOutageSimulationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetOutageSimulationApiRequest)
	}

	return r0
}

// ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOutageSimulationWithParams'
type ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call struct {
	*mock.Call
}

// GetOutageSimulationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetOutageSimulationApiParams
func (_e *ClusterOutageSimulationApi_Expecter) GetOutageSimulationWithParams(ctx interface{}, args interface{}) *ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call {
	return &ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call{Call: _e.mock.On("GetOutageSimulationWithParams", ctx, args)}
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call) Run(run func(ctx context.Context, args *admin.GetOutageSimulationApiParams)) *ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetOutageSimulationApiParams))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call) Return(_a0 admin.GetOutageSimulationApiRequest) *ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetOutageSimulationApiParams) admin.GetOutageSimulationApiRequest) *ClusterOutageSimulationApi_GetOutageSimulationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// StartOutageSimulation provides a mock function with given fields: ctx, groupId, clusterName, clusterOutageSimulation
func (_m *ClusterOutageSimulationApi) StartOutageSimulation(ctx context.Context, groupId string, clusterName string, clusterOutageSimulation *admin.ClusterOutageSimulation) admin.StartOutageSimulationApiRequest {
	ret := _m.Called(ctx, groupId, clusterName, clusterOutageSimulation)

	if len(ret) == 0 {
		panic("no return value specified for StartOutageSimulation")
	}

	var r0 admin.StartOutageSimulationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.ClusterOutageSimulation) admin.StartOutageSimulationApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName, clusterOutageSimulation)
	} else {
		r0 = ret.Get(0).(admin.StartOutageSimulationApiRequest)
	}

	return r0
}

// ClusterOutageSimulationApi_StartOutageSimulation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartOutageSimulation'
type ClusterOutageSimulationApi_StartOutageSimulation_Call struct {
	*mock.Call
}

// StartOutageSimulation is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
//   - clusterOutageSimulation *admin.ClusterOutageSimulation
func (_e *ClusterOutageSimulationApi_Expecter) StartOutageSimulation(ctx interface{}, groupId interface{}, clusterName interface{}, clusterOutageSimulation interface{}) *ClusterOutageSimulationApi_StartOutageSimulation_Call {
	return &ClusterOutageSimulationApi_StartOutageSimulation_Call{Call: _e.mock.On("StartOutageSimulation", ctx, groupId, clusterName, clusterOutageSimulation)}
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulation_Call) Run(run func(ctx context.Context, groupId string, clusterName string, clusterOutageSimulation *admin.ClusterOutageSimulation)) *ClusterOutageSimulationApi_StartOutageSimulation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.ClusterOutageSimulation))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulation_Call) Return(_a0 admin.StartOutageSimulationApiRequest) *ClusterOutageSimulationApi_StartOutageSimulation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulation_Call) RunAndReturn(run func(context.Context, string, string, *admin.ClusterOutageSimulation) admin.StartOutageSimulationApiRequest) *ClusterOutageSimulationApi_StartOutageSimulation_Call {
	_c.Call.Return(run)
	return _c
}

// StartOutageSimulationExecute provides a mock function with given fields: r
func (_m *ClusterOutageSimulationApi) StartOutageSimulationExecute(r admin.StartOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for StartOutageSimulationExecute")
	}

	var r0 *admin.ClusterOutageSimulation
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.StartOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.StartOutageSimulationApiRequest) *admin.ClusterOutageSimulation); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ClusterOutageSimulation)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.StartOutageSimulationApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.StartOutageSimulationApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterOutageSimulationApi_StartOutageSimulationExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartOutageSimulationExecute'
type ClusterOutageSimulationApi_StartOutageSimulationExecute_Call struct {
	*mock.Call
}

// StartOutageSimulationExecute is a helper method to define mock.On call
//   - r admin.StartOutageSimulationApiRequest
func (_e *ClusterOutageSimulationApi_Expecter) StartOutageSimulationExecute(r interface{}) *ClusterOutageSimulationApi_StartOutageSimulationExecute_Call {
	return &ClusterOutageSimulationApi_StartOutageSimulationExecute_Call{Call: _e.mock.On("StartOutageSimulationExecute", r)}
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulationExecute_Call) Run(run func(r admin.StartOutageSimulationApiRequest)) *ClusterOutageSimulationApi_StartOutageSimulationExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.StartOutageSimulationApiRequest))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulationExecute_Call) Return(_a0 *admin.ClusterOutageSimulation, _a1 *http.Response, _a2 error) *ClusterOutageSimulationApi_StartOutageSimulationExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulationExecute_Call) RunAndReturn(run func(admin.StartOutageSimulationApiRequest) (*admin.ClusterOutageSimulation, *http.Response, error)) *ClusterOutageSimulationApi_StartOutageSimulationExecute_Call {
	_c.Call.Return(run)
	return _c
}

// StartOutageSimulationWithParams provides a mock function with given fields: ctx, args
func (_m *ClusterOutageSimulationApi) StartOutageSimulationWithParams(ctx context.Context, args *admin.StartOutageSimulationApiParams) admin.StartOutageSimulationApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for StartOutageSimulationWithParams")
	}

	var r0 admin.StartOutageSimulationApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.StartOutageSimulationApiParams) admin.StartOutageSimulationApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.StartOutageSimulationApiRequest)
	}

	return r0
}

// ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartOutageSimulationWithParams'
type ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call struct {
	*mock.Call
}

// StartOutageSimulationWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.StartOutageSimulationApiParams
func (_e *ClusterOutageSimulationApi_Expecter) StartOutageSimulationWithParams(ctx interface{}, args interface{}) *ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call {
	return &ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call{Call: _e.mock.On("StartOutageSimulationWithParams", ctx, args)}
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call) Run(run func(ctx context.Context, args *admin.StartOutageSimulationApiParams)) *ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.StartOutageSimulationApiParams))
	})
	return _c
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call) Return(_a0 admin.StartOutageSimulationApiRequest) *ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call) RunAndReturn(run func(context.Context, *admin.StartOutageSimulationApiParams) admin.StartOutageSimulationApiRequest) *ClusterOutageSimulationApi_StartOutageSimulationWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewClusterOutageSimulationApi creates a new instance of ClusterOutageSimulationApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClusterOutageSimulationApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClusterOutageSimulationApi {
	mock := &ClusterOutageSimulationApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
