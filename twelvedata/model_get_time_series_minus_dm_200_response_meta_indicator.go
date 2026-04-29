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

// checks if the GetTimeSeriesMinusDM200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinusDM200ResponseMetaIndicator{}

// GetTimeSeriesMinusDM200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMinusDM200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The time period used for calculation in the indicator
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesMinusDM200ResponseMetaIndicator GetTimeSeriesMinusDM200ResponseMetaIndicator

// NewGetTimeSeriesMinusDM200ResponseMetaIndicator instantiates a new GetTimeSeriesMinusDM200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinusDM200ResponseMetaIndicator(name string, timePeriod int64) *GetTimeSeriesMinusDM200ResponseMetaIndicator {
	this := GetTimeSeriesMinusDM200ResponseMetaIndicator{}
	this.Name = name
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesMinusDM200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMinusDM200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinusDM200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMinusDM200ResponseMetaIndicator {
	this := GetTimeSeriesMinusDM200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMinusDM200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDM200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMinusDM200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesMinusDM200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDM200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesMinusDM200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesMinusDM200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinusDM200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesMinusDM200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMinusDM200ResponseMetaIndicator := _GetTimeSeriesMinusDM200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinusDM200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinusDM200ResponseMetaIndicator(varGetTimeSeriesMinusDM200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMinusDM200ResponseMetaIndicator struct {
	value *GetTimeSeriesMinusDM200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMinusDM200ResponseMetaIndicator) Get() *GetTimeSeriesMinusDM200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMinusDM200ResponseMetaIndicator) Set(val *GetTimeSeriesMinusDM200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinusDM200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinusDM200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinusDM200ResponseMetaIndicator(val *GetTimeSeriesMinusDM200ResponseMetaIndicator) *NullableGetTimeSeriesMinusDM200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMinusDM200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinusDM200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinusDM200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
