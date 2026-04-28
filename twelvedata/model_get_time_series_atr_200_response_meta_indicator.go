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

// checks if the GetTimeSeriesAtr200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAtr200ResponseMetaIndicator{}

// GetTimeSeriesAtr200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesAtr200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesAtr200ResponseMetaIndicator GetTimeSeriesAtr200ResponseMetaIndicator

// NewGetTimeSeriesAtr200ResponseMetaIndicator instantiates a new GetTimeSeriesAtr200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAtr200ResponseMetaIndicator(name string, timePeriod int64) *GetTimeSeriesAtr200ResponseMetaIndicator {
	this := GetTimeSeriesAtr200ResponseMetaIndicator{}
	this.Name = name
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesAtr200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesAtr200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAtr200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesAtr200ResponseMetaIndicator {
	this := GetTimeSeriesAtr200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesAtr200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAtr200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesAtr200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesAtr200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAtr200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesAtr200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesAtr200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAtr200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesAtr200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAtr200ResponseMetaIndicator := _GetTimeSeriesAtr200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAtr200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAtr200ResponseMetaIndicator(varGetTimeSeriesAtr200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesAtr200ResponseMetaIndicator struct {
	value *GetTimeSeriesAtr200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesAtr200ResponseMetaIndicator) Get() *GetTimeSeriesAtr200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesAtr200ResponseMetaIndicator) Set(val *GetTimeSeriesAtr200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAtr200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAtr200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAtr200ResponseMetaIndicator(val *GetTimeSeriesAtr200ResponseMetaIndicator) *NullableGetTimeSeriesAtr200ResponseMetaIndicator {
	return &NullableGetTimeSeriesAtr200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAtr200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAtr200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
