// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgLogIntegrationRequest Request schema for creating an organization-level log integration.
type OrgLogIntegrationRequest struct {
	// Requests a one-time backfill of historical events to this integration's sink.
	EventBackfill *EventBackfillRequest `json:"eventBackfill,omitempty"`
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

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o OrgLogIntegrationRequest) MarshalJSON() ([]byte, error) {
	type noMethod OrgLogIntegrationRequest
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewOrgLogIntegrationRequest instantiates a new OrgLogIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgLogIntegrationRequest(logTypes []string, type_ string) *OrgLogIntegrationRequest {
	this := OrgLogIntegrationRequest{}
	this.LogTypes = logTypes
	this.Type = type_
	return &this
}

// NewOrgLogIntegrationRequestWithDefaults instantiates a new OrgLogIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgLogIntegrationRequestWithDefaults() *OrgLogIntegrationRequest {
	this := OrgLogIntegrationRequest{}
	return &this
}

// GetEventBackfill returns the EventBackfill field value if set, zero value otherwise
func (o *OrgLogIntegrationRequest) GetEventBackfill() EventBackfillRequest {
	if o == nil || IsNil(o.EventBackfill) {
		var ret EventBackfillRequest
		return ret
	}
	return *o.EventBackfill
}

// GetEventBackfillOk returns a tuple with the EventBackfill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationRequest) GetEventBackfillOk() (*EventBackfillRequest, bool) {
	if o == nil || IsNil(o.EventBackfill) {
		return nil, false
	}

	return o.EventBackfill, true
}

// HasEventBackfill returns a boolean if a field has been set.
func (o *OrgLogIntegrationRequest) HasEventBackfill() bool {
	if o != nil && !IsNil(o.EventBackfill) {
		return true
	}

	return false
}

// SetEventBackfill gets a reference to the given EventBackfillRequest and assigns it to the EventBackfill field.
func (o *OrgLogIntegrationRequest) SetEventBackfill(v EventBackfillRequest) {
	o.EventBackfill = &v
	o.NullFields = removeNullField(o.NullFields, "EventBackfill")
}

// SetEventBackfillNil sets EventBackfill to an explicit JSON null when marshaled.
func (o *OrgLogIntegrationRequest) SetEventBackfillNil() {
	o.EventBackfill = nil
	o.NullFields = addNullField(o.NullFields, "EventBackfill")
}

// GetLogTypes returns the LogTypes field value
func (o *OrgLogIntegrationRequest) GetLogTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LogTypes
}

// GetLogTypesOk returns a tuple with the LogTypes field value
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationRequest) GetLogTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTypes, true
}

// SetLogTypes sets field value
func (o *OrgLogIntegrationRequest) SetLogTypes(v []string) {
	o.LogTypes = v
}

// GetType returns the Type field value
func (o *OrgLogIntegrationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrgLogIntegrationRequest) SetType(v string) {
	o.Type = v
}

// GetOtelEndpoint returns the OtelEndpoint field value if set, zero value otherwise
func (o *OrgLogIntegrationRequest) GetOtelEndpoint() string {
	if o == nil || IsNil(o.OtelEndpoint) {
		var ret string
		return ret
	}
	return *o.OtelEndpoint
}

// GetOtelEndpointOk returns a tuple with the OtelEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationRequest) GetOtelEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.OtelEndpoint) {
		return nil, false
	}

	return o.OtelEndpoint, true
}

// HasOtelEndpoint returns a boolean if a field has been set.
func (o *OrgLogIntegrationRequest) HasOtelEndpoint() bool {
	if o != nil && !IsNil(o.OtelEndpoint) {
		return true
	}

	return false
}

// SetOtelEndpoint gets a reference to the given string and assigns it to the OtelEndpoint field.
func (o *OrgLogIntegrationRequest) SetOtelEndpoint(v string) {
	o.OtelEndpoint = &v
	o.NullFields = removeNullField(o.NullFields, "OtelEndpoint")
}

// SetOtelEndpointNil sets OtelEndpoint to an explicit JSON null when marshaled.
func (o *OrgLogIntegrationRequest) SetOtelEndpointNil() {
	o.OtelEndpoint = nil
	o.NullFields = addNullField(o.NullFields, "OtelEndpoint")
}

// GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field value if set, zero value otherwise
func (o *OrgLogIntegrationRequest) GetOtelSuppliedHeaders() []OrgLogIntegrationHeader {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		var ret []OrgLogIntegrationHeader
		return ret
	}
	return *o.OtelSuppliedHeaders
}

// GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgLogIntegrationRequest) GetOtelSuppliedHeadersOk() (*[]OrgLogIntegrationHeader, bool) {
	if o == nil || IsNil(o.OtelSuppliedHeaders) {
		return nil, false
	}

	return o.OtelSuppliedHeaders, true
}

// HasOtelSuppliedHeaders returns a boolean if a field has been set.
func (o *OrgLogIntegrationRequest) HasOtelSuppliedHeaders() bool {
	if o != nil && !IsNil(o.OtelSuppliedHeaders) {
		return true
	}

	return false
}

// SetOtelSuppliedHeaders gets a reference to the given []OrgLogIntegrationHeader and assigns it to the OtelSuppliedHeaders field.
func (o *OrgLogIntegrationRequest) SetOtelSuppliedHeaders(v []OrgLogIntegrationHeader) {
	o.OtelSuppliedHeaders = &v
	o.NullFields = removeNullField(o.NullFields, "OtelSuppliedHeaders")
}

// SetOtelSuppliedHeadersNil sets OtelSuppliedHeaders to an explicit JSON null when marshaled.
func (o *OrgLogIntegrationRequest) SetOtelSuppliedHeadersNil() {
	o.OtelSuppliedHeaders = nil
	o.NullFields = addNullField(o.NullFields, "OtelSuppliedHeaders")
}
