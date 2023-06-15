// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DataLakeOnlineArchiveStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeOnlineArchiveStore{}

// DataLakeOnlineArchiveStore struct for DataLakeOnlineArchiveStore
type DataLakeOnlineArchiveStore struct {
	// ID of the Cluster the Online Archive belongs to.
	ClusterId string `json:"clusterId"`
	// Name of the Cluster the Online Archive belongs to.
	ClusterName string `json:"clusterName"`
	// ID of the Project the Online Archive belongs to.
	ProjectId string `json:"projectId"`
	// Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an `M10` or higher cluster.
	Name     *string `json:"name,omitempty"`
	Provider string  `json:"provider"`
}

// NewDataLakeOnlineArchiveStore instantiates a new DataLakeOnlineArchiveStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeOnlineArchiveStore(clusterId string, clusterName string, projectId string, provider string) *DataLakeOnlineArchiveStore {
	this := DataLakeOnlineArchiveStore{}
	this.ClusterId = clusterId
	this.ClusterName = clusterName
	this.ProjectId = projectId
	this.Provider = provider
	return &this
}

// NewDataLakeOnlineArchiveStoreWithDefaults instantiates a new DataLakeOnlineArchiveStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeOnlineArchiveStoreWithDefaults() *DataLakeOnlineArchiveStore {
	this := DataLakeOnlineArchiveStore{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *DataLakeOnlineArchiveStore) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *DataLakeOnlineArchiveStore) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *DataLakeOnlineArchiveStore) SetClusterId(v string) {
	o.ClusterId = v
}

// GetClusterName returns the ClusterName field value
func (o *DataLakeOnlineArchiveStore) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *DataLakeOnlineArchiveStore) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *DataLakeOnlineArchiveStore) SetClusterName(v string) {
	o.ClusterName = v
}

// GetProjectId returns the ProjectId field value
func (o *DataLakeOnlineArchiveStore) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DataLakeOnlineArchiveStore) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DataLakeOnlineArchiveStore) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataLakeOnlineArchiveStore) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeOnlineArchiveStore) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeOnlineArchiveStore) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeOnlineArchiveStore) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value
func (o *DataLakeOnlineArchiveStore) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *DataLakeOnlineArchiveStore) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *DataLakeOnlineArchiveStore) SetProvider(v string) {
	o.Provider = v
}

func (o DataLakeOnlineArchiveStore) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeOnlineArchiveStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterId"] = o.ClusterId
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["projectId"] = o.ProjectId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["provider"] = o.Provider
	return toSerialize, nil
}

type NullableDataLakeOnlineArchiveStore struct {
	value *DataLakeOnlineArchiveStore
	isSet bool
}

func (v NullableDataLakeOnlineArchiveStore) Get() *DataLakeOnlineArchiveStore {
	return v.value
}

func (v *NullableDataLakeOnlineArchiveStore) Set(val *DataLakeOnlineArchiveStore) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeOnlineArchiveStore) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeOnlineArchiveStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeOnlineArchiveStore(val *DataLakeOnlineArchiveStore) *NullableDataLakeOnlineArchiveStore {
	return &NullableDataLakeOnlineArchiveStore{value: val, isSet: true}
}

func (v NullableDataLakeOnlineArchiveStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeOnlineArchiveStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
