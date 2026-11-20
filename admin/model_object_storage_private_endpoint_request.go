// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ObjectStoragePrivateEndpointRequest Object Storage Private Endpoint Request.
type ObjectStoragePrivateEndpointRequest struct {
	// Human-readable label that identifies the cloud provider.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	RegionName *string `json:"regionName,omitempty"`
	// Physical location where MongoDB Cloud deploys your Azure-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. Contact MongoDB to become an early adopter.
	VpcRegionName *string `json:"vpcRegionName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Azure Cloud Provider Access Service Principal that MongoDB Cloud uses to access the Azure Storage Account. This is a Cloud Provider Access role ID from your project's Cloud Provider Access configuration (returned by the Cloud Provider Access roles API). Enabled per project via a feature flag; see the Azure PrivateLink exports documentation for early access. Contact MongoDB to become an early adopter.
	RoleId *string `json:"roleId,omitempty"`
	// URL of the Azure Storage Account for the private endpoint. For example: `https://examplestorageaccount.blob.core.windows.net/exportcontainer`. Only standard endpoints (with `blob.core.windows.net`) are supported. Enabled per project via a feature flag; see the Azure PrivateLink exports documentation for early access. Contact MongoDB to become an early adopter.
	ServiceUrl *string `json:"serviceUrl,omitempty"`
	// ARM resource ID of the Azure Storage Account that the private endpoint connects to. Optional: provide it when the Cloud Provider Access service principal is scoped to the storage account, so MongoDB Cloud reads the account directly by resource ID instead of enumerating subscriptions. Must reference the same storage account as `serviceUrl`. Enabled per project via a feature flag; see the Azure PrivateLink exports documentation for early access. Contact MongoDB to become an early adopter.
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ObjectStoragePrivateEndpointRequest) MarshalJSON() ([]byte, error) {
	type noMethod ObjectStoragePrivateEndpointRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewObjectStoragePrivateEndpointRequest instantiates a new ObjectStoragePrivateEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePrivateEndpointRequest() *ObjectStoragePrivateEndpointRequest {
	this := ObjectStoragePrivateEndpointRequest{}
	return &this
}

// NewObjectStoragePrivateEndpointRequestWithDefaults instantiates a new ObjectStoragePrivateEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePrivateEndpointRequestWithDefaults() *ObjectStoragePrivateEndpointRequest {
	this := ObjectStoragePrivateEndpointRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ObjectStoragePrivateEndpointRequest) SetCloudProvider(v string) {
	o.CloudProvider = &v
	o.NullFields = removeNullField(o.NullFields, "CloudProvider")
}

// SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetCloudProviderNil() {
	o.CloudProvider = nil
	o.NullFields = addNullField(o.NullFields, "CloudProvider")
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ObjectStoragePrivateEndpointRequest) SetRegionName(v string) {
	o.RegionName = &v
	o.NullFields = removeNullField(o.NullFields, "RegionName")
}

// SetRegionNameNil sets RegionName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetRegionNameNil() {
	o.RegionName = nil
	o.NullFields = addNullField(o.NullFields, "RegionName")
}

// GetVpcRegionName returns the VpcRegionName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetVpcRegionName() string {
	if o == nil || IsNil(o.VpcRegionName) {
		var ret string
		return ret
	}
	return *o.VpcRegionName
}

// GetVpcRegionNameOk returns a tuple with the VpcRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetVpcRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.VpcRegionName) {
		return nil, false
	}

	return o.VpcRegionName, true
}

// HasVpcRegionName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasVpcRegionName() bool {
	if o != nil && !IsNil(o.VpcRegionName) {
		return true
	}

	return false
}

// SetVpcRegionName gets a reference to the given string and assigns it to the VpcRegionName field.
func (o *ObjectStoragePrivateEndpointRequest) SetVpcRegionName(v string) {
	o.VpcRegionName = &v
	o.NullFields = removeNullField(o.NullFields, "VpcRegionName")
}

// SetVpcRegionNameNil sets VpcRegionName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetVpcRegionNameNil() {
	o.VpcRegionName = nil
	o.NullFields = addNullField(o.NullFields, "VpcRegionName")
}

// GetRoleId returns the RoleId field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}

	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *ObjectStoragePrivateEndpointRequest) SetRoleId(v string) {
	o.RoleId = &v
	o.NullFields = removeNullField(o.NullFields, "RoleId")
}

// SetRoleIdNil sets RoleId to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetRoleIdNil() {
	o.RoleId = nil
	o.NullFields = addNullField(o.NullFields, "RoleId")
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetServiceUrl() string {
	if o == nil || IsNil(o.ServiceUrl) {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUrl) {
		return nil, false
	}

	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasServiceUrl() bool {
	if o != nil && !IsNil(o.ServiceUrl) {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *ObjectStoragePrivateEndpointRequest) SetServiceUrl(v string) {
	o.ServiceUrl = &v
	o.NullFields = removeNullField(o.NullFields, "ServiceUrl")
}

// SetServiceUrlNil sets ServiceUrl to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetServiceUrlNil() {
	o.ServiceUrl = nil
	o.NullFields = addNullField(o.NullFields, "ServiceUrl")
}

// GetStorageAccountResourceId returns the StorageAccountResourceId field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointRequest) GetStorageAccountResourceId() string {
	if o == nil || IsNil(o.StorageAccountResourceId) {
		var ret string
		return ret
	}
	return *o.StorageAccountResourceId
}

// GetStorageAccountResourceIdOk returns a tuple with the StorageAccountResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointRequest) GetStorageAccountResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.StorageAccountResourceId) {
		return nil, false
	}

	return o.StorageAccountResourceId, true
}

// HasStorageAccountResourceId returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointRequest) HasStorageAccountResourceId() bool {
	if o != nil && !IsNil(o.StorageAccountResourceId) {
		return true
	}

	return false
}

// SetStorageAccountResourceId gets a reference to the given string and assigns it to the StorageAccountResourceId field.
func (o *ObjectStoragePrivateEndpointRequest) SetStorageAccountResourceId(v string) {
	o.StorageAccountResourceId = &v
	o.NullFields = removeNullField(o.NullFields, "StorageAccountResourceId")
}

// SetStorageAccountResourceIdNil sets StorageAccountResourceId to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointRequest) SetStorageAccountResourceIdNil() {
	o.StorageAccountResourceId = nil
	o.NullFields = addNullField(o.NullFields, "StorageAccountResourceId")
}
