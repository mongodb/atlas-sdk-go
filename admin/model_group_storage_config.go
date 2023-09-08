// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GroupStorageConfig struct for GroupStorageConfig
type GroupStorageConfig struct {
	Mode *string `json:"mode,omitempty"`
}

// NewGroupStorageConfig instantiates a new GroupStorageConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupStorageConfig() *GroupStorageConfig {
	this := GroupStorageConfig{}
	return &this
}

// NewGroupStorageConfigWithDefaults instantiates a new GroupStorageConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupStorageConfigWithDefaults() *GroupStorageConfig {
	this := GroupStorageConfig{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise
func (o *GroupStorageConfig) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStorageConfig) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}

	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GroupStorageConfig) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GroupStorageConfig) SetMode(v string) {
	o.Mode = &v
}

func (o GroupStorageConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupStorageConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}
