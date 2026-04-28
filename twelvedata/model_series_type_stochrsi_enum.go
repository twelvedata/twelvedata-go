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

// SeriesTypeStochrsiEnum the model 'SeriesTypeStochrsiEnum'
type SeriesTypeStochrsiEnum string

// List of SeriesTypeStochrsiEnum
const (
	SERIESTYPESTOCHRSIENUM_OPEN  SeriesTypeStochrsiEnum = "open"
	SERIESTYPESTOCHRSIENUM_HIGH  SeriesTypeStochrsiEnum = "high"
	SERIESTYPESTOCHRSIENUM_LOW   SeriesTypeStochrsiEnum = "low"
	SERIESTYPESTOCHRSIENUM_CLOSE SeriesTypeStochrsiEnum = "close"
)

// All allowed values of SeriesTypeStochrsiEnum enum
var AllowedSeriesTypeStochrsiEnumEnumValues = []SeriesTypeStochrsiEnum{
	"open",
	"high",
	"low",
	"close",
}

func (v *SeriesTypeStochrsiEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SeriesTypeStochrsiEnum(value)
	for _, existing := range AllowedSeriesTypeStochrsiEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SeriesTypeStochrsiEnum", value)
}

// NewSeriesTypeStochrsiEnumFromValue returns a pointer to a valid SeriesTypeStochrsiEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeriesTypeStochrsiEnumFromValue(v string) (*SeriesTypeStochrsiEnum, error) {
	ev := SeriesTypeStochrsiEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SeriesTypeStochrsiEnum: valid values are %v", v, AllowedSeriesTypeStochrsiEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SeriesTypeStochrsiEnum) IsValid() bool {
	for _, existing := range AllowedSeriesTypeStochrsiEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeriesTypeStochrsiEnum value
func (v SeriesTypeStochrsiEnum) Ptr() *SeriesTypeStochrsiEnum {
	return &v
}

type NullableSeriesTypeStochrsiEnum struct {
	value *SeriesTypeStochrsiEnum
	isSet bool
}

func (v NullableSeriesTypeStochrsiEnum) Get() *SeriesTypeStochrsiEnum {
	return v.value
}

func (v *NullableSeriesTypeStochrsiEnum) Set(val *SeriesTypeStochrsiEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesTypeStochrsiEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesTypeStochrsiEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesTypeStochrsiEnum(val *SeriesTypeStochrsiEnum) *NullableSeriesTypeStochrsiEnum {
	return &NullableSeriesTypeStochrsiEnum{value: val, isSet: true}
}

func (v NullableSeriesTypeStochrsiEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesTypeStochrsiEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
