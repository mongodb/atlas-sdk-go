// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelApiKeyUpdateRequest struct for AiModelApiKeyUpdateRequest
type AiModelApiKeyUpdateRequest struct {
	// A new name for the API key.
	Name string `json:"name"`
}

// NewAiModelApiKeyUpdateRequest instantiates a new AiModelApiKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelApiKeyUpdateRequest(name string) *AiModelApiKeyUpdateRequest {
	this := AiModelApiKeyUpdateRequest{}
	this.Name = name
	return &this
}

// NewAiModelApiKeyUpdateRequestWithDefaults instantiates a new AiModelApiKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelApiKeyUpdateRequestWithDefaults() *AiModelApiKeyUpdateRequest {
	this := AiModelApiKeyUpdateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AiModelApiKeyUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiModelApiKeyUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiModelApiKeyUpdateRequest) SetName(v string) {
	o.Name = v
}
