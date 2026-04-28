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

// PeriodEnum the model 'PeriodEnum'
type PeriodEnum string

// List of PeriodEnum
const (
	PERIODENUM_ANNUAL    PeriodEnum = "annual"
	PERIODENUM_QUARTERLY PeriodEnum = "quarterly"
)

// All allowed values of PeriodEnum enum
var AllowedPeriodEnumEnumValues = []PeriodEnum{
	"annual",
	"quarterly",
}

func (v *PeriodEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PeriodEnum(value)
	for _, existing := range AllowedPeriodEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PeriodEnum", value)
}

// NewPeriodEnumFromValue returns a pointer to a valid PeriodEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPeriodEnumFromValue(v string) (*PeriodEnum, error) {
	ev := PeriodEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PeriodEnum: valid values are %v", v, AllowedPeriodEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PeriodEnum) IsValid() bool {
	for _, existing := range AllowedPeriodEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PeriodEnum value
func (v PeriodEnum) Ptr() *PeriodEnum {
	return &v
}

type NullablePeriodEnum struct {
	value *PeriodEnum
	isSet bool
}

func (v NullablePeriodEnum) Get() *PeriodEnum {
	return v.value
}

func (v *NullablePeriodEnum) Set(val *PeriodEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodEnum(val *PeriodEnum) *NullablePeriodEnum {
	return &NullablePeriodEnum{value: val, isSet: true}
}

func (v NullablePeriodEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
