# QuerySamplingCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | Human-readable label that identifies the namespace (&#x60;&lt;database&gt;.&lt;collection&gt;&#x60;) on which to sample queries. | 

## Methods

### NewQuerySamplingCreateRequest

`func NewQuerySamplingCreateRequest(namespace string, ) *QuerySamplingCreateRequest`

NewQuerySamplingCreateRequest instantiates a new QuerySamplingCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySamplingCreateRequestWithDefaults

`func NewQuerySamplingCreateRequestWithDefaults() *QuerySamplingCreateRequest`

NewQuerySamplingCreateRequestWithDefaults instantiates a new QuerySamplingCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *QuerySamplingCreateRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *QuerySamplingCreateRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *QuerySamplingCreateRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


