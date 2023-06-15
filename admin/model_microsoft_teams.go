// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the MicrosoftTeams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrosoftTeams{}

// MicrosoftTeams Details to integrate one Microsoft Teams account with one MongoDB Cloud project.
type MicrosoftTeams struct {
	// Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.  **NOTE**: When you view or edit the alert for a Microsoft Teams notification, the URL appears partially redacted.
	MicrosoftTeamsWebhookUrl string `json:"microsoftTeamsWebhookUrl"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
}

// NewMicrosoftTeams instantiates a new MicrosoftTeams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftTeams(microsoftTeamsWebhookUrl string) *MicrosoftTeams {
	this := MicrosoftTeams{}
	this.MicrosoftTeamsWebhookUrl = microsoftTeamsWebhookUrl
	return &this
}

// NewMicrosoftTeamsWithDefaults instantiates a new MicrosoftTeams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftTeamsWithDefaults() *MicrosoftTeams {
	this := MicrosoftTeams{}
	return &this
}

// GetMicrosoftTeamsWebhookUrl returns the MicrosoftTeamsWebhookUrl field value
func (o *MicrosoftTeams) GetMicrosoftTeamsWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicrosoftTeamsWebhookUrl
}

// GetMicrosoftTeamsWebhookUrlOk returns a tuple with the MicrosoftTeamsWebhookUrl field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeams) GetMicrosoftTeamsWebhookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicrosoftTeamsWebhookUrl, true
}

// SetMicrosoftTeamsWebhookUrl sets field value
func (o *MicrosoftTeams) SetMicrosoftTeamsWebhookUrl(v string) {
	o.MicrosoftTeamsWebhookUrl = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MicrosoftTeams) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeams) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftTeams) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MicrosoftTeams) SetType(v string) {
	o.Type = &v
}

func (o MicrosoftTeams) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o MicrosoftTeams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["microsoftTeamsWebhookUrl"] = o.MicrosoftTeamsWebhookUrl
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMicrosoftTeams struct {
	value *MicrosoftTeams
	isSet bool
}

func (v NullableMicrosoftTeams) Get() *MicrosoftTeams {
	return v.value
}

func (v *NullableMicrosoftTeams) Set(val *MicrosoftTeams) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftTeams) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftTeams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftTeams(val *MicrosoftTeams) *NullableMicrosoftTeams {
	return &NullableMicrosoftTeams{value: val, isSet: true}
}

func (v NullableMicrosoftTeams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftTeams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
