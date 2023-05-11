# CloudProviderAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsIamRoles** | Pointer to [**[]CloudProviderAccessAWSIAMRole**](CloudProviderAccessAWSIAMRole.md) | List that contains the Amazon Web Services (AWS) IAM roles registered and authorized with MongoDB Cloud. | [optional] 

## Methods

### NewCloudProviderAccess

`func NewCloudProviderAccess() *CloudProviderAccess`

NewCloudProviderAccess instantiates a new CloudProviderAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessWithDefaults

`func NewCloudProviderAccessWithDefaults() *CloudProviderAccess`

NewCloudProviderAccessWithDefaults instantiates a new CloudProviderAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsIamRoles

`func (o *CloudProviderAccess) GetAwsIamRoles() []CloudProviderAccessAWSIAMRole`

GetAwsIamRoles returns the AwsIamRoles field if non-nil, zero value otherwise.

### GetAwsIamRolesOk

`func (o *CloudProviderAccess) GetAwsIamRolesOk() (*[]CloudProviderAccessAWSIAMRole, bool)`

GetAwsIamRolesOk returns a tuple with the AwsIamRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIamRoles

`func (o *CloudProviderAccess) SetAwsIamRoles(v []CloudProviderAccessAWSIAMRole)`

SetAwsIamRoles sets AwsIamRoles field to given value.

### HasAwsIamRoles

`func (o *CloudProviderAccess) HasAwsIamRoles() bool`

HasAwsIamRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


