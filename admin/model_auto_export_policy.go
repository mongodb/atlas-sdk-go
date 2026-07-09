// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AutoExportPolicy Policy for automatically exporting Cloud Backup Snapshots.
type AutoExportPolicy struct {
	// Unique 24-hexadecimal character string that identifies the Export Bucket.
	ExportBucketId *string `json:"exportBucketId,omitempty"`
	// Human-readable label that indicates the rate at which the export policy item occurs.
	FrequencyType *string `json:"frequencyType,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AutoExportPolicy) MarshalJSON() ([]byte, error) {
	type noMethod AutoExportPolicy
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAutoExportPolicy instantiates a new AutoExportPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoExportPolicy() *AutoExportPolicy {
	this := AutoExportPolicy{}
	return &this
}

// NewAutoExportPolicyWithDefaults instantiates a new AutoExportPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoExportPolicyWithDefaults() *AutoExportPolicy {
	this := AutoExportPolicy{}
	return &this
}

// GetExportBucketId returns the ExportBucketId field value if set, zero value otherwise
func (o *AutoExportPolicy) GetExportBucketId() string {
	if o == nil || IsNil(o.ExportBucketId) {
		var ret string
		return ret
	}
	return *o.ExportBucketId
}

// GetExportBucketIdOk returns a tuple with the ExportBucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoExportPolicy) GetExportBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportBucketId) {
		return nil, false
	}

	return o.ExportBucketId, true
}

// HasExportBucketId returns a boolean if a field has been set.
func (o *AutoExportPolicy) HasExportBucketId() bool {
	if o != nil && !IsNil(o.ExportBucketId) {
		return true
	}

	return false
}

// SetExportBucketId gets a reference to the given string and assigns it to the ExportBucketId field.
func (o *AutoExportPolicy) SetExportBucketId(v string) {
	o.ExportBucketId = &v
}

// SetExportBucketIdNil sets ExportBucketId to an explicit JSON null when marshaled.
func (o *AutoExportPolicy) SetExportBucketIdNil() {
	o.ExportBucketId = nil
	o.NullFields = append(o.NullFields, "ExportBucketId")
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise
func (o *AutoExportPolicy) GetFrequencyType() string {
	if o == nil || IsNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoExportPolicy) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyType) {
		return nil, false
	}

	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *AutoExportPolicy) HasFrequencyType() bool {
	if o != nil && !IsNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *AutoExportPolicy) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

// SetFrequencyTypeNil sets FrequencyType to an explicit JSON null when marshaled.
func (o *AutoExportPolicy) SetFrequencyTypeNil() {
	o.FrequencyType = nil
	o.NullFields = append(o.NullFields, "FrequencyType")
}
