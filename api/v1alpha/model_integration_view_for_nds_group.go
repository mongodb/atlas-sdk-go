/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"fmt"
)

// IntegrationViewForNdsGroup - struct for IntegrationViewForNdsGroup
type IntegrationViewForNdsGroup struct {
	ApiDatadogView *ApiDatadogView
	ApiMicrosoftTeamsView *ApiMicrosoftTeamsView
	ApiNewRelicView *ApiNewRelicView
	ApiOpsGenieView *ApiOpsGenieView
	ApiPagerDutyView *ApiPagerDutyView
	ApiPrometheusView *ApiPrometheusView
	ApiSlackView *ApiSlackView
	ApiVictorOpsView *ApiVictorOpsView
	ApiWebhookView *ApiWebhookView
}

// ApiDatadogViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiDatadogView wrapped in IntegrationViewForNdsGroup
func ApiDatadogViewAsIntegrationViewForNdsGroup(v *ApiDatadogView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiDatadogView: v,
	}
}

// ApiMicrosoftTeamsViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiMicrosoftTeamsView wrapped in IntegrationViewForNdsGroup
func ApiMicrosoftTeamsViewAsIntegrationViewForNdsGroup(v *ApiMicrosoftTeamsView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiMicrosoftTeamsView: v,
	}
}

// ApiNewRelicViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiNewRelicView wrapped in IntegrationViewForNdsGroup
func ApiNewRelicViewAsIntegrationViewForNdsGroup(v *ApiNewRelicView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiNewRelicView: v,
	}
}

// ApiOpsGenieViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiOpsGenieView wrapped in IntegrationViewForNdsGroup
func ApiOpsGenieViewAsIntegrationViewForNdsGroup(v *ApiOpsGenieView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiOpsGenieView: v,
	}
}

// ApiPagerDutyViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiPagerDutyView wrapped in IntegrationViewForNdsGroup
func ApiPagerDutyViewAsIntegrationViewForNdsGroup(v *ApiPagerDutyView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiPagerDutyView: v,
	}
}

// ApiPrometheusViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiPrometheusView wrapped in IntegrationViewForNdsGroup
func ApiPrometheusViewAsIntegrationViewForNdsGroup(v *ApiPrometheusView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiPrometheusView: v,
	}
}

// ApiSlackViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiSlackView wrapped in IntegrationViewForNdsGroup
func ApiSlackViewAsIntegrationViewForNdsGroup(v *ApiSlackView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiSlackView: v,
	}
}

// ApiVictorOpsViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiVictorOpsView wrapped in IntegrationViewForNdsGroup
func ApiVictorOpsViewAsIntegrationViewForNdsGroup(v *ApiVictorOpsView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiVictorOpsView: v,
	}
}

// ApiWebhookViewAsIntegrationViewForNdsGroup is a convenience function that returns ApiWebhookView wrapped in IntegrationViewForNdsGroup
func ApiWebhookViewAsIntegrationViewForNdsGroup(v *ApiWebhookView) IntegrationViewForNdsGroup {
	return IntegrationViewForNdsGroup{
		ApiWebhookView: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IntegrationViewForNdsGroup) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ApiDatadogView
	err = json.Unmarshal(data, &dst.ApiDatadogView)
	if err == nil {
		jsonApiDatadogView, _ := json.Marshal(dst.ApiDatadogView)
		if string(jsonApiDatadogView) == "{}" { // empty struct
			dst.ApiDatadogView = nil
		} else {
			match++
		}
	} else {
		dst.ApiDatadogView = nil
	}

	// try to unmarshal data into ApiMicrosoftTeamsView
	err = json.Unmarshal(data, &dst.ApiMicrosoftTeamsView)
	if err == nil {
		jsonApiMicrosoftTeamsView, _ := json.Marshal(dst.ApiMicrosoftTeamsView)
		if string(jsonApiMicrosoftTeamsView) == "{}" { // empty struct
			dst.ApiMicrosoftTeamsView = nil
		} else {
			match++
		}
	} else {
		dst.ApiMicrosoftTeamsView = nil
	}

	// try to unmarshal data into ApiNewRelicView
	err = json.Unmarshal(data, &dst.ApiNewRelicView)
	if err == nil {
		jsonApiNewRelicView, _ := json.Marshal(dst.ApiNewRelicView)
		if string(jsonApiNewRelicView) == "{}" { // empty struct
			dst.ApiNewRelicView = nil
		} else {
			match++
		}
	} else {
		dst.ApiNewRelicView = nil
	}

	// try to unmarshal data into ApiOpsGenieView
	err = json.Unmarshal(data, &dst.ApiOpsGenieView)
	if err == nil {
		jsonApiOpsGenieView, _ := json.Marshal(dst.ApiOpsGenieView)
		if string(jsonApiOpsGenieView) == "{}" { // empty struct
			dst.ApiOpsGenieView = nil
		} else {
			match++
		}
	} else {
		dst.ApiOpsGenieView = nil
	}

	// try to unmarshal data into ApiPagerDutyView
	err = json.Unmarshal(data, &dst.ApiPagerDutyView)
	if err == nil {
		jsonApiPagerDutyView, _ := json.Marshal(dst.ApiPagerDutyView)
		if string(jsonApiPagerDutyView) == "{}" { // empty struct
			dst.ApiPagerDutyView = nil
		} else {
			match++
		}
	} else {
		dst.ApiPagerDutyView = nil
	}

	// try to unmarshal data into ApiPrometheusView
	err = json.Unmarshal(data, &dst.ApiPrometheusView)
	if err == nil {
		jsonApiPrometheusView, _ := json.Marshal(dst.ApiPrometheusView)
		if string(jsonApiPrometheusView) == "{}" { // empty struct
			dst.ApiPrometheusView = nil
		} else {
			match++
		}
	} else {
		dst.ApiPrometheusView = nil
	}

	// try to unmarshal data into ApiSlackView
	err = json.Unmarshal(data, &dst.ApiSlackView)
	if err == nil {
		jsonApiSlackView, _ := json.Marshal(dst.ApiSlackView)
		if string(jsonApiSlackView) == "{}" { // empty struct
			dst.ApiSlackView = nil
		} else {
			match++
		}
	} else {
		dst.ApiSlackView = nil
	}

	// try to unmarshal data into ApiVictorOpsView
	err = json.Unmarshal(data, &dst.ApiVictorOpsView)
	if err == nil {
		jsonApiVictorOpsView, _ := json.Marshal(dst.ApiVictorOpsView)
		if string(jsonApiVictorOpsView) == "{}" { // empty struct
			dst.ApiVictorOpsView = nil
		} else {
			match++
		}
	} else {
		dst.ApiVictorOpsView = nil
	}

	// try to unmarshal data into ApiWebhookView
	err = json.Unmarshal(data, &dst.ApiWebhookView)
	if err == nil {
		jsonApiWebhookView, _ := json.Marshal(dst.ApiWebhookView)
		if string(jsonApiWebhookView) == "{}" { // empty struct
			dst.ApiWebhookView = nil
		} else {
			match++
		}
	} else {
		dst.ApiWebhookView = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ApiDatadogView = nil
		dst.ApiMicrosoftTeamsView = nil
		dst.ApiNewRelicView = nil
		dst.ApiOpsGenieView = nil
		dst.ApiPagerDutyView = nil
		dst.ApiPrometheusView = nil
		dst.ApiSlackView = nil
		dst.ApiVictorOpsView = nil
		dst.ApiWebhookView = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IntegrationViewForNdsGroup)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IntegrationViewForNdsGroup)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IntegrationViewForNdsGroup) MarshalJSON() ([]byte, error) {
	if src.ApiDatadogView != nil {
		return json.Marshal(&src.ApiDatadogView)
	}

	if src.ApiMicrosoftTeamsView != nil {
		return json.Marshal(&src.ApiMicrosoftTeamsView)
	}

	if src.ApiNewRelicView != nil {
		return json.Marshal(&src.ApiNewRelicView)
	}

	if src.ApiOpsGenieView != nil {
		return json.Marshal(&src.ApiOpsGenieView)
	}

	if src.ApiPagerDutyView != nil {
		return json.Marshal(&src.ApiPagerDutyView)
	}

	if src.ApiPrometheusView != nil {
		return json.Marshal(&src.ApiPrometheusView)
	}

	if src.ApiSlackView != nil {
		return json.Marshal(&src.ApiSlackView)
	}

	if src.ApiVictorOpsView != nil {
		return json.Marshal(&src.ApiVictorOpsView)
	}

	if src.ApiWebhookView != nil {
		return json.Marshal(&src.ApiWebhookView)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IntegrationViewForNdsGroup) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ApiDatadogView != nil {
		return obj.ApiDatadogView
	}

	if obj.ApiMicrosoftTeamsView != nil {
		return obj.ApiMicrosoftTeamsView
	}

	if obj.ApiNewRelicView != nil {
		return obj.ApiNewRelicView
	}

	if obj.ApiOpsGenieView != nil {
		return obj.ApiOpsGenieView
	}

	if obj.ApiPagerDutyView != nil {
		return obj.ApiPagerDutyView
	}

	if obj.ApiPrometheusView != nil {
		return obj.ApiPrometheusView
	}

	if obj.ApiSlackView != nil {
		return obj.ApiSlackView
	}

	if obj.ApiVictorOpsView != nil {
		return obj.ApiVictorOpsView
	}

	if obj.ApiWebhookView != nil {
		return obj.ApiWebhookView
	}

	// all schemas are nil
	return nil
}

type NullableIntegrationViewForNdsGroup struct {
	value *IntegrationViewForNdsGroup
	isSet bool
}

func (v NullableIntegrationViewForNdsGroup) Get() *IntegrationViewForNdsGroup {
	return v.value
}

func (v *NullableIntegrationViewForNdsGroup) Set(val *IntegrationViewForNdsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationViewForNdsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationViewForNdsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationViewForNdsGroup(val *IntegrationViewForNdsGroup) *NullableIntegrationViewForNdsGroup {
	return &NullableIntegrationViewForNdsGroup{value: val, isSet: true}
}

func (v NullableIntegrationViewForNdsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationViewForNdsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


