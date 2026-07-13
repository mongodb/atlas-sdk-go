// Code based on the AtlasAPI V2 OpenAPI file

package admin

// CreateEphemeralClusterRequest Create an ephemeral Atlas cluster.
type CreateEphemeralClusterRequest struct {
	// Designate the name of the ephemeral cluster.
	ClusterName *string `json:"clusterName,omitempty"`
}

// NewCreateEphemeralClusterRequest instantiates a new CreateEphemeralClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEphemeralClusterRequest() *CreateEphemeralClusterRequest {
	this := CreateEphemeralClusterRequest{}
	return &this
}

// NewCreateEphemeralClusterRequestWithDefaults instantiates a new CreateEphemeralClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEphemeralClusterRequestWithDefaults() *CreateEphemeralClusterRequest {
	this := CreateEphemeralClusterRequest{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *CreateEphemeralClusterRequest) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralClusterRequest) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *CreateEphemeralClusterRequest) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *CreateEphemeralClusterRequest) SetClusterName(v string) {
	o.ClusterName = &v
}
