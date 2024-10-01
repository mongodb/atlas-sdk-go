// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// FieldViolation struct for FieldViolation
type FieldViolation struct {
	// A description of why the request element is bad.
	Description *string `json:"description,omitempty"`
	// A path that leads to a field in the request body.
	Field *string `json:"field,omitempty"`
}

// NewFieldViolation instantiates a new FieldViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldViolation() *FieldViolation {
	this := FieldViolation{}
	return &this
}

// NewFieldViolationWithDefaults instantiates a new FieldViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldViolationWithDefaults() *FieldViolation {
	this := FieldViolation{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *FieldViolation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldViolation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FieldViolation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FieldViolation) SetDescription(v string) {
	o.Description = &v
}

// GetField returns the Field field value if set, zero value otherwise
func (o *FieldViolation) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldViolation) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}

	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *FieldViolation) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *FieldViolation) SetField(v string) {
	o.Field = &v
}

func (o *FieldViolation) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o *FieldViolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	return toSerialize, nil
}
