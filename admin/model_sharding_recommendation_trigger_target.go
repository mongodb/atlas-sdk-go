// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingRecommendationTriggerTarget One entity, such as a shard or namespace, for which the trigger fired.
type ShardingRecommendationTriggerTarget struct {
	// Name of the host that the trigger fired for. This parameter is absent when the trigger did not resolve to a single host.
	// Read only field.
	HostName *string `json:"hostName,omitempty"`
	// Human-readable name that identifies the entity for which the trigger fired.
	// Read only field.
	Name string `json:"name"`
	// Namespace, in `database.collection` form, that the trigger fired for. This parameter is absent for cluster-level and shard-level triggers.
	// Read only field.
	Namespace *string `json:"namespace,omitempty"`
	// Name of the replica set that the trigger fired for. This parameter is absent when the trigger did not resolve to a replica set.
	// Read only field.
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
	// Measured value of the trigger metric for this entity. This parameter is absent when no finite measurement was available.
	// Read only field.
	Value *float64 `json:"value,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingRecommendationTriggerTarget) MarshalJSON() ([]byte, error) {
	type noMethod ShardingRecommendationTriggerTarget
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingRecommendationTriggerTarget instantiates a new ShardingRecommendationTriggerTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingRecommendationTriggerTarget(name string) *ShardingRecommendationTriggerTarget {
	this := ShardingRecommendationTriggerTarget{}
	this.Name = name
	return &this
}

// NewShardingRecommendationTriggerTargetWithDefaults instantiates a new ShardingRecommendationTriggerTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingRecommendationTriggerTargetWithDefaults() *ShardingRecommendationTriggerTarget {
	this := ShardingRecommendationTriggerTarget{}
	return &this
}

// GetHostName returns the HostName field value if set, zero value otherwise
func (o *ShardingRecommendationTriggerTarget) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTriggerTarget) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}

	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ShardingRecommendationTriggerTarget) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ShardingRecommendationTriggerTarget) SetHostName(v string) {
	o.HostName = &v
	o.NullFields = removeNullField(o.NullFields, "HostName")
}

// SetHostNameNil sets HostName to an explicit JSON null when marshaled.
func (o *ShardingRecommendationTriggerTarget) SetHostNameNil() {
	o.HostName = nil
	o.NullFields = addNullField(o.NullFields, "HostName")
}

// GetName returns the Name field value
func (o *ShardingRecommendationTriggerTarget) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTriggerTarget) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShardingRecommendationTriggerTarget) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise
func (o *ShardingRecommendationTriggerTarget) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTriggerTarget) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}

	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ShardingRecommendationTriggerTarget) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ShardingRecommendationTriggerTarget) SetNamespace(v string) {
	o.Namespace = &v
	o.NullFields = removeNullField(o.NullFields, "Namespace")
}

// SetNamespaceNil sets Namespace to an explicit JSON null when marshaled.
func (o *ShardingRecommendationTriggerTarget) SetNamespaceNil() {
	o.Namespace = nil
	o.NullFields = addNullField(o.NullFields, "Namespace")
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *ShardingRecommendationTriggerTarget) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTriggerTarget) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *ShardingRecommendationTriggerTarget) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *ShardingRecommendationTriggerTarget) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
	o.NullFields = removeNullField(o.NullFields, "ReplicaSetName")
}

// SetReplicaSetNameNil sets ReplicaSetName to an explicit JSON null when marshaled.
func (o *ShardingRecommendationTriggerTarget) SetReplicaSetNameNil() {
	o.ReplicaSetName = nil
	o.NullFields = addNullField(o.NullFields, "ReplicaSetName")
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *ShardingRecommendationTriggerTarget) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShardingRecommendationTriggerTarget) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ShardingRecommendationTriggerTarget) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *ShardingRecommendationTriggerTarget) SetValue(v float64) {
	o.Value = &v
	o.NullFields = removeNullField(o.NullFields, "Value")
}

// SetValueNil sets Value to an explicit JSON null when marshaled.
func (o *ShardingRecommendationTriggerTarget) SetValueNil() {
	o.Value = nil
	o.NullFields = addNullField(o.NullFields, "Value")
}
