// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// EventBackfillRequest Request to backfill historical events to an external log sink.
type EventBackfillRequest struct {
	// Earliest event created date to include in the backfill, in ISO 8601 UTC format. When omitted, the entire available history is backfilled.
	EarliestCreatedDate *time.Time `json:"earliestCreatedDate,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o EventBackfillRequest) MarshalJSON() ([]byte, error) {
	type noMethod EventBackfillRequest
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewEventBackfillRequest instantiates a new EventBackfillRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBackfillRequest() *EventBackfillRequest {
	this := EventBackfillRequest{}
	return &this
}

// NewEventBackfillRequestWithDefaults instantiates a new EventBackfillRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBackfillRequestWithDefaults() *EventBackfillRequest {
	this := EventBackfillRequest{}
	return &this
}

// GetEarliestCreatedDate returns the EarliestCreatedDate field value if set, zero value otherwise
func (o *EventBackfillRequest) GetEarliestCreatedDate() time.Time {
	if o == nil || IsNil(o.EarliestCreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.EarliestCreatedDate
}

// GetEarliestCreatedDateOk returns a tuple with the EarliestCreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBackfillRequest) GetEarliestCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestCreatedDate) {
		return nil, false
	}

	return o.EarliestCreatedDate, true
}

// HasEarliestCreatedDate returns a boolean if a field has been set.
func (o *EventBackfillRequest) HasEarliestCreatedDate() bool {
	if o != nil && !IsNil(o.EarliestCreatedDate) {
		return true
	}

	return false
}

// SetEarliestCreatedDate gets a reference to the given time.Time and assigns it to the EarliestCreatedDate field.
func (o *EventBackfillRequest) SetEarliestCreatedDate(v time.Time) {
	o.EarliestCreatedDate = &v
	o.NullFields = removeNullField(o.NullFields, "EarliestCreatedDate")
}

// SetEarliestCreatedDateNil sets EarliestCreatedDate to an explicit JSON null when marshaled.
func (o *EventBackfillRequest) SetEarliestCreatedDateNil() {
	o.EarliestCreatedDate = nil
	o.NullFields = addNullField(o.NullFields, "EarliestCreatedDate")
}
