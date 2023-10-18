// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// ServiceAccountSecretRequest struct for ServiceAccountSecretRequest
type ServiceAccountSecretRequest struct {
	// Timestamp representing creation time for secret.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewServiceAccountSecretRequest instantiates a new ServiceAccountSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountSecretRequest() *ServiceAccountSecretRequest {
	this := ServiceAccountSecretRequest{}
	return &this
}

// NewServiceAccountSecretRequestWithDefaults instantiates a new ServiceAccountSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountSecretRequestWithDefaults() *ServiceAccountSecretRequest {
	this := ServiceAccountSecretRequest{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *ServiceAccountSecretRequest) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecretRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ServiceAccountSecretRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ServiceAccountSecretRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o ServiceAccountSecretRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountSecretRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}
