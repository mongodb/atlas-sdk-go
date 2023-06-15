// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AWSInterfaceEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSInterfaceEndpoint{}

// AWSInterfaceEndpoint Group of Private Endpoint settings.
type AWSInterfaceEndpoint struct {
	// Cloud service provider that serves the requested endpoint.
	CloudProvider string `json:"cloudProvider"`
	// State of the Amazon Web Service PrivateLink connection when MongoDB Cloud received this request.
	ConnectionStatus *string `json:"connectionStatus,omitempty"`
	// Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service.
	DeleteRequested *bool `json:"deleteRequested,omitempty"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the interface endpoint.
	InterfaceEndpointId *string `json:"interfaceEndpointId,omitempty"`
}

// NewAWSInterfaceEndpoint instantiates a new AWSInterfaceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSInterfaceEndpoint(cloudProvider string) *AWSInterfaceEndpoint {
	this := AWSInterfaceEndpoint{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewAWSInterfaceEndpointWithDefaults instantiates a new AWSInterfaceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSInterfaceEndpointWithDefaults() *AWSInterfaceEndpoint {
	this := AWSInterfaceEndpoint{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *AWSInterfaceEndpoint) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *AWSInterfaceEndpoint) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *AWSInterfaceEndpoint) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *AWSInterfaceEndpoint) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSInterfaceEndpoint) GetConnectionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *AWSInterfaceEndpoint) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *AWSInterfaceEndpoint) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDeleteRequested returns the DeleteRequested field value if set, zero value otherwise.
func (o *AWSInterfaceEndpoint) GetDeleteRequested() bool {
	if o == nil || IsNil(o.DeleteRequested) {
		var ret bool
		return ret
	}
	return *o.DeleteRequested
}

// GetDeleteRequestedOk returns a tuple with the DeleteRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSInterfaceEndpoint) GetDeleteRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteRequested) {
		return nil, false
	}
	return o.DeleteRequested, true
}

// HasDeleteRequested returns a boolean if a field has been set.
func (o *AWSInterfaceEndpoint) HasDeleteRequested() bool {
	if o != nil && !IsNil(o.DeleteRequested) {
		return true
	}

	return false
}

// SetDeleteRequested gets a reference to the given bool and assigns it to the DeleteRequested field.
func (o *AWSInterfaceEndpoint) SetDeleteRequested(v bool) {
	o.DeleteRequested = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AWSInterfaceEndpoint) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSInterfaceEndpoint) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AWSInterfaceEndpoint) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AWSInterfaceEndpoint) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetInterfaceEndpointId returns the InterfaceEndpointId field value if set, zero value otherwise.
func (o *AWSInterfaceEndpoint) GetInterfaceEndpointId() string {
	if o == nil || IsNil(o.InterfaceEndpointId) {
		var ret string
		return ret
	}
	return *o.InterfaceEndpointId
}

// GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSInterfaceEndpoint) GetInterfaceEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceEndpointId) {
		return nil, false
	}
	return o.InterfaceEndpointId, true
}

// HasInterfaceEndpointId returns a boolean if a field has been set.
func (o *AWSInterfaceEndpoint) HasInterfaceEndpointId() bool {
	if o != nil && !IsNil(o.InterfaceEndpointId) {
		return true
	}

	return false
}

// SetInterfaceEndpointId gets a reference to the given string and assigns it to the InterfaceEndpointId field.
func (o *AWSInterfaceEndpoint) SetInterfaceEndpointId(v string) {
	o.InterfaceEndpointId = &v
}

func (o AWSInterfaceEndpoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AWSInterfaceEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAWSInterfaceEndpoint struct {
	value *AWSInterfaceEndpoint
	isSet bool
}

func (v NullableAWSInterfaceEndpoint) Get() *AWSInterfaceEndpoint {
	return v.value
}

func (v *NullableAWSInterfaceEndpoint) Set(val *AWSInterfaceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSInterfaceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSInterfaceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSInterfaceEndpoint(val *AWSInterfaceEndpoint) *NullableAWSInterfaceEndpoint {
	return &NullableAWSInterfaceEndpoint{value: val, isSet: true}
}

func (v NullableAWSInterfaceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSInterfaceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
