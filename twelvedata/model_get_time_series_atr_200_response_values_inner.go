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

// checks if the GetTimeSeriesAtr200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAtr200ResponseValuesInner{}

// GetTimeSeriesAtr200ResponseValuesInner struct for GetTimeSeriesAtr200ResponseValuesInner
type GetTimeSeriesAtr200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ATR value
	Atr string `json:"atr"`
}

type _GetTimeSeriesAtr200ResponseValuesInner GetTimeSeriesAtr200ResponseValuesInner

// NewGetTimeSeriesAtr200ResponseValuesInner instantiates a new GetTimeSeriesAtr200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAtr200ResponseValuesInner(datetime string, atr string) *GetTimeSeriesAtr200ResponseValuesInner {
	this := GetTimeSeriesAtr200ResponseValuesInner{}
	this.Datetime = datetime
	this.Atr = atr
	return &this
}

// NewGetTimeSeriesAtr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAtr200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAtr200ResponseValuesInnerWithDefaults() *GetTimeSeriesAtr200ResponseValuesInner {
	this := GetTimeSeriesAtr200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAtr200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAtr200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAtr200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAtr returns the Atr field value
func (o *GetTimeSeriesAtr200ResponseValuesInner) GetAtr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Atr
}

// GetAtrOk returns a tuple with the Atr field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAtr200ResponseValuesInner) GetAtrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Atr, true
}

// SetAtr sets field value
func (o *GetTimeSeriesAtr200ResponseValuesInner) SetAtr(v string) {
	o.Atr = v
}

func (o GetTimeSeriesAtr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAtr200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["atr"] = o.Atr
	return toSerialize, nil
}

func (o *GetTimeSeriesAtr200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"atr",
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

	varGetTimeSeriesAtr200ResponseValuesInner := _GetTimeSeriesAtr200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAtr200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAtr200ResponseValuesInner(varGetTimeSeriesAtr200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAtr200ResponseValuesInner struct {
	value *GetTimeSeriesAtr200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAtr200ResponseValuesInner) Get() *GetTimeSeriesAtr200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAtr200ResponseValuesInner) Set(val *GetTimeSeriesAtr200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAtr200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAtr200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAtr200ResponseValuesInner(val *GetTimeSeriesAtr200ResponseValuesInner) *NullableGetTimeSeriesAtr200ResponseValuesInner {
	return &NullableGetTimeSeriesAtr200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAtr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAtr200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
