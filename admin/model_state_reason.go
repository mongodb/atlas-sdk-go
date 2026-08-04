// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StateReason State reason of the Job. This is set when the job state is \"Failed\".
type StateReason struct {
	// Error code relating to state.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Message describing error or state.
	Message *string `json:"message,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StateReason) MarshalJSON() ([]byte, error) {
	type noMethod StateReason
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStateReason instantiates a new StateReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateReason() *StateReason {
	this := StateReason{}
	return &this
}

// NewStateReasonWithDefaults instantiates a new StateReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateReasonWithDefaults() *StateReason {
	this := StateReason{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise
func (o *StateReason) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateReason) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}

	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *StateReason) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *StateReason) SetErrorCode(v string) {
	o.ErrorCode = &v
	o.NullFields = removeNullField(o.NullFields, "ErrorCode")
}

// SetErrorCodeNil sets ErrorCode to an explicit JSON null when marshaled.
func (o *StateReason) SetErrorCodeNil() {
	o.ErrorCode = nil
	o.NullFields = addNullField(o.NullFields, "ErrorCode")
}

// GetMessage returns the Message field value if set, zero value otherwise
func (o *StateReason) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateReason) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}

	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *StateReason) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *StateReason) SetMessage(v string) {
	o.Message = &v
	o.NullFields = removeNullField(o.NullFields, "Message")
}

// SetMessageNil sets Message to an explicit JSON null when marshaled.
func (o *StateReason) SetMessageNil() {
	o.Message = nil
	o.NullFields = addNullField(o.NullFields, "Message")
}
