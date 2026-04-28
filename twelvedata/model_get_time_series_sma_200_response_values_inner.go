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

// checks if the GetTimeSeriesSma200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSma200ResponseValuesInner{}

// GetTimeSeriesSma200ResponseValuesInner struct for GetTimeSeriesSma200ResponseValuesInner
type GetTimeSeriesSma200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SMA value
	Sma string `json:"sma"`
}

type _GetTimeSeriesSma200ResponseValuesInner GetTimeSeriesSma200ResponseValuesInner

// NewGetTimeSeriesSma200ResponseValuesInner instantiates a new GetTimeSeriesSma200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSma200ResponseValuesInner(datetime string, sma string) *GetTimeSeriesSma200ResponseValuesInner {
	this := GetTimeSeriesSma200ResponseValuesInner{}
	this.Datetime = datetime
	this.Sma = sma
	return &this
}

// NewGetTimeSeriesSma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSma200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSma200ResponseValuesInnerWithDefaults() *GetTimeSeriesSma200ResponseValuesInner {
	this := GetTimeSeriesSma200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSma200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSma200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSma200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSma returns the Sma field value
func (o *GetTimeSeriesSma200ResponseValuesInner) GetSma() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sma
}

// GetSmaOk returns a tuple with the Sma field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSma200ResponseValuesInner) GetSmaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sma, true
}

// SetSma sets field value
func (o *GetTimeSeriesSma200ResponseValuesInner) SetSma(v string) {
	o.Sma = v
}

func (o GetTimeSeriesSma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSma200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sma"] = o.Sma
	return toSerialize, nil
}

func (o *GetTimeSeriesSma200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sma",
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

	varGetTimeSeriesSma200ResponseValuesInner := _GetTimeSeriesSma200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesSma200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSma200ResponseValuesInner(varGetTimeSeriesSma200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSma200ResponseValuesInner struct {
	value *GetTimeSeriesSma200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSma200ResponseValuesInner) Get() *GetTimeSeriesSma200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSma200ResponseValuesInner) Set(val *GetTimeSeriesSma200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSma200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSma200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSma200ResponseValuesInner(val *GetTimeSeriesSma200ResponseValuesInner) *NullableGetTimeSeriesSma200ResponseValuesInner {
	return &NullableGetTimeSeriesSma200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSma200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
