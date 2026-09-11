// Code based on the AtlasAPI V2 OpenAPI file

package admin

// QueryStatsSummaryListResponse struct for QueryStatsSummaryListResponse
type QueryStatsSummaryListResponse struct {
	QueryStatsCollection *QueryStatsCollectionResponse `json:"queryStatsCollection,omitempty"`
	// List of query shape statistic summaries from Query Shape Insights.
	Summaries *[]QueryStatsSummary `json:"summaries,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *QueryStatsSummaryListResponse) MarshalJSON() ([]byte, error) {
	type noMethod QueryStatsSummaryListResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewQueryStatsSummaryListResponse instantiates a new QueryStatsSummaryListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryStatsSummaryListResponse() *QueryStatsSummaryListResponse {
	this := QueryStatsSummaryListResponse{}
	return &this
}

// NewQueryStatsSummaryListResponseWithDefaults instantiates a new QueryStatsSummaryListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryStatsSummaryListResponseWithDefaults() *QueryStatsSummaryListResponse {
	this := QueryStatsSummaryListResponse{}
	return &this
}

// GetQueryStatsCollection returns the QueryStatsCollection field value if set, zero value otherwise
func (o *QueryStatsSummaryListResponse) GetQueryStatsCollection() QueryStatsCollectionResponse {
	if o == nil || IsNil(o.QueryStatsCollection) {
		var ret QueryStatsCollectionResponse
		return ret
	}
	return *o.QueryStatsCollection
}

// GetQueryStatsCollectionOk returns a tuple with the QueryStatsCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummaryListResponse) GetQueryStatsCollectionOk() (*QueryStatsCollectionResponse, bool) {
	if o == nil || IsNil(o.QueryStatsCollection) {
		return nil, false
	}

	return o.QueryStatsCollection, true
}

// HasQueryStatsCollection returns a boolean if a field has been set.
func (o *QueryStatsSummaryListResponse) HasQueryStatsCollection() bool {
	if o != nil && !IsNil(o.QueryStatsCollection) {
		return true
	}

	return false
}

// SetQueryStatsCollection gets a reference to the given QueryStatsCollectionResponse and assigns it to the QueryStatsCollection field.
func (o *QueryStatsSummaryListResponse) SetQueryStatsCollection(v QueryStatsCollectionResponse) {
	o.QueryStatsCollection = &v
	o.NullFields = removeNullField(o.NullFields, "QueryStatsCollection")
}

// SetQueryStatsCollectionNil sets QueryStatsCollection to an explicit JSON null when marshaled.
func (o *QueryStatsSummaryListResponse) SetQueryStatsCollectionNil() {
	o.QueryStatsCollection = nil
	o.NullFields = addNullField(o.NullFields, "QueryStatsCollection")
}

// GetSummaries returns the Summaries field value if set, zero value otherwise
func (o *QueryStatsSummaryListResponse) GetSummaries() []QueryStatsSummary {
	if o == nil || IsNil(o.Summaries) {
		var ret []QueryStatsSummary
		return ret
	}
	return *o.Summaries
}

// GetSummariesOk returns a tuple with the Summaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryStatsSummaryListResponse) GetSummariesOk() (*[]QueryStatsSummary, bool) {
	if o == nil || IsNil(o.Summaries) {
		return nil, false
	}

	return o.Summaries, true
}

// HasSummaries returns a boolean if a field has been set.
func (o *QueryStatsSummaryListResponse) HasSummaries() bool {
	if o != nil && !IsNil(o.Summaries) {
		return true
	}

	return false
}

// SetSummaries gets a reference to the given []QueryStatsSummary and assigns it to the Summaries field.
func (o *QueryStatsSummaryListResponse) SetSummaries(v []QueryStatsSummary) {
	o.Summaries = &v
	o.NullFields = removeNullField(o.NullFields, "Summaries")
}

// SetSummariesNil sets Summaries to an explicit JSON null when marshaled.
func (o *QueryStatsSummaryListResponse) SetSummariesNil() {
	o.Summaries = nil
	o.NullFields = addNullField(o.NullFields, "Summaries")
}
