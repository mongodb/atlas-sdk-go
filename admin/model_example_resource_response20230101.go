// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ExampleResourceResponse20230101 struct for ExampleResourceResponse20230101
type ExampleResourceResponse20230101 struct {
	// Dummy data added as response.
	Data string `json:"data"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
}

// NewExampleResourceResponse20230101 instantiates a new ExampleResourceResponse20230101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleResourceResponse20230101(data string) *ExampleResourceResponse20230101 {
	this := ExampleResourceResponse20230101{}
	this.Data = data
	return &this
}

// NewExampleResourceResponse20230101WithDefaults instantiates a new ExampleResourceResponse20230101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleResourceResponse20230101WithDefaults() *ExampleResourceResponse20230101 {
	this := ExampleResourceResponse20230101{}
	return &this
}

// GetData returns the Data field value
func (o *ExampleResourceResponse20230101) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponse20230101) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExampleResourceResponse20230101) SetData(v string) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *ExampleResourceResponse20230101) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleResourceResponse20230101) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExampleResourceResponse20230101) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ExampleResourceResponse20230101) SetLinks(v []Link) {
	o.Links = &v
}
