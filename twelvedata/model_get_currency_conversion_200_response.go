// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetCurrencyConversion200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCurrencyConversion200Response{}

// GetCurrencyConversion200Response struct for GetCurrencyConversion200Response
type GetCurrencyConversion200Response struct {
	// Requested currency symbol
	Symbol string `json:"symbol"`
	// Real-time exchange rate for the corresponding symbol
	Rate float64 `json:"rate"`
	// Amount of converted currency
	Amount *float64 `json:"amount,omitempty"`
	// Unix timestamp of the rate
	Timestamp int64 `json:"timestamp"`
}

type _GetCurrencyConversion200Response GetCurrencyConversion200Response

// NewGetCurrencyConversion200Response instantiates a new GetCurrencyConversion200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrencyConversion200Response(symbol string, rate float64, timestamp int64) *GetCurrencyConversion200Response {
	this := GetCurrencyConversion200Response{}
	this.Symbol = symbol
	this.Rate = rate
	this.Timestamp = timestamp
	return &this
}

// NewGetCurrencyConversion200ResponseWithDefaults instantiates a new GetCurrencyConversion200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrencyConversion200ResponseWithDefaults() *GetCurrencyConversion200Response {
	this := GetCurrencyConversion200Response{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetCurrencyConversion200Response) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetCurrencyConversion200Response) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetCurrencyConversion200Response) SetSymbol(v string) {
	o.Symbol = v
}

// GetRate returns the Rate field value
func (o *GetCurrencyConversion200Response) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *GetCurrencyConversion200Response) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *GetCurrencyConversion200Response) SetRate(v float64) {
	o.Rate = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetCurrencyConversion200Response) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrencyConversion200Response) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetCurrencyConversion200Response) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *GetCurrencyConversion200Response) SetAmount(v float64) {
	o.Amount = &v
}

// GetTimestamp returns the Timestamp field value
func (o *GetCurrencyConversion200Response) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetCurrencyConversion200Response) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetCurrencyConversion200Response) SetTimestamp(v int64) {
	o.Timestamp = v
}

func (o GetCurrencyConversion200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrencyConversion200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["rate"] = o.Rate
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *GetCurrencyConversion200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"rate",
		"timestamp",
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

	varGetCurrencyConversion200Response := _GetCurrencyConversion200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetCurrencyConversion200Response)

	if err != nil {
		return err
	}

	*o = GetCurrencyConversion200Response(varGetCurrencyConversion200Response)

	return err
}

type NullableGetCurrencyConversion200Response struct {
	value *GetCurrencyConversion200Response
	isSet bool
}

func (v NullableGetCurrencyConversion200Response) Get() *GetCurrencyConversion200Response {
	return v.value
}

func (v *NullableGetCurrencyConversion200Response) Set(val *GetCurrencyConversion200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrencyConversion200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrencyConversion200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrencyConversion200Response(val *GetCurrencyConversion200Response) *NullableGetCurrencyConversion200Response {
	return &NullableGetCurrencyConversion200Response{value: val, isSet: true}
}

func (v NullableGetCurrencyConversion200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrencyConversion200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
