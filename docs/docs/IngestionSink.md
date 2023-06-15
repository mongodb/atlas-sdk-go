# IngestionSink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataProvider** | Pointer to **string** | Target cloud provider for this Data Lake Pipeline. | [optional] 
**MetadataRegion** | Pointer to **string** | Target cloud provider region for this Data Lake Pipeline. | [optional] 
**PartitionFields** | Pointer to [**[]PipelinesPartitionField**](PipelinesPartitionField.md) | Ordered fields used to physically organize data in the destination. | [optional] 
**Type** | Pointer to **string** | Type of ingestion destination of this Data Lake Pipeline. | [optional] [readonly] 

## Methods

### NewIngestionSink

`func NewIngestionSink() *IngestionSink`

NewIngestionSink instantiates a new IngestionSink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionSinkWithDefaults

`func NewIngestionSinkWithDefaults() *IngestionSink`

NewIngestionSinkWithDefaults instantiates a new IngestionSink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataProvider

`func (o *IngestionSink) GetMetadataProvider() string`

GetMetadataProvider returns the MetadataProvider field if non-nil, zero value otherwise.

### GetMetadataProviderOk

`func (o *IngestionSink) GetMetadataProviderOk() (*string, bool)`

GetMetadataProviderOk returns a tuple with the MetadataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProvider

`func (o *IngestionSink) SetMetadataProvider(v string)`

SetMetadataProvider sets MetadataProvider field to given value.

### HasMetadataProvider

`func (o *IngestionSink) HasMetadataProvider() bool`

HasMetadataProvider returns a boolean if a field has been set.

### GetMetadataRegion

`func (o *IngestionSink) GetMetadataRegion() string`

GetMetadataRegion returns the MetadataRegion field if non-nil, zero value otherwise.

### GetMetadataRegionOk

`func (o *IngestionSink) GetMetadataRegionOk() (*string, bool)`

GetMetadataRegionOk returns a tuple with the MetadataRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRegion

`func (o *IngestionSink) SetMetadataRegion(v string)`

SetMetadataRegion sets MetadataRegion field to given value.

### HasMetadataRegion

`func (o *IngestionSink) HasMetadataRegion() bool`

HasMetadataRegion returns a boolean if a field has been set.

### GetPartitionFields

`func (o *IngestionSink) GetPartitionFields() []PipelinesPartitionField`

GetPartitionFields returns the PartitionFields field if non-nil, zero value otherwise.

### GetPartitionFieldsOk

`func (o *IngestionSink) GetPartitionFieldsOk() (*[]PipelinesPartitionField, bool)`

GetPartitionFieldsOk returns a tuple with the PartitionFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionFields

`func (o *IngestionSink) SetPartitionFields(v []PipelinesPartitionField)`

SetPartitionFields sets PartitionFields field to given value.

### HasPartitionFields

`func (o *IngestionSink) HasPartitionFields() bool`

HasPartitionFields returns a boolean if a field has been set.

### GetType

`func (o *IngestionSink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IngestionSink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IngestionSink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IngestionSink) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


