// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgServiceAccountRequest Organization Service Account that Atlas creates for this organization. If omitted, Atlas doesn't create an organization Service Account for this organization. If specified, this object requires all body parameters. Note that API Keys cannot be specified in the same request.
type OrgServiceAccountRequest struct {
	// Human readable description for the Service Account.
	Description string `json:"description"`
	// Human-readable name for the Service Account. The name is modifiable and does not have to be unique.
	Name string `json:"name"`
	// A list of organization-level roles for the Service Account.
	Roles []string `json:"roles"`
	// The expiration time of the new Service Account secret, provided in hours. The minimum and maximum allowed expiration times are subject to change and are controlled by the organization's settings. Required unless `withoutInitialSecret` is true.
	SecretExpiresAfterHours *int `json:"secretExpiresAfterHours,omitempty"`
	// If true, creates the Service Account without generating an initial secret. `secretExpiresAfterHours` must not be set when this is true. Defaults to false, which preserves existing behavior: a secret is generated and returned in the response. Use the `CreateOrgServiceAccountSecret` endpoint to add a secret later.
	WithoutInitialSecret *bool `json:"withoutInitialSecret,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OrgServiceAccountRequest) MarshalJSON() ([]byte, error) {
	type noMethod OrgServiceAccountRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOrgServiceAccountRequest instantiates a new OrgServiceAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgServiceAccountRequest(description string, name string, roles []string) *OrgServiceAccountRequest {
	this := OrgServiceAccountRequest{}
	this.Description = description
	this.Name = name
	this.Roles = roles
	var withoutInitialSecret bool = false
	this.WithoutInitialSecret = &withoutInitialSecret
	return &this
}

// NewOrgServiceAccountRequestWithDefaults instantiates a new OrgServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgServiceAccountRequestWithDefaults() *OrgServiceAccountRequest {
	this := OrgServiceAccountRequest{}
	var withoutInitialSecret bool = false
	this.WithoutInitialSecret = &withoutInitialSecret
	return &this
}

// GetDescription returns the Description field value
func (o *OrgServiceAccountRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OrgServiceAccountRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *OrgServiceAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrgServiceAccountRequest) SetName(v string) {
	o.Name = v
}

// GetRoles returns the Roles field value
func (o *OrgServiceAccountRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *OrgServiceAccountRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetSecretExpiresAfterHours returns the SecretExpiresAfterHours field value if set, zero value otherwise
func (o *OrgServiceAccountRequest) GetSecretExpiresAfterHours() int {
	if o == nil || IsNil(o.SecretExpiresAfterHours) {
		var ret int
		return ret
	}
	return *o.SecretExpiresAfterHours
}

// GetSecretExpiresAfterHoursOk returns a tuple with the SecretExpiresAfterHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetSecretExpiresAfterHoursOk() (*int, bool) {
	if o == nil || IsNil(o.SecretExpiresAfterHours) {
		return nil, false
	}

	return o.SecretExpiresAfterHours, true
}

// HasSecretExpiresAfterHours returns a boolean if a field has been set.
func (o *OrgServiceAccountRequest) HasSecretExpiresAfterHours() bool {
	if o != nil && !IsNil(o.SecretExpiresAfterHours) {
		return true
	}

	return false
}

// SetSecretExpiresAfterHours gets a reference to the given int and assigns it to the SecretExpiresAfterHours field.
func (o *OrgServiceAccountRequest) SetSecretExpiresAfterHours(v int) {
	o.SecretExpiresAfterHours = &v
	o.NullFields = removeNullField(o.NullFields, "SecretExpiresAfterHours")
}

// SetSecretExpiresAfterHoursNil sets SecretExpiresAfterHours to an explicit JSON null when marshaled.
func (o *OrgServiceAccountRequest) SetSecretExpiresAfterHoursNil() {
	o.SecretExpiresAfterHours = nil
	o.NullFields = addNullField(o.NullFields, "SecretExpiresAfterHours")
}

// GetWithoutInitialSecret returns the WithoutInitialSecret field value if set, zero value otherwise
func (o *OrgServiceAccountRequest) GetWithoutInitialSecret() bool {
	if o == nil || IsNil(o.WithoutInitialSecret) {
		var ret bool
		return ret
	}
	return *o.WithoutInitialSecret
}

// GetWithoutInitialSecretOk returns a tuple with the WithoutInitialSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetWithoutInitialSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.WithoutInitialSecret) {
		return nil, false
	}

	return o.WithoutInitialSecret, true
}

// HasWithoutInitialSecret returns a boolean if a field has been set.
func (o *OrgServiceAccountRequest) HasWithoutInitialSecret() bool {
	if o != nil && !IsNil(o.WithoutInitialSecret) {
		return true
	}

	return false
}

// SetWithoutInitialSecret gets a reference to the given bool and assigns it to the WithoutInitialSecret field.
func (o *OrgServiceAccountRequest) SetWithoutInitialSecret(v bool) {
	o.WithoutInitialSecret = &v
	o.NullFields = removeNullField(o.NullFields, "WithoutInitialSecret")
}

// SetWithoutInitialSecretNil sets WithoutInitialSecret to an explicit JSON null when marshaled.
func (o *OrgServiceAccountRequest) SetWithoutInitialSecretNil() {
	o.WithoutInitialSecret = nil
	o.NullFields = addNullField(o.NullFields, "WithoutInitialSecret")
}
