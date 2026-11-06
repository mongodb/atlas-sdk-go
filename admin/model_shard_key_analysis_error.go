// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardKeyAnalysisError Details of the failure that ended the shard key analysis operation.
type ShardKeyAnalysisError struct {
	// Reason that the shard key analysis failed. MongoDB Cloud may report additional values in the future without a new resource version, so treat a value this description does not list as a generic failure.
	// Read only field.
	Code string `json:"code"`
	// Human-readable explanation of the failure. This message describes the classified failure and is not the raw error text from the cluster.
	// Read only field.
	Message string `json:"message"`
	// How to retry a failed shard key analysis. `BACKOFF` indicates that resubmitting the same request after a delay may succeed. `NONE` indicates that the failure describes the collection or the candidate shard key and that resubmitting the same request will fail the same way.
	// Read only field.
	RetryStrategy *string `json:"retryStrategy,omitempty"`
	// Flag that indicates whether submitting the same request again may succeed. This flag is `false` when the failure describes the collection or the candidate shard key, because the same request then fails the same way.
	// Read only field.
	Retryable bool `json:"retryable"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardKeyAnalysisError) MarshalJSON() ([]byte, error) {
	type noMethod ShardKeyAnalysisError
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardKeyAnalysisError instantiates a new ShardKeyAnalysisError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeyAnalysisError(code string, message string, retryable bool) *ShardKeyAnalysisError {
	this := ShardKeyAnalysisError{}
	this.Code = code
	this.Message = message
	this.Retryable = retryable
	return &this
}

// NewShardKeyAnalysisErrorWithDefaults instantiates a new ShardKeyAnalysisError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeyAnalysisErrorWithDefaults() *ShardKeyAnalysisError {
	this := ShardKeyAnalysisError{}
	return &this
}

// GetCode returns the Code field value
func (o *ShardKeyAnalysisError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ShardKeyAnalysisError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ShardKeyAnalysisError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ShardKeyAnalysisError) SetMessage(v string) {
	o.Message = v
}

// GetRetryStrategy returns the RetryStrategy field value if set, zero value otherwise
func (o *ShardKeyAnalysisError) GetRetryStrategy() string {
	if o == nil || IsNil(o.RetryStrategy) {
		var ret string
		return ret
	}
	return *o.RetryStrategy
}

// GetRetryStrategyOk returns a tuple with the RetryStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisError) GetRetryStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.RetryStrategy) {
		return nil, false
	}

	return o.RetryStrategy, true
}

// HasRetryStrategy returns a boolean if a field has been set.
func (o *ShardKeyAnalysisError) HasRetryStrategy() bool {
	if o != nil && !IsNil(o.RetryStrategy) {
		return true
	}

	return false
}

// SetRetryStrategy gets a reference to the given string and assigns it to the RetryStrategy field.
func (o *ShardKeyAnalysisError) SetRetryStrategy(v string) {
	o.RetryStrategy = &v
	o.NullFields = removeNullField(o.NullFields, "RetryStrategy")
}

// SetRetryStrategyNil sets RetryStrategy to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisError) SetRetryStrategyNil() {
	o.RetryStrategy = nil
	o.NullFields = addNullField(o.NullFields, "RetryStrategy")
}

// GetRetryable returns the Retryable field value
func (o *ShardKeyAnalysisError) GetRetryable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisError) GetRetryableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retryable, true
}

// SetRetryable sets field value
func (o *ShardKeyAnalysisError) SetRetryable(v bool) {
	o.Retryable = v
}
