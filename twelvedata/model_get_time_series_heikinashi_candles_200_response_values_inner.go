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

// checks if the GetTimeSeriesHeikinashiCandles200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHeikinashiCandles200ResponseValuesInner{}

// GetTimeSeriesHeikinashiCandles200ResponseValuesInner struct for GetTimeSeriesHeikinashiCandles200ResponseValuesInner
type GetTimeSeriesHeikinashiCandles200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Heikin-Ashi highs value
	Heikinhighs string `json:"heikinhighs"`
	// Heikin-Ashi opens value
	Heikinopens string `json:"heikinopens"`
	// Heikin-Ashi closes value
	Heikincloses string `json:"heikincloses"`
	// Heikin-Ashi lows value
	Heikinlows string `json:"heikinlows"`
}

type _GetTimeSeriesHeikinashiCandles200ResponseValuesInner GetTimeSeriesHeikinashiCandles200ResponseValuesInner

// NewGetTimeSeriesHeikinashiCandles200ResponseValuesInner instantiates a new GetTimeSeriesHeikinashiCandles200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHeikinashiCandles200ResponseValuesInner(datetime string, heikinhighs string, heikinopens string, heikincloses string, heikinlows string) *GetTimeSeriesHeikinashiCandles200ResponseValuesInner {
	this := GetTimeSeriesHeikinashiCandles200ResponseValuesInner{}
	this.Datetime = datetime
	this.Heikinhighs = heikinhighs
	this.Heikinopens = heikinopens
	this.Heikincloses = heikincloses
	this.Heikinlows = heikinlows
	return &this
}

// NewGetTimeSeriesHeikinashiCandles200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHeikinashiCandles200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHeikinashiCandles200ResponseValuesInnerWithDefaults() *GetTimeSeriesHeikinashiCandles200ResponseValuesInner {
	this := GetTimeSeriesHeikinashiCandles200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetHeikinhighs returns the Heikinhighs field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikinhighs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikinhighs
}

// GetHeikinhighsOk returns a tuple with the Heikinhighs field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikinhighsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikinhighs, true
}

// SetHeikinhighs sets field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) SetHeikinhighs(v string) {
	o.Heikinhighs = v
}

// GetHeikinopens returns the Heikinopens field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikinopens() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikinopens
}

// GetHeikinopensOk returns a tuple with the Heikinopens field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikinopensOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikinopens, true
}

// SetHeikinopens sets field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) SetHeikinopens(v string) {
	o.Heikinopens = v
}

// GetHeikincloses returns the Heikincloses field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikincloses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikincloses
}

// GetHeikinclosesOk returns a tuple with the Heikincloses field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikinclosesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikincloses, true
}

// SetHeikincloses sets field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) SetHeikincloses(v string) {
	o.Heikincloses = v
}

// GetHeikinlows returns the Heikinlows field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikinlows() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikinlows
}

// GetHeikinlowsOk returns a tuple with the Heikinlows field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) GetHeikinlowsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikinlows, true
}

// SetHeikinlows sets field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) SetHeikinlows(v string) {
	o.Heikinlows = v
}

func (o GetTimeSeriesHeikinashiCandles200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHeikinashiCandles200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["heikinhighs"] = o.Heikinhighs
	toSerialize["heikinopens"] = o.Heikinopens
	toSerialize["heikincloses"] = o.Heikincloses
	toSerialize["heikinlows"] = o.Heikinlows
	return toSerialize, nil
}

func (o *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"heikinhighs",
		"heikinopens",
		"heikincloses",
		"heikinlows",
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

	varGetTimeSeriesHeikinashiCandles200ResponseValuesInner := _GetTimeSeriesHeikinashiCandles200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesHeikinashiCandles200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHeikinashiCandles200ResponseValuesInner(varGetTimeSeriesHeikinashiCandles200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner struct {
	value *GetTimeSeriesHeikinashiCandles200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner) Get() *GetTimeSeriesHeikinashiCandles200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner) Set(val *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner(val *GetTimeSeriesHeikinashiCandles200ResponseValuesInner) *NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner {
	return &NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHeikinashiCandles200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
