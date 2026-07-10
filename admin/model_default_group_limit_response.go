// Code based on the AtlasAPI V2 OpenAPI file

package admin

// DefaultGroupLimitResponse General information about a project-level, user-configured limit, including its name, maximum value, default value, and a description.
type DefaultGroupLimitResponse struct {
	// Default limit value. Returns null if no default exists.
	// Read only field.
	DefaultLimit *int `json:"defaultLimit,omitempty"`
	// Human-friendly description of what the limit controls.
	// Read only field.
	Description *string `json:"description,omitempty"`
	// Maximum limit value. Returns null if no maximum exists.
	// Read only field.
	MaximumLimit *int `json:"maximumLimit,omitempty"`
	// Human-readable label that identifies the user-managed limit.
	// Read only field.
	Name *string `json:"name,omitempty"`
}

// NewDefaultGroupLimitResponse instantiates a new DefaultGroupLimitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultGroupLimitResponse() *DefaultGroupLimitResponse {
	this := DefaultGroupLimitResponse{}
	return &this
}

// NewDefaultGroupLimitResponseWithDefaults instantiates a new DefaultGroupLimitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultGroupLimitResponseWithDefaults() *DefaultGroupLimitResponse {
	this := DefaultGroupLimitResponse{}
	return &this
}

// GetDefaultLimit returns the DefaultLimit field value if set, zero value otherwise
func (o *DefaultGroupLimitResponse) GetDefaultLimit() int {
	if o == nil || IsNil(o.DefaultLimit) {
		var ret int
		return ret
	}
	return *o.DefaultLimit
}

// GetDefaultLimitOk returns a tuple with the DefaultLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultGroupLimitResponse) GetDefaultLimitOk() (*int, bool) {
	if o == nil || IsNil(o.DefaultLimit) {
		return nil, false
	}

	return o.DefaultLimit, true
}

// HasDefaultLimit returns a boolean if a field has been set.
func (o *DefaultGroupLimitResponse) HasDefaultLimit() bool {
	if o != nil && !IsNil(o.DefaultLimit) {
		return true
	}

	return false
}

// SetDefaultLimit gets a reference to the given int and assigns it to the DefaultLimit field.
func (o *DefaultGroupLimitResponse) SetDefaultLimit(v int) {
	o.DefaultLimit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *DefaultGroupLimitResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultGroupLimitResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DefaultGroupLimitResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DefaultGroupLimitResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise
func (o *DefaultGroupLimitResponse) GetMaximumLimit() int {
	if o == nil || IsNil(o.MaximumLimit) {
		var ret int
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultGroupLimitResponse) GetMaximumLimitOk() (*int, bool) {
	if o == nil || IsNil(o.MaximumLimit) {
		return nil, false
	}

	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *DefaultGroupLimitResponse) HasMaximumLimit() bool {
	if o != nil && !IsNil(o.MaximumLimit) {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given int and assigns it to the MaximumLimit field.
func (o *DefaultGroupLimitResponse) SetMaximumLimit(v int) {
	o.MaximumLimit = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *DefaultGroupLimitResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultGroupLimitResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DefaultGroupLimitResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DefaultGroupLimitResponse) SetName(v string) {
	o.Name = &v
}
