// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgLogIntegrationResponse Response schema for organization-level log integration operations.
type OrgLogIntegrationResponse struct {
	// Unique 24-character hexadecimal digit string that identifies the log integration configuration.
	// Read only field.
	Id string `json:"id"`
	// Array of log types exported by this integration.
	LogTypes []string `json:"logTypes"`
	// Human-readable label that identifies the service to which you want to integrate with Atlas. The value must match the log integration type. This value cannot be modified after the integration is created.
	Type string `json:"type"`
	// OpenTelemetry collector endpoint URL.
	OtelEndpoint *string `json:"otelEndpoint,omitempty"`
	// HTTP headers for authentication and configuration. Maximum 10 headers, total size limit 2KB. Values are redacted.
	OtelSuppliedHeaders *[]OrgLogIntegrationHeader `json:"otelSuppliedHeaders,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OrgLogIntegrationResponse) MarshalJSON() ([]byte, error) {
	type noMethod OrgLogIntegrationResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOrgLogIntegrationResponse instantiates a new OrgLogIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgLogIntegrationResponse(id string, logTypes []string, type_ string) *OrgLogIntegrationResponse {
	this := OrgLogIntegrationResponse{}
	this.Id = id
	this.LogTypes = logTypes
	this.Type = type_
	return &this
}

// NewOrgLogIntegrationResponseWithDefaults instantiates a new OrgLogIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgLogIntegrationResponseWithDefaults() *OrgLogIntegrationResponse {
	this := OrgLogIntegrationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrgLogIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrgLogIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetLogTypes returns the LogTypes field value
func (o *OrgLogIntegrationResponse) GetLogTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationResponse) GetLogTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTypes, true
}

// SetLogTypes sets field value
func (o *OrgLogIntegrationResponse) SetLogTypes(v []string) {
	o.LogTypes = v
}

// GetType returns the Type field value
func (o *OrgLogIntegrationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrgLogIntegrationResponse) SetType(v string) {
	o.Type = v
}

// GetOtelEndpoint returns the OtelEndpoint field value if set, zero value otherwise
func (o *OrgLogIntegrationResponse) GetOtelEndpoint() string {
	if o == nil || IsNil(o.OtelEndpoint) {
		var ret string
		return ret
	}
	return *o.OtelEndpoint
}

// GetOtelEndpointOk returns a tuple with the OtelEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationResponse) GetOtelEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.OtelEndpoint) {
		return nil, false
	}

	return o.OtelEndpoint, true
}

// HasOtelEndpoint returns a boolean if a field has been set.
func (o *OrgLogIntegrationResponse) HasOtelEndpoint() bool {
	if o != nil && !IsNil(o.OtelEndpoint) {
		return true
	}

	return false
}

// SetOtelEndpoint gets a reference to the given string and assigns it to the OtelEndpoint field.
func (o *OrgLogIntegrationResponse) SetOtelEndpoint(v string) {
	o.OtelEndpoint = &v
	o.NullFields = removeNullField(o.NullFields, "OtelEndpoint")
}

// SetOtelEndpointNil sets OtelEndpoint to an explicit JSON null when marshaled.
func (o *OrgLogIntegrationResponse) SetOtelEndpointNil() {
	o.OtelEndpoint = nil
	o.NullFields = addNullField(o.NullFields, "OtelEndpoint")
}

// GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field value if set, zero value otherwise
func (o *OrgLogIntegrationResponse) GetOtelSuppliedHeaders() []OrgLogIntegrationHeader {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		var ret []OrgLogIntegrationHeader
		return ret
	}
	return *o.OtelSuppliedHeaders
}

// GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationResponse) GetOtelSuppliedHeadersOk() (*[]OrgLogIntegrationHeader, bool) {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		return nil, false
	}

	return o.OtelSuppliedHeaders, true
}

// HasOtelSuppliedHeaders returns a boolean if a field has been set.
func (o *OrgLogIntegrationResponse) HasOtelSuppliedHeaders() bool {
	if o != nil && !IsNil(o.OtelSuppliedHeaders) {
		return true
	}

	return false
}

// SetOtelSuppliedHeaders gets a reference to the given []OrgLogIntegrationHeader and assigns it to the OtelSuppliedHeaders field.
func (o *OrgLogIntegrationResponse) SetOtelSuppliedHeaders(v []OrgLogIntegrationHeader) {
	o.OtelSuppliedHeaders = &v
	o.NullFields = removeNullField(o.NullFields, "OtelSuppliedHeaders")
}

// SetOtelSuppliedHeadersNil sets OtelSuppliedHeaders to an explicit JSON null when marshaled.
func (o *OrgLogIntegrationResponse) SetOtelSuppliedHeadersNil() {
	o.OtelSuppliedHeaders = nil
	o.NullFields = addNullField(o.NullFields, "OtelSuppliedHeaders")
}
