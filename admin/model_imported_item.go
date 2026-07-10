// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ImportedItem Imported Item.
type ImportedItem struct {
	// Embedding ID of the item. This will only be returned if embedding is enabled for the entity.
	EmbeddingId *string `json:"embeddingId,omitempty"`
	// ID of the item.
	Id string `json:"id"`
	// Title of the item.
	Title string `json:"title"`
}

// NewImportedItem instantiates a new ImportedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedItem(id string, title string) *ImportedItem {
	this := ImportedItem{}
	this.Id = id
	this.Title = title
	return &this
}

// NewImportedItemWithDefaults instantiates a new ImportedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedItemWithDefaults() *ImportedItem {
	this := ImportedItem{}
	return &this
}

// GetEmbeddingId returns the EmbeddingId field value if set, zero value otherwise
func (o *ImportedItem) GetEmbeddingId() string {
	if o == nil || IsNil(o.EmbeddingId) {
		var ret string
		return ret
	}
	return *o.EmbeddingId
}

// GetEmbeddingIdOk returns a tuple with the EmbeddingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportedItem) GetEmbeddingIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmbeddingId) {
		return nil, false
	}

	return o.EmbeddingId, true
}

// HasEmbeddingId returns a boolean if a field has been set.
func (o *ImportedItem) HasEmbeddingId() bool {
	if o != nil && !IsNil(o.EmbeddingId) {
		return true
	}

	return false
}

// SetEmbeddingId gets a reference to the given string and assigns it to the EmbeddingId field.
func (o *ImportedItem) SetEmbeddingId(v string) {
	o.EmbeddingId = &v
}

// GetId returns the Id field value
func (o *ImportedItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImportedItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImportedItem) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *ImportedItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ImportedItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ImportedItem) SetTitle(v string) {
	o.Title = v
}
