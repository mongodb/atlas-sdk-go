// Code based on the AtlasAPI V2 OpenAPI file

package admin

// SqlInterfaceStatusUpdateRequest struct for SqlInterfaceStatusUpdateRequest
type SqlInterfaceStatusUpdateRequest struct {
	// Current or desired SQL Interface status for the cluster. The GET endpoint returns OFF when no status is configured.
	Status string `json:"status"`
}

// NewSqlInterfaceStatusUpdateRequest instantiates a new SqlInterfaceStatusUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlInterfaceStatusUpdateRequest(status string) *SqlInterfaceStatusUpdateRequest {
	this := SqlInterfaceStatusUpdateRequest{}
	this.Status = status
	return &this
}

// NewSqlInterfaceStatusUpdateRequestWithDefaults instantiates a new SqlInterfaceStatusUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlInterfaceStatusUpdateRequestWithDefaults() *SqlInterfaceStatusUpdateRequest {
	this := SqlInterfaceStatusUpdateRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *SqlInterfaceStatusUpdateRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SqlInterfaceStatusUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SqlInterfaceStatusUpdateRequest) SetStatus(v string) {
	o.Status = v
}
