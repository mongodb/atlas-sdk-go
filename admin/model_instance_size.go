// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// InstanceSize Minimum instance size to which your cluster can automatically scale. MongoDB Cloud requires this parameter if `\"replicationSpecs[n].regionConfigs[m].autoScaling.compute.scaleDownEnabled\" : true`.
type InstanceSize string

// List of InstanceSize
const (
	INSTANCESIZE_M10       InstanceSize = "M10"
	INSTANCESIZE_M20       InstanceSize = "M20"
	INSTANCESIZE_M30       InstanceSize = "M30"
	INSTANCESIZE_M40       InstanceSize = "M40"
	INSTANCESIZE_M50       InstanceSize = "M50"
	INSTANCESIZE_M60       InstanceSize = "M60"
	INSTANCESIZE_M80       InstanceSize = "M80"
	INSTANCESIZE_M100      InstanceSize = "M100"
	INSTANCESIZE_M140      InstanceSize = "M140"
	INSTANCESIZE_M200      InstanceSize = "M200"
	INSTANCESIZE_M300      InstanceSize = "M300"
	INSTANCESIZE_R40       InstanceSize = "R40"
	INSTANCESIZE_R50       InstanceSize = "R50"
	INSTANCESIZE_R60       InstanceSize = "R60"
	INSTANCESIZE_R80       InstanceSize = "R80"
	INSTANCESIZE_R200      InstanceSize = "R200"
	INSTANCESIZE_R300      InstanceSize = "R300"
	INSTANCESIZE_R400      InstanceSize = "R400"
	INSTANCESIZE_R700      InstanceSize = "R700"
	INSTANCESIZE_M40_NVME  InstanceSize = "M40_NVME"
	INSTANCESIZE_M50_NVME  InstanceSize = "M50_NVME"
	INSTANCESIZE_M60_NVME  InstanceSize = "M60_NVME"
	INSTANCESIZE_M80_NVME  InstanceSize = "M80_NVME"
	INSTANCESIZE_M200_NVME InstanceSize = "M200_NVME"
	INSTANCESIZE_M400_NVME InstanceSize = "M400_NVME"
	INSTANCESIZE_M90       InstanceSize = "M90"
	INSTANCESIZE_M300_NVME InstanceSize = "M300_NVME"
	INSTANCESIZE_M600_NVME InstanceSize = "M600_NVME"
	INSTANCESIZE_M250      InstanceSize = "M250"
	INSTANCESIZE_M400      InstanceSize = "M400"
	INSTANCESIZE_R600      InstanceSize = "R600"
)

// All allowed values of InstanceSize enum
var AllowedInstanceSizeEnumValues = []InstanceSize{
	"M10",
	"M20",
	"M30",
	"M40",
	"M50",
	"M60",
	"M80",
	"M100",
	"M140",
	"M200",
	"M300",
	"R40",
	"R50",
	"R60",
	"R80",
	"R200",
	"R300",
	"R400",
	"R700",
	"M40_NVME",
	"M50_NVME",
	"M60_NVME",
	"M80_NVME",
	"M200_NVME",
	"M400_NVME",
	"M90",
	"M300_NVME",
	"M600_NVME",
	"M250",
	"M400",
	"R600",
}

func (v *InstanceSize) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InstanceSize(value)
	for _, existing := range AllowedInstanceSizeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InstanceSize", value)
}

// NewInstanceSizeFromValue returns a pointer to a valid InstanceSize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInstanceSizeFromValue(v string) (*InstanceSize, error) {
	ev := InstanceSize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InstanceSize: valid values are %v", v, AllowedInstanceSizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InstanceSize) IsValid() bool {
	for _, existing := range AllowedInstanceSizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InstanceSize value
func (v InstanceSize) Ptr() *InstanceSize {
	return &v
}
