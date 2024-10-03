// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// FlexBackupSettings20250101 Flex backup configuration
type FlexBackupSettings20250101 struct {
	// Flag that indicates whether backups are performed for this flex cluster. Backup uses [TODO](TODO) for flex clusters.
	// Read only field.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewFlexBackupSettings20250101 instantiates a new FlexBackupSettings20250101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexBackupSettings20250101() *FlexBackupSettings20250101 {
	this := FlexBackupSettings20250101{}
	return &this
}

// NewFlexBackupSettings20250101WithDefaults instantiates a new FlexBackupSettings20250101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexBackupSettings20250101WithDefaults() *FlexBackupSettings20250101 {
	this := FlexBackupSettings20250101{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *FlexBackupSettings20250101) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexBackupSettings20250101) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *FlexBackupSettings20250101) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *FlexBackupSettings20250101) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o *FlexBackupSettings20250101) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o *FlexBackupSettings20250101) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
