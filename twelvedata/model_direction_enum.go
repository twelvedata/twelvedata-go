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

// DirectionEnum the model 'DirectionEnum'
type DirectionEnum string

// List of DirectionEnum
const (
	DIRECTIONENUM_GAINERS DirectionEnum = "gainers"
	DIRECTIONENUM_LOSERS  DirectionEnum = "losers"
)

// All allowed values of DirectionEnum enum
var AllowedDirectionEnumEnumValues = []DirectionEnum{
	"gainers",
	"losers",
}

func (v *DirectionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DirectionEnum(value)
	for _, existing := range AllowedDirectionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DirectionEnum", value)
}

// NewDirectionEnumFromValue returns a pointer to a valid DirectionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDirectionEnumFromValue(v string) (*DirectionEnum, error) {
	ev := DirectionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DirectionEnum: valid values are %v", v, AllowedDirectionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DirectionEnum) IsValid() bool {
	for _, existing := range AllowedDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DirectionEnum value
func (v DirectionEnum) Ptr() *DirectionEnum {
	return &v
}

type NullableDirectionEnum struct {
	value *DirectionEnum
	isSet bool
}

func (v NullableDirectionEnum) Get() *DirectionEnum {
	return v.value
}

func (v *NullableDirectionEnum) Set(val *DirectionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectionEnum(val *DirectionEnum) *NullableDirectionEnum {
	return &NullableDirectionEnum{value: val, isSet: true}
}

func (v NullableDirectionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
