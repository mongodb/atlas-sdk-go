// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the LiveImportValidation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveImportValidation{}

// LiveImportValidation struct for LiveImportValidation
type LiveImportValidation struct {
	// Unique 24-hexadecimal digit string that identifies the validation.
	Id *string `json:"_id,omitempty"`
	// Reason why the validation job failed.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project to validate.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the source project.
	SourceGroupId *string `json:"sourceGroupId,omitempty"`
	// State of the specified validation job returned at the time of the request.
	Status NullableString `json:"status,omitempty"`
}

// NewLiveImportValidation instantiates a new LiveImportValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveImportValidation() *LiveImportValidation {
	this := LiveImportValidation{}
	return &this
}

// NewLiveImportValidationWithDefaults instantiates a new LiveImportValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveImportValidationWithDefaults() *LiveImportValidation {
	this := LiveImportValidation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveImportValidation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveImportValidation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveImportValidation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveImportValidation) SetId(v string) {
	o.Id = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveImportValidation) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveImportValidation) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *LiveImportValidation) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *LiveImportValidation) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *LiveImportValidation) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *LiveImportValidation) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *LiveImportValidation) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveImportValidation) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *LiveImportValidation) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *LiveImportValidation) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSourceGroupId returns the SourceGroupId field value if set, zero value otherwise.
func (o *LiveImportValidation) GetSourceGroupId() string {
	if o == nil || IsNil(o.SourceGroupId) {
		var ret string
		return ret
	}
	return *o.SourceGroupId
}

// GetSourceGroupIdOk returns a tuple with the SourceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveImportValidation) GetSourceGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceGroupId) {
		return nil, false
	}
	return o.SourceGroupId, true
}

// HasSourceGroupId returns a boolean if a field has been set.
func (o *LiveImportValidation) HasSourceGroupId() bool {
	if o != nil && !IsNil(o.SourceGroupId) {
		return true
	}

	return false
}

// SetSourceGroupId gets a reference to the given string and assigns it to the SourceGroupId field.
func (o *LiveImportValidation) SetSourceGroupId(v string) {
	o.SourceGroupId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveImportValidation) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveImportValidation) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *LiveImportValidation) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *LiveImportValidation) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *LiveImportValidation) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *LiveImportValidation) UnsetStatus() {
	o.Status.Unset()
}

func (o LiveImportValidation) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LiveImportValidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if !IsNil(o.SourceGroupId) {
		toSerialize["sourceGroupId"] = o.SourceGroupId
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return toSerialize, nil
}

type NullableLiveImportValidation struct {
	value *LiveImportValidation
	isSet bool
}

func (v NullableLiveImportValidation) Get() *LiveImportValidation {
	return v.value
}

func (v *NullableLiveImportValidation) Set(val *LiveImportValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveImportValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveImportValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveImportValidation(val *LiveImportValidation) *NullableLiveImportValidation {
	return &NullableLiveImportValidation{value: val, isSet: true}
}

func (v NullableLiveImportValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveImportValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
