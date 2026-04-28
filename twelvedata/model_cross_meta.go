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

// checks if the CrossMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrossMeta{}

// CrossMeta Json object with request general information
type CrossMeta struct {
	// Base instrument symbol
	BaseInstrument string `json:"base_instrument"`
	// Base currency
	BaseCurrency string `json:"base_currency"`
	// Base exchange
	BaseExchange string `json:"base_exchange"`
	// Interval between two consecutive points in time series
	Interval string `json:"interval"`
	// Quote instrument symbol
	QuoteInstrument string `json:"quote_instrument"`
	// Quote currency
	QuoteCurrency string `json:"quote_currency"`
	// Quote exchange
	QuoteExchange string `json:"quote_exchange"`
}

type _CrossMeta CrossMeta

// NewCrossMeta instantiates a new CrossMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossMeta(baseInstrument string, baseCurrency string, baseExchange string, interval string, quoteInstrument string, quoteCurrency string, quoteExchange string) *CrossMeta {
	this := CrossMeta{}
	this.BaseInstrument = baseInstrument
	this.BaseCurrency = baseCurrency
	this.BaseExchange = baseExchange
	this.Interval = interval
	this.QuoteInstrument = quoteInstrument
	this.QuoteCurrency = quoteCurrency
	this.QuoteExchange = quoteExchange
	return &this
}

// NewCrossMetaWithDefaults instantiates a new CrossMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossMetaWithDefaults() *CrossMeta {
	this := CrossMeta{}
	return &this
}

// GetBaseInstrument returns the BaseInstrument field value
func (o *CrossMeta) GetBaseInstrument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseInstrument
}

// GetBaseInstrumentOk returns a tuple with the BaseInstrument field value
// and a boolean to check if the value has been set.
func (o *CrossMeta) GetBaseInstrumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseInstrument, true
}

// SetBaseInstrument sets field value
func (o *CrossMeta) SetBaseInstrument(v string) {
	o.BaseInstrument = v
}

// GetBaseCurrency returns the BaseCurrency field value
func (o *CrossMeta) GetBaseCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseCurrency
}

// GetBaseCurrencyOk returns a tuple with the BaseCurrency field value
// and a boolean to check if the value has been set.
func (o *CrossMeta) GetBaseCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseCurrency, true
}

// SetBaseCurrency sets field value
func (o *CrossMeta) SetBaseCurrency(v string) {
	o.BaseCurrency = v
}

// GetBaseExchange returns the BaseExchange field value
func (o *CrossMeta) GetBaseExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseExchange
}

// GetBaseExchangeOk returns a tuple with the BaseExchange field value
// and a boolean to check if the value has been set.
func (o *CrossMeta) GetBaseExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseExchange, true
}

// SetBaseExchange sets field value
func (o *CrossMeta) SetBaseExchange(v string) {
	o.BaseExchange = v
}

// GetInterval returns the Interval field value
func (o *CrossMeta) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *CrossMeta) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *CrossMeta) SetInterval(v string) {
	o.Interval = v
}

// GetQuoteInstrument returns the QuoteInstrument field value
func (o *CrossMeta) GetQuoteInstrument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteInstrument
}

// GetQuoteInstrumentOk returns a tuple with the QuoteInstrument field value
// and a boolean to check if the value has been set.
func (o *CrossMeta) GetQuoteInstrumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteInstrument, true
}

// SetQuoteInstrument sets field value
func (o *CrossMeta) SetQuoteInstrument(v string) {
	o.QuoteInstrument = v
}

// GetQuoteCurrency returns the QuoteCurrency field value
func (o *CrossMeta) GetQuoteCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteCurrency
}

// GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field value
// and a boolean to check if the value has been set.
func (o *CrossMeta) GetQuoteCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteCurrency, true
}

// SetQuoteCurrency sets field value
func (o *CrossMeta) SetQuoteCurrency(v string) {
	o.QuoteCurrency = v
}

// GetQuoteExchange returns the QuoteExchange field value
func (o *CrossMeta) GetQuoteExchange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuoteExchange
}

// GetQuoteExchangeOk returns a tuple with the QuoteExchange field value
// and a boolean to check if the value has been set.
func (o *CrossMeta) GetQuoteExchangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuoteExchange, true
}

// SetQuoteExchange sets field value
func (o *CrossMeta) SetQuoteExchange(v string) {
	o.QuoteExchange = v
}

func (o CrossMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrossMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_instrument"] = o.BaseInstrument
	toSerialize["base_currency"] = o.BaseCurrency
	toSerialize["base_exchange"] = o.BaseExchange
	toSerialize["interval"] = o.Interval
	toSerialize["quote_instrument"] = o.QuoteInstrument
	toSerialize["quote_currency"] = o.QuoteCurrency
	toSerialize["quote_exchange"] = o.QuoteExchange
	return toSerialize, nil
}

func (o *CrossMeta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base_instrument",
		"base_currency",
		"base_exchange",
		"interval",
		"quote_instrument",
		"quote_currency",
		"quote_exchange",
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

	varCrossMeta := _CrossMeta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCrossMeta)

	if err != nil {
		return err
	}

	*o = CrossMeta(varCrossMeta)

	return err
}

type NullableCrossMeta struct {
	value *CrossMeta
	isSet bool
}

func (v NullableCrossMeta) Get() *CrossMeta {
	return v.value
}

func (v *NullableCrossMeta) Set(val *CrossMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossMeta(val *CrossMeta) *NullableCrossMeta {
	return &NullableCrossMeta{value: val, isSet: true}
}

func (v NullableCrossMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
