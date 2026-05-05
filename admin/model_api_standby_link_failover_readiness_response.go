// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiStandbyLinkFailoverReadinessResponse Response body containing a standby link failover readiness snapshot for assessing potential data loss before an unplanned failover.
type ApiStandbyLinkFailoverReadinessResponse struct {
	// Best-effort estimate of potential data loss in seconds if failover is triggered now.
	// Read only field.
	EstimatedDataLossSeconds float64 `json:"estimatedDataLossSeconds"`
	// Best-effort estimate of total expected recovery time in seconds, including Atlas-side failover overhead.
	// Read only field.
	EstimatedFailoverTimeSeconds float64 `json:"estimatedFailoverTimeSeconds"`
	// Unique 24-hexadecimal digit string that identifies the project that owns the standby link.
	// Read only field.
	GroupId string `json:"groupId"`
	// Unique 24-hexadecimal digit string that identifies the standby link described by this readiness snapshot.
	// Read only field.
	StandbyLinkId string `json:"standbyLinkId"`
}

// NewApiStandbyLinkFailoverReadinessResponse instantiates a new ApiStandbyLinkFailoverReadinessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStandbyLinkFailoverReadinessResponse(estimatedDataLossSeconds float64, estimatedFailoverTimeSeconds float64, groupId string, standbyLinkId string) *ApiStandbyLinkFailoverReadinessResponse {
	this := ApiStandbyLinkFailoverReadinessResponse{}
	this.EstimatedDataLossSeconds = estimatedDataLossSeconds
	this.EstimatedFailoverTimeSeconds = estimatedFailoverTimeSeconds
	this.GroupId = groupId
	this.StandbyLinkId = standbyLinkId
	return &this
}

// NewApiStandbyLinkFailoverReadinessResponseWithDefaults instantiates a new ApiStandbyLinkFailoverReadinessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStandbyLinkFailoverReadinessResponseWithDefaults() *ApiStandbyLinkFailoverReadinessResponse {
	this := ApiStandbyLinkFailoverReadinessResponse{}
	return &this
}

// GetEstimatedDataLossSeconds returns the EstimatedDataLossSeconds field value
func (o *ApiStandbyLinkFailoverReadinessResponse) GetEstimatedDataLossSeconds() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.EstimatedDataLossSeconds
}

// GetEstimatedDataLossSecondsOk returns a tuple with the EstimatedDataLossSeconds field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverReadinessResponse) GetEstimatedDataLossSecondsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedDataLossSeconds, true
}

// SetEstimatedDataLossSeconds sets field value
func (o *ApiStandbyLinkFailoverReadinessResponse) SetEstimatedDataLossSeconds(v float64) {
	o.EstimatedDataLossSeconds = v
}

// GetEstimatedFailoverTimeSeconds returns the EstimatedFailoverTimeSeconds field value
func (o *ApiStandbyLinkFailoverReadinessResponse) GetEstimatedFailoverTimeSeconds() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.EstimatedFailoverTimeSeconds
}

// GetEstimatedFailoverTimeSecondsOk returns a tuple with the EstimatedFailoverTimeSeconds field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverReadinessResponse) GetEstimatedFailoverTimeSecondsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedFailoverTimeSeconds, true
}

// SetEstimatedFailoverTimeSeconds sets field value
func (o *ApiStandbyLinkFailoverReadinessResponse) SetEstimatedFailoverTimeSeconds(v float64) {
	o.EstimatedFailoverTimeSeconds = v
}

// GetGroupId returns the GroupId field value
func (o *ApiStandbyLinkFailoverReadinessResponse) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverReadinessResponse) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ApiStandbyLinkFailoverReadinessResponse) SetGroupId(v string) {
	o.GroupId = v
}

// GetStandbyLinkId returns the StandbyLinkId field value
func (o *ApiStandbyLinkFailoverReadinessResponse) GetStandbyLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StandbyLinkId
}

// GetStandbyLinkIdOk returns a tuple with the StandbyLinkId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkFailoverReadinessResponse) GetStandbyLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandbyLinkId, true
}

// SetStandbyLinkId sets field value
func (o *ApiStandbyLinkFailoverReadinessResponse) SetStandbyLinkId(v string) {
	o.StandbyLinkId = v
}
