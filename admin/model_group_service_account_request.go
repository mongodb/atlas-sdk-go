// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupServiceAccountRequest struct for GroupServiceAccountRequest
type GroupServiceAccountRequest struct {
	// Human readable description for the Service Account.
	Description string `json:"description"`
	// Human-readable name for the Service Account. The name is modifiable and does not have to be unique.
	Name string `json:"name"`
	// A list of project-level roles for the Service Account.
	Roles []string `json:"roles"`
	// The expiration time of the new Service Account secret, provided in hours. The minimum and maximum allowed expiration times are subject to change and are controlled by the organization's settings. Required unless `withoutInitialSecret` is true.
	SecretExpiresAfterHours *int `json:"secretExpiresAfterHours,omitempty"`
	// If true, creates the Service Account without generating an initial secret. `secretExpiresAfterHours` must not be set when this is true. Defaults to false, which preserves existing behavior: a secret is generated and returned in the response. Use the `CreateGroupServiceAccountSecret` endpoint to add a secret later.
	WithoutInitialSecret *bool `json:"withoutInitialSecret,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GroupServiceAccountRequest) MarshalJSON() ([]byte, error) {
	type noMethod GroupServiceAccountRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGroupServiceAccountRequest instantiates a new GroupServiceAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupServiceAccountRequest(description string, name string, roles []string) *GroupServiceAccountRequest {
	this := GroupServiceAccountRequest{}
	this.Description = description
	this.Name = name
	this.Roles = roles
	var withoutInitialSecret bool = false
	this.WithoutInitialSecret = &withoutInitialSecret
	return &this
}

// NewGroupServiceAccountRequestWithDefaults instantiates a new GroupServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupServiceAccountRequestWithDefaults() *GroupServiceAccountRequest {
	this := GroupServiceAccountRequest{}
	var withoutInitialSecret bool = false
	this.WithoutInitialSecret = &withoutInitialSecret
	return &this
}

// GetDescription returns the Description field value
func (o *GroupServiceAccountRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GroupServiceAccountRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *GroupServiceAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupServiceAccountRequest) SetName(v string) {
	o.Name = v
}

// GetRoles returns the Roles field value
func (o *GroupServiceAccountRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *GroupServiceAccountRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetSecretExpiresAfterHours returns the SecretExpiresAfterHours field value if set, zero value otherwise
func (o *GroupServiceAccountRequest) GetSecretExpiresAfterHours() int {
	if o == nil || IsNil(o.SecretExpiresAfterHours) {
		var ret int
		return ret
	}
	return *o.SecretExpiresAfterHours
}

// GetSecretExpiresAfterHoursOk returns a tuple with the SecretExpiresAfterHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetSecretExpiresAfterHoursOk() (*int, bool) {
	if o == nil || IsNil(o.SecretExpiresAfterHours) {
		return nil, false
	}

	return o.SecretExpiresAfterHours, true
}

// HasSecretExpiresAfterHours returns a boolean if a field has been set.
func (o *GroupServiceAccountRequest) HasSecretExpiresAfterHours() bool {
	if o != nil && !IsNil(o.SecretExpiresAfterHours) {
		return true
	}

	return false
}

// SetSecretExpiresAfterHours gets a reference to the given int and assigns it to the SecretExpiresAfterHours field.
func (o *GroupServiceAccountRequest) SetSecretExpiresAfterHours(v int) {
	o.SecretExpiresAfterHours = &v
	o.NullFields = removeNullField(o.NullFields, "SecretExpiresAfterHours")
}

// SetSecretExpiresAfterHoursNil sets SecretExpiresAfterHours to an explicit JSON null when marshaled.
func (o *GroupServiceAccountRequest) SetSecretExpiresAfterHoursNil() {
	o.SecretExpiresAfterHours = nil
	o.NullFields = addNullField(o.NullFields, "SecretExpiresAfterHours")
}

// GetWithoutInitialSecret returns the WithoutInitialSecret field value if set, zero value otherwise
func (o *GroupServiceAccountRequest) GetWithoutInitialSecret() bool {
	if o == nil || IsNil(o.WithoutInitialSecret) {
		var ret bool
		return ret
	}
	return *o.WithoutInitialSecret
}

// GetWithoutInitialSecretOk returns a tuple with the WithoutInitialSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetWithoutInitialSecretOk() (*bool, bool) {
	if o == nil || IsNil(o.WithoutInitialSecret) {
		return nil, false
	}

	return o.WithoutInitialSecret, true
}

// HasWithoutInitialSecret returns a boolean if a field has been set.
func (o *GroupServiceAccountRequest) HasWithoutInitialSecret() bool {
	if o != nil && !IsNil(o.WithoutInitialSecret) {
		return true
	}

	return false
}

// SetWithoutInitialSecret gets a reference to the given bool and assigns it to the WithoutInitialSecret field.
func (o *GroupServiceAccountRequest) SetWithoutInitialSecret(v bool) {
	o.WithoutInitialSecret = &v
	o.NullFields = removeNullField(o.NullFields, "WithoutInitialSecret")
}

// SetWithoutInitialSecretNil sets WithoutInitialSecret to an explicit JSON null when marshaled.
func (o *GroupServiceAccountRequest) SetWithoutInitialSecretNil() {
	o.WithoutInitialSecret = nil
	o.NullFields = addNullField(o.NullFields, "WithoutInitialSecret")
}
