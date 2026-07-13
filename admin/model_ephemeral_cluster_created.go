// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// EphemeralClusterCreated Details for the ephemeral Atlas cluster that was created.
type EphemeralClusterCreated struct {
	// URL that redirects users to the Atlas registration page.
	ClaimUrl *string `json:"claimUrl,omitempty"`
	// Unique ID that can be used to check the status of an ephemeral cluster.
	ClusterId *string `json:"clusterId,omitempty"`
	// Atlas cluster connection string with SCRAM database user credentials included.
	ConnectionString *string `json:"connectionString,omitempty"`
	// Date and time when the cluster will be paused and no longer accessible until it is claimed. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Status of the ephemeral cluster.
	Status *string `json:"status,omitempty"`
	// Legal Terms of Service agreement.
	TermsOfService *string `json:"termsOfService,omitempty"`
}

// NewEphemeralClusterCreated instantiates a new EphemeralClusterCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEphemeralClusterCreated() *EphemeralClusterCreated {
	this := EphemeralClusterCreated{}
	return &this
}

// NewEphemeralClusterCreatedWithDefaults instantiates a new EphemeralClusterCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEphemeralClusterCreatedWithDefaults() *EphemeralClusterCreated {
	this := EphemeralClusterCreated{}
	return &this
}

// GetClaimUrl returns the ClaimUrl field value if set, zero value otherwise
func (o *EphemeralClusterCreated) GetClaimUrl() string {
	if o == nil || IsNil(o.ClaimUrl) {
		var ret string
		return ret
	}
	return *o.ClaimUrl
}

// GetClaimUrlOk returns a tuple with the ClaimUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterCreated) GetClaimUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimUrl) {
		return nil, false
	}

	return o.ClaimUrl, true
}

// HasClaimUrl returns a boolean if a field has been set.
func (o *EphemeralClusterCreated) HasClaimUrl() bool {
	if o != nil && !IsNil(o.ClaimUrl) {
		return true
	}

	return false
}

// SetClaimUrl gets a reference to the given string and assigns it to the ClaimUrl field.
func (o *EphemeralClusterCreated) SetClaimUrl(v string) {
	o.ClaimUrl = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *EphemeralClusterCreated) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterCreated) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *EphemeralClusterCreated) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *EphemeralClusterCreated) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise
func (o *EphemeralClusterCreated) GetConnectionString() string {
	if o == nil || IsNil(o.ConnectionString) {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterCreated) GetConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionString) {
		return nil, false
	}

	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *EphemeralClusterCreated) HasConnectionString() bool {
	if o != nil && !IsNil(o.ConnectionString) {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *EphemeralClusterCreated) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *EphemeralClusterCreated) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterCreated) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *EphemeralClusterCreated) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *EphemeralClusterCreated) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *EphemeralClusterCreated) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterCreated) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EphemeralClusterCreated) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EphemeralClusterCreated) SetStatus(v string) {
	o.Status = &v
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise
func (o *EphemeralClusterCreated) GetTermsOfService() string {
	if o == nil || IsNil(o.TermsOfService) {
		var ret string
		return ret
	}
	return *o.TermsOfService
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterCreated) GetTermsOfServiceOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfService) {
		return nil, false
	}

	return o.TermsOfService, true
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *EphemeralClusterCreated) HasTermsOfService() bool {
	if o != nil && !IsNil(o.TermsOfService) {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given string and assigns it to the TermsOfService field.
func (o *EphemeralClusterCreated) SetTermsOfService(v string) {
	o.TermsOfService = &v
}
