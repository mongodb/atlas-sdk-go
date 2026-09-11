# SystemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to [**ApiKey**](ApiKey.md) |  | [optional] 
**AppName** | **string** | Human-readable label that identifies the service from which you requested this response. | [readonly] 
**Build** | **string** | Unique 40-hexadecimal digit hash that identifies the latest git commit merged for this application. | [readonly] 
**IpAddress** | **string** | IPv4 or IPv6 address from which you requested this response. Use this value to confirm which address IP access lists evaluate for your requests. | [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Throttling** | **bool** | Flag that indicates whether someone enabled throttling on this service. | [readonly] 
**User** | Pointer to [**AuthenticatedUser**](AuthenticatedUser.md) |  | [optional] 

## Methods

### NewSystemStatus

`func NewSystemStatus(appName string, build string, ipAddress string, throttling bool, ) *SystemStatus`

NewSystemStatus instantiates a new SystemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatusWithDefaults

`func NewSystemStatusWithDefaults() *SystemStatus`

NewSystemStatusWithDefaults instantiates a new SystemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *SystemStatus) GetApiKey() ApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *SystemStatus) GetApiKeyOk() (*ApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *SystemStatus) SetApiKey(v ApiKey)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *SystemStatus) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *SystemStatus) SetApiKeyNil()`

SetApiKeyNil sets ApiKey to an explicit JSON null when marshaled, overriding any value previously set with SetApiKey. Calling SetApiKey again clears the null override.

### GetAppName

`func (o *SystemStatus) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *SystemStatus) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *SystemStatus) SetAppName(v string)`

SetAppName sets AppName field to given value.

### GetBuild

`func (o *SystemStatus) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *SystemStatus) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *SystemStatus) SetBuild(v string)`

SetBuild sets Build field to given value.

### GetIpAddress

`func (o *SystemStatus) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SystemStatus) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SystemStatus) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### GetLinks

`func (o *SystemStatus) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SystemStatus) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SystemStatus) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SystemStatus) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *SystemStatus) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetThrottling

`func (o *SystemStatus) GetThrottling() bool`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *SystemStatus) GetThrottlingOk() (*bool, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *SystemStatus) SetThrottling(v bool)`

SetThrottling sets Throttling field to given value.

### GetUser

`func (o *SystemStatus) GetUser() AuthenticatedUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SystemStatus) GetUserOk() (*AuthenticatedUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SystemStatus) SetUser(v AuthenticatedUser)`

SetUser sets User field to given value.

### HasUser

`func (o *SystemStatus) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SystemStatus) SetUserNil()`

SetUserNil sets User to an explicit JSON null when marshaled, overriding any value previously set with SetUser. Calling SetUser again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


