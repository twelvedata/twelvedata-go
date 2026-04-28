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

// checks if the GetExchangeRate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExchangeRate200Response{}

// GetExchangeRate200Response struct for GetExchangeRate200Response
type GetExchangeRate200Response struct {
	// Requested currency symbol
	Symbol string `json:"symbol"`
	// Real-time exchange rate for the corresponding symbol
	Rate float64 `json:"rate"`
	// Unix timestamp of the rate
	Timestamp int64 `json:"timestamp"`
}

type _GetExchangeRate200Response GetExchangeRate200Response

// NewGetExchangeRate200Response instantiates a new GetExchangeRate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeRate200Response(symbol string, rate float64, timestamp int64) *GetExchangeRate200Response {
	this := GetExchangeRate200Response{}
	this.Symbol = symbol
	this.Rate = rate
	this.Timestamp = timestamp
	return &this
}

// NewGetExchangeRate200ResponseWithDefaults instantiates a new GetExchangeRate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeRate200ResponseWithDefaults() *GetExchangeRate200Response {
	this := GetExchangeRate200Response{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *GetExchangeRate200Response) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRate200Response) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *GetExchangeRate200Response) SetSymbol(v string) {
	o.Symbol = v
}

// GetRate returns the Rate field value
func (o *GetExchangeRate200Response) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRate200Response) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *GetExchangeRate200Response) SetRate(v float64) {
	o.Rate = v
}

// GetTimestamp returns the Timestamp field value
func (o *GetExchangeRate200Response) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRate200Response) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *GetExchangeRate200Response) SetTimestamp(v int64) {
	o.Timestamp = v
}

func (o GetExchangeRate200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeRate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["rate"] = o.Rate
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *GetExchangeRate200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetExchangeRate200Response := _GetExchangeRate200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetExchangeRate200Response)

	if err != nil {
		return err
	}

	*o = GetExchangeRate200Response(varGetExchangeRate200Response)

	return err
}

type NullableGetExchangeRate200Response struct {
	value *GetExchangeRate200Response
	isSet bool
}

func (v NullableGetExchangeRate200Response) Get() *GetExchangeRate200Response {
	return v.value
}

func (v *NullableGetExchangeRate200Response) Set(val *GetExchangeRate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeRate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeRate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeRate200Response(val *GetExchangeRate200Response) *NullableGetExchangeRate200Response {
	return &NullableGetExchangeRate200Response{value: val, isSet: true}
}

func (v NullableGetExchangeRate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeRate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
