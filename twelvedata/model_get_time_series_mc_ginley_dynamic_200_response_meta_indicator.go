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

// checks if the GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator{}

// GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator

// NewGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator instantiates a new GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator(name string, timePeriod int64) *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator {
	this := GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator{}
	this.Name = name
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator {
	this := GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator := _GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator(varGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator struct {
	value *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) Get() *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) Set(val *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator(val *GetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) *NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
