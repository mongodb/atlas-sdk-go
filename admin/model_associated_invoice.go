// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AssociatedInvoice An invoice associated with an organization.
type AssociatedInvoice struct {
	// Unique 24-hexadecimal digit identifier for an invoice.
	InvoiceId *string `json:"invoiceId,omitempty"`
	// Unique 24-hexadecimal digit identifier for an organization.
	OrgId *string `json:"orgId,omitempty"`
}

// NewAssociatedInvoice instantiates a new AssociatedInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedInvoice() *AssociatedInvoice {
	this := AssociatedInvoice{}
	return &this
}

// NewAssociatedInvoiceWithDefaults instantiates a new AssociatedInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedInvoiceWithDefaults() *AssociatedInvoice {
	this := AssociatedInvoice{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise
func (o *AssociatedInvoice) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedInvoice) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}

	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *AssociatedInvoice) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *AssociatedInvoice) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise
func (o *AssociatedInvoice) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedInvoice) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}

	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AssociatedInvoice) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *AssociatedInvoice) SetOrgId(v string) {
	o.OrgId = &v
}
