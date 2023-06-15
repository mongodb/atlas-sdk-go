// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the ServerlessAzureTenantEndpointUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessAzureTenantEndpointUpdate{}

// ServerlessAzureTenantEndpointUpdate Updates to a serverless Azure tenant endpoint.
type ServerlessAzureTenantEndpointUpdate struct {
	// Unique string that identifies the Azure private endpoint's network interface for this private endpoint service.
	CloudProviderEndpointId *string `json:"cloudProviderEndpointId,omitempty"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIpAddress *string `json:"privateEndpointIpAddress,omitempty"`
	// Human-readable comment associated with the private endpoint.
	Comment      *string `json:"comment,omitempty"`
	ProviderName string  `json:"providerName"`
}

// NewServerlessAzureTenantEndpointUpdate instantiates a new ServerlessAzureTenantEndpointUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessAzureTenantEndpointUpdate(providerName string) *ServerlessAzureTenantEndpointUpdate {
	this := ServerlessAzureTenantEndpointUpdate{}
	this.ProviderName = providerName
	return &this
}

// NewServerlessAzureTenantEndpointUpdateWithDefaults instantiates a new ServerlessAzureTenantEndpointUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessAzureTenantEndpointUpdateWithDefaults() *ServerlessAzureTenantEndpointUpdate {
	this := ServerlessAzureTenantEndpointUpdate{}
	return &this
}

// GetCloudProviderEndpointId returns the CloudProviderEndpointId field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpointUpdate) GetCloudProviderEndpointId() string {
	if o == nil || IsNil(o.CloudProviderEndpointId) {
		var ret string
		return ret
	}
	return *o.CloudProviderEndpointId
}

// GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpointUpdate) GetCloudProviderEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProviderEndpointId) {
		return nil, false
	}
	return o.CloudProviderEndpointId, true
}

// HasCloudProviderEndpointId returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpointUpdate) HasCloudProviderEndpointId() bool {
	if o != nil && !IsNil(o.CloudProviderEndpointId) {
		return true
	}

	return false
}

// SetCloudProviderEndpointId gets a reference to the given string and assigns it to the CloudProviderEndpointId field.
func (o *ServerlessAzureTenantEndpointUpdate) SetCloudProviderEndpointId(v string) {
	o.CloudProviderEndpointId = &v
}

// GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpointUpdate) GetPrivateEndpointIpAddress() string {
	if o == nil || IsNil(o.PrivateEndpointIpAddress) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointIpAddress
}

// GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpointUpdate) GetPrivateEndpointIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointIpAddress) {
		return nil, false
	}
	return o.PrivateEndpointIpAddress, true
}

// HasPrivateEndpointIpAddress returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpointUpdate) HasPrivateEndpointIpAddress() bool {
	if o != nil && !IsNil(o.PrivateEndpointIpAddress) {
		return true
	}

	return false
}

// SetPrivateEndpointIpAddress gets a reference to the given string and assigns it to the PrivateEndpointIpAddress field.
func (o *ServerlessAzureTenantEndpointUpdate) SetPrivateEndpointIpAddress(v string) {
	o.PrivateEndpointIpAddress = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpointUpdate) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpointUpdate) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpointUpdate) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ServerlessAzureTenantEndpointUpdate) SetComment(v string) {
	o.Comment = &v
}

// GetProviderName returns the ProviderName field value
func (o *ServerlessAzureTenantEndpointUpdate) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpointUpdate) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *ServerlessAzureTenantEndpointUpdate) SetProviderName(v string) {
	o.ProviderName = v
}

func (o ServerlessAzureTenantEndpointUpdate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessAzureTenantEndpointUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProviderEndpointId) {
		toSerialize["cloudProviderEndpointId"] = o.CloudProviderEndpointId
	}
	if !IsNil(o.PrivateEndpointIpAddress) {
		toSerialize["privateEndpointIpAddress"] = o.PrivateEndpointIpAddress
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["providerName"] = o.ProviderName
	return toSerialize, nil
}

type NullableServerlessAzureTenantEndpointUpdate struct {
	value *ServerlessAzureTenantEndpointUpdate
	isSet bool
}

func (v NullableServerlessAzureTenantEndpointUpdate) Get() *ServerlessAzureTenantEndpointUpdate {
	return v.value
}

func (v *NullableServerlessAzureTenantEndpointUpdate) Set(val *ServerlessAzureTenantEndpointUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessAzureTenantEndpointUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessAzureTenantEndpointUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessAzureTenantEndpointUpdate(val *ServerlessAzureTenantEndpointUpdate) *NullableServerlessAzureTenantEndpointUpdate {
	return &NullableServerlessAzureTenantEndpointUpdate{value: val, isSet: true}
}

func (v NullableServerlessAzureTenantEndpointUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessAzureTenantEndpointUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
