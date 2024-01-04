// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ResourceTag Key-value pair that tags and categorizes a MongoDB Cloud organization, project, or cluster. For example, `environment : production`.
type ResourceTag struct {
	// Constant that defines the set of the tag. For example, `environment` in the `environment : production` tag.
	Key *string `json:"key,omitempty"`
	// Variable that belongs to the set of the tag. For example, `production` in the `environment : production` tag.
	Value *string `json:"value,omitempty"`
}

// NewResourceTag instantiates a new ResourceTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTag() *ResourceTag {
	this := ResourceTag{}
	return &this
}

// NewResourceTagWithDefaults instantiates a new ResourceTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTagWithDefaults() *ResourceTag {
	this := ResourceTag{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise
func (o *ResourceTag) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTag) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}

	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ResourceTag) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ResourceTag) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *ResourceTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResourceTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResourceTag) SetValue(v string) {
	o.Value = &v
}

func (o ResourceTag) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ResourceTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}
