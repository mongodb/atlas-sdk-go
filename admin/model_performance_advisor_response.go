// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PerformanceAdvisorResponse struct for PerformanceAdvisorResponse
type PerformanceAdvisorResponse struct {
	Content struct {
		// List of query predicates, sorts, and projections that the Performance Advisor suggests.
		// Read only field.
		Shapes *[]PerformanceAdvisorShape `json:"shapes,omitempty"`
		// List that contains the documents with information about the indexes that the Performance Advisor suggests.
		// Read only field.
		SuggestedIndexes *[]PerformanceAdvisorIndex `json:"suggestedIndexes,omitempty"`
	} `json:"content"`
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
	if o == nil || IsNil(o.Content.Shapes) {
		var ret []PerformanceAdvisorShape
		return ret
	}
	return *o.Content.Shapes
}

// GetShapesOk returns a tuple with the Shapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorResponse) GetShapesOk() (*[]PerformanceAdvisorShape, bool) {
	if o == nil || IsNil(o.Content.Shapes) {
		return nil, false
	}

	return o.Content.Shapes, true
}

// HasShapes returns a boolean if a field has been set.
func (o *PerformanceAdvisorResponse) HasShapes() bool {
	if o != nil && !IsNil(o.Content) && !IsNil(o.Content.Shapes) {
		return true
	}

	return false
}

// SetShapes gets a reference to the given []PerformanceAdvisorShape and assigns it to the Shapes field.
func (o *PerformanceAdvisorResponse) SetShapes(v []PerformanceAdvisorShape) {
	o.Content.Shapes = &v
}

// GetSuggestedIndexes returns the SuggestedIndexes field value if set, zero value otherwise
func (o *PerformanceAdvisorResponse) GetSuggestedIndexes() []PerformanceAdvisorIndex {
	if o == nil || IsNil(o.Content.SuggestedIndexes) {
		var ret []PerformanceAdvisorIndex
		return ret
	}
	return *o.Content.SuggestedIndexes
}

// GetSuggestedIndexesOk returns a tuple with the SuggestedIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceAdvisorResponse) GetSuggestedIndexesOk() (*[]PerformanceAdvisorIndex, bool) {
	if o == nil || IsNil(o.Content.SuggestedIndexes) {
		return nil, false
	}

	return o.Content.SuggestedIndexes, true
}

// HasSuggestedIndexes returns a boolean if a field has been set.
func (o *PerformanceAdvisorResponse) HasSuggestedIndexes() bool {
	if o != nil && !IsNil(o.Content) && !IsNil(o.Content.SuggestedIndexes) {
		return true
	}

	return false
}

// SetSuggestedIndexes gets a reference to the given []PerformanceAdvisorIndex and assigns it to the SuggestedIndexes field.
func (o *PerformanceAdvisorResponse) SetSuggestedIndexes(v []PerformanceAdvisorIndex) {
	o.Content.SuggestedIndexes = &v
}
