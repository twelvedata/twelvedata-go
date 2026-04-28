/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetDividends200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDividends200ResponseMeta{}

// GetDividends200ResponseMeta Json object with request general information
type GetDividends200ResponseMeta struct {
	// Ticker symbol of instrument
	Symbol *string `json:"symbol,omitempty"`
	// Name of symbol
	Name *string `json:"name,omitempty"`
	// Currency in which instrument is traded
	Currency *string `json:"currency,omitempty"`
	// Exchange where instrument is traded
	Exchange *string `json:"exchange,omitempty"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
	// Timezone of the exchange
	ExchangeTimezone *string `json:"exchange_timezone,omitempty"`
}

// NewGetDividends200ResponseMeta instantiates a new GetDividends200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDividends200ResponseMeta() *GetDividends200ResponseMeta {
	this := GetDividends200ResponseMeta{}
	return &this
}

// NewGetDividends200ResponseMetaWithDefaults instantiates a new GetDividends200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDividends200ResponseMetaWithDefaults() *GetDividends200ResponseMeta {
	this := GetDividends200ResponseMeta{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetDividends200ResponseMeta) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseMeta) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetDividends200ResponseMeta) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *GetDividends200ResponseMeta) SetSymbol(v string) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetDividends200ResponseMeta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseMeta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetDividends200ResponseMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetDividends200ResponseMeta) SetName(v string) {
	o.Name = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetDividends200ResponseMeta) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseMeta) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetDividends200ResponseMeta) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetDividends200ResponseMeta) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *GetDividends200ResponseMeta) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseMeta) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *GetDividends200ResponseMeta) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *GetDividends200ResponseMeta) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *GetDividends200ResponseMeta) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseMeta) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *GetDividends200ResponseMeta) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *GetDividends200ResponseMeta) SetMicCode(v string) {
	o.MicCode = &v
}

// GetExchangeTimezone returns the ExchangeTimezone field value if set, zero value otherwise.
func (o *GetDividends200ResponseMeta) GetExchangeTimezone() string {
	if o == nil || IsNil(o.ExchangeTimezone) {
		var ret string
		return ret
	}
	return *o.ExchangeTimezone
}

// GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDividends200ResponseMeta) GetExchangeTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeTimezone) {
		return nil, false
	}
	return o.ExchangeTimezone, true
}

// HasExchangeTimezone returns a boolean if a field has been set.
func (o *GetDividends200ResponseMeta) HasExchangeTimezone() bool {
	if o != nil && !IsNil(o.ExchangeTimezone) {
		return true
	}

	return false
}

// SetExchangeTimezone gets a reference to the given string and assigns it to the ExchangeTimezone field.
func (o *GetDividends200ResponseMeta) SetExchangeTimezone(v string) {
	o.ExchangeTimezone = &v
}

func (o GetDividends200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDividends200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	if !IsNil(o.ExchangeTimezone) {
		toSerialize["exchange_timezone"] = o.ExchangeTimezone
	}
	return toSerialize, nil
}

type NullableGetDividends200ResponseMeta struct {
	value *GetDividends200ResponseMeta
	isSet bool
}

func (v NullableGetDividends200ResponseMeta) Get() *GetDividends200ResponseMeta {
	return v.value
}

func (v *NullableGetDividends200ResponseMeta) Set(val *GetDividends200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDividends200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDividends200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDividends200ResponseMeta(val *GetDividends200ResponseMeta) *NullableGetDividends200ResponseMeta {
	return &NullableGetDividends200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetDividends200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDividends200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
