# DiskPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**PartitionName** | Pointer to **string** | Human-readable label of the disk or partition to which the measurements apply. | [optional] [readonly] 

## Methods

### NewDiskPartition

`func NewDiskPartition() *DiskPartition`

NewDiskPartition instantiates a new DiskPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskPartitionWithDefaults

`func NewDiskPartitionWithDefaults() *DiskPartition`

NewDiskPartitionWithDefaults instantiates a new DiskPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DiskPartition) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskPartition) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskPartition) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskPartition) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPartitionName

`func (o *DiskPartition) GetPartitionName() string`

GetPartitionName returns the PartitionName field if non-nil, zero value otherwise.

### GetPartitionNameOk

`func (o *DiskPartition) GetPartitionNameOk() (*string, bool)`

GetPartitionNameOk returns a tuple with the PartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionName

`func (o *DiskPartition) SetPartitionName(v string)`

SetPartitionName sets PartitionName field to given value.

### HasPartitionName

`func (o *DiskPartition) HasPartitionName() bool`

HasPartitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


