// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OperationProgress Progress made by an active operation.
type OperationProgress struct {
	// Number of completed units.
	Completed *int `json:"completed,omitempty"`
	// Total number of units.
	Total *int `json:"total,omitempty"`
	// Unit measured by completed and total.
	Unit *string `json:"unit,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OperationProgress) MarshalJSON() ([]byte, error) {
	type noMethod OperationProgress
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOperationProgress instantiates a new OperationProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationProgress() *OperationProgress {
	this := OperationProgress{}
	return &this
}

// NewOperationProgressWithDefaults instantiates a new OperationProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationProgressWithDefaults() *OperationProgress {
	this := OperationProgress{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise
func (o *OperationProgress) GetCompleted() int {
	if o == nil || IsNil(o.Completed) {
		var ret int
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationProgress) GetCompletedOk() (*int, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}

	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *OperationProgress) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given int and assigns it to the Completed field.
func (o *OperationProgress) SetCompleted(v int) {
	o.Completed = &v
	o.NullFields = removeNullField(o.NullFields, "Completed")
}

// SetCompletedNil sets Completed to an explicit JSON null when marshaled.
func (o *OperationProgress) SetCompletedNil() {
	o.Completed = nil
	o.NullFields = addNullField(o.NullFields, "Completed")
}

// GetTotal returns the Total field value if set, zero value otherwise
func (o *OperationProgress) GetTotal() int {
	if o == nil || IsNil(o.Total) {
		var ret int
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationProgress) GetTotalOk() (*int, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}

	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OperationProgress) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int and assigns it to the Total field.
func (o *OperationProgress) SetTotal(v int) {
	o.Total = &v
	o.NullFields = removeNullField(o.NullFields, "Total")
}

// SetTotalNil sets Total to an explicit JSON null when marshaled.
func (o *OperationProgress) SetTotalNil() {
	o.Total = nil
	o.NullFields = addNullField(o.NullFields, "Total")
}

// GetUnit returns the Unit field value if set, zero value otherwise
func (o *OperationProgress) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationProgress) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}

	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *OperationProgress) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *OperationProgress) SetUnit(v string) {
	o.Unit = &v
	o.NullFields = removeNullField(o.NullFields, "Unit")
}

// SetUnitNil sets Unit to an explicit JSON null when marshaled.
func (o *OperationProgress) SetUnitNil() {
	o.Unit = nil
	o.NullFields = addNullField(o.NullFields, "Unit")
}
