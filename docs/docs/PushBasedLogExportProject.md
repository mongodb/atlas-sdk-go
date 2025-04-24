# PushBasedLogExportProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The name of the bucket to which the agent will send the logs to. | [optional] 
**CreateDate** | Pointer to **time.Time** | Date and time that this feature was enabled on. | [optional] [readonly] 
**IamRoleId** | Pointer to **string** | ID of the AWS IAM role that will be used to write to the S3 bucket. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**PrefixPath** | Pointer to **string** | S3 directory in which vector will write to in order to store the logs. An empty string denotes the root directory. | [optional] 
**State** | Pointer to **string** | Describes whether or not the feature is enabled and what status it is in. | [optional] [readonly] 

## Methods

### NewPushBasedLogExportProject

`func NewPushBasedLogExportProject() *PushBasedLogExportProject`

NewPushBasedLogExportProject instantiates a new PushBasedLogExportProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushBasedLogExportProjectWithDefaults

`func NewPushBasedLogExportProjectWithDefaults() *PushBasedLogExportProject`

NewPushBasedLogExportProjectWithDefaults instantiates a new PushBasedLogExportProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *PushBasedLogExportProject) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *PushBasedLogExportProject) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *PushBasedLogExportProject) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *PushBasedLogExportProject) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.
### GetCreateDate

`func (o *PushBasedLogExportProject) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *PushBasedLogExportProject) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *PushBasedLogExportProject) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *PushBasedLogExportProject) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.
### GetIamRoleId

`func (o *PushBasedLogExportProject) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *PushBasedLogExportProject) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *PushBasedLogExportProject) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *PushBasedLogExportProject) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.
### GetLinks

`func (o *PushBasedLogExportProject) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PushBasedLogExportProject) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PushBasedLogExportProject) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PushBasedLogExportProject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetPrefixPath

`func (o *PushBasedLogExportProject) GetPrefixPath() string`

GetPrefixPath returns the PrefixPath field if non-nil, zero value otherwise.

### GetPrefixPathOk

`func (o *PushBasedLogExportProject) GetPrefixPathOk() (*string, bool)`

GetPrefixPathOk returns a tuple with the PrefixPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixPath

`func (o *PushBasedLogExportProject) SetPrefixPath(v string)`

SetPrefixPath sets PrefixPath field to given value.

### HasPrefixPath

`func (o *PushBasedLogExportProject) HasPrefixPath() bool`

HasPrefixPath returns a boolean if a field has been set.
### GetState

`func (o *PushBasedLogExportProject) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PushBasedLogExportProject) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PushBasedLogExportProject) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PushBasedLogExportProject) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


