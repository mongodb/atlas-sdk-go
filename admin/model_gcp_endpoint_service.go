// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the GCPEndpointService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCPEndpointService{}

// GCPEndpointService Group of Private Endpoint Service settings.
type GCPEndpointService struct {
	// Cloud service provider that serves the requested endpoint service.
	CloudProvider string `json:"cloudProvider"`
	// List of Google Cloud network endpoint groups that corresponds to the Private Service Connect endpoint service.
	EndpointGroupNames []string `json:"endpointGroupNames,omitempty"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	Id *string `json:"id,omitempty"`
	// Cloud provider region that manages this Private Endpoint Service.
	RegionName *string `json:"regionName,omitempty"`
	// List of Uniform Resource Locators (URLs) that identifies endpoints that MongoDB Cloud can use to access one Google Cloud Service across a Google Cloud Virtual Private Connection (VPC) network.
	ServiceAttachmentNames []string `json:"serviceAttachmentNames,omitempty"`
	// State of the Private Endpoint Service connection when MongoDB Cloud received this request.
	Status *string `json:"status,omitempty"`
}

// NewGCPEndpointService instantiates a new GCPEndpointService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPEndpointService(cloudProvider string) *GCPEndpointService {
	this := GCPEndpointService{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewGCPEndpointServiceWithDefaults instantiates a new GCPEndpointService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPEndpointServiceWithDefaults() *GCPEndpointService {
	this := GCPEndpointService{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *GCPEndpointService) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *GCPEndpointService) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *GCPEndpointService) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetEndpointGroupNames returns the EndpointGroupNames field value if set, zero value otherwise.
func (o *GCPEndpointService) GetEndpointGroupNames() []string {
	if o == nil || IsNil(o.EndpointGroupNames) {
		var ret []string
		return ret
	}
	return o.EndpointGroupNames
}

// GetEndpointGroupNamesOk returns a tuple with the EndpointGroupNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPEndpointService) GetEndpointGroupNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.EndpointGroupNames) {
		return nil, false
	}
	return o.EndpointGroupNames, true
}

// HasEndpointGroupNames returns a boolean if a field has been set.
func (o *GCPEndpointService) HasEndpointGroupNames() bool {
	if o != nil && !IsNil(o.EndpointGroupNames) {
		return true
	}

	return false
}

// SetEndpointGroupNames gets a reference to the given []string and assigns it to the EndpointGroupNames field.
func (o *GCPEndpointService) SetEndpointGroupNames(v []string) {
	o.EndpointGroupNames = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GCPEndpointService) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPEndpointService) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GCPEndpointService) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GCPEndpointService) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GCPEndpointService) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPEndpointService) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GCPEndpointService) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GCPEndpointService) SetId(v string) {
	o.Id = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *GCPEndpointService) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPEndpointService) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *GCPEndpointService) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *GCPEndpointService) SetRegionName(v string) {
	o.RegionName = &v
}

// GetServiceAttachmentNames returns the ServiceAttachmentNames field value if set, zero value otherwise.
func (o *GCPEndpointService) GetServiceAttachmentNames() []string {
	if o == nil || IsNil(o.ServiceAttachmentNames) {
		var ret []string
		return ret
	}
	return o.ServiceAttachmentNames
}

// GetServiceAttachmentNamesOk returns a tuple with the ServiceAttachmentNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPEndpointService) GetServiceAttachmentNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ServiceAttachmentNames) {
		return nil, false
	}
	return o.ServiceAttachmentNames, true
}

// HasServiceAttachmentNames returns a boolean if a field has been set.
func (o *GCPEndpointService) HasServiceAttachmentNames() bool {
	if o != nil && !IsNil(o.ServiceAttachmentNames) {
		return true
	}

	return false
}

// SetServiceAttachmentNames gets a reference to the given []string and assigns it to the ServiceAttachmentNames field.
func (o *GCPEndpointService) SetServiceAttachmentNames(v []string) {
	o.ServiceAttachmentNames = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GCPEndpointService) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPEndpointService) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GCPEndpointService) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GCPEndpointService) SetStatus(v string) {
	o.Status = &v
}

func (o GCPEndpointService) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GCPEndpointService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointGroupNames) {
		toSerialize["endpointGroupNames"] = o.EndpointGroupNames
	}
	if !IsNil(o.ServiceAttachmentNames) {
		toSerialize["serviceAttachmentNames"] = o.ServiceAttachmentNames
	}
	return toSerialize, nil
}

type NullableGCPEndpointService struct {
	value *GCPEndpointService
	isSet bool
}

func (v NullableGCPEndpointService) Get() *GCPEndpointService {
	return v.value
}

func (v *NullableGCPEndpointService) Set(val *GCPEndpointService) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPEndpointService) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPEndpointService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPEndpointService(val *GCPEndpointService) *NullableGCPEndpointService {
	return &NullableGCPEndpointService{value: val, isSet: true}
}

func (v NullableGCPEndpointService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPEndpointService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
