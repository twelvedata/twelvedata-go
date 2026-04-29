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

// checks if the GetTimeSeriesKeltner200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKeltner200ResponseMetaIndicator{}

// GetTimeSeriesKeltner200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesKeltner200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
	// The time period used for calculating the Average True Range
	AtrTimePeriod int64 `json:"atr_time_period"`
	// The factor used to adjust the indicator's sensitivity
	Multiplier int64 `json:"multiplier"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The type of moving average used
	MaType string `json:"ma_type"`
}

type _GetTimeSeriesKeltner200ResponseMetaIndicator GetTimeSeriesKeltner200ResponseMetaIndicator

// NewGetTimeSeriesKeltner200ResponseMetaIndicator instantiates a new GetTimeSeriesKeltner200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKeltner200ResponseMetaIndicator(name string, timePeriod int64, atrTimePeriod int64, multiplier int64, seriesType string, maType string) *GetTimeSeriesKeltner200ResponseMetaIndicator {
	this := GetTimeSeriesKeltner200ResponseMetaIndicator{}
	this.Name = name
	this.TimePeriod = timePeriod
	this.AtrTimePeriod = atrTimePeriod
	this.Multiplier = multiplier
	this.SeriesType = seriesType
	this.MaType = maType
	return &this
}

// NewGetTimeSeriesKeltner200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesKeltner200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKeltner200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesKeltner200ResponseMetaIndicator {
	this := GetTimeSeriesKeltner200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

// GetAtrTimePeriod returns the AtrTimePeriod field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetAtrTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AtrTimePeriod
}

// GetAtrTimePeriodOk returns a tuple with the AtrTimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetAtrTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AtrTimePeriod, true
}

// SetAtrTimePeriod sets field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetAtrTimePeriod(v int64) {
	o.AtrTimePeriod = v
}

// GetMultiplier returns the Multiplier field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMultiplier() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMultiplierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Multiplier, true
}

// SetMultiplier sets field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetMultiplier(v int64) {
	o.Multiplier = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetMaType returns the MaType field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaType
}

// GetMaTypeOk returns a tuple with the MaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) GetMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaType, true
}

// SetMaType sets field value
func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) SetMaType(v string) {
	o.MaType = v
}

func (o GetTimeSeriesKeltner200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKeltner200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["atr_time_period"] = o.AtrTimePeriod
	toSerialize["multiplier"] = o.Multiplier
	toSerialize["series_type"] = o.SeriesType
	toSerialize["ma_type"] = o.MaType
	return toSerialize, nil
}

func (o *GetTimeSeriesKeltner200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"time_period",
		"atr_time_period",
		"multiplier",
		"series_type",
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

	varGetTimeSeriesKeltner200ResponseMetaIndicator := _GetTimeSeriesKeltner200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesKeltner200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKeltner200ResponseMetaIndicator(varGetTimeSeriesKeltner200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesKeltner200ResponseMetaIndicator struct {
	value *GetTimeSeriesKeltner200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesKeltner200ResponseMetaIndicator) Get() *GetTimeSeriesKeltner200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesKeltner200ResponseMetaIndicator) Set(val *GetTimeSeriesKeltner200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKeltner200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKeltner200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKeltner200ResponseMetaIndicator(val *GetTimeSeriesKeltner200ResponseMetaIndicator) *NullableGetTimeSeriesKeltner200ResponseMetaIndicator {
	return &NullableGetTimeSeriesKeltner200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKeltner200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKeltner200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
