// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// PreferredHostname struct for PreferredHostname
type PreferredHostname struct {
	EndsWith *bool   `json:"endsWith,omitempty"`
	Id       *string `json:"id,omitempty"`
	Regexp   *bool   `json:"regexp,omitempty"`
	Value    *string `json:"value,omitempty"`
}

// NewPreferredHostname instantiates a new PreferredHostname object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferredHostname() *PreferredHostname {
	this := PreferredHostname{}
	return &this
}

// NewPreferredHostnameWithDefaults instantiates a new PreferredHostname object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferredHostnameWithDefaults() *PreferredHostname {
	this := PreferredHostname{}
	return &this
}

// GetEndsWith returns the EndsWith field value if set, zero value otherwise
func (o *PreferredHostname) GetEndsWith() bool {
	if o == nil || IsNil(o.EndsWith) {
		var ret bool
		return ret
	}
	return *o.EndsWith
}

// GetEndsWithOk returns a tuple with the EndsWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredHostname) GetEndsWithOk() (*bool, bool) {
	if o == nil || IsNil(o.EndsWith) {
		return nil, false
	}

	return o.EndsWith, true
}

// HasEndsWith returns a boolean if a field has been set.
func (o *PreferredHostname) HasEndsWith() bool {
	if o != nil && !IsNil(o.EndsWith) {
		return true
	}

	return false
}

// SetEndsWith gets a reference to the given bool and assigns it to the EndsWith field.
func (o *PreferredHostname) SetEndsWith(v bool) {
	o.EndsWith = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *PreferredHostname) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredHostname) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PreferredHostname) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PreferredHostname) SetId(v string) {
	o.Id = &v
}

// GetRegexp returns the Regexp field value if set, zero value otherwise
func (o *PreferredHostname) GetRegexp() bool {
	if o == nil || IsNil(o.Regexp) {
		var ret bool
		return ret
	}
	return *o.Regexp
}

// GetRegexpOk returns a tuple with the Regexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredHostname) GetRegexpOk() (*bool, bool) {
	if o == nil || IsNil(o.Regexp) {
		return nil, false
	}

	return o.Regexp, true
}

// HasRegexp returns a boolean if a field has been set.
func (o *PreferredHostname) HasRegexp() bool {
	if o != nil && !IsNil(o.Regexp) {
		return true
	}

	return false
}

// SetRegexp gets a reference to the given bool and assigns it to the Regexp field.
func (o *PreferredHostname) SetRegexp(v bool) {
	o.Regexp = &v
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *PreferredHostname) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredHostname) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PreferredHostname) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PreferredHostname) SetValue(v string) {
	o.Value = &v
}

func (o PreferredHostname) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PreferredHostname) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndsWith) {
		toSerialize["endsWith"] = o.EndsWith
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Regexp) {
		toSerialize["regexp"] = o.Regexp
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}
