# InvoiceReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormatSpecVersion** | Pointer to **string** | Version of the report format specification. | [optional] 
**ReportFormat** | **string** | Format of the report. | 
**ReportType** | **string** | Type of report to generate. | 

## Methods

### NewInvoiceReportRequest

`func NewInvoiceReportRequest(reportFormat string, reportType string, ) *InvoiceReportRequest`

NewInvoiceReportRequest instantiates a new InvoiceReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReportRequestWithDefaults

`func NewInvoiceReportRequestWithDefaults() *InvoiceReportRequest`

NewInvoiceReportRequestWithDefaults instantiates a new InvoiceReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormatSpecVersion

`func (o *InvoiceReportRequest) GetFormatSpecVersion() string`

GetFormatSpecVersion returns the FormatSpecVersion field if non-nil, zero value otherwise.

### GetFormatSpecVersionOk

`func (o *InvoiceReportRequest) GetFormatSpecVersionOk() (*string, bool)`

GetFormatSpecVersionOk returns a tuple with the FormatSpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatSpecVersion

`func (o *InvoiceReportRequest) SetFormatSpecVersion(v string)`

SetFormatSpecVersion sets FormatSpecVersion field to given value.

### HasFormatSpecVersion

`func (o *InvoiceReportRequest) HasFormatSpecVersion() bool`

HasFormatSpecVersion returns a boolean if a field has been set.

### SetFormatSpecVersionNil

`func (o *InvoiceReportRequest) SetFormatSpecVersionNil()`

SetFormatSpecVersionNil sets FormatSpecVersion to an explicit JSON null when marshaled, overriding any value previously set with SetFormatSpecVersion. Calling SetFormatSpecVersion again clears the null override.

### GetReportFormat

`func (o *InvoiceReportRequest) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *InvoiceReportRequest) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *InvoiceReportRequest) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### GetReportType

`func (o *InvoiceReportRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *InvoiceReportRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *InvoiceReportRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


