/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the AWSPrivateLinkConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSPrivateLinkConnection{}

// AWSPrivateLinkConnection Group of Private Endpoint Service settings.
type AWSPrivateLinkConnection struct {
	// Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service.
	EndpointServiceName *string `json:"endpointServiceName,omitempty"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	Id *string `json:"id,omitempty"`
	// List of strings that identify private endpoint interfaces applied to the specified project.
	InterfaceEndpoints []string `json:"interfaceEndpoints,omitempty"`
	// Cloud provider region that manages this Private Endpoint Service.
	RegionName *string `json:"regionName,omitempty"`
	// State of the Private Endpoint Service connection when MongoDB Cloud received this request.
	Status *string `json:"status,omitempty"`
}

// NewAWSPrivateLinkConnection instantiates a new AWSPrivateLinkConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSPrivateLinkConnection() *AWSPrivateLinkConnection {
	this := AWSPrivateLinkConnection{}
	return &this
}

// NewAWSPrivateLinkConnectionWithDefaults instantiates a new AWSPrivateLinkConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSPrivateLinkConnectionWithDefaults() *AWSPrivateLinkConnection {
	this := AWSPrivateLinkConnection{}
	return &this
}

// GetEndpointServiceName returns the EndpointServiceName field value if set, zero value otherwise.
func (o *AWSPrivateLinkConnection) GetEndpointServiceName() string {
	if o == nil || IsNil(o.EndpointServiceName) {
		var ret string
		return ret
	}
	return *o.EndpointServiceName
}

// GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConnection) GetEndpointServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointServiceName) {
		return nil, false
	}
	return o.EndpointServiceName, true
}

// HasEndpointServiceName returns a boolean if a field has been set.
func (o *AWSPrivateLinkConnection) HasEndpointServiceName() bool {
	if o != nil && !IsNil(o.EndpointServiceName) {
		return true
	}

	return false
}

// SetEndpointServiceName gets a reference to the given string and assigns it to the EndpointServiceName field.
func (o *AWSPrivateLinkConnection) SetEndpointServiceName(v string) {
	o.EndpointServiceName = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AWSPrivateLinkConnection) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConnection) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AWSPrivateLinkConnection) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AWSPrivateLinkConnection) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AWSPrivateLinkConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AWSPrivateLinkConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AWSPrivateLinkConnection) SetId(v string) {
	o.Id = &v
}

// GetInterfaceEndpoints returns the InterfaceEndpoints field value if set, zero value otherwise.
func (o *AWSPrivateLinkConnection) GetInterfaceEndpoints() []string {
	if o == nil || IsNil(o.InterfaceEndpoints) {
		var ret []string
		return ret
	}
	return o.InterfaceEndpoints
}

// GetInterfaceEndpointsOk returns a tuple with the InterfaceEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConnection) GetInterfaceEndpointsOk() ([]string, bool) {
	if o == nil || IsNil(o.InterfaceEndpoints) {
		return nil, false
	}
	return o.InterfaceEndpoints, true
}

// HasInterfaceEndpoints returns a boolean if a field has been set.
func (o *AWSPrivateLinkConnection) HasInterfaceEndpoints() bool {
	if o != nil && !IsNil(o.InterfaceEndpoints) {
		return true
	}

	return false
}

// SetInterfaceEndpoints gets a reference to the given []string and assigns it to the InterfaceEndpoints field.
func (o *AWSPrivateLinkConnection) SetInterfaceEndpoints(v []string) {
	o.InterfaceEndpoints = v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AWSPrivateLinkConnection) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConnection) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AWSPrivateLinkConnection) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AWSPrivateLinkConnection) SetRegionName(v string) {
	o.RegionName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AWSPrivateLinkConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPrivateLinkConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AWSPrivateLinkConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AWSPrivateLinkConnection) SetStatus(v string) {
	o.Status = &v
}

func (o AWSPrivateLinkConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSPrivateLinkConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: endpointServiceName is readOnly
	// skip: errorMessage is readOnly
	// skip: id is readOnly
	// skip: interfaceEndpoints is readOnly
	// skip: regionName is readOnly
	// skip: status is readOnly
	return toSerialize, nil
}

type NullableAWSPrivateLinkConnection struct {
	value *AWSPrivateLinkConnection
	isSet bool
}

func (v NullableAWSPrivateLinkConnection) Get() *AWSPrivateLinkConnection {
	return v.value
}

func (v *NullableAWSPrivateLinkConnection) Set(val *AWSPrivateLinkConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSPrivateLinkConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSPrivateLinkConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSPrivateLinkConnection(val *AWSPrivateLinkConnection) *NullableAWSPrivateLinkConnection {
	return &NullableAWSPrivateLinkConnection{value: val, isSet: true}
}

func (v NullableAWSPrivateLinkConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSPrivateLinkConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


