// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DataLakeAtlasStoreInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeAtlasStoreInstance{}

// DataLakeAtlasStoreInstance struct for DataLakeAtlasStoreInstance
type DataLakeAtlasStoreInstance struct {
	// Human-readable label of the MongoDB Cloud cluster on which the store is based.
	ClusterName *string `json:"clusterName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project.
	ProjectId      *string                           `json:"projectId,omitempty"`
	ReadPreference *DataLakeAtlasStoreReadPreference `json:"readPreference,omitempty"`
	// Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an `M10` or higher cluster.
	Name     *string `json:"name,omitempty"`
	Provider string  `json:"provider"`
}

// NewDataLakeAtlasStoreInstance instantiates a new DataLakeAtlasStoreInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeAtlasStoreInstance(provider string) *DataLakeAtlasStoreInstance {
	this := DataLakeAtlasStoreInstance{}
	this.Provider = provider
	return &this
}

// NewDataLakeAtlasStoreInstanceWithDefaults instantiates a new DataLakeAtlasStoreInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeAtlasStoreInstanceWithDefaults() *DataLakeAtlasStoreInstance {
	this := DataLakeAtlasStoreInstance{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *DataLakeAtlasStoreInstance) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreInstance) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreInstance) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *DataLakeAtlasStoreInstance) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DataLakeAtlasStoreInstance) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreInstance) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreInstance) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DataLakeAtlasStoreInstance) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetReadPreference returns the ReadPreference field value if set, zero value otherwise.
func (o *DataLakeAtlasStoreInstance) GetReadPreference() DataLakeAtlasStoreReadPreference {
	if o == nil || IsNil(o.ReadPreference) {
		var ret DataLakeAtlasStoreReadPreference
		return ret
	}
	return *o.ReadPreference
}

// GetReadPreferenceOk returns a tuple with the ReadPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreInstance) GetReadPreferenceOk() (*DataLakeAtlasStoreReadPreference, bool) {
	if o == nil || IsNil(o.ReadPreference) {
		return nil, false
	}
	return o.ReadPreference, true
}

// HasReadPreference returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreInstance) HasReadPreference() bool {
	if o != nil && !IsNil(o.ReadPreference) {
		return true
	}

	return false
}

// SetReadPreference gets a reference to the given DataLakeAtlasStoreReadPreference and assigns it to the ReadPreference field.
func (o *DataLakeAtlasStoreInstance) SetReadPreference(v DataLakeAtlasStoreReadPreference) {
	o.ReadPreference = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataLakeAtlasStoreInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeAtlasStoreInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeAtlasStoreInstance) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value
func (o *DataLakeAtlasStoreInstance) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *DataLakeAtlasStoreInstance) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *DataLakeAtlasStoreInstance) SetProvider(v string) {
	o.Provider = v
}

func (o DataLakeAtlasStoreInstance) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeAtlasStoreInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !IsNil(o.ReadPreference) {
		toSerialize["readPreference"] = o.ReadPreference
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["provider"] = o.Provider
	return toSerialize, nil
}

type NullableDataLakeAtlasStoreInstance struct {
	value *DataLakeAtlasStoreInstance
	isSet bool
}

func (v NullableDataLakeAtlasStoreInstance) Get() *DataLakeAtlasStoreInstance {
	return v.value
}

func (v *NullableDataLakeAtlasStoreInstance) Set(val *DataLakeAtlasStoreInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeAtlasStoreInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeAtlasStoreInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeAtlasStoreInstance(val *DataLakeAtlasStoreInstance) *NullableDataLakeAtlasStoreInstance {
	return &NullableDataLakeAtlasStoreInstance{value: val, isSet: true}
}

func (v NullableDataLakeAtlasStoreInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeAtlasStoreInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
