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

// checks if the GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator{}

// GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The period used for calculation in the indicator
	Period int64 `json:"period"`
	// The multiplier used for calculation in the indicator
	Multiplier int64 `json:"multiplier"`
}

type _GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator

// NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator(name string, period int64, multiplier int64) *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator {
	this := GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator{}
	this.Name = name
	this.Period = period
	this.Multiplier = multiplier
	return &this
}

// NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator {
	this := GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetPeriod returns the Period field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) GetPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) GetPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) SetPeriod(v int64) {
	o.Period = v
}

// GetMultiplier returns the Multiplier field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) GetMultiplier() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) GetMultiplierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Multiplier, true
}

// SetMultiplier sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) SetMultiplier(v int64) {
	o.Multiplier = v
}

func (o GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["period"] = o.Period
	toSerialize["multiplier"] = o.Multiplier
	return toSerialize, nil
}

func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"period",
		"multiplier",
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

	varGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator := _GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator(varGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator struct {
	value *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) Get() *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) Set(val *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator(val *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator {
	return &NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
