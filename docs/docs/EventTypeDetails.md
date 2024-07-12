# EventTypeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alertable** | Pointer to **bool** | Whether or not this event type can be configured as an alert via the API. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the event type. | [optional] [readonly] 
**EventType** | Pointer to **string** | Enum representation of the event type. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewEventTypeDetails

`func NewEventTypeDetails() *EventTypeDetails`

NewEventTypeDetails instantiates a new EventTypeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTypeDetailsWithDefaults

`func NewEventTypeDetailsWithDefaults() *EventTypeDetails`

NewEventTypeDetailsWithDefaults instantiates a new EventTypeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertable

`func (o *EventTypeDetails) GetAlertable() bool`

GetAlertable returns the Alertable field if non-nil, zero value otherwise.

### GetAlertableOk

`func (o *EventTypeDetails) GetAlertableOk() (*bool, bool)`

GetAlertableOk returns a tuple with the Alertable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertable

`func (o *EventTypeDetails) SetAlertable(v bool)`

SetAlertable sets Alertable field to given value.

### HasAlertable

`func (o *EventTypeDetails) HasAlertable() bool`

HasAlertable returns a boolean if a field has been set.
### GetDescription

`func (o *EventTypeDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventTypeDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventTypeDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventTypeDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetEventType

`func (o *EventTypeDetails) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventTypeDetails) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventTypeDetails) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventTypeDetails) HasEventType() bool`

HasEventType returns a boolean if a field has been set.
### GetLinks

`func (o *EventTypeDetails) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventTypeDetails) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventTypeDetails) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventTypeDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


