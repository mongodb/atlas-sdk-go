// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// EphemeralClusterStatusResponse Current status and details of an ephemeral Atlas cluster.
type EphemeralClusterStatusResponse struct {
	// URL that redirects users to the Atlas registration page.
	ClaimUrl *string `json:"claimUrl,omitempty"`
	// Unique ID of the ephemeral cluster.
	ClusterId *string `json:"clusterId,omitempty"`
	// Atlas cluster connection string with the SCRAM database user name included. The password is replaced by a placeholder, since it is only ever returned when the cluster is created.
	ConnectionString *string `json:"connectionString,omitempty"`
	// Date and time when the cluster will be paused and no longer accessible until it is claimed. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Status of the ephemeral cluster.
	Status *string `json:"status,omitempty"`
	// Legal Terms of Service agreement.
	TermsOfService *string `json:"termsOfService,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *EphemeralClusterStatusResponse) MarshalJSON() ([]byte, error) {
	type noMethod EphemeralClusterStatusResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewEphemeralClusterStatusResponse instantiates a new EphemeralClusterStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEphemeralClusterStatusResponse() *EphemeralClusterStatusResponse {
	this := EphemeralClusterStatusResponse{}
	return &this
}

// NewEphemeralClusterStatusResponseWithDefaults instantiates a new EphemeralClusterStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEphemeralClusterStatusResponseWithDefaults() *EphemeralClusterStatusResponse {
	this := EphemeralClusterStatusResponse{}
	return &this
}

// GetClaimUrl returns the ClaimUrl field value if set, zero value otherwise
func (o *EphemeralClusterStatusResponse) GetClaimUrl() string {
	if o == nil || IsNil(o.ClaimUrl) {
		var ret string
		return ret
	}
	return *o.ClaimUrl
}

// GetClaimUrlOk returns a tuple with the ClaimUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterStatusResponse) GetClaimUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimUrl) {
		return nil, false
	}

	return o.ClaimUrl, true
}

// HasClaimUrl returns a boolean if a field has been set.
func (o *EphemeralClusterStatusResponse) HasClaimUrl() bool {
	if o != nil && !IsNil(o.ClaimUrl) {
		return true
	}

	return false
}

// SetClaimUrl gets a reference to the given string and assigns it to the ClaimUrl field.
func (o *EphemeralClusterStatusResponse) SetClaimUrl(v string) {
	o.ClaimUrl = &v
	o.NullFields = removeNullField(o.NullFields, "ClaimUrl")
}

// SetClaimUrlNil sets ClaimUrl to an explicit JSON null when marshaled.
func (o *EphemeralClusterStatusResponse) SetClaimUrlNil() {
	o.ClaimUrl = nil
	o.NullFields = addNullField(o.NullFields, "ClaimUrl")
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *EphemeralClusterStatusResponse) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterStatusResponse) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *EphemeralClusterStatusResponse) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *EphemeralClusterStatusResponse) SetClusterId(v string) {
	o.ClusterId = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterId")
}

// SetClusterIdNil sets ClusterId to an explicit JSON null when marshaled.
func (o *EphemeralClusterStatusResponse) SetClusterIdNil() {
	o.ClusterId = nil
	o.NullFields = addNullField(o.NullFields, "ClusterId")
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise
func (o *EphemeralClusterStatusResponse) GetConnectionString() string {
	if o == nil || IsNil(o.ConnectionString) {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterStatusResponse) GetConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionString) {
		return nil, false
	}

	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *EphemeralClusterStatusResponse) HasConnectionString() bool {
	if o != nil && !IsNil(o.ConnectionString) {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *EphemeralClusterStatusResponse) SetConnectionString(v string) {
	o.ConnectionString = &v
	o.NullFields = removeNullField(o.NullFields, "ConnectionString")
}

// SetConnectionStringNil sets ConnectionString to an explicit JSON null when marshaled.
func (o *EphemeralClusterStatusResponse) SetConnectionStringNil() {
	o.ConnectionString = nil
	o.NullFields = addNullField(o.NullFields, "ConnectionString")
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *EphemeralClusterStatusResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterStatusResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *EphemeralClusterStatusResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *EphemeralClusterStatusResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
	o.NullFields = removeNullField(o.NullFields, "ExpiresAt")
}

// SetExpiresAtNil sets ExpiresAt to an explicit JSON null when marshaled.
func (o *EphemeralClusterStatusResponse) SetExpiresAtNil() {
	o.ExpiresAt = nil
	o.NullFields = addNullField(o.NullFields, "ExpiresAt")
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *EphemeralClusterStatusResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EphemeralClusterStatusResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EphemeralClusterStatusResponse) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *EphemeralClusterStatusResponse) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise
func (o *EphemeralClusterStatusResponse) GetTermsOfService() string {
	if o == nil || IsNil(o.TermsOfService) {
		var ret string
		return ret
	}
	return *o.TermsOfService
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralClusterStatusResponse) GetTermsOfServiceOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfService) {
		return nil, false
	}

	return o.TermsOfService, true
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *EphemeralClusterStatusResponse) HasTermsOfService() bool {
	if o != nil && !IsNil(o.TermsOfService) {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given string and assigns it to the TermsOfService field.
func (o *EphemeralClusterStatusResponse) SetTermsOfService(v string) {
	o.TermsOfService = &v
	o.NullFields = removeNullField(o.NullFields, "TermsOfService")
}

// SetTermsOfServiceNil sets TermsOfService to an explicit JSON null when marshaled.
func (o *EphemeralClusterStatusResponse) SetTermsOfServiceNil() {
	o.TermsOfService = nil
	o.NullFields = addNullField(o.NullFields, "TermsOfService")
}
