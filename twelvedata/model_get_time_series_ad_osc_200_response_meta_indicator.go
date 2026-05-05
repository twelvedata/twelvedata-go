// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAdOsc200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdOsc200ResponseMetaIndicator{}

// GetTimeSeriesAdOsc200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesAdOsc200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Number of periods for fast moving average
	FastPeriod int64 `json:"fast_period"`
	// Number of periods for slow moving average
	SlowPeriod int64 `json:"slow_period"`
}

type _GetTimeSeriesAdOsc200ResponseMetaIndicator GetTimeSeriesAdOsc200ResponseMetaIndicator

// NewGetTimeSeriesAdOsc200ResponseMetaIndicator instantiates a new GetTimeSeriesAdOsc200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdOsc200ResponseMetaIndicator(name string, fastPeriod int64, slowPeriod int64) *GetTimeSeriesAdOsc200ResponseMetaIndicator {
	this := GetTimeSeriesAdOsc200ResponseMetaIndicator{}
	this.Name = name
	this.FastPeriod = fastPeriod
	this.SlowPeriod = slowPeriod
	return &this
}

// NewGetTimeSeriesAdOsc200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesAdOsc200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdOsc200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesAdOsc200ResponseMetaIndicator {
	this := GetTimeSeriesAdOsc200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetFastPeriod returns the FastPeriod field value
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) GetFastPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastPeriod
}

// GetFastPeriodOk returns a tuple with the FastPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) GetFastPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastPeriod, true
}

// SetFastPeriod sets field value
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) SetFastPeriod(v int64) {
	o.FastPeriod = v
}

// GetSlowPeriod returns the SlowPeriod field value
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) GetSlowPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowPeriod
}

// GetSlowPeriodOk returns a tuple with the SlowPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) GetSlowPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowPeriod, true
}

// SetSlowPeriod sets field value
func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) SetSlowPeriod(v int64) {
	o.SlowPeriod = v
}

func (o GetTimeSeriesAdOsc200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdOsc200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fast_period"] = o.FastPeriod
	toSerialize["slow_period"] = o.SlowPeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesAdOsc200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fast_period",
		"slow_period",
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

	varGetTimeSeriesAdOsc200ResponseMetaIndicator := _GetTimeSeriesAdOsc200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAdOsc200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdOsc200ResponseMetaIndicator(varGetTimeSeriesAdOsc200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesAdOsc200ResponseMetaIndicator struct {
	value *GetTimeSeriesAdOsc200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesAdOsc200ResponseMetaIndicator) Get() *GetTimeSeriesAdOsc200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesAdOsc200ResponseMetaIndicator) Set(val *GetTimeSeriesAdOsc200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdOsc200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdOsc200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdOsc200ResponseMetaIndicator(val *GetTimeSeriesAdOsc200ResponseMetaIndicator) *NullableGetTimeSeriesAdOsc200ResponseMetaIndicator {
	return &NullableGetTimeSeriesAdOsc200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdOsc200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdOsc200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
