// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the Slack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Slack{}

// Slack Details to integrate one Slack account with one MongoDB Cloud project.
type Slack struct {
	// Key that allows MongoDB Cloud to access your Slack account.  **NOTE**: After you create a notification which requires an API or integration key, the key appears partially redacted when you:  * View or edit the alert through the Atlas UI.  * Query the alert for the notification through the Atlas Administration API.  **IMPORTANT**: Slack integrations now use the OAuth2 verification method and must  be initially configured, or updated from a legacy integration, through the Atlas  third-party service integrations page. Legacy tokens will soon no longer be  supported.
	ApiToken string `json:"apiToken"`
	// Name of the Slack channel to which MongoDB Cloud sends alert notifications.
	ChannelName NullableString `json:"channelName"`
	// Human-readable label that identifies your Slack team. Set this parameter when you configure a legacy Slack integration.
	TeamName *string `json:"teamName,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewSlack instantiates a new Slack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlack(apiToken string, channelName NullableString) *Slack {
	this := Slack{}
	this.ApiToken = apiToken
	this.ChannelName = channelName
	return &this
}

// NewSlackWithDefaults instantiates a new Slack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackWithDefaults() *Slack {
	this := Slack{}
	return &this
}

// GetApiToken returns the ApiToken field value
func (o *Slack) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *Slack) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value
func (o *Slack) SetApiToken(v string) {
	o.ApiToken = v
}

// GetChannelName returns the ChannelName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Slack) GetChannelName() string {
	if o == nil || o.ChannelName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ChannelName.Get()
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Slack) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelName.Get(), o.ChannelName.IsSet()
}

// SetChannelName sets field value
func (o *Slack) SetChannelName(v string) {
	o.ChannelName.Set(&v)
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *Slack) GetTeamName() string {
	if o == nil || IsNil(o.TeamName) {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slack) GetTeamNameOk() (*string, bool) {
	if o == nil || IsNil(o.TeamName) {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *Slack) HasTeamName() bool {
	if o != nil && !IsNil(o.TeamName) {
		return true
	}

	return false
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *Slack) SetTeamName(v string) {
	o.TeamName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Slack) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Slack) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Slack) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Slack) SetType(v string) {
	o.Type = &v
}

func (o Slack) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Slack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiToken"] = o.ApiToken
	toSerialize["channelName"] = o.ChannelName.Get()
	if !IsNil(o.TeamName) {
		toSerialize["teamName"] = o.TeamName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSlack struct {
	value *Slack
	isSet bool
}

func (v NullableSlack) Get() *Slack {
	return v.value
}

func (v *NullableSlack) Set(val *Slack) {
	v.value = val
	v.isSet = true
}

func (v NullableSlack) IsSet() bool {
	return v.isSet
}

func (v *NullableSlack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlack(val *Slack) *NullableSlack {
	return &NullableSlack{value: val, isSet: true}
}

func (v NullableSlack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
