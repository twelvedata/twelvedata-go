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

// checks if the CryptocurrencyResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CryptocurrencyResponseItem{}

// CryptocurrencyResponseItem CryptocurrencyResponseItem represents details of a cryptocurrency
type CryptocurrencyResponseItem struct {
	// Cryptocurrency pair codes with slash(/) delimiter
	Symbol string `json:"symbol"`
	// List of exchanges where the cryptocurrency is available
	AvailableExchanges []string `json:"available_exchanges"`
	// Base currency of the cryptocurrency pair
	CurrencyBase string `json:"currency_base"`
	// Quote currency of the cryptocurrency pair
	CurrencyQuote string `json:"currency_quote"`
}

type _CryptocurrencyResponseItem CryptocurrencyResponseItem

// NewCryptocurrencyResponseItem instantiates a new CryptocurrencyResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCryptocurrencyResponseItem(symbol string, availableExchanges []string, currencyBase string, currencyQuote string) *CryptocurrencyResponseItem {
	this := CryptocurrencyResponseItem{}
	this.Symbol = symbol
	this.AvailableExchanges = availableExchanges
	this.CurrencyBase = currencyBase
	this.CurrencyQuote = currencyQuote
	return &this
}

// NewCryptocurrencyResponseItemWithDefaults instantiates a new CryptocurrencyResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCryptocurrencyResponseItemWithDefaults() *CryptocurrencyResponseItem {
	this := CryptocurrencyResponseItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *CryptocurrencyResponseItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *CryptocurrencyResponseItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *CryptocurrencyResponseItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetAvailableExchanges returns the AvailableExchanges field value
func (o *CryptocurrencyResponseItem) GetAvailableExchanges() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableExchanges
}

// GetAvailableExchangesOk returns a tuple with the AvailableExchanges field value
// and a boolean to check if the value has been set.
func (o *CryptocurrencyResponseItem) GetAvailableExchangesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableExchanges, true
}

// SetAvailableExchanges sets field value
func (o *CryptocurrencyResponseItem) SetAvailableExchanges(v []string) {
	o.AvailableExchanges = v
}

// GetCurrencyBase returns the CurrencyBase field value
func (o *CryptocurrencyResponseItem) GetCurrencyBase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyBase
}

// GetCurrencyBaseOk returns a tuple with the CurrencyBase field value
// and a boolean to check if the value has been set.
func (o *CryptocurrencyResponseItem) GetCurrencyBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyBase, true
}

// SetCurrencyBase sets field value
func (o *CryptocurrencyResponseItem) SetCurrencyBase(v string) {
	o.CurrencyBase = v
}

// GetCurrencyQuote returns the CurrencyQuote field value
func (o *CryptocurrencyResponseItem) GetCurrencyQuote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyQuote
}

// GetCurrencyQuoteOk returns a tuple with the CurrencyQuote field value
// and a boolean to check if the value has been set.
func (o *CryptocurrencyResponseItem) GetCurrencyQuoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyQuote, true
}

// SetCurrencyQuote sets field value
func (o *CryptocurrencyResponseItem) SetCurrencyQuote(v string) {
	o.CurrencyQuote = v
}

func (o CryptocurrencyResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CryptocurrencyResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["available_exchanges"] = o.AvailableExchanges
	toSerialize["currency_base"] = o.CurrencyBase
	toSerialize["currency_quote"] = o.CurrencyQuote
	return toSerialize, nil
}

func (o *CryptocurrencyResponseItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"available_exchanges",
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

	varCryptocurrencyResponseItem := _CryptocurrencyResponseItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCryptocurrencyResponseItem)

	if err != nil {
		return err
	}

	*o = CryptocurrencyResponseItem(varCryptocurrencyResponseItem)

	return err
}

type NullableCryptocurrencyResponseItem struct {
	value *CryptocurrencyResponseItem
	isSet bool
}

func (v NullableCryptocurrencyResponseItem) Get() *CryptocurrencyResponseItem {
	return v.value
}

func (v *NullableCryptocurrencyResponseItem) Set(val *CryptocurrencyResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCryptocurrencyResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCryptocurrencyResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCryptocurrencyResponseItem(val *CryptocurrencyResponseItem) *NullableCryptocurrencyResponseItem {
	return &NullableCryptocurrencyResponseItem{value: val, isSet: true}
}

func (v NullableCryptocurrencyResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCryptocurrencyResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
