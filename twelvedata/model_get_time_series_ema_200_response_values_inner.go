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

// checks if the GetTimeSeriesEma200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesEma200ResponseValuesInner{}

// GetTimeSeriesEma200ResponseValuesInner struct for GetTimeSeriesEma200ResponseValuesInner
type GetTimeSeriesEma200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// EMA value
	Ema string `json:"ema"`
}

type _GetTimeSeriesEma200ResponseValuesInner GetTimeSeriesEma200ResponseValuesInner

// NewGetTimeSeriesEma200ResponseValuesInner instantiates a new GetTimeSeriesEma200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesEma200ResponseValuesInner(datetime string, ema string) *GetTimeSeriesEma200ResponseValuesInner {
	this := GetTimeSeriesEma200ResponseValuesInner{}
	this.Datetime = datetime
	this.Ema = ema
	return &this
}

// NewGetTimeSeriesEma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesEma200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesEma200ResponseValuesInnerWithDefaults() *GetTimeSeriesEma200ResponseValuesInner {
	this := GetTimeSeriesEma200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesEma200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesEma200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesEma200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetEma returns the Ema field value
func (o *GetTimeSeriesEma200ResponseValuesInner) GetEma() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ema
}

// GetEmaOk returns a tuple with the Ema field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesEma200ResponseValuesInner) GetEmaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ema, true
}

// SetEma sets field value
func (o *GetTimeSeriesEma200ResponseValuesInner) SetEma(v string) {
	o.Ema = v
}

func (o GetTimeSeriesEma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesEma200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ema"] = o.Ema
	return toSerialize, nil
}

func (o *GetTimeSeriesEma200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ema",
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

	varGetTimeSeriesEma200ResponseValuesInner := _GetTimeSeriesEma200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesEma200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesEma200ResponseValuesInner(varGetTimeSeriesEma200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesEma200ResponseValuesInner struct {
	value *GetTimeSeriesEma200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesEma200ResponseValuesInner) Get() *GetTimeSeriesEma200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesEma200ResponseValuesInner) Set(val *GetTimeSeriesEma200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesEma200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesEma200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesEma200ResponseValuesInner(val *GetTimeSeriesEma200ResponseValuesInner) *NullableGetTimeSeriesEma200ResponseValuesInner {
	return &NullableGetTimeSeriesEma200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesEma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesEma200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
