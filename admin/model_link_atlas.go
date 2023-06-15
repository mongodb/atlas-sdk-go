// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the LinkAtlas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkAtlas{}

// LinkAtlas struct for LinkAtlas
type LinkAtlas struct {
	// Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. This URL often begins with `https://cloud.mongodb.com/api/atlas`.
	Href *string `json:"href,omitempty"`
	// Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. This URL often begins with `https://cloud.mongodb.com/api/atlas`.
	Rel *string `json:"rel,omitempty"`
}

// NewLinkAtlas instantiates a new LinkAtlas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkAtlas() *LinkAtlas {
	this := LinkAtlas{}
	return &this
}

// NewLinkAtlasWithDefaults instantiates a new LinkAtlas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkAtlasWithDefaults() *LinkAtlas {
	this := LinkAtlas{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LinkAtlas) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkAtlas) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LinkAtlas) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LinkAtlas) SetHref(v string) {
	o.Href = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *LinkAtlas) GetRel() string {
	if o == nil || IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkAtlas) GetRelOk() (*string, bool) {
	if o == nil || IsNil(o.Rel) {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *LinkAtlas) HasRel() bool {
	if o != nil && !IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *LinkAtlas) SetRel(v string) {
	o.Rel = &v
}

func (o LinkAtlas) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LinkAtlas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	return toSerialize, nil
}

type NullableLinkAtlas struct {
	value *LinkAtlas
	isSet bool
}

func (v NullableLinkAtlas) Get() *LinkAtlas {
	return v.value
}

func (v *NullableLinkAtlas) Set(val *LinkAtlas) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkAtlas) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkAtlas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkAtlas(val *LinkAtlas) *NullableLinkAtlas {
	return &NullableLinkAtlas{value: val, isSet: true}
}

func (v NullableLinkAtlas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkAtlas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
