// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterCostEstimate Estimated compute cost of running a cluster.
type ClusterCostEstimate struct {
	// Human-readable notes about what the estimate does and does not cover. Suitable for rendering line-by-line to the user.
	// Read only field.
	Disclaimer []string `json:"disclaimer"`
	// Estimated compute cost per hour, in hundredths of a credit. One Atlas credit converts 1:1 to one US dollar.
	// Read only field.
	HourlyEstimateCents int64 `json:"hourlyEstimateCents"`
	// Estimated compute cost per month, in hundredths of a credit. One Atlas credit converts 1:1 to one US dollar.
	// Read only field.
	MonthlyEstimateCents int64 `json:"monthlyEstimateCents"`
}

// NewClusterCostEstimate instantiates a new ClusterCostEstimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCostEstimate(disclaimer []string, hourlyEstimateCents int64, monthlyEstimateCents int64) *ClusterCostEstimate {
	this := ClusterCostEstimate{}
	this.Disclaimer = disclaimer
	this.HourlyEstimateCents = hourlyEstimateCents
	this.MonthlyEstimateCents = monthlyEstimateCents
	return &this
}

// NewClusterCostEstimateWithDefaults instantiates a new ClusterCostEstimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCostEstimateWithDefaults() *ClusterCostEstimate {
	this := ClusterCostEstimate{}
	return &this
}

// GetDisclaimer returns the Disclaimer field value
func (o *ClusterCostEstimate) GetDisclaimer() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Disclaimer
}

// GetDisclaimerOk returns a tuple with the Disclaimer field value
// and a boolean to check if the value has been set.
func (o *ClusterCostEstimate) GetDisclaimerOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disclaimer, true
}

// SetDisclaimer sets field value
func (o *ClusterCostEstimate) SetDisclaimer(v []string) {
	o.Disclaimer = v
}

// GetHourlyEstimateCents returns the HourlyEstimateCents field value
func (o *ClusterCostEstimate) GetHourlyEstimateCents() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HourlyEstimateCents
}

// GetHourlyEstimateCentsOk returns a tuple with the HourlyEstimateCents field value
// and a boolean to check if the value has been set.
func (o *ClusterCostEstimate) GetHourlyEstimateCentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourlyEstimateCents, true
}

// SetHourlyEstimateCents sets field value
func (o *ClusterCostEstimate) SetHourlyEstimateCents(v int64) {
	o.HourlyEstimateCents = v
}

// GetMonthlyEstimateCents returns the MonthlyEstimateCents field value
func (o *ClusterCostEstimate) GetMonthlyEstimateCents() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MonthlyEstimateCents
}

// GetMonthlyEstimateCentsOk returns a tuple with the MonthlyEstimateCents field value
// and a boolean to check if the value has been set.
func (o *ClusterCostEstimate) GetMonthlyEstimateCentsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyEstimateCents, true
}

// SetMonthlyEstimateCents sets field value
func (o *ClusterCostEstimate) SetMonthlyEstimateCents(v int64) {
	o.MonthlyEstimateCents = v
}
