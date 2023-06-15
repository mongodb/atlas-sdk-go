// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// checks if the DataFederationQueryLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataFederationQueryLimit{}

// DataFederationQueryLimit Details of a query limit for Data Federation. Query limit is the limit on the amount of usage during a time period based on cost.
type DataFederationQueryLimit struct {
	// Amount that indicates the current usage of the limit.
	CurrentUsage *int64 `json:"currentUsage,omitempty"`
	// Default value of the limit.
	DefaultLimit *int64 `json:"defaultLimit,omitempty"`
	// Only used for Data Federation limits. Timestamp that indicates when this usage limit was last modified. This field uses the ISO 8601 timestamp format in UTC.
	LastModifiedDate *time.Time `json:"lastModifiedDate,omitempty"`
	// Maximum value of the limit.
	MaximumLimit *int64 `json:"maximumLimit,omitempty"`
	// Human-readable label that identifies the user-managed limit to modify.
	Name string `json:"name"`
	// Only used for Data Federation limits. Action to take when the usage limit is exceeded. If limit span is set to QUERY, this is ignored because MongoDB Cloud stops the query when it exceeds the usage limit.
	OverrunPolicy *string `json:"overrunPolicy,omitempty"`
	// Amount to set the limit to.
	Value int64 `json:"value"`
}

// NewDataFederationQueryLimit instantiates a new DataFederationQueryLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataFederationQueryLimit(name string, value int64) *DataFederationQueryLimit {
	this := DataFederationQueryLimit{}
	this.Name = name
	this.Value = value
	return &this
}

// NewDataFederationQueryLimitWithDefaults instantiates a new DataFederationQueryLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataFederationQueryLimitWithDefaults() *DataFederationQueryLimit {
	this := DataFederationQueryLimit{}
	return &this
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetCurrentUsage() int64 {
	if o == nil || IsNil(o.CurrentUsage) {
		var ret int64
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetCurrentUsageOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentUsage) {
		return nil, false
	}
	return o.CurrentUsage, true
}

// HasCurrentUsage returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasCurrentUsage() bool {
	if o != nil && !IsNil(o.CurrentUsage) {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given int64 and assigns it to the CurrentUsage field.
func (o *DataFederationQueryLimit) SetCurrentUsage(v int64) {
	o.CurrentUsage = &v
}

// GetDefaultLimit returns the DefaultLimit field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetDefaultLimit() int64 {
	if o == nil || IsNil(o.DefaultLimit) {
		var ret int64
		return ret
	}
	return *o.DefaultLimit
}

// GetDefaultLimitOk returns a tuple with the DefaultLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetDefaultLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultLimit) {
		return nil, false
	}
	return o.DefaultLimit, true
}

// HasDefaultLimit returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasDefaultLimit() bool {
	if o != nil && !IsNil(o.DefaultLimit) {
		return true
	}

	return false
}

// SetDefaultLimit gets a reference to the given int64 and assigns it to the DefaultLimit field.
func (o *DataFederationQueryLimit) SetDefaultLimit(v int64) {
	o.DefaultLimit = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *DataFederationQueryLimit) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetMaximumLimit() int64 {
	if o == nil || IsNil(o.MaximumLimit) {
		var ret int64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetMaximumLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumLimit) {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasMaximumLimit() bool {
	if o != nil && !IsNil(o.MaximumLimit) {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given int64 and assigns it to the MaximumLimit field.
func (o *DataFederationQueryLimit) SetMaximumLimit(v int64) {
	o.MaximumLimit = &v
}

// GetName returns the Name field value
func (o *DataFederationQueryLimit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataFederationQueryLimit) SetName(v string) {
	o.Name = v
}

// GetOverrunPolicy returns the OverrunPolicy field value if set, zero value otherwise.
func (o *DataFederationQueryLimit) GetOverrunPolicy() string {
	if o == nil || IsNil(o.OverrunPolicy) {
		var ret string
		return ret
	}
	return *o.OverrunPolicy
}

// GetOverrunPolicyOk returns a tuple with the OverrunPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetOverrunPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.OverrunPolicy) {
		return nil, false
	}
	return o.OverrunPolicy, true
}

// HasOverrunPolicy returns a boolean if a field has been set.
func (o *DataFederationQueryLimit) HasOverrunPolicy() bool {
	if o != nil && !IsNil(o.OverrunPolicy) {
		return true
	}

	return false
}

// SetOverrunPolicy gets a reference to the given string and assigns it to the OverrunPolicy field.
func (o *DataFederationQueryLimit) SetOverrunPolicy(v string) {
	o.OverrunPolicy = &v
}

// GetValue returns the Value field value
func (o *DataFederationQueryLimit) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DataFederationQueryLimit) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DataFederationQueryLimit) SetValue(v int64) {
	o.Value = v
}

func (o DataFederationQueryLimit) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataFederationQueryLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OverrunPolicy) {
		toSerialize["overrunPolicy"] = o.OverrunPolicy
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableDataFederationQueryLimit struct {
	value *DataFederationQueryLimit
	isSet bool
}

func (v NullableDataFederationQueryLimit) Get() *DataFederationQueryLimit {
	return v.value
}

func (v *NullableDataFederationQueryLimit) Set(val *DataFederationQueryLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFederationQueryLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFederationQueryLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFederationQueryLimit(val *DataFederationQueryLimit) *NullableDataFederationQueryLimit {
	return &NullableDataFederationQueryLimit{value: val, isSet: true}
}

func (v NullableDataFederationQueryLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFederationQueryLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
