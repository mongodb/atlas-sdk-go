# GroupSummaryStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSizeBackedUpBytes** | Pointer to **int64** |  | [optional] 
**DataSizeNoLocalBytes** | Pointer to **int64** |  | [optional] 
**LastInvoiceAmountCents** | Pointer to **int64** |  | [optional] 
**NumBillableServer** | Pointer to **int** |  | [optional] 
**NumCluster** | Pointer to **int** |  | [optional] 
**NumPaidAtlasClusters** | Pointer to **int** |  | [optional] 
**NumTotalAtlasClusters** | Pointer to **int** |  | [optional] 
**SupportLevelLabel** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupSummaryStatistics

`func NewGroupSummaryStatistics() *GroupSummaryStatistics`

NewGroupSummaryStatistics instantiates a new GroupSummaryStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSummaryStatisticsWithDefaults

`func NewGroupSummaryStatisticsWithDefaults() *GroupSummaryStatistics`

NewGroupSummaryStatisticsWithDefaults instantiates a new GroupSummaryStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSizeBackedUpBytes

`func (o *GroupSummaryStatistics) GetDataSizeBackedUpBytes() int64`

GetDataSizeBackedUpBytes returns the DataSizeBackedUpBytes field if non-nil, zero value otherwise.

### GetDataSizeBackedUpBytesOk

`func (o *GroupSummaryStatistics) GetDataSizeBackedUpBytesOk() (*int64, bool)`

GetDataSizeBackedUpBytesOk returns a tuple with the DataSizeBackedUpBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeBackedUpBytes

`func (o *GroupSummaryStatistics) SetDataSizeBackedUpBytes(v int64)`

SetDataSizeBackedUpBytes sets DataSizeBackedUpBytes field to given value.

### HasDataSizeBackedUpBytes

`func (o *GroupSummaryStatistics) HasDataSizeBackedUpBytes() bool`

HasDataSizeBackedUpBytes returns a boolean if a field has been set.
### GetDataSizeNoLocalBytes

`func (o *GroupSummaryStatistics) GetDataSizeNoLocalBytes() int64`

GetDataSizeNoLocalBytes returns the DataSizeNoLocalBytes field if non-nil, zero value otherwise.

### GetDataSizeNoLocalBytesOk

`func (o *GroupSummaryStatistics) GetDataSizeNoLocalBytesOk() (*int64, bool)`

GetDataSizeNoLocalBytesOk returns a tuple with the DataSizeNoLocalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSizeNoLocalBytes

`func (o *GroupSummaryStatistics) SetDataSizeNoLocalBytes(v int64)`

SetDataSizeNoLocalBytes sets DataSizeNoLocalBytes field to given value.

### HasDataSizeNoLocalBytes

`func (o *GroupSummaryStatistics) HasDataSizeNoLocalBytes() bool`

HasDataSizeNoLocalBytes returns a boolean if a field has been set.
### GetLastInvoiceAmountCents

`func (o *GroupSummaryStatistics) GetLastInvoiceAmountCents() int64`

GetLastInvoiceAmountCents returns the LastInvoiceAmountCents field if non-nil, zero value otherwise.

### GetLastInvoiceAmountCentsOk

`func (o *GroupSummaryStatistics) GetLastInvoiceAmountCentsOk() (*int64, bool)`

GetLastInvoiceAmountCentsOk returns a tuple with the LastInvoiceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInvoiceAmountCents

`func (o *GroupSummaryStatistics) SetLastInvoiceAmountCents(v int64)`

SetLastInvoiceAmountCents sets LastInvoiceAmountCents field to given value.

### HasLastInvoiceAmountCents

`func (o *GroupSummaryStatistics) HasLastInvoiceAmountCents() bool`

HasLastInvoiceAmountCents returns a boolean if a field has been set.
### GetNumBillableServer

`func (o *GroupSummaryStatistics) GetNumBillableServer() int`

GetNumBillableServer returns the NumBillableServer field if non-nil, zero value otherwise.

### GetNumBillableServerOk

`func (o *GroupSummaryStatistics) GetNumBillableServerOk() (*int, bool)`

GetNumBillableServerOk returns a tuple with the NumBillableServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBillableServer

`func (o *GroupSummaryStatistics) SetNumBillableServer(v int)`

SetNumBillableServer sets NumBillableServer field to given value.

### HasNumBillableServer

`func (o *GroupSummaryStatistics) HasNumBillableServer() bool`

HasNumBillableServer returns a boolean if a field has been set.
### GetNumCluster

`func (o *GroupSummaryStatistics) GetNumCluster() int`

GetNumCluster returns the NumCluster field if non-nil, zero value otherwise.

### GetNumClusterOk

`func (o *GroupSummaryStatistics) GetNumClusterOk() (*int, bool)`

GetNumClusterOk returns a tuple with the NumCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCluster

`func (o *GroupSummaryStatistics) SetNumCluster(v int)`

SetNumCluster sets NumCluster field to given value.

### HasNumCluster

`func (o *GroupSummaryStatistics) HasNumCluster() bool`

HasNumCluster returns a boolean if a field has been set.
### GetNumPaidAtlasClusters

`func (o *GroupSummaryStatistics) GetNumPaidAtlasClusters() int`

GetNumPaidAtlasClusters returns the NumPaidAtlasClusters field if non-nil, zero value otherwise.

### GetNumPaidAtlasClustersOk

`func (o *GroupSummaryStatistics) GetNumPaidAtlasClustersOk() (*int, bool)`

GetNumPaidAtlasClustersOk returns a tuple with the NumPaidAtlasClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPaidAtlasClusters

`func (o *GroupSummaryStatistics) SetNumPaidAtlasClusters(v int)`

SetNumPaidAtlasClusters sets NumPaidAtlasClusters field to given value.

### HasNumPaidAtlasClusters

`func (o *GroupSummaryStatistics) HasNumPaidAtlasClusters() bool`

HasNumPaidAtlasClusters returns a boolean if a field has been set.
### GetNumTotalAtlasClusters

`func (o *GroupSummaryStatistics) GetNumTotalAtlasClusters() int`

GetNumTotalAtlasClusters returns the NumTotalAtlasClusters field if non-nil, zero value otherwise.

### GetNumTotalAtlasClustersOk

`func (o *GroupSummaryStatistics) GetNumTotalAtlasClustersOk() (*int, bool)`

GetNumTotalAtlasClustersOk returns a tuple with the NumTotalAtlasClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTotalAtlasClusters

`func (o *GroupSummaryStatistics) SetNumTotalAtlasClusters(v int)`

SetNumTotalAtlasClusters sets NumTotalAtlasClusters field to given value.

### HasNumTotalAtlasClusters

`func (o *GroupSummaryStatistics) HasNumTotalAtlasClusters() bool`

HasNumTotalAtlasClusters returns a boolean if a field has been set.
### GetSupportLevelLabel

`func (o *GroupSummaryStatistics) GetSupportLevelLabel() string`

GetSupportLevelLabel returns the SupportLevelLabel field if non-nil, zero value otherwise.

### GetSupportLevelLabelOk

`func (o *GroupSummaryStatistics) GetSupportLevelLabelOk() (*string, bool)`

GetSupportLevelLabelOk returns a tuple with the SupportLevelLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLevelLabel

`func (o *GroupSummaryStatistics) SetSupportLevelLabel(v string)`

SetSupportLevelLabel sets SupportLevelLabel field to given value.

### HasSupportLevelLabel

`func (o *GroupSummaryStatistics) HasSupportLevelLabel() bool`

HasSupportLevelLabel returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


