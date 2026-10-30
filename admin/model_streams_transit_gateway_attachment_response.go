// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsTransitGatewayAttachmentResponse struct for StreamsTransitGatewayAttachmentResponse
type StreamsTransitGatewayAttachmentResponse struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// The AWS Transit Gateway Attachment ID.
	// Read only field.
	TgwAttachmentId *string `json:"tgwAttachmentId,omitempty"`
	// AWS transit gateway ID.
	TgwId *string `json:"tgwId,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsTransitGatewayAttachmentResponse) MarshalJSON() ([]byte, error) {
	type noMethod StreamsTransitGatewayAttachmentResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsTransitGatewayAttachmentResponse instantiates a new StreamsTransitGatewayAttachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsTransitGatewayAttachmentResponse() *StreamsTransitGatewayAttachmentResponse {
	this := StreamsTransitGatewayAttachmentResponse{}
	return &this
}

// NewStreamsTransitGatewayAttachmentResponseWithDefaults instantiates a new StreamsTransitGatewayAttachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsTransitGatewayAttachmentResponseWithDefaults() *StreamsTransitGatewayAttachmentResponse {
	this := StreamsTransitGatewayAttachmentResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsTransitGatewayAttachmentResponse) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayAttachmentResponse) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsTransitGatewayAttachmentResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsTransitGatewayAttachmentResponse) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsTransitGatewayAttachmentResponse) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetTgwAttachmentId returns the TgwAttachmentId field value if set, zero value otherwise
func (o *StreamsTransitGatewayAttachmentResponse) GetTgwAttachmentId() string {
	if o == nil || IsNil(o.TgwAttachmentId) {
		var ret string
		return ret
	}
	return *o.TgwAttachmentId
}

// GetTgwAttachmentIdOk returns a tuple with the TgwAttachmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayAttachmentResponse) GetTgwAttachmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwAttachmentId) {
		return nil, false
	}

	return o.TgwAttachmentId, true
}

// HasTgwAttachmentId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayAttachmentResponse) HasTgwAttachmentId() bool {
	if o != nil && !IsNil(o.TgwAttachmentId) {
		return true
	}

	return false
}

// SetTgwAttachmentId gets a reference to the given string and assigns it to the TgwAttachmentId field.
func (o *StreamsTransitGatewayAttachmentResponse) SetTgwAttachmentId(v string) {
	o.TgwAttachmentId = &v
	o.NullFields = removeNullField(o.NullFields, "TgwAttachmentId")
}

// SetTgwAttachmentIdNil sets TgwAttachmentId to an explicit JSON null when marshaled.
func (o *StreamsTransitGatewayAttachmentResponse) SetTgwAttachmentIdNil() {
	o.TgwAttachmentId = nil
	o.NullFields = addNullField(o.NullFields, "TgwAttachmentId")
}

// GetTgwId returns the TgwId field value if set, zero value otherwise
func (o *StreamsTransitGatewayAttachmentResponse) GetTgwId() string {
	if o == nil || IsNil(o.TgwId) {
		var ret string
		return ret
	}
	return *o.TgwId
}

// GetTgwIdOk returns a tuple with the TgwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayAttachmentResponse) GetTgwIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwId) {
		return nil, false
	}

	return o.TgwId, true
}

// HasTgwId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayAttachmentResponse) HasTgwId() bool {
	if o != nil && !IsNil(o.TgwId) {
		return true
	}

	return false
}

// SetTgwId gets a reference to the given string and assigns it to the TgwId field.
func (o *StreamsTransitGatewayAttachmentResponse) SetTgwId(v string) {
	o.TgwId = &v
	o.NullFields = removeNullField(o.NullFields, "TgwId")
}

// SetTgwIdNil sets TgwId to an explicit JSON null when marshaled.
func (o *StreamsTransitGatewayAttachmentResponse) SetTgwIdNil() {
	o.TgwId = nil
	o.NullFields = addNullField(o.NullFields, "TgwId")
}
