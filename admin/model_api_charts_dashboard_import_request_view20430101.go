// Code based on the AtlasAPI V2 OpenAPI file

package admin

import "io"

// ApiChartsDashboardImportRequestView20430101 Details to import or override a dashboard.
type ApiChartsDashboardImportRequestView20430101 struct {
	// Optional field that contains the ID of the dashboard to override. If `override=true`, you must provide this field.
	DashboardId *string `json:"dashboardId,omitempty"`
	// Flag that indicates whether to override a dashboard.
	Override *bool `json:"override,omitempty"`
	// Template of the dashboard to import, which corresponds to the returned response schema from calling the Dashboard Export API.
	Template *io.ReadCloser `json:"template,omitempty"`
}

// NewApiChartsDashboardImportRequestView20430101 instantiates a new ApiChartsDashboardImportRequestView20430101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiChartsDashboardImportRequestView20430101() *ApiChartsDashboardImportRequestView20430101 {
	this := ApiChartsDashboardImportRequestView20430101{}
	return &this
}

// NewApiChartsDashboardImportRequestView20430101WithDefaults instantiates a new ApiChartsDashboardImportRequestView20430101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiChartsDashboardImportRequestView20430101WithDefaults() *ApiChartsDashboardImportRequestView20430101 {
	this := ApiChartsDashboardImportRequestView20430101{}
	return &this
}

// GetDashboardId returns the DashboardId field value if set, zero value otherwise
func (o *ApiChartsDashboardImportRequestView20430101) GetDashboardId() string {
	if o == nil || IsNil(o.DashboardId) {
		var ret string
		return ret
	}
	return *o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiChartsDashboardImportRequestView20430101) GetDashboardIdOk() (*string, bool) {
	if o == nil || IsNil(o.DashboardId) {
		return nil, false
	}

	return o.DashboardId, true
}

// HasDashboardId returns a boolean if a field has been set.
func (o *ApiChartsDashboardImportRequestView20430101) HasDashboardId() bool {
	if o != nil && !IsNil(o.DashboardId) {
		return true
	}

	return false
}

// SetDashboardId gets a reference to the given string and assigns it to the DashboardId field.
func (o *ApiChartsDashboardImportRequestView20430101) SetDashboardId(v string) {
	o.DashboardId = &v
}

// GetOverride returns the Override field value if set, zero value otherwise
func (o *ApiChartsDashboardImportRequestView20430101) GetOverride() bool {
	if o == nil || IsNil(o.Override) {
		var ret bool
		return ret
	}
	return *o.Override
}

// GetOverrideOk returns a tuple with the Override field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiChartsDashboardImportRequestView20430101) GetOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.Override) {
		return nil, false
	}

	return o.Override, true
}

// HasOverride returns a boolean if a field has been set.
func (o *ApiChartsDashboardImportRequestView20430101) HasOverride() bool {
	if o != nil && !IsNil(o.Override) {
		return true
	}

	return false
}

// SetOverride gets a reference to the given bool and assigns it to the Override field.
func (o *ApiChartsDashboardImportRequestView20430101) SetOverride(v bool) {
	o.Override = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise
func (o *ApiChartsDashboardImportRequestView20430101) GetTemplate() io.ReadCloser {
	if o == nil || IsNil(o.Template) {
		var ret io.ReadCloser
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiChartsDashboardImportRequestView20430101) GetTemplateOk() (*io.ReadCloser, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}

	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ApiChartsDashboardImportRequestView20430101) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given io.ReadCloser and assigns it to the Template field.
func (o *ApiChartsDashboardImportRequestView20430101) SetTemplate(v io.ReadCloser) {
	o.Template = &v
}
