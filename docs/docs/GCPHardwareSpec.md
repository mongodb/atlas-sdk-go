# GCPHardwareSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSize** | Pointer to **string** | Hardware specification for the instance sizes in this region. Each instance size has a default storage and memory capacity. The instance size you select applies to all the data-bearing hosts in your instance size. | [optional] 
**NodeCount** | Pointer to **int** | Number of nodes of the given type for MongoDB Cloud to deploy to the region. | [optional] 

## Methods

### NewGCPHardwareSpec

`func NewGCPHardwareSpec() *GCPHardwareSpec`

NewGCPHardwareSpec instantiates a new GCPHardwareSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPHardwareSpecWithDefaults

`func NewGCPHardwareSpecWithDefaults() *GCPHardwareSpec`

NewGCPHardwareSpecWithDefaults instantiates a new GCPHardwareSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSize

`func (o *GCPHardwareSpec) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *GCPHardwareSpec) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *GCPHardwareSpec) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *GCPHardwareSpec) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### GetNodeCount

`func (o *GCPHardwareSpec) GetNodeCount() int`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *GCPHardwareSpec) GetNodeCountOk() (*int, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *GCPHardwareSpec) SetNodeCount(v int)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *GCPHardwareSpec) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


