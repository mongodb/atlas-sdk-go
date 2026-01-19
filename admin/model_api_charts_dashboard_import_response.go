// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiChartsDashboardImportResponse Imported dashboard and associated imported items.
type ApiChartsDashboardImportResponse struct {
	Dashboard *ImportedDashboard `json:"dashboard,omitempty"`
	// Imported items.
	Items *[]ImportedItem `json:"items,omitempty"`
}

// NewApiChartsDashboardImportResponse instantiates a new ApiChartsDashboardImportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiChartsDashboardImportResponse() *ApiChartsDashboardImportResponse {
	this := ApiChartsDashboardImportResponse{}
	return &this
}

// NewApiChartsDashboardImportResponseWithDefaults instantiates a new ApiChartsDashboardImportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiChartsDashboardImportResponseWithDefaults() *ApiChartsDashboardImportResponse {
	this := ApiChartsDashboardImportResponse{}
	return &this
}

// GetDashboard returns the Dashboard field value if set, zero value otherwise
func (o *ApiChartsDashboardImportResponse) GetDashboard() ImportedDashboard {
	if o == nil || IsNil(o.Dashboard) {
		var ret ImportedDashboard
		return ret
	}
	return *o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiChartsDashboardImportResponse) GetDashboardOk() (*ImportedDashboard, bool) {
	if o == nil || IsNil(o.Dashboard) {
		return nil, false
	}

	return o.Dashboard, true
}

// HasDashboard returns a boolean if a field has been set.
func (o *ApiChartsDashboardImportResponse) HasDashboard() bool {
	if o != nil && !IsNil(o.Dashboard) {
		return true
	}

	return false
}

// SetDashboard gets a reference to the given ImportedDashboard and assigns it to the Dashboard field.
func (o *ApiChartsDashboardImportResponse) SetDashboard(v ImportedDashboard) {
	o.Dashboard = &v
}

// GetItems returns the Items field value if set, zero value otherwise
func (o *ApiChartsDashboardImportResponse) GetItems() []ImportedItem {
	if o == nil || IsNil(o.Items) {
		var ret []ImportedItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiChartsDashboardImportResponse) GetItemsOk() (*[]ImportedItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}

	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ApiChartsDashboardImportResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ImportedItem and assigns it to the Items field.
func (o *ApiChartsDashboardImportResponse) SetItems(v []ImportedItem) {
	o.Items = &v
}
