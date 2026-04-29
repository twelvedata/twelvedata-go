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

// checks if the GetTimeSeriesAroon200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAroon200ResponseMetaIndicator{}

// GetTimeSeriesAroon200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesAroon200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesAroon200ResponseMetaIndicator GetTimeSeriesAroon200ResponseMetaIndicator

// NewGetTimeSeriesAroon200ResponseMetaIndicator instantiates a new GetTimeSeriesAroon200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAroon200ResponseMetaIndicator(name string, timePeriod int64) *GetTimeSeriesAroon200ResponseMetaIndicator {
	this := GetTimeSeriesAroon200ResponseMetaIndicator{}
	this.Name = name
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesAroon200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesAroon200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAroon200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesAroon200ResponseMetaIndicator {
	this := GetTimeSeriesAroon200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesAroon200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesAroon200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesAroon200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAroon200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesAroon200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesAroon200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAroon200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesAroon200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAroon200ResponseMetaIndicator := _GetTimeSeriesAroon200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAroon200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAroon200ResponseMetaIndicator(varGetTimeSeriesAroon200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesAroon200ResponseMetaIndicator struct {
	value *GetTimeSeriesAroon200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesAroon200ResponseMetaIndicator) Get() *GetTimeSeriesAroon200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesAroon200ResponseMetaIndicator) Set(val *GetTimeSeriesAroon200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAroon200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAroon200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAroon200ResponseMetaIndicator(val *GetTimeSeriesAroon200ResponseMetaIndicator) *NullableGetTimeSeriesAroon200ResponseMetaIndicator {
	return &NullableGetTimeSeriesAroon200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAroon200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAroon200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
