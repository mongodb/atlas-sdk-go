// Code based on the AtlasAPI V2 OpenAPI file

package admin

// TargetOrg struct for TargetOrg
type TargetOrg struct {
	// Link token that contains all the information required to complete the link.
	LinkToken string `json:"linkToken"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *TargetOrg) MarshalJSON() ([]byte, error) {
	type noMethod TargetOrg
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewTargetOrg instantiates a new TargetOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetOrg(linkToken string) *TargetOrg {
	this := TargetOrg{}
	this.LinkToken = linkToken
	return &this
}

// NewTargetOrgWithDefaults instantiates a new TargetOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetOrgWithDefaults() *TargetOrg {
	this := TargetOrg{}
	return &this
}

// GetLinkToken returns the LinkToken field value
func (o *TargetOrg) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *TargetOrg) GetLinkTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *TargetOrg) SetLinkToken(v string) {
	o.LinkToken = v
}
