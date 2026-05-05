// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetPriceTarget200ResponsePriceTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPriceTarget200ResponsePriceTarget{}

// GetPriceTarget200ResponsePriceTarget Price target information
type GetPriceTarget200ResponsePriceTarget struct {
	// Highest price target given by an analyst
	High *float64 `json:"high,omitempty"`
	// Median price target given across analysts
	Median *float64 `json:"median,omitempty"`
	// Lowest price target given by an analyst
	Low *float64 `json:"low,omitempty"`
	// Average price target given across analysts
	Average *float64 `json:"average,omitempty"`
	// Current price from of a security
	Current *float64 `json:"current,omitempty"`
	// Currency in which the price targets values are quoted
	Currency string `json:"currency"`
}

type _GetPriceTarget200ResponsePriceTarget GetPriceTarget200ResponsePriceTarget

// NewGetPriceTarget200ResponsePriceTarget instantiates a new GetPriceTarget200ResponsePriceTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPriceTarget200ResponsePriceTarget(currency string) *GetPriceTarget200ResponsePriceTarget {
	this := GetPriceTarget200ResponsePriceTarget{}
	this.Currency = currency
	return &this
}

// NewGetPriceTarget200ResponsePriceTargetWithDefaults instantiates a new GetPriceTarget200ResponsePriceTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPriceTarget200ResponsePriceTargetWithDefaults() *GetPriceTarget200ResponsePriceTarget {
	this := GetPriceTarget200ResponsePriceTarget{}
	return &this
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *GetPriceTarget200ResponsePriceTarget) GetHigh() float64 {
	if o == nil || IsNil(o.High) {
		var ret float64
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200ResponsePriceTarget) GetHighOk() (*float64, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *GetPriceTarget200ResponsePriceTarget) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given float64 and assigns it to the High field.
func (o *GetPriceTarget200ResponsePriceTarget) SetHigh(v float64) {
	o.High = &v
}

// GetMedian returns the Median field value if set, zero value otherwise.
func (o *GetPriceTarget200ResponsePriceTarget) GetMedian() float64 {
	if o == nil || IsNil(o.Median) {
		var ret float64
		return ret
	}
	return *o.Median
}

// GetMedianOk returns a tuple with the Median field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200ResponsePriceTarget) GetMedianOk() (*float64, bool) {
	if o == nil || IsNil(o.Median) {
		return nil, false
	}
	return o.Median, true
}

// HasMedian returns a boolean if a field has been set.
func (o *GetPriceTarget200ResponsePriceTarget) HasMedian() bool {
	if o != nil && !IsNil(o.Median) {
		return true
	}

	return false
}

// SetMedian gets a reference to the given float64 and assigns it to the Median field.
func (o *GetPriceTarget200ResponsePriceTarget) SetMedian(v float64) {
	o.Median = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *GetPriceTarget200ResponsePriceTarget) GetLow() float64 {
	if o == nil || IsNil(o.Low) {
		var ret float64
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200ResponsePriceTarget) GetLowOk() (*float64, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *GetPriceTarget200ResponsePriceTarget) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given float64 and assigns it to the Low field.
func (o *GetPriceTarget200ResponsePriceTarget) SetLow(v float64) {
	o.Low = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *GetPriceTarget200ResponsePriceTarget) GetAverage() float64 {
	if o == nil || IsNil(o.Average) {
		var ret float64
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200ResponsePriceTarget) GetAverageOk() (*float64, bool) {
	if o == nil || IsNil(o.Average) {
		return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *GetPriceTarget200ResponsePriceTarget) HasAverage() bool {
	if o != nil && !IsNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float64 and assigns it to the Average field.
func (o *GetPriceTarget200ResponsePriceTarget) SetAverage(v float64) {
	o.Average = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *GetPriceTarget200ResponsePriceTarget) GetCurrent() float64 {
	if o == nil || IsNil(o.Current) {
		var ret float64
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200ResponsePriceTarget) GetCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *GetPriceTarget200ResponsePriceTarget) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given float64 and assigns it to the Current field.
func (o *GetPriceTarget200ResponsePriceTarget) SetCurrent(v float64) {
	o.Current = &v
}

// GetCurrency returns the Currency field value
func (o *GetPriceTarget200ResponsePriceTarget) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetPriceTarget200ResponsePriceTarget) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetPriceTarget200ResponsePriceTarget) SetCurrency(v string) {
	o.Currency = v
}

func (o GetPriceTarget200ResponsePriceTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPriceTarget200ResponsePriceTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Median) {
		toSerialize["median"] = o.Median
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *GetPriceTarget200ResponsePriceTarget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
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

	varGetPriceTarget200ResponsePriceTarget := _GetPriceTarget200ResponsePriceTarget{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetPriceTarget200ResponsePriceTarget)

	if err != nil {
		return err
	}

	*o = GetPriceTarget200ResponsePriceTarget(varGetPriceTarget200ResponsePriceTarget)

	return err
}

type NullableGetPriceTarget200ResponsePriceTarget struct {
	value *GetPriceTarget200ResponsePriceTarget
	isSet bool
}

func (v NullableGetPriceTarget200ResponsePriceTarget) Get() *GetPriceTarget200ResponsePriceTarget {
	return v.value
}

func (v *NullableGetPriceTarget200ResponsePriceTarget) Set(val *GetPriceTarget200ResponsePriceTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPriceTarget200ResponsePriceTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPriceTarget200ResponsePriceTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPriceTarget200ResponsePriceTarget(val *GetPriceTarget200ResponsePriceTarget) *NullableGetPriceTarget200ResponsePriceTarget {
	return &NullableGetPriceTarget200ResponsePriceTarget{value: val, isSet: true}
}

func (v NullableGetPriceTarget200ResponsePriceTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPriceTarget200ResponsePriceTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
