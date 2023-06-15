// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DataLakeS3StoreSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeS3StoreSettings{}

// DataLakeS3StoreSettings struct for DataLakeS3StoreSettings
type DataLakeS3StoreSettings struct {
	// Collection of AWS S3 [storage classes](https://aws.amazon.com/s3/storage-classes/). Atlas Data Lake includes the files in these storage classes in the query results.
	AdditionalStorageClasses []string `json:"additionalStorageClasses,omitempty"`
	// Human-readable label that identifies the AWS S3 bucket. This label must exactly match the name of an S3 bucket that the data lake can access with the configured AWS Identity and Access Management (IAM) credentials.
	Bucket *string `json:"bucket,omitempty"`
	// The delimiter that separates **databases.[n].collections.[n].dataSources.[n].path** segments in the data store. MongoDB Cloud uses the delimiter to efficiently traverse S3 buckets with a hierarchical directory structure. You can specify any character supported by the S3 object keys as the delimiter. For example, you can specify an underscore (_) or a plus sign (+) or multiple characters, such as double underscores (__) as the delimiter. If omitted, defaults to `/`.
	Delimiter *string `json:"delimiter,omitempty"`
	// Flag that indicates whether to use S3 tags on the files in the given path as additional partition attributes. If set to `true`, data lake adds the S3 tags as additional partition attributes and adds new top-level BSON elements associating each tag to each document.
	IncludeTags *bool `json:"includeTags,omitempty"`
	// Prefix that MongoDB Cloud applies when searching for files in the S3 bucket. The data store prepends the value of prefix to the **databases.[n].collections.[n].dataSources.[n].path** to create the full path for files to ingest. If omitted, MongoDB Cloud searches all files from the root of the S3 bucket.
	Prefix *string `json:"prefix,omitempty"`
	// Flag that indicates whether the bucket is public. If set to `true`, MongoDB Cloud doesn't use the configured AWS Identity and Access Management (IAM) role to access the S3 bucket. If set to `false`, the configured AWS IAM role must include permissions to access the S3 bucket.
	Public *bool `json:"public,omitempty"`
	//  Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts.
	Region *string `json:"region,omitempty"`
	// Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an `M10` or higher cluster.
	Name     *string `json:"name,omitempty"`
	Provider string  `json:"provider"`
}

// NewDataLakeS3StoreSettings instantiates a new DataLakeS3StoreSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeS3StoreSettings(provider string) *DataLakeS3StoreSettings {
	this := DataLakeS3StoreSettings{}
	var includeTags bool = false
	this.IncludeTags = &includeTags
	var public bool = false
	this.Public = &public
	this.Provider = provider
	return &this
}

// NewDataLakeS3StoreSettingsWithDefaults instantiates a new DataLakeS3StoreSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeS3StoreSettingsWithDefaults() *DataLakeS3StoreSettings {
	this := DataLakeS3StoreSettings{}
	var includeTags bool = false
	this.IncludeTags = &includeTags
	var public bool = false
	this.Public = &public
	return &this
}

// GetAdditionalStorageClasses returns the AdditionalStorageClasses field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetAdditionalStorageClasses() []string {
	if o == nil || IsNil(o.AdditionalStorageClasses) {
		var ret []string
		return ret
	}
	return o.AdditionalStorageClasses
}

// GetAdditionalStorageClassesOk returns a tuple with the AdditionalStorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetAdditionalStorageClassesOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalStorageClasses) {
		return nil, false
	}
	return o.AdditionalStorageClasses, true
}

// HasAdditionalStorageClasses returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasAdditionalStorageClasses() bool {
	if o != nil && !IsNil(o.AdditionalStorageClasses) {
		return true
	}

	return false
}

// SetAdditionalStorageClasses gets a reference to the given []string and assigns it to the AdditionalStorageClasses field.
func (o *DataLakeS3StoreSettings) SetAdditionalStorageClasses(v []string) {
	o.AdditionalStorageClasses = v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *DataLakeS3StoreSettings) SetBucket(v string) {
	o.Bucket = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetDelimiter() string {
	if o == nil || IsNil(o.Delimiter) {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetDelimiterOk() (*string, bool) {
	if o == nil || IsNil(o.Delimiter) {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasDelimiter() bool {
	if o != nil && !IsNil(o.Delimiter) {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *DataLakeS3StoreSettings) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetIncludeTags() bool {
	if o == nil || IsNil(o.IncludeTags) {
		var ret bool
		return ret
	}
	return *o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetIncludeTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTags) {
		return nil, false
	}
	return o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasIncludeTags() bool {
	if o != nil && !IsNil(o.IncludeTags) {
		return true
	}

	return false
}

// SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.
func (o *DataLakeS3StoreSettings) SetIncludeTags(v bool) {
	o.IncludeTags = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *DataLakeS3StoreSettings) SetPrefix(v string) {
	o.Prefix = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *DataLakeS3StoreSettings) SetPublic(v bool) {
	o.Public = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DataLakeS3StoreSettings) SetRegion(v string) {
	o.Region = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataLakeS3StoreSettings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataLakeS3StoreSettings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataLakeS3StoreSettings) SetName(v string) {
	o.Name = &v
}

// GetProvider returns the Provider field value
func (o *DataLakeS3StoreSettings) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *DataLakeS3StoreSettings) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *DataLakeS3StoreSettings) SetProvider(v string) {
	o.Provider = v
}

func (o DataLakeS3StoreSettings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeS3StoreSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalStorageClasses) {
		toSerialize["additionalStorageClasses"] = o.AdditionalStorageClasses
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Delimiter) {
		toSerialize["delimiter"] = o.Delimiter
	}
	if !IsNil(o.IncludeTags) {
		toSerialize["includeTags"] = o.IncludeTags
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["provider"] = o.Provider
	return toSerialize, nil
}

type NullableDataLakeS3StoreSettings struct {
	value *DataLakeS3StoreSettings
	isSet bool
}

func (v NullableDataLakeS3StoreSettings) Get() *DataLakeS3StoreSettings {
	return v.value
}

func (v *NullableDataLakeS3StoreSettings) Set(val *DataLakeS3StoreSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeS3StoreSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeS3StoreSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeS3StoreSettings(val *DataLakeS3StoreSettings) *NullableDataLakeS3StoreSettings {
	return &NullableDataLakeS3StoreSettings{value: val, isSet: true}
}

func (v NullableDataLakeS3StoreSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeS3StoreSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
