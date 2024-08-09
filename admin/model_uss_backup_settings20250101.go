// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// USSBackupSettings20250101 USS backup configuration
type USSBackupSettings20250101 struct {
	// Flag that indicates whether backups are performed for this USS instance. Backup uses [TODO](TODO) for USS instances.
	// Read only field.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUSSBackupSettings20250101 instantiates a new USSBackupSettings20250101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUSSBackupSettings20250101() *USSBackupSettings20250101 {
	this := USSBackupSettings20250101{}
	return &this
}

// NewUSSBackupSettings20250101WithDefaults instantiates a new USSBackupSettings20250101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUSSBackupSettings20250101WithDefaults() *USSBackupSettings20250101 {
	this := USSBackupSettings20250101{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *USSBackupSettings20250101) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *USSBackupSettings20250101) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *USSBackupSettings20250101) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *USSBackupSettings20250101) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o USSBackupSettings20250101) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o USSBackupSettings20250101) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
