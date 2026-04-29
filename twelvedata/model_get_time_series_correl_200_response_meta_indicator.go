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

// checks if the GetTimeSeriesCorrel200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCorrel200ResponseMetaIndicator{}

// GetTimeSeriesCorrel200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesCorrel200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type used as the first part of technical indicator
	SeriesType1 string `json:"series_type_1"`
	// Price type used as the second part of technical indicator
	SeriesType2 string `json:"series_type_2"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesCorrel200ResponseMetaIndicator GetTimeSeriesCorrel200ResponseMetaIndicator

// NewGetTimeSeriesCorrel200ResponseMetaIndicator instantiates a new GetTimeSeriesCorrel200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCorrel200ResponseMetaIndicator(name string, seriesType1 string, seriesType2 string, timePeriod int64) *GetTimeSeriesCorrel200ResponseMetaIndicator {
	this := GetTimeSeriesCorrel200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType1 = seriesType1
	this.SeriesType2 = seriesType2
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesCorrel200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesCorrel200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCorrel200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesCorrel200ResponseMetaIndicator {
	this := GetTimeSeriesCorrel200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType1 returns the SeriesType1 field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetSeriesType1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType1
}

// GetSeriesType1Ok returns a tuple with the SeriesType1 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetSeriesType1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType1, true
}

// SetSeriesType1 sets field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) SetSeriesType1(v string) {
	o.SeriesType1 = v
}

// GetSeriesType2 returns the SeriesType2 field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetSeriesType2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType2
}

// GetSeriesType2Ok returns a tuple with the SeriesType2 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetSeriesType2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType2, true
}

// SetSeriesType2 sets field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) SetSeriesType2(v string) {
	o.SeriesType2 = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesCorrel200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCorrel200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type_1"] = o.SeriesType1
	toSerialize["series_type_2"] = o.SeriesType2
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesCorrel200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type_1",
		"series_type_2",
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

	varGetTimeSeriesCorrel200ResponseMetaIndicator := _GetTimeSeriesCorrel200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCorrel200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCorrel200ResponseMetaIndicator(varGetTimeSeriesCorrel200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesCorrel200ResponseMetaIndicator struct {
	value *GetTimeSeriesCorrel200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesCorrel200ResponseMetaIndicator) Get() *GetTimeSeriesCorrel200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesCorrel200ResponseMetaIndicator) Set(val *GetTimeSeriesCorrel200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCorrel200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCorrel200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCorrel200ResponseMetaIndicator(val *GetTimeSeriesCorrel200ResponseMetaIndicator) *NullableGetTimeSeriesCorrel200ResponseMetaIndicator {
	return &NullableGetTimeSeriesCorrel200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCorrel200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCorrel200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
