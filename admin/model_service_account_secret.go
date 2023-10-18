// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// ServiceAccountSecret struct for ServiceAccountSecret
type ServiceAccountSecret struct {
	// Timestamp representing creation time.
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp representing secret expiration time.
	ExpiresAt time.Time `json:"expiresAt"`
	// Unique 24-hexadecimal character string that identifies the secret.
	Id string `json:"id"`
	// Timestamp representing last time secret usage.
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	// Secret for the service account.
	Secret *string `json:"secret,omitempty"`
}

// NewServiceAccountSecret instantiates a new ServiceAccountSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountSecret(createdAt time.Time, expiresAt time.Time, id string) *ServiceAccountSecret {
	this := ServiceAccountSecret{}
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.Id = id
	return &this
}

// NewServiceAccountSecretWithDefaults instantiates a new ServiceAccountSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountSecretWithDefaults() *ServiceAccountSecret {
	this := ServiceAccountSecret{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceAccountSecret) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecret) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceAccountSecret) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *ServiceAccountSecret) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecret) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *ServiceAccountSecret) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *ServiceAccountSecret) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecret) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccountSecret) SetId(v string) {
	o.Id = v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise
func (o *ServiceAccountSecret) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecret) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}

	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ServiceAccountSecret) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *ServiceAccountSecret) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise
func (o *ServiceAccountSecret) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecret) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}

	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ServiceAccountSecret) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ServiceAccountSecret) SetSecret(v string) {
	o.Secret = &v
}

func (o ServiceAccountSecret) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}
