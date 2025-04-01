# AcknowledgeAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedUntil** | Pointer to **time.Time** | Date and time until which this alert has been acknowledged. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert. | [optional] 
**AcknowledgementComment** | Pointer to **string** | Comment that a MongoDB Cloud user submitted when acknowledging the alert. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**UnacknowledgeAlert** | Pointer to **bool** | Flag that indicates to unacknowledge a previously acknowledged alert. By default this value is set to false. If set to true, it will override the acknowledgedUntil parameter. | [optional] 

## Methods

### NewAcknowledgeAlert

`func NewAcknowledgeAlert() *AcknowledgeAlert`

NewAcknowledgeAlert instantiates a new AcknowledgeAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcknowledgeAlertWithDefaults

`func NewAcknowledgeAlertWithDefaults() *AcknowledgeAlert`

NewAcknowledgeAlertWithDefaults instantiates a new AcknowledgeAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedUntil

`func (o *AcknowledgeAlert) GetAcknowledgedUntil() time.Time`

GetAcknowledgedUntil returns the AcknowledgedUntil field if non-nil, zero value otherwise.

### GetAcknowledgedUntilOk

`func (o *AcknowledgeAlert) GetAcknowledgedUntilOk() (*time.Time, bool)`

GetAcknowledgedUntilOk returns a tuple with the AcknowledgedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedUntil

`func (o *AcknowledgeAlert) SetAcknowledgedUntil(v time.Time)`

SetAcknowledgedUntil sets AcknowledgedUntil field to given value.

### HasAcknowledgedUntil

`func (o *AcknowledgeAlert) HasAcknowledgedUntil() bool`

HasAcknowledgedUntil returns a boolean if a field has been set.
### GetAcknowledgementComment

`func (o *AcknowledgeAlert) GetAcknowledgementComment() string`

GetAcknowledgementComment returns the AcknowledgementComment field if non-nil, zero value otherwise.

### GetAcknowledgementCommentOk

`func (o *AcknowledgeAlert) GetAcknowledgementCommentOk() (*string, bool)`

GetAcknowledgementCommentOk returns a tuple with the AcknowledgementComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgementComment

`func (o *AcknowledgeAlert) SetAcknowledgementComment(v string)`

SetAcknowledgementComment sets AcknowledgementComment field to given value.

### HasAcknowledgementComment

`func (o *AcknowledgeAlert) HasAcknowledgementComment() bool`

HasAcknowledgementComment returns a boolean if a field has been set.
### GetLinks

`func (o *AcknowledgeAlert) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AcknowledgeAlert) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AcknowledgeAlert) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AcknowledgeAlert) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetUnacknowledgeAlert

`func (o *AcknowledgeAlert) GetUnacknowledgeAlert() bool`

GetUnacknowledgeAlert returns the UnacknowledgeAlert field if non-nil, zero value otherwise.

### GetUnacknowledgeAlertOk

`func (o *AcknowledgeAlert) GetUnacknowledgeAlertOk() (*bool, bool)`

GetUnacknowledgeAlertOk returns a tuple with the UnacknowledgeAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnacknowledgeAlert

`func (o *AcknowledgeAlert) SetUnacknowledgeAlert(v bool)`

SetUnacknowledgeAlert sets UnacknowledgeAlert field to given value.

### HasUnacknowledgeAlert

`func (o *AcknowledgeAlert) HasUnacknowledgeAlert() bool`

HasUnacknowledgeAlert returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


