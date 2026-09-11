// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AdaptiveSettingsUpdateRequest Adaptive Settings for a cluster, including client-provided overrides and the effective settings derived from those overrides and Atlas-managed defaults.
type AdaptiveSettingsUpdateRequest struct {
	// Map of customer-specified overrides for Adaptive Settings, applied on a best-effort basis. Each supported entry that you specify in this object takes precedence over the corresponding Atlas-managed default. For example, if a setting is enabled by default, you can add an override to disable it for your cluster.
	AdaptiveSettingsOverrides *map[string]any `json:"adaptiveSettingsOverrides,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AdaptiveSettingsUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod AdaptiveSettingsUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAdaptiveSettingsUpdateRequest instantiates a new AdaptiveSettingsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdaptiveSettingsUpdateRequest() *AdaptiveSettingsUpdateRequest {
	this := AdaptiveSettingsUpdateRequest{}
	return &this
}

// NewAdaptiveSettingsUpdateRequestWithDefaults instantiates a new AdaptiveSettingsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdaptiveSettingsUpdateRequestWithDefaults() *AdaptiveSettingsUpdateRequest {
	this := AdaptiveSettingsUpdateRequest{}
	return &this
}

// GetAdaptiveSettingsOverrides returns the AdaptiveSettingsOverrides field value if set, zero value otherwise
func (o *AdaptiveSettingsUpdateRequest) GetAdaptiveSettingsOverrides() map[string]any {
	if o == nil || IsNil(o.AdaptiveSettingsOverrides) {
		var ret map[string]any
		return ret
	}
	return *o.AdaptiveSettingsOverrides
}

// GetAdaptiveSettingsOverridesOk returns a tuple with the AdaptiveSettingsOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdaptiveSettingsUpdateRequest) GetAdaptiveSettingsOverridesOk() (*map[string]any, bool) {
	if o == nil || IsNil(o.AdaptiveSettingsOverrides) {
		return nil, false
	}

	return o.AdaptiveSettingsOverrides, true
}

// HasAdaptiveSettingsOverrides returns a boolean if a field has been set.
func (o *AdaptiveSettingsUpdateRequest) HasAdaptiveSettingsOverrides() bool {
	if o != nil && !IsNil(o.AdaptiveSettingsOverrides) {
		return true
	}

	return false
}

// SetAdaptiveSettingsOverrides gets a reference to the given map[string]any and assigns it to the AdaptiveSettingsOverrides field.
func (o *AdaptiveSettingsUpdateRequest) SetAdaptiveSettingsOverrides(v map[string]any) {
	o.AdaptiveSettingsOverrides = &v
	o.NullFields = removeNullField(o.NullFields, "AdaptiveSettingsOverrides")
}

// SetAdaptiveSettingsOverridesNil sets AdaptiveSettingsOverrides to an explicit JSON null when marshaled.
func (o *AdaptiveSettingsUpdateRequest) SetAdaptiveSettingsOverridesNil() {
	o.AdaptiveSettingsOverrides = nil
	o.NullFields = addNullField(o.NullFields, "AdaptiveSettingsOverrides")
}
