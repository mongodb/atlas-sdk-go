// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// CloudProviderAccessFeatureUsage - MongoDB Cloud features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
type CloudProviderAccessFeatureUsage struct {
	CloudProviderAccessDataLakeFeatureUsage         *CloudProviderAccessDataLakeFeatureUsage
	CloudProviderAccessEncryptionAtRestFeatureUsage *CloudProviderAccessEncryptionAtRestFeatureUsage
	CloudProviderAccessExportSnapshotFeatureUsage   *CloudProviderAccessExportSnapshotFeatureUsage
}

// CloudProviderAccessDataLakeFeatureUsageAsCloudProviderAccessFeatureUsage is a convenience function that returns CloudProviderAccessDataLakeFeatureUsage wrapped in CloudProviderAccessFeatureUsage
func CloudProviderAccessDataLakeFeatureUsageAsCloudProviderAccessFeatureUsage(v *CloudProviderAccessDataLakeFeatureUsage) CloudProviderAccessFeatureUsage {
	return CloudProviderAccessFeatureUsage{
		CloudProviderAccessDataLakeFeatureUsage: v,
	}
}

// CloudProviderAccessEncryptionAtRestFeatureUsageAsCloudProviderAccessFeatureUsage is a convenience function that returns CloudProviderAccessEncryptionAtRestFeatureUsage wrapped in CloudProviderAccessFeatureUsage
func CloudProviderAccessEncryptionAtRestFeatureUsageAsCloudProviderAccessFeatureUsage(v *CloudProviderAccessEncryptionAtRestFeatureUsage) CloudProviderAccessFeatureUsage {
	return CloudProviderAccessFeatureUsage{
		CloudProviderAccessEncryptionAtRestFeatureUsage: v,
	}
}

// CloudProviderAccessExportSnapshotFeatureUsageAsCloudProviderAccessFeatureUsage is a convenience function that returns CloudProviderAccessExportSnapshotFeatureUsage wrapped in CloudProviderAccessFeatureUsage
func CloudProviderAccessExportSnapshotFeatureUsageAsCloudProviderAccessFeatureUsage(v *CloudProviderAccessExportSnapshotFeatureUsage) CloudProviderAccessFeatureUsage {
	return CloudProviderAccessFeatureUsage{
		CloudProviderAccessExportSnapshotFeatureUsage: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CloudProviderAccessFeatureUsage) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ATLAS_DATA_LAKE'
	if jsonDict["featureType"] == "ATLAS_DATA_LAKE" {
		// try to unmarshal JSON data into CloudProviderAccessDataLakeFeatureUsage
		err = json.Unmarshal(data, &dst.CloudProviderAccessDataLakeFeatureUsage)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessDataLakeFeatureUsage, return on the first match
		} else {
			dst.CloudProviderAccessDataLakeFeatureUsage = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessFeatureUsage as CloudProviderAccessDataLakeFeatureUsage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CloudProviderAccessDataLakeFeatureUsage'
	if jsonDict["featureType"] == "CloudProviderAccessDataLakeFeatureUsage" {
		// try to unmarshal JSON data into CloudProviderAccessDataLakeFeatureUsage
		err = json.Unmarshal(data, &dst.CloudProviderAccessDataLakeFeatureUsage)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessDataLakeFeatureUsage, return on the first match
		} else {
			dst.CloudProviderAccessDataLakeFeatureUsage = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessFeatureUsage as CloudProviderAccessDataLakeFeatureUsage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CloudProviderAccessEncryptionAtRestFeatureUsage'
	if jsonDict["featureType"] == "CloudProviderAccessEncryptionAtRestFeatureUsage" {
		// try to unmarshal JSON data into CloudProviderAccessEncryptionAtRestFeatureUsage
		err = json.Unmarshal(data, &dst.CloudProviderAccessEncryptionAtRestFeatureUsage)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessEncryptionAtRestFeatureUsage, return on the first match
		} else {
			dst.CloudProviderAccessEncryptionAtRestFeatureUsage = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessFeatureUsage as CloudProviderAccessEncryptionAtRestFeatureUsage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CloudProviderAccessExportSnapshotFeatureUsage'
	if jsonDict["featureType"] == "CloudProviderAccessExportSnapshotFeatureUsage" {
		// try to unmarshal JSON data into CloudProviderAccessExportSnapshotFeatureUsage
		err = json.Unmarshal(data, &dst.CloudProviderAccessExportSnapshotFeatureUsage)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessExportSnapshotFeatureUsage, return on the first match
		} else {
			dst.CloudProviderAccessExportSnapshotFeatureUsage = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessFeatureUsage as CloudProviderAccessExportSnapshotFeatureUsage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ENCRYPTION_AT_REST'
	if jsonDict["featureType"] == "ENCRYPTION_AT_REST" {
		// try to unmarshal JSON data into CloudProviderAccessEncryptionAtRestFeatureUsage
		err = json.Unmarshal(data, &dst.CloudProviderAccessEncryptionAtRestFeatureUsage)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessEncryptionAtRestFeatureUsage, return on the first match
		} else {
			dst.CloudProviderAccessEncryptionAtRestFeatureUsage = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessFeatureUsage as CloudProviderAccessEncryptionAtRestFeatureUsage: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EXPORT_SNAPSHOT'
	if jsonDict["featureType"] == "EXPORT_SNAPSHOT" {
		// try to unmarshal JSON data into CloudProviderAccessExportSnapshotFeatureUsage
		err = json.Unmarshal(data, &dst.CloudProviderAccessExportSnapshotFeatureUsage)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessExportSnapshotFeatureUsage, return on the first match
		} else {
			dst.CloudProviderAccessExportSnapshotFeatureUsage = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessFeatureUsage as CloudProviderAccessExportSnapshotFeatureUsage: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CloudProviderAccessFeatureUsage) MarshalJSON() ([]byte, error) {
	if src.CloudProviderAccessDataLakeFeatureUsage != nil {
		return json.Marshal(&src.CloudProviderAccessDataLakeFeatureUsage)
	}

	if src.CloudProviderAccessEncryptionAtRestFeatureUsage != nil {
		return json.Marshal(&src.CloudProviderAccessEncryptionAtRestFeatureUsage)
	}

	if src.CloudProviderAccessExportSnapshotFeatureUsage != nil {
		return json.Marshal(&src.CloudProviderAccessExportSnapshotFeatureUsage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CloudProviderAccessFeatureUsage) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CloudProviderAccessDataLakeFeatureUsage != nil {
		return obj.CloudProviderAccessDataLakeFeatureUsage
	}

	if obj.CloudProviderAccessEncryptionAtRestFeatureUsage != nil {
		return obj.CloudProviderAccessEncryptionAtRestFeatureUsage
	}

	if obj.CloudProviderAccessExportSnapshotFeatureUsage != nil {
		return obj.CloudProviderAccessExportSnapshotFeatureUsage
	}

	// all schemas are nil
	return nil
}

type NullableCloudProviderAccessFeatureUsage struct {
	value *CloudProviderAccessFeatureUsage
	isSet bool
}

func (v NullableCloudProviderAccessFeatureUsage) Get() *CloudProviderAccessFeatureUsage {
	return v.value
}

func (v *NullableCloudProviderAccessFeatureUsage) Set(val *CloudProviderAccessFeatureUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessFeatureUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessFeatureUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessFeatureUsage(val *CloudProviderAccessFeatureUsage) *NullableCloudProviderAccessFeatureUsage {
	return &NullableCloudProviderAccessFeatureUsage{value: val, isSet: true}
}

func (v NullableCloudProviderAccessFeatureUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessFeatureUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
