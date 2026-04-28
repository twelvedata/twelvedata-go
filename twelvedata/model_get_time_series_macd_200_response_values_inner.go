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

// checks if the GetTimeSeriesMacd200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacd200ResponseValuesInner{}

// GetTimeSeriesMacd200ResponseValuesInner struct for GetTimeSeriesMacd200ResponseValuesInner
type GetTimeSeriesMacd200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// MACD value
	Macd string `json:"macd"`
	// MACD signal line value
	MacdSignal string `json:"macd_signal"`
	// MACD histogram value
	MacdHist string `json:"macd_hist"`
}

type _GetTimeSeriesMacd200ResponseValuesInner GetTimeSeriesMacd200ResponseValuesInner

// NewGetTimeSeriesMacd200ResponseValuesInner instantiates a new GetTimeSeriesMacd200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacd200ResponseValuesInner(datetime string, macd string, macdSignal string, macdHist string) *GetTimeSeriesMacd200ResponseValuesInner {
	this := GetTimeSeriesMacd200ResponseValuesInner{}
	this.Datetime = datetime
	this.Macd = macd
	this.MacdSignal = macdSignal
	this.MacdHist = macdHist
	return &this
}

// NewGetTimeSeriesMacd200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMacd200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacd200ResponseValuesInnerWithDefaults() *GetTimeSeriesMacd200ResponseValuesInner {
	this := GetTimeSeriesMacd200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMacd returns the Macd field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Macd
}

// GetMacdOk returns a tuple with the Macd field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Macd, true
}

// SetMacd sets field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) SetMacd(v string) {
	o.Macd = v
}

// GetMacdSignal returns the MacdSignal field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdSignal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacdSignal
}

// GetMacdSignalOk returns a tuple with the MacdSignal field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdSignalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacdSignal, true
}

// SetMacdSignal sets field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) SetMacdSignal(v string) {
	o.MacdSignal = v
}

// GetMacdHist returns the MacdHist field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdHist() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacdHist
}

// GetMacdHistOk returns a tuple with the MacdHist field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseValuesInner) GetMacdHistOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacdHist, true
}

// SetMacdHist sets field value
func (o *GetTimeSeriesMacd200ResponseValuesInner) SetMacdHist(v string) {
	o.MacdHist = v
}

func (o GetTimeSeriesMacd200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacd200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["macd"] = o.Macd
	toSerialize["macd_signal"] = o.MacdSignal
	toSerialize["macd_hist"] = o.MacdHist
	return toSerialize, nil
}

func (o *GetTimeSeriesMacd200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMacd200ResponseValuesInner := _GetTimeSeriesMacd200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMacd200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacd200ResponseValuesInner(varGetTimeSeriesMacd200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMacd200ResponseValuesInner struct {
	value *GetTimeSeriesMacd200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMacd200ResponseValuesInner) Get() *GetTimeSeriesMacd200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMacd200ResponseValuesInner) Set(val *GetTimeSeriesMacd200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacd200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacd200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacd200ResponseValuesInner(val *GetTimeSeriesMacd200ResponseValuesInner) *NullableGetTimeSeriesMacd200ResponseValuesInner {
	return &NullableGetTimeSeriesMacd200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacd200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacd200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
