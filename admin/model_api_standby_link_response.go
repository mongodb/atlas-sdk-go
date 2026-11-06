// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiStandbyLinkResponse Response body containing standby link information.
type ApiStandbyLinkResponse struct {
	// Unique 24-hexadecimal digit string that identifies the active cluster in the standby link.
	ActiveClusterUniqueId string `json:"activeClusterUniqueId"`
	// Reason the standby link is in an error state. Only present when the status is DEGRADED or FAILED.
	// Read only field.
	ErrorReason *string `json:"errorReason,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project containing the clusters in the standby link.
	// Read only field.
	GroupId string `json:"groupId"`
	// Unique 24-hexadecimal digit string that identifies the standby cluster in the standby link.
	StandbyClusterUniqueId string `json:"standbyClusterUniqueId"`
	// Unique 24-hexadecimal digit string that identifies the standby link.
	// Read only field.
	StandbyLinkId string `json:"standbyLinkId"`
	// Current status of the standby link.
	// Read only field.
	Status string `json:"status"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiStandbyLinkResponse) MarshalJSON() ([]byte, error) {
	type noMethod ApiStandbyLinkResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiStandbyLinkResponse instantiates a new ApiStandbyLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStandbyLinkResponse(activeClusterUniqueId string, groupId string, standbyClusterUniqueId string, standbyLinkId string, status string) *ApiStandbyLinkResponse {
	this := ApiStandbyLinkResponse{}
	this.ActiveClusterUniqueId = activeClusterUniqueId
	this.GroupId = groupId
	this.StandbyClusterUniqueId = standbyClusterUniqueId
	this.StandbyLinkId = standbyLinkId
	this.Status = status
	return &this
}

// NewApiStandbyLinkResponseWithDefaults instantiates a new ApiStandbyLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStandbyLinkResponseWithDefaults() *ApiStandbyLinkResponse {
	this := ApiStandbyLinkResponse{}
	return &this
}

// GetActiveClusterUniqueId returns the ActiveClusterUniqueId field value
func (o *ApiStandbyLinkResponse) GetActiveClusterUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveClusterUniqueId
}

// GetActiveClusterUniqueIdOk returns a tuple with the ActiveClusterUniqueId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkResponse) GetActiveClusterUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveClusterUniqueId, true
}

// SetActiveClusterUniqueId sets field value
func (o *ApiStandbyLinkResponse) SetActiveClusterUniqueId(v string) {
	o.ActiveClusterUniqueId = v
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise
func (o *ApiStandbyLinkResponse) GetErrorReason() string {
	if o == nil || IsNil(o.ErrorReason) {
		var ret string
		return ret
	}
	return *o.ErrorReason
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkResponse) GetErrorReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorReason) {
		return nil, false
	}

	return o.ErrorReason, true
}

// HasErrorReason returns a boolean if a field has been set.
func (o *ApiStandbyLinkResponse) HasErrorReason() bool {
	if o != nil && !IsNil(o.ErrorReason) {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given string and assigns it to the ErrorReason field.
func (o *ApiStandbyLinkResponse) SetErrorReason(v string) {
	o.ErrorReason = &v
	o.NullFields = removeNullField(o.NullFields, "ErrorReason")
}

// SetErrorReasonNil sets ErrorReason to an explicit JSON null when marshaled.
func (o *ApiStandbyLinkResponse) SetErrorReasonNil() {
	o.ErrorReason = nil
	o.NullFields = addNullField(o.NullFields, "ErrorReason")
}

// GetGroupId returns the GroupId field value
func (o *ApiStandbyLinkResponse) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkResponse) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ApiStandbyLinkResponse) SetGroupId(v string) {
	o.GroupId = v
}

// GetStandbyClusterUniqueId returns the StandbyClusterUniqueId field value
func (o *ApiStandbyLinkResponse) GetStandbyClusterUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StandbyClusterUniqueId
}

// GetStandbyClusterUniqueIdOk returns a tuple with the StandbyClusterUniqueId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkResponse) GetStandbyClusterUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandbyClusterUniqueId, true
}

// SetStandbyClusterUniqueId sets field value
func (o *ApiStandbyLinkResponse) SetStandbyClusterUniqueId(v string) {
	o.StandbyClusterUniqueId = v
}

// GetStandbyLinkId returns the StandbyLinkId field value
func (o *ApiStandbyLinkResponse) GetStandbyLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StandbyLinkId
}

// GetStandbyLinkIdOk returns a tuple with the StandbyLinkId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkResponse) GetStandbyLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandbyLinkId, true
}

// SetStandbyLinkId sets field value
func (o *ApiStandbyLinkResponse) SetStandbyLinkId(v string) {
	o.StandbyLinkId = v
}

// GetStatus returns the Status field value
func (o *ApiStandbyLinkResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiStandbyLinkResponse) SetStatus(v string) {
	o.Status = v
}
