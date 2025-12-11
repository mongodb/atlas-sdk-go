# QueryShapeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | The MongoDB command type issued for a query shape. | [optional] [readonly] 
**Namespace** | Pointer to **string** | Human-readable label that identifies the namespace on the specified host. The resource expresses this parameter value as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | [optional] [readonly] 
**QueryShape** | Pointer to **string** | A query shape is a set of specifications that group similar queries together. Specifications can include filters, sorts, projections, aggregation pipeline stages, a namespace, and others. Queries that have similar specifications have the same query shape. This field may be null if the user lacks PII view access. | [optional] [readonly] 
**QueryShapeHash** | **string** | A hexadecimal string that represents the hash of a MongoDB query shape. | [readonly] 
**Status** | **string** | The rejection status of a query shape. Use REJECTED to prevent the query shape from executing on the cluster, or UNREJECTED to allow it to execute. | 

## Methods

### NewQueryShapeResponse

`func NewQueryShapeResponse(queryShapeHash string, status string, ) *QueryShapeResponse`

NewQueryShapeResponse instantiates a new QueryShapeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryShapeResponseWithDefaults

`func NewQueryShapeResponseWithDefaults() *QueryShapeResponse`

NewQueryShapeResponseWithDefaults instantiates a new QueryShapeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *QueryShapeResponse) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *QueryShapeResponse) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *QueryShapeResponse) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *QueryShapeResponse) HasCommand() bool`

HasCommand returns a boolean if a field has been set.
### GetNamespace

`func (o *QueryShapeResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *QueryShapeResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *QueryShapeResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *QueryShapeResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.
### GetQueryShape

`func (o *QueryShapeResponse) GetQueryShape() string`

GetQueryShape returns the QueryShape field if non-nil, zero value otherwise.

### GetQueryShapeOk

`func (o *QueryShapeResponse) GetQueryShapeOk() (*string, bool)`

GetQueryShapeOk returns a tuple with the QueryShape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryShape

`func (o *QueryShapeResponse) SetQueryShape(v string)`

SetQueryShape sets QueryShape field to given value.

### HasQueryShape

`func (o *QueryShapeResponse) HasQueryShape() bool`

HasQueryShape returns a boolean if a field has been set.
### GetQueryShapeHash

`func (o *QueryShapeResponse) GetQueryShapeHash() string`

GetQueryShapeHash returns the QueryShapeHash field if non-nil, zero value otherwise.

### GetQueryShapeHashOk

`func (o *QueryShapeResponse) GetQueryShapeHashOk() (*string, bool)`

GetQueryShapeHashOk returns a tuple with the QueryShapeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryShapeHash

`func (o *QueryShapeResponse) SetQueryShapeHash(v string)`

SetQueryShapeHash sets QueryShapeHash field to given value.

### GetStatus

`func (o *QueryShapeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryShapeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryShapeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


