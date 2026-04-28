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

// checks if the GetTimeSeriesMacdSlope200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacdSlope200ResponseMetaIndicator{}

// GetTimeSeriesMacdSlope200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMacdSlope200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The shorter time period for calculation
	FastPeriod int64 `json:"fast_period"`
	// The longer time period for calculation
	SlowPeriod int64 `json:"slow_period"`
	// The time period used for generating the signal line
	SignalPeriod int64 `json:"signal_period"`
	// The time period used for calculation in the indicator
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesMacdSlope200ResponseMetaIndicator GetTimeSeriesMacdSlope200ResponseMetaIndicator

// NewGetTimeSeriesMacdSlope200ResponseMetaIndicator instantiates a new GetTimeSeriesMacdSlope200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacdSlope200ResponseMetaIndicator(name string, seriesType string, fastPeriod int64, slowPeriod int64, signalPeriod int64, timePeriod int64) *GetTimeSeriesMacdSlope200ResponseMetaIndicator {
	this := GetTimeSeriesMacdSlope200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.FastPeriod = fastPeriod
	this.SlowPeriod = slowPeriod
	this.SignalPeriod = signalPeriod
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesMacdSlope200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMacdSlope200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacdSlope200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMacdSlope200ResponseMetaIndicator {
	this := GetTimeSeriesMacdSlope200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetFastPeriod returns the FastPeriod field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetFastPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastPeriod
}

// GetFastPeriodOk returns a tuple with the FastPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetFastPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastPeriod, true
}

// SetFastPeriod sets field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) SetFastPeriod(v int64) {
	o.FastPeriod = v
}

// GetSlowPeriod returns the SlowPeriod field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetSlowPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowPeriod
}

// GetSlowPeriodOk returns a tuple with the SlowPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetSlowPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowPeriod, true
}

// SetSlowPeriod sets field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) SetSlowPeriod(v int64) {
	o.SlowPeriod = v
}

// GetSignalPeriod returns the SignalPeriod field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetSignalPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SignalPeriod
}

// GetSignalPeriodOk returns a tuple with the SignalPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetSignalPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalPeriod, true
}

// SetSignalPeriod sets field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) SetSignalPeriod(v int64) {
	o.SignalPeriod = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesMacdSlope200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacdSlope200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["fast_period"] = o.FastPeriod
	toSerialize["slow_period"] = o.SlowPeriod
	toSerialize["signal_period"] = o.SignalPeriod
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesMacdSlope200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"fast_period",
		"slow_period",
		"signal_period",
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

	varGetTimeSeriesMacdSlope200ResponseMetaIndicator := _GetTimeSeriesMacdSlope200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMacdSlope200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacdSlope200ResponseMetaIndicator(varGetTimeSeriesMacdSlope200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator struct {
	value *GetTimeSeriesMacdSlope200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator) Get() *GetTimeSeriesMacdSlope200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator) Set(val *GetTimeSeriesMacdSlope200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacdSlope200ResponseMetaIndicator(val *GetTimeSeriesMacdSlope200ResponseMetaIndicator) *NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacdSlope200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
