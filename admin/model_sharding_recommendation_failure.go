// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendationFailure One recommendation type whose generation failed. Recommendation types that completed are unaffected and still appear in the recommendations object.
type ShardingRecommendationFailure struct {
	// Stable code that classifies the failure. Codes are added over time, so tolerate codes that your client does not recognize.
	// Read only field.
	Code string `json:"code"`
	// Human-readable summary of the failure.
	// Read only field.
	Message string `json:"message"`
	// Recommendation type that failed to generate. Types are added over time, so tolerate types that your client does not recognize.
	// Read only field.
	Type string `json:"type"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingRecommendationFailure) MarshalJSON() ([]byte, error) {
	type noMethod ShardingRecommendationFailure
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingRecommendationFailure instantiates a new ShardingRecommendationFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRecommendationFailure(code string, message string, type_ string) *ShardingRecommendationFailure {
	this := ShardingRecommendationFailure{}
	this.Code = code
	this.Message = message
	this.Type = type_
	return &this
}

// NewShardingRecommendationFailureWithDefaults instantiates a new ShardingRecommendationFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRecommendationFailureWithDefaults() *ShardingRecommendationFailure {
	this := ShardingRecommendationFailure{}
	return &this
}

// GetCode returns the Code field value
func (o *ShardingRecommendationFailure) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationFailure) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ShardingRecommendationFailure) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ShardingRecommendationFailure) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationFailure) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ShardingRecommendationFailure) SetMessage(v string) {
	o.Message = v
}

// GetType returns the Type field value
func (o *ShardingRecommendationFailure) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationFailure) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShardingRecommendationFailure) SetType(v string) {
	o.Type = v
}
