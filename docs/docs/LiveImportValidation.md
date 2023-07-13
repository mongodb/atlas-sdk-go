# LiveImportValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the validation. | [optional] [readonly] 
**ErrorMessage** | Pointer to **string** | Reason why the validation job failed. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project to validate. | [optional] [readonly] 
**SourceGroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the source project. | [optional] 
**Status** | Pointer to **string** | State of the specified validation job returned at the time of the request. | [optional] [readonly] 

## Methods

### NewLiveImportValidation

`func NewLiveImportValidation() *LiveImportValidation`

NewLiveImportValidation instantiates a new LiveImportValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveImportValidationWithDefaults

`func NewLiveImportValidationWithDefaults() *LiveImportValidation`

NewLiveImportValidationWithDefaults instantiates a new LiveImportValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveImportValidation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveImportValidation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveImportValidation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveImportValidation) HasId() bool`

HasId returns a boolean if a field has been set.
### GetErrorMessage

`func (o *LiveImportValidation) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *LiveImportValidation) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *LiveImportValidation) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *LiveImportValidation) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetGroupId

`func (o *LiveImportValidation) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LiveImportValidation) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LiveImportValidation) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LiveImportValidation) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetSourceGroupId

`func (o *LiveImportValidation) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *LiveImportValidation) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *LiveImportValidation) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *LiveImportValidation) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.
### GetStatus

`func (o *LiveImportValidation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiveImportValidation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiveImportValidation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiveImportValidation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


