# NewRelic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique 40-hexadecimal digit string that identifies your New Relic account. | 
**LicenseKey** | **string** | Unique 40-hexadecimal digit string that identifies your New Relic license.  **IMPORTANT**: Effective Wednesday, June 16th, 2021, New Relic no longer supports the plugin-based integration with MongoDB. We do not recommend that you sign up for the plugin-based integration. To learn more, see the &lt;a href&#x3D;\&quot;https://discuss.newrelic.com/t/new-relic-plugin-eol-wednesday-june-16th-2021/127267\&quot; target&#x3D;\&quot;_blank\&quot;&gt;New Relic Plugin EOL Statement&lt;/a&gt; Consider configuring an alternative monitoring integration before June 16th to maintain visibility into your MongoDB deployments. | 
**ReadToken** | **string** | Query key used to access your New Relic account. | 
**Type** | Pointer to **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type. | [optional] 
**WriteToken** | **string** | Insert key associated with your New Relic account. | 

## Methods

### NewNewRelic

`func NewNewRelic(accountId string, licenseKey string, readToken string, writeToken string, ) *NewRelic`

NewNewRelic instantiates a new NewRelic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRelicWithDefaults

`func NewNewRelicWithDefaults() *NewRelic`

NewNewRelicWithDefaults instantiates a new NewRelic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NewRelic) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NewRelic) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NewRelic) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLicenseKey

`func (o *NewRelic) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *NewRelic) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *NewRelic) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.


### GetReadToken

`func (o *NewRelic) GetReadToken() string`

GetReadToken returns the ReadToken field if non-nil, zero value otherwise.

### GetReadTokenOk

`func (o *NewRelic) GetReadTokenOk() (*string, bool)`

GetReadTokenOk returns a tuple with the ReadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadToken

`func (o *NewRelic) SetReadToken(v string)`

SetReadToken sets ReadToken field to given value.


### GetType

`func (o *NewRelic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewRelic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewRelic) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewRelic) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWriteToken

`func (o *NewRelic) GetWriteToken() string`

GetWriteToken returns the WriteToken field if non-nil, zero value otherwise.

### GetWriteTokenOk

`func (o *NewRelic) GetWriteTokenOk() (*string, bool)`

GetWriteTokenOk returns a tuple with the WriteToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteToken

`func (o *NewRelic) SetWriteToken(v string)`

SetWriteToken sets WriteToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


