# Raw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **string** | Unique identifier of event type. | [optional] 
**AlertConfigId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the alert configuration related to the event. | [optional] [readonly] 
**Cid** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project in which the event occurred. | [optional] [readonly] 
**Cre** | Pointer to **time.Time** | Date and time when this event occurred. This parameter expresses its value in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the event. | [optional] 
**Gn** | Pointer to **string** | Human-readable label that identifies the project. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the event. | [optional] [readonly] 
**OrgId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the organization to which these events apply. | [optional] [readonly] 
**OrgName** | Pointer to **string** | Human-readable label that identifies the organization that contains the project. | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewRaw

`func NewRaw() *Raw`

NewRaw instantiates a new Raw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawWithDefaults

`func NewRawWithDefaults() *Raw`

NewRawWithDefaults instantiates a new Raw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *Raw) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *Raw) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *Raw) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *Raw) HasT() bool`

HasT returns a boolean if a field has been set.

### GetAlertConfigId

`func (o *Raw) GetAlertConfigId() string`

GetAlertConfigId returns the AlertConfigId field if non-nil, zero value otherwise.

### GetAlertConfigIdOk

`func (o *Raw) GetAlertConfigIdOk() (*string, bool)`

GetAlertConfigIdOk returns a tuple with the AlertConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigId

`func (o *Raw) SetAlertConfigId(v string)`

SetAlertConfigId sets AlertConfigId field to given value.

### HasAlertConfigId

`func (o *Raw) HasAlertConfigId() bool`

HasAlertConfigId returns a boolean if a field has been set.

### GetCid

`func (o *Raw) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Raw) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Raw) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *Raw) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetCre

`func (o *Raw) GetCre() time.Time`

GetCre returns the Cre field if non-nil, zero value otherwise.

### GetCreOk

`func (o *Raw) GetCreOk() (*time.Time, bool)`

GetCreOk returns a tuple with the Cre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCre

`func (o *Raw) SetCre(v time.Time)`

SetCre sets Cre field to given value.

### HasCre

`func (o *Raw) HasCre() bool`

HasCre returns a boolean if a field has been set.

### GetDescription

`func (o *Raw) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Raw) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Raw) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Raw) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGn

`func (o *Raw) GetGn() string`

GetGn returns the Gn field if non-nil, zero value otherwise.

### GetGnOk

`func (o *Raw) GetGnOk() (*string, bool)`

GetGnOk returns a tuple with the Gn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGn

`func (o *Raw) SetGn(v string)`

SetGn sets Gn field to given value.

### HasGn

`func (o *Raw) HasGn() bool`

HasGn returns a boolean if a field has been set.

### GetId

`func (o *Raw) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Raw) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Raw) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Raw) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *Raw) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Raw) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Raw) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Raw) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *Raw) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *Raw) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *Raw) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *Raw) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetSeverity

`func (o *Raw) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Raw) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Raw) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Raw) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


