// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
	"fmt"
)

// AdjustEnum the model 'AdjustEnum'
type AdjustEnum string

// List of AdjustEnum
const (
	ADJUSTENUM_ALL       AdjustEnum = "all"
	ADJUSTENUM_SPLITS    AdjustEnum = "splits"
	ADJUSTENUM_DIVIDENDS AdjustEnum = "dividends"
	ADJUSTENUM_NONE      AdjustEnum = "none"
)

// All allowed values of AdjustEnum enum
var AllowedAdjustEnumEnumValues = []AdjustEnum{
	"all",
	"splits",
	"dividends",
	"none",
}

func (v *AdjustEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdjustEnum(value)
	for _, existing := range AllowedAdjustEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdjustEnum", value)
}

// NewAdjustEnumFromValue returns a pointer to a valid AdjustEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdjustEnumFromValue(v string) (*AdjustEnum, error) {
	ev := AdjustEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdjustEnum: valid values are %v", v, AllowedAdjustEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdjustEnum) IsValid() bool {
	for _, existing := range AllowedAdjustEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdjustEnum value
func (v AdjustEnum) Ptr() *AdjustEnum {
	return &v
}

type NullableAdjustEnum struct {
	value *AdjustEnum
	isSet bool
}

func (v NullableAdjustEnum) Get() *AdjustEnum {
	return v.value
}

func (v *NullableAdjustEnum) Set(val *AdjustEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustEnum(val *AdjustEnum) *NullableAdjustEnum {
	return &NullableAdjustEnum{value: val, isSet: true}
}

func (v NullableAdjustEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
