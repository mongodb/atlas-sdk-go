// Code based on the AtlasAPI V2 OpenAPI file

package admin

// QueryShapeSeenMetadata Metadata about when a query shape was seen.
type QueryShapeSeenMetadata struct {
	// The name of the application that this query shape came from. This can be set via the MongoDB connection string. The application name is set to unknown for internal MongoDB queries.
	ApplicationName *string `json:"applicationName,omitempty"`
	// The name of the MongoDB driver that this query shape was executed from. The driver name is set to unknown for internal MongoDB queries.
	DriverName *string `json:"driverName,omitempty"`
	// The version of the MongoDB driver that this query shape was executed from. The driver version is set to unknown for internal MongoDB queries.
	DriverVersion *string `json:"driverVersion,omitempty"`
	// Unix epoch milliseconds of the time.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *QueryShapeSeenMetadata) MarshalJSON() ([]byte, error) {
	type noMethod QueryShapeSeenMetadata
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewQueryShapeSeenMetadata instantiates a new QueryShapeSeenMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryShapeSeenMetadata() *QueryShapeSeenMetadata {
	this := QueryShapeSeenMetadata{}
	return &this
}

// NewQueryShapeSeenMetadataWithDefaults instantiates a new QueryShapeSeenMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryShapeSeenMetadataWithDefaults() *QueryShapeSeenMetadata {
	this := QueryShapeSeenMetadata{}
	return &this
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise
func (o *QueryShapeSeenMetadata) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShapeSeenMetadata) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}

	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *QueryShapeSeenMetadata) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *QueryShapeSeenMetadata) SetApplicationName(v string) {
	o.ApplicationName = &v
	o.NullFields = removeNullField(o.NullFields, "ApplicationName")
}

// SetApplicationNameNil sets ApplicationName to an explicit JSON null when marshaled.
func (o *QueryShapeSeenMetadata) SetApplicationNameNil() {
	o.ApplicationName = nil
	o.NullFields = addNullField(o.NullFields, "ApplicationName")
}

// GetDriverName returns the DriverName field value if set, zero value otherwise
func (o *QueryShapeSeenMetadata) GetDriverName() string {
	if o == nil || IsNil(o.DriverName) {
		var ret string
		return ret
	}
	return *o.DriverName
}

// GetDriverNameOk returns a tuple with the DriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShapeSeenMetadata) GetDriverNameOk() (*string, bool) {
	if o == nil || IsNil(o.DriverName) {
		return nil, false
	}

	return o.DriverName, true
}

// HasDriverName returns a boolean if a field has been set.
func (o *QueryShapeSeenMetadata) HasDriverName() bool {
	if o != nil && !IsNil(o.DriverName) {
		return true
	}

	return false
}

// SetDriverName gets a reference to the given string and assigns it to the DriverName field.
func (o *QueryShapeSeenMetadata) SetDriverName(v string) {
	o.DriverName = &v
	o.NullFields = removeNullField(o.NullFields, "DriverName")
}

// SetDriverNameNil sets DriverName to an explicit JSON null when marshaled.
func (o *QueryShapeSeenMetadata) SetDriverNameNil() {
	o.DriverName = nil
	o.NullFields = addNullField(o.NullFields, "DriverName")
}

// GetDriverVersion returns the DriverVersion field value if set, zero value otherwise
func (o *QueryShapeSeenMetadata) GetDriverVersion() string {
	if o == nil || IsNil(o.DriverVersion) {
		var ret string
		return ret
	}
	return *o.DriverVersion
}

// GetDriverVersionOk returns a tuple with the DriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShapeSeenMetadata) GetDriverVersionOk() (*string, bool) {
	if o == nil || IsNil(o.DriverVersion) {
		return nil, false
	}

	return o.DriverVersion, true
}

// HasDriverVersion returns a boolean if a field has been set.
func (o *QueryShapeSeenMetadata) HasDriverVersion() bool {
	if o != nil && !IsNil(o.DriverVersion) {
		return true
	}

	return false
}

// SetDriverVersion gets a reference to the given string and assigns it to the DriverVersion field.
func (o *QueryShapeSeenMetadata) SetDriverVersion(v string) {
	o.DriverVersion = &v
	o.NullFields = removeNullField(o.NullFields, "DriverVersion")
}

// SetDriverVersionNil sets DriverVersion to an explicit JSON null when marshaled.
func (o *QueryShapeSeenMetadata) SetDriverVersionNil() {
	o.DriverVersion = nil
	o.NullFields = addNullField(o.NullFields, "DriverVersion")
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise
func (o *QueryShapeSeenMetadata) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryShapeSeenMetadata) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}

	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *QueryShapeSeenMetadata) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *QueryShapeSeenMetadata) SetTimestamp(v int64) {
	o.Timestamp = &v
	o.NullFields = removeNullField(o.NullFields, "Timestamp")
}

// SetTimestampNil sets Timestamp to an explicit JSON null when marshaled.
func (o *QueryShapeSeenMetadata) SetTimestampNil() {
	o.Timestamp = nil
	o.NullFields = addNullField(o.NullFields, "Timestamp")
}
