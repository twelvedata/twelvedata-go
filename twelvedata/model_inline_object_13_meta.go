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

// checks if the InlineObject13Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject13Meta{}

// InlineObject13Meta struct for InlineObject13Meta
type InlineObject13Meta struct {
	// The ticker symbol of an instrument for which data was requested.
	Symbol string `json:"symbol"`
	// The time gap between consecutive data points.
	Interval string `json:"interval"`
	// The currency of a traded instrument.
	Currency string `json:"currency"`
	// The timezone of the exchange where the instrument is traded.
	ExchangeTimezone string `json:"exchange_timezone"`
	// The exchange name where the instrument is traded.
	Exchange string `json:"exchange"`
	// The Market Identifier Code (MIC) of the exchange where the instrument is traded.
	MicCode string `json:"mic_code"`
	// The asset class to which the instrument belongs.
	Type      string                       `json:"type"`
	Indicator *InlineObject13MetaIndicator `json:"indicator,omitempty"`
}

type _InlineObject13Meta InlineObject13Meta

// NewInlineObject13Meta instantiates a new InlineObject13Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject13Meta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string) *InlineObject13Meta {
	this := InlineObject13Meta{}
	this.Symbol = symbol
	this.Interval = interval
	this.Currency = currency
	this.ExchangeTimezone = exchangeTimezone
	this.Exchange = exchange
	this.MicCode = micCode
	this.Type = type_
	return &this
}

// NewInlineObject13MetaWithDefaults instantiates a new InlineObject13Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject13MetaWithDefaults() *InlineObject13Meta {
	this := InlineObject13Meta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *InlineObject13Meta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *InlineObject13Meta) SetSymbol(v string) {
	o.Symbol = v
}

// GetInterval returns the Interval field value
func (o *InlineObject13Meta) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *InlineObject13Meta) SetInterval(v string) {
	o.Interval = v
}

// GetCurrency returns the Currency field value
func (o *InlineObject13Meta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InlineObject13Meta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *InlineObject13Meta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *InlineObject13Meta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

// GetExchange returns the Exchange field value
func (o *InlineObject13Meta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *InlineObject13Meta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *InlineObject13Meta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *InlineObject13Meta) SetMicCode(v string) {
	o.MicCode = v
}

// GetType returns the Type field value
func (o *InlineObject13Meta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InlineObject13Meta) SetType(v string) {
	o.Type = v
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *InlineObject13Meta) GetIndicator() InlineObject13MetaIndicator {
	if o == nil || IsNil(o.Indicator) {
		var ret InlineObject13MetaIndicator
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13Meta) GetIndicatorOk() (*InlineObject13MetaIndicator, bool) {
	if o == nil || IsNil(o.Indicator) {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *InlineObject13Meta) HasIndicator() bool {
	if o != nil && !IsNil(o.Indicator) {
		return true
	}

	return false
}

// SetIndicator gets a reference to the given InlineObject13MetaIndicator and assigns it to the Indicator field.
func (o *InlineObject13Meta) SetIndicator(v InlineObject13MetaIndicator) {
	o.Indicator = &v
}

func (o InlineObject13Meta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject13Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["interval"] = o.Interval
	toSerialize["currency"] = o.Currency
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["type"] = o.Type
	if !IsNil(o.Indicator) {
		toSerialize["indicator"] = o.Indicator
	}
	return toSerialize, nil
}

func (o *InlineObject13Meta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"interval",
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

	varInlineObject13Meta := _InlineObject13Meta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineObject13Meta)

	if err != nil {
		return err
	}

	*o = InlineObject13Meta(varInlineObject13Meta)

	return err
}

type NullableInlineObject13Meta struct {
	value *InlineObject13Meta
	isSet bool
}

func (v NullableInlineObject13Meta) Get() *InlineObject13Meta {
	return v.value
}

func (v *NullableInlineObject13Meta) Set(val *InlineObject13Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject13Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject13Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject13Meta(val *InlineObject13Meta) *NullableInlineObject13Meta {
	return &NullableInlineObject13Meta{value: val, isSet: true}
}

func (v NullableInlineObject13Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject13Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
