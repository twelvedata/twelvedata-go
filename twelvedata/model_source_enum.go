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

// SourceEnum the model 'SourceEnum'
type SourceEnum string

// List of SourceEnum
const (
	SOURCEENUM_OFAC SourceEnum = "ofac"
	SOURCEENUM_UK   SourceEnum = "uk"
	SOURCEENUM_EU   SourceEnum = "eu"
	SOURCEENUM_AU   SourceEnum = "au"
)

// All allowed values of SourceEnum enum
var AllowedSourceEnumEnumValues = []SourceEnum{
	"ofac",
	"uk",
	"eu",
	"au",
}

func (v *SourceEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SourceEnum(value)
	for _, existing := range AllowedSourceEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SourceEnum", value)
}

// NewSourceEnumFromValue returns a pointer to a valid SourceEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSourceEnumFromValue(v string) (*SourceEnum, error) {
	ev := SourceEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SourceEnum: valid values are %v", v, AllowedSourceEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SourceEnum) IsValid() bool {
	for _, existing := range AllowedSourceEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourceEnum value
func (v SourceEnum) Ptr() *SourceEnum {
	return &v
}

type NullableSourceEnum struct {
	value *SourceEnum
	isSet bool
}

func (v NullableSourceEnum) Get() *SourceEnum {
	return v.value
}

func (v *NullableSourceEnum) Set(val *SourceEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceEnum(val *SourceEnum) *NullableSourceEnum {
	return &NullableSourceEnum{value: val, isSet: true}
}

func (v NullableSourceEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
