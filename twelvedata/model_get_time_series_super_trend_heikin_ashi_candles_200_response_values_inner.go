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

// checks if the GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner{}

// GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner struct for GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner
type GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SuperTrend value
	Supertrend string `json:"supertrend"`
	// Heikin-Ashi high values
	Heikinhighs string `json:"heikinhighs"`
	// Heikin-Ashi open values
	Heikinopens string `json:"heikinopens"`
	// Heikin-Ashi close values
	Heikincloses string `json:"heikincloses"`
	// Heikin-Ashi low values
	Heikinlows string `json:"heikinlows"`
}

type _GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner

// NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner(datetime string, supertrend string, heikinhighs string, heikinopens string, heikincloses string, heikinlows string) *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner {
	this := GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner{}
	this.Datetime = datetime
	this.Supertrend = supertrend
	this.Heikinhighs = heikinhighs
	this.Heikinopens = heikinopens
	this.Heikincloses = heikincloses
	this.Heikinlows = heikinlows
	return &this
}

// NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInnerWithDefaults() *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner {
	this := GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSupertrend returns the Supertrend field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetSupertrend() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supertrend
}

// GetSupertrendOk returns a tuple with the Supertrend field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetSupertrendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supertrend, true
}

// SetSupertrend sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) SetSupertrend(v string) {
	o.Supertrend = v
}

// GetHeikinhighs returns the Heikinhighs field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikinhighs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikinhighs
}

// GetHeikinhighsOk returns a tuple with the Heikinhighs field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikinhighsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikinhighs, true
}

// SetHeikinhighs sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) SetHeikinhighs(v string) {
	o.Heikinhighs = v
}

// GetHeikinopens returns the Heikinopens field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikinopens() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikinopens
}

// GetHeikinopensOk returns a tuple with the Heikinopens field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikinopensOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikinopens, true
}

// SetHeikinopens sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) SetHeikinopens(v string) {
	o.Heikinopens = v
}

// GetHeikincloses returns the Heikincloses field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikincloses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikincloses
}

// GetHeikinclosesOk returns a tuple with the Heikincloses field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikinclosesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikincloses, true
}

// SetHeikincloses sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) SetHeikincloses(v string) {
	o.Heikincloses = v
}

// GetHeikinlows returns the Heikinlows field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikinlows() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Heikinlows
}

// GetHeikinlowsOk returns a tuple with the Heikinlows field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) GetHeikinlowsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Heikinlows, true
}

// SetHeikinlows sets field value
func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) SetHeikinlows(v string) {
	o.Heikinlows = v
}

func (o GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["supertrend"] = o.Supertrend
	toSerialize["heikinhighs"] = o.Heikinhighs
	toSerialize["heikinopens"] = o.Heikinopens
	toSerialize["heikincloses"] = o.Heikincloses
	toSerialize["heikinlows"] = o.Heikinlows
	return toSerialize, nil
}

func (o *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"supertrend",
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

	varGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner := _GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner(varGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner struct {
	value *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) Get() *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) Set(val *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner(val *GetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner {
	return &NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSuperTrendHeikinAshiCandles200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
