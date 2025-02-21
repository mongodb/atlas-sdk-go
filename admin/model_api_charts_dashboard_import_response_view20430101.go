// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiChartsDashboardImportResponseView20430101 Imported dashboard and associated imported items.
type ApiChartsDashboardImportResponseView20430101 struct {
	Dashboard *ImportedDashboard `json:"dashboard,omitempty"`
	// Imported items.
	Items *[]ImportedItem `json:"items,omitempty"`
}

// NewApiChartsDashboardImportResponseView20430101 instantiates a new ApiChartsDashboardImportResponseView20430101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiChartsDashboardImportResponseView20430101() *ApiChartsDashboardImportResponseView20430101 {
	this := ApiChartsDashboardImportResponseView20430101{}
	return &this
}

// NewApiChartsDashboardImportResponseView20430101WithDefaults instantiates a new ApiChartsDashboardImportResponseView20430101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiChartsDashboardImportResponseView20430101WithDefaults() *ApiChartsDashboardImportResponseView20430101 {
	this := ApiChartsDashboardImportResponseView20430101{}
	return &this
}

// GetDashboard returns the Dashboard field value if set, zero value otherwise
func (o *ApiChartsDashboardImportResponseView20430101) GetDashboard() ImportedDashboard {
	if o == nil || IsNil(o.Dashboard) {
		var ret ImportedDashboard
		return ret
	}
	return *o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiChartsDashboardImportResponseView20430101) GetDashboardOk() (*ImportedDashboard, bool) {
	if o == nil || IsNil(o.Dashboard) {
		return nil, false
	}

	return o.Dashboard, true
}

// HasDashboard returns a boolean if a field has been set.
func (o *ApiChartsDashboardImportResponseView20430101) HasDashboard() bool {
	if o != nil && !IsNil(o.Dashboard) {
		return true
	}

	return false
}

// SetDashboard gets a reference to the given ImportedDashboard and assigns it to the Dashboard field.
func (o *ApiChartsDashboardImportResponseView20430101) SetDashboard(v ImportedDashboard) {
	o.Dashboard = &v
}

// GetItems returns the Items field value if set, zero value otherwise
func (o *ApiChartsDashboardImportResponseView20430101) GetItems() []ImportedItem {
	if o == nil || IsNil(o.Items) {
		var ret []ImportedItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiChartsDashboardImportResponseView20430101) GetItemsOk() (*[]ImportedItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}

	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ApiChartsDashboardImportResponseView20430101) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ImportedItem and assigns it to the Items field.
func (o *ApiChartsDashboardImportResponseView20430101) SetItems(v []ImportedItem) {
	o.Items = &v
}
