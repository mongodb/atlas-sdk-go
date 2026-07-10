// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiStandbyLinkFailoverRequest Request to initiate a failover operation.
type ApiStandbyLinkFailoverRequest struct {
	// Type of failover operation to perform.
	Type string `json:"type"`
}

// NewApiStandbyLinkFailoverRequest instantiates a new ApiStandbyLinkFailoverRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStandbyLinkFailoverRequest(type_ string) *ApiStandbyLinkFailoverRequest {
	this := ApiStandbyLinkFailoverRequest{}
	this.Type = type_
	return &this
}

// NewApiStandbyLinkFailoverRequestWithDefaults instantiates a new ApiStandbyLinkFailoverRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStandbyLinkFailoverRequestWithDefaults() *ApiStandbyLinkFailoverRequest {
	this := ApiStandbyLinkFailoverRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ApiStandbyLinkFailoverRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiStandbyLinkFailoverRequest) SetType(v string) {
	o.Type = v
}
