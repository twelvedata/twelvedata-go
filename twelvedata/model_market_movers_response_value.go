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

// checks if the MarketMoversResponseValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketMoversResponseValue{}

// MarketMoversResponseValue struct for MarketMoversResponseValue
type MarketMoversResponseValue struct {
	// The exchange symbol ticker
	Symbol string `json:"symbol"`
	// The official name of the instrument
	Name string `json:"name"`
	// Exchange where instrument is traded
	Exchange *string `json:"exchange,omitempty"`
	// Market identifier code (MIC) under ISO 10383 standard
	MicCode *string `json:"mic_code,omitempty"`
	// The last updated datetime timestamp
	Datetime string `json:"datetime"`
	// The latest available price for the symbol today
	Last float64 `json:"last"`
	// The highest price for the symbol today
	High float64 `json:"high"`
	// The lowest price for the symbol today
	Low float64 `json:"low"`
	// The trading volume of the symbol today
	Volume int64 `json:"volume"`
	// The value of the change since the previous day
	Change float64 `json:"change"`
	// The percentage change since the previous day
	PercentChange float64 `json:"percent_change"`
}

type _MarketMoversResponseValue MarketMoversResponseValue

// NewMarketMoversResponseValue instantiates a new MarketMoversResponseValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketMoversResponseValue(symbol string, name string, datetime string, last float64, high float64, low float64, volume int64, change float64, percentChange float64) *MarketMoversResponseValue {
	this := MarketMoversResponseValue{}
	this.Symbol = symbol
	this.Name = name
	this.Datetime = datetime
	this.Last = last
	this.High = high
	this.Low = low
	this.Volume = volume
	this.Change = change
	this.PercentChange = percentChange
	return &this
}

// NewMarketMoversResponseValueWithDefaults instantiates a new MarketMoversResponseValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketMoversResponseValueWithDefaults() *MarketMoversResponseValue {
	this := MarketMoversResponseValue{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *MarketMoversResponseValue) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *MarketMoversResponseValue) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *MarketMoversResponseValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MarketMoversResponseValue) SetName(v string) {
	o.Name = v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *MarketMoversResponseValue) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *MarketMoversResponseValue) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *MarketMoversResponseValue) SetExchange(v string) {
	o.Exchange = &v
}

// GetMicCode returns the MicCode field value if set, zero value otherwise.
func (o *MarketMoversResponseValue) GetMicCode() string {
	if o == nil || IsNil(o.MicCode) {
		var ret string
		return ret
	}
	return *o.MicCode
}

// GetMicCodeOk returns a tuple with the MicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetMicCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MicCode) {
		return nil, false
	}
	return o.MicCode, true
}

// HasMicCode returns a boolean if a field has been set.
func (o *MarketMoversResponseValue) HasMicCode() bool {
	if o != nil && !IsNil(o.MicCode) {
		return true
	}

	return false
}

// SetMicCode gets a reference to the given string and assigns it to the MicCode field.
func (o *MarketMoversResponseValue) SetMicCode(v string) {
	o.MicCode = &v
}

// GetDatetime returns the Datetime field value
func (o *MarketMoversResponseValue) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *MarketMoversResponseValue) SetDatetime(v string) {
	o.Datetime = v
}

// GetLast returns the Last field value
func (o *MarketMoversResponseValue) GetLast() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Last
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetLastOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last, true
}

// SetLast sets field value
func (o *MarketMoversResponseValue) SetLast(v float64) {
	o.Last = v
}

// GetHigh returns the High field value
func (o *MarketMoversResponseValue) GetHigh() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetHighOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value
func (o *MarketMoversResponseValue) SetHigh(v float64) {
	o.High = v
}

// GetLow returns the Low field value
func (o *MarketMoversResponseValue) GetLow() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetLowOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value
func (o *MarketMoversResponseValue) SetLow(v float64) {
	o.Low = v
}

// GetVolume returns the Volume field value
func (o *MarketMoversResponseValue) GetVolume() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *MarketMoversResponseValue) SetVolume(v int64) {
	o.Volume = v
}

// GetChange returns the Change field value
func (o *MarketMoversResponseValue) GetChange() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Change
}

// GetChangeOk returns a tuple with the Change field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetChangeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Change, true
}

// SetChange sets field value
func (o *MarketMoversResponseValue) SetChange(v float64) {
	o.Change = v
}

// GetPercentChange returns the PercentChange field value
func (o *MarketMoversResponseValue) GetPercentChange() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.PercentChange
}

// GetPercentChangeOk returns a tuple with the PercentChange field value
// and a boolean to check if the value has been set.
func (o *MarketMoversResponseValue) GetPercentChangeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentChange, true
}

// SetPercentChange sets field value
func (o *MarketMoversResponseValue) SetPercentChange(v float64) {
	o.PercentChange = v
}

func (o MarketMoversResponseValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketMoversResponseValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.MicCode) {
		toSerialize["mic_code"] = o.MicCode
	}
	toSerialize["datetime"] = o.Datetime
	toSerialize["last"] = o.Last
	toSerialize["high"] = o.High
	toSerialize["low"] = o.Low
	toSerialize["volume"] = o.Volume
	toSerialize["change"] = o.Change
	toSerialize["percent_change"] = o.PercentChange
	return toSerialize, nil
}

func (o *MarketMoversResponseValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"datetime",
		"last",
		"high",
		"low",
		"volume",
		"change",
		"percent_change",
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

	varMarketMoversResponseValue := _MarketMoversResponseValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMarketMoversResponseValue)

	if err != nil {
		return err
	}

	*o = MarketMoversResponseValue(varMarketMoversResponseValue)

	return err
}

type NullableMarketMoversResponseValue struct {
	value *MarketMoversResponseValue
	isSet bool
}

func (v NullableMarketMoversResponseValue) Get() *MarketMoversResponseValue {
	return v.value
}

func (v *NullableMarketMoversResponseValue) Set(val *MarketMoversResponseValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketMoversResponseValue) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketMoversResponseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketMoversResponseValue(val *MarketMoversResponseValue) *NullableMarketMoversResponseValue {
	return &NullableMarketMoversResponseValue{value: val, isSet: true}
}

func (v NullableMarketMoversResponseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketMoversResponseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
