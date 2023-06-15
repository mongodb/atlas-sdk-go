// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PrivateGCPEndpointGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateGCPEndpointGroup{}

// PrivateGCPEndpointGroup Group of Private Endpoint settings.
type PrivateGCPEndpointGroup struct {
	// Cloud service provider that serves the requested endpoint.
	CloudProvider string `json:"cloudProvider"`
	// Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service.
	DeleteRequested *bool `json:"deleteRequested,omitempty"`
	// Human-readable label that identifies a set of endpoints.
	EndpointGroupName *string `json:"endpointGroupName,omitempty"`
	// List of individual private endpoints that comprise this endpoint group.
	Endpoints []GCPConsumerForwardingRule `json:"endpoints,omitempty"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// State of the Google Cloud network endpoint group when MongoDB Cloud received this request.
	Status *string `json:"status,omitempty"`
}

// NewPrivateGCPEndpointGroup instantiates a new PrivateGCPEndpointGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateGCPEndpointGroup(cloudProvider string) *PrivateGCPEndpointGroup {
	this := PrivateGCPEndpointGroup{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewPrivateGCPEndpointGroupWithDefaults instantiates a new PrivateGCPEndpointGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateGCPEndpointGroupWithDefaults() *PrivateGCPEndpointGroup {
	this := PrivateGCPEndpointGroup{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *PrivateGCPEndpointGroup) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *PrivateGCPEndpointGroup) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *PrivateGCPEndpointGroup) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetDeleteRequested returns the DeleteRequested field value if set, zero value otherwise.
func (o *PrivateGCPEndpointGroup) GetDeleteRequested() bool {
	if o == nil || IsNil(o.DeleteRequested) {
		var ret bool
		return ret
	}
	return *o.DeleteRequested
}

// GetDeleteRequestedOk returns a tuple with the DeleteRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGCPEndpointGroup) GetDeleteRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteRequested) {
		return nil, false
	}
	return o.DeleteRequested, true
}

// HasDeleteRequested returns a boolean if a field has been set.
func (o *PrivateGCPEndpointGroup) HasDeleteRequested() bool {
	if o != nil && !IsNil(o.DeleteRequested) {
		return true
	}

	return false
}

// SetDeleteRequested gets a reference to the given bool and assigns it to the DeleteRequested field.
func (o *PrivateGCPEndpointGroup) SetDeleteRequested(v bool) {
	o.DeleteRequested = &v
}

// GetEndpointGroupName returns the EndpointGroupName field value if set, zero value otherwise.
func (o *PrivateGCPEndpointGroup) GetEndpointGroupName() string {
	if o == nil || IsNil(o.EndpointGroupName) {
		var ret string
		return ret
	}
	return *o.EndpointGroupName
}

// GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGCPEndpointGroup) GetEndpointGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointGroupName) {
		return nil, false
	}
	return o.EndpointGroupName, true
}

// HasEndpointGroupName returns a boolean if a field has been set.
func (o *PrivateGCPEndpointGroup) HasEndpointGroupName() bool {
	if o != nil && !IsNil(o.EndpointGroupName) {
		return true
	}

	return false
}

// SetEndpointGroupName gets a reference to the given string and assigns it to the EndpointGroupName field.
func (o *PrivateGCPEndpointGroup) SetEndpointGroupName(v string) {
	o.EndpointGroupName = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *PrivateGCPEndpointGroup) GetEndpoints() []GCPConsumerForwardingRule {
	if o == nil || IsNil(o.Endpoints) {
		var ret []GCPConsumerForwardingRule
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGCPEndpointGroup) GetEndpointsOk() ([]GCPConsumerForwardingRule, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *PrivateGCPEndpointGroup) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []GCPConsumerForwardingRule and assigns it to the Endpoints field.
func (o *PrivateGCPEndpointGroup) SetEndpoints(v []GCPConsumerForwardingRule) {
	o.Endpoints = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PrivateGCPEndpointGroup) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGCPEndpointGroup) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PrivateGCPEndpointGroup) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PrivateGCPEndpointGroup) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PrivateGCPEndpointGroup) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateGCPEndpointGroup) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PrivateGCPEndpointGroup) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PrivateGCPEndpointGroup) SetStatus(v string) {
	o.Status = &v
}

func (o PrivateGCPEndpointGroup) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PrivateGCPEndpointGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullablePrivateGCPEndpointGroup struct {
	value *PrivateGCPEndpointGroup
	isSet bool
}

func (v NullablePrivateGCPEndpointGroup) Get() *PrivateGCPEndpointGroup {
	return v.value
}

func (v *NullablePrivateGCPEndpointGroup) Set(val *PrivateGCPEndpointGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateGCPEndpointGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateGCPEndpointGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateGCPEndpointGroup(val *PrivateGCPEndpointGroup) *NullablePrivateGCPEndpointGroup {
	return &NullablePrivateGCPEndpointGroup{value: val, isSet: true}
}

func (v NullablePrivateGCPEndpointGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateGCPEndpointGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
