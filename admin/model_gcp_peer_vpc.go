// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the GCPPeerVpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCPPeerVpc{}

// GCPPeerVpc Group of Network Peering connection settings.
type GCPPeerVpc struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId string `json:"containerId"`
	// Details of the error returned when requesting a GCP network peering resource. The resource returns `null` if the request succeeded.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Human-readable label that identifies the GCP project that contains the network that you want to peer with the MongoDB Cloud VPC.
	GcpProjectId string `json:"gcpProjectId"`
	// Unique 24-hexadecimal digit string that identifies the network peering connection.
	Id *string `json:"id,omitempty"`
	// Human-readable label that identifies the network to peer with the MongoDB Cloud VPC.
	NetworkName string `json:"networkName"`
	// Cloud service provider that serves the requested network peering connection.
	ProviderName *string `json:"providerName,omitempty"`
	// State of the network peering connection at the time you made the request.
	Status *string `json:"status,omitempty"`
}

// NewGCPPeerVpc instantiates a new GCPPeerVpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPPeerVpc(containerId string, gcpProjectId string, networkName string) *GCPPeerVpc {
	this := GCPPeerVpc{}
	this.ContainerId = containerId
	this.GcpProjectId = gcpProjectId
	this.NetworkName = networkName
	return &this
}

// NewGCPPeerVpcWithDefaults instantiates a new GCPPeerVpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPPeerVpcWithDefaults() *GCPPeerVpc {
	this := GCPPeerVpc{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *GCPPeerVpc) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *GCPPeerVpc) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *GCPPeerVpc) SetContainerId(v string) {
	o.ContainerId = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GCPPeerVpc) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPPeerVpc) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GCPPeerVpc) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GCPPeerVpc) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetGcpProjectId returns the GcpProjectId field value
func (o *GCPPeerVpc) GetGcpProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value
// and a boolean to check if the value has been set.
func (o *GCPPeerVpc) GetGcpProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcpProjectId, true
}

// SetGcpProjectId sets field value
func (o *GCPPeerVpc) SetGcpProjectId(v string) {
	o.GcpProjectId = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GCPPeerVpc) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPPeerVpc) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GCPPeerVpc) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GCPPeerVpc) SetId(v string) {
	o.Id = &v
}

// GetNetworkName returns the NetworkName field value
func (o *GCPPeerVpc) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *GCPPeerVpc) GetNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *GCPPeerVpc) SetNetworkName(v string) {
	o.NetworkName = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *GCPPeerVpc) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPPeerVpc) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *GCPPeerVpc) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *GCPPeerVpc) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GCPPeerVpc) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPPeerVpc) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GCPPeerVpc) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GCPPeerVpc) SetStatus(v string) {
	o.Status = &v
}

func (o GCPPeerVpc) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GCPPeerVpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerId"] = o.ContainerId
	toSerialize["gcpProjectId"] = o.GcpProjectId
	toSerialize["networkName"] = o.NetworkName
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	return toSerialize, nil
}

type NullableGCPPeerVpc struct {
	value *GCPPeerVpc
	isSet bool
}

func (v NullableGCPPeerVpc) Get() *GCPPeerVpc {
	return v.value
}

func (v *NullableGCPPeerVpc) Set(val *GCPPeerVpc) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPPeerVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPPeerVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPPeerVpc(val *GCPPeerVpc) *NullableGCPPeerVpc {
	return &NullableGCPPeerVpc{value: val, isSet: true}
}

func (v NullableGCPPeerVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPPeerVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
