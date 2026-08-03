// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// InvoiceReportResponse Status and details of a previously requested invoice report.
type InvoiceReportResponse struct {
	// URL to download the report. Present only when the report has succeeded.
	// Read only field.
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// Time at which the download URL expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Present only when the report has succeeded.
	// Read only field.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Reason the report failed. Present only when the report has failed.
	// Read only field.
	FailureReason *string `json:"failureReason,omitempty"`
	// Version of the report format specification.
	FormatSpecVersion *string `json:"formatSpecVersion,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the invoice.
	// Read only field.
	InvoiceId string `json:"invoiceId"`
	// Format of the generated report.
	ReportFormat string `json:"reportFormat"`
	// Unique 24-hexadecimal digit string that identifies the report.
	// Read only field.
	ReportId string `json:"reportId"`
	// Type of the generated report.
	ReportType string `json:"reportType"`
	// Current state of the report generation.
	// Read only field.
	State string `json:"state"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *InvoiceReportResponse) MarshalJSON() ([]byte, error) {
	type noMethod InvoiceReportResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewInvoiceReportResponse instantiates a new InvoiceReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceReportResponse(invoiceId string, reportFormat string, reportId string, reportType string, state string) *InvoiceReportResponse {
	this := InvoiceReportResponse{}
	this.InvoiceId = invoiceId
	this.ReportFormat = reportFormat
	this.ReportId = reportId
	this.ReportType = reportType
	this.State = state
	return &this
}

// NewInvoiceReportResponseWithDefaults instantiates a new InvoiceReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceReportResponseWithDefaults() *InvoiceReportResponse {
	this := InvoiceReportResponse{}
	return &this
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise
func (o *InvoiceReportResponse) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}

	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *InvoiceReportResponse) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *InvoiceReportResponse) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
	o.NullFields = removeNullField(o.NullFields, "DownloadUrl")
}

// SetDownloadUrlNil sets DownloadUrl to an explicit JSON null when marshaled.
func (o *InvoiceReportResponse) SetDownloadUrlNil() {
	o.DownloadUrl = nil
	o.NullFields = addNullField(o.NullFields, "DownloadUrl")
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *InvoiceReportResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *InvoiceReportResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *InvoiceReportResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
	o.NullFields = removeNullField(o.NullFields, "ExpiresAt")
}

// SetExpiresAtNil sets ExpiresAt to an explicit JSON null when marshaled.
func (o *InvoiceReportResponse) SetExpiresAtNil() {
	o.ExpiresAt = nil
	o.NullFields = addNullField(o.NullFields, "ExpiresAt")
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise
func (o *InvoiceReportResponse) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}

	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *InvoiceReportResponse) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *InvoiceReportResponse) SetFailureReason(v string) {
	o.FailureReason = &v
	o.NullFields = removeNullField(o.NullFields, "FailureReason")
}

// SetFailureReasonNil sets FailureReason to an explicit JSON null when marshaled.
func (o *InvoiceReportResponse) SetFailureReasonNil() {
	o.FailureReason = nil
	o.NullFields = addNullField(o.NullFields, "FailureReason")
}

// GetFormatSpecVersion returns the FormatSpecVersion field value if set, zero value otherwise
func (o *InvoiceReportResponse) GetFormatSpecVersion() string {
	if o == nil || IsNil(o.FormatSpecVersion) {
		var ret string
		return ret
	}
	return *o.FormatSpecVersion
}

// GetFormatSpecVersionOk returns a tuple with the FormatSpecVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetFormatSpecVersionOk() (*string, bool) {
	if o == nil || IsNil(o.FormatSpecVersion) {
		return nil, false
	}

	return o.FormatSpecVersion, true
}

// HasFormatSpecVersion returns a boolean if a field has been set.
func (o *InvoiceReportResponse) HasFormatSpecVersion() bool {
	if o != nil && !IsNil(o.FormatSpecVersion) {
		return true
	}

	return false
}

// SetFormatSpecVersion gets a reference to the given string and assigns it to the FormatSpecVersion field.
func (o *InvoiceReportResponse) SetFormatSpecVersion(v string) {
	o.FormatSpecVersion = &v
	o.NullFields = removeNullField(o.NullFields, "FormatSpecVersion")
}

// SetFormatSpecVersionNil sets FormatSpecVersion to an explicit JSON null when marshaled.
func (o *InvoiceReportResponse) SetFormatSpecVersionNil() {
	o.FormatSpecVersion = nil
	o.NullFields = addNullField(o.NullFields, "FormatSpecVersion")
}

// GetInvoiceId returns the InvoiceId field value
func (o *InvoiceReportResponse) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *InvoiceReportResponse) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetReportFormat returns the ReportFormat field value
func (o *InvoiceReportResponse) GetReportFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportFormat
}

// GetReportFormatOk returns a tuple with the ReportFormat field value
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetReportFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportFormat, true
}

// SetReportFormat sets field value
func (o *InvoiceReportResponse) SetReportFormat(v string) {
	o.ReportFormat = v
}

// GetReportId returns the ReportId field value
func (o *InvoiceReportResponse) GetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportId, true
}

// SetReportId sets field value
func (o *InvoiceReportResponse) SetReportId(v string) {
	o.ReportId = v
}

// GetReportType returns the ReportType field value
func (o *InvoiceReportResponse) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *InvoiceReportResponse) SetReportType(v string) {
	o.ReportType = v
}

// GetState returns the State field value
func (o *InvoiceReportResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InvoiceReportResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InvoiceReportResponse) SetState(v string) {
	o.State = v
}
