// Code based on the AtlasAPI V2 OpenAPI file

package admin

// Gateway Represents a service-specific gateway, such as the Atlas Gateway, with its IP addresses.
type Gateway struct {
	Ips *GatewayIpAddresses `json:"ips,omitempty"`
	// Name of the service that this gateway represents.
	// Read only field.
	Name *string `json:"name,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *Gateway) MarshalJSON() ([]byte, error) {
	type noMethod Gateway
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGateway instantiates a new Gateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateway() *Gateway {
	this := Gateway{}
	return &this
}

// NewGatewayWithDefaults instantiates a new Gateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayWithDefaults() *Gateway {
	this := Gateway{}
	return &this
}

// GetIps returns the Ips field value if set, zero value otherwise
func (o *Gateway) GetIps() GatewayIpAddresses {
	if o == nil || IsNil(o.Ips) {
		var ret GatewayIpAddresses
		return ret
	}
	return *o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetIpsOk() (*GatewayIpAddresses, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}

	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *Gateway) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given GatewayIpAddresses and assigns it to the Ips field.
func (o *Gateway) SetIps(v GatewayIpAddresses) {
	o.Ips = &v
	o.NullFields = removeNullField(o.NullFields, "Ips")
}

// SetIpsNil sets Ips to an explicit JSON null when marshaled.
func (o *Gateway) SetIpsNil() {
	o.Ips = nil
	o.NullFields = addNullField(o.NullFields, "Ips")
}

// GetName returns the Name field value if set, zero value otherwise
func (o *Gateway) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Gateway) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Gateway) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *Gateway) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}
