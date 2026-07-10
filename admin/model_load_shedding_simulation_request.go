// Code based on the AtlasAPI V2 OpenAPI file

package admin

// LoadSheddingSimulationRequest Request to start a load shedding simulation for a cluster.
type LoadSheddingSimulationRequest struct {
	// Duration of the load shedding simulation in seconds.
	DurationSeconds *int `json:"durationSeconds,omitempty"`
}

// NewLoadSheddingSimulationRequest instantiates a new LoadSheddingSimulationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadSheddingSimulationRequest() *LoadSheddingSimulationRequest {
	this := LoadSheddingSimulationRequest{}
	return &this
}

// NewLoadSheddingSimulationRequestWithDefaults instantiates a new LoadSheddingSimulationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadSheddingSimulationRequestWithDefaults() *LoadSheddingSimulationRequest {
	this := LoadSheddingSimulationRequest{}
	return &this
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise
func (o *LoadSheddingSimulationRequest) GetDurationSeconds() int {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret int
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadSheddingSimulationRequest) GetDurationSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}

	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *LoadSheddingSimulationRequest) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given int and assigns it to the DurationSeconds field.
func (o *LoadSheddingSimulationRequest) SetDurationSeconds(v int) {
	o.DurationSeconds = &v
}
