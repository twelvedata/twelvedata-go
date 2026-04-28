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

// IntervalEnum the model 'IntervalEnum'
type IntervalEnum string

// List of IntervalEnum
const (
	INTERVALENUM__1MIN   IntervalEnum = "1min"
	INTERVALENUM__5MIN   IntervalEnum = "5min"
	INTERVALENUM__15MIN  IntervalEnum = "15min"
	INTERVALENUM__30MIN  IntervalEnum = "30min"
	INTERVALENUM__45MIN  IntervalEnum = "45min"
	INTERVALENUM__1H     IntervalEnum = "1h"
	INTERVALENUM__2H     IntervalEnum = "2h"
	INTERVALENUM__4H     IntervalEnum = "4h"
	INTERVALENUM__8H     IntervalEnum = "8h"
	INTERVALENUM__1DAY   IntervalEnum = "1day"
	INTERVALENUM__1WEEK  IntervalEnum = "1week"
	INTERVALENUM__1MONTH IntervalEnum = "1month"
)

// All allowed values of IntervalEnum enum
var AllowedIntervalEnumEnumValues = []IntervalEnum{
	"1min",
	"5min",
	"15min",
	"30min",
	"45min",
	"1h",
	"2h",
	"4h",
	"8h",
	"1day",
	"1week",
	"1month",
}

func (v *IntervalEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntervalEnum(value)
	for _, existing := range AllowedIntervalEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntervalEnum", value)
}

// NewIntervalEnumFromValue returns a pointer to a valid IntervalEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntervalEnumFromValue(v string) (*IntervalEnum, error) {
	ev := IntervalEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntervalEnum: valid values are %v", v, AllowedIntervalEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntervalEnum) IsValid() bool {
	for _, existing := range AllowedIntervalEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntervalEnum value
func (v IntervalEnum) Ptr() *IntervalEnum {
	return &v
}

type NullableIntervalEnum struct {
	value *IntervalEnum
	isSet bool
}

func (v NullableIntervalEnum) Get() *IntervalEnum {
	return v.value
}

func (v *NullableIntervalEnum) Set(val *IntervalEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableIntervalEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableIntervalEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntervalEnum(val *IntervalEnum) *NullableIntervalEnum {
	return &NullableIntervalEnum{value: val, isSet: true}
}

func (v NullableIntervalEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntervalEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
