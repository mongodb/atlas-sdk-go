// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ReportGenerationRequest struct for ReportGenerationRequest
type ReportGenerationRequest struct {
	// Format of the report.
	ReportFormat *string `json:"reportFormat,omitempty"`
	// Type of report to generate.
	ReportType *string `json:"reportType,omitempty"`
	// Version of the report format specification.
	Version *string `json:"version,omitempty"`
}

// NewReportGenerationRequest instantiates a new ReportGenerationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportGenerationRequest() *ReportGenerationRequest {
	this := ReportGenerationRequest{}
	return &this
}

// NewReportGenerationRequestWithDefaults instantiates a new ReportGenerationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportGenerationRequestWithDefaults() *ReportGenerationRequest {
	this := ReportGenerationRequest{}
	return &this
}

// GetReportFormat returns the ReportFormat field value if set, zero value otherwise
func (o *ReportGenerationRequest) GetReportFormat() string {
	if o == nil || IsNil(o.ReportFormat) {
		var ret string
		return ret
	}
	return *o.ReportFormat
}

// GetReportFormatOk returns a tuple with the ReportFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGenerationRequest) GetReportFormatOk() (*string, bool) {
	if o == nil || IsNil(o.ReportFormat) {
		return nil, false
	}

	return o.ReportFormat, true
}

// HasReportFormat returns a boolean if a field has been set.
func (o *ReportGenerationRequest) HasReportFormat() bool {
	if o != nil && !IsNil(o.ReportFormat) {
		return true
	}

	return false
}

// SetReportFormat gets a reference to the given string and assigns it to the ReportFormat field.
func (o *ReportGenerationRequest) SetReportFormat(v string) {
	o.ReportFormat = &v
}

// GetReportType returns the ReportType field value if set, zero value otherwise
func (o *ReportGenerationRequest) GetReportType() string {
	if o == nil || IsNil(o.ReportType) {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGenerationRequest) GetReportTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportType) {
		return nil, false
	}

	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportGenerationRequest) HasReportType() bool {
	if o != nil && !IsNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *ReportGenerationRequest) SetReportType(v string) {
	o.ReportType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise
func (o *ReportGenerationRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGenerationRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}

	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ReportGenerationRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ReportGenerationRequest) SetVersion(v string) {
	o.Version = &v
}
