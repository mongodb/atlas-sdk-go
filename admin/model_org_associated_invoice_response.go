// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgAssociatedInvoiceResponse Response containing associated invoices for an organization.
type OrgAssociatedInvoiceResponse struct {
	// List of invoices associated with the organization for the specified period.
	// Read only field.
	AssociatedInvoices *[]AssociatedInvoice `json:"associatedInvoices,omitempty"`
	// Two-digit number that represents the month of the associated invoices.
	// Read only field.
	Month *string `json:"month,omitempty"`
	// Four-digit number that represents the year of the associated invoices.
	// Read only field.
	Year *string `json:"year,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OrgAssociatedInvoiceResponse) MarshalJSON() ([]byte, error) {
	type noMethod OrgAssociatedInvoiceResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOrgAssociatedInvoiceResponse instantiates a new OrgAssociatedInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgAssociatedInvoiceResponse() *OrgAssociatedInvoiceResponse {
	this := OrgAssociatedInvoiceResponse{}
	return &this
}

// NewOrgAssociatedInvoiceResponseWithDefaults instantiates a new OrgAssociatedInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgAssociatedInvoiceResponseWithDefaults() *OrgAssociatedInvoiceResponse {
	this := OrgAssociatedInvoiceResponse{}
	return &this
}

// GetAssociatedInvoices returns the AssociatedInvoices field value if set, zero value otherwise
func (o *OrgAssociatedInvoiceResponse) GetAssociatedInvoices() []AssociatedInvoice {
	if o == nil || IsNil(o.AssociatedInvoices) {
		var ret []AssociatedInvoice
		return ret
	}
	return *o.AssociatedInvoices
}

// GetAssociatedInvoicesOk returns a tuple with the AssociatedInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAssociatedInvoiceResponse) GetAssociatedInvoicesOk() (*[]AssociatedInvoice, bool) {
	if o == nil || IsNil(o.AssociatedInvoices) {
		return nil, false
	}

	return o.AssociatedInvoices, true
}

// HasAssociatedInvoices returns a boolean if a field has been set.
func (o *OrgAssociatedInvoiceResponse) HasAssociatedInvoices() bool {
	if o != nil && !IsNil(o.AssociatedInvoices) {
		return true
	}

	return false
}

// SetAssociatedInvoices gets a reference to the given []AssociatedInvoice and assigns it to the AssociatedInvoices field.
func (o *OrgAssociatedInvoiceResponse) SetAssociatedInvoices(v []AssociatedInvoice) {
	o.AssociatedInvoices = &v
	o.NullFields = removeNullField(o.NullFields, "AssociatedInvoices")
}

// SetAssociatedInvoicesNil sets AssociatedInvoices to an explicit JSON null when marshaled.
func (o *OrgAssociatedInvoiceResponse) SetAssociatedInvoicesNil() {
	o.AssociatedInvoices = nil
	o.NullFields = addNullField(o.NullFields, "AssociatedInvoices")
}

// GetMonth returns the Month field value if set, zero value otherwise
func (o *OrgAssociatedInvoiceResponse) GetMonth() string {
	if o == nil || IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAssociatedInvoiceResponse) GetMonthOk() (*string, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}

	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *OrgAssociatedInvoiceResponse) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *OrgAssociatedInvoiceResponse) SetMonth(v string) {
	o.Month = &v
	o.NullFields = removeNullField(o.NullFields, "Month")
}

// SetMonthNil sets Month to an explicit JSON null when marshaled.
func (o *OrgAssociatedInvoiceResponse) SetMonthNil() {
	o.Month = nil
	o.NullFields = addNullField(o.NullFields, "Month")
}

// GetYear returns the Year field value if set, zero value otherwise
func (o *OrgAssociatedInvoiceResponse) GetYear() string {
	if o == nil || IsNil(o.Year) {
		var ret string
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAssociatedInvoiceResponse) GetYearOk() (*string, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}

	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *OrgAssociatedInvoiceResponse) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given string and assigns it to the Year field.
func (o *OrgAssociatedInvoiceResponse) SetYear(v string) {
	o.Year = &v
	o.NullFields = removeNullField(o.NullFields, "Year")
}

// SetYearNil sets Year to an explicit JSON null when marshaled.
func (o *OrgAssociatedInvoiceResponse) SetYearNil() {
	o.Year = nil
	o.NullFields = addNullField(o.NullFields, "Year")
}
