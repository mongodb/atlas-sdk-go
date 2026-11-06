// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ShardingInsightsResponse Cluster-level sharding state of one sharded cluster: shard count, balancer state, config server availability, and the configured balancing window.
type ShardingInsightsResponse struct {
	// State of the cluster balancer. `BALANCING` means a balancer round is in progress. `IDLE` means the balancer is enabled and no round is in progress. `DISABLED` means the balancer is turned off. `FAILED_TO_DETERMINE` means the balancer state could not be resolved.
	// Read only field.
	BalancerState   string                          `json:"balancerState"`
	BalancingWindow ShardingInsightsBalancingWindow `json:"balancingWindow"`
	// Availability of the cluster's config servers. `AVAILABLE` means at least one config server reported recently. `UNAVAILABLE` means config servers exist but none has reported recently. `FAILED_TO_DETERMINE` means availability could not be resolved.
	// Read only field.
	ConfigServerStatus string `json:"configServerStatus"`
	// Number of shards in the cluster.
	// Read only field.
	TotalShards int `json:"totalShards"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ShardingInsightsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ShardingInsightsResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewShardingInsightsResponse instantiates a new ShardingInsightsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShardingInsightsResponse(balancerState string, balancingWindow ShardingInsightsBalancingWindow, configServerStatus string, totalShards int) *ShardingInsightsResponse {
	this := ShardingInsightsResponse{}
	this.BalancerState = balancerState
	this.BalancingWindow = balancingWindow
	this.ConfigServerStatus = configServerStatus
	this.TotalShards = totalShards
	return &this
}

// NewShardingInsightsResponseWithDefaults instantiates a new ShardingInsightsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShardingInsightsResponseWithDefaults() *ShardingInsightsResponse {
	this := ShardingInsightsResponse{}
	return &this
}

// GetBalancerState returns the BalancerState field value
func (o *ShardingInsightsResponse) GetBalancerState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalancerState
}

// GetBalancerStateOk returns a tuple with the BalancerState field value
// and a boolean to check if the value has been set.
func (o *ShardingInsightsResponse) GetBalancerStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalancerState, true
}

// SetBalancerState sets field value
func (o *ShardingInsightsResponse) SetBalancerState(v string) {
	o.BalancerState = v
}

// GetBalancingWindow returns the BalancingWindow field value
func (o *ShardingInsightsResponse) GetBalancingWindow() ShardingInsightsBalancingWindow {
	if o == nil {
		var ret ShardingInsightsBalancingWindow
		return ret
	}

	return o.BalancingWindow
}

// GetBalancingWindowOk returns a tuple with the BalancingWindow field value
// and a boolean to check if the value has been set.
func (o *ShardingInsightsResponse) GetBalancingWindowOk() (*ShardingInsightsBalancingWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalancingWindow, true
}

// SetBalancingWindow sets field value
func (o *ShardingInsightsResponse) SetBalancingWindow(v ShardingInsightsBalancingWindow) {
	o.BalancingWindow = v
}

// GetConfigServerStatus returns the ConfigServerStatus field value
func (o *ShardingInsightsResponse) GetConfigServerStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigServerStatus
}

// GetConfigServerStatusOk returns a tuple with the ConfigServerStatus field value
// and a boolean to check if the value has been set.
func (o *ShardingInsightsResponse) GetConfigServerStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigServerStatus, true
}

// SetConfigServerStatus sets field value
func (o *ShardingInsightsResponse) SetConfigServerStatus(v string) {
	o.ConfigServerStatus = v
}

// GetTotalShards returns the TotalShards field value
func (o *ShardingInsightsResponse) GetTotalShards() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TotalShards
}

// GetTotalShardsOk returns a tuple with the TotalShards field value
// and a boolean to check if the value has been set.
func (o *ShardingInsightsResponse) GetTotalShardsOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalShards, true
}

// SetTotalShards sets field value
func (o *ShardingInsightsResponse) SetTotalShards(v int) {
	o.TotalShards = v
}
