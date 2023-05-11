# DataLakeDatabaseDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInsecure** | Pointer to **bool** | Flag that validates the scheme in the specified URLs. If &#x60;true&#x60;, allows insecure &#x60;HTTP&#x60; scheme, doesn&#39;t verify the server&#39;s certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If &#x60;false&#x60;, allows secure &#x60;HTTPS&#x60; scheme only. | [optional] [default to false]
**Collection** | Pointer to **string** | Human-readable label that identifies the collection in the database. For creating a wildcard (&#x60;*&#x60;) collection, you must omit this parameter. | [optional] 
**CollectionRegex** | Pointer to **string** | Regex pattern to use for creating the wildcard (*) collection. To learn more about the regex syntax, see [Go programming language](https://pkg.go.dev/regexp). | [optional] 
**Database** | Pointer to **string** | Human-readable label that identifies the database, which contains the collection in the cluster. You must omit this parameter to generate wildcard (&#x60;*&#x60;) collections for dynamically generated databases. | [optional] 
**DatabaseRegex** | Pointer to **string** | Regex pattern to use for creating the wildcard (*) database. To learn more about the regex syntax, see [Go programming language](https://pkg.go.dev/regexp). | [optional] 
**DefaultFormat** | Pointer to **string** | File format that MongoDB Cloud uses if it encounters a file without a file extension while searching **storeName**. | [optional] 
**Path** | Pointer to **string** | File path that controls how MongoDB Cloud searches for and parses files in the **storeName** before mapping them to a collection.Specify &#x60;&#x60;/&#x60;&#x60; to capture all files and folders from the &#x60;&#x60;prefix&#x60;&#x60; path. | [optional] 
**ProvenanceFieldName** | Pointer to **string** | Name for the field that includes the provenance of the documents in the results. MongoDB Cloud returns different fields in the results for each supported provider. | [optional] 
**StoreName** | Pointer to **string** | Human-readable label that identifies the data store that MongoDB Cloud maps to the collection. | [optional] 
**Urls** | Pointer to **[]string** | URLs of the publicly accessible data files. You can&#39;t specify URLs that require authentication. Atlas Data Lake creates a partition for each URL. If empty or omitted, Data Lake uses the URLs from the store specified in the **dataSources.storeName** parameter. | [optional] 

## Methods

### NewDataLakeDatabaseDataSource

`func NewDataLakeDatabaseDataSource() *DataLakeDatabaseDataSource`

NewDataLakeDatabaseDataSource instantiates a new DataLakeDatabaseDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeDatabaseDataSourceWithDefaults

`func NewDataLakeDatabaseDataSourceWithDefaults() *DataLakeDatabaseDataSource`

NewDataLakeDatabaseDataSourceWithDefaults instantiates a new DataLakeDatabaseDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInsecure

`func (o *DataLakeDatabaseDataSource) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *DataLakeDatabaseDataSource) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *DataLakeDatabaseDataSource) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *DataLakeDatabaseDataSource) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetCollection

`func (o *DataLakeDatabaseDataSource) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *DataLakeDatabaseDataSource) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *DataLakeDatabaseDataSource) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *DataLakeDatabaseDataSource) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetCollectionRegex

`func (o *DataLakeDatabaseDataSource) GetCollectionRegex() string`

GetCollectionRegex returns the CollectionRegex field if non-nil, zero value otherwise.

### GetCollectionRegexOk

`func (o *DataLakeDatabaseDataSource) GetCollectionRegexOk() (*string, bool)`

GetCollectionRegexOk returns a tuple with the CollectionRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionRegex

`func (o *DataLakeDatabaseDataSource) SetCollectionRegex(v string)`

SetCollectionRegex sets CollectionRegex field to given value.

### HasCollectionRegex

`func (o *DataLakeDatabaseDataSource) HasCollectionRegex() bool`

HasCollectionRegex returns a boolean if a field has been set.

### GetDatabase

`func (o *DataLakeDatabaseDataSource) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DataLakeDatabaseDataSource) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DataLakeDatabaseDataSource) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *DataLakeDatabaseDataSource) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetDatabaseRegex

`func (o *DataLakeDatabaseDataSource) GetDatabaseRegex() string`

GetDatabaseRegex returns the DatabaseRegex field if non-nil, zero value otherwise.

### GetDatabaseRegexOk

`func (o *DataLakeDatabaseDataSource) GetDatabaseRegexOk() (*string, bool)`

GetDatabaseRegexOk returns a tuple with the DatabaseRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseRegex

`func (o *DataLakeDatabaseDataSource) SetDatabaseRegex(v string)`

SetDatabaseRegex sets DatabaseRegex field to given value.

### HasDatabaseRegex

`func (o *DataLakeDatabaseDataSource) HasDatabaseRegex() bool`

HasDatabaseRegex returns a boolean if a field has been set.

### GetDefaultFormat

`func (o *DataLakeDatabaseDataSource) GetDefaultFormat() string`

GetDefaultFormat returns the DefaultFormat field if non-nil, zero value otherwise.

### GetDefaultFormatOk

`func (o *DataLakeDatabaseDataSource) GetDefaultFormatOk() (*string, bool)`

GetDefaultFormatOk returns a tuple with the DefaultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormat

`func (o *DataLakeDatabaseDataSource) SetDefaultFormat(v string)`

SetDefaultFormat sets DefaultFormat field to given value.

### HasDefaultFormat

`func (o *DataLakeDatabaseDataSource) HasDefaultFormat() bool`

HasDefaultFormat returns a boolean if a field has been set.

### GetPath

`func (o *DataLakeDatabaseDataSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DataLakeDatabaseDataSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DataLakeDatabaseDataSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DataLakeDatabaseDataSource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProvenanceFieldName

`func (o *DataLakeDatabaseDataSource) GetProvenanceFieldName() string`

GetProvenanceFieldName returns the ProvenanceFieldName field if non-nil, zero value otherwise.

### GetProvenanceFieldNameOk

`func (o *DataLakeDatabaseDataSource) GetProvenanceFieldNameOk() (*string, bool)`

GetProvenanceFieldNameOk returns a tuple with the ProvenanceFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenanceFieldName

`func (o *DataLakeDatabaseDataSource) SetProvenanceFieldName(v string)`

SetProvenanceFieldName sets ProvenanceFieldName field to given value.

### HasProvenanceFieldName

`func (o *DataLakeDatabaseDataSource) HasProvenanceFieldName() bool`

HasProvenanceFieldName returns a boolean if a field has been set.

### GetStoreName

`func (o *DataLakeDatabaseDataSource) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *DataLakeDatabaseDataSource) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *DataLakeDatabaseDataSource) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *DataLakeDatabaseDataSource) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetUrls

`func (o *DataLakeDatabaseDataSource) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *DataLakeDatabaseDataSource) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *DataLakeDatabaseDataSource) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *DataLakeDatabaseDataSource) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


