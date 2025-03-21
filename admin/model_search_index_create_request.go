// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SearchIndexCreateRequest struct for SearchIndexCreateRequest
type SearchIndexCreateRequest struct {
	// Label that identifies the collection to create an Atlas Search index in.
	CollectionName *string `json:"collectionName,omitempty"`
	// Label that identifies the database that contains the collection to create an Atlas Search index in.
	Database *string `json:"database,omitempty"`
	// Label that identifies this index. Within each namespace, names of all indexes in the namespace must be unique.
	Name *string `json:"name,omitempty"`
	// Type of the index. The default type is search.
	Type       *string                                 `json:"type,omitempty"`
	Definition *BaseSearchIndexCreateRequestDefinition `json:"definition,omitempty"`
}

// NewSearchIndexCreateRequest instantiates a new SearchIndexCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexCreateRequest() *SearchIndexCreateRequest {
	this := SearchIndexCreateRequest{}
	return &this
}

// NewSearchIndexCreateRequestWithDefaults instantiates a new SearchIndexCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexCreateRequestWithDefaults() *SearchIndexCreateRequest {
	this := SearchIndexCreateRequest{}
	return &this
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise
func (o *SearchIndexCreateRequest) GetCollectionName() string {
	if o == nil || IsNil(o.CollectionName) {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetCollectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionName) {
		return nil, false
	}

	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *SearchIndexCreateRequest) HasCollectionName() bool {
	if o != nil && !IsNil(o.CollectionName) {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *SearchIndexCreateRequest) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise
func (o *SearchIndexCreateRequest) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}

	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *SearchIndexCreateRequest) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *SearchIndexCreateRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *SearchIndexCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchIndexCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchIndexCreateRequest) SetName(v string) {
	o.Name = &v
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
func (o *SearchIndexCreateRequest) GetDefinition() BaseSearchIndexCreateRequestDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret BaseSearchIndexCreateRequestDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexCreateRequest) GetDefinitionOk() (*BaseSearchIndexCreateRequestDefinition, bool) {
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

// SetDefinition gets a reference to the given BaseSearchIndexCreateRequestDefinition and assigns it to the Definition field.
func (o *SearchIndexCreateRequest) SetDefinition(v BaseSearchIndexCreateRequestDefinition) {
	o.Definition = &v
}
