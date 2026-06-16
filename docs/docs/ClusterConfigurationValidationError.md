# ClusterConfigurationValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Machine-readable error code identifying the type of validation failure. | [optional] 
**ValidationIssue** | Pointer to **string** | Description of the validation failure. | [optional] 

## Methods

### NewClusterConfigurationValidationError

`func NewClusterConfigurationValidationError() *ClusterConfigurationValidationError`

NewClusterConfigurationValidationError instantiates a new ClusterConfigurationValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationValidationErrorWithDefaults

`func NewClusterConfigurationValidationErrorWithDefaults() *ClusterConfigurationValidationError`

NewClusterConfigurationValidationErrorWithDefaults instantiates a new ClusterConfigurationValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ClusterConfigurationValidationError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ClusterConfigurationValidationError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ClusterConfigurationValidationError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ClusterConfigurationValidationError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.
### GetValidationIssue

`func (o *ClusterConfigurationValidationError) GetValidationIssue() string`

GetValidationIssue returns the ValidationIssue field if non-nil, zero value otherwise.

### GetValidationIssueOk

`func (o *ClusterConfigurationValidationError) GetValidationIssueOk() (*string, bool)`

GetValidationIssueOk returns a tuple with the ValidationIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationIssue

`func (o *ClusterConfigurationValidationError) SetValidationIssue(v string)`

SetValidationIssue sets ValidationIssue field to given value.

### HasValidationIssue

`func (o *ClusterConfigurationValidationError) HasValidationIssue() bool`

HasValidationIssue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


