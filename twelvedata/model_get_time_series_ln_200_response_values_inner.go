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

// checks if the GetTimeSeriesLn200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLn200ResponseValuesInner{}

// GetTimeSeriesLn200ResponseValuesInner struct for GetTimeSeriesLn200ResponseValuesInner
type GetTimeSeriesLn200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Natural logarithm value
	Ln string `json:"ln"`
}

type _GetTimeSeriesLn200ResponseValuesInner GetTimeSeriesLn200ResponseValuesInner

// NewGetTimeSeriesLn200ResponseValuesInner instantiates a new GetTimeSeriesLn200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLn200ResponseValuesInner(datetime string, ln string) *GetTimeSeriesLn200ResponseValuesInner {
	this := GetTimeSeriesLn200ResponseValuesInner{}
	this.Datetime = datetime
	this.Ln = ln
	return &this
}

// NewGetTimeSeriesLn200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLn200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLn200ResponseValuesInnerWithDefaults() *GetTimeSeriesLn200ResponseValuesInner {
	this := GetTimeSeriesLn200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesLn200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLn200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesLn200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetLn returns the Ln field value
func (o *GetTimeSeriesLn200ResponseValuesInner) GetLn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ln
}

// GetLnOk returns a tuple with the Ln field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLn200ResponseValuesInner) GetLnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ln, true
}

// SetLn sets field value
func (o *GetTimeSeriesLn200ResponseValuesInner) SetLn(v string) {
	o.Ln = v
}

func (o GetTimeSeriesLn200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLn200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ln"] = o.Ln
	return toSerialize, nil
}

func (o *GetTimeSeriesLn200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ln",
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

	varGetTimeSeriesLn200ResponseValuesInner := _GetTimeSeriesLn200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesLn200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLn200ResponseValuesInner(varGetTimeSeriesLn200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesLn200ResponseValuesInner struct {
	value *GetTimeSeriesLn200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesLn200ResponseValuesInner) Get() *GetTimeSeriesLn200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesLn200ResponseValuesInner) Set(val *GetTimeSeriesLn200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLn200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLn200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLn200ResponseValuesInner(val *GetTimeSeriesLn200ResponseValuesInner) *NullableGetTimeSeriesLn200ResponseValuesInner {
	return &NullableGetTimeSeriesLn200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLn200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLn200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
