// Code based on the AtlasAPI V2 OpenAPI file

package admin

// EnvelopedSchemaAdvisorResponse Response envelope that wraps the response payload in `content` and includes response metadata such as `status` and `locations`.
type EnvelopedSchemaAdvisorResponse struct {
	Content SchemaAdvisorResponse `json:"content"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// URLs of resources created by this request.
	// Read only field.
	Locations *[]string `json:"locations,omitempty"`
	// HTTP status code returned with this response.
	// Read only field.
	Status int `json:"status"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *EnvelopedSchemaAdvisorResponse) MarshalJSON() ([]byte, error) {
	type noMethod EnvelopedSchemaAdvisorResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewEnvelopedSchemaAdvisorResponse instantiates a new EnvelopedSchemaAdvisorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopedSchemaAdvisorResponse(content SchemaAdvisorResponse, status int) *EnvelopedSchemaAdvisorResponse {
	this := EnvelopedSchemaAdvisorResponse{}
	this.Content = content
	this.Status = status
	return &this
}

// NewEnvelopedSchemaAdvisorResponseWithDefaults instantiates a new EnvelopedSchemaAdvisorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopedSchemaAdvisorResponseWithDefaults() *EnvelopedSchemaAdvisorResponse {
	this := EnvelopedSchemaAdvisorResponse{}
	return &this
}

// GetContent returns the Content field value
func (o *EnvelopedSchemaAdvisorResponse) GetContent() SchemaAdvisorResponse {
	if o == nil {
		var ret SchemaAdvisorResponse
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *EnvelopedSchemaAdvisorResponse) GetContentOk() (*SchemaAdvisorResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *EnvelopedSchemaAdvisorResponse) SetContent(v SchemaAdvisorResponse) {
	o.Content = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *EnvelopedSchemaAdvisorResponse) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopedSchemaAdvisorResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvelopedSchemaAdvisorResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *EnvelopedSchemaAdvisorResponse) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *EnvelopedSchemaAdvisorResponse) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetLocations returns the Locations field value if set, zero value otherwise
func (o *EnvelopedSchemaAdvisorResponse) GetLocations() []string {
	if o == nil || IsNil(o.Locations) {
		var ret []string
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopedSchemaAdvisorResponse) GetLocationsOk() (*[]string, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}

	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *EnvelopedSchemaAdvisorResponse) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *EnvelopedSchemaAdvisorResponse) SetLocations(v []string) {
	o.Locations = &v
	o.NullFields = removeNullField(o.NullFields, "Locations")
}

// SetLocationsNil sets Locations to an explicit JSON null when marshaled.
func (o *EnvelopedSchemaAdvisorResponse) SetLocationsNil() {
	o.Locations = nil
	o.NullFields = addNullField(o.NullFields, "Locations")
}

// GetStatus returns the Status field value
func (o *EnvelopedSchemaAdvisorResponse) GetStatus() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvelopedSchemaAdvisorResponse) GetStatusOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EnvelopedSchemaAdvisorResponse) SetStatus(v int) {
	o.Status = v
}
