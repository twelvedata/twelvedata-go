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

// SeriesTypeEnum the model 'SeriesTypeEnum'
type SeriesTypeEnum string

// List of SeriesTypeEnum
const (
	SERIESTYPEENUM_CLOSE  SeriesTypeEnum = "close"
	SERIESTYPEENUM_OPEN   SeriesTypeEnum = "open"
	SERIESTYPEENUM_HIGH   SeriesTypeEnum = "high"
	SERIESTYPEENUM_LOW    SeriesTypeEnum = "low"
	SERIESTYPEENUM_VOLUME SeriesTypeEnum = "volume"
)

// All allowed values of SeriesTypeEnum enum
var AllowedSeriesTypeEnumEnumValues = []SeriesTypeEnum{
	"close",
	"open",
	"high",
	"low",
	"volume",
}

func (v *SeriesTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SeriesTypeEnum(value)
	for _, existing := range AllowedSeriesTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SeriesTypeEnum", value)
}

// NewSeriesTypeEnumFromValue returns a pointer to a valid SeriesTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeriesTypeEnumFromValue(v string) (*SeriesTypeEnum, error) {
	ev := SeriesTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SeriesTypeEnum: valid values are %v", v, AllowedSeriesTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SeriesTypeEnum) IsValid() bool {
	for _, existing := range AllowedSeriesTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeriesTypeEnum value
func (v SeriesTypeEnum) Ptr() *SeriesTypeEnum {
	return &v
}

type NullableSeriesTypeEnum struct {
	value *SeriesTypeEnum
	isSet bool
}

func (v NullableSeriesTypeEnum) Get() *SeriesTypeEnum {
	return v.value
}

func (v *NullableSeriesTypeEnum) Set(val *SeriesTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesTypeEnum(val *SeriesTypeEnum) *NullableSeriesTypeEnum {
	return &NullableSeriesTypeEnum{value: val, isSet: true}
}

func (v NullableSeriesTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
