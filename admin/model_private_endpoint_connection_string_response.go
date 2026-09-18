// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// PrivateEndpointConnectionStringResponse Private endpoint connection string for one cluster.
type PrivateEndpointConnectionStringResponse struct {
	// Private endpoint-aware connection string that uses the `mongodb://` protocol to connect to this cluster through the selected private endpoints. MongoDB Cloud returns this value when `status` is `AVAILABLE`. After you submit an update, MongoDB Cloud continues to return the existing connection string until the updated connection string becomes available.
	// Read only field.
	ConnectionString *string `json:"connectionString,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this private endpoint connection string. This value is the `connectionStringId` path parameter of this resource's endpoints.
	// Read only field.
	ConnectionStringId string `json:"connectionStringId"`
	// Date and time when MongoDB Cloud created this private endpoint connection string. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Error message that explains why this private endpoint connection string is in the `FAILED` state, including the remediation, for example when a selected endpoint was deleted mid-flight. Present only when `status` is `FAILED`.
	// Read only field.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Party that manages this private endpoint connection string. MongoDB Cloud returns `SELF_MANAGED` for connection strings that you create and maintain through this resource, and `IMPORTED` for records that preserve the private endpoint connection strings that this cluster exposed before migration. You can delete an imported record but not update it.
	// Read only field.
	ManagementType string `json:"managementType"`
	// Human-readable label of at most 64 characters that identifies this private endpoint connection string. You can change this label without changing the connection string.
	Name string `json:"name"`
	// Flag that indicates whether this private endpoint connection string uses the optimized load-balanced mode. This mode requires a sharded cluster and a private endpoint selection that satisfies the optimized connection string requirements.
	OptimizedModeEnabled *bool `json:"optimizedModeEnabled,omitempty"`
	// Set of private endpoints that this connection string covers. MongoDB Cloud stores this set as a snapshot of your selection and never adds newly created private endpoints to it. The set accepts at most one private endpoint per cloud provider region, and can include private endpoints that the cluster does not use yet. Order is not significant and MongoDB Cloud rejects duplicate values.
	PrivateEndpointIds []string `json:"privateEndpointIds"`
	// Flag that indicates whether this private endpoint connection string routes clients to a subset of the `mongos` processes of the cluster.
	SelectiveMongosEnabled *bool `json:"selectiveMongosEnabled,omitempty"`
	// Short generated pin of at most 64 characters, unique per project. Present only when `managementType` is `SELF_MANAGED`. A self-managed connection string embeds this value in its DNS names, and its replica set horizon name is `cs-` plus this value.
	// Read only field.
	Slug *string `json:"slug,omitempty"`
	// Private endpoint-aware connection string that uses the `mongodb+srv://` protocol to connect to this cluster through the selected private endpoints. MongoDB Cloud returns this value when `status` is `AVAILABLE`. After you submit an update, MongoDB Cloud continues to return the existing connection string until the updated connection string becomes available.
	// Read only field.
	SrvConnectionString *string `json:"srvConnectionString,omitempty"`
	// State of this private endpoint connection string. MongoDB Cloud returns `PENDING` while it reconfigures the cluster, `AVAILABLE` once the connection string is usable, and `FAILED` when creation cannot complete.
	// Read only field.
	Status string `json:"status"`
	// Date and time when MongoDB Cloud last updated this private endpoint connection string. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PrivateEndpointConnectionStringResponse) MarshalJSON() ([]byte, error) {
	type noMethod PrivateEndpointConnectionStringResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPrivateEndpointConnectionStringResponse instantiates a new PrivateEndpointConnectionStringResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateEndpointConnectionStringResponse(connectionStringId string, managementType string, name string, privateEndpointIds []string, status string) *PrivateEndpointConnectionStringResponse {
	this := PrivateEndpointConnectionStringResponse{}
	this.ConnectionStringId = connectionStringId
	this.ManagementType = managementType
	this.Name = name
	var optimizedModeEnabled bool = false
	this.OptimizedModeEnabled = &optimizedModeEnabled
	this.PrivateEndpointIds = privateEndpointIds
	var selectiveMongosEnabled bool = false
	this.SelectiveMongosEnabled = &selectiveMongosEnabled
	this.Status = status
	return &this
}

// NewPrivateEndpointConnectionStringResponseWithDefaults instantiates a new PrivateEndpointConnectionStringResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateEndpointConnectionStringResponseWithDefaults() *PrivateEndpointConnectionStringResponse {
	this := PrivateEndpointConnectionStringResponse{}
	var optimizedModeEnabled bool = false
	this.OptimizedModeEnabled = &optimizedModeEnabled
	var selectiveMongosEnabled bool = false
	this.SelectiveMongosEnabled = &selectiveMongosEnabled
	return &this
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetConnectionString() string {
	if o == nil || IsNil(o.ConnectionString) {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionString) {
		return nil, false
	}

	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasConnectionString() bool {
	if o != nil && !IsNil(o.ConnectionString) {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *PrivateEndpointConnectionStringResponse) SetConnectionString(v string) {
	o.ConnectionString = &v
	o.NullFields = removeNullField(o.NullFields, "ConnectionString")
}

// SetConnectionStringNil sets ConnectionString to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetConnectionStringNil() {
	o.ConnectionString = nil
	o.NullFields = addNullField(o.NullFields, "ConnectionString")
}

// GetConnectionStringId returns the ConnectionStringId field value
func (o *PrivateEndpointConnectionStringResponse) GetConnectionStringId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionStringId
}

// GetConnectionStringIdOk returns a tuple with the ConnectionStringId field value
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetConnectionStringIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionStringId, true
}

// SetConnectionStringId sets field value
func (o *PrivateEndpointConnectionStringResponse) SetConnectionStringId(v string) {
	o.ConnectionStringId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PrivateEndpointConnectionStringResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
	o.NullFields = removeNullField(o.NullFields, "CreatedAt")
}

// SetCreatedAtNil sets CreatedAt to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetCreatedAtNil() {
	o.CreatedAt = nil
	o.NullFields = addNullField(o.NullFields, "CreatedAt")
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PrivateEndpointConnectionStringResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
	o.NullFields = removeNullField(o.NullFields, "ErrorMessage")
}

// SetErrorMessageNil sets ErrorMessage to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetErrorMessageNil() {
	o.ErrorMessage = nil
	o.NullFields = addNullField(o.NullFields, "ErrorMessage")
}

// GetManagementType returns the ManagementType field value
func (o *PrivateEndpointConnectionStringResponse) GetManagementType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManagementType
}

// GetManagementTypeOk returns a tuple with the ManagementType field value
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetManagementTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementType, true
}

// SetManagementType sets field value
func (o *PrivateEndpointConnectionStringResponse) SetManagementType(v string) {
	o.ManagementType = v
}

// GetName returns the Name field value
func (o *PrivateEndpointConnectionStringResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateEndpointConnectionStringResponse) SetName(v string) {
	o.Name = v
}

// GetOptimizedModeEnabled returns the OptimizedModeEnabled field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetOptimizedModeEnabled() bool {
	if o == nil || IsNil(o.OptimizedModeEnabled) {
		var ret bool
		return ret
	}
	return *o.OptimizedModeEnabled
}

// GetOptimizedModeEnabledOk returns a tuple with the OptimizedModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetOptimizedModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OptimizedModeEnabled) {
		return nil, false
	}

	return o.OptimizedModeEnabled, true
}

// HasOptimizedModeEnabled returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasOptimizedModeEnabled() bool {
	if o != nil && !IsNil(o.OptimizedModeEnabled) {
		return true
	}

	return false
}

// SetOptimizedModeEnabled gets a reference to the given bool and assigns it to the OptimizedModeEnabled field.
func (o *PrivateEndpointConnectionStringResponse) SetOptimizedModeEnabled(v bool) {
	o.OptimizedModeEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "OptimizedModeEnabled")
}

// SetOptimizedModeEnabledNil sets OptimizedModeEnabled to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetOptimizedModeEnabledNil() {
	o.OptimizedModeEnabled = nil
	o.NullFields = addNullField(o.NullFields, "OptimizedModeEnabled")
}

// GetPrivateEndpointIds returns the PrivateEndpointIds field value
func (o *PrivateEndpointConnectionStringResponse) GetPrivateEndpointIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrivateEndpointIds
}

// GetPrivateEndpointIdsOk returns a tuple with the PrivateEndpointIds field value
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetPrivateEndpointIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateEndpointIds, true
}

// SetPrivateEndpointIds sets field value
func (o *PrivateEndpointConnectionStringResponse) SetPrivateEndpointIds(v []string) {
	o.PrivateEndpointIds = v
}

// GetSelectiveMongosEnabled returns the SelectiveMongosEnabled field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetSelectiveMongosEnabled() bool {
	if o == nil || IsNil(o.SelectiveMongosEnabled) {
		var ret bool
		return ret
	}
	return *o.SelectiveMongosEnabled
}

// GetSelectiveMongosEnabledOk returns a tuple with the SelectiveMongosEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetSelectiveMongosEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SelectiveMongosEnabled) {
		return nil, false
	}

	return o.SelectiveMongosEnabled, true
}

// HasSelectiveMongosEnabled returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasSelectiveMongosEnabled() bool {
	if o != nil && !IsNil(o.SelectiveMongosEnabled) {
		return true
	}

	return false
}

// SetSelectiveMongosEnabled gets a reference to the given bool and assigns it to the SelectiveMongosEnabled field.
func (o *PrivateEndpointConnectionStringResponse) SetSelectiveMongosEnabled(v bool) {
	o.SelectiveMongosEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "SelectiveMongosEnabled")
}

// SetSelectiveMongosEnabledNil sets SelectiveMongosEnabled to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetSelectiveMongosEnabledNil() {
	o.SelectiveMongosEnabled = nil
	o.NullFields = addNullField(o.NullFields, "SelectiveMongosEnabled")
}

// GetSlug returns the Slug field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}

	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PrivateEndpointConnectionStringResponse) SetSlug(v string) {
	o.Slug = &v
	o.NullFields = removeNullField(o.NullFields, "Slug")
}

// SetSlugNil sets Slug to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetSlugNil() {
	o.Slug = nil
	o.NullFields = addNullField(o.NullFields, "Slug")
}

// GetSrvConnectionString returns the SrvConnectionString field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetSrvConnectionString() string {
	if o == nil || IsNil(o.SrvConnectionString) {
		var ret string
		return ret
	}
	return *o.SrvConnectionString
}

// GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetSrvConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.SrvConnectionString) {
		return nil, false
	}

	return o.SrvConnectionString, true
}

// HasSrvConnectionString returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasSrvConnectionString() bool {
	if o != nil && !IsNil(o.SrvConnectionString) {
		return true
	}

	return false
}

// SetSrvConnectionString gets a reference to the given string and assigns it to the SrvConnectionString field.
func (o *PrivateEndpointConnectionStringResponse) SetSrvConnectionString(v string) {
	o.SrvConnectionString = &v
	o.NullFields = removeNullField(o.NullFields, "SrvConnectionString")
}

// SetSrvConnectionStringNil sets SrvConnectionString to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetSrvConnectionStringNil() {
	o.SrvConnectionString = nil
	o.NullFields = addNullField(o.NullFields, "SrvConnectionString")
}

// GetStatus returns the Status field value
func (o *PrivateEndpointConnectionStringResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PrivateEndpointConnectionStringResponse) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise
func (o *PrivateEndpointConnectionStringResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateEndpointConnectionStringResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}

	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PrivateEndpointConnectionStringResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PrivateEndpointConnectionStringResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
	o.NullFields = removeNullField(o.NullFields, "UpdatedAt")
}

// SetUpdatedAtNil sets UpdatedAt to an explicit JSON null when marshaled.
func (o *PrivateEndpointConnectionStringResponse) SetUpdatedAtNil() {
	o.UpdatedAt = nil
	o.NullFields = addNullField(o.NullFields, "UpdatedAt")
}
