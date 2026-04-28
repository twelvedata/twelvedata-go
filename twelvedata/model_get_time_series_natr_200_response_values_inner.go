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

// checks if the GetTimeSeriesNatr200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesNatr200ResponseValuesInner{}

// GetTimeSeriesNatr200ResponseValuesInner struct for GetTimeSeriesNatr200ResponseValuesInner
type GetTimeSeriesNatr200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// natr value
	Natr string `json:"natr"`
}

type _GetTimeSeriesNatr200ResponseValuesInner GetTimeSeriesNatr200ResponseValuesInner

// NewGetTimeSeriesNatr200ResponseValuesInner instantiates a new GetTimeSeriesNatr200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesNatr200ResponseValuesInner(datetime string, natr string) *GetTimeSeriesNatr200ResponseValuesInner {
	this := GetTimeSeriesNatr200ResponseValuesInner{}
	this.Datetime = datetime
	this.Natr = natr
	return &this
}

// NewGetTimeSeriesNatr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesNatr200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesNatr200ResponseValuesInnerWithDefaults() *GetTimeSeriesNatr200ResponseValuesInner {
	this := GetTimeSeriesNatr200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesNatr200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesNatr200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesNatr200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetNatr returns the Natr field value
func (o *GetTimeSeriesNatr200ResponseValuesInner) GetNatr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Natr
}

// GetNatrOk returns a tuple with the Natr field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesNatr200ResponseValuesInner) GetNatrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Natr, true
}

// SetNatr sets field value
func (o *GetTimeSeriesNatr200ResponseValuesInner) SetNatr(v string) {
	o.Natr = v
}

func (o GetTimeSeriesNatr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesNatr200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["natr"] = o.Natr
	return toSerialize, nil
}

func (o *GetTimeSeriesNatr200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"natr",
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

	varGetTimeSeriesNatr200ResponseValuesInner := _GetTimeSeriesNatr200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesNatr200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesNatr200ResponseValuesInner(varGetTimeSeriesNatr200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesNatr200ResponseValuesInner struct {
	value *GetTimeSeriesNatr200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesNatr200ResponseValuesInner) Get() *GetTimeSeriesNatr200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesNatr200ResponseValuesInner) Set(val *GetTimeSeriesNatr200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesNatr200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesNatr200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesNatr200ResponseValuesInner(val *GetTimeSeriesNatr200ResponseValuesInner) *NullableGetTimeSeriesNatr200ResponseValuesInner {
	return &NullableGetTimeSeriesNatr200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesNatr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesNatr200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
