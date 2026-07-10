// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelsRateLimitResponse struct for AiModelsRateLimitResponse
type AiModelsRateLimitResponse struct {
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
}

// NewAiModelsRateLimitResponse instantiates a new AiModelsRateLimitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelsRateLimitResponse() *AiModelsRateLimitResponse {
	this := AiModelsRateLimitResponse{}
	return &this
}

// NewAiModelsRateLimitResponseWithDefaults instantiates a new AiModelsRateLimitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelsRateLimitResponseWithDefaults() *AiModelsRateLimitResponse {
	this := AiModelsRateLimitResponse{}
	return &this
}

// GetModelGroupName returns the ModelGroupName field value if set, zero value otherwise
func (o *AiModelsRateLimitResponse) GetModelGroupName() string {
	if o == nil || IsNil(o.ModelGroupName) {
		var ret string
		return ret
	}
	return *o.ModelGroupName
}

// GetModelGroupNameOk returns a tuple with the ModelGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsRateLimitResponse) GetModelGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModelGroupName) {
		return nil, false
	}

	return o.ModelGroupName, true
}

// HasModelGroupName returns a boolean if a field has been set.
func (o *AiModelsRateLimitResponse) HasModelGroupName() bool {
	if o != nil && !IsNil(o.ModelGroupName) {
		return true
	}

	return false
}

// SetModelGroupName gets a reference to the given string and assigns it to the ModelGroupName field.
func (o *AiModelsRateLimitResponse) SetModelGroupName(v string) {
	o.ModelGroupName = &v
}

// GetModelNames returns the ModelNames field value if set, zero value otherwise
func (o *AiModelsRateLimitResponse) GetModelNames() []string {
	if o == nil || IsNil(o.ModelNames) {
		var ret []string
		return ret
	}
	return *o.ModelNames
}

// GetModelNamesOk returns a tuple with the ModelNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsRateLimitResponse) GetModelNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.ModelNames) {
		return nil, false
	}

	return o.ModelNames, true
}

// HasModelNames returns a boolean if a field has been set.
func (o *AiModelsRateLimitResponse) HasModelNames() bool {
	if o != nil && !IsNil(o.ModelNames) {
		return true
	}

	return false
}

// SetModelNames gets a reference to the given []string and assigns it to the ModelNames field.
func (o *AiModelsRateLimitResponse) SetModelNames(v []string) {
	o.ModelNames = &v
}

// GetRequestsPerMinuteLimit returns the RequestsPerMinuteLimit field value if set, zero value otherwise
func (o *AiModelsRateLimitResponse) GetRequestsPerMinuteLimit() int {
	if o == nil || IsNil(o.RequestsPerMinuteLimit) {
		var ret int
		return ret
	}
	return *o.RequestsPerMinuteLimit
}

// GetRequestsPerMinuteLimitOk returns a tuple with the RequestsPerMinuteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsRateLimitResponse) GetRequestsPerMinuteLimitOk() (*int, bool) {
	if o == nil || IsNil(o.RequestsPerMinuteLimit) {
		return nil, false
	}

	return o.RequestsPerMinuteLimit, true
}

// HasRequestsPerMinuteLimit returns a boolean if a field has been set.
func (o *AiModelsRateLimitResponse) HasRequestsPerMinuteLimit() bool {
	if o != nil && !IsNil(o.RequestsPerMinuteLimit) {
		return true
	}

	return false
}

// SetRequestsPerMinuteLimit gets a reference to the given int and assigns it to the RequestsPerMinuteLimit field.
func (o *AiModelsRateLimitResponse) SetRequestsPerMinuteLimit(v int) {
	o.RequestsPerMinuteLimit = &v
}

// GetTokensPerMinuteLimit returns the TokensPerMinuteLimit field value if set, zero value otherwise
func (o *AiModelsRateLimitResponse) GetTokensPerMinuteLimit() int {
	if o == nil || IsNil(o.TokensPerMinuteLimit) {
		var ret int
		return ret
	}
	return *o.TokensPerMinuteLimit
}

// GetTokensPerMinuteLimitOk returns a tuple with the TokensPerMinuteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsRateLimitResponse) GetTokensPerMinuteLimitOk() (*int, bool) {
	if o == nil || IsNil(o.TokensPerMinuteLimit) {
		return nil, false
	}

	return o.TokensPerMinuteLimit, true
}

// HasTokensPerMinuteLimit returns a boolean if a field has been set.
func (o *AiModelsRateLimitResponse) HasTokensPerMinuteLimit() bool {
	if o != nil && !IsNil(o.TokensPerMinuteLimit) {
		return true
	}

	return false
}

// SetTokensPerMinuteLimit gets a reference to the given int and assigns it to the TokensPerMinuteLimit field.
func (o *AiModelsRateLimitResponse) SetTokensPerMinuteLimit(v int) {
	o.TokensPerMinuteLimit = &v
}
