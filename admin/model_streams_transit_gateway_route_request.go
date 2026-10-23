// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsTransitGatewayRouteRequest Container for metadata needed to create a Transit Gateway route.
type StreamsTransitGatewayRouteRequest struct {
	// The route's destination CIDR.
	DestinationCidr *string `json:"destinationCidr,omitempty"`
	// The region of the Atlas VPCs where this route should take effect.
	Region *string `json:"region,omitempty"`
	// The AWS ID of the Transit Gateway through which traffic will be routed.
	TgwId *string `json:"tgwId,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsTransitGatewayRouteRequest) MarshalJSON() ([]byte, error) {
	type noMethod StreamsTransitGatewayRouteRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsTransitGatewayRouteRequest instantiates a new StreamsTransitGatewayRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsTransitGatewayRouteRequest() *StreamsTransitGatewayRouteRequest {
	this := StreamsTransitGatewayRouteRequest{}
	return &this
}

// NewStreamsTransitGatewayRouteRequestWithDefaults instantiates a new StreamsTransitGatewayRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsTransitGatewayRouteRequestWithDefaults() *StreamsTransitGatewayRouteRequest {
	this := StreamsTransitGatewayRouteRequest{}
	return &this
}

// GetDestinationCidr returns the DestinationCidr field value if set, zero value otherwise
func (o *StreamsTransitGatewayRouteRequest) GetDestinationCidr() string {
	if o == nil || IsNil(o.DestinationCidr) {
		var ret string
		return ret
	}
	return *o.DestinationCidr
}

// GetDestinationCidrOk returns a tuple with the DestinationCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayRouteRequest) GetDestinationCidrOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationCidr) {
		return nil, false
	}

	return o.DestinationCidr, true
}

// HasDestinationCidr returns a boolean if a field has been set.
func (o *StreamsTransitGatewayRouteRequest) HasDestinationCidr() bool {
	if o != nil && !IsNil(o.DestinationCidr) {
		return true
	}

	return false
}

// SetDestinationCidr gets a reference to the given string and assigns it to the DestinationCidr field.
func (o *StreamsTransitGatewayRouteRequest) SetDestinationCidr(v string) {
	o.DestinationCidr = &v
	o.NullFields = removeNullField(o.NullFields, "DestinationCidr")
}

// SetDestinationCidrNil sets DestinationCidr to an explicit JSON null when marshaled.
func (o *StreamsTransitGatewayRouteRequest) SetDestinationCidrNil() {
	o.DestinationCidr = nil
	o.NullFields = addNullField(o.NullFields, "DestinationCidr")
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *StreamsTransitGatewayRouteRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayRouteRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamsTransitGatewayRouteRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StreamsTransitGatewayRouteRequest) SetRegion(v string) {
	o.Region = &v
	o.NullFields = removeNullField(o.NullFields, "Region")
}

// SetRegionNil sets Region to an explicit JSON null when marshaled.
func (o *StreamsTransitGatewayRouteRequest) SetRegionNil() {
	o.Region = nil
	o.NullFields = addNullField(o.NullFields, "Region")
}

// GetTgwId returns the TgwId field value if set, zero value otherwise
func (o *StreamsTransitGatewayRouteRequest) GetTgwId() string {
	if o == nil || IsNil(o.TgwId) {
		var ret string
		return ret
	}
	return *o.TgwId
}

// GetTgwIdOk returns a tuple with the TgwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTransitGatewayRouteRequest) GetTgwIdOk() (*string, bool) {
	if o == nil || IsNil(o.TgwId) {
		return nil, false
	}

	return o.TgwId, true
}

// HasTgwId returns a boolean if a field has been set.
func (o *StreamsTransitGatewayRouteRequest) HasTgwId() bool {
	if o != nil && !IsNil(o.TgwId) {
		return true
	}

	return false
}

// SetTgwId gets a reference to the given string and assigns it to the TgwId field.
func (o *StreamsTransitGatewayRouteRequest) SetTgwId(v string) {
	o.TgwId = &v
	o.NullFields = removeNullField(o.NullFields, "TgwId")
}

// SetTgwIdNil sets TgwId to an explicit JSON null when marshaled.
func (o *StreamsTransitGatewayRouteRequest) SetTgwIdNil() {
	o.TgwId = nil
	o.NullFields = addNullField(o.NullFields, "TgwId")
}
