// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiSearchAutoScaling Settings that let Atlas change the Search Node tier on its own. Bounds apply to the whole deployment, and Atlas scales each region and shard independently within them. Omit to keep autoscaling off.
type ApiSearchAutoScaling struct {
	Compute *ApiSearchComputeAutoScaling `json:"compute,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiSearchAutoScaling) MarshalJSON() ([]byte, error) {
	type noMethod ApiSearchAutoScaling
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiSearchAutoScaling instantiates a new ApiSearchAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchAutoScaling() *ApiSearchAutoScaling {
	this := ApiSearchAutoScaling{}
	return &this
}

// NewApiSearchAutoScalingWithDefaults instantiates a new ApiSearchAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchAutoScalingWithDefaults() *ApiSearchAutoScaling {
	this := ApiSearchAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise
func (o *ApiSearchAutoScaling) GetCompute() ApiSearchComputeAutoScaling {
	if o == nil || IsNil(o.Compute) {
		var ret ApiSearchComputeAutoScaling
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchAutoScaling) GetComputeOk() (*ApiSearchComputeAutoScaling, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}

	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ApiSearchAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ApiSearchComputeAutoScaling and assigns it to the Compute field.
func (o *ApiSearchAutoScaling) SetCompute(v ApiSearchComputeAutoScaling) {
	o.Compute = &v
	o.NullFields = removeNullField(o.NullFields, "Compute")
}

// SetComputeNil sets Compute to an explicit JSON null when marshaled.
func (o *ApiSearchAutoScaling) SetComputeNil() {
	o.Compute = nil
	o.NullFields = addNullField(o.NullFields, "Compute")
}
