// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// SearchIndexDefinitionVersion Object which includes the version number of the index definition and the time that the index definition was created.
type SearchIndexDefinitionVersion struct {
	// The time at which this index definition was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The version number associated with this index definition when it was created.
	Version *int64 `json:"version,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *SearchIndexDefinitionVersion) MarshalJSON() ([]byte, error) {
	type noMethod SearchIndexDefinitionVersion
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewSearchIndexDefinitionVersion instantiates a new SearchIndexDefinitionVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexDefinitionVersion() *SearchIndexDefinitionVersion {
	this := SearchIndexDefinitionVersion{}
	return &this
}

// NewSearchIndexDefinitionVersionWithDefaults instantiates a new SearchIndexDefinitionVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexDefinitionVersionWithDefaults() *SearchIndexDefinitionVersion {
	this := SearchIndexDefinitionVersion{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *SearchIndexDefinitionVersion) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexDefinitionVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SearchIndexDefinitionVersion) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SearchIndexDefinitionVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
	o.NullFields = removeNullField(o.NullFields, "CreatedAt")
}

// SetCreatedAtNil sets CreatedAt to an explicit JSON null when marshaled.
func (o *SearchIndexDefinitionVersion) SetCreatedAtNil() {
	o.CreatedAt = nil
	o.NullFields = addNullField(o.NullFields, "CreatedAt")
}

// GetVersion returns the Version field value if set, zero value otherwise
func (o *SearchIndexDefinitionVersion) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchIndexDefinitionVersion) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}

	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SearchIndexDefinitionVersion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SearchIndexDefinitionVersion) SetVersion(v int64) {
	o.Version = &v
	o.NullFields = removeNullField(o.NullFields, "Version")
}

// SetVersionNil sets Version to an explicit JSON null when marshaled.
func (o *SearchIndexDefinitionVersion) SetVersionNil() {
	o.Version = nil
	o.NullFields = addNullField(o.NullFields, "Version")
}
