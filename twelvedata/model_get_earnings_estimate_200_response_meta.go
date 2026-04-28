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

// checks if the GetEarningsEstimate200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarningsEstimate200ResponseMeta{}

// GetEarningsEstimate200ResponseMeta Meta information about the symbol
type GetEarningsEstimate200ResponseMeta struct {
	// Symbol ticker of the instrument
	Symbol string `json:"symbol"`
	// Name of the instrument
	Name string `json:"name"`
	// Currency in which the instrument is traded
	Currency string `json:"currency"`
	// Timezone of the exchange
	ExchangeTimezone string `json:"exchange_timezone"`
	// Exchange where the instrument is traded
	Exchange string `json:"exchange"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode string `json:"mic_code"`
	// Type of the instrument
	Type string `json:"type"`
}

type _GetEarningsEstimate200ResponseMeta GetEarningsEstimate200ResponseMeta

// NewGetEarningsEstimate200ResponseMeta instantiates a new GetEarningsEstimate200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarningsEstimate200ResponseMeta(symbol string, name string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string) *GetEarningsEstimate200ResponseMeta {
	this := GetEarningsEstimate200ResponseMeta{}
	this.Symbol = symbol
	this.Name = name
	this.Currency = currency
	this.ExchangeTimezone = exchangeTimezone
	this.Exchange = exchange
	this.MicCode = micCode
	this.Type = type_
	return &this
}

// NewGetEarningsEstimate200ResponseMetaWithDefaults instantiates a new GetEarningsEstimate200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarningsEstimate200ResponseMetaWithDefaults() *GetEarningsEstimate200ResponseMeta {
	this := GetEarningsEstimate200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetEarningsEstimate200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetEarningsEstimate200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *GetEarningsEstimate200ResponseMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetEarningsEstimate200ResponseMeta) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *GetEarningsEstimate200ResponseMeta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetEarningsEstimate200ResponseMeta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *GetEarningsEstimate200ResponseMeta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *GetEarningsEstimate200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

// GetExchange returns the Exchange field value
func (o *GetEarningsEstimate200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetEarningsEstimate200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetEarningsEstimate200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetEarningsEstimate200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetType returns the Type field value
func (o *GetEarningsEstimate200ResponseMeta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseMeta) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetEarningsEstimate200ResponseMeta) SetType(v string) {
	o.Type = v
}

func (o GetEarningsEstimate200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarningsEstimate200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *GetEarningsEstimate200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"currency",
		"exchange_timezone",
		"exchange",
		"mic_code",
		"type",
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

	varGetEarningsEstimate200ResponseMeta := _GetEarningsEstimate200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEarningsEstimate200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetEarningsEstimate200ResponseMeta(varGetEarningsEstimate200ResponseMeta)

	return err
}

type NullableGetEarningsEstimate200ResponseMeta struct {
	value *GetEarningsEstimate200ResponseMeta
	isSet bool
}

func (v NullableGetEarningsEstimate200ResponseMeta) Get() *GetEarningsEstimate200ResponseMeta {
	return v.value
}

func (v *NullableGetEarningsEstimate200ResponseMeta) Set(val *GetEarningsEstimate200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarningsEstimate200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarningsEstimate200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarningsEstimate200ResponseMeta(val *GetEarningsEstimate200ResponseMeta) *NullableGetEarningsEstimate200ResponseMeta {
	return &NullableGetEarningsEstimate200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetEarningsEstimate200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarningsEstimate200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
