// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelRateLimitResponse struct for AiModelRateLimitResponse
type AiModelRateLimitResponse struct {
	// Cloud provider scope for this rate limit. Use \"ANY\" for cloud-agnostic scope.
	// Read only field.
	Cloud *string `json:"cloud,omitempty"`
	// Server-computed endpoint hostname derived from `cloud` and `geography`. This field is read-only and must not be supplied in request bodies.
	// Read only field.
	Endpoint *string `json:"endpoint,omitempty"`
	// Geography scope for this rate limit. Use \"ANY\" for geography-agnostic scope.
	// Read only field.
	Geography *string `json:"geography,omitempty"`
	// Identifier used to reference this model group.
	// Read only field.
	ModelGroupName *string `json:"modelGroupName,omitempty"`
	// List of embedding model names included in this model group.
	// Read only field.
	ModelNames *[]string `json:"modelNames,omitempty"`
	// The number of requests per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model.
	RequestsPerMinuteLimit *int `json:"requestsPerMinuteLimit,omitempty"`
	// The number of tokens per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model.
	TokensPerMinuteLimit *int `json:"tokensPerMinuteLimit,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AiModelRateLimitResponse) MarshalJSON() ([]byte, error) {
	type noMethod AiModelRateLimitResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAiModelRateLimitResponse instantiates a new AiModelRateLimitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelRateLimitResponse() *AiModelRateLimitResponse {
	this := AiModelRateLimitResponse{}
	return &this
}

// NewAiModelRateLimitResponseWithDefaults instantiates a new AiModelRateLimitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelRateLimitResponseWithDefaults() *AiModelRateLimitResponse {
	this := AiModelRateLimitResponse{}
	return &this
}

// GetCloud returns the Cloud field value if set, zero value otherwise
func (o *AiModelRateLimitResponse) GetCloud() string {
	if o == nil || IsNil(o.Cloud) {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelRateLimitResponse) GetCloudOk() (*string, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}

	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *AiModelRateLimitResponse) HasCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *AiModelRateLimitResponse) SetCloud(v string) {
	o.Cloud = &v
	o.NullFields = removeNullField(o.NullFields, "Cloud")
}

// SetCloudNil sets Cloud to an explicit JSON null when marshaled.
func (o *AiModelRateLimitResponse) SetCloudNil() {
	o.Cloud = nil
	o.NullFields = addNullField(o.NullFields, "Cloud")
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise
func (o *AiModelRateLimitResponse) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelRateLimitResponse) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}

	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AiModelRateLimitResponse) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *AiModelRateLimitResponse) SetEndpoint(v string) {
	o.Endpoint = &v
	o.NullFields = removeNullField(o.NullFields, "Endpoint")
}

// SetEndpointNil sets Endpoint to an explicit JSON null when marshaled.
func (o *AiModelRateLimitResponse) SetEndpointNil() {
	o.Endpoint = nil
	o.NullFields = addNullField(o.NullFields, "Endpoint")
}

// GetGeography returns the Geography field value if set, zero value otherwise
func (o *AiModelRateLimitResponse) GetGeography() string {
	if o == nil || IsNil(o.Geography) {
		var ret string
		return ret
	}
	return *o.Geography
}

// GetGeographyOk returns a tuple with the Geography field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelRateLimitResponse) GetGeographyOk() (*string, bool) {
	if o == nil || IsNil(o.Geography) {
		return nil, false
	}

	return o.Geography, true
}

// HasGeography returns a boolean if a field has been set.
func (o *AiModelRateLimitResponse) HasGeography() bool {
	if o != nil && !IsNil(o.Geography) {
		return true
	}

	return false
}

// SetGeography gets a reference to the given string and assigns it to the Geography field.
func (o *AiModelRateLimitResponse) SetGeography(v string) {
	o.Geography = &v
	o.NullFields = removeNullField(o.NullFields, "Geography")
}

// SetGeographyNil sets Geography to an explicit JSON null when marshaled.
func (o *AiModelRateLimitResponse) SetGeographyNil() {
	o.Geography = nil
	o.NullFields = addNullField(o.NullFields, "Geography")
}

// GetModelGroupName returns the ModelGroupName field value if set, zero value otherwise
func (o *AiModelRateLimitResponse) GetModelGroupName() string {
	if o == nil || IsNil(o.ModelGroupName) {
		var ret string
		return ret
	}
	return *o.ModelGroupName
}

// GetModelGroupNameOk returns a tuple with the ModelGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelRateLimitResponse) GetModelGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelGroupName) {
		return nil, false
	}

	return o.ModelGroupName, true
}

// HasModelGroupName returns a boolean if a field has been set.
func (o *AiModelRateLimitResponse) HasModelGroupName() bool {
	if o != nil && !IsNil(o.ModelGroupName) {
		return true
	}

	return false
}

// SetModelGroupName gets a reference to the given string and assigns it to the ModelGroupName field.
func (o *AiModelRateLimitResponse) SetModelGroupName(v string) {
	o.ModelGroupName = &v
	o.NullFields = removeNullField(o.NullFields, "ModelGroupName")
}

// SetModelGroupNameNil sets ModelGroupName to an explicit JSON null when marshaled.
func (o *AiModelRateLimitResponse) SetModelGroupNameNil() {
	o.ModelGroupName = nil
	o.NullFields = addNullField(o.NullFields, "ModelGroupName")
}

// GetModelNames returns the ModelNames field value if set, zero value otherwise
func (o *AiModelRateLimitResponse) GetModelNames() []string {
	if o == nil || IsNil(o.ModelNames) {
		var ret []string
		return ret
	}
	return *o.ModelNames
}

// GetModelNamesOk returns a tuple with the ModelNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelRateLimitResponse) GetModelNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.ModelNames) {
		return nil, false
	}

	return o.ModelNames, true
}

// HasModelNames returns a boolean if a field has been set.
func (o *AiModelRateLimitResponse) HasModelNames() bool {
	if o != nil && !IsNil(o.ModelNames) {
		return true
	}

	return false
}

// SetModelNames gets a reference to the given []string and assigns it to the ModelNames field.
func (o *AiModelRateLimitResponse) SetModelNames(v []string) {
	o.ModelNames = &v
	o.NullFields = removeNullField(o.NullFields, "ModelNames")
}

// SetModelNamesNil sets ModelNames to an explicit JSON null when marshaled.
func (o *AiModelRateLimitResponse) SetModelNamesNil() {
	o.ModelNames = nil
	o.NullFields = addNullField(o.NullFields, "ModelNames")
}

// GetRequestsPerMinuteLimit returns the RequestsPerMinuteLimit field value if set, zero value otherwise
func (o *AiModelRateLimitResponse) GetRequestsPerMinuteLimit() int {
	if o == nil || IsNil(o.RequestsPerMinuteLimit) {
		var ret int
		return ret
	}
	return *o.RequestsPerMinuteLimit
}

// GetRequestsPerMinuteLimitOk returns a tuple with the RequestsPerMinuteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelRateLimitResponse) GetRequestsPerMinuteLimitOk() (*int, bool) {
	if o == nil || IsNil(o.RequestsPerMinuteLimit) {
		return nil, false
	}

	return o.RequestsPerMinuteLimit, true
}

// HasRequestsPerMinuteLimit returns a boolean if a field has been set.
func (o *AiModelRateLimitResponse) HasRequestsPerMinuteLimit() bool {
	if o != nil && !IsNil(o.RequestsPerMinuteLimit) {
		return true
	}

	return false
}

// SetRequestsPerMinuteLimit gets a reference to the given int and assigns it to the RequestsPerMinuteLimit field.
func (o *AiModelRateLimitResponse) SetRequestsPerMinuteLimit(v int) {
	o.RequestsPerMinuteLimit = &v
	o.NullFields = removeNullField(o.NullFields, "RequestsPerMinuteLimit")
}

// SetRequestsPerMinuteLimitNil sets RequestsPerMinuteLimit to an explicit JSON null when marshaled.
func (o *AiModelRateLimitResponse) SetRequestsPerMinuteLimitNil() {
	o.RequestsPerMinuteLimit = nil
	o.NullFields = addNullField(o.NullFields, "RequestsPerMinuteLimit")
}

// GetTokensPerMinuteLimit returns the TokensPerMinuteLimit field value if set, zero value otherwise
func (o *AiModelRateLimitResponse) GetTokensPerMinuteLimit() int {
	if o == nil || IsNil(o.TokensPerMinuteLimit) {
		var ret int
		return ret
	}
	return *o.TokensPerMinuteLimit
}

// GetTokensPerMinuteLimitOk returns a tuple with the TokensPerMinuteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelRateLimitResponse) GetTokensPerMinuteLimitOk() (*int, bool) {
	if o == nil || IsNil(o.TokensPerMinuteLimit) {
		return nil, false
	}

	return o.TokensPerMinuteLimit, true
}

// HasTokensPerMinuteLimit returns a boolean if a field has been set.
func (o *AiModelRateLimitResponse) HasTokensPerMinuteLimit() bool {
	if o != nil && !IsNil(o.TokensPerMinuteLimit) {
		return true
	}

	return false
}

// SetTokensPerMinuteLimit gets a reference to the given int and assigns it to the TokensPerMinuteLimit field.
func (o *AiModelRateLimitResponse) SetTokensPerMinuteLimit(v int) {
	o.TokensPerMinuteLimit = &v
	o.NullFields = removeNullField(o.NullFields, "TokensPerMinuteLimit")
}

// SetTokensPerMinuteLimitNil sets TokensPerMinuteLimit to an explicit JSON null when marshaled.
func (o *AiModelRateLimitResponse) SetTokensPerMinuteLimitNil() {
	o.TokensPerMinuteLimit = nil
	o.NullFields = addNullField(o.NullFields, "TokensPerMinuteLimit")
}
