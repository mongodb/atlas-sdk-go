// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the GeoSharding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoSharding{}

// GeoSharding struct for GeoSharding
type GeoSharding struct {
	// List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.  This parameter returns an empty object if no custom zones exist.
	CustomZoneMapping *map[string]string `json:"customZoneMapping,omitempty"`
	// List that contains a namespace for a Global Cluster. MongoDB Cloud manages this cluster.
	ManagedNamespaces []ManagedNamespaces `json:"managedNamespaces,omitempty"`
}

// NewGeoSharding instantiates a new GeoSharding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoSharding() *GeoSharding {
	this := GeoSharding{}
	return &this
}

// NewGeoShardingWithDefaults instantiates a new GeoSharding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoShardingWithDefaults() *GeoSharding {
	this := GeoSharding{}
	return &this
}

// GetCustomZoneMapping returns the CustomZoneMapping field value if set, zero value otherwise.
func (o *GeoSharding) GetCustomZoneMapping() map[string]string {
	if o == nil || IsNil(o.CustomZoneMapping) {
		var ret map[string]string
		return ret
	}
	return *o.CustomZoneMapping
}

// GetCustomZoneMappingOk returns a tuple with the CustomZoneMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoSharding) GetCustomZoneMappingOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomZoneMapping) {
		return nil, false
	}
	return o.CustomZoneMapping, true
}

// HasCustomZoneMapping returns a boolean if a field has been set.
func (o *GeoSharding) HasCustomZoneMapping() bool {
	if o != nil && !IsNil(o.CustomZoneMapping) {
		return true
	}

	return false
}

// SetCustomZoneMapping gets a reference to the given map[string]string and assigns it to the CustomZoneMapping field.
func (o *GeoSharding) SetCustomZoneMapping(v map[string]string) {
	o.CustomZoneMapping = &v
}

// GetManagedNamespaces returns the ManagedNamespaces field value if set, zero value otherwise.
func (o *GeoSharding) GetManagedNamespaces() []ManagedNamespaces {
	if o == nil || IsNil(o.ManagedNamespaces) {
		var ret []ManagedNamespaces
		return ret
	}
	return o.ManagedNamespaces
}

// GetManagedNamespacesOk returns a tuple with the ManagedNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoSharding) GetManagedNamespacesOk() ([]ManagedNamespaces, bool) {
	if o == nil || IsNil(o.ManagedNamespaces) {
		return nil, false
	}
	return o.ManagedNamespaces, true
}

// HasManagedNamespaces returns a boolean if a field has been set.
func (o *GeoSharding) HasManagedNamespaces() bool {
	if o != nil && !IsNil(o.ManagedNamespaces) {
		return true
	}

	return false
}

// SetManagedNamespaces gets a reference to the given []ManagedNamespaces and assigns it to the ManagedNamespaces field.
func (o *GeoSharding) SetManagedNamespaces(v []ManagedNamespaces) {
	o.ManagedNamespaces = v
}

func (o GeoSharding) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GeoSharding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableGeoSharding struct {
	value *GeoSharding
	isSet bool
}

func (v NullableGeoSharding) Get() *GeoSharding {
	return v.value
}

func (v *NullableGeoSharding) Set(val *GeoSharding) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoSharding) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoSharding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoSharding(val *GeoSharding) *NullableGeoSharding {
	return &NullableGeoSharding{value: val, isSet: true}
}

func (v NullableGeoSharding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoSharding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}