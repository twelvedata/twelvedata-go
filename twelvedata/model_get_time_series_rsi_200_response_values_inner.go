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

// checks if the GetTimeSeriesRsi200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRsi200ResponseValuesInner{}

// GetTimeSeriesRsi200ResponseValuesInner struct for GetTimeSeriesRsi200ResponseValuesInner
type GetTimeSeriesRsi200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// RSI value
	Rsi string `json:"rsi"`
}

type _GetTimeSeriesRsi200ResponseValuesInner GetTimeSeriesRsi200ResponseValuesInner

// NewGetTimeSeriesRsi200ResponseValuesInner instantiates a new GetTimeSeriesRsi200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRsi200ResponseValuesInner(datetime string, rsi string) *GetTimeSeriesRsi200ResponseValuesInner {
	this := GetTimeSeriesRsi200ResponseValuesInner{}
	this.Datetime = datetime
	this.Rsi = rsi
	return &this
}

// NewGetTimeSeriesRsi200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRsi200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRsi200ResponseValuesInnerWithDefaults() *GetTimeSeriesRsi200ResponseValuesInner {
	this := GetTimeSeriesRsi200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesRsi200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRsi200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesRsi200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetRsi returns the Rsi field value
func (o *GetTimeSeriesRsi200ResponseValuesInner) GetRsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rsi
}

// GetRsiOk returns a tuple with the Rsi field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRsi200ResponseValuesInner) GetRsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rsi, true
}

// SetRsi sets field value
func (o *GetTimeSeriesRsi200ResponseValuesInner) SetRsi(v string) {
	o.Rsi = v
}

func (o GetTimeSeriesRsi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRsi200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["rsi"] = o.Rsi
	return toSerialize, nil
}

func (o *GetTimeSeriesRsi200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"rsi",
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

	varGetTimeSeriesRsi200ResponseValuesInner := _GetTimeSeriesRsi200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesRsi200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRsi200ResponseValuesInner(varGetTimeSeriesRsi200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesRsi200ResponseValuesInner struct {
	value *GetTimeSeriesRsi200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesRsi200ResponseValuesInner) Get() *GetTimeSeriesRsi200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesRsi200ResponseValuesInner) Set(val *GetTimeSeriesRsi200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRsi200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRsi200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRsi200ResponseValuesInner(val *GetTimeSeriesRsi200ResponseValuesInner) *NullableGetTimeSeriesRsi200ResponseValuesInner {
	return &NullableGetTimeSeriesRsi200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRsi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRsi200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
