// Code based on the AtlasAPI V2 OpenAPI file

package admin

// BaseCloudProviderAccessFeatureUsageFeatureId struct for BaseCloudProviderAccessFeatureUsageFeatureId
type BaseCloudProviderAccessFeatureUsageFeatureId struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the data lake.
	Name *string `json:"name,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Atlas log integration configuration.
	// Read only field.
	ConfigId *string `json:"configId,omitempty"`
	// Human-readable label that identifies the sink, such as an S3/GCS bucket or Azure container name.
	// Read only field.
	SinkName *string `json:"sinkName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the AWS S3 bucket to which you export your snapshots.
	// Read only field.
	ExportBucketId *string `json:"exportBucketId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the object storage private endpoint.
	// Read only field.
	PrivateEndpointId *string `json:"privateEndpointId,omitempty"`
	// Name of the AWS S3 bucket to which your logs will be exported to.
	// Read only field.
	BucketName *string `json:"bucketName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o BaseCloudProviderAccessFeatureUsageFeatureId) MarshalJSON() ([]byte, error) {
	type noMethod BaseCloudProviderAccessFeatureUsageFeatureId
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewBaseCloudProviderAccessFeatureUsageFeatureId instantiates a new BaseCloudProviderAccessFeatureUsageFeatureId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCloudProviderAccessFeatureUsageFeatureId() *BaseCloudProviderAccessFeatureUsageFeatureId {
	this := BaseCloudProviderAccessFeatureUsageFeatureId{}
	return &this
}

// NewBaseCloudProviderAccessFeatureUsageFeatureIdWithDefaults instantiates a new BaseCloudProviderAccessFeatureUsageFeatureId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseCloudProviderAccessFeatureUsageFeatureIdWithDefaults() *BaseCloudProviderAccessFeatureUsageFeatureId {
	this := BaseCloudProviderAccessFeatureUsageFeatureId{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetName returns the Name field value if set, zero value otherwise
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetName(v string) {
	o.Name = &v
	o.NullFields = removeNullField(o.NullFields, "Name")
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetNameNil() {
	o.Name = nil
	o.NullFields = addNullField(o.NullFields, "Name")
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId) {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigId) {
		return nil, false
	}

	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasConfigId() bool {
	if o != nil && !IsNil(o.ConfigId) {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetConfigId(v string) {
	o.ConfigId = &v
	o.NullFields = removeNullField(o.NullFields, "ConfigId")
}

// SetConfigIdNil sets ConfigId to an explicit JSON null when marshaled.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetConfigIdNil() {
	o.ConfigId = nil
	o.NullFields = addNullField(o.NullFields, "ConfigId")
}

// GetSinkName returns the SinkName field value if set, zero value otherwise
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetSinkName() string {
	if o == nil || IsNil(o.SinkName) {
		var ret string
		return ret
	}
	return *o.SinkName
}

// GetSinkNameOk returns a tuple with the SinkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetSinkNameOk() (*string, bool) {
	if o == nil || IsNil(o.SinkName) {
		return nil, false
	}

	return o.SinkName, true
}

// HasSinkName returns a boolean if a field has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasSinkName() bool {
	if o != nil && !IsNil(o.SinkName) {
		return true
	}

	return false
}

// SetSinkName gets a reference to the given string and assigns it to the SinkName field.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetSinkName(v string) {
	o.SinkName = &v
	o.NullFields = removeNullField(o.NullFields, "SinkName")
}

// SetSinkNameNil sets SinkName to an explicit JSON null when marshaled.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetSinkNameNil() {
	o.SinkName = nil
	o.NullFields = addNullField(o.NullFields, "SinkName")
}

// GetExportBucketId returns the ExportBucketId field value if set, zero value otherwise
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetExportBucketId() string {
	if o == nil || IsNil(o.ExportBucketId) {
		var ret string
		return ret
	}
	return *o.ExportBucketId
}

// GetExportBucketIdOk returns a tuple with the ExportBucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetExportBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportBucketId) {
		return nil, false
	}

	return o.ExportBucketId, true
}

// HasExportBucketId returns a boolean if a field has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasExportBucketId() bool {
	if o != nil && !IsNil(o.ExportBucketId) {
		return true
	}

	return false
}

// SetExportBucketId gets a reference to the given string and assigns it to the ExportBucketId field.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetExportBucketId(v string) {
	o.ExportBucketId = &v
	o.NullFields = removeNullField(o.NullFields, "ExportBucketId")
}

// SetExportBucketIdNil sets ExportBucketId to an explicit JSON null when marshaled.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetExportBucketIdNil() {
	o.ExportBucketId = nil
	o.NullFields = addNullField(o.NullFields, "ExportBucketId")
}

// GetPrivateEndpointId returns the PrivateEndpointId field value if set, zero value otherwise
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetPrivateEndpointId() string {
	if o == nil || IsNil(o.PrivateEndpointId) {
		var ret string
		return ret
	}
	return *o.PrivateEndpointId
}

// GetPrivateEndpointIdOk returns a tuple with the PrivateEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetPrivateEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateEndpointId) {
		return nil, false
	}

	return o.PrivateEndpointId, true
}

// HasPrivateEndpointId returns a boolean if a field has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasPrivateEndpointId() bool {
	if o != nil && !IsNil(o.PrivateEndpointId) {
		return true
	}

	return false
}

// SetPrivateEndpointId gets a reference to the given string and assigns it to the PrivateEndpointId field.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetPrivateEndpointId(v string) {
	o.PrivateEndpointId = &v
	o.NullFields = removeNullField(o.NullFields, "PrivateEndpointId")
}

// SetPrivateEndpointIdNil sets PrivateEndpointId to an explicit JSON null when marshaled.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetPrivateEndpointIdNil() {
	o.PrivateEndpointId = nil
	o.NullFields = addNullField(o.NullFields, "PrivateEndpointId")
}

// GetBucketName returns the BucketName field value if set, zero value otherwise
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}

	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetBucketName(v string) {
	o.BucketName = &v
	o.NullFields = removeNullField(o.NullFields, "BucketName")
}

// SetBucketNameNil sets BucketName to an explicit JSON null when marshaled.
func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetBucketNameNil() {
	o.BucketName = nil
	o.NullFields = addNullField(o.NullFields, "BucketName")
}
