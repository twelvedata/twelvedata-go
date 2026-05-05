// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
	"fmt"
)

// RangeSplitsEnum the model 'RangeSplitsEnum'
type RangeSplitsEnum string

// List of RangeSplitsEnum
const (
	RANGESPLITSENUM_LAST RangeSplitsEnum = "last"
	RANGESPLITSENUM__1M  RangeSplitsEnum = "1m"
	RANGESPLITSENUM__3M  RangeSplitsEnum = "3m"
	RANGESPLITSENUM__6M  RangeSplitsEnum = "6m"
	RANGESPLITSENUM_YTD  RangeSplitsEnum = "ytd"
	RANGESPLITSENUM__1Y  RangeSplitsEnum = "1y"
	RANGESPLITSENUM__2Y  RangeSplitsEnum = "2y"
	RANGESPLITSENUM__5Y  RangeSplitsEnum = "5y"
	RANGESPLITSENUM_FULL RangeSplitsEnum = "full"
)

// All allowed values of RangeSplitsEnum enum
var AllowedRangeSplitsEnumEnumValues = []RangeSplitsEnum{
	"last",
	"1m",
	"3m",
	"6m",
	"ytd",
	"1y",
	"2y",
	"5y",
	"full",
}

func (v *RangeSplitsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RangeSplitsEnum(value)
	for _, existing := range AllowedRangeSplitsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RangeSplitsEnum", value)
}

// NewRangeSplitsEnumFromValue returns a pointer to a valid RangeSplitsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRangeSplitsEnumFromValue(v string) (*RangeSplitsEnum, error) {
	ev := RangeSplitsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RangeSplitsEnum: valid values are %v", v, AllowedRangeSplitsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RangeSplitsEnum) IsValid() bool {
	for _, existing := range AllowedRangeSplitsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RangeSplitsEnum value
func (v RangeSplitsEnum) Ptr() *RangeSplitsEnum {
	return &v
}

type NullableRangeSplitsEnum struct {
	value *RangeSplitsEnum
	isSet bool
}

func (v NullableRangeSplitsEnum) Get() *RangeSplitsEnum {
	return v.value
}

func (v *NullableRangeSplitsEnum) Set(val *RangeSplitsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeSplitsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeSplitsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeSplitsEnum(val *RangeSplitsEnum) *NullableRangeSplitsEnum {
	return &NullableRangeSplitsEnum{value: val, isSet: true}
}

func (v NullableRangeSplitsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeSplitsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
