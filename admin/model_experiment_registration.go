// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ExperimentRegistration struct for ExperimentRegistration
type ExperimentRegistration struct {
	ExperimentId *string `json:"experimentId,omitempty"`
	Key          *string `json:"key,omitempty"`
	VariationId  *int    `json:"variationId,omitempty"`
}

// NewExperimentRegistration instantiates a new ExperimentRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentRegistration() *ExperimentRegistration {
	this := ExperimentRegistration{}
	return &this
}

// NewExperimentRegistrationWithDefaults instantiates a new ExperimentRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentRegistrationWithDefaults() *ExperimentRegistration {
	this := ExperimentRegistration{}
	return &this
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise
func (o *ExperimentRegistration) GetExperimentId() string {
	if o == nil || IsNil(o.ExperimentId) {
		var ret string
		return ret
	}
	return *o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentRegistration) GetExperimentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExperimentId) {
		return nil, false
	}

	return o.ExperimentId, true
}

// HasExperimentId returns a boolean if a field has been set.
func (o *ExperimentRegistration) HasExperimentId() bool {
	if o != nil && !IsNil(o.ExperimentId) {
		return true
	}

	return false
}

// SetExperimentId gets a reference to the given string and assigns it to the ExperimentId field.
func (o *ExperimentRegistration) SetExperimentId(v string) {
	o.ExperimentId = &v
}

// GetKey returns the Key field value if set, zero value otherwise
func (o *ExperimentRegistration) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentRegistration) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}

	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ExperimentRegistration) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ExperimentRegistration) SetKey(v string) {
	o.Key = &v
}

// GetVariationId returns the VariationId field value if set, zero value otherwise
func (o *ExperimentRegistration) GetVariationId() int {
	if o == nil || IsNil(o.VariationId) {
		var ret int
		return ret
	}
	return *o.VariationId
}

// GetVariationIdOk returns a tuple with the VariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentRegistration) GetVariationIdOk() (*int, bool) {
	if o == nil || IsNil(o.VariationId) {
		return nil, false
	}

	return o.VariationId, true
}

// HasVariationId returns a boolean if a field has been set.
func (o *ExperimentRegistration) HasVariationId() bool {
	if o != nil && !IsNil(o.VariationId) {
		return true
	}

	return false
}

// SetVariationId gets a reference to the given int and assigns it to the VariationId field.
func (o *ExperimentRegistration) SetVariationId(v int) {
	o.VariationId = &v
}

func (o ExperimentRegistration) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ExperimentRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExperimentId) {
		toSerialize["experimentId"] = o.ExperimentId
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.VariationId) {
		toSerialize["variationId"] = o.VariationId
	}
	return toSerialize, nil
}
