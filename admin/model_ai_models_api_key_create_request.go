// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelsApiKeyCreateRequest struct for AiModelsApiKeyCreateRequest
type AiModelsApiKeyCreateRequest struct {
	// A name for the new API key that will be created.
	Name string `json:"name"`
}

// NewAiModelsApiKeyCreateRequest instantiates a new AiModelsApiKeyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelsApiKeyCreateRequest(name string) *AiModelsApiKeyCreateRequest {
	this := AiModelsApiKeyCreateRequest{}
	this.Name = name
	return &this
}

// NewAiModelsApiKeyCreateRequestWithDefaults instantiates a new AiModelsApiKeyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelsApiKeyCreateRequestWithDefaults() *AiModelsApiKeyCreateRequest {
	this := AiModelsApiKeyCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AiModelsApiKeyCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiModelsApiKeyCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiModelsApiKeyCreateRequest) SetName(v string) {
	o.Name = v
}
