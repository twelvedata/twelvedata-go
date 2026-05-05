// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
	"fmt"
)

// OrderEnum the model 'OrderEnum'
type OrderEnum string

// List of OrderEnum
const (
	ORDERENUM_ASC  OrderEnum = "asc"
	ORDERENUM_DESC OrderEnum = "desc"
)

// All allowed values of OrderEnum enum
var AllowedOrderEnumEnumValues = []OrderEnum{
	"asc",
	"desc",
}

func (v *OrderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderEnum(value)
	for _, existing := range AllowedOrderEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderEnum", value)
}

// NewOrderEnumFromValue returns a pointer to a valid OrderEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderEnumFromValue(v string) (*OrderEnum, error) {
	ev := OrderEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderEnum: valid values are %v", v, AllowedOrderEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderEnum) IsValid() bool {
	for _, existing := range AllowedOrderEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderEnum value
func (v OrderEnum) Ptr() *OrderEnum {
	return &v
}

type NullableOrderEnum struct {
	value *OrderEnum
	isSet bool
}

func (v NullableOrderEnum) Get() *OrderEnum {
	return v.value
}

func (v *NullableOrderEnum) Set(val *OrderEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderEnum(val *OrderEnum) *NullableOrderEnum {
	return &NullableOrderEnum{value: val, isSet: true}
}

func (v NullableOrderEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
