// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// SearchIndexCreateRequest struct for SearchIndexCreateRequest
type SearchIndexCreateRequest struct {
	// Label that identifies the collection to create an Atlas Search index in.
	CollectionName string `json:"collectionName"`
	// Label that identifies the database that contains the collection to create an Atlas Search index in.
	Database string `json:"database"`
	// Label that identifies this index. Within each namespace, names of all indexes in the namespace must be unique.
	Name string `json:"name"`
	// Type of the index. The default type is search.
	Type       *string                      `json:"type,omitempty"`
	Definition *VectorSearchIndexDefinition `json:"definition,omitempty"`
}

// NewSearchIndexCreateRequest instantiates a new SearchIndexCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexCreateRequest(collectionName string, database string, name string) *SearchIndexCreateRequest {
	this := SearchIndexCreateRequest{}
	this.CollectionName = collectionName
	this.Database = database
	this.Name = name
	return &this
}

// NewSearchIndexCreateRequestWithDefaults instantiates a new SearchIndexCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexCreateRequestWithDefaults() *SearchIndexCreateRequest {
	this := SearchIndexCreateRequest{}
	return &this
}

// GetCollectionName returns the CollectionName field value
func (o *SearchIndexCreateRequest) GetCollectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetCollectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionName, true
}

// SetCollectionName sets field value
func (o *SearchIndexCreateRequest) SetCollectionName(v string) {
	o.CollectionName = v
}

// GetDatabase returns the Database field value
func (o *SearchIndexCreateRequest) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value
func (o *SearchIndexCreateRequest) SetDatabase(v string) {
	o.Database = v
}

// GetName returns the Name field value
func (o *SearchIndexCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SearchIndexCreateRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *SearchIndexCreateRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SearchIndexCreateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SearchIndexCreateRequest) SetType(v string) {
	o.Type = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise
func (o *SearchIndexCreateRequest) GetDefinition() VectorSearchIndexDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret VectorSearchIndexDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetDefinitionOk() (*VectorSearchIndexDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}

	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *SearchIndexCreateRequest) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given VectorSearchIndexDefinition and assigns it to the Definition field.
func (o *SearchIndexCreateRequest) SetDefinition(v VectorSearchIndexDefinition) {
	o.Definition = &v
}

func (o SearchIndexCreateRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o SearchIndexCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collectionName"] = o.CollectionName
	toSerialize["database"] = o.Database
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	return toSerialize, nil
}
