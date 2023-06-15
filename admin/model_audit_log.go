// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AuditLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditLog{}

// AuditLog struct for AuditLog
type AuditLog struct {
	// Flag that indicates whether someone set auditing to track successful authentications. This only applies to the `\"atype\" : \"authCheck\"` audit filter. Setting this parameter to `true` degrades cluster performance.
	AuditAuthorizationSuccess bool `json:"auditAuthorizationSuccess"`
	// JSON document that specifies which events to record. Escape any characters that may prevent parsing, such as single or double quotes, using a backslash (`\\`).
	AuditFilter string `json:"auditFilter"`
	// Human-readable label that displays how to configure the audit filter.
	ConfigurationType *string `json:"configurationType,omitempty"`
	// Flag that indicates whether someone enabled database auditing for the specified project.
	Enabled bool `json:"enabled"`
}

// NewAuditLog instantiates a new AuditLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditLog(auditAuthorizationSuccess bool, auditFilter string, enabled bool) *AuditLog {
	this := AuditLog{}
	this.AuditAuthorizationSuccess = auditAuthorizationSuccess
	this.AuditFilter = auditFilter
	this.Enabled = enabled
	return &this
}

// NewAuditLogWithDefaults instantiates a new AuditLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditLogWithDefaults() *AuditLog {
	this := AuditLog{}
	var auditAuthorizationSuccess bool = false
	this.AuditAuthorizationSuccess = auditAuthorizationSuccess
	var enabled bool = false
	this.Enabled = enabled
	return &this
}

// GetAuditAuthorizationSuccess returns the AuditAuthorizationSuccess field value
func (o *AuditLog) GetAuditAuthorizationSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuditAuthorizationSuccess
}

// GetAuditAuthorizationSuccessOk returns a tuple with the AuditAuthorizationSuccess field value
// and a boolean to check if the value has been set.
func (o *AuditLog) GetAuditAuthorizationSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditAuthorizationSuccess, true
}

// SetAuditAuthorizationSuccess sets field value
func (o *AuditLog) SetAuditAuthorizationSuccess(v bool) {
	o.AuditAuthorizationSuccess = v
}

// GetAuditFilter returns the AuditFilter field value
func (o *AuditLog) GetAuditFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuditFilter
}

// GetAuditFilterOk returns a tuple with the AuditFilter field value
// and a boolean to check if the value has been set.
func (o *AuditLog) GetAuditFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuditFilter, true
}

// SetAuditFilter sets field value
func (o *AuditLog) SetAuditFilter(v string) {
	o.AuditFilter = v
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise.
func (o *AuditLog) GetConfigurationType() string {
	if o == nil || IsNil(o.ConfigurationType) {
		var ret string
		return ret
	}
	return *o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditLog) GetConfigurationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationType) {
		return nil, false
	}
	return o.ConfigurationType, true
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *AuditLog) HasConfigurationType() bool {
	if o != nil && !IsNil(o.ConfigurationType) {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given string and assigns it to the ConfigurationType field.
func (o *AuditLog) SetConfigurationType(v string) {
	o.ConfigurationType = &v
}

// GetEnabled returns the Enabled field value
func (o *AuditLog) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AuditLog) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AuditLog) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AuditLog) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AuditLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auditAuthorizationSuccess"] = o.AuditAuthorizationSuccess
	toSerialize["auditFilter"] = o.AuditFilter
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAuditLog struct {
	value *AuditLog
	isSet bool
}

func (v NullableAuditLog) Get() *AuditLog {
	return v.value
}

func (v *NullableAuditLog) Set(val *AuditLog) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditLog) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditLog(val *AuditLog) *NullableAuditLog {
	return &NullableAuditLog{value: val, isSet: true}
}

func (v NullableAuditLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
