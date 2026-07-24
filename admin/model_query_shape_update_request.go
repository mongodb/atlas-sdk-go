// Code based on the AtlasAPI V2 OpenAPI file

package admin

// QueryShapeUpdateRequest Request body for modifying the rejection status of a query shape.
type QueryShapeUpdateRequest struct {
	// The rejection status of a query shape. Use REJECTED to prevent the query shape from executing on the cluster, or UNREJECTED to allow it to execute.
	Status string `json:"status"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *QueryShapeUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod QueryShapeUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewQueryShapeUpdateRequest instantiates a new QueryShapeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryShapeUpdateRequest(status string) *QueryShapeUpdateRequest {
	this := QueryShapeUpdateRequest{}
	this.Status = status
	return &this
}

// NewQueryShapeUpdateRequestWithDefaults instantiates a new QueryShapeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryShapeUpdateRequestWithDefaults() *QueryShapeUpdateRequest {
	this := QueryShapeUpdateRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *QueryShapeUpdateRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *QueryShapeUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *QueryShapeUpdateRequest) SetStatus(v string) {
	o.Status = v
}
