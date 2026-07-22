// Code based on the AtlasAPI V2 OpenAPI file

package admin

// EARPrivateEndpoint Encryption At Rest Private Endpoint.
type EARPrivateEndpoint struct {
	// Human-readable label that identifies the cloud provider for the Encryption At Rest private endpoint.
	// Read only field.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Error message for failures associated with the Encryption At Rest private endpoint.
	// Read only field.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Private Endpoint Service.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Cloud provider region in which the Encryption At Rest private endpoint is located.
	RegionName *string `json:"regionName,omitempty"`
	// State of the Encryption At Rest private endpoint.
	// Read only field.
	Status *string `json:"status,omitempty"`
	// Resource Id of the Aws Private Endpoint.
	// Read only field.
	PrivateEndpointConnectionName *string `json:"privateEndpointConnectionName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *EARPrivateEndpoint) MarshalJSON() ([]byte, error) {
	type noMethod EARPrivateEndpoint
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewEARPrivateEndpoint instantiates a new EARPrivateEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEARPrivateEndpoint() *EARPrivateEndpoint {
	this := EARPrivateEndpoint{}
	return &this
}

// NewEARPrivateEndpointWithDefaults instantiates a new EARPrivateEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEARPrivateEndpointWithDefaults() *EARPrivateEndpoint {
	this := EARPrivateEndpoint{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *EARPrivateEndpoint) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EARPrivateEndpoint) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *EARPrivateEndpoint) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *EARPrivateEndpoint) SetCloudProvider(v string) {
	o.CloudProvider = &v
	o.NullFields = removeNullField(o.NullFields, "CloudProvider")
}

// SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled.
func (o *EARPrivateEndpoint) SetCloudProviderNil() {
	o.CloudProvider = nil
	o.NullFields = addNullField(o.NullFields, "CloudProvider")
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *EARPrivateEndpoint) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EARPrivateEndpoint) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *EARPrivateEndpoint) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *EARPrivateEndpoint) SetErrorMessage(v string) {
	o.ErrorMessage = &v
	o.NullFields = removeNullField(o.NullFields, "ErrorMessage")
}

// SetErrorMessageNil sets ErrorMessage to an explicit JSON null when marshaled.
func (o *EARPrivateEndpoint) SetErrorMessageNil() {
	o.ErrorMessage = nil
	o.NullFields = addNullField(o.NullFields, "ErrorMessage")
}

// GetId returns the Id field value if set, zero value otherwise
func (o *EARPrivateEndpoint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EARPrivateEndpoint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EARPrivateEndpoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EARPrivateEndpoint) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *EARPrivateEndpoint) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *EARPrivateEndpoint) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EARPrivateEndpoint) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *EARPrivateEndpoint) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *EARPrivateEndpoint) SetRegionName(v string) {
	o.RegionName = &v
	o.NullFields = removeNullField(o.NullFields, "RegionName")
}

// SetRegionNameNil sets RegionName to an explicit JSON null when marshaled.
func (o *EARPrivateEndpoint) SetRegionNameNil() {
	o.RegionName = nil
	o.NullFields = addNullField(o.NullFields, "RegionName")
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *EARPrivateEndpoint) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EARPrivateEndpoint) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EARPrivateEndpoint) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EARPrivateEndpoint) SetStatus(v string) {
	o.Status = &v
	o.NullFields = removeNullField(o.NullFields, "Status")
}

// SetStatusNil sets Status to an explicit JSON null when marshaled.
func (o *EARPrivateEndpoint) SetStatusNil() {
	o.Status = nil
	o.NullFields = addNullField(o.NullFields, "Status")
}

// GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field value if set, zero value otherwise
func (o *EARPrivateEndpoint) GetPrivateEndpointConnectionName() string {
	if o == nil || IsNil(o.PrivateEndpointConnectionName) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointConnectionName
}

// GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EARPrivateEndpoint) GetPrivateEndpointConnectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointConnectionName) {
		return nil, false
	}

	return o.PrivateEndpointConnectionName, true
}

// HasPrivateEndpointConnectionName returns a boolean if a field has been set.
func (o *EARPrivateEndpoint) HasPrivateEndpointConnectionName() bool {
	if o != nil && !IsNil(o.PrivateEndpointConnectionName) {
		return true
	}

	return false
}

// SetPrivateEndpointConnectionName gets a reference to the given string and assigns it to the PrivateEndpointConnectionName field.
func (o *EARPrivateEndpoint) SetPrivateEndpointConnectionName(v string) {
	o.PrivateEndpointConnectionName = &v
	o.NullFields = removeNullField(o.NullFields, "PrivateEndpointConnectionName")
}

// SetPrivateEndpointConnectionNameNil sets PrivateEndpointConnectionName to an explicit JSON null when marshaled.
func (o *EARPrivateEndpoint) SetPrivateEndpointConnectionNameNil() {
	o.PrivateEndpointConnectionName = nil
	o.NullFields = addNullField(o.NullFields, "PrivateEndpointConnectionName")
}
