// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsTransitGatewayRouteResponse Container for metadata associated with a Transit Gateway route.
type StreamsTransitGatewayRouteResponse struct {
	// The route's destination CIDR.
	DestinationCidr *string `json:"destinationCidr,omitempty"`
	// The region of the Atlas VPCs where this route should take effect.
	Region *string `json:"region,omitempty"`
	// The AWS ID of the Transit Gateway through which traffic will be routed.
	TgwId *string `json:"tgwId,omitempty"`
	// The ID of the Transit Gateway route.
	// Read only field.
	TgwRouteId *string `json:"tgwRouteId,omitempty"`
}

// NewStreamsTransitGatewayRouteResponse instantiates a new StreamsTransitGatewayRouteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsTransitGatewayRouteResponse() *StreamsTransitGatewayRouteResponse {
	this := StreamsTransitGatewayRouteResponse{}
	return &this
}

// NewStreamsTransitGatewayRouteResponseWithDefaults instantiates a new StreamsTransitGatewayRouteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsTransitGatewayRouteResponseWithDefaults() *StreamsTransitGatewayRouteResponse {
	this := StreamsTransitGatewayRouteResponse{}
	return &this
}

// GetDestinationCidr returns the DestinationCidr field value if set, zero value otherwise
func (o *StreamsTransitGatewayRouteResponse) GetDestinationCidr() string {
	if o == nil || IsNil(o.DestinationCidr) {
		var ret string
		return ret
	}
	return *o.DestinationCidr
}

// GetDestinationCidrOk returns a tuple with the DestinationCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayRouteResponse) GetDestinationCidrOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationCidr) {
		return nil, false
	}

	return o.DestinationCidr, true
}

// HasDestinationCidr returns a boolean if a field has been set.
func (o *StreamsTransitGatewayRouteResponse) HasDestinationCidr() bool {
	if o != nil && !IsNil(o.DestinationCidr) {
		return true
	}

	return false
}

// SetDestinationCidr gets a reference to the given string and assigns it to the DestinationCidr field.
func (o *StreamsTransitGatewayRouteResponse) SetDestinationCidr(v string) {
	o.DestinationCidr = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *StreamsTransitGatewayRouteResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayRouteResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamsTransitGatewayRouteResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StreamsTransitGatewayRouteResponse) SetRegion(v string) {
	o.Region = &v
}

// GetTgwId returns the TgwId field value if set, zero value otherwise
func (o *StreamsTransitGatewayRouteResponse) GetTgwId() string {
	if o == nil || IsNil(o.TgwId) {
		var ret string
		return ret
	}
	return *o.TgwId
}

// GetTgwIdOk returns a tuple with the TgwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayRouteResponse) GetTgwIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwId) {
		return nil, false
	}

	return o.TgwId, true
}

// HasTgwId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayRouteResponse) HasTgwId() bool {
	if o != nil && !IsNil(o.TgwId) {
		return true
	}

	return false
}

// SetTgwId gets a reference to the given string and assigns it to the TgwId field.
func (o *StreamsTransitGatewayRouteResponse) SetTgwId(v string) {
	o.TgwId = &v
}

// GetTgwRouteId returns the TgwRouteId field value if set, zero value otherwise
func (o *StreamsTransitGatewayRouteResponse) GetTgwRouteId() string {
	if o == nil || IsNil(o.TgwRouteId) {
		var ret string
		return ret
	}
	return *o.TgwRouteId
}

// GetTgwRouteIdOk returns a tuple with the TgwRouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayRouteResponse) GetTgwRouteIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwRouteId) {
		return nil, false
	}

	return o.TgwRouteId, true
}

// HasTgwRouteId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayRouteResponse) HasTgwRouteId() bool {
	if o != nil && !IsNil(o.TgwRouteId) {
		return true
	}

	return false
}

// SetTgwRouteId gets a reference to the given string and assigns it to the TgwRouteId field.
func (o *StreamsTransitGatewayRouteResponse) SetTgwRouteId(v string) {
	o.TgwRouteId = &v
}
