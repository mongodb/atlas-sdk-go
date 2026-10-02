// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsWorkspacePrivateEndpointRequest Customer VPC endpoint to register for inbound private connectivity to Atlas Stream Processing.
type StreamsWorkspacePrivateEndpointRequest struct {
	// DNS name of the customer's VPC endpoint.
	CustomerDnsName string `json:"customerDnsName"`
	// Unique identifier of the customer's VPC endpoint.
	Id string `json:"id"`
	// Human-readable label identifying the region of the customer's VPC endpoint.
	Region string `json:"region"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsWorkspacePrivateEndpointRequest) MarshalJSON() ([]byte, error) {
	type noMethod StreamsWorkspacePrivateEndpointRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsWorkspacePrivateEndpointRequest instantiates a new StreamsWorkspacePrivateEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsWorkspacePrivateEndpointRequest(customerDnsName string, id string, region string) *StreamsWorkspacePrivateEndpointRequest {
	this := StreamsWorkspacePrivateEndpointRequest{}
	this.CustomerDnsName = customerDnsName
	this.Id = id
	this.Region = region
	return &this
}

// NewStreamsWorkspacePrivateEndpointRequestWithDefaults instantiates a new StreamsWorkspacePrivateEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsWorkspacePrivateEndpointRequestWithDefaults() *StreamsWorkspacePrivateEndpointRequest {
	this := StreamsWorkspacePrivateEndpointRequest{}
	return &this
}

// GetCustomerDnsName returns the CustomerDnsName field value
func (o *StreamsWorkspacePrivateEndpointRequest) GetCustomerDnsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerDnsName
}

// GetCustomerDnsNameOk returns a tuple with the CustomerDnsName field value
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointRequest) GetCustomerDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerDnsName, true
}

// SetCustomerDnsName sets field value
func (o *StreamsWorkspacePrivateEndpointRequest) SetCustomerDnsName(v string) {
	o.CustomerDnsName = v
}

// GetId returns the Id field value
func (o *StreamsWorkspacePrivateEndpointRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StreamsWorkspacePrivateEndpointRequest) SetId(v string) {
	o.Id = v
}

// GetRegion returns the Region field value
func (o *StreamsWorkspacePrivateEndpointRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *StreamsWorkspacePrivateEndpointRequest) SetRegion(v string) {
	o.Region = v
}
