# CreatePushBasedLogExportProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | The name of the bucket to which the agent will send the logs to. | 
**IamRoleId** | **string** | ID of the AWS IAM role that will be used to write to the S3 bucket. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**PrefixPath** | **string** | S3 directory in which vector will write to in order to store the logs. An empty string denotes the root directory. | 

## Methods

### NewCreatePushBasedLogExportProjectRequest

`func NewCreatePushBasedLogExportProjectRequest(bucketName string, iamRoleId string, prefixPath string, ) *CreatePushBasedLogExportProjectRequest`

NewCreatePushBasedLogExportProjectRequest instantiates a new CreatePushBasedLogExportProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePushBasedLogExportProjectRequestWithDefaults

`func NewCreatePushBasedLogExportProjectRequestWithDefaults() *CreatePushBasedLogExportProjectRequest`

NewCreatePushBasedLogExportProjectRequestWithDefaults instantiates a new CreatePushBasedLogExportProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *CreatePushBasedLogExportProjectRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CreatePushBasedLogExportProjectRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CreatePushBasedLogExportProjectRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### GetIamRoleId

`func (o *CreatePushBasedLogExportProjectRequest) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *CreatePushBasedLogExportProjectRequest) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *CreatePushBasedLogExportProjectRequest) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### GetLinks

`func (o *CreatePushBasedLogExportProjectRequest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreatePushBasedLogExportProjectRequest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreatePushBasedLogExportProjectRequest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreatePushBasedLogExportProjectRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetPrefixPath

`func (o *CreatePushBasedLogExportProjectRequest) GetPrefixPath() string`

GetPrefixPath returns the PrefixPath field if non-nil, zero value otherwise.

### GetPrefixPathOk

`func (o *CreatePushBasedLogExportProjectRequest) GetPrefixPathOk() (*string, bool)`

GetPrefixPathOk returns a tuple with the PrefixPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixPath

`func (o *CreatePushBasedLogExportProjectRequest) SetPrefixPath(v string)`

SetPrefixPath sets PrefixPath field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


