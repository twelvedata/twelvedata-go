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

// checks if the GetTimeSeriesAdOsc200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdOsc200ResponseMeta{}

// GetTimeSeriesAdOsc200ResponseMeta Json object with request general information
type GetTimeSeriesAdOsc200ResponseMeta struct {
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
	Type      string                                     `json:"type"`
	Indicator GetTimeSeriesAdOsc200ResponseMetaIndicator `json:"indicator"`
}

type _GetTimeSeriesAdOsc200ResponseMeta GetTimeSeriesAdOsc200ResponseMeta

// NewGetTimeSeriesAdOsc200ResponseMeta instantiates a new GetTimeSeriesAdOsc200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdOsc200ResponseMeta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string, indicator GetTimeSeriesAdOsc200ResponseMetaIndicator) *GetTimeSeriesAdOsc200ResponseMeta {
	this := GetTimeSeriesAdOsc200ResponseMeta{}
	this.Symbol = symbol
	this.Interval = interval
	this.Currency = currency
	this.ExchangeTimezone = exchangeTimezone
	this.Exchange = exchange
	this.MicCode = micCode
	this.Type = type_
	this.Indicator = indicator
	return &this
}

// NewGetTimeSeriesAdOsc200ResponseMetaWithDefaults instantiates a new GetTimeSeriesAdOsc200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdOsc200ResponseMetaWithDefaults() *GetTimeSeriesAdOsc200ResponseMeta {
	this := GetTimeSeriesAdOsc200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetInterval returns the Interval field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetInterval(v string) {
	o.Interval = v
}

// GetCurrency returns the Currency field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

// GetExchange returns the Exchange field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetType returns the Type field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetType(v string) {
	o.Type = v
}

// GetIndicator returns the Indicator field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetIndicator() GetTimeSeriesAdOsc200ResponseMetaIndicator {
	if o == nil {
		var ret GetTimeSeriesAdOsc200ResponseMetaIndicator
		return ret
	}

	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesAdOsc200ResponseMetaIndicator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value
func (o *GetTimeSeriesAdOsc200ResponseMeta) SetIndicator(v GetTimeSeriesAdOsc200ResponseMetaIndicator) {
	o.Indicator = v
}

func (o GetTimeSeriesAdOsc200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdOsc200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["interval"] = o.Interval
	toSerialize["currency"] = o.Currency
	toSerialize["exchange_timezone"] = o.ExchangeTimezone
	toSerialize["exchange"] = o.Exchange
	toSerialize["mic_code"] = o.MicCode
	toSerialize["type"] = o.Type
	toSerialize["indicator"] = o.Indicator
	return toSerialize, nil
}

func (o *GetTimeSeriesAdOsc200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
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
		"indicator",
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

	varGetTimeSeriesAdOsc200ResponseMeta := _GetTimeSeriesAdOsc200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAdOsc200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdOsc200ResponseMeta(varGetTimeSeriesAdOsc200ResponseMeta)

	return err
}

type NullableGetTimeSeriesAdOsc200ResponseMeta struct {
	value *GetTimeSeriesAdOsc200ResponseMeta
	isSet bool
}

func (v NullableGetTimeSeriesAdOsc200ResponseMeta) Get() *GetTimeSeriesAdOsc200ResponseMeta {
	return v.value
}

func (v *NullableGetTimeSeriesAdOsc200ResponseMeta) Set(val *GetTimeSeriesAdOsc200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdOsc200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdOsc200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdOsc200ResponseMeta(val *GetTimeSeriesAdOsc200ResponseMeta) *NullableGetTimeSeriesAdOsc200ResponseMeta {
	return &NullableGetTimeSeriesAdOsc200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdOsc200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdOsc200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
