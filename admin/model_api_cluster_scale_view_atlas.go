// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiClusterScaleViewAtlas Information on scaling the cluster to a tier that supports the requested feature. Present only when the feature is unavailable on the cluster's current tier.
type ApiClusterScaleViewAtlas struct {
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
func (o *ApiClusterScaleViewAtlas) MarshalJSON() ([]byte, error) {
	type noMethod ApiClusterScaleViewAtlas
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiClusterScaleViewAtlas instantiates a new ApiClusterScaleViewAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClusterScaleViewAtlas() *ApiClusterScaleViewAtlas {
	this := ApiClusterScaleViewAtlas{}
	return &this
}

// NewApiClusterScaleViewAtlasWithDefaults instantiates a new ApiClusterScaleViewAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClusterScaleViewAtlasWithDefaults() *ApiClusterScaleViewAtlas {
	this := ApiClusterScaleViewAtlas{}
	return &this
}

// GetDocsUrl returns the DocsUrl field value if set, zero value otherwise
func (o *ApiClusterScaleViewAtlas) GetDocsUrl() string {
	if o == nil || IsNil(o.DocsUrl) {
		var ret string
		return ret
	}
	return *o.DocsUrl
}

// GetDocsUrlOk returns a tuple with the DocsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterScaleViewAtlas) GetDocsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocsUrl) {
		return nil, false
	}

	return o.DocsUrl, true
}

// HasDocsUrl returns a boolean if a field has been set.
func (o *ApiClusterScaleViewAtlas) HasDocsUrl() bool {
	if o != nil && !IsNil(o.DocsUrl) {
		return true
	}

	return false
}

// SetDocsUrl gets a reference to the given string and assigns it to the DocsUrl field.
func (o *ApiClusterScaleViewAtlas) SetDocsUrl(v string) {
	o.DocsUrl = &v
	o.NullFields = removeNullField(o.NullFields, "DocsUrl")
}

// SetDocsUrlNil sets DocsUrl to an explicit JSON null when marshaled.
func (o *ApiClusterScaleViewAtlas) SetDocsUrlNil() {
	o.DocsUrl = nil
	o.NullFields = addNullField(o.NullFields, "DocsUrl")
}

// GetMinimumTier returns the MinimumTier field value if set, zero value otherwise
func (o *ApiClusterScaleViewAtlas) GetMinimumTier() string {
	if o == nil || IsNil(o.MinimumTier) {
		var ret string
		return ret
	}
	return *o.MinimumTier
}

// GetMinimumTierOk returns a tuple with the MinimumTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterScaleViewAtlas) GetMinimumTierOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumTier) {
		return nil, false
	}

	return o.MinimumTier, true
}

// HasMinimumTier returns a boolean if a field has been set.
func (o *ApiClusterScaleViewAtlas) HasMinimumTier() bool {
	if o != nil && !IsNil(o.MinimumTier) {
		return true
	}

	return false
}

// SetMinimumTier gets a reference to the given string and assigns it to the MinimumTier field.
func (o *ApiClusterScaleViewAtlas) SetMinimumTier(v string) {
	o.MinimumTier = &v
	o.NullFields = removeNullField(o.NullFields, "MinimumTier")
}

// SetMinimumTierNil sets MinimumTier to an explicit JSON null when marshaled.
func (o *ApiClusterScaleViewAtlas) SetMinimumTierNil() {
	o.MinimumTier = nil
	o.NullFields = addNullField(o.NullFields, "MinimumTier")
}

// GetScaleUrl returns the ScaleUrl field value if set, zero value otherwise
func (o *ApiClusterScaleViewAtlas) GetScaleUrl() string {
	if o == nil || IsNil(o.ScaleUrl) {
		var ret string
		return ret
	}
	return *o.ScaleUrl
}

// GetScaleUrlOk returns a tuple with the ScaleUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClusterScaleViewAtlas) GetScaleUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ScaleUrl) {
		return nil, false
	}

	return o.ScaleUrl, true
}

// HasScaleUrl returns a boolean if a field has been set.
func (o *ApiClusterScaleViewAtlas) HasScaleUrl() bool {
	if o != nil && !IsNil(o.ScaleUrl) {
		return true
	}

	return false
}

// SetScaleUrl gets a reference to the given string and assigns it to the ScaleUrl field.
func (o *ApiClusterScaleViewAtlas) SetScaleUrl(v string) {
	o.ScaleUrl = &v
	o.NullFields = removeNullField(o.NullFields, "ScaleUrl")
}

// SetScaleUrlNil sets ScaleUrl to an explicit JSON null when marshaled.
func (o *ApiClusterScaleViewAtlas) SetScaleUrlNil() {
	o.ScaleUrl = nil
	o.NullFields = addNullField(o.NullFields, "ScaleUrl")
}
