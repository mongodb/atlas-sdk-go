// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelsApiKeyUpdateRequest struct for AiModelsApiKeyUpdateRequest
type AiModelsApiKeyUpdateRequest struct {
	// A new name for the API key.
	Name string `json:"name"`
}

// NewAiModelsApiKeyUpdateRequest instantiates a new AiModelsApiKeyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelsApiKeyUpdateRequest(name string) *AiModelsApiKeyUpdateRequest {
	this := AiModelsApiKeyUpdateRequest{}
	this.Name = name
	return &this
}

// NewAiModelsApiKeyUpdateRequestWithDefaults instantiates a new AiModelsApiKeyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelsApiKeyUpdateRequestWithDefaults() *AiModelsApiKeyUpdateRequest {
	this := AiModelsApiKeyUpdateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AiModelsApiKeyUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiModelsApiKeyUpdateRequest) SetName(v string) {
	o.Name = v
}
