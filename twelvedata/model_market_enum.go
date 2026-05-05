// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
	"fmt"
)

// MarketEnum the model 'MarketEnum'
type MarketEnum string

// List of MarketEnum
const (
	MARKETENUM_STOCKS       MarketEnum = "stocks"
	MARKETENUM_ETF          MarketEnum = "etf"
	MARKETENUM_MUTUAL_FUNDS MarketEnum = "mutual_funds"
	MARKETENUM_FOREX        MarketEnum = "forex"
	MARKETENUM_CRYPTO       MarketEnum = "crypto"
)

// All allowed values of MarketEnum enum
var AllowedMarketEnumEnumValues = []MarketEnum{
	"stocks",
	"etf",
	"mutual_funds",
	"forex",
	"crypto",
}

func (v *MarketEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarketEnum(value)
	for _, existing := range AllowedMarketEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarketEnum", value)
}

// NewMarketEnumFromValue returns a pointer to a valid MarketEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarketEnumFromValue(v string) (*MarketEnum, error) {
	ev := MarketEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarketEnum: valid values are %v", v, AllowedMarketEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarketEnum) IsValid() bool {
	for _, existing := range AllowedMarketEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MarketEnum value
func (v MarketEnum) Ptr() *MarketEnum {
	return &v
}

type NullableMarketEnum struct {
	value *MarketEnum
	isSet bool
}

func (v NullableMarketEnum) Get() *MarketEnum {
	return v.value
}

func (v *NullableMarketEnum) Set(val *MarketEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketEnum(val *MarketEnum) *NullableMarketEnum {
	return &NullableMarketEnum{value: val, isSet: true}
}

func (v NullableMarketEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
