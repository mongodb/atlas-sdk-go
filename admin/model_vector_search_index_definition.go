// Code based on the AtlasAPI V2 OpenAPI file

package admin

// VectorSearchIndexDefinition The vector search index definition set by the user.
type VectorSearchIndexDefinition struct {
	// Settings that configure the fields, one per object, to index. You must define at least one \"vector\" type field. You can optionally define \"filter\" type fields also.
	Fields *[]interface{} `json:"fields,omitempty"`
}

// NewVectorSearchIndexDefinition instantiates a new VectorSearchIndexDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVectorSearchIndexDefinition() *VectorSearchIndexDefinition {
	this := VectorSearchIndexDefinition{}
	return &this
}

// NewVectorSearchIndexDefinitionWithDefaults instantiates a new VectorSearchIndexDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVectorSearchIndexDefinitionWithDefaults() *VectorSearchIndexDefinition {
	this := VectorSearchIndexDefinition{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise
func (o *VectorSearchIndexDefinition) GetFields() []interface{} {
	if o == nil || IsNil(o.Fields) {
		var ret []interface{}
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VectorSearchIndexDefinition) GetFieldsOk() (*[]interface{}, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}

	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *VectorSearchIndexDefinition) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []interface{} and assigns it to the Fields field.
func (o *VectorSearchIndexDefinition) SetFields(v []interface{}) {
	o.Fields = &v
}
