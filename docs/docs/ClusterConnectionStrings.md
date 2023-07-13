# ClusterConnectionStrings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsPrivateLink** | Pointer to **map[string]string** | Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related &#x60;mongodb://&#x60; connection string that you use to connect to MongoDB Cloud through the interface endpoint that the key names. | [optional] [readonly] 
**AwsPrivateLinkSrv** | Pointer to **map[string]string** | Private endpoint-aware connection strings that use AWS-hosted clusters with Amazon Web Services (AWS) PrivateLink. Each key identifies an Amazon Web Services (AWS) interface endpoint. Each value identifies the related &#x60;mongodb://&#x60; connection string that you use to connect to Atlas through the interface endpoint that the key names. | [optional] [readonly] 
**Private** | Pointer to **string** | Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. This connection string uses the &#x60;mongodb+srv://&#x60; protocol. The resource returns this parameter once someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don&#39;t need to append the seed list or change the URI if the nodes change. Use this URI format if your driver supports it. If it doesn&#39;t, use connectionStrings.private. For Amazon Web Services (AWS) clusters, this resource returns this parameter only if you enable custom DNS. | [optional] [readonly] 
**PrivateEndpoint** | Pointer to [**[]ClusterDescriptionConnectionStringsPrivateEndpoint**](ClusterDescriptionConnectionStringsPrivateEndpoint.md) | List of private endpoint-aware connection strings that you can use to connect to this cluster through a private endpoint. This parameter returns only if you deployed a private endpoint to all regions to which you deployed this clusters&#39; nodes. | [optional] [readonly] 
**PrivateSrv** | Pointer to **string** | Network peering connection strings for each interface Virtual Private Cloud (VPC) endpoint that you configured to connect to this cluster. This connection string uses the &#x60;mongodb+srv://&#x60; protocol. The resource returns this parameter when someone creates a network peering connection to this cluster. This protocol tells the application to look up the host seed list in the Domain Name System (DNS). This list synchronizes with the nodes in a cluster. If the connection string uses this Uniform Resource Identifier (URI) format, you don&#39;t need to append the seed list or change the Uniform Resource Identifier (URI) if the nodes change. Use this Uniform Resource Identifier (URI) format if your driver supports it. If it doesn&#39;t, use &#x60;connectionStrings.private&#x60;. For Amazon Web Services (AWS) clusters, this parameter returns only if you [enable custom DNS](https://docs.atlas.mongodb.com/reference/api/aws-custom-dns-update/). | [optional] [readonly] 
**Standard** | Pointer to **string** | Public connection string that you can use to connect to this cluster. This connection string uses the &#x60;mongodb://&#x60; protocol. | [optional] [readonly] 
**StandardSrv** | Pointer to **string** | Public connection string that you can use to connect to this cluster. This connection string uses the &#x60;mongodb+srv://&#x60; protocol. | [optional] [readonly] 

## Methods

### NewClusterConnectionStrings

`func NewClusterConnectionStrings() *ClusterConnectionStrings`

NewClusterConnectionStrings instantiates a new ClusterConnectionStrings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConnectionStringsWithDefaults

`func NewClusterConnectionStringsWithDefaults() *ClusterConnectionStrings`

NewClusterConnectionStringsWithDefaults instantiates a new ClusterConnectionStrings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsPrivateLink

`func (o *ClusterConnectionStrings) GetAwsPrivateLink() map[string]string`

GetAwsPrivateLink returns the AwsPrivateLink field if non-nil, zero value otherwise.

### GetAwsPrivateLinkOk

`func (o *ClusterConnectionStrings) GetAwsPrivateLinkOk() (*map[string]string, bool)`

GetAwsPrivateLinkOk returns a tuple with the AwsPrivateLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsPrivateLink

`func (o *ClusterConnectionStrings) SetAwsPrivateLink(v map[string]string)`

SetAwsPrivateLink sets AwsPrivateLink field to given value.

### HasAwsPrivateLink

`func (o *ClusterConnectionStrings) HasAwsPrivateLink() bool`

HasAwsPrivateLink returns a boolean if a field has been set.
### GetAwsPrivateLinkSrv

`func (o *ClusterConnectionStrings) GetAwsPrivateLinkSrv() map[string]string`

GetAwsPrivateLinkSrv returns the AwsPrivateLinkSrv field if non-nil, zero value otherwise.

### GetAwsPrivateLinkSrvOk

`func (o *ClusterConnectionStrings) GetAwsPrivateLinkSrvOk() (*map[string]string, bool)`

GetAwsPrivateLinkSrvOk returns a tuple with the AwsPrivateLinkSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsPrivateLinkSrv

`func (o *ClusterConnectionStrings) SetAwsPrivateLinkSrv(v map[string]string)`

SetAwsPrivateLinkSrv sets AwsPrivateLinkSrv field to given value.

### HasAwsPrivateLinkSrv

`func (o *ClusterConnectionStrings) HasAwsPrivateLinkSrv() bool`

HasAwsPrivateLinkSrv returns a boolean if a field has been set.
### GetPrivate

`func (o *ClusterConnectionStrings) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ClusterConnectionStrings) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ClusterConnectionStrings) SetPrivate(v string)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ClusterConnectionStrings) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.
### GetPrivateEndpoint

`func (o *ClusterConnectionStrings) GetPrivateEndpoint() []ClusterDescriptionConnectionStringsPrivateEndpoint`

GetPrivateEndpoint returns the PrivateEndpoint field if non-nil, zero value otherwise.

### GetPrivateEndpointOk

`func (o *ClusterConnectionStrings) GetPrivateEndpointOk() (*[]ClusterDescriptionConnectionStringsPrivateEndpoint, bool)`

GetPrivateEndpointOk returns a tuple with the PrivateEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpoint

`func (o *ClusterConnectionStrings) SetPrivateEndpoint(v []ClusterDescriptionConnectionStringsPrivateEndpoint)`

SetPrivateEndpoint sets PrivateEndpoint field to given value.

### HasPrivateEndpoint

`func (o *ClusterConnectionStrings) HasPrivateEndpoint() bool`

HasPrivateEndpoint returns a boolean if a field has been set.
### GetPrivateSrv

`func (o *ClusterConnectionStrings) GetPrivateSrv() string`

GetPrivateSrv returns the PrivateSrv field if non-nil, zero value otherwise.

### GetPrivateSrvOk

`func (o *ClusterConnectionStrings) GetPrivateSrvOk() (*string, bool)`

GetPrivateSrvOk returns a tuple with the PrivateSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSrv

`func (o *ClusterConnectionStrings) SetPrivateSrv(v string)`

SetPrivateSrv sets PrivateSrv field to given value.

### HasPrivateSrv

`func (o *ClusterConnectionStrings) HasPrivateSrv() bool`

HasPrivateSrv returns a boolean if a field has been set.
### GetStandard

`func (o *ClusterConnectionStrings) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *ClusterConnectionStrings) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *ClusterConnectionStrings) SetStandard(v string)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *ClusterConnectionStrings) HasStandard() bool`

HasStandard returns a boolean if a field has been set.
### GetStandardSrv

`func (o *ClusterConnectionStrings) GetStandardSrv() string`

GetStandardSrv returns the StandardSrv field if non-nil, zero value otherwise.

### GetStandardSrvOk

`func (o *ClusterConnectionStrings) GetStandardSrvOk() (*string, bool)`

GetStandardSrvOk returns a tuple with the StandardSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardSrv

`func (o *ClusterConnectionStrings) SetStandardSrv(v string)`

SetStandardSrv sets StandardSrv field to given value.

### HasStandardSrv

`func (o *ClusterConnectionStrings) HasStandardSrv() bool`

HasStandardSrv returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


