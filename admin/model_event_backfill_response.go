// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// EventBackfillResponse State of a one-time backfill of historical events to an external log sink.
type EventBackfillResponse struct {
	// Date after which events are covered by the live stream rather than the backfill, in ISO 8601 UTC format.
	// Read only field.
	BackfillCutoffDate *time.Time `json:"backfillCutoffDate,omitempty"`
	// Earliest event created date included in the backfill, in ISO 8601 UTC format. Null when the entire available history was backfilled.
	EarliestCreatedDate *time.Time `json:"earliestCreatedDate,omitempty"`
	// Earliest failed event date, recorded when a backfill fails so a future retry can be bounded, in ISO 8601 UTC format.
	// Read only field.
	EarliestFailedDate *time.Time `json:"earliestFailedDate,omitempty"`
	// Reason the backfill failed: historical events could not be read, the external platform rejected events, or some events were never confirmed as delivered. Null unless the backfill failed.
	// Read only field.
	FailureReason *string `json:"failureReason,omitempty"`
	// Status of the backfill.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *EventBackfillResponse) MarshalJSON() ([]byte, error) {
	type noMethod EventBackfillResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewEventBackfillResponse instantiates a new EventBackfillResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBackfillResponse() *EventBackfillResponse {
	this := EventBackfillResponse{}
	return &this
}

// NewEventBackfillResponseWithDefaults instantiates a new EventBackfillResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBackfillResponseWithDefaults() *EventBackfillResponse {
	this := EventBackfillResponse{}
	return &this
}

// GetBackfillCutoffDate returns the BackfillCutoffDate field value if set, zero value otherwise
func (o *EventBackfillResponse) GetBackfillCutoffDate() time.Time {
	if o == nil || IsNil(o.BackfillCutoffDate) {
		var ret time.Time
		return ret
	}
	return *o.BackfillCutoffDate
}

// GetBackfillCutoffDateOk returns a tuple with the BackfillCutoffDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBackfillResponse) GetBackfillCutoffDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BackfillCutoffDate) {
		return nil, false
	}

	return o.BackfillCutoffDate, true
}

// HasBackfillCutoffDate returns a boolean if a field has been set.
func (o *EventBackfillResponse) HasBackfillCutoffDate() bool {
	if o != nil && !IsNil(o.BackfillCutoffDate) {
		return true
	}

	return false
}

// SetBackfillCutoffDate gets a reference to the given time.Time and assigns it to the BackfillCutoffDate field.
func (o *EventBackfillResponse) SetBackfillCutoffDate(v time.Time) {
	o.BackfillCutoffDate = &v
	o.NullFields = removeNullField(o.NullFields, "BackfillCutoffDate")
}

// SetBackfillCutoffDateNil sets BackfillCutoffDate to an explicit JSON null when marshaled.
func (o *EventBackfillResponse) SetBackfillCutoffDateNil() {
	o.BackfillCutoffDate = nil
	o.NullFields = addNullField(o.NullFields, "BackfillCutoffDate")
}

// GetEarliestCreatedDate returns the EarliestCreatedDate field value if set, zero value otherwise
func (o *EventBackfillResponse) GetEarliestCreatedDate() time.Time {
	if o == nil || IsNil(o.EarliestCreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.EarliestCreatedDate
}

// GetEarliestCreatedDateOk returns a tuple with the EarliestCreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBackfillResponse) GetEarliestCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestCreatedDate) {
		return nil, false
	}

	return o.EarliestCreatedDate, true
}

// HasEarliestCreatedDate returns a boolean if a field has been set.
func (o *EventBackfillResponse) HasEarliestCreatedDate() bool {
	if o != nil && !IsNil(o.EarliestCreatedDate) {
		return true
	}

	return false
}

// SetEarliestCreatedDate gets a reference to the given time.Time and assigns it to the EarliestCreatedDate field.
func (o *EventBackfillResponse) SetEarliestCreatedDate(v time.Time) {
	o.EarliestCreatedDate = &v
	o.NullFields = removeNullField(o.NullFields, "EarliestCreatedDate")
}

// SetEarliestCreatedDateNil sets EarliestCreatedDate to an explicit JSON null when marshaled.
func (o *EventBackfillResponse) SetEarliestCreatedDateNil() {
	o.EarliestCreatedDate = nil
	o.NullFields = addNullField(o.NullFields, "EarliestCreatedDate")
}

// GetEarliestFailedDate returns the EarliestFailedDate field value if set, zero value otherwise
func (o *EventBackfillResponse) GetEarliestFailedDate() time.Time {
	if o == nil || IsNil(o.EarliestFailedDate) {
		var ret time.Time
		return ret
	}
	return *o.EarliestFailedDate
}

// GetEarliestFailedDateOk returns a tuple with the EarliestFailedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBackfillResponse) GetEarliestFailedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestFailedDate) {
		return nil, false
	}

	return o.EarliestFailedDate, true
}

// HasEarliestFailedDate returns a boolean if a field has been set.
func (o *EventBackfillResponse) HasEarliestFailedDate() bool {
	if o != nil && !IsNil(o.EarliestFailedDate) {
		return true
	}

	return false
}

// SetEarliestFailedDate gets a reference to the given time.Time and assigns it to the EarliestFailedDate field.
func (o *EventBackfillResponse) SetEarliestFailedDate(v time.Time) {
	o.EarliestFailedDate = &v
	o.NullFields = removeNullField(o.NullFields, "EarliestFailedDate")
}

// SetEarliestFailedDateNil sets EarliestFailedDate to an explicit JSON null when marshaled.
func (o *EventBackfillResponse) SetEarliestFailedDateNil() {
	o.EarliestFailedDate = nil
	o.NullFields = addNullField(o.NullFields, "EarliestFailedDate")
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise
func (o *EventBackfillResponse) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBackfillResponse) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}

	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *EventBackfillResponse) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *EventBackfillResponse) SetFailureReason(v string) {
	o.FailureReason = &v
	o.NullFields = removeNullField(o.NullFields, "FailureReason")
}

// SetFailureReasonNil sets FailureReason to an explicit JSON null when marshaled.
func (o *EventBackfillResponse) SetFailureReasonNil() {
	o.FailureReason = nil
	o.NullFields = addNullField(o.NullFields, "FailureReason")
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *EventBackfillResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBackfillResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventBackfillResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventBackfillResponse) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *EventBackfillResponse) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}
