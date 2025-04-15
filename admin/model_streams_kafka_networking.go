// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// StreamsKafkaNetworking Networking Access Type can either be 'PUBLIC' (default) or VPC. VPC type is in public preview, please file a support ticket to enable VPC Network Access.
type StreamsKafkaNetworking struct {
	Access *StreamsKafkaNetworkingAccess `json:"access,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
}

// NewStreamsKafkaNetworking instantiates a new StreamsKafkaNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsKafkaNetworking() *StreamsKafkaNetworking {
	this := StreamsKafkaNetworking{}
	return &this
}

// NewStreamsKafkaNetworkingWithDefaults instantiates a new StreamsKafkaNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsKafkaNetworkingWithDefaults() *StreamsKafkaNetworking {
	this := StreamsKafkaNetworking{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise
func (o *StreamsKafkaNetworking) GetAccess() StreamsKafkaNetworkingAccess {
	if o == nil || IsNil(o.Access) {
		var ret StreamsKafkaNetworkingAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaNetworking) GetAccessOk() (*StreamsKafkaNetworkingAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}

	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *StreamsKafkaNetworking) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given StreamsKafkaNetworkingAccess and assigns it to the Access field.
func (o *StreamsKafkaNetworking) SetAccess(v StreamsKafkaNetworkingAccess) {
	o.Access = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsKafkaNetworking) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsKafkaNetworking) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsKafkaNetworking) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsKafkaNetworking) SetLinks(v []Link) {
	o.Links = &v
}

func (o StreamsKafkaNetworking) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o StreamsKafkaNetworking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}
