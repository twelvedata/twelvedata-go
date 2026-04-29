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

// checks if the GetTimeSeriesUltOsc200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesUltOsc200ResponseValuesInner{}

// GetTimeSeriesUltOsc200ResponseValuesInner struct for GetTimeSeriesUltOsc200ResponseValuesInner
type GetTimeSeriesUltOsc200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Ultimate Oscillator value
	Ultosc string `json:"ultosc"`
}

type _GetTimeSeriesUltOsc200ResponseValuesInner GetTimeSeriesUltOsc200ResponseValuesInner

// NewGetTimeSeriesUltOsc200ResponseValuesInner instantiates a new GetTimeSeriesUltOsc200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesUltOsc200ResponseValuesInner(datetime string, ultosc string) *GetTimeSeriesUltOsc200ResponseValuesInner {
	this := GetTimeSeriesUltOsc200ResponseValuesInner{}
	this.Datetime = datetime
	this.Ultosc = ultosc
	return &this
}

// NewGetTimeSeriesUltOsc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesUltOsc200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesUltOsc200ResponseValuesInnerWithDefaults() *GetTimeSeriesUltOsc200ResponseValuesInner {
	this := GetTimeSeriesUltOsc200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesUltOsc200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetUltosc returns the Ultosc field value
func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetUltosc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ultosc
}

// GetUltoscOk returns a tuple with the Ultosc field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetUltoscOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ultosc, true
}

// SetUltosc sets field value
func (o *GetTimeSeriesUltOsc200ResponseValuesInner) SetUltosc(v string) {
	o.Ultosc = v
}

func (o GetTimeSeriesUltOsc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesUltOsc200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ultosc"] = o.Ultosc
	return toSerialize, nil
}

func (o *GetTimeSeriesUltOsc200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ultosc",
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

	varGetTimeSeriesUltOsc200ResponseValuesInner := _GetTimeSeriesUltOsc200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesUltOsc200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesUltOsc200ResponseValuesInner(varGetTimeSeriesUltOsc200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesUltOsc200ResponseValuesInner struct {
	value *GetTimeSeriesUltOsc200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesUltOsc200ResponseValuesInner) Get() *GetTimeSeriesUltOsc200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesUltOsc200ResponseValuesInner) Set(val *GetTimeSeriesUltOsc200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesUltOsc200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesUltOsc200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesUltOsc200ResponseValuesInner(val *GetTimeSeriesUltOsc200ResponseValuesInner) *NullableGetTimeSeriesUltOsc200ResponseValuesInner {
	return &NullableGetTimeSeriesUltOsc200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesUltOsc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesUltOsc200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
