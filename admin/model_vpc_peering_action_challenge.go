// Code based on the AtlasAPI V2 OpenAPI file

package admin

// VPCPeeringActionChallenge Container for elements used to challenge the user before taking certain actions on VPC Peering connections.
type VPCPeeringActionChallenge struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// The AWS requester account ID.
	RequesterAccountId *string `json:"requesterAccountId,omitempty"`
	// The AWS requester VPC ID.
	RequesterVpcId *string `json:"requesterVpcId,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *VPCPeeringActionChallenge) MarshalJSON() ([]byte, error) {
	type noMethod VPCPeeringActionChallenge
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewVPCPeeringActionChallenge instantiates a new VPCPeeringActionChallenge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVPCPeeringActionChallenge() *VPCPeeringActionChallenge {
	this := VPCPeeringActionChallenge{}
	return &this
}

// NewVPCPeeringActionChallengeWithDefaults instantiates a new VPCPeeringActionChallenge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVPCPeeringActionChallengeWithDefaults() *VPCPeeringActionChallenge {
	this := VPCPeeringActionChallenge{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *VPCPeeringActionChallenge) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringActionChallenge) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VPCPeeringActionChallenge) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *VPCPeeringActionChallenge) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *VPCPeeringActionChallenge) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetRequesterAccountId returns the RequesterAccountId field value if set, zero value otherwise
func (o *VPCPeeringActionChallenge) GetRequesterAccountId() string {
	if o == nil || IsNil(o.RequesterAccountId) {
		var ret string
		return ret
	}
	return *o.RequesterAccountId
}

// GetRequesterAccountIdOk returns a tuple with the RequesterAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringActionChallenge) GetRequesterAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterAccountId) {
		return nil, false
	}

	return o.RequesterAccountId, true
}

// HasRequesterAccountId returns a boolean if a field has been set.
func (o *VPCPeeringActionChallenge) HasRequesterAccountId() bool {
	if o != nil && !IsNil(o.RequesterAccountId) {
		return true
	}

	return false
}

// SetRequesterAccountId gets a reference to the given string and assigns it to the RequesterAccountId field.
func (o *VPCPeeringActionChallenge) SetRequesterAccountId(v string) {
	o.RequesterAccountId = &v
	o.NullFields = removeNullField(o.NullFields, "RequesterAccountId")
}

// SetRequesterAccountIdNil sets RequesterAccountId to an explicit JSON null when marshaled.
func (o *VPCPeeringActionChallenge) SetRequesterAccountIdNil() {
	o.RequesterAccountId = nil
	o.NullFields = addNullField(o.NullFields, "RequesterAccountId")
}

// GetRequesterVpcId returns the RequesterVpcId field value if set, zero value otherwise
func (o *VPCPeeringActionChallenge) GetRequesterVpcId() string {
	if o == nil || IsNil(o.RequesterVpcId) {
		var ret string
		return ret
	}
	return *o.RequesterVpcId
}

// GetRequesterVpcIdOk returns a tuple with the RequesterVpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringActionChallenge) GetRequesterVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterVpcId) {
		return nil, false
	}

	return o.RequesterVpcId, true
}

// HasRequesterVpcId returns a boolean if a field has been set.
func (o *VPCPeeringActionChallenge) HasRequesterVpcId() bool {
	if o != nil && !IsNil(o.RequesterVpcId) {
		return true
	}

	return false
}

// SetRequesterVpcId gets a reference to the given string and assigns it to the RequesterVpcId field.
func (o *VPCPeeringActionChallenge) SetRequesterVpcId(v string) {
	o.RequesterVpcId = &v
	o.NullFields = removeNullField(o.NullFields, "RequesterVpcId")
}

// SetRequesterVpcIdNil sets RequesterVpcId to an explicit JSON null when marshaled.
func (o *VPCPeeringActionChallenge) SetRequesterVpcIdNil() {
	o.RequesterVpcId = nil
	o.NullFields = addNullField(o.NullFields, "RequesterVpcId")
}
