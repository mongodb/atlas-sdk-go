// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SchemaAdvisorResponse struct for SchemaAdvisorResponse
type SchemaAdvisorResponse struct {
	// List that contains the documents with information about the schema advice that Performance Advisor suggests.
	// Read only field.
	Recommendations *[]SchemaAdvisorItemRecommendation `json:"recommendations,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *SchemaAdvisorResponse) MarshalJSON() ([]byte, error) {
	type noMethod SchemaAdvisorResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewSchemaAdvisorResponse instantiates a new SchemaAdvisorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaAdvisorResponse() *SchemaAdvisorResponse {
	this := SchemaAdvisorResponse{}
	return &this
}

// NewSchemaAdvisorResponseWithDefaults instantiates a new SchemaAdvisorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaAdvisorResponseWithDefaults() *SchemaAdvisorResponse {
	this := SchemaAdvisorResponse{}
	return &this
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise
func (o *SchemaAdvisorResponse) GetRecommendations() []SchemaAdvisorItemRecommendation {
	if o == nil || IsNil(o.Recommendations) {
		var ret []SchemaAdvisorItemRecommendation
		return ret
	}
	return *o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAdvisorResponse) GetRecommendationsOk() (*[]SchemaAdvisorItemRecommendation, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}

	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *SchemaAdvisorResponse) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []SchemaAdvisorItemRecommendation and assigns it to the Recommendations field.
func (o *SchemaAdvisorResponse) SetRecommendations(v []SchemaAdvisorItemRecommendation) {
	o.Recommendations = &v
	o.NullFields = removeNullField(o.NullFields, "Recommendations")
}

// SetRecommendationsNil sets Recommendations to an explicit JSON null when marshaled.
func (o *SchemaAdvisorResponse) SetRecommendationsNil() {
	o.Recommendations = nil
	o.NullFields = addNullField(o.NullFields, "Recommendations")
}
