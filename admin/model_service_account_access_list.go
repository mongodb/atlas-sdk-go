// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// ServiceAccountAccessList struct for ServiceAccountAccessList
type ServiceAccountAccessList struct {
	// Time of creation.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Last address where the service account was used from.
	LastUsedAddress *string `json:"lastUsedAddress,omitempty"`
	// Last time the access list was used.
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	// An IP address or CIDR notation.
	NetworkAddress *string `json:"networkAddress,omitempty"`
	// Count of requests using this network address.
	RequestCount *int64 `json:"requestCount,omitempty"`
}

// NewServiceAccountAccessList instantiates a new ServiceAccountAccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountAccessList() *ServiceAccountAccessList {
	this := ServiceAccountAccessList{}
	return &this
}

// NewServiceAccountAccessListWithDefaults instantiates a new ServiceAccountAccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountAccessListWithDefaults() *ServiceAccountAccessList {
	this := ServiceAccountAccessList{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *ServiceAccountAccessList) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessList) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountAccessList) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountAccessList) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUsedAddress returns the LastUsedAddress field value if set, zero value otherwise
func (o *ServiceAccountAccessList) GetLastUsedAddress() string {
	if o == nil || IsNil(o.LastUsedAddress) {
		var ret string
		return ret
	}
	return *o.LastUsedAddress
}

// GetLastUsedAddressOk returns a tuple with the LastUsedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessList) GetLastUsedAddressOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsedAddress) {
		return nil, false
	}

	return o.LastUsedAddress, true
}

// HasLastUsedAddress returns a boolean if a field has been set.
func (o *ServiceAccountAccessList) HasLastUsedAddress() bool {
	if o != nil && !IsNil(o.LastUsedAddress) {
		return true
	}

	return false
}

// SetLastUsedAddress gets a reference to the given string and assigns it to the LastUsedAddress field.
func (o *ServiceAccountAccessList) SetLastUsedAddress(v string) {
	o.LastUsedAddress = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise
func (o *ServiceAccountAccessList) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessList) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}

	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ServiceAccountAccessList) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *ServiceAccountAccessList) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetNetworkAddress returns the NetworkAddress field value if set, zero value otherwise
func (o *ServiceAccountAccessList) GetNetworkAddress() string {
	if o == nil || IsNil(o.NetworkAddress) {
		var ret string
		return ret
	}
	return *o.NetworkAddress
}

// GetNetworkAddressOk returns a tuple with the NetworkAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessList) GetNetworkAddressOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkAddress) {
		return nil, false
	}

	return o.NetworkAddress, true
}

// HasNetworkAddress returns a boolean if a field has been set.
func (o *ServiceAccountAccessList) HasNetworkAddress() bool {
	if o != nil && !IsNil(o.NetworkAddress) {
		return true
	}

	return false
}

// SetNetworkAddress gets a reference to the given string and assigns it to the NetworkAddress field.
func (o *ServiceAccountAccessList) SetNetworkAddress(v string) {
	o.NetworkAddress = &v
}

// GetRequestCount returns the RequestCount field value if set, zero value otherwise
func (o *ServiceAccountAccessList) GetRequestCount() int64 {
	if o == nil || IsNil(o.RequestCount) {
		var ret int64
		return ret
	}
	return *o.RequestCount
}

// GetRequestCountOk returns a tuple with the RequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessList) GetRequestCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RequestCount) {
		return nil, false
	}

	return o.RequestCount, true
}

// HasRequestCount returns a boolean if a field has been set.
func (o *ServiceAccountAccessList) HasRequestCount() bool {
	if o != nil && !IsNil(o.RequestCount) {
		return true
	}

	return false
}

// SetRequestCount gets a reference to the given int64 and assigns it to the RequestCount field.
func (o *ServiceAccountAccessList) SetRequestCount(v int64) {
	o.RequestCount = &v
}

func (o ServiceAccountAccessList) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountAccessList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkAddress) {
		toSerialize["networkAddress"] = o.NetworkAddress
	}
	return toSerialize, nil
}
