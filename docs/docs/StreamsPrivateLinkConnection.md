# StreamsPrivateLinkConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Private Link connection. | [optional] [readonly] 
**Arn** | Pointer to **string** | Amazon Resource Name (ARN). Required for AWS Provider and MSK vendor. | [optional] 
**AzureResourceIds** | Pointer to **[]string** | Azure Resource IDs of each availability zone for the Azure Confluent cluster. | [optional] 
**DnsDomain** | Pointer to **string** | The domain hostname. Required for the following provider and vendor combinations: - AWS provider with CONFLUENT vendor. - AZURE provider with EVENTHUB or CONFLUENT vendor. | [optional] 
**DnsSubDomain** | Pointer to **[]string** | Sub-Domain name of Confluent cluster. These are typically your availability zones. Required for AWS Provider and CONFLUENT vendor, if your AWS CONFLUENT cluster doesn&#39;t use subdomains, you must set this to the empty array []. | [optional] 
**ErrorMessage** | Pointer to **string** | Error message if the state is FAILED. | [optional] [readonly] 
**GcpServiceAttachmentUris** | Pointer to **[]string** | Service Attachment URIs of each availability zone for the GCP Confluent cluster. | [optional] 
**InterfaceEndpointId** | Pointer to **string** | Interface endpoint ID that is created from the service endpoint ID provided. | [optional] [readonly] 
**InterfaceEndpointName** | Pointer to **string** | Interface endpoint name that is created from the service endpoint ID provided. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Provider** | **string** | Provider where the Kafka cluster is deployed. Valid values are AWS and AZURE. | 
**ProviderAccountId** | Pointer to **string** | Account ID from the cloud provider. | [optional] [readonly] 
**Region** | Pointer to **string** | The region of the Provider’s cluster. See [AZURE](https://www.mongodb.com/docs/atlas/reference/microsoft-azure/#stream-processing-instances) and [AWS](https://www.mongodb.com/docs/atlas/reference/amazon-aws/#stream-processing-instances) supported regions. | [optional] 
**ServiceEndpointId** | Pointer to **string** | For AZURE EVENTHUB, this is the [namespace endpoint ID](https://learn.microsoft.com/en-us/rest/api/eventhub/namespaces/get). For AWS CONFLUENT cluster, this is the [VPC Endpoint service name](https://docs.confluent.io/cloud/current/networking/private-links/aws-privatelink.html). | [optional] 
**State** | Pointer to **string** | State the connection is in. | [optional] [readonly] 
**Vendor** | Pointer to **string** | Vendor that manages the cloud service. The following are the vendor values per provider: - AWS -- MSK for AWS MSK Kafka clusters -- CONFLUENT for Confluent Kafka clusters on AWS -- KINESIS for AWS Kinesis Data Streams (coming soon).  - Azure -- EVENTHUB for Azure EventHub. -- CONFLUENT for the Confluent Kafka clusters on Azure  **NOTE** Omitting the vendor field will default to using the GENERIC vendor. | [optional] 

## Methods

### NewStreamsPrivateLinkConnection

`func NewStreamsPrivateLinkConnection(provider string, ) *StreamsPrivateLinkConnection`

NewStreamsPrivateLinkConnection instantiates a new StreamsPrivateLinkConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsPrivateLinkConnectionWithDefaults

`func NewStreamsPrivateLinkConnectionWithDefaults() *StreamsPrivateLinkConnection`

NewStreamsPrivateLinkConnectionWithDefaults instantiates a new StreamsPrivateLinkConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamsPrivateLinkConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsPrivateLinkConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsPrivateLinkConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamsPrivateLinkConnection) HasId() bool`

HasId returns a boolean if a field has been set.
### GetArn

`func (o *StreamsPrivateLinkConnection) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *StreamsPrivateLinkConnection) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *StreamsPrivateLinkConnection) SetArn(v string)`

SetArn sets Arn field to given value.

### HasArn

`func (o *StreamsPrivateLinkConnection) HasArn() bool`

HasArn returns a boolean if a field has been set.
### GetAzureResourceIds

`func (o *StreamsPrivateLinkConnection) GetAzureResourceIds() []string`

GetAzureResourceIds returns the AzureResourceIds field if non-nil, zero value otherwise.

### GetAzureResourceIdsOk

`func (o *StreamsPrivateLinkConnection) GetAzureResourceIdsOk() (*[]string, bool)`

GetAzureResourceIdsOk returns a tuple with the AzureResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceIds

`func (o *StreamsPrivateLinkConnection) SetAzureResourceIds(v []string)`

SetAzureResourceIds sets AzureResourceIds field to given value.

### HasAzureResourceIds

`func (o *StreamsPrivateLinkConnection) HasAzureResourceIds() bool`

HasAzureResourceIds returns a boolean if a field has been set.
### GetDnsDomain

`func (o *StreamsPrivateLinkConnection) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *StreamsPrivateLinkConnection) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *StreamsPrivateLinkConnection) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *StreamsPrivateLinkConnection) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.
### GetDnsSubDomain

`func (o *StreamsPrivateLinkConnection) GetDnsSubDomain() []string`

GetDnsSubDomain returns the DnsSubDomain field if non-nil, zero value otherwise.

### GetDnsSubDomainOk

`func (o *StreamsPrivateLinkConnection) GetDnsSubDomainOk() (*[]string, bool)`

GetDnsSubDomainOk returns a tuple with the DnsSubDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSubDomain

`func (o *StreamsPrivateLinkConnection) SetDnsSubDomain(v []string)`

SetDnsSubDomain sets DnsSubDomain field to given value.

### HasDnsSubDomain

`func (o *StreamsPrivateLinkConnection) HasDnsSubDomain() bool`

HasDnsSubDomain returns a boolean if a field has been set.
### GetErrorMessage

`func (o *StreamsPrivateLinkConnection) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *StreamsPrivateLinkConnection) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *StreamsPrivateLinkConnection) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *StreamsPrivateLinkConnection) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetGcpServiceAttachmentUris

`func (o *StreamsPrivateLinkConnection) GetGcpServiceAttachmentUris() []string`

GetGcpServiceAttachmentUris returns the GcpServiceAttachmentUris field if non-nil, zero value otherwise.

### GetGcpServiceAttachmentUrisOk

`func (o *StreamsPrivateLinkConnection) GetGcpServiceAttachmentUrisOk() (*[]string, bool)`

GetGcpServiceAttachmentUrisOk returns a tuple with the GcpServiceAttachmentUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAttachmentUris

`func (o *StreamsPrivateLinkConnection) SetGcpServiceAttachmentUris(v []string)`

SetGcpServiceAttachmentUris sets GcpServiceAttachmentUris field to given value.

### HasGcpServiceAttachmentUris

`func (o *StreamsPrivateLinkConnection) HasGcpServiceAttachmentUris() bool`

HasGcpServiceAttachmentUris returns a boolean if a field has been set.
### GetInterfaceEndpointId

`func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointId() string`

GetInterfaceEndpointId returns the InterfaceEndpointId field if non-nil, zero value otherwise.

### GetInterfaceEndpointIdOk

`func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointIdOk() (*string, bool)`

GetInterfaceEndpointIdOk returns a tuple with the InterfaceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpointId

`func (o *StreamsPrivateLinkConnection) SetInterfaceEndpointId(v string)`

SetInterfaceEndpointId sets InterfaceEndpointId field to given value.

### HasInterfaceEndpointId

`func (o *StreamsPrivateLinkConnection) HasInterfaceEndpointId() bool`

HasInterfaceEndpointId returns a boolean if a field has been set.
### GetInterfaceEndpointName

`func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointName() string`

GetInterfaceEndpointName returns the InterfaceEndpointName field if non-nil, zero value otherwise.

### GetInterfaceEndpointNameOk

`func (o *StreamsPrivateLinkConnection) GetInterfaceEndpointNameOk() (*string, bool)`

GetInterfaceEndpointNameOk returns a tuple with the InterfaceEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpointName

`func (o *StreamsPrivateLinkConnection) SetInterfaceEndpointName(v string)`

SetInterfaceEndpointName sets InterfaceEndpointName field to given value.

### HasInterfaceEndpointName

`func (o *StreamsPrivateLinkConnection) HasInterfaceEndpointName() bool`

HasInterfaceEndpointName returns a boolean if a field has been set.
### GetLinks

`func (o *StreamsPrivateLinkConnection) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsPrivateLinkConnection) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsPrivateLinkConnection) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsPrivateLinkConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetProvider

`func (o *StreamsPrivateLinkConnection) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *StreamsPrivateLinkConnection) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *StreamsPrivateLinkConnection) SetProvider(v string)`

SetProvider sets Provider field to given value.

### GetProviderAccountId

`func (o *StreamsPrivateLinkConnection) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *StreamsPrivateLinkConnection) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *StreamsPrivateLinkConnection) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.

### HasProviderAccountId

`func (o *StreamsPrivateLinkConnection) HasProviderAccountId() bool`

HasProviderAccountId returns a boolean if a field has been set.
### GetRegion

`func (o *StreamsPrivateLinkConnection) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StreamsPrivateLinkConnection) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StreamsPrivateLinkConnection) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *StreamsPrivateLinkConnection) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetServiceEndpointId

`func (o *StreamsPrivateLinkConnection) GetServiceEndpointId() string`

GetServiceEndpointId returns the ServiceEndpointId field if non-nil, zero value otherwise.

### GetServiceEndpointIdOk

`func (o *StreamsPrivateLinkConnection) GetServiceEndpointIdOk() (*string, bool)`

GetServiceEndpointIdOk returns a tuple with the ServiceEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpointId

`func (o *StreamsPrivateLinkConnection) SetServiceEndpointId(v string)`

SetServiceEndpointId sets ServiceEndpointId field to given value.

### HasServiceEndpointId

`func (o *StreamsPrivateLinkConnection) HasServiceEndpointId() bool`

HasServiceEndpointId returns a boolean if a field has been set.
### GetState

`func (o *StreamsPrivateLinkConnection) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamsPrivateLinkConnection) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamsPrivateLinkConnection) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StreamsPrivateLinkConnection) HasState() bool`

HasState returns a boolean if a field has been set.
### GetVendor

`func (o *StreamsPrivateLinkConnection) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *StreamsPrivateLinkConnection) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *StreamsPrivateLinkConnection) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *StreamsPrivateLinkConnection) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


