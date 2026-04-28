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

// checks if the GetTimeSeriesMacdExt200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMacdExt200ResponseMetaIndicator{}

// GetTimeSeriesMacdExt200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMacdExt200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The shorter time period for calculation
	FastPeriod int64 `json:"fast_period"`
	// The type of fast moving average used in the calculation
	FastMaType string `json:"fast_ma_type"`
	// The longer time period for calculation
	SlowPeriod int64 `json:"slow_period"`
	// The type of slow moving average used in the calculation
	SlowMaType string `json:"slow_ma_type"`
	// The time period used for generating the signal line
	SignalPeriod int64 `json:"signal_period"`
	// The type of moving average used for generating the signal line
	SignalMaType string `json:"signal_ma_type"`
}

type _GetTimeSeriesMacdExt200ResponseMetaIndicator GetTimeSeriesMacdExt200ResponseMetaIndicator

// NewGetTimeSeriesMacdExt200ResponseMetaIndicator instantiates a new GetTimeSeriesMacdExt200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMacdExt200ResponseMetaIndicator(name string, seriesType string, fastPeriod int64, fastMaType string, slowPeriod int64, slowMaType string, signalPeriod int64, signalMaType string) *GetTimeSeriesMacdExt200ResponseMetaIndicator {
	this := GetTimeSeriesMacdExt200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.FastPeriod = fastPeriod
	this.FastMaType = fastMaType
	this.SlowPeriod = slowPeriod
	this.SlowMaType = slowMaType
	this.SignalPeriod = signalPeriod
	this.SignalMaType = signalMaType
	return &this
}

// NewGetTimeSeriesMacdExt200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMacdExt200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMacdExt200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMacdExt200ResponseMetaIndicator {
	this := GetTimeSeriesMacdExt200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetFastPeriod returns the FastPeriod field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastPeriod
}

// GetFastPeriodOk returns a tuple with the FastPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastPeriod, true
}

// SetFastPeriod sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetFastPeriod(v int64) {
	o.FastPeriod = v
}

// GetFastMaType returns the FastMaType field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FastMaType
}

// GetFastMaTypeOk returns a tuple with the FastMaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastMaType, true
}

// SetFastMaType sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetFastMaType(v string) {
	o.FastMaType = v
}

// GetSlowPeriod returns the SlowPeriod field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowPeriod
}

// GetSlowPeriodOk returns a tuple with the SlowPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowPeriod, true
}

// SetSlowPeriod sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSlowPeriod(v int64) {
	o.SlowPeriod = v
}

// GetSlowMaType returns the SlowMaType field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlowMaType
}

// GetSlowMaTypeOk returns a tuple with the SlowMaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowMaType, true
}

// SetSlowMaType sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSlowMaType(v string) {
	o.SlowMaType = v
}

// GetSignalPeriod returns the SignalPeriod field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SignalPeriod
}

// GetSignalPeriodOk returns a tuple with the SignalPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalPeriod, true
}

// SetSignalPeriod sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSignalPeriod(v int64) {
	o.SignalPeriod = v
}

// GetSignalMaType returns the SignalMaType field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignalMaType
}

// GetSignalMaTypeOk returns a tuple with the SignalMaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalMaType, true
}

// SetSignalMaType sets field value
func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSignalMaType(v string) {
	o.SignalMaType = v
}

func (o GetTimeSeriesMacdExt200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMacdExt200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["fast_period"] = o.FastPeriod
	toSerialize["fast_ma_type"] = o.FastMaType
	toSerialize["slow_period"] = o.SlowPeriod
	toSerialize["slow_ma_type"] = o.SlowMaType
	toSerialize["signal_period"] = o.SignalPeriod
	toSerialize["signal_ma_type"] = o.SignalMaType
	return toSerialize, nil
}

func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"fast_period",
		"fast_ma_type",
		"slow_period",
		"slow_ma_type",
		"signal_period",
		"signal_ma_type",
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

	varGetTimeSeriesMacdExt200ResponseMetaIndicator := _GetTimeSeriesMacdExt200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMacdExt200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMacdExt200ResponseMetaIndicator(varGetTimeSeriesMacdExt200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMacdExt200ResponseMetaIndicator struct {
	value *GetTimeSeriesMacdExt200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMacdExt200ResponseMetaIndicator) Get() *GetTimeSeriesMacdExt200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMacdExt200ResponseMetaIndicator) Set(val *GetTimeSeriesMacdExt200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMacdExt200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMacdExt200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMacdExt200ResponseMetaIndicator(val *GetTimeSeriesMacdExt200ResponseMetaIndicator) *NullableGetTimeSeriesMacdExt200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMacdExt200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMacdExt200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMacdExt200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
