// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardAlert Open alert on one of the shard's nodes.
type ShardAlert struct {
	// Human-readable description of the condition that triggered this alert.
	// Read only field.
	Message string `json:"message"`
	// Severity that MongoDB Cloud assigned to this alert.
	// Read only field.
	Severity string `json:"severity"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardAlert) MarshalJSON() ([]byte, error) {
	type noMethod ShardAlert
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardAlert instantiates a new ShardAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardAlert(message string, severity string) *ShardAlert {
	this := ShardAlert{}
	this.Message = message
	this.Severity = severity
	return &this
}

// NewShardAlertWithDefaults instantiates a new ShardAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardAlertWithDefaults() *ShardAlert {
	this := ShardAlert{}
	return &this
}

// GetMessage returns the Message field value
func (o *ShardAlert) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ShardAlert) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ShardAlert) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value
func (o *ShardAlert) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *ShardAlert) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *ShardAlert) SetSeverity(v string) {
	o.Severity = v
}
