// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiClusterScale Information on scaling the cluster to a tier that supports the requested feature. Present only when the feature is unavailable on the cluster's current tier.
type ApiClusterScale struct {
	// Documentation page describing feature availability by cluster tier.
	// Read only field.
	DocsUrl *string `json:"docsUrl,omitempty"`
	// Lowest cluster tier that supports the requested feature.
	// Read only field.
	MinimumTier *string `json:"minimumTier,omitempty"`
	// Link to the Atlas UI page that scales the cluster, with the minimum supported tier pre-selected.
	// Read only field.
	ScaleUrl *string `json:"scaleUrl,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiClusterScale) MarshalJSON() ([]byte, error) {
	type noMethod ApiClusterScale
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiClusterScale instantiates a new ApiClusterScale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClusterScale() *ApiClusterScale {
	this := ApiClusterScale{}
	return &this
}

// NewApiClusterScaleWithDefaults instantiates a new ApiClusterScale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClusterScaleWithDefaults() *ApiClusterScale {
	this := ApiClusterScale{}
	return &this
}

// GetDocsUrl returns the DocsUrl field value if set, zero value otherwise
func (o *ApiClusterScale) GetDocsUrl() string {
	if o == nil || IsNil(o.DocsUrl) {
		var ret string
		return ret
	}
	return *o.DocsUrl
}

// GetDocsUrlOk returns a tuple with the DocsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterScale) GetDocsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocsUrl) {
		return nil, false
	}

	return o.DocsUrl, true
}

// HasDocsUrl returns a boolean if a field has been set.
func (o *ApiClusterScale) HasDocsUrl() bool {
	if o != nil && !IsNil(o.DocsUrl) {
		return true
	}

	return false
}

// SetDocsUrl gets a reference to the given string and assigns it to the DocsUrl field.
func (o *ApiClusterScale) SetDocsUrl(v string) {
	o.DocsUrl = &v
	o.NullFields = removeNullField(o.NullFields, "DocsUrl")
}

// SetDocsUrlNil sets DocsUrl to an explicit JSON null when marshaled.
func (o *ApiClusterScale) SetDocsUrlNil() {
	o.DocsUrl = nil
	o.NullFields = addNullField(o.NullFields, "DocsUrl")
}

// GetMinimumTier returns the MinimumTier field value if set, zero value otherwise
func (o *ApiClusterScale) GetMinimumTier() string {
	if o == nil || IsNil(o.MinimumTier) {
		var ret string
		return ret
	}
	return *o.MinimumTier
}

// GetMinimumTierOk returns a tuple with the MinimumTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterScale) GetMinimumTierOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumTier) {
		return nil, false
	}

	return o.MinimumTier, true
}

// HasMinimumTier returns a boolean if a field has been set.
func (o *ApiClusterScale) HasMinimumTier() bool {
	if o != nil && !IsNil(o.MinimumTier) {
		return true
	}

	return false
}

// SetMinimumTier gets a reference to the given string and assigns it to the MinimumTier field.
func (o *ApiClusterScale) SetMinimumTier(v string) {
	o.MinimumTier = &v
	o.NullFields = removeNullField(o.NullFields, "MinimumTier")
}

// SetMinimumTierNil sets MinimumTier to an explicit JSON null when marshaled.
func (o *ApiClusterScale) SetMinimumTierNil() {
	o.MinimumTier = nil
	o.NullFields = addNullField(o.NullFields, "MinimumTier")
}

// GetScaleUrl returns the ScaleUrl field value if set, zero value otherwise
func (o *ApiClusterScale) GetScaleUrl() string {
	if o == nil || IsNil(o.ScaleUrl) {
		var ret string
		return ret
	}
	return *o.ScaleUrl
}

// GetScaleUrlOk returns a tuple with the ScaleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterScale) GetScaleUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ScaleUrl) {
		return nil, false
	}

	return o.ScaleUrl, true
}

// HasScaleUrl returns a boolean if a field has been set.
func (o *ApiClusterScale) HasScaleUrl() bool {
	if o != nil && !IsNil(o.ScaleUrl) {
		return true
	}

	return false
}

// SetScaleUrl gets a reference to the given string and assigns it to the ScaleUrl field.
func (o *ApiClusterScale) SetScaleUrl(v string) {
	o.ScaleUrl = &v
	o.NullFields = removeNullField(o.NullFields, "ScaleUrl")
}

// SetScaleUrlNil sets ScaleUrl to an explicit JSON null when marshaled.
func (o *ApiClusterScale) SetScaleUrlNil() {
	o.ScaleUrl = nil
	o.NullFields = addNullField(o.NullFields, "ScaleUrl")
}
