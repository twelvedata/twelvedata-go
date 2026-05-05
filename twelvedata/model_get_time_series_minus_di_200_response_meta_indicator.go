// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMinusDI200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinusDI200ResponseMetaIndicator{}

// GetTimeSeriesMinusDI200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMinusDI200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesMinusDI200ResponseMetaIndicator GetTimeSeriesMinusDI200ResponseMetaIndicator

// NewGetTimeSeriesMinusDI200ResponseMetaIndicator instantiates a new GetTimeSeriesMinusDI200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinusDI200ResponseMetaIndicator(name string, timePeriod int64) *GetTimeSeriesMinusDI200ResponseMetaIndicator {
	this := GetTimeSeriesMinusDI200ResponseMetaIndicator{}
	this.Name = name
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesMinusDI200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMinusDI200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinusDI200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMinusDI200ResponseMetaIndicator {
	this := GetTimeSeriesMinusDI200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMinusDI200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDI200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMinusDI200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesMinusDI200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDI200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesMinusDI200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesMinusDI200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinusDI200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesMinusDI200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"time_period",
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

	varGetTimeSeriesMinusDI200ResponseMetaIndicator := _GetTimeSeriesMinusDI200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinusDI200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinusDI200ResponseMetaIndicator(varGetTimeSeriesMinusDI200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMinusDI200ResponseMetaIndicator struct {
	value *GetTimeSeriesMinusDI200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMinusDI200ResponseMetaIndicator) Get() *GetTimeSeriesMinusDI200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMinusDI200ResponseMetaIndicator) Set(val *GetTimeSeriesMinusDI200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinusDI200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinusDI200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinusDI200ResponseMetaIndicator(val *GetTimeSeriesMinusDI200ResponseMetaIndicator) *NullableGetTimeSeriesMinusDI200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMinusDI200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinusDI200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinusDI200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
