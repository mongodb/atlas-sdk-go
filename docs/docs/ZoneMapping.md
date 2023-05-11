# ZoneMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Code that represents a location that maps to a zone in your global cluster. MongoDB Cloud represents this location with a ISO 3166-2 location and subdivision codes when possible. | 
**Zone** | **string** | Human-readable label that identifies the zone in your global cluster. This zone maps to a location code. | 

## Methods

### NewZoneMapping

`func NewZoneMapping(location string, zone string, ) *ZoneMapping`

NewZoneMapping instantiates a new ZoneMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneMappingWithDefaults

`func NewZoneMappingWithDefaults() *ZoneMapping`

NewZoneMappingWithDefaults instantiates a new ZoneMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ZoneMapping) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ZoneMapping) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ZoneMapping) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetZone

`func (o *ZoneMapping) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ZoneMapping) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ZoneMapping) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


