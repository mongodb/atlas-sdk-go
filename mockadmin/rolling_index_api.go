// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20240805003/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RollingIndexApi is an autogenerated mock type for the RollingIndexApi type
type RollingIndexApi struct {
	mock.Mock
}

type RollingIndexApi_Expecter struct {
	mock *mock.Mock
}

func (_m *RollingIndexApi) EXPECT() *RollingIndexApi_Expecter {
	return &RollingIndexApi_Expecter{mock: &_m.Mock}
}

// CreateRollingIndex provides a mock function with given fields: ctx, groupId, clusterName, databaseRollingIndexRequest
func (_m *RollingIndexApi) CreateRollingIndex(ctx context.Context, groupId string, clusterName string, databaseRollingIndexRequest *admin.DatabaseRollingIndexRequest) admin.CreateRollingIndexApiRequest {
	ret := _m.Called(ctx, groupId, clusterName, databaseRollingIndexRequest)

	if len(ret) == 0 {
		panic("no return value specified for CreateRollingIndex")
	}

	var r0 admin.CreateRollingIndexApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.DatabaseRollingIndexRequest) admin.CreateRollingIndexApiRequest); ok {
		r0 = rf(ctx, groupId, clusterName, databaseRollingIndexRequest)
	} else {
		r0 = ret.Get(0).(admin.CreateRollingIndexApiRequest)
	}

	return r0
}

// RollingIndexApi_CreateRollingIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRollingIndex'
type RollingIndexApi_CreateRollingIndex_Call struct {
	*mock.Call
}

// CreateRollingIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - clusterName string
//   - databaseRollingIndexRequest *admin.DatabaseRollingIndexRequest
func (_e *RollingIndexApi_Expecter) CreateRollingIndex(ctx interface{}, groupId interface{}, clusterName interface{}, databaseRollingIndexRequest interface{}) *RollingIndexApi_CreateRollingIndex_Call {
	return &RollingIndexApi_CreateRollingIndex_Call{Call: _e.mock.On("CreateRollingIndex", ctx, groupId, clusterName, databaseRollingIndexRequest)}
}

func (_c *RollingIndexApi_CreateRollingIndex_Call) Run(run func(ctx context.Context, groupId string, clusterName string, databaseRollingIndexRequest *admin.DatabaseRollingIndexRequest)) *RollingIndexApi_CreateRollingIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.DatabaseRollingIndexRequest))
	})
	return _c
}

func (_c *RollingIndexApi_CreateRollingIndex_Call) Return(_a0 admin.CreateRollingIndexApiRequest) *RollingIndexApi_CreateRollingIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RollingIndexApi_CreateRollingIndex_Call) RunAndReturn(run func(context.Context, string, string, *admin.DatabaseRollingIndexRequest) admin.CreateRollingIndexApiRequest) *RollingIndexApi_CreateRollingIndex_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRollingIndexExecute provides a mock function with given fields: r
func (_m *RollingIndexApi) CreateRollingIndexExecute(r admin.CreateRollingIndexApiRequest) (*http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for CreateRollingIndexExecute")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(admin.CreateRollingIndexApiRequest) (*http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.CreateRollingIndexApiRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.CreateRollingIndexApiRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RollingIndexApi_CreateRollingIndexExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRollingIndexExecute'
type RollingIndexApi_CreateRollingIndexExecute_Call struct {
	*mock.Call
}

// CreateRollingIndexExecute is a helper method to define mock.On call
//   - r admin.CreateRollingIndexApiRequest
func (_e *RollingIndexApi_Expecter) CreateRollingIndexExecute(r interface{}) *RollingIndexApi_CreateRollingIndexExecute_Call {
	return &RollingIndexApi_CreateRollingIndexExecute_Call{Call: _e.mock.On("CreateRollingIndexExecute", r)}
}

func (_c *RollingIndexApi_CreateRollingIndexExecute_Call) Run(run func(r admin.CreateRollingIndexApiRequest)) *RollingIndexApi_CreateRollingIndexExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.CreateRollingIndexApiRequest))
	})
	return _c
}

func (_c *RollingIndexApi_CreateRollingIndexExecute_Call) Return(_a0 *http.Response, _a1 error) *RollingIndexApi_CreateRollingIndexExecute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RollingIndexApi_CreateRollingIndexExecute_Call) RunAndReturn(run func(admin.CreateRollingIndexApiRequest) (*http.Response, error)) *RollingIndexApi_CreateRollingIndexExecute_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRollingIndexWithParams provides a mock function with given fields: ctx, args
func (_m *RollingIndexApi) CreateRollingIndexWithParams(ctx context.Context, args *admin.CreateRollingIndexApiParams) admin.CreateRollingIndexApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for CreateRollingIndexWithParams")
	}

	var r0 admin.CreateRollingIndexApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.CreateRollingIndexApiParams) admin.CreateRollingIndexApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.CreateRollingIndexApiRequest)
	}

	return r0
}

// RollingIndexApi_CreateRollingIndexWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRollingIndexWithParams'
type RollingIndexApi_CreateRollingIndexWithParams_Call struct {
	*mock.Call
}

// CreateRollingIndexWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.CreateRollingIndexApiParams
func (_e *RollingIndexApi_Expecter) CreateRollingIndexWithParams(ctx interface{}, args interface{}) *RollingIndexApi_CreateRollingIndexWithParams_Call {
	return &RollingIndexApi_CreateRollingIndexWithParams_Call{Call: _e.mock.On("CreateRollingIndexWithParams", ctx, args)}
}

func (_c *RollingIndexApi_CreateRollingIndexWithParams_Call) Run(run func(ctx context.Context, args *admin.CreateRollingIndexApiParams)) *RollingIndexApi_CreateRollingIndexWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.CreateRollingIndexApiParams))
	})
	return _c
}

func (_c *RollingIndexApi_CreateRollingIndexWithParams_Call) Return(_a0 admin.CreateRollingIndexApiRequest) *RollingIndexApi_CreateRollingIndexWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RollingIndexApi_CreateRollingIndexWithParams_Call) RunAndReturn(run func(context.Context, *admin.CreateRollingIndexApiParams) admin.CreateRollingIndexApiRequest) *RollingIndexApi_CreateRollingIndexWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewRollingIndexApi creates a new instance of RollingIndexApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRollingIndexApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *RollingIndexApi {
	mock := &RollingIndexApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
