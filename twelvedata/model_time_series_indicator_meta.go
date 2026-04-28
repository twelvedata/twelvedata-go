/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the TimeSeriesIndicatorMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeSeriesIndicatorMeta{}

// TimeSeriesIndicatorMeta Common metadata fields for time series indicator responses
type TimeSeriesIndicatorMeta struct {
	// The ticker symbol of an instrument for which data was requested.
	Symbol *string `json:"symbol,omitempty"`
	// The time gap between consecutive data points.
	Interval *string `json:"interval,omitempty"`
	// The currency of a traded instrument.
	Currency *string `json:"currency,omitempty"`
	// The timezone of the exchange where the instrument is traded.
	ExchangeTimezone *string `json:"exchange_timezone,omitempty"`
	// The exchange name where the instrument is traded.
	Exchange *string `json:"exchange,omitempty"`
	// The Market Identifier Code (MIC) of the exchange where the instrument is traded.
	MicCode *string `json:"mic_code,omitempty"`
	// The asset class to which the instrument belongs.
	Type *string `json:"type,omitempty"`
}

// NewTimeSeriesIndicatorMeta instantiates a new TimeSeriesIndicatorMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSeriesIndicatorMeta() *TimeSeriesIndicatorMeta {
	this := TimeSeriesIndicatorMeta{}
	return &this
}

// NewTimeSeriesIndicatorMetaWithDefaults instantiates a new TimeSeriesIndicatorMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSeriesIndicatorMetaWithDefaults() *TimeSeriesIndicatorMeta {
	this := TimeSeriesIndicatorMeta{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TimeSeriesIndicatorMeta) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesIndicatorMeta) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TimeSeriesIndicatorMeta) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TimeSeriesIndicatorMeta) SetSymbol(v string) {
	o.Symbol = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *TimeSeriesIndicatorMeta) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesIndicatorMeta) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *TimeSeriesIndicatorMeta) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *TimeSeriesIndicatorMeta) SetInterval(v string) {
	o.Interval = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TimeSeriesIndicatorMeta) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesIndicatorMeta) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TimeSeriesIndicatorMeta) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TimeSeriesIndicatorMeta) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchangeTimezone returns the ExchangeTimezone field value if set, zero value otherwise.
func (o *TimeSeriesIndicatorMeta) GetExchangeTimezone() string {
	if o == nil || IsNil(o.ExchangeTimezone) {
		var ret string
		return ret
	}
	return *o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesIndicatorMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeTimezone) {
		return nil, false
	}
	return o.ExchangeTimezone, true
}

// HasExchangeTimezone returns a boolean if a field has been set.
func (o *TimeSeriesIndicatorMeta) HasExchangeTimezone() bool {
	if o != nil && !IsNil(o.ExchangeTimezone) {
		return true
	}

	return false
}

// SetExchangeTimezone gets a reference to the given string and assigns it to the ExchangeTimezone field.
func (o *TimeSeriesIndicatorMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *TimeSeriesIndicatorMeta) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesIndicatorMeta) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *TimeSeriesIndicatorMeta) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *TimeSeriesIndicatorMeta) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *TimeSeriesIndicatorMeta) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesIndicatorMeta) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *TimeSeriesIndicatorMeta) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *TimeSeriesIndicatorMeta) SetMicCode(v string) {
	o.MicCode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TimeSeriesIndicatorMeta) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSeriesIndicatorMeta) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TimeSeriesIndicatorMeta) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TimeSeriesIndicatorMeta) SetType(v string) {
	o.Type = &v
}

func (o TimeSeriesIndicatorMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeSeriesIndicatorMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExchangeTimezone) {
		toSerialize["exchange_timezone"] = o.ExchangeTimezone
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableTimeSeriesIndicatorMeta struct {
	value *TimeSeriesIndicatorMeta
	isSet bool
}

func (v NullableTimeSeriesIndicatorMeta) Get() *TimeSeriesIndicatorMeta {
	return v.value
}

func (v *NullableTimeSeriesIndicatorMeta) Set(val *TimeSeriesIndicatorMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSeriesIndicatorMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSeriesIndicatorMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSeriesIndicatorMeta(val *TimeSeriesIndicatorMeta) *NullableTimeSeriesIndicatorMeta {
	return &NullableTimeSeriesIndicatorMeta{value: val, isSet: true}
}

func (v NullableTimeSeriesIndicatorMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSeriesIndicatorMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
