# CustomCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | MongoDB find query that selects documents to archive. The specified query follows the syntax of the &#x60;db.collection.find(query)&#x60; command. This query can&#39;t use the empty document (&#x60;{}&#x60;) to return all documents. Set this parameter when **\&quot;criteria.type\&quot; : \&quot;CUSTOM\&quot;**. | 
**Type** | Pointer to **string** | Means by which MongoDB Cloud selects data to archive. Data can be chosen using the age of the data or a MongoDB query. **DATE** selects documents to archive based on a date. **CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn&#39;t support **CUSTOM** when &#x60;\&quot;collectionType\&quot;: \&quot;TIMESERIES\&quot;&#x60;. | [optional] 

## Methods

### NewCustomCriteria

`func NewCustomCriteria(query string, ) *CustomCriteria`

NewCustomCriteria instantiates a new CustomCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCriteriaWithDefaults

`func NewCustomCriteriaWithDefaults() *CustomCriteria`

NewCustomCriteriaWithDefaults instantiates a new CustomCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CustomCriteria) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CustomCriteria) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CustomCriteria) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *CustomCriteria) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomCriteria) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomCriteria) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomCriteria) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


