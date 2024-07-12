# StreamsDLQ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coll** | Pointer to **string** | Name of the collection that will be used for the DLQ. | [optional] 
**ConnectionName** | Pointer to **string** | Connection name that will be used to write DLQ messages to. Has to be an Atlas connection. | [optional] 
**Db** | Pointer to **string** | Name of the database that will be used for the DLQ. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewStreamsDLQ

`func NewStreamsDLQ() *StreamsDLQ`

NewStreamsDLQ instantiates a new StreamsDLQ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsDLQWithDefaults

`func NewStreamsDLQWithDefaults() *StreamsDLQ`

NewStreamsDLQWithDefaults instantiates a new StreamsDLQ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColl

`func (o *StreamsDLQ) GetColl() string`

GetColl returns the Coll field if non-nil, zero value otherwise.

### GetCollOk

`func (o *StreamsDLQ) GetCollOk() (*string, bool)`

GetCollOk returns a tuple with the Coll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColl

`func (o *StreamsDLQ) SetColl(v string)`

SetColl sets Coll field to given value.

### HasColl

`func (o *StreamsDLQ) HasColl() bool`

HasColl returns a boolean if a field has been set.
### GetConnectionName

`func (o *StreamsDLQ) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *StreamsDLQ) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *StreamsDLQ) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *StreamsDLQ) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.
### GetDb

`func (o *StreamsDLQ) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *StreamsDLQ) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *StreamsDLQ) SetDb(v string)`

SetDb sets Db field to given value.

### HasDb

`func (o *StreamsDLQ) HasDb() bool`

HasDb returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsDLQ) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsDLQ) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsDLQ) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsDLQ) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


