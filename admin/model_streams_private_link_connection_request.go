// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsPrivateLinkConnectionRequest Request body for updating a Private Link connection.
type StreamsPrivateLinkConnectionRequest struct {
	// The domain hostname for the AWS Confluent Serverless Private Link connection. Allowed only when no domain is currently set, or when the connection is in `IDLE` state.
	DnsDomain *string `json:"dnsDomain,omitempty"`
}

// NewStreamsPrivateLinkConnectionRequest instantiates a new StreamsPrivateLinkConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsPrivateLinkConnectionRequest() *StreamsPrivateLinkConnectionRequest {
	this := StreamsPrivateLinkConnectionRequest{}
	return &this
}

// NewStreamsPrivateLinkConnectionRequestWithDefaults instantiates a new StreamsPrivateLinkConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsPrivateLinkConnectionRequestWithDefaults() *StreamsPrivateLinkConnectionRequest {
	this := StreamsPrivateLinkConnectionRequest{}
	return &this
}

// GetDnsDomain returns the DnsDomain field value if set, zero value otherwise
func (o *StreamsPrivateLinkConnectionRequest) GetDnsDomain() string {
	if o == nil || IsNil(o.DnsDomain) {
		var ret string
		return ret
	}
	return *o.DnsDomain
}

// GetDnsDomainOk returns a tuple with the DnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsPrivateLinkConnectionRequest) GetDnsDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DnsDomain) {
		return nil, false
	}

	return o.DnsDomain, true
}

// HasDnsDomain returns a boolean if a field has been set.
func (o *StreamsPrivateLinkConnectionRequest) HasDnsDomain() bool {
	if o != nil && !IsNil(o.DnsDomain) {
		return true
	}

	return false
}

// SetDnsDomain gets a reference to the given string and assigns it to the DnsDomain field.
func (o *StreamsPrivateLinkConnectionRequest) SetDnsDomain(v string) {
	o.DnsDomain = &v
}
