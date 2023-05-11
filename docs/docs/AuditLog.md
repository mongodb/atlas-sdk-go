# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditAuthorizationSuccess** | **bool** | Flag that indicates whether someone set auditing to track successful authentications. This only applies to the &#x60;\&quot;atype\&quot; : \&quot;authCheck\&quot;&#x60; audit filter. Setting this parameter to &#x60;true&#x60; degrades cluster performance. | [default to false]
**AuditFilter** | **string** | JSON document that specifies which events to record. Escape any characters that may prevent parsing, such as single or double quotes, using a backslash (&#x60;\\&#x60;). | 
**ConfigurationType** | Pointer to **string** | Human-readable label that displays how to configure the audit filter. | [optional] [readonly] 
**Enabled** | **bool** | Flag that indicates whether someone enabled database auditing for the specified project. | [default to false]

## Methods

### NewAuditLog

`func NewAuditLog(auditAuthorizationSuccess bool, auditFilter string, enabled bool, ) *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditAuthorizationSuccess

`func (o *AuditLog) GetAuditAuthorizationSuccess() bool`

GetAuditAuthorizationSuccess returns the AuditAuthorizationSuccess field if non-nil, zero value otherwise.

### GetAuditAuthorizationSuccessOk

`func (o *AuditLog) GetAuditAuthorizationSuccessOk() (*bool, bool)`

GetAuditAuthorizationSuccessOk returns a tuple with the AuditAuthorizationSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditAuthorizationSuccess

`func (o *AuditLog) SetAuditAuthorizationSuccess(v bool)`

SetAuditAuthorizationSuccess sets AuditAuthorizationSuccess field to given value.


### GetAuditFilter

`func (o *AuditLog) GetAuditFilter() string`

GetAuditFilter returns the AuditFilter field if non-nil, zero value otherwise.

### GetAuditFilterOk

`func (o *AuditLog) GetAuditFilterOk() (*string, bool)`

GetAuditFilterOk returns a tuple with the AuditFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditFilter

`func (o *AuditLog) SetAuditFilter(v string)`

SetAuditFilter sets AuditFilter field to given value.


### GetConfigurationType

`func (o *AuditLog) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *AuditLog) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *AuditLog) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *AuditLog) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetEnabled

`func (o *AuditLog) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuditLog) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuditLog) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


