// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelsApiKeyResponse struct for AiModelsApiKeyResponse
type AiModelsApiKeyResponse struct {
	// Identifier used to reference this API key in admin API calls.
	// Read only field.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// UTC date when the API key was created. This parameter is formatted as an ISO 8601 timestamp.
	// Read only field.
	CreatedAt *string `json:"createdAt,omitempty"`
	// Name of the user that created this API key. If no user name is available, the user ID is returned.
	// Read only field.
	CreatedBy *string `json:"createdBy,omitempty"`
	// ID of the Atlas group this API key belongs to.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// UTC date when the API key was last used. This parameter is formatted as an ISO 8601 timestamp.
	// Read only field.
	LastUsedAt *string `json:"lastUsedAt,omitempty"`
	// A partially obfuscated version of the API key secret returned when the API key was created.
	// Read only field.
	MaskedSecret *string `json:"maskedSecret,omitempty"`
	// Arbitrary string identifier assigned to this API key for convenient identification.
	Name *string `json:"name,omitempty"`
	// The full API key secret used for interacting with the embedding / reranking service. Note: this will only be fully populated in the response to a create API key request. Responses to get, list, and update requests will not include the secret.
	// Read only field.
	Secret *string `json:"secret,omitempty"`
	// A string describing the current status of the API key.
	// Read only field.
	Status *string `json:"status,omitempty"`
}

// NewAiModelsApiKeyResponse instantiates a new AiModelsApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelsApiKeyResponse() *AiModelsApiKeyResponse {
	this := AiModelsApiKeyResponse{}
	return &this
}

// NewAiModelsApiKeyResponseWithDefaults instantiates a new AiModelsApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelsApiKeyResponseWithDefaults() *AiModelsApiKeyResponse {
	this := AiModelsApiKeyResponse{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetApiKeyId() string {
	if o == nil || IsNil(o.ApiKeyId) {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetApiKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKeyId) {
		return nil, false
	}

	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasApiKeyId() bool {
	if o != nil && !IsNil(o.ApiKeyId) {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *AiModelsApiKeyResponse) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AiModelsApiKeyResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}

	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AiModelsApiKeyResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AiModelsApiKeyResponse) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetLastUsedAt() string {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret string
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetLastUsedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}

	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given string and assigns it to the LastUsedAt field.
func (o *AiModelsApiKeyResponse) SetLastUsedAt(v string) {
	o.LastUsedAt = &v
}

// GetMaskedSecret returns the MaskedSecret field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetMaskedSecret() string {
	if o == nil || IsNil(o.MaskedSecret) {
		var ret string
		return ret
	}
	return *o.MaskedSecret
}

// GetMaskedSecretOk returns a tuple with the MaskedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetMaskedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.MaskedSecret) {
		return nil, false
	}

	return o.MaskedSecret, true
}

// HasMaskedSecret returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasMaskedSecret() bool {
	if o != nil && !IsNil(o.MaskedSecret) {
		return true
	}

	return false
}

// SetMaskedSecret gets a reference to the given string and assigns it to the MaskedSecret field.
func (o *AiModelsApiKeyResponse) SetMaskedSecret(v string) {
	o.MaskedSecret = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiModelsApiKeyResponse) SetName(v string) {
	o.Name = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}

	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AiModelsApiKeyResponse) SetSecret(v string) {
	o.Secret = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *AiModelsApiKeyResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiModelsApiKeyResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiModelsApiKeyResponse) SetStatus(v string) {
	o.Status = &v
}
