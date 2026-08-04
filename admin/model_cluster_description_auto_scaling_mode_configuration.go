// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterDescriptionAutoScalingModeConfiguration Contains the internal configuration of AutoScaling on sharded clusters.
type ClusterDescriptionAutoScalingModeConfiguration struct {
	// Describes whether cluster nodes scale together across all shards or independently.
	AutoScalingMode *string `json:"autoScalingMode,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ClusterDescriptionAutoScalingModeConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod ClusterDescriptionAutoScalingModeConfiguration
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewClusterDescriptionAutoScalingModeConfiguration instantiates a new ClusterDescriptionAutoScalingModeConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDescriptionAutoScalingModeConfiguration() *ClusterDescriptionAutoScalingModeConfiguration {
	this := ClusterDescriptionAutoScalingModeConfiguration{}
	return &this
}

// NewClusterDescriptionAutoScalingModeConfigurationWithDefaults instantiates a new ClusterDescriptionAutoScalingModeConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDescriptionAutoScalingModeConfigurationWithDefaults() *ClusterDescriptionAutoScalingModeConfiguration {
	this := ClusterDescriptionAutoScalingModeConfiguration{}
	return &this
}

// GetAutoScalingMode returns the AutoScalingMode field value if set, zero value otherwise
func (o *ClusterDescriptionAutoScalingModeConfiguration) GetAutoScalingMode() string {
	if o == nil || IsNil(o.AutoScalingMode) {
		var ret string
		return ret
	}
	return *o.AutoScalingMode
}

// GetAutoScalingModeOk returns a tuple with the AutoScalingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDescriptionAutoScalingModeConfiguration) GetAutoScalingModeOk() (*string, bool) {
	if o == nil || IsNil(o.AutoScalingMode) {
		return nil, false
	}

	return o.AutoScalingMode, true
}

// HasAutoScalingMode returns a boolean if a field has been set.
func (o *ClusterDescriptionAutoScalingModeConfiguration) HasAutoScalingMode() bool {
	if o != nil && !IsNil(o.AutoScalingMode) {
		return true
	}

	return false
}

// SetAutoScalingMode gets a reference to the given string and assigns it to the AutoScalingMode field.
func (o *ClusterDescriptionAutoScalingModeConfiguration) SetAutoScalingMode(v string) {
	o.AutoScalingMode = &v
	o.NullFields = removeNullField(o.NullFields, "AutoScalingMode")
}

// SetAutoScalingModeNil sets AutoScalingMode to an explicit JSON null when marshaled.
func (o *ClusterDescriptionAutoScalingModeConfiguration) SetAutoScalingModeNil() {
	o.AutoScalingMode = nil
	o.NullFields = addNullField(o.NullFields, "AutoScalingMode")
}
