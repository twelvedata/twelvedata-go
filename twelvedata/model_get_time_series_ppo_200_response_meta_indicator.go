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

// checks if the GetTimeSeriesPpo200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPpo200ResponseMetaIndicator{}

// GetTimeSeriesPpo200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesPpo200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The shorter time period for calculation
	FastPeriod int64 `json:"fast_period"`
	// The longer time period for calculation
	SlowPeriod int64 `json:"slow_period"`
	// The type of moving average used
	MaType string `json:"ma_type"`
}

type _GetTimeSeriesPpo200ResponseMetaIndicator GetTimeSeriesPpo200ResponseMetaIndicator

// NewGetTimeSeriesPpo200ResponseMetaIndicator instantiates a new GetTimeSeriesPpo200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPpo200ResponseMetaIndicator(name string, seriesType string, fastPeriod int64, slowPeriod int64, maType string) *GetTimeSeriesPpo200ResponseMetaIndicator {
	this := GetTimeSeriesPpo200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.FastPeriod = fastPeriod
	this.SlowPeriod = slowPeriod
	this.MaType = maType
	return &this
}

// NewGetTimeSeriesPpo200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesPpo200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPpo200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesPpo200ResponseMetaIndicator {
	this := GetTimeSeriesPpo200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetFastPeriod returns the FastPeriod field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetFastPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastPeriod
}

// GetFastPeriodOk returns a tuple with the FastPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetFastPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastPeriod, true
}

// SetFastPeriod sets field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) SetFastPeriod(v int64) {
	o.FastPeriod = v
}

// GetSlowPeriod returns the SlowPeriod field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetSlowPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowPeriod
}

// GetSlowPeriodOk returns a tuple with the SlowPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetSlowPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowPeriod, true
}

// SetSlowPeriod sets field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) SetSlowPeriod(v int64) {
	o.SlowPeriod = v
}

// GetMaType returns the MaType field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaType
}

// GetMaTypeOk returns a tuple with the MaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) GetMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaType, true
}

// SetMaType sets field value
func (o *GetTimeSeriesPpo200ResponseMetaIndicator) SetMaType(v string) {
	o.MaType = v
}

func (o GetTimeSeriesPpo200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPpo200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["fast_period"] = o.FastPeriod
	toSerialize["slow_period"] = o.SlowPeriod
	toSerialize["ma_type"] = o.MaType
	return toSerialize, nil
}

func (o *GetTimeSeriesPpo200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"fast_period",
		"slow_period",
		"ma_type",
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

	varGetTimeSeriesPpo200ResponseMetaIndicator := _GetTimeSeriesPpo200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesPpo200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPpo200ResponseMetaIndicator(varGetTimeSeriesPpo200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesPpo200ResponseMetaIndicator struct {
	value *GetTimeSeriesPpo200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesPpo200ResponseMetaIndicator) Get() *GetTimeSeriesPpo200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesPpo200ResponseMetaIndicator) Set(val *GetTimeSeriesPpo200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPpo200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPpo200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPpo200ResponseMetaIndicator(val *GetTimeSeriesPpo200ResponseMetaIndicator) *NullableGetTimeSeriesPpo200ResponseMetaIndicator {
	return &NullableGetTimeSeriesPpo200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPpo200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPpo200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
