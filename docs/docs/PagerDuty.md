# PagerDuty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | PagerDuty region that indicates the API Uniform Resource Locator (URL) to use. | [optional] 
**ServiceKey** | **string** | Service key associated with your PagerDuty account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API. | 
**Type** | Pointer to **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type. | [optional] 

## Methods

### NewPagerDuty

`func NewPagerDuty(serviceKey string, ) *PagerDuty`

NewPagerDuty instantiates a new PagerDuty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagerDutyWithDefaults

`func NewPagerDutyWithDefaults() *PagerDuty`

NewPagerDutyWithDefaults instantiates a new PagerDuty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *PagerDuty) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PagerDuty) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PagerDuty) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PagerDuty) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceKey

`func (o *PagerDuty) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *PagerDuty) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *PagerDuty) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetType

`func (o *PagerDuty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PagerDuty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PagerDuty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PagerDuty) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


