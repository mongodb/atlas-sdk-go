// Code based on the AtlasAPI V2 OpenAPI file

package admin

// CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId Identifying characteristics about the Atlas log integration linked to this cloud provider access role.
type CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId struct {
	// Unique 24-hexadecimal digit string that identifies the Atlas log integration configuration.
	// Read only field.
	ConfigId *string `json:"configId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies your project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the sink, such as an S3/GCS bucket or Azure container name.
	// Read only field.
	SinkName *string `json:"sinkName,omitempty"`
}

// NewCloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId instantiates a new CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId() *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId {
	this := CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId{}
	return &this
}

// NewCloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureIdWithDefaults instantiates a new CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureIdWithDefaults() *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId {
	this := CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId{}
	return &this
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId) {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) GetConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigId) {
		return nil, false
	}

	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) HasConfigId() bool {
	if o != nil && !IsNil(o.ConfigId) {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) SetGroupId(v string) {
	o.GroupId = &v
}

// GetSinkName returns the SinkName field value if set, zero value otherwise
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) GetSinkName() string {
	if o == nil || IsNil(o.SinkName) {
		var ret string
		return ret
	}
	return *o.SinkName
}

// GetSinkNameOk returns a tuple with the SinkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) GetSinkNameOk() (*string, bool) {
	if o == nil || IsNil(o.SinkName) {
		return nil, false
	}

	return o.SinkName, true
}

// HasSinkName returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) HasSinkName() bool {
	if o != nil && !IsNil(o.SinkName) {
		return true
	}

	return false
}

// SetSinkName gets a reference to the given string and assigns it to the SinkName field.
func (o *CloudProviderAccessFeatureUsageAtlasLogIntegrationFeatureId) SetSinkName(v string) {
	o.SinkName = &v
}
