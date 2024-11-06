# CreateDataProcessRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the Cloud service provider where you wish to store your archived data. **AZURE** or **GCP** may be selected only if it is the Cloud service provider for the cluster and no archives for any other cloud provider have been created for the cluster. **GCP** is currently only supported in [private preview](https://www.mongodb.com/community/forums/t/invitation-to-participate-in-the-private-preview-program-of-online-archive-on-gcp). | [optional] 
**Region** | Pointer to **string** | Human-readable label that identifies the geographic location of the region where you wish to store your archived data. | [optional] 

## Methods

### NewCreateDataProcessRegion

`func NewCreateDataProcessRegion() *CreateDataProcessRegion`

NewCreateDataProcessRegion instantiates a new CreateDataProcessRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataProcessRegionWithDefaults

`func NewCreateDataProcessRegionWithDefaults() *CreateDataProcessRegion`

NewCreateDataProcessRegionWithDefaults instantiates a new CreateDataProcessRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateDataProcessRegion) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateDataProcessRegion) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateDataProcessRegion) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CreateDataProcessRegion) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetRegion

`func (o *CreateDataProcessRegion) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateDataProcessRegion) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateDataProcessRegion) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateDataProcessRegion) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


