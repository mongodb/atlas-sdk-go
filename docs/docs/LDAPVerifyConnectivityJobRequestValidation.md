# LDAPVerifyConnectivityJobRequestValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Human-readable string that indicates the result of this verification test. | [optional] [readonly] 
**ValidationType** | Pointer to **string** | Human-readable label that identifies this verification test that MongoDB Cloud runs. | [optional] [readonly] 

## Methods

### NewLDAPVerifyConnectivityJobRequestValidation

`func NewLDAPVerifyConnectivityJobRequestValidation() *LDAPVerifyConnectivityJobRequestValidation`

NewLDAPVerifyConnectivityJobRequestValidation instantiates a new LDAPVerifyConnectivityJobRequestValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPVerifyConnectivityJobRequestValidationWithDefaults

`func NewLDAPVerifyConnectivityJobRequestValidationWithDefaults() *LDAPVerifyConnectivityJobRequestValidation`

NewLDAPVerifyConnectivityJobRequestValidationWithDefaults instantiates a new LDAPVerifyConnectivityJobRequestValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LDAPVerifyConnectivityJobRequestValidation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LDAPVerifyConnectivityJobRequestValidation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LDAPVerifyConnectivityJobRequestValidation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LDAPVerifyConnectivityJobRequestValidation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetValidationType

`func (o *LDAPVerifyConnectivityJobRequestValidation) GetValidationType() string`

GetValidationType returns the ValidationType field if non-nil, zero value otherwise.

### GetValidationTypeOk

`func (o *LDAPVerifyConnectivityJobRequestValidation) GetValidationTypeOk() (*string, bool)`

GetValidationTypeOk returns a tuple with the ValidationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationType

`func (o *LDAPVerifyConnectivityJobRequestValidation) SetValidationType(v string)`

SetValidationType sets ValidationType field to given value.

### HasValidationType

`func (o *LDAPVerifyConnectivityJobRequestValidation) HasValidationType() bool`

HasValidationType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


