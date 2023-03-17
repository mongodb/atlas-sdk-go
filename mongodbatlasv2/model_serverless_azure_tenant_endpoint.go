/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ServerlessAzureTenantEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessAzureTenantEndpoint{}

// ServerlessAzureTenantEndpoint View for a serverless Azure tenant endpoint.
type ServerlessAzureTenantEndpoint struct {
	// Unique 24-hexadecimal digit string that identifies the private endpoint.
	Id *string `json:"_id,omitempty"`
	// Unique string that identifies the Azure private endpoint's network interface that someone added to this private endpoint service.
	CloudProviderEndpointId *string `json:"cloudProviderEndpointId,omitempty"`
	// Human-readable comment associated with the private endpoint.
	Comment *string `json:"comment,omitempty"`
	// Unique string that identifies the Azure private endpoint service. MongoDB Cloud returns null while it creates the endpoint service.
	EndpointServiceName *string `json:"endpointServiceName,omitempty"`
	// Human-readable error message that indicates error condition associated with establishing the private endpoint connection.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIpAddress *string `json:"privateEndpointIpAddress,omitempty"`
	// Root-relative path that identifies the Azure Private Link Service that MongoDB Cloud manages. MongoDB Cloud returns null while it creates the endpoint service.
	PrivateLinkServiceResourceId *string `json:"privateLinkServiceResourceId,omitempty"`
	// Human-readable label that identifies the cloud service provider.
	ProviderName *string `json:"providerName,omitempty"`
	// Human-readable label that indicates the current operating status of the private endpoint.
	Status *string `json:"status,omitempty"`
}

// NewServerlessAzureTenantEndpoint instantiates a new ServerlessAzureTenantEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessAzureTenantEndpoint() *ServerlessAzureTenantEndpoint {
	this := ServerlessAzureTenantEndpoint{}
	return &this
}

// NewServerlessAzureTenantEndpointWithDefaults instantiates a new ServerlessAzureTenantEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessAzureTenantEndpointWithDefaults() *ServerlessAzureTenantEndpoint {
	this := ServerlessAzureTenantEndpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerlessAzureTenantEndpoint) SetId(v string) {
	o.Id = &v
}

// GetCloudProviderEndpointId returns the CloudProviderEndpointId field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetCloudProviderEndpointId() string {
	if o == nil || IsNil(o.CloudProviderEndpointId) {
		var ret string
		return ret
	}
	return *o.CloudProviderEndpointId
}

// GetCloudProviderEndpointIdOk returns a tuple with the CloudProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetCloudProviderEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProviderEndpointId) {
		return nil, false
	}
	return o.CloudProviderEndpointId, true
}

// HasCloudProviderEndpointId returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasCloudProviderEndpointId() bool {
	if o != nil && !IsNil(o.CloudProviderEndpointId) {
		return true
	}

	return false
}

// SetCloudProviderEndpointId gets a reference to the given string and assigns it to the CloudProviderEndpointId field.
func (o *ServerlessAzureTenantEndpoint) SetCloudProviderEndpointId(v string) {
	o.CloudProviderEndpointId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ServerlessAzureTenantEndpoint) SetComment(v string) {
	o.Comment = &v
}

// GetEndpointServiceName returns the EndpointServiceName field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetEndpointServiceName() string {
	if o == nil || IsNil(o.EndpointServiceName) {
		var ret string
		return ret
	}
	return *o.EndpointServiceName
}

// GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetEndpointServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointServiceName) {
		return nil, false
	}
	return o.EndpointServiceName, true
}

// HasEndpointServiceName returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasEndpointServiceName() bool {
	if o != nil && !IsNil(o.EndpointServiceName) {
		return true
	}

	return false
}

// SetEndpointServiceName gets a reference to the given string and assigns it to the EndpointServiceName field.
func (o *ServerlessAzureTenantEndpoint) SetEndpointServiceName(v string) {
	o.EndpointServiceName = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ServerlessAzureTenantEndpoint) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetPrivateEndpointIpAddress returns the PrivateEndpointIpAddress field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetPrivateEndpointIpAddress() string {
	if o == nil || IsNil(o.PrivateEndpointIpAddress) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointIpAddress
}

// GetPrivateEndpointIpAddressOk returns a tuple with the PrivateEndpointIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetPrivateEndpointIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointIpAddress) {
		return nil, false
	}
	return o.PrivateEndpointIpAddress, true
}

// HasPrivateEndpointIpAddress returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasPrivateEndpointIpAddress() bool {
	if o != nil && !IsNil(o.PrivateEndpointIpAddress) {
		return true
	}

	return false
}

// SetPrivateEndpointIpAddress gets a reference to the given string and assigns it to the PrivateEndpointIpAddress field.
func (o *ServerlessAzureTenantEndpoint) SetPrivateEndpointIpAddress(v string) {
	o.PrivateEndpointIpAddress = &v
}

// GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetPrivateLinkServiceResourceId() string {
	if o == nil || IsNil(o.PrivateLinkServiceResourceId) {
		var ret string
		return ret
	}
	return *o.PrivateLinkServiceResourceId
}

// GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetPrivateLinkServiceResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateLinkServiceResourceId) {
		return nil, false
	}
	return o.PrivateLinkServiceResourceId, true
}

// HasPrivateLinkServiceResourceId returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasPrivateLinkServiceResourceId() bool {
	if o != nil && !IsNil(o.PrivateLinkServiceResourceId) {
		return true
	}

	return false
}

// SetPrivateLinkServiceResourceId gets a reference to the given string and assigns it to the PrivateLinkServiceResourceId field.
func (o *ServerlessAzureTenantEndpoint) SetPrivateLinkServiceResourceId(v string) {
	o.PrivateLinkServiceResourceId = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ServerlessAzureTenantEndpoint) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServerlessAzureTenantEndpoint) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessAzureTenantEndpoint) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServerlessAzureTenantEndpoint) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ServerlessAzureTenantEndpoint) SetStatus(v string) {
	o.Status = &v
}

func (o ServerlessAzureTenantEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerlessAzureTenantEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: _id is readOnly
	// skip: cloudProviderEndpointId is readOnly
	// skip: comment is readOnly
	// skip: endpointServiceName is readOnly
	// skip: errorMessage is readOnly
	// skip: privateEndpointIpAddress is readOnly
	// skip: privateLinkServiceResourceId is readOnly
	// skip: providerName is readOnly
	// skip: status is readOnly
	return toSerialize, nil
}

type NullableServerlessAzureTenantEndpoint struct {
	value *ServerlessAzureTenantEndpoint
	isSet bool
}

func (v NullableServerlessAzureTenantEndpoint) Get() *ServerlessAzureTenantEndpoint {
	return v.value
}

func (v *NullableServerlessAzureTenantEndpoint) Set(val *ServerlessAzureTenantEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessAzureTenantEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessAzureTenantEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessAzureTenantEndpoint(val *ServerlessAzureTenantEndpoint) *NullableServerlessAzureTenantEndpoint {
	return &NullableServerlessAzureTenantEndpoint{value: val, isSet: true}
}

func (v NullableServerlessAzureTenantEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessAzureTenantEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


