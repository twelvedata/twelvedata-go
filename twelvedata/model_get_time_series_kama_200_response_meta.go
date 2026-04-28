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

// checks if the GetTimeSeriesKama200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKama200ResponseMeta{}

// GetTimeSeriesKama200ResponseMeta Json object with request general information
type GetTimeSeriesKama200ResponseMeta struct {
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
	Type      string                                    `json:"type"`
	Indicator GetTimeSeriesKama200ResponseMetaIndicator `json:"indicator"`
}

type _GetTimeSeriesKama200ResponseMeta GetTimeSeriesKama200ResponseMeta

// NewGetTimeSeriesKama200ResponseMeta instantiates a new GetTimeSeriesKama200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKama200ResponseMeta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string, indicator GetTimeSeriesKama200ResponseMetaIndicator) *GetTimeSeriesKama200ResponseMeta {
	this := GetTimeSeriesKama200ResponseMeta{}
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

// NewGetTimeSeriesKama200ResponseMetaWithDefaults instantiates a new GetTimeSeriesKama200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKama200ResponseMetaWithDefaults() *GetTimeSeriesKama200ResponseMeta {
	this := GetTimeSeriesKama200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetTimeSeriesKama200ResponseMeta) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetSymbol(v string) {
	o.Symbol = v
}

// GetInterval returns the Interval field value
func (o *GetTimeSeriesKama200ResponseMeta) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetInterval(v string) {
	o.Interval = v
}

// GetCurrency returns the Currency field value
func (o *GetTimeSeriesKama200ResponseMeta) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetCurrency(v string) {
	o.Currency = v
}

// GetExchangeTimezone returns the ExchangeTimezone field value
func (o *GetTimeSeriesKama200ResponseMeta) GetExchangeTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExchangeTimezone, true
}

// SetExchangeTimezone sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = v
}

// GetExchange returns the Exchange field value
func (o *GetTimeSeriesKama200ResponseMeta) GetExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exchange, true
}

// SetExchange sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetExchange(v string) {
	o.Exchange = v
}

// GetMicCode returns the MicCode field value
func (o *GetTimeSeriesKama200ResponseMeta) GetMicCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MicCode, true
}

// SetMicCode sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetMicCode(v string) {
	o.MicCode = v
}

// GetType returns the Type field value
func (o *GetTimeSeriesKama200ResponseMeta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetType(v string) {
	o.Type = v
}

// GetIndicator returns the Indicator field value
func (o *GetTimeSeriesKama200ResponseMeta) GetIndicator() GetTimeSeriesKama200ResponseMetaIndicator {
	if o == nil {
		var ret GetTimeSeriesKama200ResponseMetaIndicator
		return ret
	}

	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesKama200ResponseMetaIndicator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value
func (o *GetTimeSeriesKama200ResponseMeta) SetIndicator(v GetTimeSeriesKama200ResponseMetaIndicator) {
	o.Indicator = v
}

func (o GetTimeSeriesKama200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKama200ResponseMeta) ToMap() (map[string]interface{}, error) {
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

func (o *GetTimeSeriesKama200ResponseMeta) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesKama200ResponseMeta := _GetTimeSeriesKama200ResponseMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesKama200ResponseMeta)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKama200ResponseMeta(varGetTimeSeriesKama200ResponseMeta)

	return err
}

type NullableGetTimeSeriesKama200ResponseMeta struct {
	value *GetTimeSeriesKama200ResponseMeta
	isSet bool
}

func (v NullableGetTimeSeriesKama200ResponseMeta) Get() *GetTimeSeriesKama200ResponseMeta {
	return v.value
}

func (v *NullableGetTimeSeriesKama200ResponseMeta) Set(val *GetTimeSeriesKama200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKama200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKama200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKama200ResponseMeta(val *GetTimeSeriesKama200ResponseMeta) *NullableGetTimeSeriesKama200ResponseMeta {
	return &NullableGetTimeSeriesKama200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKama200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKama200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
