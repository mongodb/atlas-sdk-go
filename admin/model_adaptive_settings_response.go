// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AdaptiveSettingsResponse Adaptive Settings for a cluster, including client-provided overrides and the effective settings derived from those overrides and Atlas-managed defaults.
type AdaptiveSettingsResponse struct {
	// Map of customer-specified overrides for Adaptive Settings, applied on a best-effort basis. Each supported entry that you specify in this object takes precedence over the corresponding Atlas-managed default. For example, if a setting is enabled by default, you can add an override to disable it for your cluster.
	AdaptiveSettingsOverrides *map[string]any `json:"adaptiveSettingsOverrides,omitempty"`
	// The effective state of Adaptive Settings currently applied to your cluster, based on your overrides and Atlas-managed defaults. Atlas-managed defaults can vary by MongoDB version, so the same setting may default differently across clusters running different versions. If you set an override for a setting that your cluster's current MongoDB version doesn't support, the override doesn't take effect and the effective value reflects the Atlas-managed default instead.
	// Read only field.
	EffectiveAdaptiveSettings map[string]any `json:"effectiveAdaptiveSettings"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AdaptiveSettingsResponse) MarshalJSON() ([]byte, error) {
	type noMethod AdaptiveSettingsResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAdaptiveSettingsResponse instantiates a new AdaptiveSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdaptiveSettingsResponse(effectiveAdaptiveSettings map[string]any) *AdaptiveSettingsResponse {
	this := AdaptiveSettingsResponse{}
	this.EffectiveAdaptiveSettings = effectiveAdaptiveSettings
	return &this
}

// NewAdaptiveSettingsResponseWithDefaults instantiates a new AdaptiveSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdaptiveSettingsResponseWithDefaults() *AdaptiveSettingsResponse {
	this := AdaptiveSettingsResponse{}
	return &this
}

// GetAdaptiveSettingsOverrides returns the AdaptiveSettingsOverrides field value if set, zero value otherwise
func (o *AdaptiveSettingsResponse) GetAdaptiveSettingsOverrides() map[string]any {
	if o == nil || IsNil(o.AdaptiveSettingsOverrides) {
		var ret map[string]any
		return ret
	}
	return *o.AdaptiveSettingsOverrides
}

// GetAdaptiveSettingsOverridesOk returns a tuple with the AdaptiveSettingsOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdaptiveSettingsResponse) GetAdaptiveSettingsOverridesOk() (*map[string]any, bool) {
	if o == nil || IsNil(o.AdaptiveSettingsOverrides) {
		return nil, false
	}

	return o.AdaptiveSettingsOverrides, true
}

// HasAdaptiveSettingsOverrides returns a boolean if a field has been set.
func (o *AdaptiveSettingsResponse) HasAdaptiveSettingsOverrides() bool {
	if o != nil && !IsNil(o.AdaptiveSettingsOverrides) {
		return true
	}

	return false
}

// SetAdaptiveSettingsOverrides gets a reference to the given map[string]any and assigns it to the AdaptiveSettingsOverrides field.
func (o *AdaptiveSettingsResponse) SetAdaptiveSettingsOverrides(v map[string]any) {
	o.AdaptiveSettingsOverrides = &v
	o.NullFields = removeNullField(o.NullFields, "AdaptiveSettingsOverrides")
}

// SetAdaptiveSettingsOverridesNil sets AdaptiveSettingsOverrides to an explicit JSON null when marshaled.
func (o *AdaptiveSettingsResponse) SetAdaptiveSettingsOverridesNil() {
	o.AdaptiveSettingsOverrides = nil
	o.NullFields = addNullField(o.NullFields, "AdaptiveSettingsOverrides")
}

// GetEffectiveAdaptiveSettings returns the EffectiveAdaptiveSettings field value
func (o *AdaptiveSettingsResponse) GetEffectiveAdaptiveSettings() map[string]any {
	if o == nil {
		var ret map[string]any
		return ret
	}

	return o.EffectiveAdaptiveSettings
}

// GetEffectiveAdaptiveSettingsOk returns a tuple with the EffectiveAdaptiveSettings field value
// and a boolean to check if the value has been set.
func (o *AdaptiveSettingsResponse) GetEffectiveAdaptiveSettingsOk() (map[string]any, bool) {
	if o == nil {
		var ret map[string]any
		return ret, false
	}
	return o.EffectiveAdaptiveSettings, true
}

// SetEffectiveAdaptiveSettings sets field value
func (o *AdaptiveSettingsResponse) SetEffectiveAdaptiveSettings(v map[string]any) {
	o.EffectiveAdaptiveSettings = v
}
