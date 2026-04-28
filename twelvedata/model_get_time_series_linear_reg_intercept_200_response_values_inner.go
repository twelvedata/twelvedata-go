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

// checks if the GetTimeSeriesLinearRegIntercept200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLinearRegIntercept200ResponseValuesInner{}

// GetTimeSeriesLinearRegIntercept200ResponseValuesInner struct for GetTimeSeriesLinearRegIntercept200ResponseValuesInner
type GetTimeSeriesLinearRegIntercept200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Linear Regression Intercept value
	Linearregintercept string `json:"linearregintercept"`
}

type _GetTimeSeriesLinearRegIntercept200ResponseValuesInner GetTimeSeriesLinearRegIntercept200ResponseValuesInner

// NewGetTimeSeriesLinearRegIntercept200ResponseValuesInner instantiates a new GetTimeSeriesLinearRegIntercept200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLinearRegIntercept200ResponseValuesInner(datetime string, linearregintercept string) *GetTimeSeriesLinearRegIntercept200ResponseValuesInner {
	this := GetTimeSeriesLinearRegIntercept200ResponseValuesInner{}
	this.Datetime = datetime
	this.Linearregintercept = linearregintercept
	return &this
}

// NewGetTimeSeriesLinearRegIntercept200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLinearRegIntercept200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLinearRegIntercept200ResponseValuesInnerWithDefaults() *GetTimeSeriesLinearRegIntercept200ResponseValuesInner {
	this := GetTimeSeriesLinearRegIntercept200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetLinearregintercept returns the Linearregintercept field value
func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetLinearregintercept() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Linearregintercept
}

// GetLinearreginterceptOk returns a tuple with the Linearregintercept field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetLinearreginterceptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linearregintercept, true
}

// SetLinearregintercept sets field value
func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) SetLinearregintercept(v string) {
	o.Linearregintercept = v
}

func (o GetTimeSeriesLinearRegIntercept200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLinearRegIntercept200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["linearregintercept"] = o.Linearregintercept
	return toSerialize, nil
}

func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"linearregintercept",
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

	varGetTimeSeriesLinearRegIntercept200ResponseValuesInner := _GetTimeSeriesLinearRegIntercept200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesLinearRegIntercept200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLinearRegIntercept200ResponseValuesInner(varGetTimeSeriesLinearRegIntercept200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner struct {
	value *GetTimeSeriesLinearRegIntercept200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner) Get() *GetTimeSeriesLinearRegIntercept200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner) Set(val *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner(val *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) *NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner {
	return &NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLinearRegIntercept200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
