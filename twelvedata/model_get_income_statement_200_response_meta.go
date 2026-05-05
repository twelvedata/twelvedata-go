// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetIncomeStatement200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIncomeStatement200ResponseMeta{}

// GetIncomeStatement200ResponseMeta Meta information about the instrument
type GetIncomeStatement200ResponseMeta struct {
	// Symbol ticker of the instrument
	Symbol string `json:"symbol"`
	// Name of the company
	Name string `json:"name"`
	// Currency of the instrument according to the ISO 4217 standard
	Currency string `json:"currency"`
	// Exchange where the instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Exchange timezone
	ExchangeTimezone string `json:"exchange_timezone"`
	// Period of the income statement (Annual or Quarterly)
	Period string `json:"period"`
}

type _GetIncomeStatement200ResponseMeta GetIncomeStatement200ResponseMeta

// NewGetIncomeStatement200ResponseMeta instantiates a new GetIncomeStatement200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIncomeStatement200ResponseMeta(symbol string, name string, currency string, exchange string, micCode string, exchangeTimezone string, period string) *GetIncomeStatement200ResponseMeta {
	this := GetIncomeStatement200ResponseMeta{}
	this.Symbol = symbol
	this.Name = name
	this.Currency = currency
	this.Exchange = exchange
	this.MicCode = micCode
	this.ExchangeTimezone = exchangeTimezone
	this.Period = period
	return &this
}

// NewGetIncomeStatement200ResponseMetaWithDefaults instantiates a new GetIncomeStatement200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIncomeStatement200ResponseMetaWithDefaults() *GetIncomeStatement200ResponseMeta {
	this := GetIncomeStatement200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetIncomeStatement200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetIncomeStatement200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetIncomeStatement200ResponseMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200ResponseMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetIncomeStatement200ResponseMeta) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *GetIncomeStatement200ResponseMeta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetIncomeStatement200ResponseMeta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchange returns the Exchange field value
func (o *GetIncomeStatement200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetIncomeStatement200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetIncomeStatement200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetIncomeStatement200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *GetIncomeStatement200ResponseMeta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *GetIncomeStatement200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

// GetPeriod returns the Period field value
func (o *GetIncomeStatement200ResponseMeta) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *GetIncomeStatement200ResponseMeta) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *GetIncomeStatement200ResponseMeta) SetPeriod(v string) {
	o.Period = v
}

func (o GetIncomeStatement200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIncomeStatement200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	toSerialize["period"] = o.Period
	return toSerialize, nil
}

func (o *GetIncomeStatement200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
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
		"period",
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

	varGetIncomeStatement200ResponseMeta := _GetIncomeStatement200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetIncomeStatement200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetIncomeStatement200ResponseMeta(varGetIncomeStatement200ResponseMeta)

	return err
}

type NullableGetIncomeStatement200ResponseMeta struct {
	value *GetIncomeStatement200ResponseMeta
	isSet bool
}

func (v NullableGetIncomeStatement200ResponseMeta) Get() *GetIncomeStatement200ResponseMeta {
	return v.value
}

func (v *NullableGetIncomeStatement200ResponseMeta) Set(val *GetIncomeStatement200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIncomeStatement200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIncomeStatement200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIncomeStatement200ResponseMeta(val *GetIncomeStatement200ResponseMeta) *NullableGetIncomeStatement200ResponseMeta {
	return &NullableGetIncomeStatement200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetIncomeStatement200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIncomeStatement200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
