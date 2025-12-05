// Code based on the AtlasAPI V2 OpenAPI file

package admin

// Info Information about the MongoDB Atlas Administration API OpenAPI Specification.
type Info struct {
	// Description of the MongoDB Atlas Administration API.
	Description *string  `json:"description,omitempty"`
	License     *License `json:"license,omitempty"`
	// Terms of Service URL.
	TermsOfService *string `json:"termsOfService,omitempty"`
	// Title of the MongoDB Atlas Administration API.
	Title *string `json:"title,omitempty"`
	// Version of the MongoDB Atlas Administration API.
	Version *string `json:"version,omitempty"`
}

// NewInfo instantiates a new Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfo() *Info {
	this := Info{}
	return &this
}

// NewInfoWithDefaults instantiates a new Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoWithDefaults() *Info {
	this := Info{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *Info) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Info) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Info) SetDescription(v string) {
	o.Description = &v
}

// GetLicense returns the License field value if set, zero value otherwise
func (o *Info) GetLicense() License {
	if o == nil || IsNil(o.License) {
		var ret License
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetLicenseOk() (*License, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}

	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *Info) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given License and assigns it to the License field.
func (o *Info) SetLicense(v License) {
	o.License = &v
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise
func (o *Info) GetTermsOfService() string {
	if o == nil || IsNil(o.TermsOfService) {
		var ret string
		return ret
	}
	return *o.TermsOfService
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetTermsOfServiceOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfService) {
		return nil, false
	}

	return o.TermsOfService, true
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *Info) HasTermsOfService() bool {
	if o != nil && !IsNil(o.TermsOfService) {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given string and assigns it to the TermsOfService field.
func (o *Info) SetTermsOfService(v string) {
	o.TermsOfService = &v
}

// GetTitle returns the Title field value if set, zero value otherwise
func (o *Info) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}

	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Info) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Info) SetTitle(v string) {
	o.Title = &v
}

// GetVersion returns the Version field value if set, zero value otherwise
func (o *Info) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Info) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}

	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Info) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Info) SetVersion(v string) {
	o.Version = &v
}
