// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// DataFederationAzureCloudProviderConfig Configuration for running Data Federation in Azure.
type DataFederationAzureCloudProviderConfig struct {
	// The App ID generated by Atlas for the Service Principal's access policy.
	// Read only field.
	AtlasAppId *string `json:"atlasAppId,omitempty"`
	// Unique identifier of the role that Data Federation can use to access the data stores.Required if specifying cloudProviderConfig.
	RoleId string `json:"roleId"`
	// The ID of the Service Principal for which there is an access policyfor Atlas to access Azure resources.
	// Read only field.
	ServicePrincipalId *string `json:"servicePrincipalId,omitempty"`
	// The Azure Active Directory / Entra ID tenant ID associated with the Service Principal.
	// Read only field.
	TenantId *string `json:"tenantId,omitempty"`
}

// NewDataFederationAzureCloudProviderConfig instantiates a new DataFederationAzureCloudProviderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataFederationAzureCloudProviderConfig(roleId string) *DataFederationAzureCloudProviderConfig {
	this := DataFederationAzureCloudProviderConfig{}
	this.RoleId = roleId
	return &this
}

// NewDataFederationAzureCloudProviderConfigWithDefaults instantiates a new DataFederationAzureCloudProviderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataFederationAzureCloudProviderConfigWithDefaults() *DataFederationAzureCloudProviderConfig {
	this := DataFederationAzureCloudProviderConfig{}
	return &this
}

// GetAtlasAppId returns the AtlasAppId field value if set, zero value otherwise
func (o *DataFederationAzureCloudProviderConfig) GetAtlasAppId() string {
	if o == nil || IsNil(o.AtlasAppId) {
		var ret string
		return ret
	}
	return *o.AtlasAppId
}

// GetAtlasAppIdOk returns a tuple with the AtlasAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationAzureCloudProviderConfig) GetAtlasAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAppId) {
		return nil, false
	}

	return o.AtlasAppId, true
}

// HasAtlasAppId returns a boolean if a field has been set.
func (o *DataFederationAzureCloudProviderConfig) HasAtlasAppId() bool {
	if o != nil && !IsNil(o.AtlasAppId) {
		return true
	}

	return false
}

// SetAtlasAppId gets a reference to the given string and assigns it to the AtlasAppId field.
func (o *DataFederationAzureCloudProviderConfig) SetAtlasAppId(v string) {
	o.AtlasAppId = &v
}

// GetRoleId returns the RoleId field value
func (o *DataFederationAzureCloudProviderConfig) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *DataFederationAzureCloudProviderConfig) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *DataFederationAzureCloudProviderConfig) SetRoleId(v string) {
	o.RoleId = v
}

// GetServicePrincipalId returns the ServicePrincipalId field value if set, zero value otherwise
func (o *DataFederationAzureCloudProviderConfig) GetServicePrincipalId() string {
	if o == nil || IsNil(o.ServicePrincipalId) {
		var ret string
		return ret
	}
	return *o.ServicePrincipalId
}

// GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationAzureCloudProviderConfig) GetServicePrincipalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePrincipalId) {
		return nil, false
	}

	return o.ServicePrincipalId, true
}

// HasServicePrincipalId returns a boolean if a field has been set.
func (o *DataFederationAzureCloudProviderConfig) HasServicePrincipalId() bool {
	if o != nil && !IsNil(o.ServicePrincipalId) {
		return true
	}

	return false
}

// SetServicePrincipalId gets a reference to the given string and assigns it to the ServicePrincipalId field.
func (o *DataFederationAzureCloudProviderConfig) SetServicePrincipalId(v string) {
	o.ServicePrincipalId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise
func (o *DataFederationAzureCloudProviderConfig) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationAzureCloudProviderConfig) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}

	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DataFederationAzureCloudProviderConfig) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DataFederationAzureCloudProviderConfig) SetTenantId(v string) {
	o.TenantId = &v
}

func (o *DataFederationAzureCloudProviderConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o *DataFederationAzureCloudProviderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roleId"] = o.RoleId
	return toSerialize, nil
}
