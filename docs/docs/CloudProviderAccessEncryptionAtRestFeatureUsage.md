# CloudProviderAccessEncryptionAtRestFeatureUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureId** | Pointer to **map[string]interface{}** | Object that contains the identifying characteristics of the Amazon Web Services (AWS) Key Management Service (KMS). This field always returns a null value. | [optional] 
**FeatureType** | Pointer to **string** | Human-readable label that describes one MongoDB Cloud feature linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role. | [optional] [readonly] 

## Methods

### NewCloudProviderAccessEncryptionAtRestFeatureUsage

`func NewCloudProviderAccessEncryptionAtRestFeatureUsage() *CloudProviderAccessEncryptionAtRestFeatureUsage`

NewCloudProviderAccessEncryptionAtRestFeatureUsage instantiates a new CloudProviderAccessEncryptionAtRestFeatureUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessEncryptionAtRestFeatureUsageWithDefaults

`func NewCloudProviderAccessEncryptionAtRestFeatureUsageWithDefaults() *CloudProviderAccessEncryptionAtRestFeatureUsage`

NewCloudProviderAccessEncryptionAtRestFeatureUsageWithDefaults instantiates a new CloudProviderAccessEncryptionAtRestFeatureUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureId

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureId() map[string]interface{}`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureIdOk() (*map[string]interface{}, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) SetFeatureId(v map[string]interface{})`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### SetFeatureIdNil

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) SetFeatureIdNil(b bool)`

 SetFeatureIdNil sets the value for FeatureId to be an explicit nil

### UnsetFeatureId
`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) UnsetFeatureId()`

UnsetFeatureId ensures that no value is present for FeatureId, not even an explicit nil
### GetFeatureType

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


