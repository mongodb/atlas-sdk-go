# CustomZoneMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomZoneMappings** | Pointer to [**[]ZoneMapping**](ZoneMapping.md) | List that contains comma-separated key value pairs to map zones to geographic regions. These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to the human-readable label for the desired custom zone. MongoDB Cloud maps the ISO 3166-1a2 code to the nearest geographical zone by default. Include this parameter to override the default mappings.  This parameter returns an empty object if no custom zones exist. | [optional] 

## Methods

### NewCustomZoneMappings

`func NewCustomZoneMappings() *CustomZoneMappings`

NewCustomZoneMappings instantiates a new CustomZoneMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomZoneMappingsWithDefaults

`func NewCustomZoneMappingsWithDefaults() *CustomZoneMappings`

NewCustomZoneMappingsWithDefaults instantiates a new CustomZoneMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomZoneMappings

`func (o *CustomZoneMappings) GetCustomZoneMappings() []ZoneMapping`

GetCustomZoneMappings returns the CustomZoneMappings field if non-nil, zero value otherwise.

### GetCustomZoneMappingsOk

`func (o *CustomZoneMappings) GetCustomZoneMappingsOk() (*[]ZoneMapping, bool)`

GetCustomZoneMappingsOk returns a tuple with the CustomZoneMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomZoneMappings

`func (o *CustomZoneMappings) SetCustomZoneMappings(v []ZoneMapping)`

SetCustomZoneMappings sets CustomZoneMappings field to given value.

### HasCustomZoneMappings

`func (o *CustomZoneMappings) HasCustomZoneMappings() bool`

HasCustomZoneMappings returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


