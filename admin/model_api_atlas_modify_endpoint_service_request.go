// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasModifyEndpointServiceRequest struct for ApiAtlasModifyEndpointServiceRequest
type ApiAtlasModifyEndpointServiceRequest struct {
	// Human-readable label that identifies the cloud service provider for the private endpoint service which you want to update.
	CloudProvider string `json:"cloudProvider"`
	// List of regions that the endpoint service supports. Native cross region support is implemented for AWS only.
	SupportedRemoteRegions []string `json:"supportedRemoteRegions"`
}

// NewApiAtlasModifyEndpointServiceRequest instantiates a new ApiAtlasModifyEndpointServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasModifyEndpointServiceRequest(cloudProvider string, supportedRemoteRegions []string) *ApiAtlasModifyEndpointServiceRequest {
	this := ApiAtlasModifyEndpointServiceRequest{}
	this.CloudProvider = cloudProvider
	this.SupportedRemoteRegions = supportedRemoteRegions
	return &this
}

// NewApiAtlasModifyEndpointServiceRequestWithDefaults instantiates a new ApiAtlasModifyEndpointServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasModifyEndpointServiceRequestWithDefaults() *ApiAtlasModifyEndpointServiceRequest {
	this := ApiAtlasModifyEndpointServiceRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *ApiAtlasModifyEndpointServiceRequest) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasModifyEndpointServiceRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ApiAtlasModifyEndpointServiceRequest) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetSupportedRemoteRegions returns the SupportedRemoteRegions field value
func (o *ApiAtlasModifyEndpointServiceRequest) GetSupportedRemoteRegions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupportedRemoteRegions
}

// GetSupportedRemoteRegionsOk returns a tuple with the SupportedRemoteRegions field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasModifyEndpointServiceRequest) GetSupportedRemoteRegionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedRemoteRegions, true
}

// SetSupportedRemoteRegions sets field value
func (o *ApiAtlasModifyEndpointServiceRequest) SetSupportedRemoteRegions(v []string) {
	o.SupportedRemoteRegions = v
}
