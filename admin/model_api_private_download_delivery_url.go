// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiPrivateDownloadDeliveryUrl One Uniform Resource Locator (URL) that points to the compressed snapshot files for manual download and the corresponding private endpoint.
type ApiPrivateDownloadDeliveryUrl struct {
	// One Uniform Resource Locator that points to the compressed snapshot files for manual download.
	DeliveryUrl *string `json:"deliveryUrl,omitempty"`
	// Unique 22-character alphanumeric string that identifies the private endpoint.
	EndpointId *string `json:"endpointId,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ApiPrivateDownloadDeliveryUrl) MarshalJSON() ([]byte, error) {
	type noMethod ApiPrivateDownloadDeliveryUrl
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewApiPrivateDownloadDeliveryUrl instantiates a new ApiPrivateDownloadDeliveryUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPrivateDownloadDeliveryUrl() *ApiPrivateDownloadDeliveryUrl {
	this := ApiPrivateDownloadDeliveryUrl{}
	return &this
}

// NewApiPrivateDownloadDeliveryUrlWithDefaults instantiates a new ApiPrivateDownloadDeliveryUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPrivateDownloadDeliveryUrlWithDefaults() *ApiPrivateDownloadDeliveryUrl {
	this := ApiPrivateDownloadDeliveryUrl{}
	return &this
}

// GetDeliveryUrl returns the DeliveryUrl field value if set, zero value otherwise
func (o *ApiPrivateDownloadDeliveryUrl) GetDeliveryUrl() string {
	if o == nil || IsNil(o.DeliveryUrl) {
		var ret string
		return ret
	}
	return *o.DeliveryUrl
}

// GetDeliveryUrlOk returns a tuple with the DeliveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivateDownloadDeliveryUrl) GetDeliveryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryUrl) {
		return nil, false
	}

	return o.DeliveryUrl, true
}

// HasDeliveryUrl returns a boolean if a field has been set.
func (o *ApiPrivateDownloadDeliveryUrl) HasDeliveryUrl() bool {
	if o != nil && !IsNil(o.DeliveryUrl) {
		return true
	}

	return false
}

// SetDeliveryUrl gets a reference to the given string and assigns it to the DeliveryUrl field.
func (o *ApiPrivateDownloadDeliveryUrl) SetDeliveryUrl(v string) {
	o.DeliveryUrl = &v
}

// SetDeliveryUrlNil sets DeliveryUrl to an explicit JSON null when marshaled.
func (o *ApiPrivateDownloadDeliveryUrl) SetDeliveryUrlNil() {
	o.DeliveryUrl = nil
	o.NullFields = append(o.NullFields, "DeliveryUrl")
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise
func (o *ApiPrivateDownloadDeliveryUrl) GetEndpointId() string {
	if o == nil || IsNil(o.EndpointId) {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPrivateDownloadDeliveryUrl) GetEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointId) {
		return nil, false
	}

	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *ApiPrivateDownloadDeliveryUrl) HasEndpointId() bool {
	if o != nil && !IsNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *ApiPrivateDownloadDeliveryUrl) SetEndpointId(v string) {
	o.EndpointId = &v
}

// SetEndpointIdNil sets EndpointId to an explicit JSON null when marshaled.
func (o *ApiPrivateDownloadDeliveryUrl) SetEndpointIdNil() {
	o.EndpointId = nil
	o.NullFields = append(o.NullFields, "EndpointId")
}
