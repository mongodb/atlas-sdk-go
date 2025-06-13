// Code based on the AtlasAPI V2 OpenAPI file

package admin

// License License information of the MongoDB Atlas Administration API.
type License struct {
	// Name of the license.
	Name *string `json:"name,omitempty"`
	// URL of the license.
	Url *string `json:"url,omitempty"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense() *License {
	this := License{}
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise
func (o *License) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *License) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *License) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise
func (o *License) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}

	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *License) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *License) SetUrl(v string) {
	o.Url = &v
}
