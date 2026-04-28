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

// PeriodEarningsEnum the model 'PeriodEarningsEnum'
type PeriodEarningsEnum string

// List of PeriodEarningsEnum
const (
	PERIODEARNINGSENUM_LATEST PeriodEarningsEnum = "latest"
	PERIODEARNINGSENUM_NEXT   PeriodEarningsEnum = "next"
)

// All allowed values of PeriodEarningsEnum enum
var AllowedPeriodEarningsEnumEnumValues = []PeriodEarningsEnum{
	"latest",
	"next",
}

func (v *PeriodEarningsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PeriodEarningsEnum(value)
	for _, existing := range AllowedPeriodEarningsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PeriodEarningsEnum", value)
}

// NewPeriodEarningsEnumFromValue returns a pointer to a valid PeriodEarningsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPeriodEarningsEnumFromValue(v string) (*PeriodEarningsEnum, error) {
	ev := PeriodEarningsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PeriodEarningsEnum: valid values are %v", v, AllowedPeriodEarningsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PeriodEarningsEnum) IsValid() bool {
	for _, existing := range AllowedPeriodEarningsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PeriodEarningsEnum value
func (v PeriodEarningsEnum) Ptr() *PeriodEarningsEnum {
	return &v
}

type NullablePeriodEarningsEnum struct {
	value *PeriodEarningsEnum
	isSet bool
}

func (v NullablePeriodEarningsEnum) Get() *PeriodEarningsEnum {
	return v.value
}

func (v *NullablePeriodEarningsEnum) Set(val *PeriodEarningsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodEarningsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodEarningsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodEarningsEnum(val *PeriodEarningsEnum) *NullablePeriodEarningsEnum {
	return &NullablePeriodEarningsEnum{value: val, isSet: true}
}

func (v NullablePeriodEarningsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodEarningsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
