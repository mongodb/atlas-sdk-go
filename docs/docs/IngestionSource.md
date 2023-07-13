# IngestionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of ingestion source of this Data Lake Pipeline. | [optional] 
**ClusterName** | Pointer to **string** | Human-readable name that identifies the cluster. | [optional] 
**CollectionName** | Pointer to **string** | Human-readable name that identifies the collection. | [optional] 
**DatabaseName** | Pointer to **string** | Human-readable name that identifies the database. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**PolicyItemId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies a policy item. | [optional] 

## Methods

### NewIngestionSource

`func NewIngestionSource() *IngestionSource`

NewIngestionSource instantiates a new IngestionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSourceWithDefaults

`func NewIngestionSourceWithDefaults() *IngestionSource`

NewIngestionSourceWithDefaults instantiates a new IngestionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IngestionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IngestionSource) HasType() bool`

HasType returns a boolean if a field has been set.
### GetClusterName

`func (o *IngestionSource) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *IngestionSource) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *IngestionSource) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *IngestionSource) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetCollectionName

`func (o *IngestionSource) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *IngestionSource) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *IngestionSource) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *IngestionSource) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.
### GetDatabaseName

`func (o *IngestionSource) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *IngestionSource) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *IngestionSource) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *IngestionSource) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.
### GetGroupId

`func (o *IngestionSource) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IngestionSource) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IngestionSource) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *IngestionSource) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetPolicyItemId

`func (o *IngestionSource) GetPolicyItemId() string`

GetPolicyItemId returns the PolicyItemId field if non-nil, zero value otherwise.

### GetPolicyItemIdOk

`func (o *IngestionSource) GetPolicyItemIdOk() (*string, bool)`

GetPolicyItemIdOk returns a tuple with the PolicyItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItemId

`func (o *IngestionSource) SetPolicyItemId(v string)`

SetPolicyItemId sets PolicyItemId field to given value.

### HasPolicyItemId

`func (o *IngestionSource) HasPolicyItemId() bool`

HasPolicyItemId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


