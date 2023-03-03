/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the GCPCloudProviderContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GCPCloudProviderContainer{}

// GCPCloudProviderContainer Collection of settings that configures the network container for a virtual private connection on Amazon Web Services.
type GCPCloudProviderContainer struct {
	// IP addresses expressed in Classless Inter-Domain Routing (CIDR) notation that MongoDB Cloud uses for the network peering containers in your project. MongoDB Cloud assigns all of the project's clusters deployed to this cloud provider an IP address from this range. MongoDB Cloud locks this value if an M10 or greater cluster or a network peering connection exists in this project.  These CIDR blocks must fall within the ranges reserved per RFC 1918. GCP further limits the block to a lower bound of the `/18` range.  To modify the CIDR block, the target project cannot have:  - Any M10 or greater clusters - Any other VPC peering connections   You can also create a new project and create a network peering connection to set the desired MongoDB Cloud network peering container CIDR block for that project. MongoDB Cloud limits the number of MongoDB nodes per network peering connection based on the CIDR block and the region selected for the project.   **Example:** A project in an Google Cloud (GCP) region supporting three availability zones and an MongoDB CIDR network peering container block of limit of `/24` equals 27 three-node replica sets.
	AtlasCidrBlock string `json:"atlasCidrBlock"`
	// Unique string that identifies the GCP project in which MongoDB Cloud clusters in this network peering container exist. The response returns **null** if no clusters exist in this network peering container.
	GcpProjectId *string `json:"gcpProjectId,omitempty"`
	// Human-readable label that identifies the network in which MongoDB Cloud clusters in this network peering container exist. MongoDB Cloud returns **null** if no clusters exist in this network peering container.
	NetworkName *string `json:"networkName,omitempty"`
	// List of GCP regions to which you want to deploy this MongoDB Cloud network peering container.  In this MongoDB Cloud project, you can deploy clusters only to the GCP regions in this list. To deploy MongoDB Cloud clusters to other GCP regions, create additional projects.
	Regions []string `json:"regions,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering container.
	Id *string `json:"id,omitempty"`
	// Cloud service provider that serves the requested network peering containers.
	ProviderName *string `json:"providerName,omitempty"`
	// Flag that indicates whether MongoDB Cloud clusters exist in the specified network peering container.
	Provisioned *bool `json:"provisioned,omitempty"`
}

// NewGCPCloudProviderContainer instantiates a new GCPCloudProviderContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPCloudProviderContainer(atlasCidrBlock string) *GCPCloudProviderContainer {
	this := GCPCloudProviderContainer{}
	this.AtlasCidrBlock = atlasCidrBlock
	return &this
}

// NewGCPCloudProviderContainerWithDefaults instantiates a new GCPCloudProviderContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPCloudProviderContainerWithDefaults() *GCPCloudProviderContainer {
	this := GCPCloudProviderContainer{}
	return &this
}

// GetAtlasCidrBlock returns the AtlasCidrBlock field value
func (o *GCPCloudProviderContainer) GetAtlasCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AtlasCidrBlock
}

// GetAtlasCidrBlockOk returns a tuple with the AtlasCidrBlock field value
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainer) GetAtlasCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AtlasCidrBlock, true
}

// SetAtlasCidrBlock sets field value
func (o *GCPCloudProviderContainer) SetAtlasCidrBlock(v string) {
	o.AtlasCidrBlock = v
}

// GetGcpProjectId returns the GcpProjectId field value if set, zero value otherwise.
func (o *GCPCloudProviderContainer) GetGcpProjectId() string {
	if o == nil || IsNil(o.GcpProjectId) {
		var ret string
		return ret
	}
	return *o.GcpProjectId
}

// GetGcpProjectIdOk returns a tuple with the GcpProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainer) GetGcpProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectId) {
		return nil, false
	}
	return o.GcpProjectId, true
}

// HasGcpProjectId returns a boolean if a field has been set.
func (o *GCPCloudProviderContainer) HasGcpProjectId() bool {
	if o != nil && !IsNil(o.GcpProjectId) {
		return true
	}

	return false
}

// SetGcpProjectId gets a reference to the given string and assigns it to the GcpProjectId field.
func (o *GCPCloudProviderContainer) SetGcpProjectId(v string) {
	o.GcpProjectId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *GCPCloudProviderContainer) GetNetworkName() string {
	if o == nil || IsNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainer) GetNetworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkName) {
		return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *GCPCloudProviderContainer) HasNetworkName() bool {
	if o != nil && !IsNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *GCPCloudProviderContainer) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *GCPCloudProviderContainer) GetRegions() []string {
	if o == nil || IsNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainer) GetRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *GCPCloudProviderContainer) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *GCPCloudProviderContainer) SetRegions(v []string) {
	o.Regions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GCPCloudProviderContainer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GCPCloudProviderContainer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GCPCloudProviderContainer) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *GCPCloudProviderContainer) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainer) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *GCPCloudProviderContainer) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *GCPCloudProviderContainer) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *GCPCloudProviderContainer) GetProvisioned() bool {
	if o == nil || IsNil(o.Provisioned) {
		var ret bool
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPCloudProviderContainer) GetProvisionedOk() (*bool, bool) {
	if o == nil || IsNil(o.Provisioned) {
		return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *GCPCloudProviderContainer) HasProvisioned() bool {
	if o != nil && !IsNil(o.Provisioned) {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given bool and assigns it to the Provisioned field.
func (o *GCPCloudProviderContainer) SetProvisioned(v bool) {
	o.Provisioned = &v
}

func (o GCPCloudProviderContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GCPCloudProviderContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["atlasCidrBlock"] = o.AtlasCidrBlock
	// skip: gcpProjectId is readOnly
	// skip: networkName is readOnly
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	// skip: id is readOnly
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	// skip: provisioned is readOnly
	return toSerialize, nil
}

type NullableGCPCloudProviderContainer struct {
	value *GCPCloudProviderContainer
	isSet bool
}

func (v NullableGCPCloudProviderContainer) Get() *GCPCloudProviderContainer {
	return v.value
}

func (v *NullableGCPCloudProviderContainer) Set(val *GCPCloudProviderContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPCloudProviderContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPCloudProviderContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPCloudProviderContainer(val *GCPCloudProviderContainer) *NullableGCPCloudProviderContainer {
	return &NullableGCPCloudProviderContainer{value: val, isSet: true}
}

func (v NullableGCPCloudProviderContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPCloudProviderContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


