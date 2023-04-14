/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the ZoneMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneMapping{}

// ZoneMapping Human-readable label that identifies the subset of a global cluster.
type ZoneMapping struct {
	// Code that represents a location that maps to a zone in your global cluster. MongoDB Cloud represents this location with a ISO 3166-2 location and subdivision codes when possible.
	Location string `json:"location"`
	// Human-readable label that identifies the zone in your global cluster. This zone maps to a location code.
	Zone string `json:"zone"`
}

// NewZoneMapping instantiates a new ZoneMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneMapping(location string, zone string) *ZoneMapping {
	this := ZoneMapping{}
	this.Location = location
	this.Zone = zone
	return &this
}

// NewZoneMappingWithDefaults instantiates a new ZoneMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneMappingWithDefaults() *ZoneMapping {
	this := ZoneMapping{}
	return &this
}

// GetLocation returns the Location field value
func (o *ZoneMapping) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ZoneMapping) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ZoneMapping) SetLocation(v string) {
	o.Location = v
}

// GetZone returns the Zone field value
func (o *ZoneMapping) GetZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *ZoneMapping) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *ZoneMapping) SetZone(v string) {
	o.Zone = v
}

func (o ZoneMapping) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ZoneMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	toSerialize["zone"] = o.Zone
	return toSerialize, nil
}

type NullableZoneMapping struct {
	value *ZoneMapping
	isSet bool
}

func (v NullableZoneMapping) Get() *ZoneMapping {
	return v.value
}

func (v *NullableZoneMapping) Set(val *ZoneMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneMapping(val *ZoneMapping) *NullableZoneMapping {
	return &NullableZoneMapping{value: val, isSet: true}
}

func (v NullableZoneMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


