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

// checks if the GetTimeSeriesMacd200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacd200ResponseMetaIndicator{}

// GetTimeSeriesMacd200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMacd200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Fast period value
	FastPeriod int64 `json:"fast_period"`
	// Slow period value
	SlowPeriod int64 `json:"slow_period"`
	// Signal period value
	SignalPeriod int64 `json:"signal_period"`
}

type _GetTimeSeriesMacd200ResponseMetaIndicator GetTimeSeriesMacd200ResponseMetaIndicator

// NewGetTimeSeriesMacd200ResponseMetaIndicator instantiates a new GetTimeSeriesMacd200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacd200ResponseMetaIndicator(name string, seriesType string, fastPeriod int64, slowPeriod int64, signalPeriod int64) *GetTimeSeriesMacd200ResponseMetaIndicator {
	this := GetTimeSeriesMacd200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.FastPeriod = fastPeriod
	this.SlowPeriod = slowPeriod
	this.SignalPeriod = signalPeriod
	return &this
}

// NewGetTimeSeriesMacd200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMacd200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacd200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMacd200ResponseMetaIndicator {
	this := GetTimeSeriesMacd200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetFastPeriod returns the FastPeriod field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetFastPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastPeriod
}

// GetFastPeriodOk returns a tuple with the FastPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetFastPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastPeriod, true
}

// SetFastPeriod sets field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) SetFastPeriod(v int64) {
	o.FastPeriod = v
}

// GetSlowPeriod returns the SlowPeriod field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetSlowPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowPeriod
}

// GetSlowPeriodOk returns a tuple with the SlowPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetSlowPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowPeriod, true
}

// SetSlowPeriod sets field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) SetSlowPeriod(v int64) {
	o.SlowPeriod = v
}

// GetSignalPeriod returns the SignalPeriod field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetSignalPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SignalPeriod
}

// GetSignalPeriodOk returns a tuple with the SignalPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) GetSignalPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalPeriod, true
}

// SetSignalPeriod sets field value
func (o *GetTimeSeriesMacd200ResponseMetaIndicator) SetSignalPeriod(v int64) {
	o.SignalPeriod = v
}

func (o GetTimeSeriesMacd200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacd200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["fast_period"] = o.FastPeriod
	toSerialize["slow_period"] = o.SlowPeriod
	toSerialize["signal_period"] = o.SignalPeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesMacd200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"fast_period",
		"slow_period",
		"signal_period",
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

	varGetTimeSeriesMacd200ResponseMetaIndicator := _GetTimeSeriesMacd200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMacd200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacd200ResponseMetaIndicator(varGetTimeSeriesMacd200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMacd200ResponseMetaIndicator struct {
	value *GetTimeSeriesMacd200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMacd200ResponseMetaIndicator) Get() *GetTimeSeriesMacd200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMacd200ResponseMetaIndicator) Set(val *GetTimeSeriesMacd200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacd200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacd200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacd200ResponseMetaIndicator(val *GetTimeSeriesMacd200ResponseMetaIndicator) *NullableGetTimeSeriesMacd200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMacd200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacd200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacd200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
