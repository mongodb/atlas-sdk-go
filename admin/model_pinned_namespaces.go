// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PinnedNamespaces Pinned namespaces view for cluster.
type PinnedNamespaces struct {
	// Unique 24-hexadecimal digit string that identifies the request cluster.
	// Read only field.
	ClusterId *string `json:"clusterId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the request project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// List of all pinned namespaces.
	// Read only field.
	PinnedNamespaces []string `json:"pinnedNamespaces"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PinnedNamespaces) MarshalJSON() ([]byte, error) {
	type noMethod PinnedNamespaces
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPinnedNamespaces instantiates a new PinnedNamespaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinnedNamespaces(pinnedNamespaces []string) *PinnedNamespaces {
	this := PinnedNamespaces{}
	this.PinnedNamespaces = pinnedNamespaces
	return &this
}

// NewPinnedNamespacesWithDefaults instantiates a new PinnedNamespaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinnedNamespacesWithDefaults() *PinnedNamespaces {
	this := PinnedNamespaces{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *PinnedNamespaces) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinnedNamespaces) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *PinnedNamespaces) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *PinnedNamespaces) SetClusterId(v string) {
	o.ClusterId = &v
}

// SetClusterIdNil sets ClusterId to an explicit JSON null when marshaled.
func (o *PinnedNamespaces) SetClusterIdNil() {
	o.ClusterId = nil
	o.NullFields = append(o.NullFields, "ClusterId")
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *PinnedNamespaces) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinnedNamespaces) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *PinnedNamespaces) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *PinnedNamespaces) SetGroupId(v string) {
	o.GroupId = &v
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *PinnedNamespaces) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = append(o.NullFields, "GroupId")
}

// GetPinnedNamespaces returns the PinnedNamespaces field value
func (o *PinnedNamespaces) GetPinnedNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PinnedNamespaces
}

// GetPinnedNamespacesOk returns a tuple with the PinnedNamespaces field value
// and a boolean to check if the value has been set.
func (o *PinnedNamespaces) GetPinnedNamespacesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PinnedNamespaces, true
}

// SetPinnedNamespaces sets field value
func (o *PinnedNamespaces) SetPinnedNamespaces(v []string) {
	o.PinnedNamespaces = v
}
