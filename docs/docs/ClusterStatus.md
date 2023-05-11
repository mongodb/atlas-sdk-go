# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeStatus** | Pointer to **string** | State of cluster at the time of this request. Atlas returns **Applied** if it completed adding a user to, or removing a user from, your cluster. Atlas returns **Pending** if it&#39;s still making the requested user changes. When status is **Pending**, new users can&#39;t log in. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewClusterStatus

`func NewClusterStatus() *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeStatus

`func (o *ClusterStatus) GetChangeStatus() string`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *ClusterStatus) GetChangeStatusOk() (*string, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *ClusterStatus) SetChangeStatus(v string)`

SetChangeStatus sets ChangeStatus field to given value.

### HasChangeStatus

`func (o *ClusterStatus) HasChangeStatus() bool`

HasChangeStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ClusterStatus) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClusterStatus) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClusterStatus) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClusterStatus) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


