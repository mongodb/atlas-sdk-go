// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsWorkspacePrivateEndpointResponse Customer VPC endpoint registered for inbound private connectivity to Atlas Stream Processing.
type StreamsWorkspacePrivateEndpointResponse struct {
	// DNS name of the customer's VPC endpoint.
	// Read only field.
	CustomerDnsName *string `json:"customerDnsName,omitempty"`
	// Name of the Atlas endpoint service the customer's VPC endpoint connects to. Absent when no endpoint service is configured for the endpoint's region.
	// Read only field.
	EndpointServiceName *string `json:"endpointServiceName,omitempty"`
	// Unique identifier of the customer's VPC endpoint.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Cloud provider hosting the customer's VPC endpoint.
	// Read only field.
	Provider *string `json:"provider,omitempty"`
	// Human-readable label identifying the region of the customer's VPC endpoint.
	// Read only field.
	Region *string `json:"region,omitempty"`
	// Status of the endpoint registration.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsWorkspacePrivateEndpointResponse) MarshalJSON() ([]byte, error) {
	type noMethod StreamsWorkspacePrivateEndpointResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsWorkspacePrivateEndpointResponse instantiates a new StreamsWorkspacePrivateEndpointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsWorkspacePrivateEndpointResponse() *StreamsWorkspacePrivateEndpointResponse {
	this := StreamsWorkspacePrivateEndpointResponse{}
	return &this
}

// NewStreamsWorkspacePrivateEndpointResponseWithDefaults instantiates a new StreamsWorkspacePrivateEndpointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsWorkspacePrivateEndpointResponseWithDefaults() *StreamsWorkspacePrivateEndpointResponse {
	this := StreamsWorkspacePrivateEndpointResponse{}
	return &this
}

// GetCustomerDnsName returns the CustomerDnsName field value if set, zero value otherwise
func (o *StreamsWorkspacePrivateEndpointResponse) GetCustomerDnsName() string {
	if o == nil || IsNil(o.CustomerDnsName) {
		var ret string
		return ret
	}
	return *o.CustomerDnsName
}

// GetCustomerDnsNameOk returns a tuple with the CustomerDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) GetCustomerDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerDnsName) {
		return nil, false
	}

	return o.CustomerDnsName, true
}

// HasCustomerDnsName returns a boolean if a field has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) HasCustomerDnsName() bool {
	if o != nil && !IsNil(o.CustomerDnsName) {
		return true
	}

	return false
}

// SetCustomerDnsName gets a reference to the given string and assigns it to the CustomerDnsName field.
func (o *StreamsWorkspacePrivateEndpointResponse) SetCustomerDnsName(v string) {
	o.CustomerDnsName = &v
	o.NullFields = removeNullField(o.NullFields, "CustomerDnsName")
}

// SetCustomerDnsNameNil sets CustomerDnsName to an explicit JSON null when marshaled.
func (o *StreamsWorkspacePrivateEndpointResponse) SetCustomerDnsNameNil() {
	o.CustomerDnsName = nil
	o.NullFields = addNullField(o.NullFields, "CustomerDnsName")
}

// GetEndpointServiceName returns the EndpointServiceName field value if set, zero value otherwise
func (o *StreamsWorkspacePrivateEndpointResponse) GetEndpointServiceName() string {
	if o == nil || IsNil(o.EndpointServiceName) {
		var ret string
		return ret
	}
	return *o.EndpointServiceName
}

// GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) GetEndpointServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointServiceName) {
		return nil, false
	}

	return o.EndpointServiceName, true
}

// HasEndpointServiceName returns a boolean if a field has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) HasEndpointServiceName() bool {
	if o != nil && !IsNil(o.EndpointServiceName) {
		return true
	}

	return false
}

// SetEndpointServiceName gets a reference to the given string and assigns it to the EndpointServiceName field.
func (o *StreamsWorkspacePrivateEndpointResponse) SetEndpointServiceName(v string) {
	o.EndpointServiceName = &v
	o.NullFields = removeNullField(o.NullFields, "EndpointServiceName")
}

// SetEndpointServiceNameNil sets EndpointServiceName to an explicit JSON null when marshaled.
func (o *StreamsWorkspacePrivateEndpointResponse) SetEndpointServiceNameNil() {
	o.EndpointServiceName = nil
	o.NullFields = addNullField(o.NullFields, "EndpointServiceName")
}

// GetId returns the Id field value if set, zero value otherwise
func (o *StreamsWorkspacePrivateEndpointResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StreamsWorkspacePrivateEndpointResponse) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *StreamsWorkspacePrivateEndpointResponse) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetProvider returns the Provider field value if set, zero value otherwise
func (o *StreamsWorkspacePrivateEndpointResponse) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}

	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *StreamsWorkspacePrivateEndpointResponse) SetProvider(v string) {
	o.Provider = &v
	o.NullFields = removeNullField(o.NullFields, "Provider")
}

// SetProviderNil sets Provider to an explicit JSON null when marshaled.
func (o *StreamsWorkspacePrivateEndpointResponse) SetProviderNil() {
	o.Provider = nil
	o.NullFields = addNullField(o.NullFields, "Provider")
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *StreamsWorkspacePrivateEndpointResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StreamsWorkspacePrivateEndpointResponse) SetRegion(v string) {
	o.Region = &v
	o.NullFields = removeNullField(o.NullFields, "Region")
}

// SetRegionNil sets Region to an explicit JSON null when marshaled.
func (o *StreamsWorkspacePrivateEndpointResponse) SetRegionNil() {
	o.Region = nil
	o.NullFields = addNullField(o.NullFields, "Region")
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *StreamsWorkspacePrivateEndpointResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StreamsWorkspacePrivateEndpointResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StreamsWorkspacePrivateEndpointResponse) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *StreamsWorkspacePrivateEndpointResponse) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}
