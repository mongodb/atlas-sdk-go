// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AWSCloudProviderContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSCloudProviderContainer{}

// AWSCloudProviderContainer Collection of settings that configures the network container for a virtual private connection on Amazon Web Services.
type AWSCloudProviderContainer struct {
	// IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project's clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. AWS and Azure further limit the block to between the `/24` and  `/21` ranges.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Amazon Web Services (AWS) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of `/24` equals 27 three-node replica sets.
	AtlasCidrBlock *string `json:"atlasCidrBlock,omitempty"`
	// Geographic area that Amazon Web Services (AWS) defines to which MongoDB Cloud deployed this network peering container.
	RegionName string `json:"regionName"`
	// Unique string that identifies the MongoDB Cloud VPC on AWS.
	VpcId *string `json:"vpcId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering container.
	Id *string `json:"id,omitempty"`
	// Cloud service provider that serves the requested network peering containers.
	ProviderName *string `json:"providerName,omitempty"`
	// Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container.
	Provisioned *bool `json:"provisioned,omitempty"`
}

// NewAWSCloudProviderContainer instantiates a new AWSCloudProviderContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSCloudProviderContainer(regionName string) *AWSCloudProviderContainer {
	this := AWSCloudProviderContainer{}
	this.RegionName = regionName
	return &this
}

// NewAWSCloudProviderContainerWithDefaults instantiates a new AWSCloudProviderContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSCloudProviderContainerWithDefaults() *AWSCloudProviderContainer {
	this := AWSCloudProviderContainer{}
	return &this
}

// GetAtlasCidrBlock returns the AtlasCidrBlock field value if set, zero value otherwise.
func (o *AWSCloudProviderContainer) GetAtlasCidrBlock() string {
	if o == nil || IsNil(o.AtlasCidrBlock) {
		var ret string
		return ret
	}
	return *o.AtlasCidrBlock
}

// GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool) {
	if o == nil || IsNil(o.AtlasCidrBlock) {
		return nil, false
	}
	return o.AtlasCidrBlock, true
}

// HasAtlasCidrBlock returns a boolean if a field has been set.
func (o *AWSCloudProviderContainer) HasAtlasCidrBlock() bool {
	if o != nil && !IsNil(o.AtlasCidrBlock) {
		return true
	}

	return false
}

// SetAtlasCidrBlock gets a reference to the given string and assigns it to the AtlasCidrBlock field.
func (o *AWSCloudProviderContainer) SetAtlasCidrBlock(v string) {
	o.AtlasCidrBlock = &v
}

// GetRegionName returns the RegionName field value
func (o *AWSCloudProviderContainer) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderContainer) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *AWSCloudProviderContainer) SetRegionName(v string) {
	o.RegionName = v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *AWSCloudProviderContainer) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderContainer) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *AWSCloudProviderContainer) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *AWSCloudProviderContainer) SetVpcId(v string) {
	o.VpcId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AWSCloudProviderContainer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderContainer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AWSCloudProviderContainer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AWSCloudProviderContainer) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *AWSCloudProviderContainer) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderContainer) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *AWSCloudProviderContainer) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *AWSCloudProviderContainer) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *AWSCloudProviderContainer) GetProvisioned() bool {
	if o == nil || IsNil(o.Provisioned) {
		var ret bool
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudProviderContainer) GetProvisionedOk() (*bool, bool) {
	if o == nil || IsNil(o.Provisioned) {
		return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *AWSCloudProviderContainer) HasProvisioned() bool {
	if o != nil && !IsNil(o.Provisioned) {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given bool and assigns it to the Provisioned field.
func (o *AWSCloudProviderContainer) SetProvisioned(v bool) {
	o.Provisioned = &v
}

func (o AWSCloudProviderContainer) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AWSCloudProviderContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AtlasCidrBlock) {
		toSerialize["atlasCidrBlock"] = o.AtlasCidrBlock
	}
	toSerialize["regionName"] = o.RegionName
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	return toSerialize, nil
}

type NullableAWSCloudProviderContainer struct {
	value *AWSCloudProviderContainer
	isSet bool
}

func (v NullableAWSCloudProviderContainer) Get() *AWSCloudProviderContainer {
	return v.value
}

func (v *NullableAWSCloudProviderContainer) Set(val *AWSCloudProviderContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSCloudProviderContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSCloudProviderContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSCloudProviderContainer(val *AWSCloudProviderContainer) *NullableAWSCloudProviderContainer {
	return &NullableAWSCloudProviderContainer{value: val, isSet: true}
}

func (v NullableAWSCloudProviderContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSCloudProviderContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
