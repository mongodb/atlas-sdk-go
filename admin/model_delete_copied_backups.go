// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DeleteCopiedBackups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCopiedBackups{}

// DeleteCopiedBackups Deleted copy setting whose backup copies need to also be deleted.
type DeleteCopiedBackups struct {
	// Human-readable label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Target region for the deleted copy setting whose backup copies you want to delete. Please supply the 'Atlas Region' which can be found under [Cloud Providers](https://www.mongodb.com/docs/atlas/reference/cloud-providers/) 'regions' link.
	RegionName *string `json:"regionName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica setclusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster in One Project and consult the replicationSpecs array [Return One Cluster in One Project](#operation/getLegacyCluster).
	ReplicationSpecId *string `json:"replicationSpecId,omitempty"`
}

// NewDeleteCopiedBackups instantiates a new DeleteCopiedBackups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCopiedBackups() *DeleteCopiedBackups {
	this := DeleteCopiedBackups{}
	return &this
}

// NewDeleteCopiedBackupsWithDefaults instantiates a new DeleteCopiedBackups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCopiedBackupsWithDefaults() *DeleteCopiedBackups {
	this := DeleteCopiedBackups{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *DeleteCopiedBackups) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCopiedBackups) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *DeleteCopiedBackups) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *DeleteCopiedBackups) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *DeleteCopiedBackups) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCopiedBackups) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *DeleteCopiedBackups) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *DeleteCopiedBackups) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise.
func (o *DeleteCopiedBackups) GetReplicationSpecId() string {
	if o == nil || IsNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCopiedBackups) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationSpecId) {
		return nil, false
	}
	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *DeleteCopiedBackups) HasReplicationSpecId() bool {
	if o != nil && !IsNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *DeleteCopiedBackups) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

func (o DeleteCopiedBackups) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DeleteCopiedBackups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.ReplicationSpecId) {
		toSerialize["replicationSpecId"] = o.ReplicationSpecId
	}
	return toSerialize, nil
}

type NullableDeleteCopiedBackups struct {
	value *DeleteCopiedBackups
	isSet bool
}

func (v NullableDeleteCopiedBackups) Get() *DeleteCopiedBackups {
	return v.value
}

func (v *NullableDeleteCopiedBackups) Set(val *DeleteCopiedBackups) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCopiedBackups) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCopiedBackups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCopiedBackups(val *DeleteCopiedBackups) *NullableDeleteCopiedBackups {
	return &NullableDeleteCopiedBackups{value: val, isSet: true}
}

func (v NullableDeleteCopiedBackups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCopiedBackups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
