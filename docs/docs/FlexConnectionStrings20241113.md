# FlexConnectionStrings20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | Pointer to **string** | Public connection string that you can use to connect to this cluster. This connection string uses the &#x60;mongodb://&#x60; protocol. | [optional] [readonly] 
**StandardSrv** | Pointer to **string** | Public connection string that you can use to connect to this flex cluster. This connection string uses the &#x60;mongodb+srv://&#x60; protocol. | [optional] [readonly] 

## Methods

### NewFlexConnectionStrings20241113

`func NewFlexConnectionStrings20241113() *FlexConnectionStrings20241113`

NewFlexConnectionStrings20241113 instantiates a new FlexConnectionStrings20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexConnectionStrings20241113WithDefaults

`func NewFlexConnectionStrings20241113WithDefaults() *FlexConnectionStrings20241113`

NewFlexConnectionStrings20241113WithDefaults instantiates a new FlexConnectionStrings20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *FlexConnectionStrings20241113) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *FlexConnectionStrings20241113) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *FlexConnectionStrings20241113) SetStandard(v string)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *FlexConnectionStrings20241113) HasStandard() bool`

HasStandard returns a boolean if a field has been set.
### GetStandardSrv

`func (o *FlexConnectionStrings20241113) GetStandardSrv() string`

GetStandardSrv returns the StandardSrv field if non-nil, zero value otherwise.

### GetStandardSrvOk

`func (o *FlexConnectionStrings20241113) GetStandardSrvOk() (*string, bool)`

GetStandardSrvOk returns a tuple with the StandardSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardSrv

`func (o *FlexConnectionStrings20241113) SetStandardSrv(v string)`

SetStandardSrv sets StandardSrv field to given value.

### HasStandardSrv

`func (o *FlexConnectionStrings20241113) HasStandardSrv() bool`

HasStandardSrv returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


