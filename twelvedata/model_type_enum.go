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

// TypeEnum the model 'TypeEnum'
type TypeEnum string

// List of TypeEnum
const (
	TYPEENUM_AMERICAN_DEPOSITARY_RECEIPT TypeEnum = "American Depositary Receipt"
	TYPEENUM_BOND                        TypeEnum = "Bond"
	TYPEENUM_BOND_FUND                   TypeEnum = "Bond Fund"
	TYPEENUM_CLOSED_END_FUND             TypeEnum = "Closed-end Fund"
	TYPEENUM_COMMON_STOCK                TypeEnum = "Common Stock"
	TYPEENUM_DEPOSITARY_RECEIPT          TypeEnum = "Depositary Receipt"
	TYPEENUM_DIGITAL_CURRENCY            TypeEnum = "Digital Currency"
	TYPEENUM_ETF                         TypeEnum = "ETF"
	TYPEENUM_EXCHANGE_TRADED_NOTE        TypeEnum = "Exchange-Traded Note"
	TYPEENUM_GLOBAL_DEPOSITARY_RECEIPT   TypeEnum = "Global Depositary Receipt"
	TYPEENUM_LIMITED_PARTNERSHIP         TypeEnum = "Limited Partnership"
	TYPEENUM_MUTUAL_FUND                 TypeEnum = "Mutual Fund"
	TYPEENUM_PHYSICAL_CURRENCY           TypeEnum = "Physical Currency"
	TYPEENUM_PREFERRED_STOCK             TypeEnum = "Preferred Stock"
	TYPEENUM_REIT                        TypeEnum = "REIT"
	TYPEENUM_RIGHT                       TypeEnum = "Right"
	TYPEENUM_STRUCTURED_PRODUCT          TypeEnum = "Structured Product"
	TYPEENUM_TRUST                       TypeEnum = "Trust"
	TYPEENUM_UNIT                        TypeEnum = "Unit"
	TYPEENUM_WARRANT                     TypeEnum = "Warrant"
)

// All allowed values of TypeEnum enum
var AllowedTypeEnumEnumValues = []TypeEnum{
	"American Depositary Receipt",
	"Bond",
	"Bond Fund",
	"Closed-end Fund",
	"Common Stock",
	"Depositary Receipt",
	"Digital Currency",
	"ETF",
	"Exchange-Traded Note",
	"Global Depositary Receipt",
	"Limited Partnership",
	"Mutual Fund",
	"Physical Currency",
	"Preferred Stock",
	"REIT",
	"Right",
	"Structured Product",
	"Trust",
	"Unit",
	"Warrant",
}

func (v *TypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeEnum(value)
	for _, existing := range AllowedTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeEnum", value)
}

// NewTypeEnumFromValue returns a pointer to a valid TypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeEnumFromValue(v string) (*TypeEnum, error) {
	ev := TypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypeEnum: valid values are %v", v, AllowedTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeEnum) IsValid() bool {
	for _, existing := range AllowedTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TypeEnum value
func (v TypeEnum) Ptr() *TypeEnum {
	return &v
}

type NullableTypeEnum struct {
	value *TypeEnum
	isSet bool
}

func (v NullableTypeEnum) Get() *TypeEnum {
	return v.value
}

func (v *NullableTypeEnum) Set(val *TypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypeEnum(val *TypeEnum) *NullableTypeEnum {
	return &NullableTypeEnum{value: val, isSet: true}
}

func (v NullableTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
