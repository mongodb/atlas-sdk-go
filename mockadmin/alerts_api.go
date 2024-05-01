// Code generated by mockery. DO NOT EDIT.

package mockadmin

import (
	context "context"

	admin "go.mongodb.org/atlas-sdk/v20231115013/admin"

	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// AlertsApi is an autogenerated mock type for the AlertsApi type
type AlertsApi struct {
	mock.Mock
}

type AlertsApi_Expecter struct {
	mock *mock.Mock
}

func (_m *AlertsApi) EXPECT() *AlertsApi_Expecter {
	return &AlertsApi_Expecter{mock: &_m.Mock}
}

// AcknowledgeAlert provides a mock function with given fields: ctx, groupId, alertId, alertViewForNdsGroup
func (_m *AlertsApi) AcknowledgeAlert(ctx context.Context, groupId string, alertId string, alertViewForNdsGroup *admin.AlertViewForNdsGroup) admin.AcknowledgeAlertApiRequest {
	ret := _m.Called(ctx, groupId, alertId, alertViewForNdsGroup)

	if len(ret) == 0 {
		panic("no return value specified for AcknowledgeAlert")
	}

	var r0 admin.AcknowledgeAlertApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *admin.AlertViewForNdsGroup) admin.AcknowledgeAlertApiRequest); ok {
		r0 = rf(ctx, groupId, alertId, alertViewForNdsGroup)
	} else {
		r0 = ret.Get(0).(admin.AcknowledgeAlertApiRequest)
	}

	return r0
}

// AlertsApi_AcknowledgeAlert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AcknowledgeAlert'
type AlertsApi_AcknowledgeAlert_Call struct {
	*mock.Call
}

// AcknowledgeAlert is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - alertId string
//   - alertViewForNdsGroup *admin.AlertViewForNdsGroup
func (_e *AlertsApi_Expecter) AcknowledgeAlert(ctx interface{}, groupId interface{}, alertId interface{}, alertViewForNdsGroup interface{}) *AlertsApi_AcknowledgeAlert_Call {
	return &AlertsApi_AcknowledgeAlert_Call{Call: _e.mock.On("AcknowledgeAlert", ctx, groupId, alertId, alertViewForNdsGroup)}
}

func (_c *AlertsApi_AcknowledgeAlert_Call) Run(run func(ctx context.Context, groupId string, alertId string, alertViewForNdsGroup *admin.AlertViewForNdsGroup)) *AlertsApi_AcknowledgeAlert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*admin.AlertViewForNdsGroup))
	})
	return _c
}

func (_c *AlertsApi_AcknowledgeAlert_Call) Return(_a0 admin.AcknowledgeAlertApiRequest) *AlertsApi_AcknowledgeAlert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_AcknowledgeAlert_Call) RunAndReturn(run func(context.Context, string, string, *admin.AlertViewForNdsGroup) admin.AcknowledgeAlertApiRequest) *AlertsApi_AcknowledgeAlert_Call {
	_c.Call.Return(run)
	return _c
}

// AcknowledgeAlertExecute provides a mock function with given fields: r
func (_m *AlertsApi) AcknowledgeAlertExecute(r admin.AcknowledgeAlertApiRequest) (*admin.AlertViewForNdsGroup, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for AcknowledgeAlertExecute")
	}

	var r0 *admin.AlertViewForNdsGroup
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.AcknowledgeAlertApiRequest) (*admin.AlertViewForNdsGroup, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.AcknowledgeAlertApiRequest) *admin.AlertViewForNdsGroup); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AlertViewForNdsGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.AcknowledgeAlertApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.AcknowledgeAlertApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AlertsApi_AcknowledgeAlertExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AcknowledgeAlertExecute'
type AlertsApi_AcknowledgeAlertExecute_Call struct {
	*mock.Call
}

// AcknowledgeAlertExecute is a helper method to define mock.On call
//   - r admin.AcknowledgeAlertApiRequest
func (_e *AlertsApi_Expecter) AcknowledgeAlertExecute(r interface{}) *AlertsApi_AcknowledgeAlertExecute_Call {
	return &AlertsApi_AcknowledgeAlertExecute_Call{Call: _e.mock.On("AcknowledgeAlertExecute", r)}
}

func (_c *AlertsApi_AcknowledgeAlertExecute_Call) Run(run func(r admin.AcknowledgeAlertApiRequest)) *AlertsApi_AcknowledgeAlertExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.AcknowledgeAlertApiRequest))
	})
	return _c
}

func (_c *AlertsApi_AcknowledgeAlertExecute_Call) Return(_a0 *admin.AlertViewForNdsGroup, _a1 *http.Response, _a2 error) *AlertsApi_AcknowledgeAlertExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AlertsApi_AcknowledgeAlertExecute_Call) RunAndReturn(run func(admin.AcknowledgeAlertApiRequest) (*admin.AlertViewForNdsGroup, *http.Response, error)) *AlertsApi_AcknowledgeAlertExecute_Call {
	_c.Call.Return(run)
	return _c
}

// AcknowledgeAlertWithParams provides a mock function with given fields: ctx, args
func (_m *AlertsApi) AcknowledgeAlertWithParams(ctx context.Context, args *admin.AcknowledgeAlertApiParams) admin.AcknowledgeAlertApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for AcknowledgeAlertWithParams")
	}

	var r0 admin.AcknowledgeAlertApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.AcknowledgeAlertApiParams) admin.AcknowledgeAlertApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.AcknowledgeAlertApiRequest)
	}

	return r0
}

// AlertsApi_AcknowledgeAlertWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AcknowledgeAlertWithParams'
type AlertsApi_AcknowledgeAlertWithParams_Call struct {
	*mock.Call
}

// AcknowledgeAlertWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.AcknowledgeAlertApiParams
func (_e *AlertsApi_Expecter) AcknowledgeAlertWithParams(ctx interface{}, args interface{}) *AlertsApi_AcknowledgeAlertWithParams_Call {
	return &AlertsApi_AcknowledgeAlertWithParams_Call{Call: _e.mock.On("AcknowledgeAlertWithParams", ctx, args)}
}

func (_c *AlertsApi_AcknowledgeAlertWithParams_Call) Run(run func(ctx context.Context, args *admin.AcknowledgeAlertApiParams)) *AlertsApi_AcknowledgeAlertWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.AcknowledgeAlertApiParams))
	})
	return _c
}

func (_c *AlertsApi_AcknowledgeAlertWithParams_Call) Return(_a0 admin.AcknowledgeAlertApiRequest) *AlertsApi_AcknowledgeAlertWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_AcknowledgeAlertWithParams_Call) RunAndReturn(run func(context.Context, *admin.AcknowledgeAlertApiParams) admin.AcknowledgeAlertApiRequest) *AlertsApi_AcknowledgeAlertWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlert provides a mock function with given fields: ctx, groupId, alertId
func (_m *AlertsApi) GetAlert(ctx context.Context, groupId string, alertId string) admin.GetAlertApiRequest {
	ret := _m.Called(ctx, groupId, alertId)

	if len(ret) == 0 {
		panic("no return value specified for GetAlert")
	}

	var r0 admin.GetAlertApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.GetAlertApiRequest); ok {
		r0 = rf(ctx, groupId, alertId)
	} else {
		r0 = ret.Get(0).(admin.GetAlertApiRequest)
	}

	return r0
}

// AlertsApi_GetAlert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlert'
type AlertsApi_GetAlert_Call struct {
	*mock.Call
}

// GetAlert is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - alertId string
func (_e *AlertsApi_Expecter) GetAlert(ctx interface{}, groupId interface{}, alertId interface{}) *AlertsApi_GetAlert_Call {
	return &AlertsApi_GetAlert_Call{Call: _e.mock.On("GetAlert", ctx, groupId, alertId)}
}

func (_c *AlertsApi_GetAlert_Call) Run(run func(ctx context.Context, groupId string, alertId string)) *AlertsApi_GetAlert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AlertsApi_GetAlert_Call) Return(_a0 admin.GetAlertApiRequest) *AlertsApi_GetAlert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_GetAlert_Call) RunAndReturn(run func(context.Context, string, string) admin.GetAlertApiRequest) *AlertsApi_GetAlert_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlertExecute provides a mock function with given fields: r
func (_m *AlertsApi) GetAlertExecute(r admin.GetAlertApiRequest) (*admin.AlertViewForNdsGroup, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for GetAlertExecute")
	}

	var r0 *admin.AlertViewForNdsGroup
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.GetAlertApiRequest) (*admin.AlertViewForNdsGroup, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.GetAlertApiRequest) *admin.AlertViewForNdsGroup); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.AlertViewForNdsGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.GetAlertApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.GetAlertApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AlertsApi_GetAlertExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlertExecute'
type AlertsApi_GetAlertExecute_Call struct {
	*mock.Call
}

// GetAlertExecute is a helper method to define mock.On call
//   - r admin.GetAlertApiRequest
func (_e *AlertsApi_Expecter) GetAlertExecute(r interface{}) *AlertsApi_GetAlertExecute_Call {
	return &AlertsApi_GetAlertExecute_Call{Call: _e.mock.On("GetAlertExecute", r)}
}

func (_c *AlertsApi_GetAlertExecute_Call) Run(run func(r admin.GetAlertApiRequest)) *AlertsApi_GetAlertExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.GetAlertApiRequest))
	})
	return _c
}

func (_c *AlertsApi_GetAlertExecute_Call) Return(_a0 *admin.AlertViewForNdsGroup, _a1 *http.Response, _a2 error) *AlertsApi_GetAlertExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AlertsApi_GetAlertExecute_Call) RunAndReturn(run func(admin.GetAlertApiRequest) (*admin.AlertViewForNdsGroup, *http.Response, error)) *AlertsApi_GetAlertExecute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlertWithParams provides a mock function with given fields: ctx, args
func (_m *AlertsApi) GetAlertWithParams(ctx context.Context, args *admin.GetAlertApiParams) admin.GetAlertApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for GetAlertWithParams")
	}

	var r0 admin.GetAlertApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetAlertApiParams) admin.GetAlertApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.GetAlertApiRequest)
	}

	return r0
}

// AlertsApi_GetAlertWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlertWithParams'
type AlertsApi_GetAlertWithParams_Call struct {
	*mock.Call
}

// GetAlertWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.GetAlertApiParams
func (_e *AlertsApi_Expecter) GetAlertWithParams(ctx interface{}, args interface{}) *AlertsApi_GetAlertWithParams_Call {
	return &AlertsApi_GetAlertWithParams_Call{Call: _e.mock.On("GetAlertWithParams", ctx, args)}
}

func (_c *AlertsApi_GetAlertWithParams_Call) Run(run func(ctx context.Context, args *admin.GetAlertApiParams)) *AlertsApi_GetAlertWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.GetAlertApiParams))
	})
	return _c
}

func (_c *AlertsApi_GetAlertWithParams_Call) Return(_a0 admin.GetAlertApiRequest) *AlertsApi_GetAlertWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_GetAlertWithParams_Call) RunAndReturn(run func(context.Context, *admin.GetAlertApiParams) admin.GetAlertApiRequest) *AlertsApi_GetAlertWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlerts provides a mock function with given fields: ctx, groupId
func (_m *AlertsApi) ListAlerts(ctx context.Context, groupId string) admin.ListAlertsApiRequest {
	ret := _m.Called(ctx, groupId)

	if len(ret) == 0 {
		panic("no return value specified for ListAlerts")
	}

	var r0 admin.ListAlertsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) admin.ListAlertsApiRequest); ok {
		r0 = rf(ctx, groupId)
	} else {
		r0 = ret.Get(0).(admin.ListAlertsApiRequest)
	}

	return r0
}

// AlertsApi_ListAlerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlerts'
type AlertsApi_ListAlerts_Call struct {
	*mock.Call
}

// ListAlerts is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
func (_e *AlertsApi_Expecter) ListAlerts(ctx interface{}, groupId interface{}) *AlertsApi_ListAlerts_Call {
	return &AlertsApi_ListAlerts_Call{Call: _e.mock.On("ListAlerts", ctx, groupId)}
}

func (_c *AlertsApi_ListAlerts_Call) Run(run func(ctx context.Context, groupId string)) *AlertsApi_ListAlerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *AlertsApi_ListAlerts_Call) Return(_a0 admin.ListAlertsApiRequest) *AlertsApi_ListAlerts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_ListAlerts_Call) RunAndReturn(run func(context.Context, string) admin.ListAlertsApiRequest) *AlertsApi_ListAlerts_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlertsByAlertConfigurationId provides a mock function with given fields: ctx, groupId, alertConfigId
func (_m *AlertsApi) ListAlertsByAlertConfigurationId(ctx context.Context, groupId string, alertConfigId string) admin.ListAlertsByAlertConfigurationIdApiRequest {
	ret := _m.Called(ctx, groupId, alertConfigId)

	if len(ret) == 0 {
		panic("no return value specified for ListAlertsByAlertConfigurationId")
	}

	var r0 admin.ListAlertsByAlertConfigurationIdApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, string, string) admin.ListAlertsByAlertConfigurationIdApiRequest); ok {
		r0 = rf(ctx, groupId, alertConfigId)
	} else {
		r0 = ret.Get(0).(admin.ListAlertsByAlertConfigurationIdApiRequest)
	}

	return r0
}

// AlertsApi_ListAlertsByAlertConfigurationId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlertsByAlertConfigurationId'
type AlertsApi_ListAlertsByAlertConfigurationId_Call struct {
	*mock.Call
}

// ListAlertsByAlertConfigurationId is a helper method to define mock.On call
//   - ctx context.Context
//   - groupId string
//   - alertConfigId string
func (_e *AlertsApi_Expecter) ListAlertsByAlertConfigurationId(ctx interface{}, groupId interface{}, alertConfigId interface{}) *AlertsApi_ListAlertsByAlertConfigurationId_Call {
	return &AlertsApi_ListAlertsByAlertConfigurationId_Call{Call: _e.mock.On("ListAlertsByAlertConfigurationId", ctx, groupId, alertConfigId)}
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationId_Call) Run(run func(ctx context.Context, groupId string, alertConfigId string)) *AlertsApi_ListAlertsByAlertConfigurationId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationId_Call) Return(_a0 admin.ListAlertsByAlertConfigurationIdApiRequest) *AlertsApi_ListAlertsByAlertConfigurationId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationId_Call) RunAndReturn(run func(context.Context, string, string) admin.ListAlertsByAlertConfigurationIdApiRequest) *AlertsApi_ListAlertsByAlertConfigurationId_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlertsByAlertConfigurationIdExecute provides a mock function with given fields: r
func (_m *AlertsApi) ListAlertsByAlertConfigurationIdExecute(r admin.ListAlertsByAlertConfigurationIdApiRequest) (*admin.PaginatedAlert, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListAlertsByAlertConfigurationIdExecute")
	}

	var r0 *admin.PaginatedAlert
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListAlertsByAlertConfigurationIdApiRequest) (*admin.PaginatedAlert, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListAlertsByAlertConfigurationIdApiRequest) *admin.PaginatedAlert); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedAlert)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListAlertsByAlertConfigurationIdApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListAlertsByAlertConfigurationIdApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlertsByAlertConfigurationIdExecute'
type AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call struct {
	*mock.Call
}

// ListAlertsByAlertConfigurationIdExecute is a helper method to define mock.On call
//   - r admin.ListAlertsByAlertConfigurationIdApiRequest
func (_e *AlertsApi_Expecter) ListAlertsByAlertConfigurationIdExecute(r interface{}) *AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call {
	return &AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call{Call: _e.mock.On("ListAlertsByAlertConfigurationIdExecute", r)}
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call) Run(run func(r admin.ListAlertsByAlertConfigurationIdApiRequest)) *AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListAlertsByAlertConfigurationIdApiRequest))
	})
	return _c
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call) Return(_a0 *admin.PaginatedAlert, _a1 *http.Response, _a2 error) *AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call) RunAndReturn(run func(admin.ListAlertsByAlertConfigurationIdApiRequest) (*admin.PaginatedAlert, *http.Response, error)) *AlertsApi_ListAlertsByAlertConfigurationIdExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlertsByAlertConfigurationIdWithParams provides a mock function with given fields: ctx, args
func (_m *AlertsApi) ListAlertsByAlertConfigurationIdWithParams(ctx context.Context, args *admin.ListAlertsByAlertConfigurationIdApiParams) admin.ListAlertsByAlertConfigurationIdApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListAlertsByAlertConfigurationIdWithParams")
	}

	var r0 admin.ListAlertsByAlertConfigurationIdApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListAlertsByAlertConfigurationIdApiParams) admin.ListAlertsByAlertConfigurationIdApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListAlertsByAlertConfigurationIdApiRequest)
	}

	return r0
}

// AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlertsByAlertConfigurationIdWithParams'
type AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call struct {
	*mock.Call
}

// ListAlertsByAlertConfigurationIdWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListAlertsByAlertConfigurationIdApiParams
func (_e *AlertsApi_Expecter) ListAlertsByAlertConfigurationIdWithParams(ctx interface{}, args interface{}) *AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call {
	return &AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call{Call: _e.mock.On("ListAlertsByAlertConfigurationIdWithParams", ctx, args)}
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call) Run(run func(ctx context.Context, args *admin.ListAlertsByAlertConfigurationIdApiParams)) *AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListAlertsByAlertConfigurationIdApiParams))
	})
	return _c
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call) Return(_a0 admin.ListAlertsByAlertConfigurationIdApiRequest) *AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListAlertsByAlertConfigurationIdApiParams) admin.ListAlertsByAlertConfigurationIdApiRequest) *AlertsApi_ListAlertsByAlertConfigurationIdWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlertsExecute provides a mock function with given fields: r
func (_m *AlertsApi) ListAlertsExecute(r admin.ListAlertsApiRequest) (*admin.PaginatedAlert, *http.Response, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for ListAlertsExecute")
	}

	var r0 *admin.PaginatedAlert
	var r1 *http.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(admin.ListAlertsApiRequest) (*admin.PaginatedAlert, *http.Response, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(admin.ListAlertsApiRequest) *admin.PaginatedAlert); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.PaginatedAlert)
		}
	}

	if rf, ok := ret.Get(1).(func(admin.ListAlertsApiRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(admin.ListAlertsApiRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// AlertsApi_ListAlertsExecute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlertsExecute'
type AlertsApi_ListAlertsExecute_Call struct {
	*mock.Call
}

// ListAlertsExecute is a helper method to define mock.On call
//   - r admin.ListAlertsApiRequest
func (_e *AlertsApi_Expecter) ListAlertsExecute(r interface{}) *AlertsApi_ListAlertsExecute_Call {
	return &AlertsApi_ListAlertsExecute_Call{Call: _e.mock.On("ListAlertsExecute", r)}
}

func (_c *AlertsApi_ListAlertsExecute_Call) Run(run func(r admin.ListAlertsApiRequest)) *AlertsApi_ListAlertsExecute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(admin.ListAlertsApiRequest))
	})
	return _c
}

func (_c *AlertsApi_ListAlertsExecute_Call) Return(_a0 *admin.PaginatedAlert, _a1 *http.Response, _a2 error) *AlertsApi_ListAlertsExecute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *AlertsApi_ListAlertsExecute_Call) RunAndReturn(run func(admin.ListAlertsApiRequest) (*admin.PaginatedAlert, *http.Response, error)) *AlertsApi_ListAlertsExecute_Call {
	_c.Call.Return(run)
	return _c
}

// ListAlertsWithParams provides a mock function with given fields: ctx, args
func (_m *AlertsApi) ListAlertsWithParams(ctx context.Context, args *admin.ListAlertsApiParams) admin.ListAlertsApiRequest {
	ret := _m.Called(ctx, args)

	if len(ret) == 0 {
		panic("no return value specified for ListAlertsWithParams")
	}

	var r0 admin.ListAlertsApiRequest
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListAlertsApiParams) admin.ListAlertsApiRequest); ok {
		r0 = rf(ctx, args)
	} else {
		r0 = ret.Get(0).(admin.ListAlertsApiRequest)
	}

	return r0
}

// AlertsApi_ListAlertsWithParams_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAlertsWithParams'
type AlertsApi_ListAlertsWithParams_Call struct {
	*mock.Call
}

// ListAlertsWithParams is a helper method to define mock.On call
//   - ctx context.Context
//   - args *admin.ListAlertsApiParams
func (_e *AlertsApi_Expecter) ListAlertsWithParams(ctx interface{}, args interface{}) *AlertsApi_ListAlertsWithParams_Call {
	return &AlertsApi_ListAlertsWithParams_Call{Call: _e.mock.On("ListAlertsWithParams", ctx, args)}
}

func (_c *AlertsApi_ListAlertsWithParams_Call) Run(run func(ctx context.Context, args *admin.ListAlertsApiParams)) *AlertsApi_ListAlertsWithParams_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.ListAlertsApiParams))
	})
	return _c
}

func (_c *AlertsApi_ListAlertsWithParams_Call) Return(_a0 admin.ListAlertsApiRequest) *AlertsApi_ListAlertsWithParams_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AlertsApi_ListAlertsWithParams_Call) RunAndReturn(run func(context.Context, *admin.ListAlertsApiParams) admin.ListAlertsApiRequest) *AlertsApi_ListAlertsWithParams_Call {
	_c.Call.Return(run)
	return _c
}

// NewAlertsApi creates a new instance of AlertsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAlertsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *AlertsApi {
	mock := &AlertsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
