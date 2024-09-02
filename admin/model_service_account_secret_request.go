// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ServiceAccountSecretRequest struct for ServiceAccountSecretRequest
type ServiceAccountSecretRequest struct {
	// The expiration time of the new Service Account secret. The expiration is provided in hours.
	SecretExpiresAfterHours int `json:"secretExpiresAfterHours"`
}

// NewServiceAccountSecretRequest instantiates a new ServiceAccountSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountSecretRequest(secretExpiresAfterHours int) *ServiceAccountSecretRequest {
	this := ServiceAccountSecretRequest{}
	this.SecretExpiresAfterHours = secretExpiresAfterHours
	return &this
}

// NewServiceAccountSecretRequestWithDefaults instantiates a new ServiceAccountSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountSecretRequestWithDefaults() *ServiceAccountSecretRequest {
	this := ServiceAccountSecretRequest{}
	return &this
}

// GetSecretExpiresAfterHours returns the SecretExpiresAfterHours field value
func (o *ServiceAccountSecretRequest) GetSecretExpiresAfterHours() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.SecretExpiresAfterHours
}

// GetSecretExpiresAfterHoursOk returns a tuple with the SecretExpiresAfterHours field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecretRequest) GetSecretExpiresAfterHoursOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretExpiresAfterHours, true
}

// SetSecretExpiresAfterHours sets field value
func (o *ServiceAccountSecretRequest) SetSecretExpiresAfterHours(v int) {
	o.SecretExpiresAfterHours = v
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
	toSerialize["secretExpiresAfterHours"] = o.SecretExpiresAfterHours
	return toSerialize, nil
}
