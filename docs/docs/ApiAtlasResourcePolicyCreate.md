# ApiAtlasResourcePolicyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable label that describes the atlas resource policy. | 
**Policies** | [**[]ApiAtlasPolicyCreate**](ApiAtlasPolicyCreate.md) | List of policies that make up the atlas resource policy. | 

## Methods

### NewApiAtlasResourcePolicyCreate

`func NewApiAtlasResourcePolicyCreate(name string, policies []ApiAtlasPolicyCreate, ) *ApiAtlasResourcePolicyCreate`

NewApiAtlasResourcePolicyCreate instantiates a new ApiAtlasResourcePolicyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasResourcePolicyCreateWithDefaults

`func NewApiAtlasResourcePolicyCreateWithDefaults() *ApiAtlasResourcePolicyCreate`

NewApiAtlasResourcePolicyCreateWithDefaults instantiates a new ApiAtlasResourcePolicyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiAtlasResourcePolicyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiAtlasResourcePolicyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiAtlasResourcePolicyCreate) SetName(v string)`

SetName sets Name field to given value.

### GetPolicies

`func (o *ApiAtlasResourcePolicyCreate) GetPolicies() []ApiAtlasPolicyCreate`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ApiAtlasResourcePolicyCreate) GetPoliciesOk() (*[]ApiAtlasPolicyCreate, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ApiAtlasResourcePolicyCreate) SetPolicies(v []ApiAtlasPolicyCreate)`

SetPolicies sets Policies field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


