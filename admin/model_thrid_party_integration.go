// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// ThridPartyIntegration - Collection of settings that describe third-party integrations.
type ThridPartyIntegration struct {
	Datadog        *Datadog
	MicrosoftTeams *MicrosoftTeams
	NewRelic       *NewRelic
	OpsGenie       *OpsGenie
	PagerDuty      *PagerDuty
	Prometheus     *Prometheus
	Slack          *Slack
	VictorOps      *VictorOps
	Webhook        *Webhook
}

// DatadogAsThridPartyIntegration is a convenience function that returns Datadog wrapped in ThridPartyIntegration
func DatadogAsThridPartyIntegration(v *Datadog) ThridPartyIntegration {
	return ThridPartyIntegration{
		Datadog: v,
	}
}

// MicrosoftTeamsAsThridPartyIntegration is a convenience function that returns MicrosoftTeams wrapped in ThridPartyIntegration
func MicrosoftTeamsAsThridPartyIntegration(v *MicrosoftTeams) ThridPartyIntegration {
	return ThridPartyIntegration{
		MicrosoftTeams: v,
	}
}

// NewRelicAsThridPartyIntegration is a convenience function that returns NewRelic wrapped in ThridPartyIntegration
func NewRelicAsThridPartyIntegration(v *NewRelic) ThridPartyIntegration {
	return ThridPartyIntegration{
		NewRelic: v,
	}
}

// OpsGenieAsThridPartyIntegration is a convenience function that returns OpsGenie wrapped in ThridPartyIntegration
func OpsGenieAsThridPartyIntegration(v *OpsGenie) ThridPartyIntegration {
	return ThridPartyIntegration{
		OpsGenie: v,
	}
}

// PagerDutyAsThridPartyIntegration is a convenience function that returns PagerDuty wrapped in ThridPartyIntegration
func PagerDutyAsThridPartyIntegration(v *PagerDuty) ThridPartyIntegration {
	return ThridPartyIntegration{
		PagerDuty: v,
	}
}

// PrometheusAsThridPartyIntegration is a convenience function that returns Prometheus wrapped in ThridPartyIntegration
func PrometheusAsThridPartyIntegration(v *Prometheus) ThridPartyIntegration {
	return ThridPartyIntegration{
		Prometheus: v,
	}
}

// SlackAsThridPartyIntegration is a convenience function that returns Slack wrapped in ThridPartyIntegration
func SlackAsThridPartyIntegration(v *Slack) ThridPartyIntegration {
	return ThridPartyIntegration{
		Slack: v,
	}
}

// VictorOpsAsThridPartyIntegration is a convenience function that returns VictorOps wrapped in ThridPartyIntegration
func VictorOpsAsThridPartyIntegration(v *VictorOps) ThridPartyIntegration {
	return ThridPartyIntegration{
		VictorOps: v,
	}
}

// WebhookAsThridPartyIntegration is a convenience function that returns Webhook wrapped in ThridPartyIntegration
func WebhookAsThridPartyIntegration(v *Webhook) ThridPartyIntegration {
	return ThridPartyIntegration{
		Webhook: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ThridPartyIntegration) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DATADOG'
	if jsonDict["type"] == "DATADOG" {
		// try to unmarshal JSON data into Datadog
		err = json.Unmarshal(data, &dst.Datadog)
		if err == nil {
			return nil // data stored in dst.Datadog, return on the first match
		} else {
			dst.Datadog = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Datadog: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Datadog'
	if jsonDict["type"] == "Datadog" {
		// try to unmarshal JSON data into Datadog
		err = json.Unmarshal(data, &dst.Datadog)
		if err == nil {
			return nil // data stored in dst.Datadog, return on the first match
		} else {
			dst.Datadog = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Datadog: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MICROSOFT_TEAMS'
	if jsonDict["type"] == "MICROSOFT_TEAMS" {
		// try to unmarshal JSON data into MicrosoftTeams
		err = json.Unmarshal(data, &dst.MicrosoftTeams)
		if err == nil {
			return nil // data stored in dst.MicrosoftTeams, return on the first match
		} else {
			dst.MicrosoftTeams = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as MicrosoftTeams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MicrosoftTeams'
	if jsonDict["type"] == "MicrosoftTeams" {
		// try to unmarshal JSON data into MicrosoftTeams
		err = json.Unmarshal(data, &dst.MicrosoftTeams)
		if err == nil {
			return nil // data stored in dst.MicrosoftTeams, return on the first match
		} else {
			dst.MicrosoftTeams = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as MicrosoftTeams: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NEW_RELIC'
	if jsonDict["type"] == "NEW_RELIC" {
		// try to unmarshal JSON data into NewRelic
		err = json.Unmarshal(data, &dst.NewRelic)
		if err == nil {
			return nil // data stored in dst.NewRelic, return on the first match
		} else {
			dst.NewRelic = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as NewRelic: %s", err.Error())
		}
	}

	// check if the discriminator value is 'NewRelic'
	if jsonDict["type"] == "NewRelic" {
		// try to unmarshal JSON data into NewRelic
		err = json.Unmarshal(data, &dst.NewRelic)
		if err == nil {
			return nil // data stored in dst.NewRelic, return on the first match
		} else {
			dst.NewRelic = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as NewRelic: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OPS_GENIE'
	if jsonDict["type"] == "OPS_GENIE" {
		// try to unmarshal JSON data into OpsGenie
		err = json.Unmarshal(data, &dst.OpsGenie)
		if err == nil {
			return nil // data stored in dst.OpsGenie, return on the first match
		} else {
			dst.OpsGenie = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as OpsGenie: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OpsGenie'
	if jsonDict["type"] == "OpsGenie" {
		// try to unmarshal JSON data into OpsGenie
		err = json.Unmarshal(data, &dst.OpsGenie)
		if err == nil {
			return nil // data stored in dst.OpsGenie, return on the first match
		} else {
			dst.OpsGenie = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as OpsGenie: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PAGER_DUTY'
	if jsonDict["type"] == "PAGER_DUTY" {
		// try to unmarshal JSON data into PagerDuty
		err = json.Unmarshal(data, &dst.PagerDuty)
		if err == nil {
			return nil // data stored in dst.PagerDuty, return on the first match
		} else {
			dst.PagerDuty = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as PagerDuty: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PROMETHEUS'
	if jsonDict["type"] == "PROMETHEUS" {
		// try to unmarshal JSON data into Prometheus
		err = json.Unmarshal(data, &dst.Prometheus)
		if err == nil {
			return nil // data stored in dst.Prometheus, return on the first match
		} else {
			dst.Prometheus = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Prometheus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PagerDuty'
	if jsonDict["type"] == "PagerDuty" {
		// try to unmarshal JSON data into PagerDuty
		err = json.Unmarshal(data, &dst.PagerDuty)
		if err == nil {
			return nil // data stored in dst.PagerDuty, return on the first match
		} else {
			dst.PagerDuty = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as PagerDuty: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Prometheus'
	if jsonDict["type"] == "Prometheus" {
		// try to unmarshal JSON data into Prometheus
		err = json.Unmarshal(data, &dst.Prometheus)
		if err == nil {
			return nil // data stored in dst.Prometheus, return on the first match
		} else {
			dst.Prometheus = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Prometheus: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SLACK'
	if jsonDict["type"] == "SLACK" {
		// try to unmarshal JSON data into Slack
		err = json.Unmarshal(data, &dst.Slack)
		if err == nil {
			return nil // data stored in dst.Slack, return on the first match
		} else {
			dst.Slack = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Slack: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Slack'
	if jsonDict["type"] == "Slack" {
		// try to unmarshal JSON data into Slack
		err = json.Unmarshal(data, &dst.Slack)
		if err == nil {
			return nil // data stored in dst.Slack, return on the first match
		} else {
			dst.Slack = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Slack: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VICTOR_OPS'
	if jsonDict["type"] == "VICTOR_OPS" {
		// try to unmarshal JSON data into VictorOps
		err = json.Unmarshal(data, &dst.VictorOps)
		if err == nil {
			return nil // data stored in dst.VictorOps, return on the first match
		} else {
			dst.VictorOps = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as VictorOps: %s", err.Error())
		}
	}

	// check if the discriminator value is 'VictorOps'
	if jsonDict["type"] == "VictorOps" {
		// try to unmarshal JSON data into VictorOps
		err = json.Unmarshal(data, &dst.VictorOps)
		if err == nil {
			return nil // data stored in dst.VictorOps, return on the first match
		} else {
			dst.VictorOps = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as VictorOps: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WEBHOOK'
	if jsonDict["type"] == "WEBHOOK" {
		// try to unmarshal JSON data into Webhook
		err = json.Unmarshal(data, &dst.Webhook)
		if err == nil {
			return nil // data stored in dst.Webhook, return on the first match
		} else {
			dst.Webhook = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Webhook: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Webhook'
	if jsonDict["type"] == "Webhook" {
		// try to unmarshal JSON data into Webhook
		err = json.Unmarshal(data, &dst.Webhook)
		if err == nil {
			return nil // data stored in dst.Webhook, return on the first match
		} else {
			dst.Webhook = nil
			return fmt.Errorf("failed to unmarshal ThridPartyIntegration as Webhook: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ThridPartyIntegration) MarshalJSON() ([]byte, error) {
	if src.Datadog != nil {
		return json.Marshal(&src.Datadog)
	}

	if src.MicrosoftTeams != nil {
		return json.Marshal(&src.MicrosoftTeams)
	}

	if src.NewRelic != nil {
		return json.Marshal(&src.NewRelic)
	}

	if src.OpsGenie != nil {
		return json.Marshal(&src.OpsGenie)
	}

	if src.PagerDuty != nil {
		return json.Marshal(&src.PagerDuty)
	}

	if src.Prometheus != nil {
		return json.Marshal(&src.Prometheus)
	}

	if src.Slack != nil {
		return json.Marshal(&src.Slack)
	}

	if src.VictorOps != nil {
		return json.Marshal(&src.VictorOps)
	}

	if src.Webhook != nil {
		return json.Marshal(&src.Webhook)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ThridPartyIntegration) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Datadog != nil {
		return obj.Datadog
	}

	if obj.MicrosoftTeams != nil {
		return obj.MicrosoftTeams
	}

	if obj.NewRelic != nil {
		return obj.NewRelic
	}

	if obj.OpsGenie != nil {
		return obj.OpsGenie
	}

	if obj.PagerDuty != nil {
		return obj.PagerDuty
	}

	if obj.Prometheus != nil {
		return obj.Prometheus
	}

	if obj.Slack != nil {
		return obj.Slack
	}

	if obj.VictorOps != nil {
		return obj.VictorOps
	}

	if obj.Webhook != nil {
		return obj.Webhook
	}

	// all schemas are nil
	return nil
}

type NullableThridPartyIntegration struct {
	value *ThridPartyIntegration
	isSet bool
}

func (v NullableThridPartyIntegration) Get() *ThridPartyIntegration {
	return v.value
}

func (v *NullableThridPartyIntegration) Set(val *ThridPartyIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableThridPartyIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableThridPartyIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThridPartyIntegration(val *ThridPartyIntegration) *NullableThridPartyIntegration {
	return &NullableThridPartyIntegration{value: val, isSet: true}
}

func (v NullableThridPartyIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThridPartyIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
