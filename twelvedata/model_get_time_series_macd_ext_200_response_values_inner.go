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

// checks if the GetTimeSeriesMacdExt200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacdExt200ResponseValuesInner{}

// GetTimeSeriesMacdExt200ResponseValuesInner struct for GetTimeSeriesMacdExt200ResponseValuesInner
type GetTimeSeriesMacdExt200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// MACD value
	Macd string `json:"macd"`
	// MACD signal line value
	MacdSignal string `json:"macd_signal"`
	// MACD histogram value
	MacdHist string `json:"macd_hist"`
}

type _GetTimeSeriesMacdExt200ResponseValuesInner GetTimeSeriesMacdExt200ResponseValuesInner

// NewGetTimeSeriesMacdExt200ResponseValuesInner instantiates a new GetTimeSeriesMacdExt200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacdExt200ResponseValuesInner(datetime string, macd string, macdSignal string, macdHist string) *GetTimeSeriesMacdExt200ResponseValuesInner {
	this := GetTimeSeriesMacdExt200ResponseValuesInner{}
	this.Datetime = datetime
	this.Macd = macd
	this.MacdSignal = macdSignal
	this.MacdHist = macdHist
	return &this
}

// NewGetTimeSeriesMacdExt200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMacdExt200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacdExt200ResponseValuesInnerWithDefaults() *GetTimeSeriesMacdExt200ResponseValuesInner {
	this := GetTimeSeriesMacdExt200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMacd returns the Macd field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Macd
}

// GetMacdOk returns a tuple with the Macd field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Macd, true
}

// SetMacd sets field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetMacd(v string) {
	o.Macd = v
}

// GetMacdSignal returns the MacdSignal field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdSignal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacdSignal
}

// GetMacdSignalOk returns a tuple with the MacdSignal field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdSignalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacdSignal, true
}

// SetMacdSignal sets field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetMacdSignal(v string) {
	o.MacdSignal = v
}

// GetMacdHist returns the MacdHist field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdHist() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacdHist
}

// GetMacdHistOk returns a tuple with the MacdHist field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdHistOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacdHist, true
}

// SetMacdHist sets field value
func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetMacdHist(v string) {
	o.MacdHist = v
}

func (o GetTimeSeriesMacdExt200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacdExt200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["macd"] = o.Macd
	toSerialize["macd_signal"] = o.MacdSignal
	toSerialize["macd_hist"] = o.MacdHist
	return toSerialize, nil
}

func (o *GetTimeSeriesMacdExt200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"macd",
		"macd_signal",
		"macd_hist",
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

	varGetTimeSeriesMacdExt200ResponseValuesInner := _GetTimeSeriesMacdExt200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMacdExt200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacdExt200ResponseValuesInner(varGetTimeSeriesMacdExt200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMacdExt200ResponseValuesInner struct {
	value *GetTimeSeriesMacdExt200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMacdExt200ResponseValuesInner) Get() *GetTimeSeriesMacdExt200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMacdExt200ResponseValuesInner) Set(val *GetTimeSeriesMacdExt200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacdExt200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacdExt200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacdExt200ResponseValuesInner(val *GetTimeSeriesMacdExt200ResponseValuesInner) *NullableGetTimeSeriesMacdExt200ResponseValuesInner {
	return &NullableGetTimeSeriesMacdExt200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacdExt200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacdExt200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
