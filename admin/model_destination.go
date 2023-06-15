// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destination{}

// Destination Document that describes the destination of the migration.
type Destination struct {
	// Label that identifies the destination cluster.
	ClusterName string `json:"clusterName"`
	// Unique 24-hexadecimal digit string that identifies the destination project.
	GroupId string `json:"groupId"`
	// The network type to use between the migration host and the target cluster.
	HostnameSchemaType string `json:"hostnameSchemaType"`
	// Represents the endpoint to use when the host schema type is `PRIVATE_LINK`.
	PrivateLinkId *string `json:"privateLinkId,omitempty"`
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestination(clusterName string, groupId string, hostnameSchemaType string) *Destination {
	this := Destination{}
	this.ClusterName = clusterName
	this.GroupId = groupId
	this.HostnameSchemaType = hostnameSchemaType
	return &this
}

// NewDestinationWithDefaults instantiates a new Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationWithDefaults() *Destination {
	this := Destination{}
	var hostnameSchemaType string = "PUBLIC"
	this.HostnameSchemaType = hostnameSchemaType
	return &this
}

// GetClusterName returns the ClusterName field value
func (o *Destination) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *Destination) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *Destination) SetClusterName(v string) {
	o.ClusterName = v
}

// GetGroupId returns the GroupId field value
func (o *Destination) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *Destination) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *Destination) SetGroupId(v string) {
	o.GroupId = v
}

// GetHostnameSchemaType returns the HostnameSchemaType field value
func (o *Destination) GetHostnameSchemaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostnameSchemaType
}

// GetHostnameSchemaTypeOk returns a tuple with the HostnameSchemaType field value
// and a boolean to check if the value has been set.
func (o *Destination) GetHostnameSchemaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostnameSchemaType, true
}

// SetHostnameSchemaType sets field value
func (o *Destination) SetHostnameSchemaType(v string) {
	o.HostnameSchemaType = v
}

// GetPrivateLinkId returns the PrivateLinkId field value if set, zero value otherwise.
func (o *Destination) GetPrivateLinkId() string {
	if o == nil || IsNil(o.PrivateLinkId) {
		var ret string
		return ret
	}
	return *o.PrivateLinkId
}

// GetPrivateLinkIdOk returns a tuple with the PrivateLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetPrivateLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateLinkId) {
		return nil, false
	}
	return o.PrivateLinkId, true
}

// HasPrivateLinkId returns a boolean if a field has been set.
func (o *Destination) HasPrivateLinkId() bool {
	if o != nil && !IsNil(o.PrivateLinkId) {
		return true
	}

	return false
}

// SetPrivateLinkId gets a reference to the given string and assigns it to the PrivateLinkId field.
func (o *Destination) SetPrivateLinkId(v string) {
	o.PrivateLinkId = &v
}

func (o Destination) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["groupId"] = o.GroupId
	toSerialize["hostnameSchemaType"] = o.HostnameSchemaType
	if !IsNil(o.PrivateLinkId) {
		toSerialize["privateLinkId"] = o.PrivateLinkId
	}
	return toSerialize, nil
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
