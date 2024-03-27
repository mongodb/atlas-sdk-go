// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// ServiceAccountSecretDetails struct for ServiceAccountSecretDetails
type ServiceAccountSecretDetails struct {
	// Timestamp representing creation time.
	// Read only field.
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp representing secret expiration time.
	// Read only field.
	ExpiresAt time.Time `json:"expiresAt"`
	// Unique 24-hexadecimal character string that identifies the secret.
	// Read only field.
	Id string `json:"id"`
	// Timestamp representing last time secret usage.
	// Read only field.
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	// Value of client secret with masked values of the first 20 characters.
	// Read only field.
	MaskedSecretValue string `json:"maskedSecretValue"`
}

// NewServiceAccountSecretDetails instantiates a new ServiceAccountSecretDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountSecretDetails(createdAt time.Time, expiresAt time.Time, id string, maskedSecretValue string) *ServiceAccountSecretDetails {
	this := ServiceAccountSecretDetails{}
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.Id = id
	this.MaskedSecretValue = maskedSecretValue
	return &this
}

// NewServiceAccountSecretDetailsWithDefaults instantiates a new ServiceAccountSecretDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountSecretDetailsWithDefaults() *ServiceAccountSecretDetails {
	this := ServiceAccountSecretDetails{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceAccountSecretDetails) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecretDetails) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceAccountSecretDetails) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *ServiceAccountSecretDetails) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecretDetails) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *ServiceAccountSecretDetails) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *ServiceAccountSecretDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecretDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAccountSecretDetails) SetId(v string) {
	o.Id = v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise
func (o *ServiceAccountSecretDetails) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecretDetails) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}

	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ServiceAccountSecretDetails) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *ServiceAccountSecretDetails) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetMaskedSecretValue returns the MaskedSecretValue field value
func (o *ServiceAccountSecretDetails) GetMaskedSecretValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskedSecretValue
}

// GetMaskedSecretValueOk returns a tuple with the MaskedSecretValue field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountSecretDetails) GetMaskedSecretValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskedSecretValue, true
}

// SetMaskedSecretValue sets field value
func (o *ServiceAccountSecretDetails) SetMaskedSecretValue(v string) {
	o.MaskedSecretValue = v
}

func (o ServiceAccountSecretDetails) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountSecretDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}
