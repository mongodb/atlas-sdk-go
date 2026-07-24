// Code based on the AtlasAPI V2 OpenAPI file

package admin

// LDAPVerifyConnectivityJobRequestValidation One test that MongoDB Cloud runs to test verification of the provided Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration details.
type LDAPVerifyConnectivityJobRequestValidation struct {
	// Human-readable string that indicates the result of this verification test.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// Human-readable label that identifies this verification test that MongoDB Cloud runs.
	// Read only field.
	ValidationType *string `json:"validationType,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *LDAPVerifyConnectivityJobRequestValidation) MarshalJSON() ([]byte, error) {
	type noMethod LDAPVerifyConnectivityJobRequestValidation
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewLDAPVerifyConnectivityJobRequestValidation instantiates a new LDAPVerifyConnectivityJobRequestValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPVerifyConnectivityJobRequestValidation() *LDAPVerifyConnectivityJobRequestValidation {
	this := LDAPVerifyConnectivityJobRequestValidation{}
	return &this
}

// NewLDAPVerifyConnectivityJobRequestValidationWithDefaults instantiates a new LDAPVerifyConnectivityJobRequestValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPVerifyConnectivityJobRequestValidationWithDefaults() *LDAPVerifyConnectivityJobRequestValidation {
	this := LDAPVerifyConnectivityJobRequestValidation{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *LDAPVerifyConnectivityJobRequestValidation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LDAPVerifyConnectivityJobRequestValidation) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *LDAPVerifyConnectivityJobRequestValidation) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}

// GetValidationType returns the ValidationType field value if set, zero value otherwise
func (o *LDAPVerifyConnectivityJobRequestValidation) GetValidationType() string {
	if o == nil || IsNil(o.ValidationType) {
		var ret string
		return ret
	}
	return *o.ValidationType
}

// GetValidationTypeOk returns a tuple with the ValidationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) GetValidationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationType) {
		return nil, false
	}

	return o.ValidationType, true
}

// HasValidationType returns a boolean if a field has been set.
func (o *LDAPVerifyConnectivityJobRequestValidation) HasValidationType() bool {
	if o != nil && !IsNil(o.ValidationType) {
		return true
	}

	return false
}

// SetValidationType gets a reference to the given string and assigns it to the ValidationType field.
func (o *LDAPVerifyConnectivityJobRequestValidation) SetValidationType(v string) {
	o.ValidationType = &v
	o.NullFields = removeNullField(o.NullFields, "ValidationType")
}

// SetValidationTypeNil sets ValidationType to an explicit JSON null when marshaled.
func (o *LDAPVerifyConnectivityJobRequestValidation) SetValidationTypeNil() {
	o.ValidationType = nil
	o.NullFields = addNullField(o.NullFields, "ValidationType")
}
