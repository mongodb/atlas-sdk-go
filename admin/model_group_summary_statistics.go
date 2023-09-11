// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GroupSummaryStatistics struct for GroupSummaryStatistics
type GroupSummaryStatistics struct {
	DataSizeBackedUpBytes  *int64  `json:"dataSizeBackedUpBytes,omitempty"`
	DataSizeNoLocalBytes   *int64  `json:"dataSizeNoLocalBytes,omitempty"`
	LastInvoiceAmountCents *int64  `json:"lastInvoiceAmountCents,omitempty"`
	NumBillableServer      *int    `json:"numBillableServer,omitempty"`
	NumCluster             *int    `json:"numCluster,omitempty"`
	NumPaidAtlasClusters   *int    `json:"numPaidAtlasClusters,omitempty"`
	NumTotalAtlasClusters  *int    `json:"numTotalAtlasClusters,omitempty"`
	SupportLevelLabel      *string `json:"supportLevelLabel,omitempty"`
}

// NewGroupSummaryStatistics instantiates a new GroupSummaryStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupSummaryStatistics() *GroupSummaryStatistics {
	this := GroupSummaryStatistics{}
	return &this
}

// NewGroupSummaryStatisticsWithDefaults instantiates a new GroupSummaryStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupSummaryStatisticsWithDefaults() *GroupSummaryStatistics {
	this := GroupSummaryStatistics{}
	return &this
}

// GetDataSizeBackedUpBytes returns the DataSizeBackedUpBytes field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetDataSizeBackedUpBytes() int64 {
	if o == nil || IsNil(o.DataSizeBackedUpBytes) {
		var ret int64
		return ret
	}
	return *o.DataSizeBackedUpBytes
}

// GetDataSizeBackedUpBytesOk returns a tuple with the DataSizeBackedUpBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetDataSizeBackedUpBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DataSizeBackedUpBytes) {
		return nil, false
	}

	return o.DataSizeBackedUpBytes, true
}

// HasDataSizeBackedUpBytes returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasDataSizeBackedUpBytes() bool {
	if o != nil && !IsNil(o.DataSizeBackedUpBytes) {
		return true
	}

	return false
}

// SetDataSizeBackedUpBytes gets a reference to the given int64 and assigns it to the DataSizeBackedUpBytes field.
func (o *GroupSummaryStatistics) SetDataSizeBackedUpBytes(v int64) {
	o.DataSizeBackedUpBytes = &v
}

// GetDataSizeNoLocalBytes returns the DataSizeNoLocalBytes field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetDataSizeNoLocalBytes() int64 {
	if o == nil || IsNil(o.DataSizeNoLocalBytes) {
		var ret int64
		return ret
	}
	return *o.DataSizeNoLocalBytes
}

// GetDataSizeNoLocalBytesOk returns a tuple with the DataSizeNoLocalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetDataSizeNoLocalBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.DataSizeNoLocalBytes) {
		return nil, false
	}

	return o.DataSizeNoLocalBytes, true
}

// HasDataSizeNoLocalBytes returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasDataSizeNoLocalBytes() bool {
	if o != nil && !IsNil(o.DataSizeNoLocalBytes) {
		return true
	}

	return false
}

// SetDataSizeNoLocalBytes gets a reference to the given int64 and assigns it to the DataSizeNoLocalBytes field.
func (o *GroupSummaryStatistics) SetDataSizeNoLocalBytes(v int64) {
	o.DataSizeNoLocalBytes = &v
}

// GetLastInvoiceAmountCents returns the LastInvoiceAmountCents field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetLastInvoiceAmountCents() int64 {
	if o == nil || IsNil(o.LastInvoiceAmountCents) {
		var ret int64
		return ret
	}
	return *o.LastInvoiceAmountCents
}

// GetLastInvoiceAmountCentsOk returns a tuple with the LastInvoiceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetLastInvoiceAmountCentsOk() (*int64, bool) {
	if o == nil || IsNil(o.LastInvoiceAmountCents) {
		return nil, false
	}

	return o.LastInvoiceAmountCents, true
}

// HasLastInvoiceAmountCents returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasLastInvoiceAmountCents() bool {
	if o != nil && !IsNil(o.LastInvoiceAmountCents) {
		return true
	}

	return false
}

// SetLastInvoiceAmountCents gets a reference to the given int64 and assigns it to the LastInvoiceAmountCents field.
func (o *GroupSummaryStatistics) SetLastInvoiceAmountCents(v int64) {
	o.LastInvoiceAmountCents = &v
}

// GetNumBillableServer returns the NumBillableServer field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetNumBillableServer() int {
	if o == nil || IsNil(o.NumBillableServer) {
		var ret int
		return ret
	}
	return *o.NumBillableServer
}

// GetNumBillableServerOk returns a tuple with the NumBillableServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetNumBillableServerOk() (*int, bool) {
	if o == nil || IsNil(o.NumBillableServer) {
		return nil, false
	}

	return o.NumBillableServer, true
}

// HasNumBillableServer returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasNumBillableServer() bool {
	if o != nil && !IsNil(o.NumBillableServer) {
		return true
	}

	return false
}

// SetNumBillableServer gets a reference to the given int and assigns it to the NumBillableServer field.
func (o *GroupSummaryStatistics) SetNumBillableServer(v int) {
	o.NumBillableServer = &v
}

// GetNumCluster returns the NumCluster field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetNumCluster() int {
	if o == nil || IsNil(o.NumCluster) {
		var ret int
		return ret
	}
	return *o.NumCluster
}

// GetNumClusterOk returns a tuple with the NumCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetNumClusterOk() (*int, bool) {
	if o == nil || IsNil(o.NumCluster) {
		return nil, false
	}

	return o.NumCluster, true
}

// HasNumCluster returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasNumCluster() bool {
	if o != nil && !IsNil(o.NumCluster) {
		return true
	}

	return false
}

// SetNumCluster gets a reference to the given int and assigns it to the NumCluster field.
func (o *GroupSummaryStatistics) SetNumCluster(v int) {
	o.NumCluster = &v
}

// GetNumPaidAtlasClusters returns the NumPaidAtlasClusters field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetNumPaidAtlasClusters() int {
	if o == nil || IsNil(o.NumPaidAtlasClusters) {
		var ret int
		return ret
	}
	return *o.NumPaidAtlasClusters
}

// GetNumPaidAtlasClustersOk returns a tuple with the NumPaidAtlasClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetNumPaidAtlasClustersOk() (*int, bool) {
	if o == nil || IsNil(o.NumPaidAtlasClusters) {
		return nil, false
	}

	return o.NumPaidAtlasClusters, true
}

// HasNumPaidAtlasClusters returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasNumPaidAtlasClusters() bool {
	if o != nil && !IsNil(o.NumPaidAtlasClusters) {
		return true
	}

	return false
}

// SetNumPaidAtlasClusters gets a reference to the given int and assigns it to the NumPaidAtlasClusters field.
func (o *GroupSummaryStatistics) SetNumPaidAtlasClusters(v int) {
	o.NumPaidAtlasClusters = &v
}

// GetNumTotalAtlasClusters returns the NumTotalAtlasClusters field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetNumTotalAtlasClusters() int {
	if o == nil || IsNil(o.NumTotalAtlasClusters) {
		var ret int
		return ret
	}
	return *o.NumTotalAtlasClusters
}

// GetNumTotalAtlasClustersOk returns a tuple with the NumTotalAtlasClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetNumTotalAtlasClustersOk() (*int, bool) {
	if o == nil || IsNil(o.NumTotalAtlasClusters) {
		return nil, false
	}

	return o.NumTotalAtlasClusters, true
}

// HasNumTotalAtlasClusters returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasNumTotalAtlasClusters() bool {
	if o != nil && !IsNil(o.NumTotalAtlasClusters) {
		return true
	}

	return false
}

// SetNumTotalAtlasClusters gets a reference to the given int and assigns it to the NumTotalAtlasClusters field.
func (o *GroupSummaryStatistics) SetNumTotalAtlasClusters(v int) {
	o.NumTotalAtlasClusters = &v
}

// GetSupportLevelLabel returns the SupportLevelLabel field value if set, zero value otherwise
func (o *GroupSummaryStatistics) GetSupportLevelLabel() string {
	if o == nil || IsNil(o.SupportLevelLabel) {
		var ret string
		return ret
	}
	return *o.SupportLevelLabel
}

// GetSupportLevelLabelOk returns a tuple with the SupportLevelLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupSummaryStatistics) GetSupportLevelLabelOk() (*string, bool) {
	if o == nil || IsNil(o.SupportLevelLabel) {
		return nil, false
	}

	return o.SupportLevelLabel, true
}

// HasSupportLevelLabel returns a boolean if a field has been set.
func (o *GroupSummaryStatistics) HasSupportLevelLabel() bool {
	if o != nil && !IsNil(o.SupportLevelLabel) {
		return true
	}

	return false
}

// SetSupportLevelLabel gets a reference to the given string and assigns it to the SupportLevelLabel field.
func (o *GroupSummaryStatistics) SetSupportLevelLabel(v string) {
	o.SupportLevelLabel = &v
}

func (o GroupSummaryStatistics) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupSummaryStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataSizeBackedUpBytes) {
		toSerialize["dataSizeBackedUpBytes"] = o.DataSizeBackedUpBytes
	}
	if !IsNil(o.DataSizeNoLocalBytes) {
		toSerialize["dataSizeNoLocalBytes"] = o.DataSizeNoLocalBytes
	}
	if !IsNil(o.LastInvoiceAmountCents) {
		toSerialize["lastInvoiceAmountCents"] = o.LastInvoiceAmountCents
	}
	if !IsNil(o.NumBillableServer) {
		toSerialize["numBillableServer"] = o.NumBillableServer
	}
	if !IsNil(o.NumCluster) {
		toSerialize["numCluster"] = o.NumCluster
	}
	if !IsNil(o.NumPaidAtlasClusters) {
		toSerialize["numPaidAtlasClusters"] = o.NumPaidAtlasClusters
	}
	if !IsNil(o.NumTotalAtlasClusters) {
		toSerialize["numTotalAtlasClusters"] = o.NumTotalAtlasClusters
	}
	if !IsNil(o.SupportLevelLabel) {
		toSerialize["supportLevelLabel"] = o.SupportLevelLabel
	}
	return toSerialize, nil
}
