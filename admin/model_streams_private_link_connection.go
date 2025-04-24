// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsPrivateLinkConnection Container for metadata needed to create a Private Link connection.
type StreamsPrivateLinkConnection struct {
	// The ID of the Private Link connection.
	// Read only field.
	Id *string `json:"_id,omitempty"`
	// Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty"`
	// Domain name of Privatelink connected cluster.
	DnsDomain *string `json:"dnsDomain,omitempty"`
	// Sub-Domain name of Confluent cluster. These are typically your availability zones.
	DnsSubDomain *[]string `json:"dnsSubDomain,omitempty"`
	// Error message if the state is FAILED.
	// Read only field.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Interface endpoint ID that is created from the service endpoint ID provided.
	// Read only field.
	InterfaceEndpointId *string `json:"interfaceEndpointId,omitempty"`
	// Interface endpoint name that is created from the service endpoint ID provided.
	// Read only field.
	InterfaceEndpointName *string `json:"interfaceEndpointName,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Provider where the Kafka cluster is deployed.
	Provider *string `json:"provider,omitempty"`
	// Account ID from the cloud provider.
	// Read only field.
	ProviderAccountId *string `json:"providerAccountId,omitempty"`
	// Domain name of Confluent cluster.
	Region *string `json:"region,omitempty"`
	// Service Endpoint ID.
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty"`
	// State the connection is in.
	// Read only field.
	State *string `json:"state,omitempty"`
	// Vendor who manages the Kafka cluster.
	Vendor *string `json:"vendor,omitempty"`
}

// NewStreamsPrivateLinkConnection instantiates a new StreamsPrivateLinkConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsPrivateLinkConnection() *StreamsPrivateLinkConnection {
	this := StreamsPrivateLinkConnection{}
	return &this
}

// NewStreamsPrivateLinkConnectionWithDefaults instantiates a new StreamsPrivateLinkConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsPrivateLinkConnectionWithDefaults() *StreamsPrivateLinkConnection {
	this := StreamsPrivateLinkConnection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StreamsPrivateLinkConnection) SetId(v string) {
	o.Id = &v
}

// GetArn returns the Arn field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetArn() string {
	if o == nil || IsNil(o.Arn) {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetArnOk() (*string, bool) {
	if o == nil || IsNil(o.Arn) {
		return nil, false
	}

	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasArn() bool {
	if o != nil && !IsNil(o.Arn) {
		return true
	}

	return false
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *StreamsPrivateLinkConnection) SetArn(v string) {
	o.Arn = &v
}

// GetDnsDomain returns the DnsDomain field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetDnsDomain() string {
	if o == nil || IsNil(o.DnsDomain) {
		var ret string
		return ret
	}
	return *o.DnsDomain
}

// GetDnsDomainOk returns a tuple with the DnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetDnsDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DnsDomain) {
		return nil, false
	}

	return o.DnsDomain, true
}

// HasDnsDomain returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasDnsDomain() bool {
	if o != nil && !IsNil(o.DnsDomain) {
		return true
	}

	return false
}

// SetDnsDomain gets a reference to the given string and assigns it to the DnsDomain field.
func (o *StreamsPrivateLinkConnection) SetDnsDomain(v string) {
	o.DnsDomain = &v
}

// GetDnsSubDomain returns the DnsSubDomain field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetDnsSubDomain() []string {
	if o == nil || IsNil(o.DnsSubDomain) {
		var ret []string
		return ret
	}
	return *o.DnsSubDomain
}

// GetDnsSubDomainOk returns a tuple with the DnsSubDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetDnsSubDomainOk() (*[]string, bool) {
	if o == nil || IsNil(o.DnsSubDomain) {
		return nil, false
	}

	return o.DnsSubDomain, true
}

// HasDnsSubDomain returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasDnsSubDomain() bool {
	if o != nil && !IsNil(o.DnsSubDomain) {
		return true
	}

	return false
}

// SetDnsSubDomain gets a reference to the given []string and assigns it to the DnsSubDomain field.
func (o *StreamsPrivateLinkConnection) SetDnsSubDomain(v []string) {
	o.DnsSubDomain = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *StreamsPrivateLinkConnection) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetInterfaceEndpointId returns the InterfaceEndpointId field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointId() string {
	if o == nil || IsNil(o.InterfaceEndpointId) {
		var ret string
		return ret
	}
	return *o.InterfaceEndpointId
}

// GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceEndpointId) {
		return nil, false
	}

	return o.InterfaceEndpointId, true
}

// HasInterfaceEndpointId returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasInterfaceEndpointId() bool {
	if o != nil && !IsNil(o.InterfaceEndpointId) {
		return true
	}

	return false
}

// SetInterfaceEndpointId gets a reference to the given string and assigns it to the InterfaceEndpointId field.
func (o *StreamsPrivateLinkConnection) SetInterfaceEndpointId(v string) {
	o.InterfaceEndpointId = &v
}

// GetInterfaceEndpointName returns the InterfaceEndpointName field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointName() string {
	if o == nil || IsNil(o.InterfaceEndpointName) {
		var ret string
		return ret
	}
	return *o.InterfaceEndpointName
}

// GetInterfaceEndpointNameOk returns a tuple with the InterfaceEndpointName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointNameOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceEndpointName) {
		return nil, false
	}

	return o.InterfaceEndpointName, true
}

// HasInterfaceEndpointName returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasInterfaceEndpointName() bool {
	if o != nil && !IsNil(o.InterfaceEndpointName) {
		return true
	}

	return false
}

// SetInterfaceEndpointName gets a reference to the given string and assigns it to the InterfaceEndpointName field.
func (o *StreamsPrivateLinkConnection) SetInterfaceEndpointName(v string) {
	o.InterfaceEndpointName = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsPrivateLinkConnection) SetLinks(v []Link) {
	o.Links = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}

	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *StreamsPrivateLinkConnection) SetProvider(v string) {
	o.Provider = &v
}

// GetProviderAccountId returns the ProviderAccountId field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetProviderAccountId() string {
	if o == nil || IsNil(o.ProviderAccountId) {
		var ret string
		return ret
	}
	return *o.ProviderAccountId
}

// GetProviderAccountIdOk returns a tuple with the ProviderAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetProviderAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderAccountId) {
		return nil, false
	}

	return o.ProviderAccountId, true
}

// HasProviderAccountId returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasProviderAccountId() bool {
	if o != nil && !IsNil(o.ProviderAccountId) {
		return true
	}

	return false
}

// SetProviderAccountId gets a reference to the given string and assigns it to the ProviderAccountId field.
func (o *StreamsPrivateLinkConnection) SetProviderAccountId(v string) {
	o.ProviderAccountId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StreamsPrivateLinkConnection) SetRegion(v string) {
	o.Region = &v
}

// GetServiceEndpointId returns the ServiceEndpointId field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetServiceEndpointId() string {
	if o == nil || IsNil(o.ServiceEndpointId) {
		var ret string
		return ret
	}
	return *o.ServiceEndpointId
}

// GetServiceEndpointIdOk returns a tuple with the ServiceEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetServiceEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceEndpointId) {
		return nil, false
	}

	return o.ServiceEndpointId, true
}

// HasServiceEndpointId returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasServiceEndpointId() bool {
	if o != nil && !IsNil(o.ServiceEndpointId) {
		return true
	}

	return false
}

// SetServiceEndpointId gets a reference to the given string and assigns it to the ServiceEndpointId field.
func (o *StreamsPrivateLinkConnection) SetServiceEndpointId(v string) {
	o.ServiceEndpointId = &v
}

// GetState returns the State field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StreamsPrivateLinkConnection) SetState(v string) {
	o.State = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnection) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnection) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}

	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnection) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *StreamsPrivateLinkConnection) SetVendor(v string) {
	o.Vendor = &v
}
