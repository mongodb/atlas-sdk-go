// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// ApiStandbyLinkFailoverResponse Response containing failover operation details.
type ApiStandbyLinkFailoverResponse struct {
	// Date and time when the failover operation completed. Only present when status is COMPLETE. This value is in the ISO 8601 date and time format in UTC.
	// Read only field.
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Date and time when the failover operation was created. This value is in the ISO 8601 date and time format in UTC.
	// Read only field.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the failover operation.
	// Read only field.
	FailoverId *string `json:"failoverId,omitempty"`
	// Current stage of the failover operation. Indicates progression from initiation to completion.
	// Read only field.
	FailoverStage *string `json:"failoverStage,omitempty"`
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

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}

	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *ApiStandbyLinkFailoverResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}

	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *ApiStandbyLinkFailoverResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetFailoverId returns the FailoverId field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetFailoverId() string {
	if o == nil || IsNil(o.FailoverId) {
		var ret string
		return ret
	}
	return *o.FailoverId
}

// GetFailoverIdOk returns a tuple with the FailoverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetFailoverIdOk() (*string, bool) {
	if o == nil || IsNil(o.FailoverId) {
		return nil, false
	}

	return o.FailoverId, true
}

// HasFailoverId returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasFailoverId() bool {
	if o != nil && !IsNil(o.FailoverId) {
		return true
	}

	return false
}

// SetFailoverId gets a reference to the given string and assigns it to the FailoverId field.
func (o *ApiStandbyLinkFailoverResponse) SetFailoverId(v string) {
	o.FailoverId = &v
}

// GetFailoverStage returns the FailoverStage field value if set, zero value otherwise
func (o *ApiStandbyLinkFailoverResponse) GetFailoverStage() string {
	if o == nil || IsNil(o.FailoverStage) {
		var ret string
		return ret
	}
	return *o.FailoverStage
}

// GetFailoverStageOk returns a tuple with the FailoverStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverResponse) GetFailoverStageOk() (*string, bool) {
	if o == nil || IsNil(o.FailoverStage) {
		return nil, false
	}

	return o.FailoverStage, true
}

// HasFailoverStage returns a boolean if a field has been set.
func (o *ApiStandbyLinkFailoverResponse) HasFailoverStage() bool {
	if o != nil && !IsNil(o.FailoverStage) {
		return true
	}

	return false
}

// SetFailoverStage gets a reference to the given string and assigns it to the FailoverStage field.
func (o *ApiStandbyLinkFailoverResponse) SetFailoverStage(v string) {
	o.FailoverStage = &v
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
