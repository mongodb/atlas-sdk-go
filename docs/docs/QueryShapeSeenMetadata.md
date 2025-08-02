# QueryShapeSeenMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** | The name of the application that this query shape came from. This can be set via the MongoDB connection string. The application name is set to unknown for internal MongoDB queries. | [optional] 
**DriverName** | Pointer to **string** | The name of the MongoDB driver that this query shape was executed from. The driver name is set to unknown for internal MongoDB queries. | [optional] 
**DriverVersion** | Pointer to **string** | The version of the MongoDB driver that this query shape was executed from. The driver version is set to unknown for internal MongoDB queries. | [optional] 
**Timestamp** | Pointer to **int64** | Unix epoch milliseconds of the time. | [optional] 

## Methods

### NewQueryShapeSeenMetadata

`func NewQueryShapeSeenMetadata() *QueryShapeSeenMetadata`

NewQueryShapeSeenMetadata instantiates a new QueryShapeSeenMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryShapeSeenMetadataWithDefaults

`func NewQueryShapeSeenMetadataWithDefaults() *QueryShapeSeenMetadata`

NewQueryShapeSeenMetadataWithDefaults instantiates a new QueryShapeSeenMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *QueryShapeSeenMetadata) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *QueryShapeSeenMetadata) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *QueryShapeSeenMetadata) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *QueryShapeSeenMetadata) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.
### GetDriverName

`func (o *QueryShapeSeenMetadata) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *QueryShapeSeenMetadata) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *QueryShapeSeenMetadata) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *QueryShapeSeenMetadata) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.
### GetDriverVersion

`func (o *QueryShapeSeenMetadata) GetDriverVersion() string`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *QueryShapeSeenMetadata) GetDriverVersionOk() (*string, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *QueryShapeSeenMetadata) SetDriverVersion(v string)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *QueryShapeSeenMetadata) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.
### GetTimestamp

`func (o *QueryShapeSeenMetadata) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *QueryShapeSeenMetadata) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *QueryShapeSeenMetadata) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *QueryShapeSeenMetadata) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


