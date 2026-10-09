// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OperationError Details about a failed operation.
type OperationError struct {
	// Machine-readable error code.
	Code string `json:"code"`
	// Additional error details.
	Details any `json:"details,omitempty"`
	// Human-readable error message.
	Message string `json:"message"`
	// Recommended strategy for retrying the operation.
	RetryStrategy *string `json:"retryStrategy,omitempty"`
	// Flag that indicates whether the operation can be retried.
	Retryable bool `json:"retryable"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OperationError) MarshalJSON() ([]byte, error) {
	type noMethod OperationError
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOperationError instantiates a new OperationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationError(code string, message string, retryable bool) *OperationError {
	this := OperationError{}
	this.Code = code
	this.Message = message
	this.Retryable = retryable
	return &this
}

// NewOperationErrorWithDefaults instantiates a new OperationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationErrorWithDefaults() *OperationError {
	this := OperationError{}
	return &this
}

// GetCode returns the Code field value
func (o *OperationError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OperationError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OperationError) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value if set, zero value otherwise
func (o *OperationError) GetDetails() any {
	if o == nil || IsNil(o.Details) {
		var ret any
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationError) GetDetailsOk() (any, bool) {
	if o == nil || IsNil(o.Details) {
		var ret any
		return ret, false
	}

	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OperationError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given any and assigns it to the Details field.
func (o *OperationError) SetDetails(v any) {
	o.Details = v
	o.NullFields = removeNullField(o.NullFields, "Details")
}

// SetDetailsNil sets Details to an explicit JSON null when marshaled.
func (o *OperationError) SetDetailsNil() {
	o.Details = nil
	o.NullFields = addNullField(o.NullFields, "Details")
}

// GetMessage returns the Message field value
func (o *OperationError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OperationError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OperationError) SetMessage(v string) {
	o.Message = v
}

// GetRetryStrategy returns the RetryStrategy field value if set, zero value otherwise
func (o *OperationError) GetRetryStrategy() string {
	if o == nil || IsNil(o.RetryStrategy) {
		var ret string
		return ret
	}
	return *o.RetryStrategy
}

// GetRetryStrategyOk returns a tuple with the RetryStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationError) GetRetryStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.RetryStrategy) {
		return nil, false
	}

	return o.RetryStrategy, true
}

// HasRetryStrategy returns a boolean if a field has been set.
func (o *OperationError) HasRetryStrategy() bool {
	if o != nil && !IsNil(o.RetryStrategy) {
		return true
	}

	return false
}

// SetRetryStrategy gets a reference to the given string and assigns it to the RetryStrategy field.
func (o *OperationError) SetRetryStrategy(v string) {
	o.RetryStrategy = &v
	o.NullFields = removeNullField(o.NullFields, "RetryStrategy")
}

// SetRetryStrategyNil sets RetryStrategy to an explicit JSON null when marshaled.
func (o *OperationError) SetRetryStrategyNil() {
	o.RetryStrategy = nil
	o.NullFields = addNullField(o.NullFields, "RetryStrategy")
}

// GetRetryable returns the Retryable field value
func (o *OperationError) GetRetryable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value
// and a boolean to check if the value has been set.
func (o *OperationError) GetRetryableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retryable, true
}

// SetRetryable sets field value
func (o *OperationError) SetRetryable(v bool) {
	o.Retryable = v
}
