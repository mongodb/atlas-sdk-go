# CloudProviderAccessGCPServiceAccountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDate** | Pointer to **time.Time** | Date and time when this Google Service Account was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FeatureUsages** | Pointer to [**[]CloudProviderAccessFeatureUsage**](CloudProviderAccessFeatureUsage.md) | List that contains application features associated with this Google Service Account. | [optional] [readonly] 
**GcpServiceAccountForAtlas** | Pointer to **string** | Email address for the Google Service Account created by Atlas. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the role. | [optional] [readonly] 
**Status** | Pointer to **string** | Provision status of the service account. | [optional] [readonly] 

## Methods

### NewCloudProviderAccessGCPServiceAccountAllOf

`func NewCloudProviderAccessGCPServiceAccountAllOf() *CloudProviderAccessGCPServiceAccountAllOf`

NewCloudProviderAccessGCPServiceAccountAllOf instantiates a new CloudProviderAccessGCPServiceAccountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessGCPServiceAccountAllOfWithDefaults

`func NewCloudProviderAccessGCPServiceAccountAllOfWithDefaults() *CloudProviderAccessGCPServiceAccountAllOf`

NewCloudProviderAccessGCPServiceAccountAllOfWithDefaults instantiates a new CloudProviderAccessGCPServiceAccountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CloudProviderAccessGCPServiceAccountAllOf) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CloudProviderAccessGCPServiceAccountAllOf) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.
### GetFeatureUsages

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetFeatureUsages() []CloudProviderAccessFeatureUsage`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetFeatureUsagesOk() (*[]CloudProviderAccessFeatureUsage, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *CloudProviderAccessGCPServiceAccountAllOf) SetFeatureUsages(v []CloudProviderAccessFeatureUsage)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *CloudProviderAccessGCPServiceAccountAllOf) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.
### GetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetGcpServiceAccountForAtlas() string`

GetGcpServiceAccountForAtlas returns the GcpServiceAccountForAtlas field if non-nil, zero value otherwise.

### GetGcpServiceAccountForAtlasOk

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetGcpServiceAccountForAtlasOk() (*string, bool)`

GetGcpServiceAccountForAtlasOk returns a tuple with the GcpServiceAccountForAtlas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountForAtlas

`func (o *CloudProviderAccessGCPServiceAccountAllOf) SetGcpServiceAccountForAtlas(v string)`

SetGcpServiceAccountForAtlas sets GcpServiceAccountForAtlas field to given value.

### HasGcpServiceAccountForAtlas

`func (o *CloudProviderAccessGCPServiceAccountAllOf) HasGcpServiceAccountForAtlas() bool`

HasGcpServiceAccountForAtlas returns a boolean if a field has been set.
### GetRoleId

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CloudProviderAccessGCPServiceAccountAllOf) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CloudProviderAccessGCPServiceAccountAllOf) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetStatus

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudProviderAccessGCPServiceAccountAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudProviderAccessGCPServiceAccountAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudProviderAccessGCPServiceAccountAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


