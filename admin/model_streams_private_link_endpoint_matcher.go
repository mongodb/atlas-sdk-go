// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsPrivateLinkEndpointMatcher Rules to apply when comparing a private link endpoint against this alert configuration.
type StreamsPrivateLinkEndpointMatcher struct {
	// Name of the parameter in the target object that MongoDB Cloud checks. The parameter must match all rules for MongoDB Cloud to check for alert configurations.
	FieldName string `json:"fieldName"`
	// Comparison operator to apply when checking the current metric value against **matcher[n].value**. The `REGEX` operator only supports inclusive matches. Use the `NOT_CONTAINS` operator to exclude values.
	Operator string `json:"operator"`
	// Value to match or exceed using the specified `matchers.operator`.
	Value string `json:"value"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsPrivateLinkEndpointMatcher) MarshalJSON() ([]byte, error) {
	type noMethod StreamsPrivateLinkEndpointMatcher
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsPrivateLinkEndpointMatcher instantiates a new StreamsPrivateLinkEndpointMatcher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsPrivateLinkEndpointMatcher(fieldName string, operator string, value string) *StreamsPrivateLinkEndpointMatcher {
	this := StreamsPrivateLinkEndpointMatcher{}
	this.FieldName = fieldName
	this.Operator = operator
	this.Value = value
	return &this
}

// NewStreamsPrivateLinkEndpointMatcherWithDefaults instantiates a new StreamsPrivateLinkEndpointMatcher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsPrivateLinkEndpointMatcherWithDefaults() *StreamsPrivateLinkEndpointMatcher {
	this := StreamsPrivateLinkEndpointMatcher{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *StreamsPrivateLinkEndpointMatcher) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkEndpointMatcher) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *StreamsPrivateLinkEndpointMatcher) SetFieldName(v string) {
	o.FieldName = v
}

// GetOperator returns the Operator field value
func (o *StreamsPrivateLinkEndpointMatcher) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkEndpointMatcher) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *StreamsPrivateLinkEndpointMatcher) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value
func (o *StreamsPrivateLinkEndpointMatcher) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkEndpointMatcher) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *StreamsPrivateLinkEndpointMatcher) SetValue(v string) {
	o.Value = v
}
