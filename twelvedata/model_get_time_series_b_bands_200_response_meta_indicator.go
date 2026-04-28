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

// checks if the GetTimeSeriesBBands200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesBBands200ResponseMetaIndicator{}

// GetTimeSeriesBBands200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesBBands200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
	// Number of standard deviations
	Sd float64 `json:"sd"`
	// Moving average type
	MaType string `json:"ma_type"`
}

type _GetTimeSeriesBBands200ResponseMetaIndicator GetTimeSeriesBBands200ResponseMetaIndicator

// NewGetTimeSeriesBBands200ResponseMetaIndicator instantiates a new GetTimeSeriesBBands200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesBBands200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, sd float64, maType string) *GetTimeSeriesBBands200ResponseMetaIndicator {
	this := GetTimeSeriesBBands200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	this.Sd = sd
	this.MaType = maType
	return &this
}

// NewGetTimeSeriesBBands200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesBBands200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesBBands200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesBBands200ResponseMetaIndicator {
	this := GetTimeSeriesBBands200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

// GetSd returns the Sd field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetSd() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Sd
}

// GetSdOk returns a tuple with the Sd field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetSdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sd, true
}

// SetSd sets field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) SetSd(v float64) {
	o.Sd = v
}

// GetMaType returns the MaType field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaType
}

// GetMaTypeOk returns a tuple with the MaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) GetMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaType, true
}

// SetMaType sets field value
func (o *GetTimeSeriesBBands200ResponseMetaIndicator) SetMaType(v string) {
	o.MaType = v
}

func (o GetTimeSeriesBBands200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesBBands200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["sd"] = o.Sd
	toSerialize["ma_type"] = o.MaType
	return toSerialize, nil
}

func (o *GetTimeSeriesBBands200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"time_period",
		"sd",
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

	varGetTimeSeriesBBands200ResponseMetaIndicator := _GetTimeSeriesBBands200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesBBands200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesBBands200ResponseMetaIndicator(varGetTimeSeriesBBands200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesBBands200ResponseMetaIndicator struct {
	value *GetTimeSeriesBBands200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesBBands200ResponseMetaIndicator) Get() *GetTimeSeriesBBands200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesBBands200ResponseMetaIndicator) Set(val *GetTimeSeriesBBands200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesBBands200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesBBands200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesBBands200ResponseMetaIndicator(val *GetTimeSeriesBBands200ResponseMetaIndicator) *NullableGetTimeSeriesBBands200ResponseMetaIndicator {
	return &NullableGetTimeSeriesBBands200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesBBands200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesBBands200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
