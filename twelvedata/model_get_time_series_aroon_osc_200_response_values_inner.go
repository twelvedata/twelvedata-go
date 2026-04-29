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

// checks if the GetTimeSeriesAroonOsc200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAroonOsc200ResponseValuesInner{}

// GetTimeSeriesAroonOsc200ResponseValuesInner struct for GetTimeSeriesAroonOsc200ResponseValuesInner
type GetTimeSeriesAroonOsc200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Aroon oscillator value
	Aroonosc string `json:"aroonosc"`
}

type _GetTimeSeriesAroonOsc200ResponseValuesInner GetTimeSeriesAroonOsc200ResponseValuesInner

// NewGetTimeSeriesAroonOsc200ResponseValuesInner instantiates a new GetTimeSeriesAroonOsc200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAroonOsc200ResponseValuesInner(datetime string, aroonosc string) *GetTimeSeriesAroonOsc200ResponseValuesInner {
	this := GetTimeSeriesAroonOsc200ResponseValuesInner{}
	this.Datetime = datetime
	this.Aroonosc = aroonosc
	return &this
}

// NewGetTimeSeriesAroonOsc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAroonOsc200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAroonOsc200ResponseValuesInnerWithDefaults() *GetTimeSeriesAroonOsc200ResponseValuesInner {
	this := GetTimeSeriesAroonOsc200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAroonosc returns the Aroonosc field value
func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetAroonosc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aroonosc
}

// GetAroonoscOk returns a tuple with the Aroonosc field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) GetAroonoscOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aroonosc, true
}

// SetAroonosc sets field value
func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) SetAroonosc(v string) {
	o.Aroonosc = v
}

func (o GetTimeSeriesAroonOsc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAroonOsc200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["aroonosc"] = o.Aroonosc
	return toSerialize, nil
}

func (o *GetTimeSeriesAroonOsc200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"aroonosc",
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

	varGetTimeSeriesAroonOsc200ResponseValuesInner := _GetTimeSeriesAroonOsc200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAroonOsc200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAroonOsc200ResponseValuesInner(varGetTimeSeriesAroonOsc200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAroonOsc200ResponseValuesInner struct {
	value *GetTimeSeriesAroonOsc200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAroonOsc200ResponseValuesInner) Get() *GetTimeSeriesAroonOsc200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAroonOsc200ResponseValuesInner) Set(val *GetTimeSeriesAroonOsc200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAroonOsc200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAroonOsc200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAroonOsc200ResponseValuesInner(val *GetTimeSeriesAroonOsc200ResponseValuesInner) *NullableGetTimeSeriesAroonOsc200ResponseValuesInner {
	return &NullableGetTimeSeriesAroonOsc200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAroonOsc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAroonOsc200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
