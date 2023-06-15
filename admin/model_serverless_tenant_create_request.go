// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the ServerlessTenantCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessTenantCreateRequest{}

// ServerlessTenantCreateRequest struct for ServerlessTenantCreateRequest
type ServerlessTenantCreateRequest struct {
	// Human-readable comment associated with the private endpoint.
	Comment *string `json:"comment,omitempty"`
}

// NewServerlessTenantCreateRequest instantiates a new ServerlessTenantCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessTenantCreateRequest() *ServerlessTenantCreateRequest {
	this := ServerlessTenantCreateRequest{}
	return &this
}

// NewServerlessTenantCreateRequestWithDefaults instantiates a new ServerlessTenantCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessTenantCreateRequestWithDefaults() *ServerlessTenantCreateRequest {
	this := ServerlessTenantCreateRequest{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ServerlessTenantCreateRequest) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessTenantCreateRequest) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ServerlessTenantCreateRequest) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ServerlessTenantCreateRequest) SetComment(v string) {
	o.Comment = &v
}

func (o ServerlessTenantCreateRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServerlessTenantCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableServerlessTenantCreateRequest struct {
	value *ServerlessTenantCreateRequest
	isSet bool
}

func (v NullableServerlessTenantCreateRequest) Get() *ServerlessTenantCreateRequest {
	return v.value
}

func (v *NullableServerlessTenantCreateRequest) Set(val *ServerlessTenantCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessTenantCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessTenantCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessTenantCreateRequest(val *ServerlessTenantCreateRequest) *NullableServerlessTenantCreateRequest {
	return &NullableServerlessTenantCreateRequest{value: val, isSet: true}
}

func (v NullableServerlessTenantCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessTenantCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
