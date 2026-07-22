// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelsRateLimitUpdateRequest struct for AiModelsRateLimitUpdateRequest
type AiModelsRateLimitUpdateRequest struct {
	// The number of requests per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model.
	RequestsPerMinuteLimit int `json:"requestsPerMinuteLimit"`
	// The number of tokens per minute allowed for this model group. Must be a positive integer. Cannot be more than the organization level limit for this group model.
	TokensPerMinuteLimit int `json:"tokensPerMinuteLimit"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AiModelsRateLimitUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod AiModelsRateLimitUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAiModelsRateLimitUpdateRequest instantiates a new AiModelsRateLimitUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelsRateLimitUpdateRequest(requestsPerMinuteLimit int, tokensPerMinuteLimit int) *AiModelsRateLimitUpdateRequest {
	this := AiModelsRateLimitUpdateRequest{}
	this.RequestsPerMinuteLimit = requestsPerMinuteLimit
	this.TokensPerMinuteLimit = tokensPerMinuteLimit
	return &this
}

// NewAiModelsRateLimitUpdateRequestWithDefaults instantiates a new AiModelsRateLimitUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelsRateLimitUpdateRequestWithDefaults() *AiModelsRateLimitUpdateRequest {
	this := AiModelsRateLimitUpdateRequest{}
	return &this
}

// GetRequestsPerMinuteLimit returns the RequestsPerMinuteLimit field value
func (o *AiModelsRateLimitUpdateRequest) GetRequestsPerMinuteLimit() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.RequestsPerMinuteLimit
}

// GetRequestsPerMinuteLimitOk returns a tuple with the RequestsPerMinuteLimit field value
// and a boolean to check if the value has been set.
func (o *AiModelsRateLimitUpdateRequest) GetRequestsPerMinuteLimitOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPerMinuteLimit, true
}

// SetRequestsPerMinuteLimit sets field value
func (o *AiModelsRateLimitUpdateRequest) SetRequestsPerMinuteLimit(v int) {
	o.RequestsPerMinuteLimit = v
}

// GetTokensPerMinuteLimit returns the TokensPerMinuteLimit field value
func (o *AiModelsRateLimitUpdateRequest) GetTokensPerMinuteLimit() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TokensPerMinuteLimit
}

// GetTokensPerMinuteLimitOk returns a tuple with the TokensPerMinuteLimit field value
// and a boolean to check if the value has been set.
func (o *AiModelsRateLimitUpdateRequest) GetTokensPerMinuteLimitOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokensPerMinuteLimit, true
}

// SetTokensPerMinuteLimit sets field value
func (o *AiModelsRateLimitUpdateRequest) SetTokensPerMinuteLimit(v int) {
	o.TokensPerMinuteLimit = v
}
