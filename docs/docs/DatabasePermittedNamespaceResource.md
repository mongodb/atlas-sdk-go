# DatabasePermittedNamespaceResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **bool** | Flag that indicates whether to grant the action on the cluster resource. If &#x60;true&#x60;, MongoDB Cloud ignores the **actions.resources.collection** and **actions.resources.db** parameters. | 
**Collection** | **string** | Human-readable label that identifies the collection on which you grant the action to one MongoDB user. If you don&#39;t set this parameter, you grant the action to all collections in the database specified in the **actions.resources.db** parameter. If you set &#x60;\&quot;actions.resources.cluster\&quot; : true&#x60;, MongoDB Cloud ignores this parameter. | 
**Db** | **string** | Human-readable label that identifies the database on which you grant the action to one MongoDB user. If you set &#x60;\&quot;actions.resources.cluster\&quot; : true&#x60;, MongoDB Cloud ignores this parameter. | 

## Methods

### NewDatabasePermittedNamespaceResource

`func NewDatabasePermittedNamespaceResource(cluster bool, collection string, db string, ) *DatabasePermittedNamespaceResource`

NewDatabasePermittedNamespaceResource instantiates a new DatabasePermittedNamespaceResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabasePermittedNamespaceResourceWithDefaults

`func NewDatabasePermittedNamespaceResourceWithDefaults() *DatabasePermittedNamespaceResource`

NewDatabasePermittedNamespaceResourceWithDefaults instantiates a new DatabasePermittedNamespaceResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DatabasePermittedNamespaceResource) GetCluster() bool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DatabasePermittedNamespaceResource) GetClusterOk() (*bool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DatabasePermittedNamespaceResource) SetCluster(v bool)`

SetCluster sets Cluster field to given value.

### GetCollection

`func (o *DatabasePermittedNamespaceResource) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *DatabasePermittedNamespaceResource) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *DatabasePermittedNamespaceResource) SetCollection(v string)`

SetCollection sets Collection field to given value.

### GetDb

`func (o *DatabasePermittedNamespaceResource) GetDb() string`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *DatabasePermittedNamespaceResource) GetDbOk() (*string, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *DatabasePermittedNamespaceResource) SetDb(v string)`

SetDb sets Db field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


