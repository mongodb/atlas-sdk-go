// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MongoDBAccessLogs Authentication attempt, one per object, made against the cluster.
type MongoDBAccessLogs struct {
	// Flag that indicates whether the response should return successful authentication attempts only.
	AuthResult *bool `json:"authResult,omitempty"`
	// Database against which someone attempted to authenticate.
	// Read only field.
	AuthSource *string `json:"authSource,omitempty"`
	// Reason that the authentication failed. Null if authentication succeeded.
	// Read only field.
	FailureReason *string `json:"failureReason,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the hostname of the target node that received the authentication attempt.
	// Read only field.
	Hostname *string `json:"hostname,omitempty"`
	// Internet Protocol address that attempted to authenticate with the database.
	// Read only field.
	IpAddress *string `json:"ipAddress,omitempty"`
	// Text of the host log concerning the authentication attempt.
	// Read only field.
	LogLine *string `json:"logLine,omitempty"`
	// Date and time when someone made this authentication attempt. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	// Read only field.
	Timestamp *string `json:"timestamp,omitempty"`
	// Username used to authenticate against the database.
	// Read only field.
	Username *string `json:"username,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *MongoDBAccessLogs) MarshalJSON() ([]byte, error) {
	type noMethod MongoDBAccessLogs
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewMongoDBAccessLogs instantiates a new MongoDBAccessLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongoDBAccessLogs() *MongoDBAccessLogs {
	this := MongoDBAccessLogs{}
	return &this
}

// NewMongoDBAccessLogsWithDefaults instantiates a new MongoDBAccessLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongoDBAccessLogsWithDefaults() *MongoDBAccessLogs {
	this := MongoDBAccessLogs{}
	return &this
}

// GetAuthResult returns the AuthResult field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetAuthResult() bool {
	if o == nil || IsNil(o.AuthResult) {
		var ret bool
		return ret
	}
	return *o.AuthResult
}

// GetAuthResultOk returns a tuple with the AuthResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetAuthResultOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthResult) {
		return nil, false
	}

	return o.AuthResult, true
}

// HasAuthResult returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasAuthResult() bool {
	if o != nil && !IsNil(o.AuthResult) {
		return true
	}

	return false
}

// SetAuthResult gets a reference to the given bool and assigns it to the AuthResult field.
func (o *MongoDBAccessLogs) SetAuthResult(v bool) {
	o.AuthResult = &v
	o.NullFields = removeNullField(o.NullFields, "AuthResult")
}

// SetAuthResultNil sets AuthResult to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetAuthResultNil() {
	o.AuthResult = nil
	o.NullFields = addNullField(o.NullFields, "AuthResult")
}

// GetAuthSource returns the AuthSource field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetAuthSource() string {
	if o == nil || IsNil(o.AuthSource) {
		var ret string
		return ret
	}
	return *o.AuthSource
}

// GetAuthSourceOk returns a tuple with the AuthSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetAuthSourceOk() (*string, bool) {
	if o == nil || IsNil(o.AuthSource) {
		return nil, false
	}

	return o.AuthSource, true
}

// HasAuthSource returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasAuthSource() bool {
	if o != nil && !IsNil(o.AuthSource) {
		return true
	}

	return false
}

// SetAuthSource gets a reference to the given string and assigns it to the AuthSource field.
func (o *MongoDBAccessLogs) SetAuthSource(v string) {
	o.AuthSource = &v
	o.NullFields = removeNullField(o.NullFields, "AuthSource")
}

// SetAuthSourceNil sets AuthSource to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetAuthSourceNil() {
	o.AuthSource = nil
	o.NullFields = addNullField(o.NullFields, "AuthSource")
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}

	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *MongoDBAccessLogs) SetFailureReason(v string) {
	o.FailureReason = &v
	o.NullFields = removeNullField(o.NullFields, "FailureReason")
}

// SetFailureReasonNil sets FailureReason to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetFailureReasonNil() {
	o.FailureReason = nil
	o.NullFields = addNullField(o.NullFields, "FailureReason")
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *MongoDBAccessLogs) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetHostname returns the Hostname field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}

	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *MongoDBAccessLogs) SetHostname(v string) {
	o.Hostname = &v
	o.NullFields = removeNullField(o.NullFields, "Hostname")
}

// SetHostnameNil sets Hostname to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetHostnameNil() {
	o.Hostname = nil
	o.NullFields = addNullField(o.NullFields, "Hostname")
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}

	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *MongoDBAccessLogs) SetIpAddress(v string) {
	o.IpAddress = &v
	o.NullFields = removeNullField(o.NullFields, "IpAddress")
}

// SetIpAddressNil sets IpAddress to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetIpAddressNil() {
	o.IpAddress = nil
	o.NullFields = addNullField(o.NullFields, "IpAddress")
}

// GetLogLine returns the LogLine field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetLogLine() string {
	if o == nil || IsNil(o.LogLine) {
		var ret string
		return ret
	}
	return *o.LogLine
}

// GetLogLineOk returns a tuple with the LogLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetLogLineOk() (*string, bool) {
	if o == nil || IsNil(o.LogLine) {
		return nil, false
	}

	return o.LogLine, true
}

// HasLogLine returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasLogLine() bool {
	if o != nil && !IsNil(o.LogLine) {
		return true
	}

	return false
}

// SetLogLine gets a reference to the given string and assigns it to the LogLine field.
func (o *MongoDBAccessLogs) SetLogLine(v string) {
	o.LogLine = &v
	o.NullFields = removeNullField(o.NullFields, "LogLine")
}

// SetLogLineNil sets LogLine to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetLogLineNil() {
	o.LogLine = nil
	o.NullFields = addNullField(o.NullFields, "LogLine")
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}

	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *MongoDBAccessLogs) SetTimestamp(v string) {
	o.Timestamp = &v
	o.NullFields = removeNullField(o.NullFields, "Timestamp")
}

// SetTimestampNil sets Timestamp to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetTimestampNil() {
	o.Timestamp = nil
	o.NullFields = addNullField(o.NullFields, "Timestamp")
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *MongoDBAccessLogs) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoDBAccessLogs) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *MongoDBAccessLogs) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *MongoDBAccessLogs) SetUsername(v string) {
	o.Username = &v
	o.NullFields = removeNullField(o.NullFields, "Username")
}

// SetUsernameNil sets Username to an explicit JSON null when marshaled.
func (o *MongoDBAccessLogs) SetUsernameNil() {
	o.Username = nil
	o.NullFields = addNullField(o.NullFields, "Username")
}
