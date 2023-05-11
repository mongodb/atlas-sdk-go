# ExampleResourceResponseView20230201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to **string** | Dummy additional field added to the response. | [optional] 
**Data** | Pointer to **[]string** | Array that contains the dummy metadata. | [optional] 
**Description** | **string** | Dummy description added as response. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewExampleResourceResponseView20230201

`func NewExampleResourceResponseView20230201(description string, ) *ExampleResourceResponseView20230201`

NewExampleResourceResponseView20230201 instantiates a new ExampleResourceResponseView20230201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExampleResourceResponseView20230201WithDefaults

`func NewExampleResourceResponseView20230201WithDefaults() *ExampleResourceResponseView20230201`

NewExampleResourceResponseView20230201WithDefaults instantiates a new ExampleResourceResponseView20230201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *ExampleResourceResponseView20230201) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ExampleResourceResponseView20230201) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ExampleResourceResponseView20230201) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ExampleResourceResponseView20230201) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetData

`func (o *ExampleResourceResponseView20230201) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExampleResourceResponseView20230201) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExampleResourceResponseView20230201) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *ExampleResourceResponseView20230201) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *ExampleResourceResponseView20230201) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExampleResourceResponseView20230201) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExampleResourceResponseView20230201) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLinks

`func (o *ExampleResourceResponseView20230201) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExampleResourceResponseView20230201) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExampleResourceResponseView20230201) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExampleResourceResponseView20230201) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


