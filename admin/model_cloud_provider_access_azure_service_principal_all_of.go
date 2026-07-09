// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// CloudProviderAccessAzureServicePrincipalAllOf struct for CloudProviderAccessAzureServicePrincipalAllOf
type CloudProviderAccessAzureServicePrincipalAllOf struct {
	// Unique 24-hexadecimal digit string that identifies the role.
	// Read only field.
	Id *string `json:"_id,omitempty"`
	// Azure Active Directory Application ID of Atlas. This field is optional and will be derived from the Azure subscription if not provided.
	AtlasAzureAppId *string `json:"atlasAzureAppId,omitempty"`
	// Date and time when this Azure Service Principal was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// List that contains application features associated with this Azure Service Principal.
	// Read only field.
	FeatureUsages *[]CloudProviderAccessFeatureUsage `json:"featureUsages,omitempty"`
	// Date and time when this Azure Service Principal was last updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// UUID string that identifies the Azure Service Principal.
	ServicePrincipalId *string `json:"servicePrincipalId,omitempty"`
	// UUID String that identifies the Azure Active Directory Tenant ID.
	TenantId *string `json:"tenantId,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) MarshalJSON() ([]byte, error) {
	type noMethod CloudProviderAccessAzureServicePrincipalAllOf
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewCloudProviderAccessAzureServicePrincipalAllOf instantiates a new CloudProviderAccessAzureServicePrincipalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessAzureServicePrincipalAllOf() *CloudProviderAccessAzureServicePrincipalAllOf {
	this := CloudProviderAccessAzureServicePrincipalAllOf{}
	return &this
}

// NewCloudProviderAccessAzureServicePrincipalAllOfWithDefaults instantiates a new CloudProviderAccessAzureServicePrincipalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessAzureServicePrincipalAllOfWithDefaults() *CloudProviderAccessAzureServicePrincipalAllOf {
	this := CloudProviderAccessAzureServicePrincipalAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetId(v string) {
	o.Id = &v
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetIdNil() {
	o.Id = nil
	o.NullFields = append(o.NullFields, "Id")
}

// GetAtlasAzureAppId returns the AtlasAzureAppId field value if set, zero value otherwise
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetAtlasAzureAppId() string {
	if o == nil || IsNil(o.AtlasAzureAppId) {
		var ret string
		return ret
	}
	return *o.AtlasAzureAppId
}

// GetAtlasAzureAppIdOk returns a tuple with the AtlasAzureAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetAtlasAzureAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasAzureAppId) {
		return nil, false
	}

	return o.AtlasAzureAppId, true
}

// HasAtlasAzureAppId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasAtlasAzureAppId() bool {
	if o != nil && !IsNil(o.AtlasAzureAppId) {
		return true
	}

	return false
}

// SetAtlasAzureAppId gets a reference to the given string and assigns it to the AtlasAzureAppId field.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetAtlasAzureAppId(v string) {
	o.AtlasAzureAppId = &v
}

// SetAtlasAzureAppIdNil sets AtlasAzureAppId to an explicit JSON null when marshaled.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetAtlasAzureAppIdNil() {
	o.AtlasAzureAppId = nil
	o.NullFields = append(o.NullFields, "AtlasAzureAppId")
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}

	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// SetCreatedDateNil sets CreatedDate to an explicit JSON null when marshaled.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetCreatedDateNil() {
	o.CreatedDate = nil
	o.NullFields = append(o.NullFields, "CreatedDate")
}

// GetFeatureUsages returns the FeatureUsages field value if set, zero value otherwise
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetFeatureUsages() []CloudProviderAccessFeatureUsage {
	if o == nil || IsNil(o.FeatureUsages) {
		var ret []CloudProviderAccessFeatureUsage
		return ret
	}
	return *o.FeatureUsages
}

// GetFeatureUsagesOk returns a tuple with the FeatureUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool) {
	if o == nil || IsNil(o.FeatureUsages) {
		return nil, false
	}

	return o.FeatureUsages, true
}

// HasFeatureUsages returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasFeatureUsages() bool {
	if o != nil && !IsNil(o.FeatureUsages) {
		return true
	}

	return false
}

// SetFeatureUsages gets a reference to the given []CloudProviderAccessFeatureUsage and assigns it to the FeatureUsages field.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetFeatureUsages(v []CloudProviderAccessFeatureUsage) {
	o.FeatureUsages = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}

	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// SetLastUpdatedDateNil sets LastUpdatedDate to an explicit JSON null when marshaled.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetLastUpdatedDateNil() {
	o.LastUpdatedDate = nil
	o.NullFields = append(o.NullFields, "LastUpdatedDate")
}

// GetServicePrincipalId returns the ServicePrincipalId field value if set, zero value otherwise
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetServicePrincipalId() string {
	if o == nil || IsNil(o.ServicePrincipalId) {
		var ret string
		return ret
	}
	return *o.ServicePrincipalId
}

// GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetServicePrincipalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePrincipalId) {
		return nil, false
	}

	return o.ServicePrincipalId, true
}

// HasServicePrincipalId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasServicePrincipalId() bool {
	if o != nil && !IsNil(o.ServicePrincipalId) {
		return true
	}

	return false
}

// SetServicePrincipalId gets a reference to the given string and assigns it to the ServicePrincipalId field.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetServicePrincipalId(v string) {
	o.ServicePrincipalId = &v
}

// SetServicePrincipalIdNil sets ServicePrincipalId to an explicit JSON null when marshaled.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetServicePrincipalIdNil() {
	o.ServicePrincipalId = nil
	o.NullFields = append(o.NullFields, "ServicePrincipalId")
}

// GetTenantId returns the TenantId field value if set, zero value otherwise
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}

	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetTenantId(v string) {
	o.TenantId = &v
}

// SetTenantIdNil sets TenantId to an explicit JSON null when marshaled.
func (o *CloudProviderAccessAzureServicePrincipalAllOf) SetTenantIdNil() {
	o.TenantId = nil
	o.NullFields = append(o.NullFields, "TenantId")
}
