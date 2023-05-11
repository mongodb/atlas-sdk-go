# ExampleResourceResponseView20230101

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | Dummy data added as response. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewExampleResourceResponseView20230101

`func NewExampleResourceResponseView20230101(data string, ) *ExampleResourceResponseView20230101`

NewExampleResourceResponseView20230101 instantiates a new ExampleResourceResponseView20230101 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExampleResourceResponseView20230101WithDefaults

`func NewExampleResourceResponseView20230101WithDefaults() *ExampleResourceResponseView20230101`

NewExampleResourceResponseView20230101WithDefaults instantiates a new ExampleResourceResponseView20230101 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExampleResourceResponseView20230101) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExampleResourceResponseView20230101) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExampleResourceResponseView20230101) SetData(v string)`

SetData sets Data field to given value.


### GetLinks

`func (o *ExampleResourceResponseView20230101) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExampleResourceResponseView20230101) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExampleResourceResponseView20230101) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExampleResourceResponseView20230101) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


