# ResourcePolicyReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal character string that identifies the Atlas Resource Policy. | [readonly] 
**Name** | **string** | Human-readable label that identifies the Atlas Resource Policy. | [readonly] 

## Methods

### NewResourcePolicyReference

`func NewResourcePolicyReference(id string, name string, ) *ResourcePolicyReference`

NewResourcePolicyReference instantiates a new ResourcePolicyReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePolicyReferenceWithDefaults

`func NewResourcePolicyReferenceWithDefaults() *ResourcePolicyReference`

NewResourcePolicyReferenceWithDefaults instantiates a new ResourcePolicyReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcePolicyReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcePolicyReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcePolicyReference) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *ResourcePolicyReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourcePolicyReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourcePolicyReference) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


