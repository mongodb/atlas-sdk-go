# CloudProviderAccessFeatureUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureId** | Pointer to [**CloudProviderAccessFeatureUsageExportSnapshotFeatureId**](CloudProviderAccessFeatureUsageExportSnapshotFeatureId.md) |  | [optional] 
**FeatureType** | Pointer to **string** | Human-readable label that describes one MongoDB Cloud feature linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role. | [optional] [readonly] 

## Methods

### NewCloudProviderAccessFeatureUsage

`func NewCloudProviderAccessFeatureUsage() *CloudProviderAccessFeatureUsage`

NewCloudProviderAccessFeatureUsage instantiates a new CloudProviderAccessFeatureUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessFeatureUsageWithDefaults

`func NewCloudProviderAccessFeatureUsageWithDefaults() *CloudProviderAccessFeatureUsage`

NewCloudProviderAccessFeatureUsageWithDefaults instantiates a new CloudProviderAccessFeatureUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureId

`func (o *CloudProviderAccessFeatureUsage) GetFeatureId() CloudProviderAccessFeatureUsageExportSnapshotFeatureId`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CloudProviderAccessFeatureUsage) GetFeatureIdOk() (*CloudProviderAccessFeatureUsageExportSnapshotFeatureId, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CloudProviderAccessFeatureUsage) SetFeatureId(v CloudProviderAccessFeatureUsageExportSnapshotFeatureId)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CloudProviderAccessFeatureUsage) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetFeatureType

`func (o *CloudProviderAccessFeatureUsage) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *CloudProviderAccessFeatureUsage) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *CloudProviderAccessFeatureUsage) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *CloudProviderAccessFeatureUsage) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


