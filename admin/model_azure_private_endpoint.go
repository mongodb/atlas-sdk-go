// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AzurePrivateEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzurePrivateEndpoint{}

// AzurePrivateEndpoint Group of Private Endpoint settings.
type AzurePrivateEndpoint struct {
	// Cloud service provider that serves the requested endpoint.
	CloudProvider string `json:"cloudProvider"`
	// Flag that indicates whether MongoDB Cloud received a request to remove the specified private endpoint from the private endpoint service.
	DeleteRequested *bool `json:"deleteRequested,omitempty"`
	// Error message returned when requesting private connection resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Human-readable label that MongoDB Cloud generates that identifies the private endpoint connection.
	PrivateEndpointConnectionName *string `json:"privateEndpointConnectionName,omitempty"`
	// IPv4 address of the private endpoint in your Azure VNet that someone added to this private endpoint service.
	PrivateEndpointIPAddress *string `json:"privateEndpointIPAddress,omitempty"`
	// Unique string that identifies the Azure private endpoint's network interface that someone added to this private endpoint service.
	PrivateEndpointResourceId *string `json:"privateEndpointResourceId,omitempty"`
	// State of the Azure Private Link Service connection when MongoDB Cloud received this request.
	Status *string `json:"status,omitempty"`
}

// NewAzurePrivateEndpoint instantiates a new AzurePrivateEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurePrivateEndpoint(cloudProvider string) *AzurePrivateEndpoint {
	this := AzurePrivateEndpoint{}
	this.CloudProvider = cloudProvider
	return &this
}

// NewAzurePrivateEndpointWithDefaults instantiates a new AzurePrivateEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurePrivateEndpointWithDefaults() *AzurePrivateEndpoint {
	this := AzurePrivateEndpoint{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *AzurePrivateEndpoint) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *AzurePrivateEndpoint) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *AzurePrivateEndpoint) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetDeleteRequested returns the DeleteRequested field value if set, zero value otherwise.
func (o *AzurePrivateEndpoint) GetDeleteRequested() bool {
	if o == nil || IsNil(o.DeleteRequested) {
		var ret bool
		return ret
	}
	return *o.DeleteRequested
}

// GetDeleteRequestedOk returns a tuple with the DeleteRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateEndpoint) GetDeleteRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteRequested) {
		return nil, false
	}
	return o.DeleteRequested, true
}

// HasDeleteRequested returns a boolean if a field has been set.
func (o *AzurePrivateEndpoint) HasDeleteRequested() bool {
	if o != nil && !IsNil(o.DeleteRequested) {
		return true
	}

	return false
}

// SetDeleteRequested gets a reference to the given bool and assigns it to the DeleteRequested field.
func (o *AzurePrivateEndpoint) SetDeleteRequested(v bool) {
	o.DeleteRequested = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AzurePrivateEndpoint) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateEndpoint) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AzurePrivateEndpoint) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AzurePrivateEndpoint) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field value if set, zero value otherwise.
func (o *AzurePrivateEndpoint) GetPrivateEndpointConnectionName() string {
	if o == nil || IsNil(o.PrivateEndpointConnectionName) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointConnectionName
}

// GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateEndpoint) GetPrivateEndpointConnectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointConnectionName) {
		return nil, false
	}
	return o.PrivateEndpointConnectionName, true
}

// HasPrivateEndpointConnectionName returns a boolean if a field has been set.
func (o *AzurePrivateEndpoint) HasPrivateEndpointConnectionName() bool {
	if o != nil && !IsNil(o.PrivateEndpointConnectionName) {
		return true
	}

	return false
}

// SetPrivateEndpointConnectionName gets a reference to the given string and assigns it to the PrivateEndpointConnectionName field.
func (o *AzurePrivateEndpoint) SetPrivateEndpointConnectionName(v string) {
	o.PrivateEndpointConnectionName = &v
}

// GetPrivateEndpointIPAddress returns the PrivateEndpointIPAddress field value if set, zero value otherwise.
func (o *AzurePrivateEndpoint) GetPrivateEndpointIPAddress() string {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointIPAddress
}

// GetPrivateEndpointIPAddressOk returns a tuple with the PrivateEndpointIPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateEndpoint) GetPrivateEndpointIPAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointIPAddress) {
		return nil, false
	}
	return o.PrivateEndpointIPAddress, true
}

// HasPrivateEndpointIPAddress returns a boolean if a field has been set.
func (o *AzurePrivateEndpoint) HasPrivateEndpointIPAddress() bool {
	if o != nil && !IsNil(o.PrivateEndpointIPAddress) {
		return true
	}

	return false
}

// SetPrivateEndpointIPAddress gets a reference to the given string and assigns it to the PrivateEndpointIPAddress field.
func (o *AzurePrivateEndpoint) SetPrivateEndpointIPAddress(v string) {
	o.PrivateEndpointIPAddress = &v
}

// GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field value if set, zero value otherwise.
func (o *AzurePrivateEndpoint) GetPrivateEndpointResourceId() string {
	if o == nil || IsNil(o.PrivateEndpointResourceId) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointResourceId
}

// GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateEndpoint) GetPrivateEndpointResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointResourceId) {
		return nil, false
	}
	return o.PrivateEndpointResourceId, true
}

// HasPrivateEndpointResourceId returns a boolean if a field has been set.
func (o *AzurePrivateEndpoint) HasPrivateEndpointResourceId() bool {
	if o != nil && !IsNil(o.PrivateEndpointResourceId) {
		return true
	}

	return false
}

// SetPrivateEndpointResourceId gets a reference to the given string and assigns it to the PrivateEndpointResourceId field.
func (o *AzurePrivateEndpoint) SetPrivateEndpointResourceId(v string) {
	o.PrivateEndpointResourceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AzurePrivateEndpoint) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrivateEndpoint) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AzurePrivateEndpoint) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AzurePrivateEndpoint) SetStatus(v string) {
	o.Status = &v
}

func (o AzurePrivateEndpoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AzurePrivateEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrivateEndpointIPAddress) {
		toSerialize["privateEndpointIPAddress"] = o.PrivateEndpointIPAddress
	}
	return toSerialize, nil
}

type NullableAzurePrivateEndpoint struct {
	value *AzurePrivateEndpoint
	isSet bool
}

func (v NullableAzurePrivateEndpoint) Get() *AzurePrivateEndpoint {
	return v.value
}

func (v *NullableAzurePrivateEndpoint) Set(val *AzurePrivateEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurePrivateEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurePrivateEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurePrivateEndpoint(val *AzurePrivateEndpoint) *NullableAzurePrivateEndpoint {
	return &NullableAzurePrivateEndpoint{value: val, isSet: true}
}

func (v NullableAzurePrivateEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurePrivateEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
