// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// DataLakeStoreSettings - Group of settings that define where the data is stored.
type DataLakeStoreSettings struct {
	DataLakeAtlasStoreInstance *DataLakeAtlasStoreInstance
	DataLakeHTTPStore          *DataLakeHTTPStore
	DataLakeOnlineArchiveStore *DataLakeOnlineArchiveStore
	DataLakeS3StoreSettings    *DataLakeS3StoreSettings
}

// DataLakeAtlasStoreInstanceAsDataLakeStoreSettings is a convenience function that returns DataLakeAtlasStoreInstance wrapped in DataLakeStoreSettings
func DataLakeAtlasStoreInstanceAsDataLakeStoreSettings(v *DataLakeAtlasStoreInstance) DataLakeStoreSettings {
	return DataLakeStoreSettings{
		DataLakeAtlasStoreInstance: v,
	}
}

// DataLakeHTTPStoreAsDataLakeStoreSettings is a convenience function that returns DataLakeHTTPStore wrapped in DataLakeStoreSettings
func DataLakeHTTPStoreAsDataLakeStoreSettings(v *DataLakeHTTPStore) DataLakeStoreSettings {
	return DataLakeStoreSettings{
		DataLakeHTTPStore: v,
	}
}

// DataLakeOnlineArchiveStoreAsDataLakeStoreSettings is a convenience function that returns DataLakeOnlineArchiveStore wrapped in DataLakeStoreSettings
func DataLakeOnlineArchiveStoreAsDataLakeStoreSettings(v *DataLakeOnlineArchiveStore) DataLakeStoreSettings {
	return DataLakeStoreSettings{
		DataLakeOnlineArchiveStore: v,
	}
}

// DataLakeS3StoreSettingsAsDataLakeStoreSettings is a convenience function that returns DataLakeS3StoreSettings wrapped in DataLakeStoreSettings
func DataLakeS3StoreSettingsAsDataLakeStoreSettings(v *DataLakeS3StoreSettings) DataLakeStoreSettings {
	return DataLakeStoreSettings{
		DataLakeS3StoreSettings: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataLakeStoreSettings) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DataLakeAtlasStoreInstance'
	if jsonDict["provider"] == "DataLakeAtlasStoreInstance" {
		// try to unmarshal JSON data into DataLakeAtlasStoreInstance
		err = json.Unmarshal(data, &dst.DataLakeAtlasStoreInstance)
		if err == nil {
			return nil // data stored in dst.DataLakeAtlasStoreInstance, return on the first match
		} else {
			dst.DataLakeAtlasStoreInstance = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeAtlasStoreInstance: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DataLakeHTTPStore'
	if jsonDict["provider"] == "DataLakeHTTPStore" {
		// try to unmarshal JSON data into DataLakeHTTPStore
		err = json.Unmarshal(data, &dst.DataLakeHTTPStore)
		if err == nil {
			return nil // data stored in dst.DataLakeHTTPStore, return on the first match
		} else {
			dst.DataLakeHTTPStore = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeHTTPStore: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DataLakeOnlineArchiveStore'
	if jsonDict["provider"] == "DataLakeOnlineArchiveStore" {
		// try to unmarshal JSON data into DataLakeOnlineArchiveStore
		err = json.Unmarshal(data, &dst.DataLakeOnlineArchiveStore)
		if err == nil {
			return nil // data stored in dst.DataLakeOnlineArchiveStore, return on the first match
		} else {
			dst.DataLakeOnlineArchiveStore = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeOnlineArchiveStore: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DataLakeS3StoreSettings'
	if jsonDict["provider"] == "DataLakeS3StoreSettings" {
		// try to unmarshal JSON data into DataLakeS3StoreSettings
		err = json.Unmarshal(data, &dst.DataLakeS3StoreSettings)
		if err == nil {
			return nil // data stored in dst.DataLakeS3StoreSettings, return on the first match
		} else {
			dst.DataLakeS3StoreSettings = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeS3StoreSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'atlas'
	if jsonDict["provider"] == "atlas" {
		// try to unmarshal JSON data into DataLakeAtlasStoreInstance
		err = json.Unmarshal(data, &dst.DataLakeAtlasStoreInstance)
		if err == nil {
			return nil // data stored in dst.DataLakeAtlasStoreInstance, return on the first match
		} else {
			dst.DataLakeAtlasStoreInstance = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeAtlasStoreInstance: %s", err.Error())
		}
	}

	// check if the discriminator value is 'http'
	if jsonDict["provider"] == "http" {
		// try to unmarshal JSON data into DataLakeHTTPStore
		err = json.Unmarshal(data, &dst.DataLakeHTTPStore)
		if err == nil {
			return nil // data stored in dst.DataLakeHTTPStore, return on the first match
		} else {
			dst.DataLakeHTTPStore = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeHTTPStore: %s", err.Error())
		}
	}

	// check if the discriminator value is 'online_archive'
	if jsonDict["provider"] == "online_archive" {
		// try to unmarshal JSON data into DataLakeOnlineArchiveStore
		err = json.Unmarshal(data, &dst.DataLakeOnlineArchiveStore)
		if err == nil {
			return nil // data stored in dst.DataLakeOnlineArchiveStore, return on the first match
		} else {
			dst.DataLakeOnlineArchiveStore = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeOnlineArchiveStore: %s", err.Error())
		}
	}

	// check if the discriminator value is 's3'
	if jsonDict["provider"] == "s3" {
		// try to unmarshal JSON data into DataLakeS3StoreSettings
		err = json.Unmarshal(data, &dst.DataLakeS3StoreSettings)
		if err == nil {
			return nil // data stored in dst.DataLakeS3StoreSettings, return on the first match
		} else {
			dst.DataLakeS3StoreSettings = nil
			return fmt.Errorf("failed to unmarshal DataLakeStoreSettings as DataLakeS3StoreSettings: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataLakeStoreSettings) MarshalJSON() ([]byte, error) {
	if src.DataLakeAtlasStoreInstance != nil {
		return json.Marshal(&src.DataLakeAtlasStoreInstance)
	}

	if src.DataLakeHTTPStore != nil {
		return json.Marshal(&src.DataLakeHTTPStore)
	}

	if src.DataLakeOnlineArchiveStore != nil {
		return json.Marshal(&src.DataLakeOnlineArchiveStore)
	}

	if src.DataLakeS3StoreSettings != nil {
		return json.Marshal(&src.DataLakeS3StoreSettings)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataLakeStoreSettings) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DataLakeAtlasStoreInstance != nil {
		return obj.DataLakeAtlasStoreInstance
	}

	if obj.DataLakeHTTPStore != nil {
		return obj.DataLakeHTTPStore
	}

	if obj.DataLakeOnlineArchiveStore != nil {
		return obj.DataLakeOnlineArchiveStore
	}

	if obj.DataLakeS3StoreSettings != nil {
		return obj.DataLakeS3StoreSettings
	}

	// all schemas are nil
	return nil
}

type NullableDataLakeStoreSettings struct {
	value *DataLakeStoreSettings
	isSet bool
}

func (v NullableDataLakeStoreSettings) Get() *DataLakeStoreSettings {
	return v.value
}

func (v *NullableDataLakeStoreSettings) Set(val *DataLakeStoreSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeStoreSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeStoreSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeStoreSettings(val *DataLakeStoreSettings) *NullableDataLakeStoreSettings {
	return &NullableDataLakeStoreSettings{value: val, isSet: true}
}

func (v NullableDataLakeStoreSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeStoreSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
