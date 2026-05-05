// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesTsf200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTsf200ResponseValuesInner{}

// GetTimeSeriesTsf200ResponseValuesInner struct for GetTimeSeriesTsf200ResponseValuesInner
type GetTimeSeriesTsf200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// TSF value
	Tsf string `json:"tsf"`
}

type _GetTimeSeriesTsf200ResponseValuesInner GetTimeSeriesTsf200ResponseValuesInner

// NewGetTimeSeriesTsf200ResponseValuesInner instantiates a new GetTimeSeriesTsf200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTsf200ResponseValuesInner(datetime string, tsf string) *GetTimeSeriesTsf200ResponseValuesInner {
	this := GetTimeSeriesTsf200ResponseValuesInner{}
	this.Datetime = datetime
	this.Tsf = tsf
	return &this
}

// NewGetTimeSeriesTsf200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTsf200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTsf200ResponseValuesInnerWithDefaults() *GetTimeSeriesTsf200ResponseValuesInner {
	this := GetTimeSeriesTsf200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesTsf200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTsf200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesTsf200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTsf returns the Tsf field value
func (o *GetTimeSeriesTsf200ResponseValuesInner) GetTsf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tsf
}

// GetTsfOk returns a tuple with the Tsf field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTsf200ResponseValuesInner) GetTsfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tsf, true
}

// SetTsf sets field value
func (o *GetTimeSeriesTsf200ResponseValuesInner) SetTsf(v string) {
	o.Tsf = v
}

func (o GetTimeSeriesTsf200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTsf200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["tsf"] = o.Tsf
	return toSerialize, nil
}

func (o *GetTimeSeriesTsf200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"tsf",
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

	varGetTimeSeriesTsf200ResponseValuesInner := _GetTimeSeriesTsf200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTsf200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTsf200ResponseValuesInner(varGetTimeSeriesTsf200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesTsf200ResponseValuesInner struct {
	value *GetTimeSeriesTsf200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesTsf200ResponseValuesInner) Get() *GetTimeSeriesTsf200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesTsf200ResponseValuesInner) Set(val *GetTimeSeriesTsf200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTsf200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTsf200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTsf200ResponseValuesInner(val *GetTimeSeriesTsf200ResponseValuesInner) *NullableGetTimeSeriesTsf200ResponseValuesInner {
	return &NullableGetTimeSeriesTsf200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTsf200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTsf200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
