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

// checks if the GetTimeSeriesStdDev200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStdDev200ResponseMetaIndicator{}

// GetTimeSeriesStdDev200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesStdDev200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The time period used for calculation in the indicator
	TimePeriod int64 `json:"time_period"`
	// The standard deviation applied in the calculation
	Sd float64 `json:"sd"`
}

type _GetTimeSeriesStdDev200ResponseMetaIndicator GetTimeSeriesStdDev200ResponseMetaIndicator

// NewGetTimeSeriesStdDev200ResponseMetaIndicator instantiates a new GetTimeSeriesStdDev200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStdDev200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, sd float64) *GetTimeSeriesStdDev200ResponseMetaIndicator {
	this := GetTimeSeriesStdDev200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	this.Sd = sd
	return &this
}

// NewGetTimeSeriesStdDev200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesStdDev200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStdDev200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesStdDev200ResponseMetaIndicator {
	this := GetTimeSeriesStdDev200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

// GetSd returns the Sd field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSd() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Sd
}

// GetSdOk returns a tuple with the Sd field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sd, true
}

// SetSd sets field value
func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetSd(v float64) {
	o.Sd = v
}

func (o GetTimeSeriesStdDev200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStdDev200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["sd"] = o.Sd
	return toSerialize, nil
}

func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"time_period",
		"sd",
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

	varGetTimeSeriesStdDev200ResponseMetaIndicator := _GetTimeSeriesStdDev200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStdDev200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStdDev200ResponseMetaIndicator(varGetTimeSeriesStdDev200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesStdDev200ResponseMetaIndicator struct {
	value *GetTimeSeriesStdDev200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesStdDev200ResponseMetaIndicator) Get() *GetTimeSeriesStdDev200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesStdDev200ResponseMetaIndicator) Set(val *GetTimeSeriesStdDev200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStdDev200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStdDev200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStdDev200ResponseMetaIndicator(val *GetTimeSeriesStdDev200ResponseMetaIndicator) *NullableGetTimeSeriesStdDev200ResponseMetaIndicator {
	return &NullableGetTimeSeriesStdDev200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStdDev200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStdDev200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
