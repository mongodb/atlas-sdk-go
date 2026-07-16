// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AiModelApiKeyCreateRequest struct for AiModelApiKeyCreateRequest
type AiModelApiKeyCreateRequest struct {
	// Cloud provider scope for this API key. Must be \"ANY\". Additional cloud values will be supported in future API versions.
	// Write only field.
	Cloud string `json:"cloud"`
	// Geography scope for this API key. Must be \"ANY\". Additional geography values will be supported in future API versions.
	// Write only field.
	Geography string `json:"geography"`
	// A name for the new API key that will be created.
	Name string `json:"name"`
}

// NewAiModelApiKeyCreateRequest instantiates a new AiModelApiKeyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModelApiKeyCreateRequest(cloud string, geography string, name string) *AiModelApiKeyCreateRequest {
	this := AiModelApiKeyCreateRequest{}
	this.Cloud = cloud
	this.Geography = geography
	this.Name = name
	return &this
}

// NewAiModelApiKeyCreateRequestWithDefaults instantiates a new AiModelApiKeyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelApiKeyCreateRequestWithDefaults() *AiModelApiKeyCreateRequest {
	this := AiModelApiKeyCreateRequest{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AiModelApiKeyCreateRequest) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AiModelApiKeyCreateRequest) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AiModelApiKeyCreateRequest) SetCloud(v string) {
	o.Cloud = v
}

// GetGeography returns the Geography field value
func (o *AiModelApiKeyCreateRequest) GetGeography() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Geography
}

// GetGeographyOk returns a tuple with the Geography field value
// and a boolean to check if the value has been set.
func (o *AiModelApiKeyCreateRequest) GetGeographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Geography, true
}

// SetGeography sets field value
func (o *AiModelApiKeyCreateRequest) SetGeography(v string) {
	o.Geography = v
}

// GetName returns the Name field value
func (o *AiModelApiKeyCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiModelApiKeyCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiModelApiKeyCreateRequest) SetName(v string) {
	o.Name = v
}
