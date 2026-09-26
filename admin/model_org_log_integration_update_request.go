// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgLogIntegrationUpdateRequest Request schema for updating an organization-level log integration.
type OrgLogIntegrationUpdateRequest struct {
	// Array of log types exported by this integration.
	LogTypes []string `json:"logTypes"`
	// Human-readable label that identifies the service to which you want to integrate with Atlas. The value must match the log integration type. This value cannot be modified after the integration is created.
	Type string `json:"type"`
	// OpenTelemetry collector endpoint URL. Must be HTTPS and not exceed 2048 characters.
	OtelEndpoint *string `json:"otelEndpoint,omitempty"`
	// HTTP headers for authentication and configuration. Maximum 10 headers, total size limit 2KB.
	OtelSuppliedHeaders *[]OrgLogIntegrationHeader `json:"otelSuppliedHeaders,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OrgLogIntegrationUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod OrgLogIntegrationUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOrgLogIntegrationUpdateRequest instantiates a new OrgLogIntegrationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgLogIntegrationUpdateRequest(logTypes []string, type_ string) *OrgLogIntegrationUpdateRequest {
	this := OrgLogIntegrationUpdateRequest{}
	this.LogTypes = logTypes
	this.Type = type_
	return &this
}

// NewOrgLogIntegrationUpdateRequestWithDefaults instantiates a new OrgLogIntegrationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgLogIntegrationUpdateRequestWithDefaults() *OrgLogIntegrationUpdateRequest {
	this := OrgLogIntegrationUpdateRequest{}
	return &this
}

// GetLogTypes returns the LogTypes field value
func (o *OrgLogIntegrationUpdateRequest) GetLogTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationUpdateRequest) GetLogTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTypes, true
}

// SetLogTypes sets field value
func (o *OrgLogIntegrationUpdateRequest) SetLogTypes(v []string) {
	o.LogTypes = v
}

// GetType returns the Type field value
func (o *OrgLogIntegrationUpdateRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationUpdateRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrgLogIntegrationUpdateRequest) SetType(v string) {
	o.Type = v
}

// GetOtelEndpoint returns the OtelEndpoint field value if set, zero value otherwise
func (o *OrgLogIntegrationUpdateRequest) GetOtelEndpoint() string {
	if o == nil || IsNil(o.OtelEndpoint) {
		var ret string
		return ret
	}
	return *o.OtelEndpoint
}

// GetOtelEndpointOk returns a tuple with the OtelEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationUpdateRequest) GetOtelEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.OtelEndpoint) {
		return nil, false
	}

	return o.OtelEndpoint, true
}

// HasOtelEndpoint returns a boolean if a field has been set.
func (o *OrgLogIntegrationUpdateRequest) HasOtelEndpoint() bool {
	if o != nil && !IsNil(o.OtelEndpoint) {
		return true
	}

	return false
}

// SetOtelEndpoint gets a reference to the given string and assigns it to the OtelEndpoint field.
func (o *OrgLogIntegrationUpdateRequest) SetOtelEndpoint(v string) {
	o.OtelEndpoint = &v
	o.NullFields = removeNullField(o.NullFields, "OtelEndpoint")
}

// SetOtelEndpointNil sets OtelEndpoint to an explicit JSON null when marshaled.
func (o *OrgLogIntegrationUpdateRequest) SetOtelEndpointNil() {
	o.OtelEndpoint = nil
	o.NullFields = addNullField(o.NullFields, "OtelEndpoint")
}

// GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field value if set, zero value otherwise
func (o *OrgLogIntegrationUpdateRequest) GetOtelSuppliedHeaders() []OrgLogIntegrationHeader {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		var ret []OrgLogIntegrationHeader
		return ret
	}
	return *o.OtelSuppliedHeaders
}

// GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationUpdateRequest) GetOtelSuppliedHeadersOk() (*[]OrgLogIntegrationHeader, bool) {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		return nil, false
	}

	return o.OtelSuppliedHeaders, true
}

// HasOtelSuppliedHeaders returns a boolean if a field has been set.
func (o *OrgLogIntegrationUpdateRequest) HasOtelSuppliedHeaders() bool {
	if o != nil && !IsNil(o.OtelSuppliedHeaders) {
		return true
	}

	return false
}

// SetOtelSuppliedHeaders gets a reference to the given []OrgLogIntegrationHeader and assigns it to the OtelSuppliedHeaders field.
func (o *OrgLogIntegrationUpdateRequest) SetOtelSuppliedHeaders(v []OrgLogIntegrationHeader) {
	o.OtelSuppliedHeaders = &v
	o.NullFields = removeNullField(o.NullFields, "OtelSuppliedHeaders")
}

// SetOtelSuppliedHeadersNil sets OtelSuppliedHeaders to an explicit JSON null when marshaled.
func (o *OrgLogIntegrationUpdateRequest) SetOtelSuppliedHeadersNil() {
	o.OtelSuppliedHeaders = nil
	o.NullFields = addNullField(o.NullFields, "OtelSuppliedHeaders")
}
