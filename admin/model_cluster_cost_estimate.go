// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterCostEstimate Estimated compute cost of running a cluster.
type ClusterCostEstimate struct {
	// Human-readable notes about what the estimate does and does not cover. Suitable for rendering to the user. Omitted for free-tier (M0) estimates.
	// Read only field.
	Disclaimer *string `json:"disclaimer,omitempty"`
	// Estimated compute cost per hour, in Atlas credits. One Atlas credit converts 1:1 to one US dollar.
	// Read only field.
	HourlyEstimateCredits float32 `json:"hourlyEstimateCredits"`
	// Estimated compute cost per month, in Atlas credits. One Atlas credit converts 1:1 to one US dollar.
	// Read only field.
	MonthlyEstimateCredits float32 `json:"monthlyEstimateCredits"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ClusterCostEstimate) MarshalJSON() ([]byte, error) {
	type noMethod ClusterCostEstimate
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewClusterCostEstimate instantiates a new ClusterCostEstimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCostEstimate(hourlyEstimateCredits float32, monthlyEstimateCredits float32) *ClusterCostEstimate {
	this := ClusterCostEstimate{}
	this.HourlyEstimateCredits = hourlyEstimateCredits
	this.MonthlyEstimateCredits = monthlyEstimateCredits
	return &this
}

// NewClusterCostEstimateWithDefaults instantiates a new ClusterCostEstimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCostEstimateWithDefaults() *ClusterCostEstimate {
	this := ClusterCostEstimate{}
	return &this
}

// GetDisclaimer returns the Disclaimer field value if set, zero value otherwise
func (o *ClusterCostEstimate) GetDisclaimer() string {
	if o == nil || IsNil(o.Disclaimer) {
		var ret string
		return ret
	}
	return *o.Disclaimer
}

// GetDisclaimerOk returns a tuple with the Disclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCostEstimate) GetDisclaimerOk() (*string, bool) {
	if o == nil || IsNil(o.Disclaimer) {
		return nil, false
	}

	return o.Disclaimer, true
}

// HasDisclaimer returns a boolean if a field has been set.
func (o *ClusterCostEstimate) HasDisclaimer() bool {
	if o != nil && !IsNil(o.Disclaimer) {
		return true
	}

	return false
}

// SetDisclaimer gets a reference to the given string and assigns it to the Disclaimer field.
func (o *ClusterCostEstimate) SetDisclaimer(v string) {
	o.Disclaimer = &v
	o.NullFields = removeNullField(o.NullFields, "Disclaimer")
}

// SetDisclaimerNil sets Disclaimer to an explicit JSON null when marshaled.
func (o *ClusterCostEstimate) SetDisclaimerNil() {
	o.Disclaimer = nil
	o.NullFields = addNullField(o.NullFields, "Disclaimer")
}

// GetHourlyEstimateCredits returns the HourlyEstimateCredits field value
func (o *ClusterCostEstimate) GetHourlyEstimateCredits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HourlyEstimateCredits
}

// GetHourlyEstimateCreditsOk returns a tuple with the HourlyEstimateCredits field value
// and a boolean to check if the value has been set.
func (o *ClusterCostEstimate) GetHourlyEstimateCreditsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourlyEstimateCredits, true
}

// SetHourlyEstimateCredits sets field value
func (o *ClusterCostEstimate) SetHourlyEstimateCredits(v float32) {
	o.HourlyEstimateCredits = v
}

// GetMonthlyEstimateCredits returns the MonthlyEstimateCredits field value
func (o *ClusterCostEstimate) GetMonthlyEstimateCredits() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MonthlyEstimateCredits
}

// GetMonthlyEstimateCreditsOk returns a tuple with the MonthlyEstimateCredits field value
// and a boolean to check if the value has been set.
func (o *ClusterCostEstimate) GetMonthlyEstimateCreditsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyEstimateCredits, true
}

// SetMonthlyEstimateCredits sets field value
func (o *ClusterCostEstimate) SetMonthlyEstimateCredits(v float32) {
	o.MonthlyEstimateCredits = v
}
