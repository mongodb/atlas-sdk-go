// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the PeriodicCpsSnapshotSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeriodicCpsSnapshotSource{}

// PeriodicCpsSnapshotSource Scheduled Cloud Provider Snapshot as Source for a Data Lake Pipeline.
type PeriodicCpsSnapshotSource struct {
	// Human-readable name that identifies the cluster.
	ClusterName *string `json:"clusterName,omitempty"`
	// Human-readable name that identifies the collection.
	CollectionName *string `json:"collectionName,omitempty"`
	// Human-readable name that identifies the database.
	DatabaseName *string `json:"databaseName,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal character string that identifies a policy item.
	PolicyItemId *string `json:"policyItemId,omitempty"`
	// Type of ingestion source of this Data Lake Pipeline.
	Type *string `json:"type,omitempty"`
}

// NewPeriodicCpsSnapshotSource instantiates a new PeriodicCpsSnapshotSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriodicCpsSnapshotSource() *PeriodicCpsSnapshotSource {
	this := PeriodicCpsSnapshotSource{}
	return &this
}

// NewPeriodicCpsSnapshotSourceWithDefaults instantiates a new PeriodicCpsSnapshotSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodicCpsSnapshotSourceWithDefaults() *PeriodicCpsSnapshotSource {
	this := PeriodicCpsSnapshotSource{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *PeriodicCpsSnapshotSource) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicCpsSnapshotSource) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *PeriodicCpsSnapshotSource) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *PeriodicCpsSnapshotSource) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise.
func (o *PeriodicCpsSnapshotSource) GetCollectionName() string {
	if o == nil || IsNil(o.CollectionName) {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicCpsSnapshotSource) GetCollectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionName) {
		return nil, false
	}
	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *PeriodicCpsSnapshotSource) HasCollectionName() bool {
	if o != nil && !IsNil(o.CollectionName) {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *PeriodicCpsSnapshotSource) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *PeriodicCpsSnapshotSource) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicCpsSnapshotSource) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *PeriodicCpsSnapshotSource) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *PeriodicCpsSnapshotSource) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *PeriodicCpsSnapshotSource) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicCpsSnapshotSource) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PeriodicCpsSnapshotSource) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PeriodicCpsSnapshotSource) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPolicyItemId returns the PolicyItemId field value if set, zero value otherwise.
func (o *PeriodicCpsSnapshotSource) GetPolicyItemId() string {
	if o == nil || IsNil(o.PolicyItemId) {
		var ret string
		return ret
	}
	return *o.PolicyItemId
}

// GetPolicyItemIdOk returns a tuple with the PolicyItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicCpsSnapshotSource) GetPolicyItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyItemId) {
		return nil, false
	}
	return o.PolicyItemId, true
}

// HasPolicyItemId returns a boolean if a field has been set.
func (o *PeriodicCpsSnapshotSource) HasPolicyItemId() bool {
	if o != nil && !IsNil(o.PolicyItemId) {
		return true
	}

	return false
}

// SetPolicyItemId gets a reference to the given string and assigns it to the PolicyItemId field.
func (o *PeriodicCpsSnapshotSource) SetPolicyItemId(v string) {
	o.PolicyItemId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PeriodicCpsSnapshotSource) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicCpsSnapshotSource) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PeriodicCpsSnapshotSource) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PeriodicCpsSnapshotSource) SetType(v string) {
	o.Type = &v
}

func (o PeriodicCpsSnapshotSource) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PeriodicCpsSnapshotSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !IsNil(o.CollectionName) {
		toSerialize["collectionName"] = o.CollectionName
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !IsNil(o.PolicyItemId) {
		toSerialize["policyItemId"] = o.PolicyItemId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePeriodicCpsSnapshotSource struct {
	value *PeriodicCpsSnapshotSource
	isSet bool
}

func (v NullablePeriodicCpsSnapshotSource) Get() *PeriodicCpsSnapshotSource {
	return v.value
}

func (v *NullablePeriodicCpsSnapshotSource) Set(val *PeriodicCpsSnapshotSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodicCpsSnapshotSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodicCpsSnapshotSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodicCpsSnapshotSource(val *PeriodicCpsSnapshotSource) *NullablePeriodicCpsSnapshotSource {
	return &NullablePeriodicCpsSnapshotSource{value: val, isSet: true}
}

func (v NullablePeriodicCpsSnapshotSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodicCpsSnapshotSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
