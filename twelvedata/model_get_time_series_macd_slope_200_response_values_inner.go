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

// checks if the GetTimeSeriesMacdSlope200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacdSlope200ResponseValuesInner{}

// GetTimeSeriesMacdSlope200ResponseValuesInner struct for GetTimeSeriesMacdSlope200ResponseValuesInner
type GetTimeSeriesMacdSlope200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// MACD slope value
	MacdSlope string `json:"macd_slope"`
	// MACD signal slope value
	MacdSignalSlope string `json:"macd_signal_slope"`
	// MACD histogram slope value
	MacdHistSlope string `json:"macd_hist_slope"`
}

type _GetTimeSeriesMacdSlope200ResponseValuesInner GetTimeSeriesMacdSlope200ResponseValuesInner

// NewGetTimeSeriesMacdSlope200ResponseValuesInner instantiates a new GetTimeSeriesMacdSlope200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacdSlope200ResponseValuesInner(datetime string, macdSlope string, macdSignalSlope string, macdHistSlope string) *GetTimeSeriesMacdSlope200ResponseValuesInner {
	this := GetTimeSeriesMacdSlope200ResponseValuesInner{}
	this.Datetime = datetime
	this.MacdSlope = macdSlope
	this.MacdSignalSlope = macdSignalSlope
	this.MacdHistSlope = macdHistSlope
	return &this
}

// NewGetTimeSeriesMacdSlope200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMacdSlope200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacdSlope200ResponseValuesInnerWithDefaults() *GetTimeSeriesMacdSlope200ResponseValuesInner {
	this := GetTimeSeriesMacdSlope200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMacdSlope returns the MacdSlope field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSlope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacdSlope
}

// GetMacdSlopeOk returns a tuple with the MacdSlope field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSlopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacdSlope, true
}

// SetMacdSlope sets field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetMacdSlope(v string) {
	o.MacdSlope = v
}

// GetMacdSignalSlope returns the MacdSignalSlope field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSignalSlope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacdSignalSlope
}

// GetMacdSignalSlopeOk returns a tuple with the MacdSignalSlope field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSignalSlopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacdSignalSlope, true
}

// SetMacdSignalSlope sets field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetMacdSignalSlope(v string) {
	o.MacdSignalSlope = v
}

// GetMacdHistSlope returns the MacdHistSlope field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdHistSlope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacdHistSlope
}

// GetMacdHistSlopeOk returns a tuple with the MacdHistSlope field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdHistSlopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacdHistSlope, true
}

// SetMacdHistSlope sets field value
func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetMacdHistSlope(v string) {
	o.MacdHistSlope = v
}

func (o GetTimeSeriesMacdSlope200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacdSlope200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["macd_slope"] = o.MacdSlope
	toSerialize["macd_signal_slope"] = o.MacdSignalSlope
	toSerialize["macd_hist_slope"] = o.MacdHistSlope
	return toSerialize, nil
}

func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"macd_slope",
		"macd_signal_slope",
		"macd_hist_slope",
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

	varGetTimeSeriesMacdSlope200ResponseValuesInner := _GetTimeSeriesMacdSlope200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMacdSlope200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacdSlope200ResponseValuesInner(varGetTimeSeriesMacdSlope200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMacdSlope200ResponseValuesInner struct {
	value *GetTimeSeriesMacdSlope200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMacdSlope200ResponseValuesInner) Get() *GetTimeSeriesMacdSlope200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMacdSlope200ResponseValuesInner) Set(val *GetTimeSeriesMacdSlope200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacdSlope200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacdSlope200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacdSlope200ResponseValuesInner(val *GetTimeSeriesMacdSlope200ResponseValuesInner) *NullableGetTimeSeriesMacdSlope200ResponseValuesInner {
	return &NullableGetTimeSeriesMacdSlope200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacdSlope200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacdSlope200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
