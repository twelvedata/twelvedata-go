// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesUltOsc200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesUltOsc200ResponseMetaIndicator{}

// GetTimeSeriesUltOsc200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesUltOsc200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The first time period used for calculation in the indicator
	TimePeriod1 int64 `json:"time_period_1"`
	// The second time period used for calculation in the indicator
	TimePeriod2 int64 `json:"time_period_2"`
	// The third time period used for calculation in the indicator
	TimePeriod3 int64 `json:"time_period_3"`
}

type _GetTimeSeriesUltOsc200ResponseMetaIndicator GetTimeSeriesUltOsc200ResponseMetaIndicator

// NewGetTimeSeriesUltOsc200ResponseMetaIndicator instantiates a new GetTimeSeriesUltOsc200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesUltOsc200ResponseMetaIndicator(name string, timePeriod1 int64, timePeriod2 int64, timePeriod3 int64) *GetTimeSeriesUltOsc200ResponseMetaIndicator {
	this := GetTimeSeriesUltOsc200ResponseMetaIndicator{}
	this.Name = name
	this.TimePeriod1 = timePeriod1
	this.TimePeriod2 = timePeriod2
	this.TimePeriod3 = timePeriod3
	return &this
}

// NewGetTimeSeriesUltOsc200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesUltOsc200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesUltOsc200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesUltOsc200ResponseMetaIndicator {
	this := GetTimeSeriesUltOsc200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetTimePeriod1 returns the TimePeriod1 field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetTimePeriod1() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod1
}

// GetTimePeriod1Ok returns a tuple with the TimePeriod1 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetTimePeriod1Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod1, true
}

// SetTimePeriod1 sets field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) SetTimePeriod1(v int64) {
	o.TimePeriod1 = v
}

// GetTimePeriod2 returns the TimePeriod2 field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetTimePeriod2() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod2
}

// GetTimePeriod2Ok returns a tuple with the TimePeriod2 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetTimePeriod2Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod2, true
}

// SetTimePeriod2 sets field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) SetTimePeriod2(v int64) {
	o.TimePeriod2 = v
}

// GetTimePeriod3 returns the TimePeriod3 field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetTimePeriod3() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod3
}

// GetTimePeriod3Ok returns a tuple with the TimePeriod3 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) GetTimePeriod3Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod3, true
}

// SetTimePeriod3 sets field value
func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) SetTimePeriod3(v int64) {
	o.TimePeriod3 = v
}

func (o GetTimeSeriesUltOsc200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesUltOsc200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["time_period_1"] = o.TimePeriod1
	toSerialize["time_period_2"] = o.TimePeriod2
	toSerialize["time_period_3"] = o.TimePeriod3
	return toSerialize, nil
}

func (o *GetTimeSeriesUltOsc200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"time_period_1",
		"time_period_2",
		"time_period_3",
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

	varGetTimeSeriesUltOsc200ResponseMetaIndicator := _GetTimeSeriesUltOsc200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesUltOsc200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesUltOsc200ResponseMetaIndicator(varGetTimeSeriesUltOsc200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesUltOsc200ResponseMetaIndicator struct {
	value *GetTimeSeriesUltOsc200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesUltOsc200ResponseMetaIndicator) Get() *GetTimeSeriesUltOsc200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesUltOsc200ResponseMetaIndicator) Set(val *GetTimeSeriesUltOsc200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesUltOsc200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesUltOsc200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesUltOsc200ResponseMetaIndicator(val *GetTimeSeriesUltOsc200ResponseMetaIndicator) *NullableGetTimeSeriesUltOsc200ResponseMetaIndicator {
	return &NullableGetTimeSeriesUltOsc200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesUltOsc200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesUltOsc200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
