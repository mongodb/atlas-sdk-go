# DataLakeHTTPStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInsecure** | Pointer to **bool** | Flag that validates the scheme in the specified URLs. If &#x60;true&#x60;, allows insecure &#x60;HTTP&#x60; scheme, doesn&#39;t verify the server&#39;s certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If &#x60;false&#x60;, allows secure &#x60;HTTPS&#x60; scheme only. | [optional] [default to false]
**DefaultFormat** | Pointer to **string** | Default format that Data Lake assumes if it encounters a file without an extension while searching the &#x60;storeName&#x60;. If omitted, Data Lake attempts to detect the file type by processing a few bytes of the file. The specified format only applies to the URLs specified in the **databases.[n].collections.[n].dataSources** object. | [optional] 
**Urls** | Pointer to **[]string** | Comma-separated list of publicly accessible HTTP URLs where data is stored. You can&#39;t specify URLs that require authentication. | [optional] 
**Name** | Pointer to **string** | Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an &#x60;M10&#x60; or higher cluster. | [optional] 
**Provider** | **string** |  | 

## Methods

### NewDataLakeHTTPStore

`func NewDataLakeHTTPStore(provider string, ) *DataLakeHTTPStore`

NewDataLakeHTTPStore instantiates a new DataLakeHTTPStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeHTTPStoreWithDefaults

`func NewDataLakeHTTPStoreWithDefaults() *DataLakeHTTPStore`

NewDataLakeHTTPStoreWithDefaults instantiates a new DataLakeHTTPStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *DataLakeHTTPStore) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *DataLakeHTTPStore) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *DataLakeHTTPStore) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *DataLakeHTTPStore) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetDefaultFormat

`func (o *DataLakeHTTPStore) GetDefaultFormat() string`

GetDefaultFormat returns the DefaultFormat field if non-nil, zero value otherwise.

### GetDefaultFormatOk

`func (o *DataLakeHTTPStore) GetDefaultFormatOk() (*string, bool)`

GetDefaultFormatOk returns a tuple with the DefaultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormat

`func (o *DataLakeHTTPStore) SetDefaultFormat(v string)`

SetDefaultFormat sets DefaultFormat field to given value.

### HasDefaultFormat

`func (o *DataLakeHTTPStore) HasDefaultFormat() bool`

HasDefaultFormat returns a boolean if a field has been set.

### GetUrls

`func (o *DataLakeHTTPStore) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *DataLakeHTTPStore) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *DataLakeHTTPStore) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *DataLakeHTTPStore) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetName

`func (o *DataLakeHTTPStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeHTTPStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeHTTPStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeHTTPStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DataLakeHTTPStore) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DataLakeHTTPStore) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DataLakeHTTPStore) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


