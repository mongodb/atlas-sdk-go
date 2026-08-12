// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GatewayIpAddresses IP addresses for a specific gateway, organized by direction and cloud provider.
type GatewayIpAddresses struct {
	Inbound  *InboundControlPlaneCloudProviderIPAddresses  `json:"inbound,omitempty"`
	Outbound *OutboundControlPlaneCloudProviderIPAddresses `json:"outbound,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GatewayIpAddresses) MarshalJSON() ([]byte, error) {
	type noMethod GatewayIpAddresses
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGatewayIpAddresses instantiates a new GatewayIpAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayIpAddresses() *GatewayIpAddresses {
	this := GatewayIpAddresses{}
	return &this
}

// NewGatewayIpAddressesWithDefaults instantiates a new GatewayIpAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayIpAddressesWithDefaults() *GatewayIpAddresses {
	this := GatewayIpAddresses{}
	return &this
}

// GetInbound returns the Inbound field value if set, zero value otherwise
func (o *GatewayIpAddresses) GetInbound() InboundControlPlaneCloudProviderIPAddresses {
	if o == nil || IsNil(o.Inbound) {
		var ret InboundControlPlaneCloudProviderIPAddresses
		return ret
	}
	return *o.Inbound
}

// GetInboundOk returns a tuple with the Inbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayIpAddresses) GetInboundOk() (*InboundControlPlaneCloudProviderIPAddresses, bool) {
	if o == nil || IsNil(o.Inbound) {
		return nil, false
	}

	return o.Inbound, true
}

// HasInbound returns a boolean if a field has been set.
func (o *GatewayIpAddresses) HasInbound() bool {
	if o != nil && !IsNil(o.Inbound) {
		return true
	}

	return false
}

// SetInbound gets a reference to the given InboundControlPlaneCloudProviderIPAddresses and assigns it to the Inbound field.
func (o *GatewayIpAddresses) SetInbound(v InboundControlPlaneCloudProviderIPAddresses) {
	o.Inbound = &v
	o.NullFields = removeNullField(o.NullFields, "Inbound")
}

// SetInboundNil sets Inbound to an explicit JSON null when marshaled.
func (o *GatewayIpAddresses) SetInboundNil() {
	o.Inbound = nil
	o.NullFields = addNullField(o.NullFields, "Inbound")
}

// GetOutbound returns the Outbound field value if set, zero value otherwise
func (o *GatewayIpAddresses) GetOutbound() OutboundControlPlaneCloudProviderIPAddresses {
	if o == nil || IsNil(o.Outbound) {
		var ret OutboundControlPlaneCloudProviderIPAddresses
		return ret
	}
	return *o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayIpAddresses) GetOutboundOk() (*OutboundControlPlaneCloudProviderIPAddresses, bool) {
	if o == nil || IsNil(o.Outbound) {
		return nil, false
	}

	return o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *GatewayIpAddresses) HasOutbound() bool {
	if o != nil && !IsNil(o.Outbound) {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given OutboundControlPlaneCloudProviderIPAddresses and assigns it to the Outbound field.
func (o *GatewayIpAddresses) SetOutbound(v OutboundControlPlaneCloudProviderIPAddresses) {
	o.Outbound = &v
	o.NullFields = removeNullField(o.NullFields, "Outbound")
}

// SetOutboundNil sets Outbound to an explicit JSON null when marshaled.
func (o *GatewayIpAddresses) SetOutboundNil() {
	o.Outbound = nil
	o.NullFields = addNullField(o.NullFields, "Outbound")
}
