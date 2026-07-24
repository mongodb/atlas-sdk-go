// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OutboundControlPlaneCloudProviderIPAddresses List of outbound IP addresses from the Atlas control plane, categorized by cloud provider. If your network allows inbound HTTP requests only from specific IP addresses, you must allow access from the following IP addresses so that Atlas can communicate with your webhooks and KMS.
type OutboundControlPlaneCloudProviderIPAddresses struct {
	// Control plane IP addresses in AWS. Each key identifies an Amazon Web Services (AWS) region. Each value identifies control plane IP addresses in the AWS region.
	// Read only field.
	Aws *map[string][]string `json:"aws,omitempty"`
	// Control plane IP addresses in Azure. Each key identifies an Azure region. Each value identifies control plane IP addresses in the Azure region.
	// Read only field.
	Azure *map[string][]string `json:"azure,omitempty"`
	// Control plane IP addresses in GCP. Each key identifies a Google Cloud (GCP) region. Each value identifies control plane IP addresses in the GCP region.
	// Read only field.
	Gcp *map[string][]string `json:"gcp,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OutboundControlPlaneCloudProviderIPAddresses) MarshalJSON() ([]byte, error) {
	type noMethod OutboundControlPlaneCloudProviderIPAddresses
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOutboundControlPlaneCloudProviderIPAddresses instantiates a new OutboundControlPlaneCloudProviderIPAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundControlPlaneCloudProviderIPAddresses() *OutboundControlPlaneCloudProviderIPAddresses {
	this := OutboundControlPlaneCloudProviderIPAddresses{}
	return &this
}

// NewOutboundControlPlaneCloudProviderIPAddressesWithDefaults instantiates a new OutboundControlPlaneCloudProviderIPAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundControlPlaneCloudProviderIPAddressesWithDefaults() *OutboundControlPlaneCloudProviderIPAddresses {
	this := OutboundControlPlaneCloudProviderIPAddresses{}
	return &this
}

// GetAws returns the Aws field value if set, zero value otherwise
func (o *OutboundControlPlaneCloudProviderIPAddresses) GetAws() map[string][]string {
	if o == nil || IsNil(o.Aws) {
		var ret map[string][]string
		return ret
	}
	return *o.Aws
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundControlPlaneCloudProviderIPAddresses) GetAwsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Aws) {
		return nil, false
	}

	return o.Aws, true
}

// HasAws returns a boolean if a field has been set.
func (o *OutboundControlPlaneCloudProviderIPAddresses) HasAws() bool {
	if o != nil && !IsNil(o.Aws) {
		return true
	}

	return false
}

// SetAws gets a reference to the given map[string][]string and assigns it to the Aws field.
func (o *OutboundControlPlaneCloudProviderIPAddresses) SetAws(v map[string][]string) {
	o.Aws = &v
	o.NullFields = removeNullField(o.NullFields, "Aws")
}

// SetAwsNil sets Aws to an explicit JSON null when marshaled.
func (o *OutboundControlPlaneCloudProviderIPAddresses) SetAwsNil() {
	o.Aws = nil
	o.NullFields = addNullField(o.NullFields, "Aws")
}

// GetAzure returns the Azure field value if set, zero value otherwise
func (o *OutboundControlPlaneCloudProviderIPAddresses) GetAzure() map[string][]string {
	if o == nil || IsNil(o.Azure) {
		var ret map[string][]string
		return ret
	}
	return *o.Azure
}

// GetAzureOk returns a tuple with the Azure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundControlPlaneCloudProviderIPAddresses) GetAzureOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Azure) {
		return nil, false
	}

	return o.Azure, true
}

// HasAzure returns a boolean if a field has been set.
func (o *OutboundControlPlaneCloudProviderIPAddresses) HasAzure() bool {
	if o != nil && !IsNil(o.Azure) {
		return true
	}

	return false
}

// SetAzure gets a reference to the given map[string][]string and assigns it to the Azure field.
func (o *OutboundControlPlaneCloudProviderIPAddresses) SetAzure(v map[string][]string) {
	o.Azure = &v
	o.NullFields = removeNullField(o.NullFields, "Azure")
}

// SetAzureNil sets Azure to an explicit JSON null when marshaled.
func (o *OutboundControlPlaneCloudProviderIPAddresses) SetAzureNil() {
	o.Azure = nil
	o.NullFields = addNullField(o.NullFields, "Azure")
}

// GetGcp returns the Gcp field value if set, zero value otherwise
func (o *OutboundControlPlaneCloudProviderIPAddresses) GetGcp() map[string][]string {
	if o == nil || IsNil(o.Gcp) {
		var ret map[string][]string
		return ret
	}
	return *o.Gcp
}

// GetGcpOk returns a tuple with the Gcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundControlPlaneCloudProviderIPAddresses) GetGcpOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Gcp) {
		return nil, false
	}

	return o.Gcp, true
}

// HasGcp returns a boolean if a field has been set.
func (o *OutboundControlPlaneCloudProviderIPAddresses) HasGcp() bool {
	if o != nil && !IsNil(o.Gcp) {
		return true
	}

	return false
}

// SetGcp gets a reference to the given map[string][]string and assigns it to the Gcp field.
func (o *OutboundControlPlaneCloudProviderIPAddresses) SetGcp(v map[string][]string) {
	o.Gcp = &v
	o.NullFields = removeNullField(o.NullFields, "Gcp")
}

// SetGcpNil sets Gcp to an explicit JSON null when marshaled.
func (o *OutboundControlPlaneCloudProviderIPAddresses) SetGcpNil() {
	o.Gcp = nil
	o.NullFields = addNullField(o.NullFields, "Gcp")
}
