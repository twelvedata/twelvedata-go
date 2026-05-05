// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetEarnings200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarnings200ResponseMeta{}

// GetEarnings200ResponseMeta Metadata about the instrument
type GetEarnings200ResponseMeta struct {
	// Symbol of the instrument
	Symbol string `json:"symbol"`
	// Name of the company
	Name string `json:"name"`
	// Currency in which the instrument is traded
	Currency string `json:"currency"`
	// Exchange where the instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Timezone of the exchange
	ExchangeTimezone string `json:"exchange_timezone"`
}

type _GetEarnings200ResponseMeta GetEarnings200ResponseMeta

// NewGetEarnings200ResponseMeta instantiates a new GetEarnings200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarnings200ResponseMeta(symbol string, name string, currency string, exchange string, micCode string, exchangeTimezone string) *GetEarnings200ResponseMeta {
	this := GetEarnings200ResponseMeta{}
	this.Symbol = symbol
	this.Name = name
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.ExchangeTimezone = exchangeTimezone
	return &this
}

// NewGetEarnings200ResponseMetaWithDefaults instantiates a new GetEarnings200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarnings200ResponseMetaWithDefaults() *GetEarnings200ResponseMeta {
	this := GetEarnings200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetEarnings200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetEarnings200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetEarnings200ResponseMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200ResponseMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetEarnings200ResponseMeta) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *GetEarnings200ResponseMeta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetEarnings200ResponseMeta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *GetEarnings200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetEarnings200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetEarnings200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetEarnings200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *GetEarnings200ResponseMeta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *GetEarnings200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *GetEarnings200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

func (o GetEarnings200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarnings200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	return toSerialize, nil
}

func (o *GetEarnings200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
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

	varGetEarnings200ResponseMeta := _GetEarnings200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEarnings200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetEarnings200ResponseMeta(varGetEarnings200ResponseMeta)

	return err
}

type NullableGetEarnings200ResponseMeta struct {
	value *GetEarnings200ResponseMeta
	isSet bool
}

func (v NullableGetEarnings200ResponseMeta) Get() *GetEarnings200ResponseMeta {
	return v.value
}

func (v *NullableGetEarnings200ResponseMeta) Set(val *GetEarnings200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarnings200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarnings200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarnings200ResponseMeta(val *GetEarnings200ResponseMeta) *NullableGetEarnings200ResponseMeta {
	return &NullableGetEarnings200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetEarnings200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarnings200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
