// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// EncryptionAtRestPrivateEndpointRequest Create or delete private endpoints from the specified regions list.
type EncryptionAtRestPrivateEndpointRequest struct {
	// Cloud service provider name.
	CloudProvider string `json:"cloudProvider"`
	// List of regions.
	Regions *[]string `json:"regions,omitempty"`
}

// NewEncryptionAtRestPrivateEndpointRequest instantiates a new EncryptionAtRestPrivateEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionAtRestPrivateEndpointRequest(cloudProvider string) *EncryptionAtRestPrivateEndpointRequest {
	this := EncryptionAtRestPrivateEndpointRequest{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewEncryptionAtRestPrivateEndpointRequestWithDefaults instantiates a new EncryptionAtRestPrivateEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionAtRestPrivateEndpointRequestWithDefaults() *EncryptionAtRestPrivateEndpointRequest {
	this := EncryptionAtRestPrivateEndpointRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *EncryptionAtRestPrivateEndpointRequest) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *EncryptionAtRestPrivateEndpointRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *EncryptionAtRestPrivateEndpointRequest) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetRegions returns the Regions field value if set, zero value otherwise
func (o *EncryptionAtRestPrivateEndpointRequest) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionAtRestPrivateEndpointRequest) GetRegionsOk() (*[]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}

	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *EncryptionAtRestPrivateEndpointRequest) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *EncryptionAtRestPrivateEndpointRequest) SetRegions(v []string) {
	o.Regions = &v
}

func (o EncryptionAtRestPrivateEndpointRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o EncryptionAtRestPrivateEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return toSerialize, nil
}
