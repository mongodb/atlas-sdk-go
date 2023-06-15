// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DatabasePermittedNamespaceResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabasePermittedNamespaceResource{}

// DatabasePermittedNamespaceResource Namespace to which this database user has access.
type DatabasePermittedNamespaceResource struct {
	// Flag that indicates whether to grant the action on the cluster resource. If `true`, MongoDB Cloud ignores the **actions.resources.collection** and **actions.resources.db** parameters.
	Cluster bool `json:"cluster"`
	// Human-readable label that identifies the collection on which you grant the action to one MongoDB user. If you don't set this parameter, you grant the action to all collections in the database specified in the **actions.resources.db** parameter. If you set `\"actions.resources.cluster\" : true`, MongoDB Cloud ignores this parameter.
	Collection string `json:"collection"`
	// Human-readable label that identifies the database on which you grant the action to one MongoDB user. If you set `\"actions.resources.cluster\" : true`, MongoDB Cloud ignores this parameter.
	Db string `json:"db"`
}

// NewDatabasePermittedNamespaceResource instantiates a new DatabasePermittedNamespaceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabasePermittedNamespaceResource(cluster bool, collection string, db string) *DatabasePermittedNamespaceResource {
	this := DatabasePermittedNamespaceResource{}
	this.Cluster = cluster
	this.Collection = collection
	this.Db = db
	return &this
}

// NewDatabasePermittedNamespaceResourceWithDefaults instantiates a new DatabasePermittedNamespaceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabasePermittedNamespaceResourceWithDefaults() *DatabasePermittedNamespaceResource {
	this := DatabasePermittedNamespaceResource{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *DatabasePermittedNamespaceResource) GetCluster() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *DatabasePermittedNamespaceResource) GetClusterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *DatabasePermittedNamespaceResource) SetCluster(v bool) {
	o.Cluster = v
}

// GetCollection returns the Collection field value
func (o *DatabasePermittedNamespaceResource) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *DatabasePermittedNamespaceResource) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *DatabasePermittedNamespaceResource) SetCollection(v string) {
	o.Collection = v
}

// GetDb returns the Db field value
func (o *DatabasePermittedNamespaceResource) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *DatabasePermittedNamespaceResource) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *DatabasePermittedNamespaceResource) SetDb(v string) {
	o.Db = v
}

func (o DatabasePermittedNamespaceResource) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabasePermittedNamespaceResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["collection"] = o.Collection
	toSerialize["db"] = o.Db
	return toSerialize, nil
}

type NullableDatabasePermittedNamespaceResource struct {
	value *DatabasePermittedNamespaceResource
	isSet bool
}

func (v NullableDatabasePermittedNamespaceResource) Get() *DatabasePermittedNamespaceResource {
	return v.value
}

func (v *NullableDatabasePermittedNamespaceResource) Set(val *DatabasePermittedNamespaceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabasePermittedNamespaceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabasePermittedNamespaceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabasePermittedNamespaceResource(val *DatabasePermittedNamespaceResource) *NullableDatabasePermittedNamespaceResource {
	return &NullableDatabasePermittedNamespaceResource{value: val, isSet: true}
}

func (v NullableDatabasePermittedNamespaceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabasePermittedNamespaceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
