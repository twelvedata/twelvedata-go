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

// RangeEnum the model 'RangeEnum'
type RangeEnum string

// List of RangeEnum
const (
	RANGEENUM_LAST RangeEnum = "last"
	RANGEENUM_NEXT RangeEnum = "next"
	RANGEENUM__1M  RangeEnum = "1m"
	RANGEENUM__3M  RangeEnum = "3m"
	RANGEENUM__6M  RangeEnum = "6m"
	RANGEENUM_YTD  RangeEnum = "ytd"
	RANGEENUM__1Y  RangeEnum = "1y"
	RANGEENUM__2Y  RangeEnum = "2y"
	RANGEENUM__5Y  RangeEnum = "5y"
	RANGEENUM_FULL RangeEnum = "full"
)

// All allowed values of RangeEnum enum
var AllowedRangeEnumEnumValues = []RangeEnum{
	"last",
	"next",
	"1m",
	"3m",
	"6m",
	"ytd",
	"1y",
	"2y",
	"5y",
	"full",
}

func (v *RangeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RangeEnum(value)
	for _, existing := range AllowedRangeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RangeEnum", value)
}

// NewRangeEnumFromValue returns a pointer to a valid RangeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRangeEnumFromValue(v string) (*RangeEnum, error) {
	ev := RangeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RangeEnum: valid values are %v", v, AllowedRangeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RangeEnum) IsValid() bool {
	for _, existing := range AllowedRangeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RangeEnum value
func (v RangeEnum) Ptr() *RangeEnum {
	return &v
}

type NullableRangeEnum struct {
	value *RangeEnum
	isSet bool
}

func (v NullableRangeEnum) Get() *RangeEnum {
	return v.value
}

func (v *NullableRangeEnum) Set(val *RangeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeEnum(val *RangeEnum) *NullableRangeEnum {
	return &NullableRangeEnum{value: val, isSet: true}
}

func (v NullableRangeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
