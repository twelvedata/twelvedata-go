// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesWma200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesWma200ResponseValuesInner{}

// GetTimeSeriesWma200ResponseValuesInner struct for GetTimeSeriesWma200ResponseValuesInner
type GetTimeSeriesWma200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// WMA value
	Wma string `json:"wma"`
}

type _GetTimeSeriesWma200ResponseValuesInner GetTimeSeriesWma200ResponseValuesInner

// NewGetTimeSeriesWma200ResponseValuesInner instantiates a new GetTimeSeriesWma200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesWma200ResponseValuesInner(datetime string, wma string) *GetTimeSeriesWma200ResponseValuesInner {
	this := GetTimeSeriesWma200ResponseValuesInner{}
	this.Datetime = datetime
	this.Wma = wma
	return &this
}

// NewGetTimeSeriesWma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesWma200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesWma200ResponseValuesInnerWithDefaults() *GetTimeSeriesWma200ResponseValuesInner {
	this := GetTimeSeriesWma200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesWma200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWma200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesWma200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetWma returns the Wma field value
func (o *GetTimeSeriesWma200ResponseValuesInner) GetWma() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wma
}

// GetWmaOk returns a tuple with the Wma field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWma200ResponseValuesInner) GetWmaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wma, true
}

// SetWma sets field value
func (o *GetTimeSeriesWma200ResponseValuesInner) SetWma(v string) {
	o.Wma = v
}

func (o GetTimeSeriesWma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesWma200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["wma"] = o.Wma
	return toSerialize, nil
}

func (o *GetTimeSeriesWma200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"wma",
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

	varGetTimeSeriesWma200ResponseValuesInner := _GetTimeSeriesWma200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesWma200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesWma200ResponseValuesInner(varGetTimeSeriesWma200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesWma200ResponseValuesInner struct {
	value *GetTimeSeriesWma200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesWma200ResponseValuesInner) Get() *GetTimeSeriesWma200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesWma200ResponseValuesInner) Set(val *GetTimeSeriesWma200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesWma200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesWma200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesWma200ResponseValuesInner(val *GetTimeSeriesWma200ResponseValuesInner) *NullableGetTimeSeriesWma200ResponseValuesInner {
	return &NullableGetTimeSeriesWma200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesWma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesWma200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
