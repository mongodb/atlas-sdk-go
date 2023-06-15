// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DateCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateCriteria{}

// DateCriteria **DATE criteria.type**.
type DateCriteria struct {
	// Indexed database parameter that stores the date that determines when data moves to the online archive. MongoDB Cloud archives the data when the current date exceeds the date in this database parameter plus the number of days specified through the **expireAfterDays** parameter. Set this parameter when you set `\"criteria.type\" : \"DATE\"`.
	DateField *string `json:"dateField,omitempty"`
	// Syntax used to write the date after which data moves to the online archive. Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when **\"criteria.type\" : \"DATE\"**. You must set **\"criteria.type\" : \"DATE\"** if **\"collectionType\": \"TIMESERIES\"**.
	DateFormat *string `json:"dateFormat,omitempty"`
	// Number of days after the value in the **criteria.dateField** when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set **\"criteria.type\" : \"DATE\"**.
	ExpireAfterDays *int `json:"expireAfterDays,omitempty"`
	// Means by which MongoDB Cloud selects data to archive. Data can be chosen using the age of the data or a MongoDB query. **DATE** selects documents to archive based on a date. **CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **CUSTOM** when `\"collectionType\": \"TIMESERIES\"`.
	Type *string `json:"type,omitempty"`
}

// NewDateCriteria instantiates a new DateCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateCriteria() *DateCriteria {
	this := DateCriteria{}
	var dateFormat string = "ISODATE"
	this.DateFormat = &dateFormat
	return &this
}

// NewDateCriteriaWithDefaults instantiates a new DateCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateCriteriaWithDefaults() *DateCriteria {
	this := DateCriteria{}
	var dateFormat string = "ISODATE"
	this.DateFormat = &dateFormat
	return &this
}

// GetDateField returns the DateField field value if set, zero value otherwise.
func (o *DateCriteria) GetDateField() string {
	if o == nil || IsNil(o.DateField) {
		var ret string
		return ret
	}
	return *o.DateField
}

// GetDateFieldOk returns a tuple with the DateField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateCriteria) GetDateFieldOk() (*string, bool) {
	if o == nil || IsNil(o.DateField) {
		return nil, false
	}
	return o.DateField, true
}

// HasDateField returns a boolean if a field has been set.
func (o *DateCriteria) HasDateField() bool {
	if o != nil && !IsNil(o.DateField) {
		return true
	}

	return false
}

// SetDateField gets a reference to the given string and assigns it to the DateField field.
func (o *DateCriteria) SetDateField(v string) {
	o.DateField = &v
}

// GetDateFormat returns the DateFormat field value if set, zero value otherwise.
func (o *DateCriteria) GetDateFormat() string {
	if o == nil || IsNil(o.DateFormat) {
		var ret string
		return ret
	}
	return *o.DateFormat
}

// GetDateFormatOk returns a tuple with the DateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateCriteria) GetDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DateFormat) {
		return nil, false
	}
	return o.DateFormat, true
}

// HasDateFormat returns a boolean if a field has been set.
func (o *DateCriteria) HasDateFormat() bool {
	if o != nil && !IsNil(o.DateFormat) {
		return true
	}

	return false
}

// SetDateFormat gets a reference to the given string and assigns it to the DateFormat field.
func (o *DateCriteria) SetDateFormat(v string) {
	o.DateFormat = &v
}

// GetExpireAfterDays returns the ExpireAfterDays field value if set, zero value otherwise.
func (o *DateCriteria) GetExpireAfterDays() int {
	if o == nil || IsNil(o.ExpireAfterDays) {
		var ret int
		return ret
	}
	return *o.ExpireAfterDays
}

// GetExpireAfterDaysOk returns a tuple with the ExpireAfterDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateCriteria) GetExpireAfterDaysOk() (*int, bool) {
	if o == nil || IsNil(o.ExpireAfterDays) {
		return nil, false
	}
	return o.ExpireAfterDays, true
}

// HasExpireAfterDays returns a boolean if a field has been set.
func (o *DateCriteria) HasExpireAfterDays() bool {
	if o != nil && !IsNil(o.ExpireAfterDays) {
		return true
	}

	return false
}

// SetExpireAfterDays gets a reference to the given int and assigns it to the ExpireAfterDays field.
func (o *DateCriteria) SetExpireAfterDays(v int) {
	o.ExpireAfterDays = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DateCriteria) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateCriteria) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DateCriteria) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DateCriteria) SetType(v string) {
	o.Type = &v
}

func (o DateCriteria) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DateCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateField) {
		toSerialize["dateField"] = o.DateField
	}
	if !IsNil(o.DateFormat) {
		toSerialize["dateFormat"] = o.DateFormat
	}
	if !IsNil(o.ExpireAfterDays) {
		toSerialize["expireAfterDays"] = o.ExpireAfterDays
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDateCriteria struct {
	value *DateCriteria
	isSet bool
}

func (v NullableDateCriteria) Get() *DateCriteria {
	return v.value
}

func (v *NullableDateCriteria) Set(val *DateCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableDateCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableDateCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateCriteria(val *DateCriteria) *NullableDateCriteria {
	return &NullableDateCriteria{value: val, isSet: true}
}

func (v NullableDateCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
