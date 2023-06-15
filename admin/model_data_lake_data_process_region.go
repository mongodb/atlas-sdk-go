// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DataLakeDataProcessRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeDataProcessRegion{}

// DataLakeDataProcessRegion Information about the cloud provider region to which the data lake routes client connections. MongoDB Cloud supports AWS only.
type DataLakeDataProcessRegion struct {
	// Name of the cloud service that hosts the data lake's data stores.
	CloudProvider string `json:"cloudProvider"`
	// Atlas Data Lake Regions.
	Region string `json:"region"`
}

// NewDataLakeDataProcessRegion instantiates a new DataLakeDataProcessRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeDataProcessRegion(cloudProvider string, region string) *DataLakeDataProcessRegion {
	this := DataLakeDataProcessRegion{}
	this.CloudProvider = cloudProvider
	this.Region = region
	return &this
}

// NewDataLakeDataProcessRegionWithDefaults instantiates a new DataLakeDataProcessRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeDataProcessRegionWithDefaults() *DataLakeDataProcessRegion {
	this := DataLakeDataProcessRegion{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *DataLakeDataProcessRegion) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *DataLakeDataProcessRegion) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *DataLakeDataProcessRegion) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetRegion returns the Region field value
func (o *DataLakeDataProcessRegion) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *DataLakeDataProcessRegion) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *DataLakeDataProcessRegion) SetRegion(v string) {
	o.Region = v
}

func (o DataLakeDataProcessRegion) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DataLakeDataProcessRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloudProvider"] = o.CloudProvider
	toSerialize["region"] = o.Region
	return toSerialize, nil
}

type NullableDataLakeDataProcessRegion struct {
	value *DataLakeDataProcessRegion
	isSet bool
}

func (v NullableDataLakeDataProcessRegion) Get() *DataLakeDataProcessRegion {
	return v.value
}

func (v *NullableDataLakeDataProcessRegion) Set(val *DataLakeDataProcessRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeDataProcessRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeDataProcessRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeDataProcessRegion(val *DataLakeDataProcessRegion) *NullableDataLakeDataProcessRegion {
	return &NullableDataLakeDataProcessRegion{value: val, isSet: true}
}

func (v NullableDataLakeDataProcessRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeDataProcessRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
