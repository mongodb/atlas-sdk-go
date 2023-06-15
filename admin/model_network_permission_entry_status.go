// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the NetworkPermissionEntryStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkPermissionEntryStatus{}

// NetworkPermissionEntryStatus struct for NetworkPermissionEntryStatus
type NetworkPermissionEntryStatus struct {
	// State of the access list entry when MongoDB Cloud made this request.  | Status | Activity | |---|---| | `ACTIVE` | This access list entry applies to all relevant cloud providers. | | `PENDING` | MongoDB Cloud has started to add access list entry. This access list entry may not apply to all cloud providers at the time of this request. | | `FAILED` | MongoDB Cloud didn't succeed in adding this access list entry. |
	STATUS string `json:"STATUS"`
}

// NewNetworkPermissionEntryStatus instantiates a new NetworkPermissionEntryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPermissionEntryStatus(sTATUS string) *NetworkPermissionEntryStatus {
	this := NetworkPermissionEntryStatus{}
	this.STATUS = sTATUS
	return &this
}

// NewNetworkPermissionEntryStatusWithDefaults instantiates a new NetworkPermissionEntryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPermissionEntryStatusWithDefaults() *NetworkPermissionEntryStatus {
	this := NetworkPermissionEntryStatus{}
	return &this
}

// GetSTATUS returns the STATUS field value
func (o *NetworkPermissionEntryStatus) GetSTATUS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.STATUS
}

// GetSTATUSOk returns a tuple with the STATUS field value
// and a boolean to check if the value has been set.
func (o *NetworkPermissionEntryStatus) GetSTATUSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.STATUS, true
}

// SetSTATUS sets field value
func (o *NetworkPermissionEntryStatus) SetSTATUS(v string) {
	o.STATUS = v
}

func (o NetworkPermissionEntryStatus) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o NetworkPermissionEntryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableNetworkPermissionEntryStatus struct {
	value *NetworkPermissionEntryStatus
	isSet bool
}

func (v NullableNetworkPermissionEntryStatus) Get() *NetworkPermissionEntryStatus {
	return v.value
}

func (v *NullableNetworkPermissionEntryStatus) Set(val *NetworkPermissionEntryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPermissionEntryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPermissionEntryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPermissionEntryStatus(val *NetworkPermissionEntryStatus) *NullableNetworkPermissionEntryStatus {
	return &NullableNetworkPermissionEntryStatus{value: val, isSet: true}
}

func (v NullableNetworkPermissionEntryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPermissionEntryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
