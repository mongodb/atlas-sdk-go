# DLSIngestionSink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataProvider** | Pointer to **string** | Target cloud provider for this Data Lake Pipeline. | [optional] 
**MetadataRegion** | Pointer to **string** | Target cloud provider region for this Data Lake Pipeline. | [optional] 
**PartitionFields** | Pointer to [**[]PipelinesPartitionField**](PipelinesPartitionField.md) | Ordered fields used to physically organize data in the destination. | [optional] 
**Type** | Pointer to **string** | Type of ingestion destination of this Data Lake Pipeline. | [optional] [readonly] 

## Methods

### NewDLSIngestionSink

`func NewDLSIngestionSink() *DLSIngestionSink`

NewDLSIngestionSink instantiates a new DLSIngestionSink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDLSIngestionSinkWithDefaults

`func NewDLSIngestionSinkWithDefaults() *DLSIngestionSink`

NewDLSIngestionSinkWithDefaults instantiates a new DLSIngestionSink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataProvider

`func (o *DLSIngestionSink) GetMetadataProvider() string`

GetMetadataProvider returns the MetadataProvider field if non-nil, zero value otherwise.

### GetMetadataProviderOk

`func (o *DLSIngestionSink) GetMetadataProviderOk() (*string, bool)`

GetMetadataProviderOk returns a tuple with the MetadataProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProvider

`func (o *DLSIngestionSink) SetMetadataProvider(v string)`

SetMetadataProvider sets MetadataProvider field to given value.

### HasMetadataProvider

`func (o *DLSIngestionSink) HasMetadataProvider() bool`

HasMetadataProvider returns a boolean if a field has been set.

### GetMetadataRegion

`func (o *DLSIngestionSink) GetMetadataRegion() string`

GetMetadataRegion returns the MetadataRegion field if non-nil, zero value otherwise.

### GetMetadataRegionOk

`func (o *DLSIngestionSink) GetMetadataRegionOk() (*string, bool)`

GetMetadataRegionOk returns a tuple with the MetadataRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataRegion

`func (o *DLSIngestionSink) SetMetadataRegion(v string)`

SetMetadataRegion sets MetadataRegion field to given value.

### HasMetadataRegion

`func (o *DLSIngestionSink) HasMetadataRegion() bool`

HasMetadataRegion returns a boolean if a field has been set.

### GetPartitionFields

`func (o *DLSIngestionSink) GetPartitionFields() []PipelinesPartitionField`

GetPartitionFields returns the PartitionFields field if non-nil, zero value otherwise.

### GetPartitionFieldsOk

`func (o *DLSIngestionSink) GetPartitionFieldsOk() (*[]PipelinesPartitionField, bool)`

GetPartitionFieldsOk returns a tuple with the PartitionFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionFields

`func (o *DLSIngestionSink) SetPartitionFields(v []PipelinesPartitionField)`

SetPartitionFields sets PartitionFields field to given value.

### HasPartitionFields

`func (o *DLSIngestionSink) HasPartitionFields() bool`

HasPartitionFields returns a boolean if a field has been set.

### GetType

`func (o *DLSIngestionSink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DLSIngestionSink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DLSIngestionSink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DLSIngestionSink) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


