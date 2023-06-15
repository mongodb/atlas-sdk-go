// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AWSPeerVpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSPeerVpc{}

// AWSPeerVpc Group of Network Peering connection settings.
type AWSPeerVpc struct {
	// Amazon Web Services (AWS) region where the Virtual Peering Connection (VPC) that you peered with the MongoDB Cloud VPC resides. The resource returns `null` if your VPC and the MongoDB Cloud VPC reside in the same region.
	AccepterRegionName string `json:"accepterRegionName"`
	// Unique twelve-digit string that identifies the Amazon Web Services (AWS) account that owns the VPC that you peered with the MongoDB Cloud VPC.
	AwsAccountId string `json:"awsAccountId"`
	// Unique string that identifies the peering connection on AWS.
	ConnectionId *string `json:"connectionId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection.
	ContainerId string `json:"containerId"`
	// Type of error that can be returned when requesting an Amazon Web Services (AWS) peering connection. The resource returns `null` if the request succeeded.
	ErrorStateName *string `json:"errorStateName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the network peering connection.
	Id *string `json:"id,omitempty"`
	// Cloud service provider that serves the requested network peering connection.
	ProviderName *string `json:"providerName,omitempty"`
	// Internet Protocol (IP) addresses expressed in Classless Inter-Domain Routing (CIDR) notation of the VPC's subnet that you want to peer with the MongoDB Cloud VPC.
	RouteTableCidrBlock string `json:"routeTableCidrBlock"`
	// State of the network peering connection at the time you made the request.
	StatusName *string `json:"statusName,omitempty"`
	// Unique string that identifies the VPC on Amazon Web Services (AWS) that you want to peer with the MongoDB Cloud VPC.
	VpcId string `json:"vpcId"`
}

// NewAWSPeerVpc instantiates a new AWSPeerVpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSPeerVpc(accepterRegionName string, awsAccountId string, containerId string, routeTableCidrBlock string, vpcId string) *AWSPeerVpc {
	this := AWSPeerVpc{}
	this.AccepterRegionName = accepterRegionName
	this.AwsAccountId = awsAccountId
	this.ContainerId = containerId
	this.RouteTableCidrBlock = routeTableCidrBlock
	this.VpcId = vpcId
	return &this
}

// NewAWSPeerVpcWithDefaults instantiates a new AWSPeerVpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSPeerVpcWithDefaults() *AWSPeerVpc {
	this := AWSPeerVpc{}
	return &this
}

// GetAccepterRegionName returns the AccepterRegionName field value
func (o *AWSPeerVpc) GetAccepterRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccepterRegionName
}

// GetAccepterRegionNameOk returns a tuple with the AccepterRegionName field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetAccepterRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccepterRegionName, true
}

// SetAccepterRegionName sets field value
func (o *AWSPeerVpc) SetAccepterRegionName(v string) {
	o.AccepterRegionName = v
}

// GetAwsAccountId returns the AwsAccountId field value
func (o *AWSPeerVpc) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value
func (o *AWSPeerVpc) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *AWSPeerVpc) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId) {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionId) {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *AWSPeerVpc) HasConnectionId() bool {
	if o != nil && !IsNil(o.ConnectionId) {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *AWSPeerVpc) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetContainerId returns the ContainerId field value
func (o *AWSPeerVpc) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *AWSPeerVpc) SetContainerId(v string) {
	o.ContainerId = v
}

// GetErrorStateName returns the ErrorStateName field value if set, zero value otherwise.
func (o *AWSPeerVpc) GetErrorStateName() string {
	if o == nil || IsNil(o.ErrorStateName) {
		var ret string
		return ret
	}
	return *o.ErrorStateName
}

// GetErrorStateNameOk returns a tuple with the ErrorStateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetErrorStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorStateName) {
		return nil, false
	}
	return o.ErrorStateName, true
}

// HasErrorStateName returns a boolean if a field has been set.
func (o *AWSPeerVpc) HasErrorStateName() bool {
	if o != nil && !IsNil(o.ErrorStateName) {
		return true
	}

	return false
}

// SetErrorStateName gets a reference to the given string and assigns it to the ErrorStateName field.
func (o *AWSPeerVpc) SetErrorStateName(v string) {
	o.ErrorStateName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AWSPeerVpc) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AWSPeerVpc) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AWSPeerVpc) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *AWSPeerVpc) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *AWSPeerVpc) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *AWSPeerVpc) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRouteTableCidrBlock returns the RouteTableCidrBlock field value
func (o *AWSPeerVpc) GetRouteTableCidrBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteTableCidrBlock
}

// GetRouteTableCidrBlockOk returns a tuple with the RouteTableCidrBlock field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetRouteTableCidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteTableCidrBlock, true
}

// SetRouteTableCidrBlock sets field value
func (o *AWSPeerVpc) SetRouteTableCidrBlock(v string) {
	o.RouteTableCidrBlock = v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *AWSPeerVpc) GetStatusName() string {
	if o == nil || IsNil(o.StatusName) {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetStatusNameOk() (*string, bool) {
	if o == nil || IsNil(o.StatusName) {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *AWSPeerVpc) HasStatusName() bool {
	if o != nil && !IsNil(o.StatusName) {
		return true
	}

	return false
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *AWSPeerVpc) SetStatusName(v string) {
	o.StatusName = &v
}

// GetVpcId returns the VpcId field value
func (o *AWSPeerVpc) GetVpcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value
// and a boolean to check if the value has been set.
func (o *AWSPeerVpc) GetVpcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpcId, true
}

// SetVpcId sets field value
func (o *AWSPeerVpc) SetVpcId(v string) {
	o.VpcId = v
}

func (o AWSPeerVpc) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AWSPeerVpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accepterRegionName"] = o.AccepterRegionName
	toSerialize["awsAccountId"] = o.AwsAccountId
	toSerialize["containerId"] = o.ContainerId
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	toSerialize["routeTableCidrBlock"] = o.RouteTableCidrBlock
	toSerialize["vpcId"] = o.VpcId
	return toSerialize, nil
}

type NullableAWSPeerVpc struct {
	value *AWSPeerVpc
	isSet bool
}

func (v NullableAWSPeerVpc) Get() *AWSPeerVpc {
	return v.value
}

func (v *NullableAWSPeerVpc) Set(val *AWSPeerVpc) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSPeerVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSPeerVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSPeerVpc(val *AWSPeerVpc) *NullableAWSPeerVpc {
	return &NullableAWSPeerVpc{value: val, isSet: true}
}

func (v NullableAWSPeerVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSPeerVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
