// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// VPCPeeringConnection Represents a VPC Peering connection on AWS.
type VPCPeeringConnection struct {
	// Internal VPC Peering Connection ID.
	Id *string `json:"_id,omitempty"`
	// The account ID responsible for accepting the request.
	AccepterAccountId *string `json:"accepterAccountId,omitempty"`
	// The CIDR block for the accepter VPC.
	AccepterCidr *string `json:"accepterCidr,omitempty"`
	// The VPC ID accepting the request.
	AccepterVpcId *string `json:"accepterVpcId,omitempty"`
	// The status in the cloud provider for this connection.
	CloudStatus *string `json:"cloudStatus,omitempty"`
	// The time when the VPC Peering Connection request expires. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	// The internal project ID.
	GroupId *string `json:"groupId,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// The local status of the VPC Peering Connection.
	LocalStatus *string `json:"localStatus,omitempty"`
	// Unique VPC Peering Connection name.
	Name *string `json:"name,omitempty"`
	// The account ID requesting the VPC Peering connection.
	RequesterAccountId *string `json:"requesterAccountId,omitempty"`
	// The CIDR block for the requesting VPC.
	RequesterCidr *string `json:"requesterCidr,omitempty"`
	// The VPC ID requesting the VPC Peering connection.
	RequesterVpcId *string `json:"requesterVpcId,omitempty"`
	// A status message.
	StatusMessage *string `json:"statusMessage,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *VPCPeeringConnection) MarshalJSON() ([]byte, error) {
	type noMethod VPCPeeringConnection
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewVPCPeeringConnection instantiates a new VPCPeeringConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVPCPeeringConnection() *VPCPeeringConnection {
	this := VPCPeeringConnection{}
	return &this
}

// NewVPCPeeringConnectionWithDefaults instantiates a new VPCPeeringConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVPCPeeringConnectionWithDefaults() *VPCPeeringConnection {
	this := VPCPeeringConnection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VPCPeeringConnection) SetId(v string) {
	o.Id = &v
	o.NullFields = removeNullField(o.NullFields, "Id")
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetIdNil() {
	o.Id = nil
	o.NullFields = addNullField(o.NullFields, "Id")
}

// GetAccepterAccountId returns the AccepterAccountId field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetAccepterAccountId() string {
	if o == nil || IsNil(o.AccepterAccountId) {
		var ret string
		return ret
	}
	return *o.AccepterAccountId
}

// GetAccepterAccountIdOk returns a tuple with the AccepterAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetAccepterAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccepterAccountId) {
		return nil, false
	}

	return o.AccepterAccountId, true
}

// HasAccepterAccountId returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasAccepterAccountId() bool {
	if o != nil && !IsNil(o.AccepterAccountId) {
		return true
	}

	return false
}

// SetAccepterAccountId gets a reference to the given string and assigns it to the AccepterAccountId field.
func (o *VPCPeeringConnection) SetAccepterAccountId(v string) {
	o.AccepterAccountId = &v
	o.NullFields = removeNullField(o.NullFields, "AccepterAccountId")
}

// SetAccepterAccountIdNil sets AccepterAccountId to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetAccepterAccountIdNil() {
	o.AccepterAccountId = nil
	o.NullFields = addNullField(o.NullFields, "AccepterAccountId")
}

// GetAccepterCidr returns the AccepterCidr field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetAccepterCidr() string {
	if o == nil || IsNil(o.AccepterCidr) {
		var ret string
		return ret
	}
	return *o.AccepterCidr
}

// GetAccepterCidrOk returns a tuple with the AccepterCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetAccepterCidrOk() (*string, bool) {
	if o == nil || IsNil(o.AccepterCidr) {
		return nil, false
	}

	return o.AccepterCidr, true
}

// HasAccepterCidr returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasAccepterCidr() bool {
	if o != nil && !IsNil(o.AccepterCidr) {
		return true
	}

	return false
}

// SetAccepterCidr gets a reference to the given string and assigns it to the AccepterCidr field.
func (o *VPCPeeringConnection) SetAccepterCidr(v string) {
	o.AccepterCidr = &v
	o.NullFields = removeNullField(o.NullFields, "AccepterCidr")
}

// SetAccepterCidrNil sets AccepterCidr to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetAccepterCidrNil() {
	o.AccepterCidr = nil
	o.NullFields = addNullField(o.NullFields, "AccepterCidr")
}

// GetAccepterVpcId returns the AccepterVpcId field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetAccepterVpcId() string {
	if o == nil || IsNil(o.AccepterVpcId) {
		var ret string
		return ret
	}
	return *o.AccepterVpcId
}

// GetAccepterVpcIdOk returns a tuple with the AccepterVpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetAccepterVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccepterVpcId) {
		return nil, false
	}

	return o.AccepterVpcId, true
}

// HasAccepterVpcId returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasAccepterVpcId() bool {
	if o != nil && !IsNil(o.AccepterVpcId) {
		return true
	}

	return false
}

// SetAccepterVpcId gets a reference to the given string and assigns it to the AccepterVpcId field.
func (o *VPCPeeringConnection) SetAccepterVpcId(v string) {
	o.AccepterVpcId = &v
	o.NullFields = removeNullField(o.NullFields, "AccepterVpcId")
}

// SetAccepterVpcIdNil sets AccepterVpcId to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetAccepterVpcIdNil() {
	o.AccepterVpcId = nil
	o.NullFields = addNullField(o.NullFields, "AccepterVpcId")
}

// GetCloudStatus returns the CloudStatus field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetCloudStatus() string {
	if o == nil || IsNil(o.CloudStatus) {
		var ret string
		return ret
	}
	return *o.CloudStatus
}

// GetCloudStatusOk returns a tuple with the CloudStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetCloudStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CloudStatus) {
		return nil, false
	}

	return o.CloudStatus, true
}

// HasCloudStatus returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasCloudStatus() bool {
	if o != nil && !IsNil(o.CloudStatus) {
		return true
	}

	return false
}

// SetCloudStatus gets a reference to the given string and assigns it to the CloudStatus field.
func (o *VPCPeeringConnection) SetCloudStatus(v string) {
	o.CloudStatus = &v
	o.NullFields = removeNullField(o.NullFields, "CloudStatus")
}

// SetCloudStatusNil sets CloudStatus to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetCloudStatusNil() {
	o.CloudStatus = nil
	o.NullFields = addNullField(o.NullFields, "CloudStatus")
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}

	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *VPCPeeringConnection) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
	o.NullFields = removeNullField(o.NullFields, "ExpirationTime")
}

// SetExpirationTimeNil sets ExpirationTime to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetExpirationTimeNil() {
	o.ExpirationTime = nil
	o.NullFields = addNullField(o.NullFields, "ExpirationTime")
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *VPCPeeringConnection) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *VPCPeeringConnection) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetLocalStatus returns the LocalStatus field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetLocalStatus() string {
	if o == nil || IsNil(o.LocalStatus) {
		var ret string
		return ret
	}
	return *o.LocalStatus
}

// GetLocalStatusOk returns a tuple with the LocalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetLocalStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LocalStatus) {
		return nil, false
	}

	return o.LocalStatus, true
}

// HasLocalStatus returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasLocalStatus() bool {
	if o != nil && !IsNil(o.LocalStatus) {
		return true
	}

	return false
}

// SetLocalStatus gets a reference to the given string and assigns it to the LocalStatus field.
func (o *VPCPeeringConnection) SetLocalStatus(v string) {
	o.LocalStatus = &v
	o.NullFields = removeNullField(o.NullFields, "LocalStatus")
}

// SetLocalStatusNil sets LocalStatus to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetLocalStatusNil() {
	o.LocalStatus = nil
	o.NullFields = addNullField(o.NullFields, "LocalStatus")
}

// GetName returns the Name field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VPCPeeringConnection) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetRequesterAccountId returns the RequesterAccountId field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetRequesterAccountId() string {
	if o == nil || IsNil(o.RequesterAccountId) {
		var ret string
		return ret
	}
	return *o.RequesterAccountId
}

// GetRequesterAccountIdOk returns a tuple with the RequesterAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetRequesterAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterAccountId) {
		return nil, false
	}

	return o.RequesterAccountId, true
}

// HasRequesterAccountId returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasRequesterAccountId() bool {
	if o != nil && !IsNil(o.RequesterAccountId) {
		return true
	}

	return false
}

// SetRequesterAccountId gets a reference to the given string and assigns it to the RequesterAccountId field.
func (o *VPCPeeringConnection) SetRequesterAccountId(v string) {
	o.RequesterAccountId = &v
	o.NullFields = removeNullField(o.NullFields, "RequesterAccountId")
}

// SetRequesterAccountIdNil sets RequesterAccountId to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetRequesterAccountIdNil() {
	o.RequesterAccountId = nil
	o.NullFields = addNullField(o.NullFields, "RequesterAccountId")
}

// GetRequesterCidr returns the RequesterCidr field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetRequesterCidr() string {
	if o == nil || IsNil(o.RequesterCidr) {
		var ret string
		return ret
	}
	return *o.RequesterCidr
}

// GetRequesterCidrOk returns a tuple with the RequesterCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetRequesterCidrOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterCidr) {
		return nil, false
	}

	return o.RequesterCidr, true
}

// HasRequesterCidr returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasRequesterCidr() bool {
	if o != nil && !IsNil(o.RequesterCidr) {
		return true
	}

	return false
}

// SetRequesterCidr gets a reference to the given string and assigns it to the RequesterCidr field.
func (o *VPCPeeringConnection) SetRequesterCidr(v string) {
	o.RequesterCidr = &v
	o.NullFields = removeNullField(o.NullFields, "RequesterCidr")
}

// SetRequesterCidrNil sets RequesterCidr to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetRequesterCidrNil() {
	o.RequesterCidr = nil
	o.NullFields = addNullField(o.NullFields, "RequesterCidr")
}

// GetRequesterVpcId returns the RequesterVpcId field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetRequesterVpcId() string {
	if o == nil || IsNil(o.RequesterVpcId) {
		var ret string
		return ret
	}
	return *o.RequesterVpcId
}

// GetRequesterVpcIdOk returns a tuple with the RequesterVpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetRequesterVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterVpcId) {
		return nil, false
	}

	return o.RequesterVpcId, true
}

// HasRequesterVpcId returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasRequesterVpcId() bool {
	if o != nil && !IsNil(o.RequesterVpcId) {
		return true
	}

	return false
}

// SetRequesterVpcId gets a reference to the given string and assigns it to the RequesterVpcId field.
func (o *VPCPeeringConnection) SetRequesterVpcId(v string) {
	o.RequesterVpcId = &v
	o.NullFields = removeNullField(o.NullFields, "RequesterVpcId")
}

// SetRequesterVpcIdNil sets RequesterVpcId to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetRequesterVpcIdNil() {
	o.RequesterVpcId = nil
	o.NullFields = addNullField(o.NullFields, "RequesterVpcId")
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise
func (o *VPCPeeringConnection) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VPCPeeringConnection) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}

	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *VPCPeeringConnection) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *VPCPeeringConnection) SetStatusMessage(v string) {
	o.StatusMessage = &v
	o.NullFields = removeNullField(o.NullFields, "StatusMessage")
}

// SetStatusMessageNil sets StatusMessage to an explicit JSON null when marshaled.
func (o *VPCPeeringConnection) SetStatusMessageNil() {
	o.StatusMessage = nil
	o.NullFields = addNullField(o.NullFields, "StatusMessage")
}
