// Code based on the AtlasAPI V2 OpenAPI file

package admin

// DataLakeAtlasStoreReadPreference MongoDB Cloud cluster read preference, which describes how to route read requests to the cluster.
type DataLakeAtlasStoreReadPreference struct {
	// Maximum replication lag, or **staleness**, for reads from secondaries.
	MaxStalenessSeconds *int `json:"maxStalenessSeconds,omitempty"`
	// Read preference mode that specifies to which replica set member to route the read requests.
	Mode *string `json:"mode,omitempty"`
	// List that contains tag sets or tag specification documents. If specified, Atlas Data Lake routes read requests to replica set member or members that are associated with the specified tags.
	TagSets *[][]DataLakeAtlasStoreReadPreferenceTag `json:"tagSets,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *DataLakeAtlasStoreReadPreference) MarshalJSON() ([]byte, error) {
	type noMethod DataLakeAtlasStoreReadPreference
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewDataLakeAtlasStoreReadPreference instantiates a new DataLakeAtlasStoreReadPreference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeAtlasStoreReadPreference() *DataLakeAtlasStoreReadPreference {
	this := DataLakeAtlasStoreReadPreference{}
	return &this
}

// NewDataLakeAtlasStoreReadPreferenceWithDefaults instantiates a new DataLakeAtlasStoreReadPreference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeAtlasStoreReadPreferenceWithDefaults() *DataLakeAtlasStoreReadPreference {
	this := DataLakeAtlasStoreReadPreference{}
	return &this
}

// GetMaxStalenessSeconds returns the MaxStalenessSeconds field value if set, zero value otherwise
func (o *DataLakeAtlasStoreReadPreference) GetMaxStalenessSeconds() int {
	if o == nil || IsNil(o.MaxStalenessSeconds) {
		var ret int
		return ret
	}
	return *o.MaxStalenessSeconds
}

// GetMaxStalenessSecondsOk returns a tuple with the MaxStalenessSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreReadPreference) GetMaxStalenessSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.MaxStalenessSeconds) {
		return nil, false
	}

	return o.MaxStalenessSeconds, true
}

// HasMaxStalenessSeconds returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreReadPreference) HasMaxStalenessSeconds() bool {
	if o != nil && !IsNil(o.MaxStalenessSeconds) {
		return true
	}

	return false
}

// SetMaxStalenessSeconds gets a reference to the given int and assigns it to the MaxStalenessSeconds field.
func (o *DataLakeAtlasStoreReadPreference) SetMaxStalenessSeconds(v int) {
	o.MaxStalenessSeconds = &v
	o.NullFields = removeNullField(o.NullFields, "MaxStalenessSeconds")
}

// SetMaxStalenessSecondsNil sets MaxStalenessSeconds to an explicit JSON null when marshaled.
func (o *DataLakeAtlasStoreReadPreference) SetMaxStalenessSecondsNil() {
	o.MaxStalenessSeconds = nil
	o.NullFields = addNullField(o.NullFields, "MaxStalenessSeconds")
}

// GetMode returns the Mode field value if set, zero value otherwise
func (o *DataLakeAtlasStoreReadPreference) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreReadPreference) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}

	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreReadPreference) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *DataLakeAtlasStoreReadPreference) SetMode(v string) {
	o.Mode = &v
	o.NullFields = removeNullField(o.NullFields, "Mode")
}

// SetModeNil sets Mode to an explicit JSON null when marshaled.
func (o *DataLakeAtlasStoreReadPreference) SetModeNil() {
	o.Mode = nil
	o.NullFields = addNullField(o.NullFields, "Mode")
}

// GetTagSets returns the TagSets field value if set, zero value otherwise
func (o *DataLakeAtlasStoreReadPreference) GetTagSets() [][]DataLakeAtlasStoreReadPreferenceTag {
	if o == nil || IsNil(o.TagSets) {
		var ret [][]DataLakeAtlasStoreReadPreferenceTag
		return ret
	}
	return *o.TagSets
}

// GetTagSetsOk returns a tuple with the TagSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreReadPreference) GetTagSetsOk() (*[][]DataLakeAtlasStoreReadPreferenceTag, bool) {
	if o == nil || IsNil(o.TagSets) {
		return nil, false
	}

	return o.TagSets, true
}

// HasTagSets returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreReadPreference) HasTagSets() bool {
	if o != nil && !IsNil(o.TagSets) {
		return true
	}

	return false
}

// SetTagSets gets a reference to the given [][]DataLakeAtlasStoreReadPreferenceTag and assigns it to the TagSets field.
func (o *DataLakeAtlasStoreReadPreference) SetTagSets(v [][]DataLakeAtlasStoreReadPreferenceTag) {
	o.TagSets = &v
	o.NullFields = removeNullField(o.NullFields, "TagSets")
}

// SetTagSetsNil sets TagSets to an explicit JSON null when marshaled.
func (o *DataLakeAtlasStoreReadPreference) SetTagSetsNil() {
	o.TagSets = nil
	o.NullFields = addNullField(o.NullFields, "TagSets")
}
