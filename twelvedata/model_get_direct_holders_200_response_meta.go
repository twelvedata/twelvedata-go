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

// checks if the GetDirectHolders200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDirectHolders200ResponseMeta{}

// GetDirectHolders200ResponseMeta Meta information about the financial instrument
type GetDirectHolders200ResponseMeta struct {
	// Ticker symbol of instrument
	Symbol string `json:"symbol"`
	// Name of symbol
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

type _GetDirectHolders200ResponseMeta GetDirectHolders200ResponseMeta

// NewGetDirectHolders200ResponseMeta instantiates a new GetDirectHolders200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDirectHolders200ResponseMeta(symbol string, name string, currency string, exchange string, micCode string, exchangeTimezone string) *GetDirectHolders200ResponseMeta {
	this := GetDirectHolders200ResponseMeta{}
	this.Symbol = symbol
	this.Name = name
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.ExchangeTimezone = exchangeTimezone
	return &this
}

// NewGetDirectHolders200ResponseMetaWithDefaults instantiates a new GetDirectHolders200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDirectHolders200ResponseMetaWithDefaults() *GetDirectHolders200ResponseMeta {
	this := GetDirectHolders200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetDirectHolders200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetDirectHolders200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetDirectHolders200ResponseMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200ResponseMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetDirectHolders200ResponseMeta) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *GetDirectHolders200ResponseMeta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetDirectHolders200ResponseMeta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *GetDirectHolders200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetDirectHolders200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetDirectHolders200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetDirectHolders200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *GetDirectHolders200ResponseMeta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *GetDirectHolders200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *GetDirectHolders200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

func (o GetDirectHolders200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDirectHolders200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	return toSerialize, nil
}

func (o *GetDirectHolders200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
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

	varGetDirectHolders200ResponseMeta := _GetDirectHolders200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetDirectHolders200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetDirectHolders200ResponseMeta(varGetDirectHolders200ResponseMeta)

	return err
}

type NullableGetDirectHolders200ResponseMeta struct {
	value *GetDirectHolders200ResponseMeta
	isSet bool
}

func (v NullableGetDirectHolders200ResponseMeta) Get() *GetDirectHolders200ResponseMeta {
	return v.value
}

func (v *NullableGetDirectHolders200ResponseMeta) Set(val *GetDirectHolders200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDirectHolders200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDirectHolders200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDirectHolders200ResponseMeta(val *GetDirectHolders200ResponseMeta) *NullableGetDirectHolders200ResponseMeta {
	return &NullableGetDirectHolders200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetDirectHolders200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDirectHolders200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
