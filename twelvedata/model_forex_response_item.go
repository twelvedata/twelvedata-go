/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ForexResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForexResponseItem{}

// ForexResponseItem ForexResponseItem represents details of a forex pair
type ForexResponseItem struct {
	// Currency pair according to ISO 4217 standard codes with slash(/) delimiter
	Symbol string `json:"symbol"`
	// Group to which currency pair belongs to, could be: Major, Minor, Exotic and Exotic-Cross
	CurrencyGroup string `json:"currency_group"`
	// Base currency name according to ISO 4217 standard
	CurrencyBase string `json:"currency_base"`
	// Quote currency name according to ISO 4217 standard
	CurrencyQuote string `json:"currency_quote"`
}

type _ForexResponseItem ForexResponseItem

// NewForexResponseItem instantiates a new ForexResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForexResponseItem(symbol string, currencyGroup string, currencyBase string, currencyQuote string) *ForexResponseItem {
	this := ForexResponseItem{}
	this.Symbol = symbol
	this.CurrencyGroup = currencyGroup
	this.CurrencyBase = currencyBase
	this.CurrencyQuote = currencyQuote
	return &this
}

// NewForexResponseItemWithDefaults instantiates a new ForexResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForexResponseItemWithDefaults() *ForexResponseItem {
	this := ForexResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ForexResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ForexResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ForexResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetCurrencyGroup returns the CurrencyGroup field value
func (o *ForexResponseItem) GetCurrencyGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyGroup
}

// GetCurrencyGroupOk returns a tuple with the CurrencyGroup field value
// and a boolean to check if the value has been set.
func (o *ForexResponseItem) GetCurrencyGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyGroup, true
}

// SetCurrencyGroup sets field value
func (o *ForexResponseItem) SetCurrencyGroup(v string) {
	o.CurrencyGroup = v
}

// GetCurrencyBase returns the CurrencyBase field value
func (o *ForexResponseItem) GetCurrencyBase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyBase
}

// GetCurrencyBaseOk returns a tuple with the CurrencyBase field value
// and a boolean to check if the value has been set.
func (o *ForexResponseItem) GetCurrencyBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyBase, true
}

// SetCurrencyBase sets field value
func (o *ForexResponseItem) SetCurrencyBase(v string) {
	o.CurrencyBase = v
}

// GetCurrencyQuote returns the CurrencyQuote field value
func (o *ForexResponseItem) GetCurrencyQuote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyQuote
}

// GetCurrencyQuoteOk returns a tuple with the CurrencyQuote field value
// and a boolean to check if the value has been set.
func (o *ForexResponseItem) GetCurrencyQuoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyQuote, true
}

// SetCurrencyQuote sets field value
func (o *ForexResponseItem) SetCurrencyQuote(v string) {
	o.CurrencyQuote = v
}

func (o ForexResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForexResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["currency_group"] = o.CurrencyGroup
	toSerialize["currency_base"] = o.CurrencyBase
	toSerialize["currency_quote"] = o.CurrencyQuote
	return toSerialize, nil
}

func (o *ForexResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"currency_group",
		"currency_base",
		"currency_quote",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varForexResponseItem := _ForexResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varForexResponseItem)

	if err != nil {
		return err
	}

	*o = ForexResponseItem(varForexResponseItem)

	return err
}

type NullableForexResponseItem struct {
	value *ForexResponseItem
	isSet bool
}

func (v NullableForexResponseItem) Get() *ForexResponseItem {
	return v.value
}

func (v *NullableForexResponseItem) Set(val *ForexResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableForexResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableForexResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForexResponseItem(val *ForexResponseItem) *NullableForexResponseItem {
	return &NullableForexResponseItem{value: val, isSet: true}
}

func (v NullableForexResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForexResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
