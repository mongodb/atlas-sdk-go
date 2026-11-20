// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ObjectStoragePrivateEndpointResponse Object Storage Private Endpoint Response.
type ObjectStoragePrivateEndpointResponse struct {
	// Human-readable label that identifies the cloud provider.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Error message for failures associated with the Object Storage private endpoint.
	// Read only field.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// State of the Object Storage private endpoint.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// Connection name of the Private Endpoint.
	// Read only field.
	PrivateEndpointConnectionName *string `json:"privateEndpointConnectionName,omitempty"`
	// Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	RegionName *string `json:"regionName,omitempty"`
	// Physical location where MongoDB Cloud deploys your Azure-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. Contact MongoDB to become an early adopter.
	VpcRegionName *string `json:"vpcRegionName,omitempty"`
	// Cloud provider resource ID of the private endpoint. Returned for Azure private endpoints once the endpoint has been provisioned.
	// Read only field.
	PrivateEndpointResourceId *string `json:"privateEndpointResourceId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Cloud Provider Access role used to provision the private endpoint and re-validate storage account access. Returned for Azure private endpoints only.
	RoleId *string `json:"roleId,omitempty"`
	// URL of the Azure Storage Account for the private endpoint, exactly as supplied at creation. Returned for Azure private endpoints only.
	ServiceUrl *string `json:"serviceUrl,omitempty"`
	// Name of the Azure Storage Account the private endpoint is attached to. Returned for Azure private endpoints only.
	// Read only field.
	StorageAccountName *string `json:"storageAccountName,omitempty"`
	// ARM resource ID of the Azure Storage Account that the private endpoint connects to. Returned for Azure private endpoints only.
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ObjectStoragePrivateEndpointResponse) MarshalJSON() ([]byte, error) {
	type noMethod ObjectStoragePrivateEndpointResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewObjectStoragePrivateEndpointResponse instantiates a new ObjectStoragePrivateEndpointResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectStoragePrivateEndpointResponse() *ObjectStoragePrivateEndpointResponse {
	this := ObjectStoragePrivateEndpointResponse{}
	return &this
}

// NewObjectStoragePrivateEndpointResponseWithDefaults instantiates a new ObjectStoragePrivateEndpointResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectStoragePrivateEndpointResponseWithDefaults() *ObjectStoragePrivateEndpointResponse {
	this := ObjectStoragePrivateEndpointResponse{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ObjectStoragePrivateEndpointResponse) SetCloudProvider(v string) {
	o.CloudProvider = &v
	o.NullFields = removeNullField(o.NullFields, "CloudProvider")
}

// SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetCloudProviderNil() {
	o.CloudProvider = nil
	o.NullFields = addNullField(o.NullFields, "CloudProvider")
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ObjectStoragePrivateEndpointResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
	o.NullFields = removeNullField(o.NullFields, "ErrorMessage")
}

// SetErrorMessageNil sets ErrorMessage to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetErrorMessageNil() {
	o.ErrorMessage = nil
	o.NullFields = addNullField(o.NullFields, "ErrorMessage")
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectStoragePrivateEndpointResponse) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ObjectStoragePrivateEndpointResponse) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}

// GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetPrivateEndpointConnectionName() string {
	if o == nil || IsNil(o.PrivateEndpointConnectionName) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointConnectionName
}

// GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetPrivateEndpointConnectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointConnectionName) {
		return nil, false
	}

	return o.PrivateEndpointConnectionName, true
}

// HasPrivateEndpointConnectionName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasPrivateEndpointConnectionName() bool {
	if o != nil && !IsNil(o.PrivateEndpointConnectionName) {
		return true
	}

	return false
}

// SetPrivateEndpointConnectionName gets a reference to the given string and assigns it to the PrivateEndpointConnectionName field.
func (o *ObjectStoragePrivateEndpointResponse) SetPrivateEndpointConnectionName(v string) {
	o.PrivateEndpointConnectionName = &v
	o.NullFields = removeNullField(o.NullFields, "PrivateEndpointConnectionName")
}

// SetPrivateEndpointConnectionNameNil sets PrivateEndpointConnectionName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetPrivateEndpointConnectionNameNil() {
	o.PrivateEndpointConnectionName = nil
	o.NullFields = addNullField(o.NullFields, "PrivateEndpointConnectionName")
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ObjectStoragePrivateEndpointResponse) SetRegionName(v string) {
	o.RegionName = &v
	o.NullFields = removeNullField(o.NullFields, "RegionName")
}

// SetRegionNameNil sets RegionName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetRegionNameNil() {
	o.RegionName = nil
	o.NullFields = addNullField(o.NullFields, "RegionName")
}

// GetVpcRegionName returns the VpcRegionName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetVpcRegionName() string {
	if o == nil || IsNil(o.VpcRegionName) {
		var ret string
		return ret
	}
	return *o.VpcRegionName
}

// GetVpcRegionNameOk returns a tuple with the VpcRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetVpcRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.VpcRegionName) {
		return nil, false
	}

	return o.VpcRegionName, true
}

// HasVpcRegionName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasVpcRegionName() bool {
	if o != nil && !IsNil(o.VpcRegionName) {
		return true
	}

	return false
}

// SetVpcRegionName gets a reference to the given string and assigns it to the VpcRegionName field.
func (o *ObjectStoragePrivateEndpointResponse) SetVpcRegionName(v string) {
	o.VpcRegionName = &v
	o.NullFields = removeNullField(o.NullFields, "VpcRegionName")
}

// SetVpcRegionNameNil sets VpcRegionName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetVpcRegionNameNil() {
	o.VpcRegionName = nil
	o.NullFields = addNullField(o.NullFields, "VpcRegionName")
}

// GetPrivateEndpointResourceId returns the PrivateEndpointResourceId field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetPrivateEndpointResourceId() string {
	if o == nil || IsNil(o.PrivateEndpointResourceId) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointResourceId
}

// GetPrivateEndpointResourceIdOk returns a tuple with the PrivateEndpointResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetPrivateEndpointResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointResourceId) {
		return nil, false
	}

	return o.PrivateEndpointResourceId, true
}

// HasPrivateEndpointResourceId returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasPrivateEndpointResourceId() bool {
	if o != nil && !IsNil(o.PrivateEndpointResourceId) {
		return true
	}

	return false
}

// SetPrivateEndpointResourceId gets a reference to the given string and assigns it to the PrivateEndpointResourceId field.
func (o *ObjectStoragePrivateEndpointResponse) SetPrivateEndpointResourceId(v string) {
	o.PrivateEndpointResourceId = &v
	o.NullFields = removeNullField(o.NullFields, "PrivateEndpointResourceId")
}

// SetPrivateEndpointResourceIdNil sets PrivateEndpointResourceId to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetPrivateEndpointResourceIdNil() {
	o.PrivateEndpointResourceId = nil
	o.NullFields = addNullField(o.NullFields, "PrivateEndpointResourceId")
}

// GetRoleId returns the RoleId field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}

	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *ObjectStoragePrivateEndpointResponse) SetRoleId(v string) {
	o.RoleId = &v
	o.NullFields = removeNullField(o.NullFields, "RoleId")
}

// SetRoleIdNil sets RoleId to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetRoleIdNil() {
	o.RoleId = nil
	o.NullFields = addNullField(o.NullFields, "RoleId")
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetServiceUrl() string {
	if o == nil || IsNil(o.ServiceUrl) {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUrl) {
		return nil, false
	}

	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasServiceUrl() bool {
	if o != nil && !IsNil(o.ServiceUrl) {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *ObjectStoragePrivateEndpointResponse) SetServiceUrl(v string) {
	o.ServiceUrl = &v
	o.NullFields = removeNullField(o.NullFields, "ServiceUrl")
}

// SetServiceUrlNil sets ServiceUrl to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetServiceUrlNil() {
	o.ServiceUrl = nil
	o.NullFields = addNullField(o.NullFields, "ServiceUrl")
}

// GetStorageAccountName returns the StorageAccountName field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetStorageAccountName() string {
	if o == nil || IsNil(o.StorageAccountName) {
		var ret string
		return ret
	}
	return *o.StorageAccountName
}

// GetStorageAccountNameOk returns a tuple with the StorageAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetStorageAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.StorageAccountName) {
		return nil, false
	}

	return o.StorageAccountName, true
}

// HasStorageAccountName returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasStorageAccountName() bool {
	if o != nil && !IsNil(o.StorageAccountName) {
		return true
	}

	return false
}

// SetStorageAccountName gets a reference to the given string and assigns it to the StorageAccountName field.
func (o *ObjectStoragePrivateEndpointResponse) SetStorageAccountName(v string) {
	o.StorageAccountName = &v
	o.NullFields = removeNullField(o.NullFields, "StorageAccountName")
}

// SetStorageAccountNameNil sets StorageAccountName to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetStorageAccountNameNil() {
	o.StorageAccountName = nil
	o.NullFields = addNullField(o.NullFields, "StorageAccountName")
}

// GetStorageAccountResourceId returns the StorageAccountResourceId field value if set, zero value otherwise
func (o *ObjectStoragePrivateEndpointResponse) GetStorageAccountResourceId() string {
	if o == nil || IsNil(o.StorageAccountResourceId) {
		var ret string
		return ret
	}
	return *o.StorageAccountResourceId
}

// GetStorageAccountResourceIdOk returns a tuple with the StorageAccountResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectStoragePrivateEndpointResponse) GetStorageAccountResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.StorageAccountResourceId) {
		return nil, false
	}

	return o.StorageAccountResourceId, true
}

// HasStorageAccountResourceId returns a boolean if a field has been set.
func (o *ObjectStoragePrivateEndpointResponse) HasStorageAccountResourceId() bool {
	if o != nil && !IsNil(o.StorageAccountResourceId) {
		return true
	}

	return false
}

// SetStorageAccountResourceId gets a reference to the given string and assigns it to the StorageAccountResourceId field.
func (o *ObjectStoragePrivateEndpointResponse) SetStorageAccountResourceId(v string) {
	o.StorageAccountResourceId = &v
	o.NullFields = removeNullField(o.NullFields, "StorageAccountResourceId")
}

// SetStorageAccountResourceIdNil sets StorageAccountResourceId to an explicit JSON null when marshaled.
func (o *ObjectStoragePrivateEndpointResponse) SetStorageAccountResourceIdNil() {
	o.StorageAccountResourceId = nil
	o.NullFields = addNullField(o.NullFields, "StorageAccountResourceId")
}
