# NDSLDAPVerifyConnectivityJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project associated with this Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Request** | Pointer to [**NDSLDAPVerifyConnectivityJobRequestParams**](NDSLDAPVerifyConnectivityJobRequestParams.md) |  | [optional] 
**RequestId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this request to verify an Lightweight Directory Access Protocol (LDAP) configuration. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable string that indicates the status of the Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration. | [optional] [readonly] 
**Validations** | Pointer to [**[]NDSLDAPVerifyConnectivityJobRequestValidation**](NDSLDAPVerifyConnectivityJobRequestValidation.md) | List that contains the validation messages related to the verification of the provided Lightweight Directory Access Protocol (LDAP) over Transport Layer Security (TLS) configuration details. The list contains a document for each test that MongoDB Cloud runs. MongoDB Cloud stops running tests after the first failure. | [optional] [readonly] 

## Methods

### NewNDSLDAPVerifyConnectivityJobRequest

`func NewNDSLDAPVerifyConnectivityJobRequest() *NDSLDAPVerifyConnectivityJobRequest`

NewNDSLDAPVerifyConnectivityJobRequest instantiates a new NDSLDAPVerifyConnectivityJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNDSLDAPVerifyConnectivityJobRequestWithDefaults

`func NewNDSLDAPVerifyConnectivityJobRequestWithDefaults() *NDSLDAPVerifyConnectivityJobRequest`

NewNDSLDAPVerifyConnectivityJobRequestWithDefaults instantiates a new NDSLDAPVerifyConnectivityJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NDSLDAPVerifyConnectivityJobRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NDSLDAPVerifyConnectivityJobRequest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLinks

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NDSLDAPVerifyConnectivityJobRequest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NDSLDAPVerifyConnectivityJobRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRequest

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequest() NDSLDAPVerifyConnectivityJobRequestParams`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequestOk() (*NDSLDAPVerifyConnectivityJobRequestParams, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *NDSLDAPVerifyConnectivityJobRequest) SetRequest(v NDSLDAPVerifyConnectivityJobRequestParams)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *NDSLDAPVerifyConnectivityJobRequest) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetRequestId

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *NDSLDAPVerifyConnectivityJobRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *NDSLDAPVerifyConnectivityJobRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NDSLDAPVerifyConnectivityJobRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NDSLDAPVerifyConnectivityJobRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidations

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetValidations() []NDSLDAPVerifyConnectivityJobRequestValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *NDSLDAPVerifyConnectivityJobRequest) GetValidationsOk() (*[]NDSLDAPVerifyConnectivityJobRequestValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *NDSLDAPVerifyConnectivityJobRequest) SetValidations(v []NDSLDAPVerifyConnectivityJobRequestValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *NDSLDAPVerifyConnectivityJobRequest) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


