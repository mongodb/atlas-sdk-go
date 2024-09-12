# ApiAtlasPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | A string that defines the permissions for the policy. The syntax used is the Cedar Policy language. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the policy. | [optional] [readonly] 

## Methods

### NewApiAtlasPolicy

`func NewApiAtlasPolicy() *ApiAtlasPolicy`

NewApiAtlasPolicy instantiates a new ApiAtlasPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasPolicyWithDefaults

`func NewApiAtlasPolicyWithDefaults() *ApiAtlasPolicy`

NewApiAtlasPolicyWithDefaults instantiates a new ApiAtlasPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ApiAtlasPolicy) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ApiAtlasPolicy) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ApiAtlasPolicy) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ApiAtlasPolicy) HasBody() bool`

HasBody returns a boolean if a field has been set.
### GetId

`func (o *ApiAtlasPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


