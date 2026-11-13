// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// OperationResponse Status of a long-running operation.
type OperationResponse struct {
	// Creation time in ISO 8601 timestamp format in UTC.
	// Read only field.
	CreatedAt time.Time `json:"createdAt"`
	// Name of a custom operation. Required for CUSTOM operations.
	// Read only field.
	CustomMethod *string         `json:"customMethod,omitempty"`
	Error        *OperationError `json:"error,omitempty"`
	// Estimated completion time in ISO 8601 timestamp format in UTC.
	// Read only field.
	EstimatedCompletionTime *time.Time `json:"estimatedCompletionTime,omitempty"`
	// Expiration time in ISO 8601 timestamp format in UTC.
	// Read only field.
	ExpiresAt time.Time `json:"expiresAt"`
	// Unique identifier of the operation.
	// Read only field.
	OperationId string `json:"operationId"`
	// Type of mutation performed by the operation.
	// Read only field.
	OperationType string             `json:"operationType"`
	Progress      *OperationProgress `json:"progress,omitempty"`
	// URI of the affected resource. Required when status is SUCCEEDED.
	// Read only field.
	ResultHref *string `json:"resultHref,omitempty"`
	// Recommended polling delay. Required while the operation is active.
	// Read only field.
	RetryAfterSeconds *int `json:"retryAfterSeconds,omitempty"`
	// State of the operation.
	// Read only field.
	Status string `json:"status"`
	// Human-readable operation status.
	// Read only field.
	StatusMessage *string `json:"statusMessage,omitempty"`
	// Last update time in ISO 8601 timestamp format in UTC.
	// Read only field.
	UpdatedAt time.Time `json:"updatedAt"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OperationResponse) MarshalJSON() ([]byte, error) {
	type noMethod OperationResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOperationResponse instantiates a new OperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponse(createdAt time.Time, expiresAt time.Time, operationId string, operationType string, status string, updatedAt time.Time) *OperationResponse {
	this := OperationResponse{}
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.OperationId = operationId
	this.OperationType = operationType
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewOperationResponseWithDefaults instantiates a new OperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseWithDefaults() *OperationResponse {
	this := OperationResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *OperationResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OperationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCustomMethod returns the CustomMethod field value if set, zero value otherwise
func (o *OperationResponse) GetCustomMethod() string {
	if o == nil || IsNil(o.CustomMethod) {
		var ret string
		return ret
	}
	return *o.CustomMethod
}

// GetCustomMethodOk returns a tuple with the CustomMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetCustomMethodOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMethod) {
		return nil, false
	}

	return o.CustomMethod, true
}

// HasCustomMethod returns a boolean if a field has been set.
func (o *OperationResponse) HasCustomMethod() bool {
	if o != nil && !IsNil(o.CustomMethod) {
		return true
	}

	return false
}

// SetCustomMethod gets a reference to the given string and assigns it to the CustomMethod field.
func (o *OperationResponse) SetCustomMethod(v string) {
	o.CustomMethod = &v
	o.NullFields = removeNullField(o.NullFields, "CustomMethod")
}

// SetCustomMethodNil sets CustomMethod to an explicit JSON null when marshaled.
func (o *OperationResponse) SetCustomMethodNil() {
	o.CustomMethod = nil
	o.NullFields = addNullField(o.NullFields, "CustomMethod")
}

// GetError returns the Error field value if set, zero value otherwise
func (o *OperationResponse) GetError() OperationError {
	if o == nil || IsNil(o.Error) {
		var ret OperationError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetErrorOk() (*OperationError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}

	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OperationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given OperationError and assigns it to the Error field.
func (o *OperationResponse) SetError(v OperationError) {
	o.Error = &v
	o.NullFields = removeNullField(o.NullFields, "Error")
}

// SetErrorNil sets Error to an explicit JSON null when marshaled.
func (o *OperationResponse) SetErrorNil() {
	o.Error = nil
	o.NullFields = addNullField(o.NullFields, "Error")
}

// GetEstimatedCompletionTime returns the EstimatedCompletionTime field value if set, zero value otherwise
func (o *OperationResponse) GetEstimatedCompletionTime() time.Time {
	if o == nil || IsNil(o.EstimatedCompletionTime) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedCompletionTime
}

// GetEstimatedCompletionTimeOk returns a tuple with the EstimatedCompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetEstimatedCompletionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EstimatedCompletionTime) {
		return nil, false
	}

	return o.EstimatedCompletionTime, true
}

// HasEstimatedCompletionTime returns a boolean if a field has been set.
func (o *OperationResponse) HasEstimatedCompletionTime() bool {
	if o != nil && !IsNil(o.EstimatedCompletionTime) {
		return true
	}

	return false
}

// SetEstimatedCompletionTime gets a reference to the given time.Time and assigns it to the EstimatedCompletionTime field.
func (o *OperationResponse) SetEstimatedCompletionTime(v time.Time) {
	o.EstimatedCompletionTime = &v
	o.NullFields = removeNullField(o.NullFields, "EstimatedCompletionTime")
}

// SetEstimatedCompletionTimeNil sets EstimatedCompletionTime to an explicit JSON null when marshaled.
func (o *OperationResponse) SetEstimatedCompletionTimeNil() {
	o.EstimatedCompletionTime = nil
	o.NullFields = addNullField(o.NullFields, "EstimatedCompletionTime")
}

// GetExpiresAt returns the ExpiresAt field value
func (o *OperationResponse) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *OperationResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetOperationId returns the OperationId field value
func (o *OperationResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *OperationResponse) SetOperationId(v string) {
	o.OperationId = v
}

// GetOperationType returns the OperationType field value
func (o *OperationResponse) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *OperationResponse) SetOperationType(v string) {
	o.OperationType = v
}

// GetProgress returns the Progress field value if set, zero value otherwise
func (o *OperationResponse) GetProgress() OperationProgress {
	if o == nil || IsNil(o.Progress) {
		var ret OperationProgress
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetProgressOk() (*OperationProgress, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}

	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *OperationResponse) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given OperationProgress and assigns it to the Progress field.
func (o *OperationResponse) SetProgress(v OperationProgress) {
	o.Progress = &v
	o.NullFields = removeNullField(o.NullFields, "Progress")
}

// SetProgressNil sets Progress to an explicit JSON null when marshaled.
func (o *OperationResponse) SetProgressNil() {
	o.Progress = nil
	o.NullFields = addNullField(o.NullFields, "Progress")
}

// GetResultHref returns the ResultHref field value if set, zero value otherwise
func (o *OperationResponse) GetResultHref() string {
	if o == nil || IsNil(o.ResultHref) {
		var ret string
		return ret
	}
	return *o.ResultHref
}

// GetResultHrefOk returns a tuple with the ResultHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetResultHrefOk() (*string, bool) {
	if o == nil || IsNil(o.ResultHref) {
		return nil, false
	}

	return o.ResultHref, true
}

// HasResultHref returns a boolean if a field has been set.
func (o *OperationResponse) HasResultHref() bool {
	if o != nil && !IsNil(o.ResultHref) {
		return true
	}

	return false
}

// SetResultHref gets a reference to the given string and assigns it to the ResultHref field.
func (o *OperationResponse) SetResultHref(v string) {
	o.ResultHref = &v
	o.NullFields = removeNullField(o.NullFields, "ResultHref")
}

// SetResultHrefNil sets ResultHref to an explicit JSON null when marshaled.
func (o *OperationResponse) SetResultHrefNil() {
	o.ResultHref = nil
	o.NullFields = addNullField(o.NullFields, "ResultHref")
}

// GetRetryAfterSeconds returns the RetryAfterSeconds field value if set, zero value otherwise
func (o *OperationResponse) GetRetryAfterSeconds() int {
	if o == nil || IsNil(o.RetryAfterSeconds) {
		var ret int
		return ret
	}
	return *o.RetryAfterSeconds
}

// GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetRetryAfterSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.RetryAfterSeconds) {
		return nil, false
	}

	return o.RetryAfterSeconds, true
}

// HasRetryAfterSeconds returns a boolean if a field has been set.
func (o *OperationResponse) HasRetryAfterSeconds() bool {
	if o != nil && !IsNil(o.RetryAfterSeconds) {
		return true
	}

	return false
}

// SetRetryAfterSeconds gets a reference to the given int and assigns it to the RetryAfterSeconds field.
func (o *OperationResponse) SetRetryAfterSeconds(v int) {
	o.RetryAfterSeconds = &v
	o.NullFields = removeNullField(o.NullFields, "RetryAfterSeconds")
}

// SetRetryAfterSecondsNil sets RetryAfterSeconds to an explicit JSON null when marshaled.
func (o *OperationResponse) SetRetryAfterSecondsNil() {
	o.RetryAfterSeconds = nil
	o.NullFields = addNullField(o.NullFields, "RetryAfterSeconds")
}

// GetStatus returns the Status field value
func (o *OperationResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OperationResponse) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise
func (o *OperationResponse) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}

	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *OperationResponse) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *OperationResponse) SetStatusMessage(v string) {
	o.StatusMessage = &v
	o.NullFields = removeNullField(o.NullFields, "StatusMessage")
}

// SetStatusMessageNil sets StatusMessage to an explicit JSON null when marshaled.
func (o *OperationResponse) SetStatusMessageNil() {
	o.StatusMessage = nil
	o.NullFields = addNullField(o.NullFields, "StatusMessage")
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OperationResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OperationResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}
