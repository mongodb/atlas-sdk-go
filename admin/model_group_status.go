// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// GroupStatus struct for GroupStatus
type GroupStatus struct {
	DisplayName *string    `json:"displayName,omitempty"`
	Since       *time.Time `json:"since,omitempty"`
	Status      *string    `json:"status,omitempty"`
}

// NewGroupStatus instantiates a new GroupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupStatus() *GroupStatus {
	this := GroupStatus{}
	return &this
}

// NewGroupStatusWithDefaults instantiates a new GroupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupStatusWithDefaults() *GroupStatus {
	this := GroupStatus{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise
func (o *GroupStatus) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStatus) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}

	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GroupStatus) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GroupStatus) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetSince returns the Since field value if set, zero value otherwise
func (o *GroupStatus) GetSince() time.Time {
	if o == nil || IsNil(o.Since) {
		var ret time.Time
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStatus) GetSinceOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Since) {
		return nil, false
	}

	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *GroupStatus) HasSince() bool {
	if o != nil && !IsNil(o.Since) {
		return true
	}

	return false
}

// SetSince gets a reference to the given time.Time and assigns it to the Since field.
func (o *GroupStatus) SetSince(v time.Time) {
	o.Since = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *GroupStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GroupStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GroupStatus) SetStatus(v string) {
	o.Status = &v
}

func (o GroupStatus) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Since) {
		toSerialize["since"] = o.Since
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}
