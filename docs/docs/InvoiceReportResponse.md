# InvoiceReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadUrl** | Pointer to **string** | URL to download the report. Present only when the report has succeeded. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Time at which the download URL expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC. Present only when the report has succeeded. | [optional] [readonly] 
**FailureReason** | Pointer to **string** | Reason the report failed. Present only when the report has failed. | [optional] [readonly] 
**FormatSpecVersion** | Pointer to **string** | Version of the report format specification. | [optional] 
**InvoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice. | [readonly] 
**ReportFormat** | **string** | Format of the generated report. | 
**ReportId** | **string** | Unique 24-hexadecimal digit string that identifies the report. | [readonly] 
**ReportType** | **string** | Type of the generated report. | 
**State** | **string** | Current state of the report generation. | [readonly] 

## Methods

### NewInvoiceReportResponse

`func NewInvoiceReportResponse(invoiceId string, reportFormat string, reportId string, reportType string, state string, ) *InvoiceReportResponse`

NewInvoiceReportResponse instantiates a new InvoiceReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReportResponseWithDefaults

`func NewInvoiceReportResponseWithDefaults() *InvoiceReportResponse`

NewInvoiceReportResponseWithDefaults instantiates a new InvoiceReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadUrl

`func (o *InvoiceReportResponse) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *InvoiceReportResponse) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *InvoiceReportResponse) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *InvoiceReportResponse) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### SetDownloadUrlNil

`func (o *InvoiceReportResponse) SetDownloadUrlNil()`

SetDownloadUrlNil sets DownloadUrl to an explicit JSON null when marshaled, overriding any value previously set with SetDownloadUrl. Calling SetDownloadUrl again clears the null override.

### GetExpiresAt

`func (o *InvoiceReportResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InvoiceReportResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InvoiceReportResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InvoiceReportResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *InvoiceReportResponse) SetExpiresAtNil()`

SetExpiresAtNil sets ExpiresAt to an explicit JSON null when marshaled, overriding any value previously set with SetExpiresAt. Calling SetExpiresAt again clears the null override.

### GetFailureReason

`func (o *InvoiceReportResponse) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *InvoiceReportResponse) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *InvoiceReportResponse) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *InvoiceReportResponse) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *InvoiceReportResponse) SetFailureReasonNil()`

SetFailureReasonNil sets FailureReason to an explicit JSON null when marshaled, overriding any value previously set with SetFailureReason. Calling SetFailureReason again clears the null override.

### GetFormatSpecVersion

`func (o *InvoiceReportResponse) GetFormatSpecVersion() string`

GetFormatSpecVersion returns the FormatSpecVersion field if non-nil, zero value otherwise.

### GetFormatSpecVersionOk

`func (o *InvoiceReportResponse) GetFormatSpecVersionOk() (*string, bool)`

GetFormatSpecVersionOk returns a tuple with the FormatSpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatSpecVersion

`func (o *InvoiceReportResponse) SetFormatSpecVersion(v string)`

SetFormatSpecVersion sets FormatSpecVersion field to given value.

### HasFormatSpecVersion

`func (o *InvoiceReportResponse) HasFormatSpecVersion() bool`

HasFormatSpecVersion returns a boolean if a field has been set.

### SetFormatSpecVersionNil

`func (o *InvoiceReportResponse) SetFormatSpecVersionNil()`

SetFormatSpecVersionNil sets FormatSpecVersion to an explicit JSON null when marshaled, overriding any value previously set with SetFormatSpecVersion. Calling SetFormatSpecVersion again clears the null override.

### GetInvoiceId

`func (o *InvoiceReportResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceReportResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceReportResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### GetReportFormat

`func (o *InvoiceReportResponse) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *InvoiceReportResponse) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *InvoiceReportResponse) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### GetReportId

`func (o *InvoiceReportResponse) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *InvoiceReportResponse) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *InvoiceReportResponse) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### GetReportType

`func (o *InvoiceReportResponse) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *InvoiceReportResponse) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *InvoiceReportResponse) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### GetState

`func (o *InvoiceReportResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InvoiceReportResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InvoiceReportResponse) SetState(v string)`

SetState sets State field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


