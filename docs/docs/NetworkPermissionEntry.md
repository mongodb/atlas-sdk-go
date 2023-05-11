# NetworkPermissionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsSecurityGroup** | Pointer to **string** | Unique string of the Amazon Web Services (AWS) security group that you want to add to the project&#39;s IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. You must configure Virtual Private Connection (VPC) peering for your project before you can add an AWS security group to an IP access list. You cannot set AWS security groups as temporary access list entries. Don&#39;t set this parameter if you set **cidrBlock** or **ipAddress**. | [optional] 
**CidrBlock** | Pointer to **string** | Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that you want to add to the project&#39;s IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. Don&#39;t set this parameter if you set **awsSecurityGroup** or **ipAddress**. | [optional] 
**Comment** | Pointer to **string** | Remark that explains the purpose or scope of this IP access list entry. | [optional] 
**DeleteAfterDate** | Pointer to **time.Time** | Date and time after which MongoDB Cloud deletes the temporary access list entry. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. The date must be later than the current date but no later than one week after you submit this request. The resource returns this parameter if you specified an expiration date when creating this IP access list entry. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that contains the IP access list to which you want to add one or more entries. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IP address that you want to add to the project&#39;s IP access list. Your IP access list entry can be one **awsSecurityGroup**, one **cidrBlock**, or one **ipAddress**. Don&#39;t set this parameter if you set **awsSecurityGroup** or **cidrBlock**. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewNetworkPermissionEntry

`func NewNetworkPermissionEntry() *NetworkPermissionEntry`

NewNetworkPermissionEntry instantiates a new NetworkPermissionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPermissionEntryWithDefaults

`func NewNetworkPermissionEntryWithDefaults() *NetworkPermissionEntry`

NewNetworkPermissionEntryWithDefaults instantiates a new NetworkPermissionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsSecurityGroup

`func (o *NetworkPermissionEntry) GetAwsSecurityGroup() string`

GetAwsSecurityGroup returns the AwsSecurityGroup field if non-nil, zero value otherwise.

### GetAwsSecurityGroupOk

`func (o *NetworkPermissionEntry) GetAwsSecurityGroupOk() (*string, bool)`

GetAwsSecurityGroupOk returns a tuple with the AwsSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecurityGroup

`func (o *NetworkPermissionEntry) SetAwsSecurityGroup(v string)`

SetAwsSecurityGroup sets AwsSecurityGroup field to given value.

### HasAwsSecurityGroup

`func (o *NetworkPermissionEntry) HasAwsSecurityGroup() bool`

HasAwsSecurityGroup returns a boolean if a field has been set.

### GetCidrBlock

`func (o *NetworkPermissionEntry) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *NetworkPermissionEntry) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *NetworkPermissionEntry) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *NetworkPermissionEntry) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetComment

`func (o *NetworkPermissionEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NetworkPermissionEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NetworkPermissionEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NetworkPermissionEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeleteAfterDate

`func (o *NetworkPermissionEntry) GetDeleteAfterDate() time.Time`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *NetworkPermissionEntry) GetDeleteAfterDateOk() (*time.Time, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *NetworkPermissionEntry) SetDeleteAfterDate(v time.Time)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *NetworkPermissionEntry) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *NetworkPermissionEntry) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NetworkPermissionEntry) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NetworkPermissionEntry) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NetworkPermissionEntry) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkPermissionEntry) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkPermissionEntry) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkPermissionEntry) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkPermissionEntry) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLinks

`func (o *NetworkPermissionEntry) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NetworkPermissionEntry) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NetworkPermissionEntry) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NetworkPermissionEntry) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


