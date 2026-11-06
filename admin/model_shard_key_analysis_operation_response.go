// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// ShardKeyAnalysisOperationResponse Progress of one shard key analysis. Poll this resource until `status` reaches a terminal value, then follow `resultHref` to read the analysis.
type ShardKeyAnalysisOperationResponse struct {
	// Date and time when MongoDB Cloud accepted this operation (ISO 8601 format in UTC).
	// Read only field.
	CreatedAt time.Time              `json:"createdAt"`
	Error     *ShardKeyAnalysisError `json:"error,omitempty"`
	// Date and time when MongoDB Cloud stops returning this operation (ISO 8601 format in UTC). Reads after this time return an HTTP 404 status. This value derives from `updatedAt`, so it advances with each state change while the operation is running and settles only once the operation is terminal.
	// Read only field.
	ExpiresAt time.Time `json:"expiresAt"`
	// Unique 24-hexadecimal digit string that identifies this operation. When the operation succeeds, the analysis it created carries the same identifier.
	// Read only field.
	OperationId string `json:"operationId"`
	// Kind of change the operation makes. Submitting a shard key analysis creates an analysis resource, so the operation type is always `CREATE`.
	// Read only field.
	OperationType string `json:"operationType"`
	// URI of the analysis this operation created. MongoDB Cloud returns this parameter only when `status` is `SUCCEEDED`.
	// Read only field.
	ResultHref *string `json:"resultHref,omitempty"`
	// Number of seconds to wait before polling this operation again. MongoDB Cloud returns this parameter only while `status` is not terminal.
	// Read only field.
	RetryAfterSeconds *int `json:"retryAfterSeconds,omitempty"`
	// State of the shard key analysis operation. `PENDING` and `IN_PROGRESS` are non-terminal; `SUCCEEDED`, `FAILED`, `CANCELED` and `SUPERSEDED` are terminal. MongoDB Cloud currently reports only `PENDING`, `SUCCEEDED` and `FAILED`. The remaining values are reserved and may be reported in the future without a new resource version.
	// Read only field.
	Status string `json:"status"`
	// Date and time when this operation last changed state (ISO 8601 format in UTC). Once the operation is terminal, this is when it finished.
	// Read only field.
	UpdatedAt time.Time `json:"updatedAt"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardKeyAnalysisOperationResponse) MarshalJSON() ([]byte, error) {
	type noMethod ShardKeyAnalysisOperationResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardKeyAnalysisOperationResponse instantiates a new ShardKeyAnalysisOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardKeyAnalysisOperationResponse(createdAt time.Time, expiresAt time.Time, operationId string, operationType string, status string, updatedAt time.Time) *ShardKeyAnalysisOperationResponse {
	this := ShardKeyAnalysisOperationResponse{}
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.OperationId = operationId
	this.OperationType = operationType
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewShardKeyAnalysisOperationResponseWithDefaults instantiates a new ShardKeyAnalysisOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardKeyAnalysisOperationResponseWithDefaults() *ShardKeyAnalysisOperationResponse {
	this := ShardKeyAnalysisOperationResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ShardKeyAnalysisOperationResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ShardKeyAnalysisOperationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetError returns the Error field value if set, zero value otherwise
func (o *ShardKeyAnalysisOperationResponse) GetError() ShardKeyAnalysisError {
	if o == nil || IsNil(o.Error) {
		var ret ShardKeyAnalysisError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetErrorOk() (*ShardKeyAnalysisError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}

	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ShardKeyAnalysisOperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ShardKeyAnalysisError and assigns it to the Error field.
func (o *ShardKeyAnalysisOperationResponse) SetError(v ShardKeyAnalysisError) {
	o.Error = &v
	o.NullFields = removeNullField(o.NullFields, "Error")
}

// SetErrorNil sets Error to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisOperationResponse) SetErrorNil() {
	o.Error = nil
	o.NullFields = addNullField(o.NullFields, "Error")
}

// GetExpiresAt returns the ExpiresAt field value
func (o *ShardKeyAnalysisOperationResponse) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *ShardKeyAnalysisOperationResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetOperationId returns the OperationId field value
func (o *ShardKeyAnalysisOperationResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ShardKeyAnalysisOperationResponse) SetOperationId(v string) {
	o.OperationId = v
}

// GetOperationType returns the OperationType field value
func (o *ShardKeyAnalysisOperationResponse) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *ShardKeyAnalysisOperationResponse) SetOperationType(v string) {
	o.OperationType = v
}

// GetResultHref returns the ResultHref field value if set, zero value otherwise
func (o *ShardKeyAnalysisOperationResponse) GetResultHref() string {
	if o == nil || IsNil(o.ResultHref) {
		var ret string
		return ret
	}
	return *o.ResultHref
}

// GetResultHrefOk returns a tuple with the ResultHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetResultHrefOk() (*string, bool) {
	if o == nil || IsNil(o.ResultHref) {
		return nil, false
	}

	return o.ResultHref, true
}

// HasResultHref returns a boolean if a field has been set.
func (o *ShardKeyAnalysisOperationResponse) HasResultHref() bool {
	if o != nil && !IsNil(o.ResultHref) {
		return true
	}

	return false
}

// SetResultHref gets a reference to the given string and assigns it to the ResultHref field.
func (o *ShardKeyAnalysisOperationResponse) SetResultHref(v string) {
	o.ResultHref = &v
	o.NullFields = removeNullField(o.NullFields, "ResultHref")
}

// SetResultHrefNil sets ResultHref to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisOperationResponse) SetResultHrefNil() {
	o.ResultHref = nil
	o.NullFields = addNullField(o.NullFields, "ResultHref")
}

// GetRetryAfterSeconds returns the RetryAfterSeconds field value if set, zero value otherwise
func (o *ShardKeyAnalysisOperationResponse) GetRetryAfterSeconds() int {
	if o == nil || IsNil(o.RetryAfterSeconds) {
		var ret int
		return ret
	}
	return *o.RetryAfterSeconds
}

// GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetRetryAfterSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.RetryAfterSeconds) {
		return nil, false
	}

	return o.RetryAfterSeconds, true
}

// HasRetryAfterSeconds returns a boolean if a field has been set.
func (o *ShardKeyAnalysisOperationResponse) HasRetryAfterSeconds() bool {
	if o != nil && !IsNil(o.RetryAfterSeconds) {
		return true
	}

	return false
}

// SetRetryAfterSeconds gets a reference to the given int and assigns it to the RetryAfterSeconds field.
func (o *ShardKeyAnalysisOperationResponse) SetRetryAfterSeconds(v int) {
	o.RetryAfterSeconds = &v
	o.NullFields = removeNullField(o.NullFields, "RetryAfterSeconds")
}

// SetRetryAfterSecondsNil sets RetryAfterSeconds to an explicit JSON null when marshaled.
func (o *ShardKeyAnalysisOperationResponse) SetRetryAfterSecondsNil() {
	o.RetryAfterSeconds = nil
	o.NullFields = addNullField(o.NullFields, "RetryAfterSeconds")
}

// GetStatus returns the Status field value
func (o *ShardKeyAnalysisOperationResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ShardKeyAnalysisOperationResponse) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ShardKeyAnalysisOperationResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ShardKeyAnalysisOperationResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ShardKeyAnalysisOperationResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}
