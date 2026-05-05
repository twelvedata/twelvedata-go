// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesTrima200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTrima200ResponseValuesInner{}

// GetTimeSeriesTrima200ResponseValuesInner struct for GetTimeSeriesTrima200ResponseValuesInner
type GetTimeSeriesTrima200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// TRIMA value
	Trima string `json:"trima"`
}

type _GetTimeSeriesTrima200ResponseValuesInner GetTimeSeriesTrima200ResponseValuesInner

// NewGetTimeSeriesTrima200ResponseValuesInner instantiates a new GetTimeSeriesTrima200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTrima200ResponseValuesInner(datetime string, trima string) *GetTimeSeriesTrima200ResponseValuesInner {
	this := GetTimeSeriesTrima200ResponseValuesInner{}
	this.Datetime = datetime
	this.Trima = trima
	return &this
}

// NewGetTimeSeriesTrima200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTrima200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTrima200ResponseValuesInnerWithDefaults() *GetTimeSeriesTrima200ResponseValuesInner {
	this := GetTimeSeriesTrima200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesTrima200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTrima200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesTrima200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTrima returns the Trima field value
func (o *GetTimeSeriesTrima200ResponseValuesInner) GetTrima() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trima
}

// GetTrimaOk returns a tuple with the Trima field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTrima200ResponseValuesInner) GetTrimaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trima, true
}

// SetTrima sets field value
func (o *GetTimeSeriesTrima200ResponseValuesInner) SetTrima(v string) {
	o.Trima = v
}

func (o GetTimeSeriesTrima200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTrima200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["trima"] = o.Trima
	return toSerialize, nil
}

func (o *GetTimeSeriesTrima200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"trima",
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

	varGetTimeSeriesTrima200ResponseValuesInner := _GetTimeSeriesTrima200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTrima200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTrima200ResponseValuesInner(varGetTimeSeriesTrima200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesTrima200ResponseValuesInner struct {
	value *GetTimeSeriesTrima200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesTrima200ResponseValuesInner) Get() *GetTimeSeriesTrima200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesTrima200ResponseValuesInner) Set(val *GetTimeSeriesTrima200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTrima200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTrima200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTrima200ResponseValuesInner(val *GetTimeSeriesTrima200ResponseValuesInner) *NullableGetTimeSeriesTrima200ResponseValuesInner {
	return &NullableGetTimeSeriesTrima200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTrima200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTrima200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
