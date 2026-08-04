// Code based on the AtlasAPI V2 OpenAPI file

package admin

// InvoiceReportRequest struct for InvoiceReportRequest
type InvoiceReportRequest struct {
	// Version of the report format specification.
	FormatSpecVersion *string `json:"formatSpecVersion,omitempty"`
	// Format of the report.
	ReportFormat string `json:"reportFormat"`
	// Type of report to generate.
	ReportType string `json:"reportType"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *InvoiceReportRequest) MarshalJSON() ([]byte, error) {
	type noMethod InvoiceReportRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewInvoiceReportRequest instantiates a new InvoiceReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceReportRequest(reportFormat string, reportType string) *InvoiceReportRequest {
	this := InvoiceReportRequest{}
	this.ReportFormat = reportFormat
	this.ReportType = reportType
	return &this
}

// NewInvoiceReportRequestWithDefaults instantiates a new InvoiceReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceReportRequestWithDefaults() *InvoiceReportRequest {
	this := InvoiceReportRequest{}
	return &this
}

// GetFormatSpecVersion returns the FormatSpecVersion field value if set, zero value otherwise
func (o *InvoiceReportRequest) GetFormatSpecVersion() string {
	if o == nil || IsNil(o.FormatSpecVersion) {
		var ret string
		return ret
	}
	return *o.FormatSpecVersion
}

// GetFormatSpecVersionOk returns a tuple with the FormatSpecVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReportRequest) GetFormatSpecVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FormatSpecVersion) {
		return nil, false
	}

	return o.FormatSpecVersion, true
}

// HasFormatSpecVersion returns a boolean if a field has been set.
func (o *InvoiceReportRequest) HasFormatSpecVersion() bool {
	if o != nil && !IsNil(o.FormatSpecVersion) {
		return true
	}

	return false
}

// SetFormatSpecVersion gets a reference to the given string and assigns it to the FormatSpecVersion field.
func (o *InvoiceReportRequest) SetFormatSpecVersion(v string) {
	o.FormatSpecVersion = &v
	o.NullFields = removeNullField(o.NullFields, "FormatSpecVersion")
}

// SetFormatSpecVersionNil sets FormatSpecVersion to an explicit JSON null when marshaled.
func (o *InvoiceReportRequest) SetFormatSpecVersionNil() {
	o.FormatSpecVersion = nil
	o.NullFields = addNullField(o.NullFields, "FormatSpecVersion")
}

// GetReportFormat returns the ReportFormat field value
func (o *InvoiceReportRequest) GetReportFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFormat
}

// GetReportFormatOk returns a tuple with the ReportFormat field value
// and a boolean to check if the value has been set.
func (o *InvoiceReportRequest) GetReportFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportFormat, true
}

// SetReportFormat sets field value
func (o *InvoiceReportRequest) SetReportFormat(v string) {
	o.ReportFormat = v
}

// GetReportType returns the ReportType field value
func (o *InvoiceReportRequest) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *InvoiceReportRequest) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *InvoiceReportRequest) SetReportType(v string) {
	o.ReportType = v
}
