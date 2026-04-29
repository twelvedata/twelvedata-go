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

// checks if the GetInsiderTransactions200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsiderTransactions200ResponseMeta{}

// GetInsiderTransactions200ResponseMeta Metadata about the instrument
type GetInsiderTransactions200ResponseMeta struct {
	// Ticker symbol of instrument
	Symbol string `json:"symbol"`
	// Name of the company
	Name string `json:"name"`
	// Currency of the instrument according to the ISO 4217 standard
	Currency string `json:"currency"`
	// Exchange where instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Timezone of the exchange
	ExchangeTimezone string `json:"exchange_timezone"`
}

type _GetInsiderTransactions200ResponseMeta GetInsiderTransactions200ResponseMeta

// NewGetInsiderTransactions200ResponseMeta instantiates a new GetInsiderTransactions200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsiderTransactions200ResponseMeta(symbol string, name string, currency string, exchange string, micCode string, exchangeTimezone string) *GetInsiderTransactions200ResponseMeta {
	this := GetInsiderTransactions200ResponseMeta{}
	this.Symbol = symbol
	this.Name = name
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.ExchangeTimezone = exchangeTimezone
	return &this
}

// NewGetInsiderTransactions200ResponseMetaWithDefaults instantiates a new GetInsiderTransactions200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsiderTransactions200ResponseMetaWithDefaults() *GetInsiderTransactions200ResponseMeta {
	this := GetInsiderTransactions200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetInsiderTransactions200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetInsiderTransactions200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetInsiderTransactions200ResponseMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetInsiderTransactions200ResponseMeta) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *GetInsiderTransactions200ResponseMeta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetInsiderTransactions200ResponseMeta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *GetInsiderTransactions200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetInsiderTransactions200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetInsiderTransactions200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetInsiderTransactions200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *GetInsiderTransactions200ResponseMeta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *GetInsiderTransactions200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *GetInsiderTransactions200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

func (o GetInsiderTransactions200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsiderTransactions200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	return toSerialize, nil
}

func (o *GetInsiderTransactions200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"currency",
		"exchange",
		"mic_code",
		"exchange_timezone",
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

	varGetInsiderTransactions200ResponseMeta := _GetInsiderTransactions200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetInsiderTransactions200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetInsiderTransactions200ResponseMeta(varGetInsiderTransactions200ResponseMeta)

	return err
}

type NullableGetInsiderTransactions200ResponseMeta struct {
	value *GetInsiderTransactions200ResponseMeta
	isSet bool
}

func (v NullableGetInsiderTransactions200ResponseMeta) Get() *GetInsiderTransactions200ResponseMeta {
	return v.value
}

func (v *NullableGetInsiderTransactions200ResponseMeta) Set(val *GetInsiderTransactions200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsiderTransactions200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsiderTransactions200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsiderTransactions200ResponseMeta(val *GetInsiderTransactions200ResponseMeta) *NullableGetInsiderTransactions200ResponseMeta {
	return &NullableGetInsiderTransactions200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetInsiderTransactions200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsiderTransactions200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
