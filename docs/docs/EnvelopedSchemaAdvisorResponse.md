# EnvelopedSchemaAdvisorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**SchemaAdvisorResponse**](SchemaAdvisorResponse.md) |  | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Locations** | Pointer to **[]string** | URLs of resources created by this request. | [optional] [readonly] 
**Status** | **int** | HTTP status code returned with this response. | [readonly] 

## Methods

### NewEnvelopedSchemaAdvisorResponse

`func NewEnvelopedSchemaAdvisorResponse(content SchemaAdvisorResponse, status int, ) *EnvelopedSchemaAdvisorResponse`

NewEnvelopedSchemaAdvisorResponse instantiates a new EnvelopedSchemaAdvisorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvelopedSchemaAdvisorResponseWithDefaults

`func NewEnvelopedSchemaAdvisorResponseWithDefaults() *EnvelopedSchemaAdvisorResponse`

NewEnvelopedSchemaAdvisorResponseWithDefaults instantiates a new EnvelopedSchemaAdvisorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *EnvelopedSchemaAdvisorResponse) GetContent() SchemaAdvisorResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EnvelopedSchemaAdvisorResponse) GetContentOk() (*SchemaAdvisorResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EnvelopedSchemaAdvisorResponse) SetContent(v SchemaAdvisorResponse)`

SetContent sets Content field to given value.

### GetLinks

`func (o *EnvelopedSchemaAdvisorResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvelopedSchemaAdvisorResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvelopedSchemaAdvisorResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvelopedSchemaAdvisorResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *EnvelopedSchemaAdvisorResponse) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetLocations

`func (o *EnvelopedSchemaAdvisorResponse) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *EnvelopedSchemaAdvisorResponse) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *EnvelopedSchemaAdvisorResponse) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *EnvelopedSchemaAdvisorResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *EnvelopedSchemaAdvisorResponse) SetLocationsNil()`

SetLocationsNil sets Locations to an explicit JSON null when marshaled, overriding any value previously set with SetLocations. Calling SetLocations again clears the null override.

### GetStatus

`func (o *EnvelopedSchemaAdvisorResponse) GetStatus() int`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvelopedSchemaAdvisorResponse) GetStatusOk() (*int, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvelopedSchemaAdvisorResponse) SetStatus(v int)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


