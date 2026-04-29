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

// checks if the GetTimeSeriesLog10200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLog10200ResponseValuesInner{}

// GetTimeSeriesLog10200ResponseValuesInner struct for GetTimeSeriesLog10200ResponseValuesInner
type GetTimeSeriesLog10200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Log10 value
	Log10 string `json:"log10"`
}

type _GetTimeSeriesLog10200ResponseValuesInner GetTimeSeriesLog10200ResponseValuesInner

// NewGetTimeSeriesLog10200ResponseValuesInner instantiates a new GetTimeSeriesLog10200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLog10200ResponseValuesInner(datetime string, log10 string) *GetTimeSeriesLog10200ResponseValuesInner {
	this := GetTimeSeriesLog10200ResponseValuesInner{}
	this.Datetime = datetime
	this.Log10 = log10
	return &this
}

// NewGetTimeSeriesLog10200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLog10200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLog10200ResponseValuesInnerWithDefaults() *GetTimeSeriesLog10200ResponseValuesInner {
	this := GetTimeSeriesLog10200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesLog10200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLog10200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesLog10200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetLog10 returns the Log10 field value
func (o *GetTimeSeriesLog10200ResponseValuesInner) GetLog10() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Log10
}

// GetLog10Ok returns a tuple with the Log10 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLog10200ResponseValuesInner) GetLog10Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log10, true
}

// SetLog10 sets field value
func (o *GetTimeSeriesLog10200ResponseValuesInner) SetLog10(v string) {
	o.Log10 = v
}

func (o GetTimeSeriesLog10200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLog10200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["log10"] = o.Log10
	return toSerialize, nil
}

func (o *GetTimeSeriesLog10200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"log10",
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

	varGetTimeSeriesLog10200ResponseValuesInner := _GetTimeSeriesLog10200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesLog10200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLog10200ResponseValuesInner(varGetTimeSeriesLog10200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesLog10200ResponseValuesInner struct {
	value *GetTimeSeriesLog10200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesLog10200ResponseValuesInner) Get() *GetTimeSeriesLog10200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesLog10200ResponseValuesInner) Set(val *GetTimeSeriesLog10200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLog10200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLog10200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLog10200ResponseValuesInner(val *GetTimeSeriesLog10200ResponseValuesInner) *NullableGetTimeSeriesLog10200ResponseValuesInner {
	return &NullableGetTimeSeriesLog10200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLog10200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLog10200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
