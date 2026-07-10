// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SqlInterfaceStatusResponse struct for SqlInterfaceStatusResponse
type SqlInterfaceStatusResponse struct {
	// Current or desired SQL Interface status for the cluster. The GET endpoint returns OFF when no status is configured.
	Status string `json:"status"`
	// Whether the cluster tier and type support SQL Interface.
	// Read only field.
	Supported *bool `json:"supported,omitempty"`
	// Machine-readable reason why SQL Interface isn't supported for this cluster. Present only when supported is false.
	// Read only field.
	UnsupportedReason *string `json:"unsupportedReason,omitempty"`
}

// NewSqlInterfaceStatusResponse instantiates a new SqlInterfaceStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlInterfaceStatusResponse(status string) *SqlInterfaceStatusResponse {
	this := SqlInterfaceStatusResponse{}
	this.Status = status
	return &this
}

// NewSqlInterfaceStatusResponseWithDefaults instantiates a new SqlInterfaceStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlInterfaceStatusResponseWithDefaults() *SqlInterfaceStatusResponse {
	this := SqlInterfaceStatusResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *SqlInterfaceStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SqlInterfaceStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SqlInterfaceStatusResponse) SetStatus(v string) {
	o.Status = v
}

// GetSupported returns the Supported field value if set, zero value otherwise
func (o *SqlInterfaceStatusResponse) GetSupported() bool {
	if o == nil || IsNil(o.Supported) {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlInterfaceStatusResponse) GetSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.Supported) {
		return nil, false
	}

	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *SqlInterfaceStatusResponse) HasSupported() bool {
	if o != nil && !IsNil(o.Supported) {
		return true
	}

	return false
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *SqlInterfaceStatusResponse) SetSupported(v bool) {
	o.Supported = &v
}

// GetUnsupportedReason returns the UnsupportedReason field value if set, zero value otherwise
func (o *SqlInterfaceStatusResponse) GetUnsupportedReason() string {
	if o == nil || IsNil(o.UnsupportedReason) {
		var ret string
		return ret
	}
	return *o.UnsupportedReason
}

// GetUnsupportedReasonOk returns a tuple with the UnsupportedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlInterfaceStatusResponse) GetUnsupportedReasonOk() (*string, bool) {
	if o == nil || IsNil(o.UnsupportedReason) {
		return nil, false
	}

	return o.UnsupportedReason, true
}

// HasUnsupportedReason returns a boolean if a field has been set.
func (o *SqlInterfaceStatusResponse) HasUnsupportedReason() bool {
	if o != nil && !IsNil(o.UnsupportedReason) {
		return true
	}

	return false
}

// SetUnsupportedReason gets a reference to the given string and assigns it to the UnsupportedReason field.
func (o *SqlInterfaceStatusResponse) SetUnsupportedReason(v string) {
	o.UnsupportedReason = &v
}
