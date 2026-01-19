// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiStandbyLinkFailoverResponse Response containing failover operation details.
type ApiStandbyLinkFailoverResponse struct {
	// Unique 24-hexadecimal digit string that identifies the failover operation.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the standby link.
	// Read only field.
	StandbyLinkId *string `json:"standbyLinkId,omitempty"`
	// Current status of the failover operation.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// Type of failover operation to perform.
	Type *string `json:"type,omitempty"`
}

// NewApiStandbyLinkFailoverResponse instantiates a new ApiStandbyLinkFailoverResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStandbyLinkFailoverResponse() *ApiStandbyLinkFailoverResponse {
	this := ApiStandbyLinkFailoverResponse{}
	return &this
}

// NewApiStandbyLinkFailoverResponseWithDefaults instantiates a new ApiStandbyLinkFailoverResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStandbyLinkFailoverResponseWithDefaults() *ApiStandbyLinkFailoverResponse {
	this := ApiStandbyLinkFailoverResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiStandbyLinkFailoverResponse) SetId(v string) {
	o.Id = &v
}

// GetStandbyLinkId returns the StandbyLinkId field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetStandbyLinkId() string {
	if o == nil || IsNil(o.StandbyLinkId) {
		var ret string
		return ret
	}
	return *o.StandbyLinkId
}

// GetStandbyLinkIdOk returns a tuple with the StandbyLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetStandbyLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandbyLinkId) {
		return nil, false
	}

	return o.StandbyLinkId, true
}

// HasStandbyLinkId returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasStandbyLinkId() bool {
	if o != nil && !IsNil(o.StandbyLinkId) {
		return true
	}

	return false
}

// SetStandbyLinkId gets a reference to the given string and assigns it to the StandbyLinkId field.
func (o *ApiStandbyLinkFailoverResponse) SetStandbyLinkId(v string) {
	o.StandbyLinkId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiStandbyLinkFailoverResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiStandbyLinkFailoverResponse) SetType(v string) {
	o.Type = &v
}
