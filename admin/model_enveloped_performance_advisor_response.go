// Code based on the AtlasAPI V2 OpenAPI file

package admin

// EnvelopedPerformanceAdvisorResponse Response envelope that wraps the response payload in `content` and includes response metadata such as `status` and `locations`.
type EnvelopedPerformanceAdvisorResponse struct {
	Content PerformanceAdvisorResponse `json:"content"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// URLs of resources created by this request.
	// Read only field.
	Locations *[]string `json:"locations,omitempty"`
	// HTTP status code returned with this response.
	// Read only field.
	Status int `json:"status"`
}

// NewEnvelopedPerformanceAdvisorResponse instantiates a new EnvelopedPerformanceAdvisorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopedPerformanceAdvisorResponse(content PerformanceAdvisorResponse, status int) *EnvelopedPerformanceAdvisorResponse {
	this := EnvelopedPerformanceAdvisorResponse{}
	this.Content = content
	this.Status = status
	return &this
}

// NewEnvelopedPerformanceAdvisorResponseWithDefaults instantiates a new EnvelopedPerformanceAdvisorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopedPerformanceAdvisorResponseWithDefaults() *EnvelopedPerformanceAdvisorResponse {
	this := EnvelopedPerformanceAdvisorResponse{}
	return &this
}

// GetContent returns the Content field value
func (o *EnvelopedPerformanceAdvisorResponse) GetContent() PerformanceAdvisorResponse {
	if o == nil {
		var ret PerformanceAdvisorResponse
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *EnvelopedPerformanceAdvisorResponse) GetContentOk() (*PerformanceAdvisorResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *EnvelopedPerformanceAdvisorResponse) SetContent(v PerformanceAdvisorResponse) {
	o.Content = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *EnvelopedPerformanceAdvisorResponse) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopedPerformanceAdvisorResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvelopedPerformanceAdvisorResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *EnvelopedPerformanceAdvisorResponse) SetLinks(v []Link) {
	o.Links = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise
func (o *EnvelopedPerformanceAdvisorResponse) GetLocations() []string {
	if o == nil || IsNil(o.Locations) {
		var ret []string
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopedPerformanceAdvisorResponse) GetLocationsOk() (*[]string, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}

	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *EnvelopedPerformanceAdvisorResponse) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *EnvelopedPerformanceAdvisorResponse) SetLocations(v []string) {
	o.Locations = &v
}

// GetStatus returns the Status field value
func (o *EnvelopedPerformanceAdvisorResponse) GetStatus() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvelopedPerformanceAdvisorResponse) GetStatusOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EnvelopedPerformanceAdvisorResponse) SetStatus(v int) {
	o.Status = v
}
