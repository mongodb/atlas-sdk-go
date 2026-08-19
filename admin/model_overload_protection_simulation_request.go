// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OverloadProtectionSimulationRequest Request to start an overload protection simulation for a cluster.
type OverloadProtectionSimulationRequest struct {
	// Duration of the overload protection simulation in seconds.
	DurationSeconds int `json:"durationSeconds"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OverloadProtectionSimulationRequest) MarshalJSON() ([]byte, error) {
	type noMethod OverloadProtectionSimulationRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOverloadProtectionSimulationRequest instantiates a new OverloadProtectionSimulationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverloadProtectionSimulationRequest(durationSeconds int) *OverloadProtectionSimulationRequest {
	this := OverloadProtectionSimulationRequest{}
	this.DurationSeconds = durationSeconds
	return &this
}

// NewOverloadProtectionSimulationRequestWithDefaults instantiates a new OverloadProtectionSimulationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverloadProtectionSimulationRequestWithDefaults() *OverloadProtectionSimulationRequest {
	this := OverloadProtectionSimulationRequest{}
	return &this
}

// GetDurationSeconds returns the DurationSeconds field value
func (o *OverloadProtectionSimulationRequest) GetDurationSeconds() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *OverloadProtectionSimulationRequest) GetDurationSecondsOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value
func (o *OverloadProtectionSimulationRequest) SetDurationSeconds(v int) {
	o.DurationSeconds = v
}
