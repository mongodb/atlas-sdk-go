# Validation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the validation. | [optional] [readonly] 
**ErrorMessage** | Pointer to **NullableString** | Reason why the validation job failed. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to validate. | [optional] [readonly] 
**SourceGroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the source project. | [optional] 
**Status** | Pointer to **NullableString** | State of the specified validation job returned at the time of the request. | [optional] [readonly] 

## Methods

### NewValidation

`func NewValidation() *Validation`

NewValidation instantiates a new Validation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationWithDefaults

`func NewValidationWithDefaults() *Validation`

NewValidationWithDefaults instantiates a new Validation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Validation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Validation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Validation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Validation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Validation) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Validation) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Validation) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Validation) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Validation) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Validation) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetGroupId

`func (o *Validation) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Validation) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Validation) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Validation) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSourceGroupId

`func (o *Validation) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *Validation) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *Validation) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *Validation) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### GetStatus

`func (o *Validation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Validation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Validation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Validation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Validation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Validation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


