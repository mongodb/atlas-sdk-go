/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"time"
)

// checks if the ApiRestoreJobDeliveryView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRestoreJobDeliveryView{}

// ApiRestoreJobDeliveryView Method and details that indicate how to deliver the restored snapshot data.
type ApiRestoreJobDeliveryView struct {
	// Header name to use when downloading the restore, used with `\"delivery.methodName\" : \"HTTP\"`.
	AuthHeader *string `json:"authHeader,omitempty"`
	// Header value to use when downloading the restore, used with `\"delivery.methodName\" : \"HTTP\"`.
	AuthValue *string `json:"authValue,omitempty"`
	// Number of hours after the restore job completes that indicates when the Uniform Resource Locator (URL) for the snapshot download file expires. The resource returns this parameter when `\"delivery.methodName\" : \"HTTP\"`.
	ExpirationHours *int32 `json:"expirationHours,omitempty"`
	// Date and time when the Uniform Resource Locator (URL) for the snapshot download file expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter when `\"delivery.methodName\" : \"HTTP\"`.
	Expires *time.Time `json:"expires,omitempty"`
	// Positive integer that indicates how many times you can use the Uniform Resource Locator (URL) for the snapshot download file. The resource returns this parameter when `\"delivery.methodName\" : \"HTTP\"`.
	MaxDownloads *int32 `json:"maxDownloads,omitempty"`
	// Human-readable label that identifies the means for delivering the data. If you set `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`, you must also set: **delivery.targetGroupId** and **delivery.targetClusterName** or **delivery.targetClusterId**. The response returns `\"delivery.methodName\" : \"HTTP\"` as an automated restore uses HyperText Transport Protocol (HTTP) to deliver the restore job to the target host.
	MethodName string `json:"methodName"`
	// State of the downloadable snapshot file when MongoDB Cloud received this request.
	StatusName *string `json:"statusName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the target cluster. Use the **clusterId** returned in the response body of the **Get All Snapshots** and **Get a Snapshot** endpoints. This parameter applies when `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`.   If the target cluster doesn't have backup enabled, two resources return parameters with empty values:  - **Get All Snapshots** endpoint returns an empty results array without **clusterId** elements - **Get a Snapshot** endpoint doesn't return a **clusterId** parameter.  To return a response with the **clusterId** parameter, either use the **delivery.targetClusterName** parameter or enable backup on the target cluster.
	TargetClusterId *string `json:"targetClusterId,omitempty"`
	// Human-readable label that identifies the target cluster. Use the **clusterName** returned in the response body of the **Get All Snapshots** and **Get a Snapshot** endpoints.  This parameter applies when `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`.  If the target cluster doesn't have backup enabled, two resources return parameters with empty values:  - **Get All Snapshots** endpoint returns an empty results array without **clusterId** elements - **Get a Snapshot** endpoint doesn't return a **clusterId** parameter.  To return a response with the **clusterId** parameter, either use the **delivery.targetClusterName** parameter or enable backup on the target cluster.
	TargetClusterName *string `json:"targetClusterName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project that contains the destination cluster for the restore job. The resource returns this parameter when `\"delivery.methodName\" : \"AUTOMATED_RESTORE\"`.
	TargetGroupId *string `json:"targetGroupId,omitempty"`
	// Uniform Resource Locator (URL) from which you can download the restored snapshot data. Url includes the verification key. The resource returns this parameter when `\"delivery.methodName\" : \"HTTP\"`.
	// Deprecated
	Url *string `json:"url,omitempty"`
	// Uniform Resource Locator (URL) from which you can download the restored snapshot data. This should be preferred over **url**. The verification key must be sent as an HTTP header. The resource returns this parameter when `\"delivery.methodName\" : \"HTTP\"`.
	UrlV2 *string `json:"urlV2,omitempty"`
}

// NewApiRestoreJobDeliveryView instantiates a new ApiRestoreJobDeliveryView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRestoreJobDeliveryView() *ApiRestoreJobDeliveryView {
	this := ApiRestoreJobDeliveryView{}
	return &this
}

// NewApiRestoreJobDeliveryViewWithDefaults instantiates a new ApiRestoreJobDeliveryView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRestoreJobDeliveryViewWithDefaults() *ApiRestoreJobDeliveryView {
	this := ApiRestoreJobDeliveryView{}
	return &this
}

// GetAuthHeader returns the AuthHeader field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetAuthHeader() string {
	if o == nil || IsNil(o.AuthHeader) {
		var ret string
		return ret
	}
	return *o.AuthHeader
}

// GetAuthHeaderOk returns a tuple with the AuthHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetAuthHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.AuthHeader) {
		return nil, false
	}
	return o.AuthHeader, true
}

// HasAuthHeader returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasAuthHeader() bool {
	if o != nil && !IsNil(o.AuthHeader) {
		return true
	}

	return false
}

// SetAuthHeader gets a reference to the given string and assigns it to the AuthHeader field.
func (o *ApiRestoreJobDeliveryView) SetAuthHeader(v string) {
	o.AuthHeader = &v
}

// GetAuthValue returns the AuthValue field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetAuthValue() string {
	if o == nil || IsNil(o.AuthValue) {
		var ret string
		return ret
	}
	return *o.AuthValue
}

// GetAuthValueOk returns a tuple with the AuthValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetAuthValueOk() (*string, bool) {
	if o == nil || IsNil(o.AuthValue) {
		return nil, false
	}
	return o.AuthValue, true
}

// HasAuthValue returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasAuthValue() bool {
	if o != nil && !IsNil(o.AuthValue) {
		return true
	}

	return false
}

// SetAuthValue gets a reference to the given string and assigns it to the AuthValue field.
func (o *ApiRestoreJobDeliveryView) SetAuthValue(v string) {
	o.AuthValue = &v
}

// GetExpirationHours returns the ExpirationHours field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetExpirationHours() int32 {
	if o == nil || IsNil(o.ExpirationHours) {
		var ret int32
		return ret
	}
	return *o.ExpirationHours
}

// GetExpirationHoursOk returns a tuple with the ExpirationHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetExpirationHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpirationHours) {
		return nil, false
	}
	return o.ExpirationHours, true
}

// HasExpirationHours returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasExpirationHours() bool {
	if o != nil && !IsNil(o.ExpirationHours) {
		return true
	}

	return false
}

// SetExpirationHours gets a reference to the given int32 and assigns it to the ExpirationHours field.
func (o *ApiRestoreJobDeliveryView) SetExpirationHours(v int32) {
	o.ExpirationHours = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *ApiRestoreJobDeliveryView) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetMaxDownloads returns the MaxDownloads field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetMaxDownloads() int32 {
	if o == nil || IsNil(o.MaxDownloads) {
		var ret int32
		return ret
	}
	return *o.MaxDownloads
}

// GetMaxDownloadsOk returns a tuple with the MaxDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetMaxDownloadsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDownloads) {
		return nil, false
	}
	return o.MaxDownloads, true
}

// HasMaxDownloads returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasMaxDownloads() bool {
	if o != nil && !IsNil(o.MaxDownloads) {
		return true
	}

	return false
}

// SetMaxDownloads gets a reference to the given int32 and assigns it to the MaxDownloads field.
func (o *ApiRestoreJobDeliveryView) SetMaxDownloads(v int32) {
	o.MaxDownloads = &v
}

// GetMethodName returns the MethodName field value
func (o *ApiRestoreJobDeliveryView) GetMethodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetMethodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodName, true
}

// SetMethodName sets field value
func (o *ApiRestoreJobDeliveryView) SetMethodName(v string) {
	o.MethodName = v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetStatusName() string {
	if o == nil || IsNil(o.StatusName) {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetStatusNameOk() (*string, bool) {
	if o == nil || IsNil(o.StatusName) {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasStatusName() bool {
	if o != nil && !IsNil(o.StatusName) {
		return true
	}

	return false
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *ApiRestoreJobDeliveryView) SetStatusName(v string) {
	o.StatusName = &v
}

// GetTargetClusterId returns the TargetClusterId field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetTargetClusterId() string {
	if o == nil || IsNil(o.TargetClusterId) {
		var ret string
		return ret
	}
	return *o.TargetClusterId
}

// GetTargetClusterIdOk returns a tuple with the TargetClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetTargetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetClusterId) {
		return nil, false
	}
	return o.TargetClusterId, true
}

// HasTargetClusterId returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasTargetClusterId() bool {
	if o != nil && !IsNil(o.TargetClusterId) {
		return true
	}

	return false
}

// SetTargetClusterId gets a reference to the given string and assigns it to the TargetClusterId field.
func (o *ApiRestoreJobDeliveryView) SetTargetClusterId(v string) {
	o.TargetClusterId = &v
}

// GetTargetClusterName returns the TargetClusterName field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetTargetClusterName() string {
	if o == nil || IsNil(o.TargetClusterName) {
		var ret string
		return ret
	}
	return *o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetTargetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetClusterName) {
		return nil, false
	}
	return o.TargetClusterName, true
}

// HasTargetClusterName returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasTargetClusterName() bool {
	if o != nil && !IsNil(o.TargetClusterName) {
		return true
	}

	return false
}

// SetTargetClusterName gets a reference to the given string and assigns it to the TargetClusterName field.
func (o *ApiRestoreJobDeliveryView) SetTargetClusterName(v string) {
	o.TargetClusterName = &v
}

// GetTargetGroupId returns the TargetGroupId field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetTargetGroupId() string {
	if o == nil || IsNil(o.TargetGroupId) {
		var ret string
		return ret
	}
	return *o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetTargetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetGroupId) {
		return nil, false
	}
	return o.TargetGroupId, true
}

// HasTargetGroupId returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasTargetGroupId() bool {
	if o != nil && !IsNil(o.TargetGroupId) {
		return true
	}

	return false
}

// SetTargetGroupId gets a reference to the given string and assigns it to the TargetGroupId field.
func (o *ApiRestoreJobDeliveryView) SetTargetGroupId(v string) {
	o.TargetGroupId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
// Deprecated
func (o *ApiRestoreJobDeliveryView) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApiRestoreJobDeliveryView) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
// Deprecated
func (o *ApiRestoreJobDeliveryView) SetUrl(v string) {
	o.Url = &v
}

// GetUrlV2 returns the UrlV2 field value if set, zero value otherwise.
func (o *ApiRestoreJobDeliveryView) GetUrlV2() string {
	if o == nil || IsNil(o.UrlV2) {
		var ret string
		return ret
	}
	return *o.UrlV2
}

// GetUrlV2Ok returns a tuple with the UrlV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRestoreJobDeliveryView) GetUrlV2Ok() (*string, bool) {
	if o == nil || IsNil(o.UrlV2) {
		return nil, false
	}
	return o.UrlV2, true
}

// HasUrlV2 returns a boolean if a field has been set.
func (o *ApiRestoreJobDeliveryView) HasUrlV2() bool {
	if o != nil && !IsNil(o.UrlV2) {
		return true
	}

	return false
}

// SetUrlV2 gets a reference to the given string and assigns it to the UrlV2 field.
func (o *ApiRestoreJobDeliveryView) SetUrlV2(v string) {
	o.UrlV2 = &v
}

func (o ApiRestoreJobDeliveryView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRestoreJobDeliveryView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: authHeader is readOnly
	// skip: authValue is readOnly
	if !IsNil(o.ExpirationHours) {
		toSerialize["expirationHours"] = o.ExpirationHours
	}
	// skip: expires is readOnly
	if !IsNil(o.MaxDownloads) {
		toSerialize["maxDownloads"] = o.MaxDownloads
	}
	toSerialize["methodName"] = o.MethodName
	// skip: statusName is readOnly
	if !IsNil(o.TargetClusterId) {
		toSerialize["targetClusterId"] = o.TargetClusterId
	}
	if !IsNil(o.TargetClusterName) {
		toSerialize["targetClusterName"] = o.TargetClusterName
	}
	if !IsNil(o.TargetGroupId) {
		toSerialize["targetGroupId"] = o.TargetGroupId
	}
	// skip: url is readOnly
	// skip: urlV2 is readOnly
	return toSerialize, nil
}

type NullableApiRestoreJobDeliveryView struct {
	value *ApiRestoreJobDeliveryView
	isSet bool
}

func (v NullableApiRestoreJobDeliveryView) Get() *ApiRestoreJobDeliveryView {
	return v.value
}

func (v *NullableApiRestoreJobDeliveryView) Set(val *ApiRestoreJobDeliveryView) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRestoreJobDeliveryView) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRestoreJobDeliveryView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRestoreJobDeliveryView(val *ApiRestoreJobDeliveryView) *NullableApiRestoreJobDeliveryView {
	return &NullableApiRestoreJobDeliveryView{value: val, isSet: true}
}

func (v NullableApiRestoreJobDeliveryView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRestoreJobDeliveryView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


