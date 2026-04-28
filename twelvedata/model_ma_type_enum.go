/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
	"fmt"
)

// MaTypeEnum the model 'MaTypeEnum'
type MaTypeEnum string

// List of MaTypeEnum
const (
	MATYPEENUM_SMA   MaTypeEnum = "SMA"
	MATYPEENUM_EMA   MaTypeEnum = "EMA"
	MATYPEENUM_WMA   MaTypeEnum = "WMA"
	MATYPEENUM_DEMA  MaTypeEnum = "DEMA"
	MATYPEENUM_TEMA  MaTypeEnum = "TEMA"
	MATYPEENUM_TRIMA MaTypeEnum = "TRIMA"
	MATYPEENUM_KAMA  MaTypeEnum = "KAMA"
	MATYPEENUM_MAMA  MaTypeEnum = "MAMA"
	MATYPEENUM_T3_MA MaTypeEnum = "T3MA"
)

// All allowed values of MaTypeEnum enum
var AllowedMaTypeEnumEnumValues = []MaTypeEnum{
	"SMA",
	"EMA",
	"WMA",
	"DEMA",
	"TEMA",
	"TRIMA",
	"KAMA",
	"MAMA",
	"T3MA",
}

func (v *MaTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MaTypeEnum(value)
	for _, existing := range AllowedMaTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MaTypeEnum", value)
}

// NewMaTypeEnumFromValue returns a pointer to a valid MaTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMaTypeEnumFromValue(v string) (*MaTypeEnum, error) {
	ev := MaTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MaTypeEnum: valid values are %v", v, AllowedMaTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MaTypeEnum) IsValid() bool {
	for _, existing := range AllowedMaTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MaTypeEnum value
func (v MaTypeEnum) Ptr() *MaTypeEnum {
	return &v
}

type NullableMaTypeEnum struct {
	value *MaTypeEnum
	isSet bool
}

func (v NullableMaTypeEnum) Get() *MaTypeEnum {
	return v.value
}

func (v *NullableMaTypeEnum) Set(val *MaTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMaTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMaTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaTypeEnum(val *MaTypeEnum) *NullableMaTypeEnum {
	return &NullableMaTypeEnum{value: val, isSet: true}
}

func (v NullableMaTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
