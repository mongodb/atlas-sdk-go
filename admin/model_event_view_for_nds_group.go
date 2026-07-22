// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// EventViewForNdsGroup struct for EventViewForNdsGroup
type EventViewForNdsGroup struct {
	// Unique 24-hexadecimal digit string that identifies the API Key that triggered the event. If this resource returns this parameter, it doesn't return the `userId` parameter.
	// Read only field.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// Date and time when this event occurred. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	Created           *time.Time `json:"created,omitempty"`
	DelegatePrincipal *Principal `json:"delegatePrincipal,omitempty"`
	// Unique identifier of event type.
	EventTypeName *string `json:"eventTypeName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The `eventId` identifies the specific event.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Flag that indicates whether a MongoDB employee triggered the specified event.
	// Read only field.
	IsGlobalAdmin *bool `json:"isGlobalAdmin,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	// Read only field.
	OrgId *string `json:"orgId,omitempty"`
	// Public part of the API key that triggered the event. If this resource returns this parameter, it doesn't return the **username** parameter.
	// Read only field.
	PublicKey *string `json:"publicKey,omitempty"`
	Raw       *Raw    `json:"raw,omitempty"`
	// IPv4 or IPv6 address from which the user triggered this event.
	// Read only field.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn't return the `apiKeyId` parameter.
	// Read only field.
	UserId *string `json:"userId,omitempty"`
	// Email address for the user who triggered this event. If this resource returns this parameter, it doesn't return the `publicApiKey` parameter.
	// Read only field.
	Username *string `json:"username,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert associated with the event.
	// Read only field.
	AlertId *string `json:"alertId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the alert configuration associated with the `alertId`.
	// Read only field.
	AlertConfigId *string `json:"alertConfigId,omitempty"`
	// Public part of the API key that this event targets.
	// Read only field.
	TargetPublicKey *string `json:"targetPublicKey,omitempty"`
	// Entry in the list of source host addresses that the API key accepts and this event targets.
	// Read only field.
	WhitelistEntry *string `json:"whitelistEntry,omitempty"`
	// Unique 24-hexadecimal digit string that identifies of the invoice associated with the event.
	// Read only field.
	InvoiceId *string `json:"invoiceId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the invoice payment associated with this event.
	// Read only field.
	PaymentId *string `json:"paymentId,omitempty"`
	// Human-readable label of the shard associated with the event.
	// Read only field.
	ShardName *string `json:"shardName,omitempty"`
	// Human-readable label of the collection on which the event occurred. The resource returns this parameter when the `eventTypeName` includes `DATA_EXPLORER`.
	// Read only field.
	Collection *string `json:"collection,omitempty"`
	// Human-readable label of the database on which this incident occurred. The resource returns this parameter when `\"eventTypeName\" : \"DATA_EXPLORER\"` or `\"eventTypeName\" : \"DATA_EXPLORER_CRUD\"`.
	// Read only field.
	Database *string `json:"database,omitempty"`
	// Action that the database attempted to execute when the event triggered. The response returns this parameter when `eventTypeName\" : \"DATA_EXPLORER\"`.
	// Read only field.
	OpType *string `json:"opType,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Data Explorer session associated with the event.
	// Read only field.
	SessionId *string `json:"sessionId,omitempty"`
	// Desk location of MongoDB employee associated with the event.
	// Read only field.
	DeskLocation *string `json:"deskLocation,omitempty"`
	// Identifier of MongoDB employee associated with the event.
	// Read only field.
	EmployeeIdentifier *string `json:"employeeIdentifier,omitempty"`
	// IANA port on which the MongoDB process listens for requests.
	// Read only field.
	Port *int `json:"port,omitempty"`
	// Human-readable label of the replica set associated with the event.
	// Read only field.
	ReplicaSetName *string            `json:"replicaSetName,omitempty"`
	CurrentValue   *NumberMetricValue `json:"currentValue,omitempty"`
	// Human-readable label of the metric associated with the `alertId`. This field may change type of `currentValue` field.
	// Read only field.
	MetricName *string `json:"metricName,omitempty"`
	// The username of the MongoDB User that was created, deleted, or edited.
	// Read only field.
	DbUserUsername *string `json:"dbUserUsername,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the endpoint associated with this event.
	// Read only field.
	EndpointId *string `json:"endpointId,omitempty"`
	// Unique identification string that the cloud provider uses to identify the private endpoint.
	// Read only field.
	ProviderEndpointId *string `json:"providerEndpointId,omitempty"`
	// Fully qualified domain name (FQDN) of the host associated with the event.
	// Read only field.
	Hostname *string `json:"hostname,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization team associated with this event.
	// Read only field.
	TeamId *string `json:"teamId,omitempty"`
	// Email address for the console user that this event targets. The resource returns this parameter when `\"eventTypeName\" : \"USER\"`.
	// Read only field.
	TargetUsername *string `json:"targetUsername,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the resource associated with the event.
	// Read only field.
	ResourceId *string `json:"resourceId,omitempty"`
	// Unique identifier of resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// Name of the stream processing workspace associated with the event.
	// Read only field.
	InstanceName *string `json:"instanceName,omitempty"`
	// Tier the stream processor scaled from.
	// Read only field.
	FromTier *string `json:"fromTier,omitempty"`
	// Username of the user who modified the stream processor.
	// Read only field.
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// Error message linked to the stream processor associated with the event.
	// Read only field.
	ProcessorErrorMsg *string `json:"processorErrorMsg,omitempty"`
	// Name of the stream processor associated with the event.
	// Read only field.
	ProcessorName *string `json:"processorName,omitempty"`
	// State of the stream processor associated with the event.
	// Read only field.
	ProcessorState *string `json:"processorState,omitempty"`
	// Reason for the autoscale event.
	// Read only field.
	Reason *string `json:"reason,omitempty"`
	// Tier the stream processor scaled to.
	// Read only field.
	ToTier *string `json:"toTier,omitempty"`
	// Unique 24-hexadecimal character string that identifies the resource policy.
	// Read only field.
	ResourcePolicyId *string `json:"resourcePolicyId,omitempty"`
	// String representation of the violated resource policy ids.
	ViolatedPolicies *[]string `json:"violatedPolicies,omitempty"`
	// Resource policy action taken by the user and evaluated against the currently active policies.
	// Read only field.
	ViolationAction *string `json:"violationAction,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *EventViewForNdsGroup) MarshalJSON() ([]byte, error) {
	type noMethod EventViewForNdsGroup
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewEventViewForNdsGroup instantiates a new EventViewForNdsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventViewForNdsGroup() *EventViewForNdsGroup {
	this := EventViewForNdsGroup{}
	return &this
}

// NewEventViewForNdsGroupWithDefaults instantiates a new EventViewForNdsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventViewForNdsGroupWithDefaults() *EventViewForNdsGroup {
	this := EventViewForNdsGroup{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetApiKeyId() string {
	if o == nil || IsNil(o.ApiKeyId) {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetApiKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKeyId) {
		return nil, false
	}

	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasApiKeyId() bool {
	if o != nil && !IsNil(o.ApiKeyId) {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *EventViewForNdsGroup) SetApiKeyId(v string) {
	o.ApiKeyId = &v
	o.NullFields = removeNullField(o.NullFields, "ApiKeyId")
}

// SetApiKeyIdNil sets ApiKeyId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetApiKeyIdNil() {
	o.ApiKeyId = nil
	o.NullFields = addNullField(o.NullFields, "ApiKeyId")
}

// GetCreated returns the Created field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}

	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EventViewForNdsGroup) SetCreated(v time.Time) {
	o.Created = &v
	o.NullFields = removeNullField(o.NullFields, "Created")
}

// SetCreatedNil sets Created to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetCreatedNil() {
	o.Created = nil
	o.NullFields = addNullField(o.NullFields, "Created")
}

// GetDelegatePrincipal returns the DelegatePrincipal field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetDelegatePrincipal() Principal {
	if o == nil || IsNil(o.DelegatePrincipal) {
		var ret Principal
		return ret
	}
	return *o.DelegatePrincipal
}

// GetDelegatePrincipalOk returns a tuple with the DelegatePrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetDelegatePrincipalOk() (*Principal, bool) {
	if o == nil || IsNil(o.DelegatePrincipal) {
		return nil, false
	}

	return o.DelegatePrincipal, true
}

// HasDelegatePrincipal returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasDelegatePrincipal() bool {
	if o != nil && !IsNil(o.DelegatePrincipal) {
		return true
	}

	return false
}

// SetDelegatePrincipal gets a reference to the given Principal and assigns it to the DelegatePrincipal field.
func (o *EventViewForNdsGroup) SetDelegatePrincipal(v Principal) {
	o.DelegatePrincipal = &v
	o.NullFields = removeNullField(o.NullFields, "DelegatePrincipal")
}

// SetDelegatePrincipalNil sets DelegatePrincipal to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetDelegatePrincipalNil() {
	o.DelegatePrincipal = nil
	o.NullFields = addNullField(o.NullFields, "DelegatePrincipal")
}

// GetEventTypeName returns the EventTypeName field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetEventTypeName() string {
	if o == nil || IsNil(o.EventTypeName) {
		var ret string
		return ret
	}
	return *o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetEventTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventTypeName) {
		return nil, false
	}

	return o.EventTypeName, true
}

// HasEventTypeName returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasEventTypeName() bool {
	if o != nil && !IsNil(o.EventTypeName) {
		return true
	}

	return false
}

// SetEventTypeName gets a reference to the given string and assigns it to the EventTypeName field.
func (o *EventViewForNdsGroup) SetEventTypeName(v string) {
	o.EventTypeName = &v
	o.NullFields = removeNullField(o.NullFields, "EventTypeName")
}

// SetEventTypeNameNil sets EventTypeName to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetEventTypeNameNil() {
	o.EventTypeName = nil
	o.NullFields = addNullField(o.NullFields, "EventTypeName")
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *EventViewForNdsGroup) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetId returns the Id field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventViewForNdsGroup) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetIsGlobalAdmin returns the IsGlobalAdmin field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetIsGlobalAdmin() bool {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		var ret bool
		return ret
	}
	return *o.IsGlobalAdmin
}

// GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetIsGlobalAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		return nil, false
	}

	return o.IsGlobalAdmin, true
}

// HasIsGlobalAdmin returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasIsGlobalAdmin() bool {
	if o != nil && !IsNil(o.IsGlobalAdmin) {
		return true
	}

	return false
}

// SetIsGlobalAdmin gets a reference to the given bool and assigns it to the IsGlobalAdmin field.
func (o *EventViewForNdsGroup) SetIsGlobalAdmin(v bool) {
	o.IsGlobalAdmin = &v
	o.NullFields = removeNullField(o.NullFields, "IsGlobalAdmin")
}

// SetIsGlobalAdminNil sets IsGlobalAdmin to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetIsGlobalAdminNil() {
	o.IsGlobalAdmin = nil
	o.NullFields = addNullField(o.NullFields, "IsGlobalAdmin")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *EventViewForNdsGroup) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetOrgId returns the OrgId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}

	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *EventViewForNdsGroup) SetOrgId(v string) {
	o.OrgId = &v
	o.NullFields = removeNullField(o.NullFields, "OrgId")
}

// SetOrgIdNil sets OrgId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetOrgIdNil() {
	o.OrgId = nil
	o.NullFields = addNullField(o.NullFields, "OrgId")
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}

	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *EventViewForNdsGroup) SetPublicKey(v string) {
	o.PublicKey = &v
	o.NullFields = removeNullField(o.NullFields, "PublicKey")
}

// SetPublicKeyNil sets PublicKey to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetPublicKeyNil() {
	o.PublicKey = nil
	o.NullFields = addNullField(o.NullFields, "PublicKey")
}

// GetRaw returns the Raw field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetRaw() Raw {
	if o == nil || IsNil(o.Raw) {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetRawOk() (*Raw, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}

	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *EventViewForNdsGroup) SetRaw(v Raw) {
	o.Raw = &v
	o.NullFields = removeNullField(o.NullFields, "Raw")
}

// SetRawNil sets Raw to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetRawNil() {
	o.Raw = nil
	o.NullFields = addNullField(o.NullFields, "Raw")
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetRemoteAddress() string {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetRemoteAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}

	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *EventViewForNdsGroup) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
	o.NullFields = removeNullField(o.NullFields, "RemoteAddress")
}

// SetRemoteAddressNil sets RemoteAddress to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetRemoteAddressNil() {
	o.RemoteAddress = nil
	o.NullFields = addNullField(o.NullFields, "RemoteAddress")
}

// GetUserId returns the UserId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}

	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EventViewForNdsGroup) SetUserId(v string) {
	o.UserId = &v
	o.NullFields = removeNullField(o.NullFields, "UserId")
}

// SetUserIdNil sets UserId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetUserIdNil() {
	o.UserId = nil
	o.NullFields = addNullField(o.NullFields, "UserId")
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EventViewForNdsGroup) SetUsername(v string) {
	o.Username = &v
	o.NullFields = removeNullField(o.NullFields, "Username")
}

// SetUsernameNil sets Username to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetUsernameNil() {
	o.Username = nil
	o.NullFields = addNullField(o.NullFields, "Username")
}

// GetAlertId returns the AlertId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetAlertId() string {
	if o == nil || IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetAlertIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertId) {
		return nil, false
	}

	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasAlertId() bool {
	if o != nil && !IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *EventViewForNdsGroup) SetAlertId(v string) {
	o.AlertId = &v
	o.NullFields = removeNullField(o.NullFields, "AlertId")
}

// SetAlertIdNil sets AlertId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetAlertIdNil() {
	o.AlertId = nil
	o.NullFields = addNullField(o.NullFields, "AlertId")
}

// GetAlertConfigId returns the AlertConfigId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetAlertConfigId() string {
	if o == nil || IsNil(o.AlertConfigId) {
		var ret string
		return ret
	}
	return *o.AlertConfigId
}

// GetAlertConfigIdOk returns a tuple with the AlertConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetAlertConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertConfigId) {
		return nil, false
	}

	return o.AlertConfigId, true
}

// HasAlertConfigId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasAlertConfigId() bool {
	if o != nil && !IsNil(o.AlertConfigId) {
		return true
	}

	return false
}

// SetAlertConfigId gets a reference to the given string and assigns it to the AlertConfigId field.
func (o *EventViewForNdsGroup) SetAlertConfigId(v string) {
	o.AlertConfigId = &v
	o.NullFields = removeNullField(o.NullFields, "AlertConfigId")
}

// SetAlertConfigIdNil sets AlertConfigId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetAlertConfigIdNil() {
	o.AlertConfigId = nil
	o.NullFields = addNullField(o.NullFields, "AlertConfigId")
}

// GetTargetPublicKey returns the TargetPublicKey field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetTargetPublicKey() string {
	if o == nil || IsNil(o.TargetPublicKey) {
		var ret string
		return ret
	}
	return *o.TargetPublicKey
}

// GetTargetPublicKeyOk returns a tuple with the TargetPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetTargetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPublicKey) {
		return nil, false
	}

	return o.TargetPublicKey, true
}

// HasTargetPublicKey returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasTargetPublicKey() bool {
	if o != nil && !IsNil(o.TargetPublicKey) {
		return true
	}

	return false
}

// SetTargetPublicKey gets a reference to the given string and assigns it to the TargetPublicKey field.
func (o *EventViewForNdsGroup) SetTargetPublicKey(v string) {
	o.TargetPublicKey = &v
	o.NullFields = removeNullField(o.NullFields, "TargetPublicKey")
}

// SetTargetPublicKeyNil sets TargetPublicKey to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetTargetPublicKeyNil() {
	o.TargetPublicKey = nil
	o.NullFields = addNullField(o.NullFields, "TargetPublicKey")
}

// GetWhitelistEntry returns the WhitelistEntry field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetWhitelistEntry() string {
	if o == nil || IsNil(o.WhitelistEntry) {
		var ret string
		return ret
	}
	return *o.WhitelistEntry
}

// GetWhitelistEntryOk returns a tuple with the WhitelistEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetWhitelistEntryOk() (*string, bool) {
	if o == nil || IsNil(o.WhitelistEntry) {
		return nil, false
	}

	return o.WhitelistEntry, true
}

// HasWhitelistEntry returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasWhitelistEntry() bool {
	if o != nil && !IsNil(o.WhitelistEntry) {
		return true
	}

	return false
}

// SetWhitelistEntry gets a reference to the given string and assigns it to the WhitelistEntry field.
func (o *EventViewForNdsGroup) SetWhitelistEntry(v string) {
	o.WhitelistEntry = &v
	o.NullFields = removeNullField(o.NullFields, "WhitelistEntry")
}

// SetWhitelistEntryNil sets WhitelistEntry to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetWhitelistEntryNil() {
	o.WhitelistEntry = nil
	o.NullFields = addNullField(o.NullFields, "WhitelistEntry")
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}

	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *EventViewForNdsGroup) SetInvoiceId(v string) {
	o.InvoiceId = &v
	o.NullFields = removeNullField(o.NullFields, "InvoiceId")
}

// SetInvoiceIdNil sets InvoiceId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetInvoiceIdNil() {
	o.InvoiceId = nil
	o.NullFields = addNullField(o.NullFields, "InvoiceId")
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}

	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *EventViewForNdsGroup) SetPaymentId(v string) {
	o.PaymentId = &v
	o.NullFields = removeNullField(o.NullFields, "PaymentId")
}

// SetPaymentIdNil sets PaymentId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetPaymentIdNil() {
	o.PaymentId = nil
	o.NullFields = addNullField(o.NullFields, "PaymentId")
}

// GetShardName returns the ShardName field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetShardName() string {
	if o == nil || IsNil(o.ShardName) {
		var ret string
		return ret
	}
	return *o.ShardName
}

// GetShardNameOk returns a tuple with the ShardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetShardNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShardName) {
		return nil, false
	}

	return o.ShardName, true
}

// HasShardName returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasShardName() bool {
	if o != nil && !IsNil(o.ShardName) {
		return true
	}

	return false
}

// SetShardName gets a reference to the given string and assigns it to the ShardName field.
func (o *EventViewForNdsGroup) SetShardName(v string) {
	o.ShardName = &v
	o.NullFields = removeNullField(o.NullFields, "ShardName")
}

// SetShardNameNil sets ShardName to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetShardNameNil() {
	o.ShardName = nil
	o.NullFields = addNullField(o.NullFields, "ShardName")
}

// GetCollection returns the Collection field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetCollection() string {
	if o == nil || IsNil(o.Collection) {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetCollectionOk() (*string, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}

	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *EventViewForNdsGroup) SetCollection(v string) {
	o.Collection = &v
	o.NullFields = removeNullField(o.NullFields, "Collection")
}

// SetCollectionNil sets Collection to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetCollectionNil() {
	o.Collection = nil
	o.NullFields = addNullField(o.NullFields, "Collection")
}

// GetDatabase returns the Database field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}

	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *EventViewForNdsGroup) SetDatabase(v string) {
	o.Database = &v
	o.NullFields = removeNullField(o.NullFields, "Database")
}

// SetDatabaseNil sets Database to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetDatabaseNil() {
	o.Database = nil
	o.NullFields = addNullField(o.NullFields, "Database")
}

// GetOpType returns the OpType field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetOpType() string {
	if o == nil || IsNil(o.OpType) {
		var ret string
		return ret
	}
	return *o.OpType
}

// GetOpTypeOk returns a tuple with the OpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetOpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OpType) {
		return nil, false
	}

	return o.OpType, true
}

// HasOpType returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasOpType() bool {
	if o != nil && !IsNil(o.OpType) {
		return true
	}

	return false
}

// SetOpType gets a reference to the given string and assigns it to the OpType field.
func (o *EventViewForNdsGroup) SetOpType(v string) {
	o.OpType = &v
	o.NullFields = removeNullField(o.NullFields, "OpType")
}

// SetOpTypeNil sets OpType to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetOpTypeNil() {
	o.OpType = nil
	o.NullFields = addNullField(o.NullFields, "OpType")
}

// GetSessionId returns the SessionId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}

	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *EventViewForNdsGroup) SetSessionId(v string) {
	o.SessionId = &v
	o.NullFields = removeNullField(o.NullFields, "SessionId")
}

// SetSessionIdNil sets SessionId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetSessionIdNil() {
	o.SessionId = nil
	o.NullFields = addNullField(o.NullFields, "SessionId")
}

// GetDeskLocation returns the DeskLocation field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetDeskLocation() string {
	if o == nil || IsNil(o.DeskLocation) {
		var ret string
		return ret
	}
	return *o.DeskLocation
}

// GetDeskLocationOk returns a tuple with the DeskLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetDeskLocationOk() (*string, bool) {
	if o == nil || IsNil(o.DeskLocation) {
		return nil, false
	}

	return o.DeskLocation, true
}

// HasDeskLocation returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasDeskLocation() bool {
	if o != nil && !IsNil(o.DeskLocation) {
		return true
	}

	return false
}

// SetDeskLocation gets a reference to the given string and assigns it to the DeskLocation field.
func (o *EventViewForNdsGroup) SetDeskLocation(v string) {
	o.DeskLocation = &v
	o.NullFields = removeNullField(o.NullFields, "DeskLocation")
}

// SetDeskLocationNil sets DeskLocation to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetDeskLocationNil() {
	o.DeskLocation = nil
	o.NullFields = addNullField(o.NullFields, "DeskLocation")
}

// GetEmployeeIdentifier returns the EmployeeIdentifier field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetEmployeeIdentifier() string {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		var ret string
		return ret
	}
	return *o.EmployeeIdentifier
}

// GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetEmployeeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeIdentifier) {
		return nil, false
	}

	return o.EmployeeIdentifier, true
}

// HasEmployeeIdentifier returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasEmployeeIdentifier() bool {
	if o != nil && !IsNil(o.EmployeeIdentifier) {
		return true
	}

	return false
}

// SetEmployeeIdentifier gets a reference to the given string and assigns it to the EmployeeIdentifier field.
func (o *EventViewForNdsGroup) SetEmployeeIdentifier(v string) {
	o.EmployeeIdentifier = &v
	o.NullFields = removeNullField(o.NullFields, "EmployeeIdentifier")
}

// SetEmployeeIdentifierNil sets EmployeeIdentifier to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetEmployeeIdentifierNil() {
	o.EmployeeIdentifier = nil
	o.NullFields = addNullField(o.NullFields, "EmployeeIdentifier")
}

// GetPort returns the Port field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetPort() int {
	if o == nil || IsNil(o.Port) {
		var ret int
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetPortOk() (*int, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}

	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int and assigns it to the Port field.
func (o *EventViewForNdsGroup) SetPort(v int) {
	o.Port = &v
	o.NullFields = removeNullField(o.NullFields, "Port")
}

// SetPortNil sets Port to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetPortNil() {
	o.Port = nil
	o.NullFields = addNullField(o.NullFields, "Port")
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *EventViewForNdsGroup) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
	o.NullFields = removeNullField(o.NullFields, "ReplicaSetName")
}

// SetReplicaSetNameNil sets ReplicaSetName to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetReplicaSetNameNil() {
	o.ReplicaSetName = nil
	o.NullFields = addNullField(o.NullFields, "ReplicaSetName")
}

// GetCurrentValue returns the CurrentValue field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetCurrentValue() NumberMetricValue {
	if o == nil || IsNil(o.CurrentValue) {
		var ret NumberMetricValue
		return ret
	}
	return *o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetCurrentValueOk() (*NumberMetricValue, bool) {
	if o == nil || IsNil(o.CurrentValue) {
		return nil, false
	}

	return o.CurrentValue, true
}

// HasCurrentValue returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasCurrentValue() bool {
	if o != nil && !IsNil(o.CurrentValue) {
		return true
	}

	return false
}

// SetCurrentValue gets a reference to the given NumberMetricValue and assigns it to the CurrentValue field.
func (o *EventViewForNdsGroup) SetCurrentValue(v NumberMetricValue) {
	o.CurrentValue = &v
	o.NullFields = removeNullField(o.NullFields, "CurrentValue")
}

// SetCurrentValueNil sets CurrentValue to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetCurrentValueNil() {
	o.CurrentValue = nil
	o.NullFields = addNullField(o.NullFields, "CurrentValue")
}

// GetMetricName returns the MetricName field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}

	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *EventViewForNdsGroup) SetMetricName(v string) {
	o.MetricName = &v
	o.NullFields = removeNullField(o.NullFields, "MetricName")
}

// SetMetricNameNil sets MetricName to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetMetricNameNil() {
	o.MetricName = nil
	o.NullFields = addNullField(o.NullFields, "MetricName")
}

// GetDbUserUsername returns the DbUserUsername field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetDbUserUsername() string {
	if o == nil || IsNil(o.DbUserUsername) {
		var ret string
		return ret
	}
	return *o.DbUserUsername
}

// GetDbUserUsernameOk returns a tuple with the DbUserUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetDbUserUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.DbUserUsername) {
		return nil, false
	}

	return o.DbUserUsername, true
}

// HasDbUserUsername returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasDbUserUsername() bool {
	if o != nil && !IsNil(o.DbUserUsername) {
		return true
	}

	return false
}

// SetDbUserUsername gets a reference to the given string and assigns it to the DbUserUsername field.
func (o *EventViewForNdsGroup) SetDbUserUsername(v string) {
	o.DbUserUsername = &v
	o.NullFields = removeNullField(o.NullFields, "DbUserUsername")
}

// SetDbUserUsernameNil sets DbUserUsername to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetDbUserUsernameNil() {
	o.DbUserUsername = nil
	o.NullFields = addNullField(o.NullFields, "DbUserUsername")
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetEndpointId() string {
	if o == nil || IsNil(o.EndpointId) {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointId) {
		return nil, false
	}

	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasEndpointId() bool {
	if o != nil && !IsNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *EventViewForNdsGroup) SetEndpointId(v string) {
	o.EndpointId = &v
	o.NullFields = removeNullField(o.NullFields, "EndpointId")
}

// SetEndpointIdNil sets EndpointId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetEndpointIdNil() {
	o.EndpointId = nil
	o.NullFields = addNullField(o.NullFields, "EndpointId")
}

// GetProviderEndpointId returns the ProviderEndpointId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetProviderEndpointId() string {
	if o == nil || IsNil(o.ProviderEndpointId) {
		var ret string
		return ret
	}
	return *o.ProviderEndpointId
}

// GetProviderEndpointIdOk returns a tuple with the ProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetProviderEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderEndpointId) {
		return nil, false
	}

	return o.ProviderEndpointId, true
}

// HasProviderEndpointId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasProviderEndpointId() bool {
	if o != nil && !IsNil(o.ProviderEndpointId) {
		return true
	}

	return false
}

// SetProviderEndpointId gets a reference to the given string and assigns it to the ProviderEndpointId field.
func (o *EventViewForNdsGroup) SetProviderEndpointId(v string) {
	o.ProviderEndpointId = &v
	o.NullFields = removeNullField(o.NullFields, "ProviderEndpointId")
}

// SetProviderEndpointIdNil sets ProviderEndpointId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetProviderEndpointIdNil() {
	o.ProviderEndpointId = nil
	o.NullFields = addNullField(o.NullFields, "ProviderEndpointId")
}

// GetHostname returns the Hostname field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}

	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *EventViewForNdsGroup) SetHostname(v string) {
	o.Hostname = &v
	o.NullFields = removeNullField(o.NullFields, "Hostname")
}

// SetHostnameNil sets Hostname to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetHostnameNil() {
	o.Hostname = nil
	o.NullFields = addNullField(o.NullFields, "Hostname")
}

// GetTeamId returns the TeamId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}

	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *EventViewForNdsGroup) SetTeamId(v string) {
	o.TeamId = &v
	o.NullFields = removeNullField(o.NullFields, "TeamId")
}

// SetTeamIdNil sets TeamId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetTeamIdNil() {
	o.TeamId = nil
	o.NullFields = addNullField(o.NullFields, "TeamId")
}

// GetTargetUsername returns the TargetUsername field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetTargetUsername() string {
	if o == nil || IsNil(o.TargetUsername) {
		var ret string
		return ret
	}
	return *o.TargetUsername
}

// GetTargetUsernameOk returns a tuple with the TargetUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetTargetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUsername) {
		return nil, false
	}

	return o.TargetUsername, true
}

// HasTargetUsername returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasTargetUsername() bool {
	if o != nil && !IsNil(o.TargetUsername) {
		return true
	}

	return false
}

// SetTargetUsername gets a reference to the given string and assigns it to the TargetUsername field.
func (o *EventViewForNdsGroup) SetTargetUsername(v string) {
	o.TargetUsername = &v
	o.NullFields = removeNullField(o.NullFields, "TargetUsername")
}

// SetTargetUsernameNil sets TargetUsername to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetTargetUsernameNil() {
	o.TargetUsername = nil
	o.NullFields = addNullField(o.NullFields, "TargetUsername")
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}

	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *EventViewForNdsGroup) SetResourceId(v string) {
	o.ResourceId = &v
	o.NullFields = removeNullField(o.NullFields, "ResourceId")
}

// SetResourceIdNil sets ResourceId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetResourceIdNil() {
	o.ResourceId = nil
	o.NullFields = addNullField(o.NullFields, "ResourceId")
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}

	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *EventViewForNdsGroup) SetResourceType(v string) {
	o.ResourceType = &v
	o.NullFields = removeNullField(o.NullFields, "ResourceType")
}

// SetResourceTypeNil sets ResourceType to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetResourceTypeNil() {
	o.ResourceType = nil
	o.NullFields = addNullField(o.NullFields, "ResourceType")
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName) {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceName) {
		return nil, false
	}

	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasInstanceName() bool {
	if o != nil && !IsNil(o.InstanceName) {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *EventViewForNdsGroup) SetInstanceName(v string) {
	o.InstanceName = &v
	o.NullFields = removeNullField(o.NullFields, "InstanceName")
}

// SetInstanceNameNil sets InstanceName to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetInstanceNameNil() {
	o.InstanceName = nil
	o.NullFields = addNullField(o.NullFields, "InstanceName")
}

// GetFromTier returns the FromTier field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetFromTier() string {
	if o == nil || IsNil(o.FromTier) {
		var ret string
		return ret
	}
	return *o.FromTier
}

// GetFromTierOk returns a tuple with the FromTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetFromTierOk() (*string, bool) {
	if o == nil || IsNil(o.FromTier) {
		return nil, false
	}

	return o.FromTier, true
}

// HasFromTier returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasFromTier() bool {
	if o != nil && !IsNil(o.FromTier) {
		return true
	}

	return false
}

// SetFromTier gets a reference to the given string and assigns it to the FromTier field.
func (o *EventViewForNdsGroup) SetFromTier(v string) {
	o.FromTier = &v
	o.NullFields = removeNullField(o.NullFields, "FromTier")
}

// SetFromTierNil sets FromTier to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetFromTierNil() {
	o.FromTier = nil
	o.NullFields = addNullField(o.NullFields, "FromTier")
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}

	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *EventViewForNdsGroup) SetModifiedBy(v string) {
	o.ModifiedBy = &v
	o.NullFields = removeNullField(o.NullFields, "ModifiedBy")
}

// SetModifiedByNil sets ModifiedBy to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetModifiedByNil() {
	o.ModifiedBy = nil
	o.NullFields = addNullField(o.NullFields, "ModifiedBy")
}

// GetProcessorErrorMsg returns the ProcessorErrorMsg field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetProcessorErrorMsg() string {
	if o == nil || IsNil(o.ProcessorErrorMsg) {
		var ret string
		return ret
	}
	return *o.ProcessorErrorMsg
}

// GetProcessorErrorMsgOk returns a tuple with the ProcessorErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetProcessorErrorMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorErrorMsg) {
		return nil, false
	}

	return o.ProcessorErrorMsg, true
}

// HasProcessorErrorMsg returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasProcessorErrorMsg() bool {
	if o != nil && !IsNil(o.ProcessorErrorMsg) {
		return true
	}

	return false
}

// SetProcessorErrorMsg gets a reference to the given string and assigns it to the ProcessorErrorMsg field.
func (o *EventViewForNdsGroup) SetProcessorErrorMsg(v string) {
	o.ProcessorErrorMsg = &v
	o.NullFields = removeNullField(o.NullFields, "ProcessorErrorMsg")
}

// SetProcessorErrorMsgNil sets ProcessorErrorMsg to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetProcessorErrorMsgNil() {
	o.ProcessorErrorMsg = nil
	o.NullFields = addNullField(o.NullFields, "ProcessorErrorMsg")
}

// GetProcessorName returns the ProcessorName field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetProcessorName() string {
	if o == nil || IsNil(o.ProcessorName) {
		var ret string
		return ret
	}
	return *o.ProcessorName
}

// GetProcessorNameOk returns a tuple with the ProcessorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetProcessorNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorName) {
		return nil, false
	}

	return o.ProcessorName, true
}

// HasProcessorName returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasProcessorName() bool {
	if o != nil && !IsNil(o.ProcessorName) {
		return true
	}

	return false
}

// SetProcessorName gets a reference to the given string and assigns it to the ProcessorName field.
func (o *EventViewForNdsGroup) SetProcessorName(v string) {
	o.ProcessorName = &v
	o.NullFields = removeNullField(o.NullFields, "ProcessorName")
}

// SetProcessorNameNil sets ProcessorName to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetProcessorNameNil() {
	o.ProcessorName = nil
	o.NullFields = addNullField(o.NullFields, "ProcessorName")
}

// GetProcessorState returns the ProcessorState field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetProcessorState() string {
	if o == nil || IsNil(o.ProcessorState) {
		var ret string
		return ret
	}
	return *o.ProcessorState
}

// GetProcessorStateOk returns a tuple with the ProcessorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetProcessorStateOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessorState) {
		return nil, false
	}

	return o.ProcessorState, true
}

// HasProcessorState returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasProcessorState() bool {
	if o != nil && !IsNil(o.ProcessorState) {
		return true
	}

	return false
}

// SetProcessorState gets a reference to the given string and assigns it to the ProcessorState field.
func (o *EventViewForNdsGroup) SetProcessorState(v string) {
	o.ProcessorState = &v
	o.NullFields = removeNullField(o.NullFields, "ProcessorState")
}

// SetProcessorStateNil sets ProcessorState to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetProcessorStateNil() {
	o.ProcessorState = nil
	o.NullFields = addNullField(o.NullFields, "ProcessorState")
}

// GetReason returns the Reason field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}

	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EventViewForNdsGroup) SetReason(v string) {
	o.Reason = &v
	o.NullFields = removeNullField(o.NullFields, "Reason")
}

// SetReasonNil sets Reason to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetReasonNil() {
	o.Reason = nil
	o.NullFields = addNullField(o.NullFields, "Reason")
}

// GetToTier returns the ToTier field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetToTier() string {
	if o == nil || IsNil(o.ToTier) {
		var ret string
		return ret
	}
	return *o.ToTier
}

// GetToTierOk returns a tuple with the ToTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetToTierOk() (*string, bool) {
	if o == nil || IsNil(o.ToTier) {
		return nil, false
	}

	return o.ToTier, true
}

// HasToTier returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasToTier() bool {
	if o != nil && !IsNil(o.ToTier) {
		return true
	}

	return false
}

// SetToTier gets a reference to the given string and assigns it to the ToTier field.
func (o *EventViewForNdsGroup) SetToTier(v string) {
	o.ToTier = &v
	o.NullFields = removeNullField(o.NullFields, "ToTier")
}

// SetToTierNil sets ToTier to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetToTierNil() {
	o.ToTier = nil
	o.NullFields = addNullField(o.NullFields, "ToTier")
}

// GetResourcePolicyId returns the ResourcePolicyId field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetResourcePolicyId() string {
	if o == nil || IsNil(o.ResourcePolicyId) {
		var ret string
		return ret
	}
	return *o.ResourcePolicyId
}

// GetResourcePolicyIdOk returns a tuple with the ResourcePolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetResourcePolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePolicyId) {
		return nil, false
	}

	return o.ResourcePolicyId, true
}

// HasResourcePolicyId returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasResourcePolicyId() bool {
	if o != nil && !IsNil(o.ResourcePolicyId) {
		return true
	}

	return false
}

// SetResourcePolicyId gets a reference to the given string and assigns it to the ResourcePolicyId field.
func (o *EventViewForNdsGroup) SetResourcePolicyId(v string) {
	o.ResourcePolicyId = &v
	o.NullFields = removeNullField(o.NullFields, "ResourcePolicyId")
}

// SetResourcePolicyIdNil sets ResourcePolicyId to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetResourcePolicyIdNil() {
	o.ResourcePolicyId = nil
	o.NullFields = addNullField(o.NullFields, "ResourcePolicyId")
}

// GetViolatedPolicies returns the ViolatedPolicies field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetViolatedPolicies() []string {
	if o == nil || IsNil(o.ViolatedPolicies) {
		var ret []string
		return ret
	}
	return *o.ViolatedPolicies
}

// GetViolatedPoliciesOk returns a tuple with the ViolatedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetViolatedPoliciesOk() (*[]string, bool) {
	if o == nil || IsNil(o.ViolatedPolicies) {
		return nil, false
	}

	return o.ViolatedPolicies, true
}

// HasViolatedPolicies returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasViolatedPolicies() bool {
	if o != nil && !IsNil(o.ViolatedPolicies) {
		return true
	}

	return false
}

// SetViolatedPolicies gets a reference to the given []string and assigns it to the ViolatedPolicies field.
func (o *EventViewForNdsGroup) SetViolatedPolicies(v []string) {
	o.ViolatedPolicies = &v
	o.NullFields = removeNullField(o.NullFields, "ViolatedPolicies")
}

// SetViolatedPoliciesNil sets ViolatedPolicies to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetViolatedPoliciesNil() {
	o.ViolatedPolicies = nil
	o.NullFields = addNullField(o.NullFields, "ViolatedPolicies")
}

// GetViolationAction returns the ViolationAction field value if set, zero value otherwise
func (o *EventViewForNdsGroup) GetViolationAction() string {
	if o == nil || IsNil(o.ViolationAction) {
		var ret string
		return ret
	}
	return *o.ViolationAction
}

// GetViolationActionOk returns a tuple with the ViolationAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventViewForNdsGroup) GetViolationActionOk() (*string, bool) {
	if o == nil || IsNil(o.ViolationAction) {
		return nil, false
	}

	return o.ViolationAction, true
}

// HasViolationAction returns a boolean if a field has been set.
func (o *EventViewForNdsGroup) HasViolationAction() bool {
	if o != nil && !IsNil(o.ViolationAction) {
		return true
	}

	return false
}

// SetViolationAction gets a reference to the given string and assigns it to the ViolationAction field.
func (o *EventViewForNdsGroup) SetViolationAction(v string) {
	o.ViolationAction = &v
	o.NullFields = removeNullField(o.NullFields, "ViolationAction")
}

// SetViolationActionNil sets ViolationAction to an explicit JSON null when marshaled.
func (o *EventViewForNdsGroup) SetViolationActionNil() {
	o.ViolationAction = nil
	o.NullFields = addNullField(o.NullFields, "ViolationAction")
}
