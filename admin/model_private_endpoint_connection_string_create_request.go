// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PrivateEndpointConnectionStringCreateRequest Private endpoint connection string to create for one cluster.
type PrivateEndpointConnectionStringCreateRequest struct {
	// Human-readable label of at most 64 characters that identifies this private endpoint connection string. You can change this label without changing the connection string.
	Name string `json:"name"`
	// Flag that indicates whether this private endpoint connection string uses the optimized load-balanced mode. This mode requires a sharded cluster and a private endpoint selection that satisfies the optimized connection string requirements.
	OptimizedModeEnabled *bool `json:"optimizedModeEnabled,omitempty"`
	// Set of private endpoints that this connection string covers. MongoDB Cloud stores this set as a snapshot of your selection and never adds newly created private endpoints to it. The set accepts at most one private endpoint per cloud provider region, and can include private endpoints that the cluster does not use yet. Order is not significant and MongoDB Cloud rejects duplicate values.
	PrivateEndpointIds []string `json:"privateEndpointIds"`
	// Flag that indicates whether this private endpoint connection string routes clients to a subset of the `mongos` processes of the cluster.
	SelectiveMongosEnabled *bool `json:"selectiveMongosEnabled,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PrivateEndpointConnectionStringCreateRequest) MarshalJSON() ([]byte, error) {
	type noMethod PrivateEndpointConnectionStringCreateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPrivateEndpointConnectionStringCreateRequest instantiates a new PrivateEndpointConnectionStringCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateEndpointConnectionStringCreateRequest(name string, privateEndpointIds []string) *PrivateEndpointConnectionStringCreateRequest {
	this := PrivateEndpointConnectionStringCreateRequest{}
	this.Name = name
	var optimizedModeEnabled bool = false
	this.OptimizedModeEnabled = &optimizedModeEnabled
	this.PrivateEndpointIds = privateEndpointIds
	var selectiveMongosEnabled bool = false
	this.SelectiveMongosEnabled = &selectiveMongosEnabled
	return &this
}

// NewPrivateEndpointConnectionStringCreateRequestWithDefaults instantiates a new PrivateEndpointConnectionStringCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateEndpointConnectionStringCreateRequestWithDefaults() *PrivateEndpointConnectionStringCreateRequest {
	this := PrivateEndpointConnectionStringCreateRequest{}
	var optimizedModeEnabled bool = false
	this.OptimizedModeEnabled = &optimizedModeEnabled
	var selectiveMongosEnabled bool = false
	this.SelectiveMongosEnabled = &selectiveMongosEnabled
	return &this
}

// GetName returns the Name field value
func (o *PrivateEndpointConnectionStringCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateEndpointConnectionStringCreateRequest) SetName(v string) {
	o.Name = v
}

// GetOptimizedModeEnabled returns the OptimizedModeEnabled field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringCreateRequest) GetOptimizedModeEnabled() bool {
	if o == nil || IsNil(o.OptimizedModeEnabled) {
		var ret bool
		return ret
	}
	return *o.OptimizedModeEnabled
}

// GetOptimizedModeEnabledOk returns a tuple with the OptimizedModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringCreateRequest) GetOptimizedModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OptimizedModeEnabled) {
		return nil, false
	}

	return o.OptimizedModeEnabled, true
}

// HasOptimizedModeEnabled returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringCreateRequest) HasOptimizedModeEnabled() bool {
	if o != nil && !IsNil(o.OptimizedModeEnabled) {
		return true
	}

	return false
}

// SetOptimizedModeEnabled gets a reference to the given bool and assigns it to the OptimizedModeEnabled field.
func (o *PrivateEndpointConnectionStringCreateRequest) SetOptimizedModeEnabled(v bool) {
	o.OptimizedModeEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "OptimizedModeEnabled")
}

// SetOptimizedModeEnabledNil sets OptimizedModeEnabled to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringCreateRequest) SetOptimizedModeEnabledNil() {
	o.OptimizedModeEnabled = nil
	o.NullFields = addNullField(o.NullFields, "OptimizedModeEnabled")
}

// GetPrivateEndpointIds returns the PrivateEndpointIds field value
func (o *PrivateEndpointConnectionStringCreateRequest) GetPrivateEndpointIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateEndpointIds
}

// GetPrivateEndpointIdsOk returns a tuple with the PrivateEndpointIds field value
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringCreateRequest) GetPrivateEndpointIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateEndpointIds, true
}

// SetPrivateEndpointIds sets field value
func (o *PrivateEndpointConnectionStringCreateRequest) SetPrivateEndpointIds(v []string) {
	o.PrivateEndpointIds = v
}

// GetSelectiveMongosEnabled returns the SelectiveMongosEnabled field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringCreateRequest) GetSelectiveMongosEnabled() bool {
	if o == nil || IsNil(o.SelectiveMongosEnabled) {
		var ret bool
		return ret
	}
	return *o.SelectiveMongosEnabled
}

// GetSelectiveMongosEnabledOk returns a tuple with the SelectiveMongosEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringCreateRequest) GetSelectiveMongosEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SelectiveMongosEnabled) {
		return nil, false
	}

	return o.SelectiveMongosEnabled, true
}

// HasSelectiveMongosEnabled returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringCreateRequest) HasSelectiveMongosEnabled() bool {
	if o != nil && !IsNil(o.SelectiveMongosEnabled) {
		return true
	}

	return false
}

// SetSelectiveMongosEnabled gets a reference to the given bool and assigns it to the SelectiveMongosEnabled field.
func (o *PrivateEndpointConnectionStringCreateRequest) SetSelectiveMongosEnabled(v bool) {
	o.SelectiveMongosEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "SelectiveMongosEnabled")
}

// SetSelectiveMongosEnabledNil sets SelectiveMongosEnabled to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringCreateRequest) SetSelectiveMongosEnabledNil() {
	o.SelectiveMongosEnabled = nil
	o.NullFields = addNullField(o.NullFields, "SelectiveMongosEnabled")
}
