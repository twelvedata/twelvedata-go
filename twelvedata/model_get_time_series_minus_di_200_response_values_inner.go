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

// checks if the GetTimeSeriesMinusDI200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinusDI200ResponseValuesInner{}

// GetTimeSeriesMinusDI200ResponseValuesInner struct for GetTimeSeriesMinusDI200ResponseValuesInner
type GetTimeSeriesMinusDI200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Minus_di value
	MinusDi string `json:"minus_di"`
}

type _GetTimeSeriesMinusDI200ResponseValuesInner GetTimeSeriesMinusDI200ResponseValuesInner

// NewGetTimeSeriesMinusDI200ResponseValuesInner instantiates a new GetTimeSeriesMinusDI200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinusDI200ResponseValuesInner(datetime string, minusDi string) *GetTimeSeriesMinusDI200ResponseValuesInner {
	this := GetTimeSeriesMinusDI200ResponseValuesInner{}
	this.Datetime = datetime
	this.MinusDi = minusDi
	return &this
}

// NewGetTimeSeriesMinusDI200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinusDI200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinusDI200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinusDI200ResponseValuesInner {
	this := GetTimeSeriesMinusDI200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMinusDI200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMinusDi returns the MinusDi field value
func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetMinusDi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinusDi
}

// GetMinusDiOk returns a tuple with the MinusDi field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDI200ResponseValuesInner) GetMinusDiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinusDi, true
}

// SetMinusDi sets field value
func (o *GetTimeSeriesMinusDI200ResponseValuesInner) SetMinusDi(v string) {
	o.MinusDi = v
}

func (o GetTimeSeriesMinusDI200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinusDI200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["minus_di"] = o.MinusDi
	return toSerialize, nil
}

func (o *GetTimeSeriesMinusDI200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"minus_di",
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

	varGetTimeSeriesMinusDI200ResponseValuesInner := _GetTimeSeriesMinusDI200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMinusDI200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinusDI200ResponseValuesInner(varGetTimeSeriesMinusDI200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMinusDI200ResponseValuesInner struct {
	value *GetTimeSeriesMinusDI200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMinusDI200ResponseValuesInner) Get() *GetTimeSeriesMinusDI200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMinusDI200ResponseValuesInner) Set(val *GetTimeSeriesMinusDI200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinusDI200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinusDI200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinusDI200ResponseValuesInner(val *GetTimeSeriesMinusDI200ResponseValuesInner) *NullableGetTimeSeriesMinusDI200ResponseValuesInner {
	return &NullableGetTimeSeriesMinusDI200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinusDI200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinusDI200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
