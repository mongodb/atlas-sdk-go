// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PerformanceAdvisorResponse Response that contains Performance Advisor suggested indexes and query shapes.
type PerformanceAdvisorResponse struct {
	// List of query predicates, sorts, and projections that the Performance Advisor suggests.
	// Read only field.
	Shapes *[]PerformanceAdvisorShape `json:"shapes,omitempty"`
	// List that contains the documents with information about the indexes that the Performance Advisor suggests.
	// Read only field.
	SuggestedIndexes *[]PerformanceAdvisorIndex `json:"suggestedIndexes,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PerformanceAdvisorResponse) MarshalJSON() ([]byte, error) {
	type noMethod PerformanceAdvisorResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPerformanceAdvisorResponse instantiates a new PerformanceAdvisorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceAdvisorResponse() *PerformanceAdvisorResponse {
	this := PerformanceAdvisorResponse{}
	return &this
}

// NewPerformanceAdvisorResponseWithDefaults instantiates a new PerformanceAdvisorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceAdvisorResponseWithDefaults() *PerformanceAdvisorResponse {
	this := PerformanceAdvisorResponse{}
	return &this
}

// GetShapes returns the Shapes field value if set, zero value otherwise
func (o *PerformanceAdvisorResponse) GetShapes() []PerformanceAdvisorShape {
	if o == nil || IsNil(o.Shapes) {
		var ret []PerformanceAdvisorShape
		return ret
	}
	return *o.Shapes
}

// GetShapesOk returns a tuple with the Shapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorResponse) GetShapesOk() (*[]PerformanceAdvisorShape, bool) {
	if o == nil || IsNil(o.Shapes) {
		return nil, false
	}

	return o.Shapes, true
}

// HasShapes returns a boolean if a field has been set.
func (o *PerformanceAdvisorResponse) HasShapes() bool {
	if o != nil && !IsNil(o.Shapes) {
		return true
	}

	return false
}

// SetShapes gets a reference to the given []PerformanceAdvisorShape and assigns it to the Shapes field.
func (o *PerformanceAdvisorResponse) SetShapes(v []PerformanceAdvisorShape) {
	o.Shapes = &v
	o.NullFields = removeNullField(o.NullFields, "Shapes")
}

// SetShapesNil sets Shapes to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorResponse) SetShapesNil() {
	o.Shapes = nil
	o.NullFields = addNullField(o.NullFields, "Shapes")
}

// GetSuggestedIndexes returns the SuggestedIndexes field value if set, zero value otherwise
func (o *PerformanceAdvisorResponse) GetSuggestedIndexes() []PerformanceAdvisorIndex {
	if o == nil || IsNil(o.SuggestedIndexes) {
		var ret []PerformanceAdvisorIndex
		return ret
	}
	return *o.SuggestedIndexes
}

// GetSuggestedIndexesOk returns a tuple with the SuggestedIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorResponse) GetSuggestedIndexesOk() (*[]PerformanceAdvisorIndex, bool) {
	if o == nil || IsNil(o.SuggestedIndexes) {
		return nil, false
	}

	return o.SuggestedIndexes, true
}

// HasSuggestedIndexes returns a boolean if a field has been set.
func (o *PerformanceAdvisorResponse) HasSuggestedIndexes() bool {
	if o != nil && !IsNil(o.SuggestedIndexes) {
		return true
	}

	return false
}

// SetSuggestedIndexes gets a reference to the given []PerformanceAdvisorIndex and assigns it to the SuggestedIndexes field.
func (o *PerformanceAdvisorResponse) SetSuggestedIndexes(v []PerformanceAdvisorIndex) {
	o.SuggestedIndexes = &v
	o.NullFields = removeNullField(o.NullFields, "SuggestedIndexes")
}

// SetSuggestedIndexesNil sets SuggestedIndexes to an explicit JSON null when marshaled.
func (o *PerformanceAdvisorResponse) SetSuggestedIndexesNil() {
	o.SuggestedIndexes = nil
	o.NullFields = addNullField(o.NullFields, "SuggestedIndexes")
}
