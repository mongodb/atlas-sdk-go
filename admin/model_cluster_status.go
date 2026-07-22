// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterStatus struct for ClusterStatus
type ClusterStatus struct {
	// State of cluster at the time of this request. Atlas returns **Applied** if it completed adding a user to, or removing a user from, your cluster. Atlas returns **Pending** if it's still making the requested user changes. When status is **Pending**, new users can't log in.
	ChangeStatus *string `json:"changeStatus,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ClusterStatus) MarshalJSON() ([]byte, error) {
	type noMethod ClusterStatus
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewClusterStatus instantiates a new ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatus() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// NewClusterStatusWithDefaults instantiates a new ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusWithDefaults() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// GetChangeStatus returns the ChangeStatus field value if set, zero value otherwise
func (o *ClusterStatus) GetChangeStatus() string {
	if o == nil || IsNil(o.ChangeStatus) {
		var ret string
		return ret
	}
	return *o.ChangeStatus
}

// GetChangeStatusOk returns a tuple with the ChangeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetChangeStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeStatus) {
		return nil, false
	}

	return o.ChangeStatus, true
}

// HasChangeStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasChangeStatus() bool {
	if o != nil && !IsNil(o.ChangeStatus) {
		return true
	}

	return false
}

// SetChangeStatus gets a reference to the given string and assigns it to the ChangeStatus field.
func (o *ClusterStatus) SetChangeStatus(v string) {
	o.ChangeStatus = &v
	o.NullFields = removeNullField(o.NullFields, "ChangeStatus")
}

// SetChangeStatusNil sets ChangeStatus to an explicit JSON null when marshaled.
func (o *ClusterStatus) SetChangeStatusNil() {
	o.ChangeStatus = nil
	o.NullFields = addNullField(o.NullFields, "ChangeStatus")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *ClusterStatus) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ClusterStatus) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ClusterStatus) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *ClusterStatus) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}
