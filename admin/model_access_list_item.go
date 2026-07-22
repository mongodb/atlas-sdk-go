// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AccessListItem struct for AccessListItem
type AccessListItem struct {
	// Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that found in this project's access list.
	// Read only field.
	CidrBlock *string `json:"cidrBlock,omitempty"`
	// IP address included in the API access list.
	// Read only field.
	IpAddress string `json:"ipAddress"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AccessListItem) MarshalJSON() ([]byte, error) {
	type noMethod AccessListItem
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAccessListItem instantiates a new AccessListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessListItem(ipAddress string) *AccessListItem {
	this := AccessListItem{}
	this.IpAddress = ipAddress
	return &this
}

// NewAccessListItemWithDefaults instantiates a new AccessListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessListItemWithDefaults() *AccessListItem {
	this := AccessListItem{}
	return &this
}

// GetCidrBlock returns the CidrBlock field value if set, zero value otherwise
func (o *AccessListItem) GetCidrBlock() string {
	if o == nil || IsNil(o.CidrBlock) {
		var ret string
		return ret
	}
	return *o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessListItem) GetCidrBlockOk() (*string, bool) {
	if o == nil || IsNil(o.CidrBlock) {
		return nil, false
	}

	return o.CidrBlock, true
}

// HasCidrBlock returns a boolean if a field has been set.
func (o *AccessListItem) HasCidrBlock() bool {
	if o != nil && !IsNil(o.CidrBlock) {
		return true
	}

	return false
}

// SetCidrBlock gets a reference to the given string and assigns it to the CidrBlock field.
func (o *AccessListItem) SetCidrBlock(v string) {
	o.CidrBlock = &v
	o.NullFields = removeNullField(o.NullFields, "CidrBlock")
}

// SetCidrBlockNil sets CidrBlock to an explicit JSON null when marshaled.
func (o *AccessListItem) SetCidrBlockNil() {
	o.CidrBlock = nil
	o.NullFields = addNullField(o.NullFields, "CidrBlock")
}

// GetIpAddress returns the IpAddress field value
func (o *AccessListItem) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *AccessListItem) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *AccessListItem) SetIpAddress(v string) {
	o.IpAddress = v
}
