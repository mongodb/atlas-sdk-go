/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// Notification - One target that MongoDB Cloud sends notifications when an alert triggers.
type Notification struct {
	DatadogNotification *DatadogNotification
	EmailNotification *EmailNotification
	GroupNotification *GroupNotification
	HipChatNotification *HipChatNotification
	MicrosoftTeamsNotification *MicrosoftTeamsNotification
	NDSNotification *NDSNotification
	OpsGenieNotification *OpsGenieNotification
	OrgNotification *OrgNotification
	PagerDutyNotification *PagerDutyNotification
	SMSNotification *SMSNotification
	SlackNotification *SlackNotification
	SummaryNotification *SummaryNotification
	TeamNotification *TeamNotification
	UserNotification *UserNotification
	VictorOpsNotification *VictorOpsNotification
	WebhookNotification *WebhookNotification
}

// DatadogNotificationAsNotification is a convenience function that returns DatadogNotification wrapped in Notification
func DatadogNotificationAsNotification(v *DatadogNotification) Notification {
	return Notification{
		DatadogNotification: v,
	}
}

// EmailNotificationAsNotification is a convenience function that returns EmailNotification wrapped in Notification
func EmailNotificationAsNotification(v *EmailNotification) Notification {
	return Notification{
		EmailNotification: v,
	}
}

// GroupNotificationAsNotification is a convenience function that returns GroupNotification wrapped in Notification
func GroupNotificationAsNotification(v *GroupNotification) Notification {
	return Notification{
		GroupNotification: v,
	}
}

// HipChatNotificationAsNotification is a convenience function that returns HipChatNotification wrapped in Notification
func HipChatNotificationAsNotification(v *HipChatNotification) Notification {
	return Notification{
		HipChatNotification: v,
	}
}

// MicrosoftTeamsNotificationAsNotification is a convenience function that returns MicrosoftTeamsNotification wrapped in Notification
func MicrosoftTeamsNotificationAsNotification(v *MicrosoftTeamsNotification) Notification {
	return Notification{
		MicrosoftTeamsNotification: v,
	}
}

// NDSNotificationAsNotification is a convenience function that returns NDSNotification wrapped in Notification
func NDSNotificationAsNotification(v *NDSNotification) Notification {
	return Notification{
		NDSNotification: v,
	}
}

// OpsGenieNotificationAsNotification is a convenience function that returns OpsGenieNotification wrapped in Notification
func OpsGenieNotificationAsNotification(v *OpsGenieNotification) Notification {
	return Notification{
		OpsGenieNotification: v,
	}
}

// OrgNotificationAsNotification is a convenience function that returns OrgNotification wrapped in Notification
func OrgNotificationAsNotification(v *OrgNotification) Notification {
	return Notification{
		OrgNotification: v,
	}
}

// PagerDutyNotificationAsNotification is a convenience function that returns PagerDutyNotification wrapped in Notification
func PagerDutyNotificationAsNotification(v *PagerDutyNotification) Notification {
	return Notification{
		PagerDutyNotification: v,
	}
}

// SMSNotificationAsNotification is a convenience function that returns SMSNotification wrapped in Notification
func SMSNotificationAsNotification(v *SMSNotification) Notification {
	return Notification{
		SMSNotification: v,
	}
}

// SlackNotificationAsNotification is a convenience function that returns SlackNotification wrapped in Notification
func SlackNotificationAsNotification(v *SlackNotification) Notification {
	return Notification{
		SlackNotification: v,
	}
}

// SummaryNotificationAsNotification is a convenience function that returns SummaryNotification wrapped in Notification
func SummaryNotificationAsNotification(v *SummaryNotification) Notification {
	return Notification{
		SummaryNotification: v,
	}
}

// TeamNotificationAsNotification is a convenience function that returns TeamNotification wrapped in Notification
func TeamNotificationAsNotification(v *TeamNotification) Notification {
	return Notification{
		TeamNotification: v,
	}
}

// UserNotificationAsNotification is a convenience function that returns UserNotification wrapped in Notification
func UserNotificationAsNotification(v *UserNotification) Notification {
	return Notification{
		UserNotification: v,
	}
}

// VictorOpsNotificationAsNotification is a convenience function that returns VictorOpsNotification wrapped in Notification
func VictorOpsNotificationAsNotification(v *VictorOpsNotification) Notification {
	return Notification{
		VictorOpsNotification: v,
	}
}

// WebhookNotificationAsNotification is a convenience function that returns WebhookNotification wrapped in Notification
func WebhookNotificationAsNotification(v *WebhookNotification) Notification {
	return Notification{
		WebhookNotification: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Notification) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DatadogNotification
	err = json.Unmarshal(data, &dst.DatadogNotification)
	if err == nil {
		jsonDatadogNotification, _ := json.Marshal(dst.DatadogNotification)
		if string(jsonDatadogNotification) == "{}" { // empty struct
			dst.DatadogNotification = nil
		} else {
			match++
		}
	} else {
		dst.DatadogNotification = nil
	}

	// try to unmarshal data into EmailNotification
	err = json.Unmarshal(data, &dst.EmailNotification)
	if err == nil {
		jsonEmailNotification, _ := json.Marshal(dst.EmailNotification)
		if string(jsonEmailNotification) == "{}" { // empty struct
			dst.EmailNotification = nil
		} else {
			match++
		}
	} else {
		dst.EmailNotification = nil
	}

	// try to unmarshal data into GroupNotification
	err = json.Unmarshal(data, &dst.GroupNotification)
	if err == nil {
		jsonGroupNotification, _ := json.Marshal(dst.GroupNotification)
		if string(jsonGroupNotification) == "{}" { // empty struct
			dst.GroupNotification = nil
		} else {
			match++
		}
	} else {
		dst.GroupNotification = nil
	}

	// try to unmarshal data into HipChatNotification
	err = json.Unmarshal(data, &dst.HipChatNotification)
	if err == nil {
		jsonHipChatNotification, _ := json.Marshal(dst.HipChatNotification)
		if string(jsonHipChatNotification) == "{}" { // empty struct
			dst.HipChatNotification = nil
		} else {
			match++
		}
	} else {
		dst.HipChatNotification = nil
	}

	// try to unmarshal data into MicrosoftTeamsNotification
	err = json.Unmarshal(data, &dst.MicrosoftTeamsNotification)
	if err == nil {
		jsonMicrosoftTeamsNotification, _ := json.Marshal(dst.MicrosoftTeamsNotification)
		if string(jsonMicrosoftTeamsNotification) == "{}" { // empty struct
			dst.MicrosoftTeamsNotification = nil
		} else {
			match++
		}
	} else {
		dst.MicrosoftTeamsNotification = nil
	}

	// try to unmarshal data into NDSNotification
	err = json.Unmarshal(data, &dst.NDSNotification)
	if err == nil {
		jsonNDSNotification, _ := json.Marshal(dst.NDSNotification)
		if string(jsonNDSNotification) == "{}" { // empty struct
			dst.NDSNotification = nil
		} else {
			match++
		}
	} else {
		dst.NDSNotification = nil
	}

	// try to unmarshal data into OpsGenieNotification
	err = json.Unmarshal(data, &dst.OpsGenieNotification)
	if err == nil {
		jsonOpsGenieNotification, _ := json.Marshal(dst.OpsGenieNotification)
		if string(jsonOpsGenieNotification) == "{}" { // empty struct
			dst.OpsGenieNotification = nil
		} else {
			match++
		}
	} else {
		dst.OpsGenieNotification = nil
	}

	// try to unmarshal data into OrgNotification
	err = json.Unmarshal(data, &dst.OrgNotification)
	if err == nil {
		jsonOrgNotification, _ := json.Marshal(dst.OrgNotification)
		if string(jsonOrgNotification) == "{}" { // empty struct
			dst.OrgNotification = nil
		} else {
			match++
		}
	} else {
		dst.OrgNotification = nil
	}

	// try to unmarshal data into PagerDutyNotification
	err = json.Unmarshal(data, &dst.PagerDutyNotification)
	if err == nil {
		jsonPagerDutyNotification, _ := json.Marshal(dst.PagerDutyNotification)
		if string(jsonPagerDutyNotification) == "{}" { // empty struct
			dst.PagerDutyNotification = nil
		} else {
			match++
		}
	} else {
		dst.PagerDutyNotification = nil
	}

	// try to unmarshal data into SMSNotification
	err = json.Unmarshal(data, &dst.SMSNotification)
	if err == nil {
		jsonSMSNotification, _ := json.Marshal(dst.SMSNotification)
		if string(jsonSMSNotification) == "{}" { // empty struct
			dst.SMSNotification = nil
		} else {
			match++
		}
	} else {
		dst.SMSNotification = nil
	}

	// try to unmarshal data into SlackNotification
	err = json.Unmarshal(data, &dst.SlackNotification)
	if err == nil {
		jsonSlackNotification, _ := json.Marshal(dst.SlackNotification)
		if string(jsonSlackNotification) == "{}" { // empty struct
			dst.SlackNotification = nil
		} else {
			match++
		}
	} else {
		dst.SlackNotification = nil
	}

	// try to unmarshal data into SummaryNotification
	err = json.Unmarshal(data, &dst.SummaryNotification)
	if err == nil {
		jsonSummaryNotification, _ := json.Marshal(dst.SummaryNotification)
		if string(jsonSummaryNotification) == "{}" { // empty struct
			dst.SummaryNotification = nil
		} else {
			match++
		}
	} else {
		dst.SummaryNotification = nil
	}

	// try to unmarshal data into TeamNotification
	err = json.Unmarshal(data, &dst.TeamNotification)
	if err == nil {
		jsonTeamNotification, _ := json.Marshal(dst.TeamNotification)
		if string(jsonTeamNotification) == "{}" { // empty struct
			dst.TeamNotification = nil
		} else {
			match++
		}
	} else {
		dst.TeamNotification = nil
	}

	// try to unmarshal data into UserNotification
	err = json.Unmarshal(data, &dst.UserNotification)
	if err == nil {
		jsonUserNotification, _ := json.Marshal(dst.UserNotification)
		if string(jsonUserNotification) == "{}" { // empty struct
			dst.UserNotification = nil
		} else {
			match++
		}
	} else {
		dst.UserNotification = nil
	}

	// try to unmarshal data into VictorOpsNotification
	err = json.Unmarshal(data, &dst.VictorOpsNotification)
	if err == nil {
		jsonVictorOpsNotification, _ := json.Marshal(dst.VictorOpsNotification)
		if string(jsonVictorOpsNotification) == "{}" { // empty struct
			dst.VictorOpsNotification = nil
		} else {
			match++
		}
	} else {
		dst.VictorOpsNotification = nil
	}

	// try to unmarshal data into WebhookNotification
	err = json.Unmarshal(data, &dst.WebhookNotification)
	if err == nil {
		jsonWebhookNotification, _ := json.Marshal(dst.WebhookNotification)
		if string(jsonWebhookNotification) == "{}" { // empty struct
			dst.WebhookNotification = nil
		} else {
			match++
		}
	} else {
		dst.WebhookNotification = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DatadogNotification = nil
		dst.EmailNotification = nil
		dst.GroupNotification = nil
		dst.HipChatNotification = nil
		dst.MicrosoftTeamsNotification = nil
		dst.NDSNotification = nil
		dst.OpsGenieNotification = nil
		dst.OrgNotification = nil
		dst.PagerDutyNotification = nil
		dst.SMSNotification = nil
		dst.SlackNotification = nil
		dst.SummaryNotification = nil
		dst.TeamNotification = nil
		dst.UserNotification = nil
		dst.VictorOpsNotification = nil
		dst.WebhookNotification = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Notification)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Notification)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Notification) MarshalJSON() ([]byte, error) {
	if src.DatadogNotification != nil {
		return json.Marshal(&src.DatadogNotification)
	}

	if src.EmailNotification != nil {
		return json.Marshal(&src.EmailNotification)
	}

	if src.GroupNotification != nil {
		return json.Marshal(&src.GroupNotification)
	}

	if src.HipChatNotification != nil {
		return json.Marshal(&src.HipChatNotification)
	}

	if src.MicrosoftTeamsNotification != nil {
		return json.Marshal(&src.MicrosoftTeamsNotification)
	}

	if src.NDSNotification != nil {
		return json.Marshal(&src.NDSNotification)
	}

	if src.OpsGenieNotification != nil {
		return json.Marshal(&src.OpsGenieNotification)
	}

	if src.OrgNotification != nil {
		return json.Marshal(&src.OrgNotification)
	}

	if src.PagerDutyNotification != nil {
		return json.Marshal(&src.PagerDutyNotification)
	}

	if src.SMSNotification != nil {
		return json.Marshal(&src.SMSNotification)
	}

	if src.SlackNotification != nil {
		return json.Marshal(&src.SlackNotification)
	}

	if src.SummaryNotification != nil {
		return json.Marshal(&src.SummaryNotification)
	}

	if src.TeamNotification != nil {
		return json.Marshal(&src.TeamNotification)
	}

	if src.UserNotification != nil {
		return json.Marshal(&src.UserNotification)
	}

	if src.VictorOpsNotification != nil {
		return json.Marshal(&src.VictorOpsNotification)
	}

	if src.WebhookNotification != nil {
		return json.Marshal(&src.WebhookNotification)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Notification) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DatadogNotification != nil {
		return obj.DatadogNotification
	}

	if obj.EmailNotification != nil {
		return obj.EmailNotification
	}

	if obj.GroupNotification != nil {
		return obj.GroupNotification
	}

	if obj.HipChatNotification != nil {
		return obj.HipChatNotification
	}

	if obj.MicrosoftTeamsNotification != nil {
		return obj.MicrosoftTeamsNotification
	}

	if obj.NDSNotification != nil {
		return obj.NDSNotification
	}

	if obj.OpsGenieNotification != nil {
		return obj.OpsGenieNotification
	}

	if obj.OrgNotification != nil {
		return obj.OrgNotification
	}

	if obj.PagerDutyNotification != nil {
		return obj.PagerDutyNotification
	}

	if obj.SMSNotification != nil {
		return obj.SMSNotification
	}

	if obj.SlackNotification != nil {
		return obj.SlackNotification
	}

	if obj.SummaryNotification != nil {
		return obj.SummaryNotification
	}

	if obj.TeamNotification != nil {
		return obj.TeamNotification
	}

	if obj.UserNotification != nil {
		return obj.UserNotification
	}

	if obj.VictorOpsNotification != nil {
		return obj.VictorOpsNotification
	}

	if obj.WebhookNotification != nil {
		return obj.WebhookNotification
	}

	// all schemas are nil
	return nil
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


