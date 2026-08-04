# Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of this principal. | [optional] 
**Name** | Pointer to **string** | The human-readable name of this principal. | [optional] 
**OnBehalfOf** | Pointer to [**Principal**](Principal.md) |  | [optional] 

## Methods

### NewPrincipal

`func NewPrincipal() *Principal`

NewPrincipal instantiates a new Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithDefaults

`func NewPrincipalWithDefaults() *Principal`

NewPrincipalWithDefaults instantiates a new Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Principal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Principal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Principal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Principal) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Principal) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetName

`func (o *Principal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Principal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Principal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Principal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Principal) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.

### GetOnBehalfOf

`func (o *Principal) GetOnBehalfOf() Principal`

GetOnBehalfOf returns the OnBehalfOf field if non-nil, zero value otherwise.

### GetOnBehalfOfOk

`func (o *Principal) GetOnBehalfOfOk() (*Principal, bool)`

GetOnBehalfOfOk returns a tuple with the OnBehalfOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBehalfOf

`func (o *Principal) SetOnBehalfOf(v Principal)`

SetOnBehalfOf sets OnBehalfOf field to given value.

### HasOnBehalfOf

`func (o *Principal) HasOnBehalfOf() bool`

HasOnBehalfOf returns a boolean if a field has been set.

### SetOnBehalfOfNil

`func (o *Principal) SetOnBehalfOfNil()`

SetOnBehalfOfNil sets OnBehalfOf to an explicit JSON null when marshaled, overriding any value previously set with SetOnBehalfOf. Calling SetOnBehalfOf again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


