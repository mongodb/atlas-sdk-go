// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiStandbyLinkRequest Request body for creating a standby link between clusters.
type ApiStandbyLinkRequest struct {
	// Unique 24-hexadecimal digit string that identifies the active cluster in the standby link.
	ActiveClusterUniqueId string `json:"activeClusterUniqueId"`
	// Unique 24-hexadecimal digit string that identifies the standby cluster in the standby link.
	StandbyClusterUniqueId string `json:"standbyClusterUniqueId"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiStandbyLinkRequest) MarshalJSON() ([]byte, error) {
	type noMethod ApiStandbyLinkRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiStandbyLinkRequest instantiates a new ApiStandbyLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStandbyLinkRequest(activeClusterUniqueId string, standbyClusterUniqueId string) *ApiStandbyLinkRequest {
	this := ApiStandbyLinkRequest{}
	this.ActiveClusterUniqueId = activeClusterUniqueId
	this.StandbyClusterUniqueId = standbyClusterUniqueId
	return &this
}

// NewApiStandbyLinkRequestWithDefaults instantiates a new ApiStandbyLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStandbyLinkRequestWithDefaults() *ApiStandbyLinkRequest {
	this := ApiStandbyLinkRequest{}
	return &this
}

// GetActiveClusterUniqueId returns the ActiveClusterUniqueId field value
func (o *ApiStandbyLinkRequest) GetActiveClusterUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveClusterUniqueId
}

// GetActiveClusterUniqueIdOk returns a tuple with the ActiveClusterUniqueId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkRequest) GetActiveClusterUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveClusterUniqueId, true
}

// SetActiveClusterUniqueId sets field value
func (o *ApiStandbyLinkRequest) SetActiveClusterUniqueId(v string) {
	o.ActiveClusterUniqueId = v
}

// GetStandbyClusterUniqueId returns the StandbyClusterUniqueId field value
func (o *ApiStandbyLinkRequest) GetStandbyClusterUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StandbyClusterUniqueId
}

// GetStandbyClusterUniqueIdOk returns a tuple with the StandbyClusterUniqueId field value
// and a boolean to check if the value has been set.
func (o *ApiStandbyLinkRequest) GetStandbyClusterUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandbyClusterUniqueId, true
}

// SetStandbyClusterUniqueId sets field value
func (o *ApiStandbyLinkRequest) SetStandbyClusterUniqueId(v string) {
	o.StandbyClusterUniqueId = v
}
