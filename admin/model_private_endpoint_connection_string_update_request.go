// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PrivateEndpointConnectionStringUpdateRequest Private endpoint connection string fields to update for one cluster. Only `name` and `privateEndpointIds` can change; a request that carries any other field, including the connection method properties (`optimizedModeEnabled`, `selectiveMongosEnabled`), is rejected with a 400 naming the offending fields.
type PrivateEndpointConnectionStringUpdateRequest struct {
	// Human-readable label of at most 64 characters that identifies this private endpoint connection string. You can change this label without changing the connection string.
	Name *string `json:"name,omitempty"`
	// Set of private endpoints that this connection string covers. MongoDB Cloud replaces the stored snapshot with this set and never adds newly created private endpoints to it. The set accepts at most one private endpoint per cloud provider region, and can include private endpoints that the cluster does not use yet. Order is not significant and MongoDB Cloud rejects duplicate values.
	PrivateEndpointIds *[]string `json:"privateEndpointIds,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PrivateEndpointConnectionStringUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod PrivateEndpointConnectionStringUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPrivateEndpointConnectionStringUpdateRequest instantiates a new PrivateEndpointConnectionStringUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateEndpointConnectionStringUpdateRequest() *PrivateEndpointConnectionStringUpdateRequest {
	this := PrivateEndpointConnectionStringUpdateRequest{}
	return &this
}

// NewPrivateEndpointConnectionStringUpdateRequestWithDefaults instantiates a new PrivateEndpointConnectionStringUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateEndpointConnectionStringUpdateRequestWithDefaults() *PrivateEndpointConnectionStringUpdateRequest {
	this := PrivateEndpointConnectionStringUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrivateEndpointConnectionStringUpdateRequest) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringUpdateRequest) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetPrivateEndpointIds returns the PrivateEndpointIds field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringUpdateRequest) GetPrivateEndpointIds() []string {
	if o == nil || IsNil(o.PrivateEndpointIds) {
		var ret []string
		return ret
	}
	return *o.PrivateEndpointIds
}

// GetPrivateEndpointIdsOk returns a tuple with the PrivateEndpointIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringUpdateRequest) GetPrivateEndpointIdsOk() (*[]string, bool) {
	if o == nil || IsNil(o.PrivateEndpointIds) {
		return nil, false
	}

	return o.PrivateEndpointIds, true
}

// HasPrivateEndpointIds returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringUpdateRequest) HasPrivateEndpointIds() bool {
	if o != nil && !IsNil(o.PrivateEndpointIds) {
		return true
	}

	return false
}

// SetPrivateEndpointIds gets a reference to the given []string and assigns it to the PrivateEndpointIds field.
func (o *PrivateEndpointConnectionStringUpdateRequest) SetPrivateEndpointIds(v []string) {
	o.PrivateEndpointIds = &v
	o.NullFields = removeNullField(o.NullFields, "PrivateEndpointIds")
}

// SetPrivateEndpointIdsNil sets PrivateEndpointIds to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringUpdateRequest) SetPrivateEndpointIdsNil() {
	o.PrivateEndpointIds = nil
	o.NullFields = addNullField(o.NullFields, "PrivateEndpointIds")
}
