# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return. | [optional] [readonly] 
**Complete** | Pointer to **bool** | Flag that indicates whether the snapshot exists. This flag returns &#x60;false&#x60; while MongoDB Cloud creates the snapshot. | [optional] [readonly] 
**Created** | Pointer to [**BSONTimestamp**](BSONTimestamp.md) |  | [optional] 
**DoNotDelete** | Pointer to **bool** | Flag that indicates whether someone can delete this snapshot. You can&#39;t set &#x60;\&quot;doNotDelete\&quot; : true&#x60; and set a timestamp for **expires** in the same request. | [optional] 
**Expires** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the snapshot. If &#x60;\&quot;doNotDelete\&quot; : true&#x60;, MongoDB Cloud removes any value set for this parameter. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns the snapshots. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [optional] [readonly] 
**LastOplogAppliedTimestamp** | Pointer to [**BSONTimestamp**](BSONTimestamp.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Parts** | Pointer to [**[]SnapshotPart**](SnapshotPart.md) | Metadata that describes the complete snapshot.  - For a replica set, this array contains a single document. - For a sharded cluster, this array contains one document for each shard plus one document for the config host. | [optional] [readonly] 

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *Snapshot) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Snapshot) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Snapshot) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Snapshot) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetComplete

`func (o *Snapshot) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *Snapshot) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *Snapshot) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *Snapshot) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetCreated

`func (o *Snapshot) GetCreated() BSONTimestamp`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Snapshot) GetCreatedOk() (*BSONTimestamp, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Snapshot) SetCreated(v BSONTimestamp)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Snapshot) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDoNotDelete

`func (o *Snapshot) GetDoNotDelete() bool`

GetDoNotDelete returns the DoNotDelete field if non-nil, zero value otherwise.

### GetDoNotDeleteOk

`func (o *Snapshot) GetDoNotDeleteOk() (*bool, bool)`

GetDoNotDeleteOk returns a tuple with the DoNotDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDelete

`func (o *Snapshot) SetDoNotDelete(v bool)`

SetDoNotDelete sets DoNotDelete field to given value.

### HasDoNotDelete

`func (o *Snapshot) HasDoNotDelete() bool`

HasDoNotDelete returns a boolean if a field has been set.

### GetExpires

`func (o *Snapshot) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Snapshot) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Snapshot) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Snapshot) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetGroupId

`func (o *Snapshot) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Snapshot) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Snapshot) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Snapshot) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *Snapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Snapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Snapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Snapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastOplogAppliedTimestamp

`func (o *Snapshot) GetLastOplogAppliedTimestamp() BSONTimestamp`

GetLastOplogAppliedTimestamp returns the LastOplogAppliedTimestamp field if non-nil, zero value otherwise.

### GetLastOplogAppliedTimestampOk

`func (o *Snapshot) GetLastOplogAppliedTimestampOk() (*BSONTimestamp, bool)`

GetLastOplogAppliedTimestampOk returns a tuple with the LastOplogAppliedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOplogAppliedTimestamp

`func (o *Snapshot) SetLastOplogAppliedTimestamp(v BSONTimestamp)`

SetLastOplogAppliedTimestamp sets LastOplogAppliedTimestamp field to given value.

### HasLastOplogAppliedTimestamp

`func (o *Snapshot) HasLastOplogAppliedTimestamp() bool`

HasLastOplogAppliedTimestamp returns a boolean if a field has been set.

### GetLinks

`func (o *Snapshot) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Snapshot) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Snapshot) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Snapshot) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetParts

`func (o *Snapshot) GetParts() []SnapshotPart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *Snapshot) GetPartsOk() (*[]SnapshotPart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *Snapshot) SetParts(v []SnapshotPart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *Snapshot) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


