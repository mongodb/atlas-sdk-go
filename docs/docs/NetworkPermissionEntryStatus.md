# NetworkPermissionEntryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**STATUS** | **string** | State of the access list entry when MongoDB Cloud made this request.  | Status | Activity | |---|---| | &#x60;ACTIVE&#x60; | This access list entry applies to all relevant cloud providers. | | &#x60;PENDING&#x60; | MongoDB Cloud has started to add access list entry. This access list entry may not apply to all cloud providers at the time of this request. | | &#x60;FAILED&#x60; | MongoDB Cloud didn&#39;t succeed in adding this access list entry. |  | [readonly] 

## Methods

### NewNetworkPermissionEntryStatus

`func NewNetworkPermissionEntryStatus(sTATUS string, ) *NetworkPermissionEntryStatus`

NewNetworkPermissionEntryStatus instantiates a new NetworkPermissionEntryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPermissionEntryStatusWithDefaults

`func NewNetworkPermissionEntryStatusWithDefaults() *NetworkPermissionEntryStatus`

NewNetworkPermissionEntryStatusWithDefaults instantiates a new NetworkPermissionEntryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSTATUS

`func (o *NetworkPermissionEntryStatus) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *NetworkPermissionEntryStatus) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *NetworkPermissionEntryStatus) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


