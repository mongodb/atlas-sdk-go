# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. This URL often begins with &#x60;https://cloud.mongodb.com/api/atlas&#x60;. | [optional] 
**Rel** | Pointer to **string** | Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. This URL often begins with &#x60;https://cloud.mongodb.com/api/atlas&#x60;. | [optional] 

## Methods

### NewLink

`func NewLink() *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Link) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Link) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Link) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Link) HasHref() bool`

HasHref returns a boolean if a field has been set.
### GetRel

`func (o *Link) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *Link) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *Link) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *Link) HasRel() bool`

HasRel returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


