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

// checks if the GetTimeSeriesPercentB200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPercentB200ResponseMetaIndicator{}

// GetTimeSeriesPercentB200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesPercentB200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The time period used for calculation in the indicator
	TimePeriod int64 `json:"time_period"`
	// The standard deviation applied in the calculation
	Sd float64 `json:"sd"`
	// The type of moving average used
	MaType string `json:"ma_type"`
}

type _GetTimeSeriesPercentB200ResponseMetaIndicator GetTimeSeriesPercentB200ResponseMetaIndicator

// NewGetTimeSeriesPercentB200ResponseMetaIndicator instantiates a new GetTimeSeriesPercentB200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPercentB200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, sd float64, maType string) *GetTimeSeriesPercentB200ResponseMetaIndicator {
	this := GetTimeSeriesPercentB200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	this.Sd = sd
	this.MaType = maType
	return &this
}

// NewGetTimeSeriesPercentB200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesPercentB200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPercentB200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesPercentB200ResponseMetaIndicator {
	this := GetTimeSeriesPercentB200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

// GetSd returns the Sd field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetSd() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Sd
}

// GetSdOk returns a tuple with the Sd field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetSdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sd, true
}

// SetSd sets field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) SetSd(v float64) {
	o.Sd = v
}

// GetMaType returns the MaType field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaType
}

// GetMaTypeOk returns a tuple with the MaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) GetMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaType, true
}

// SetMaType sets field value
func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) SetMaType(v string) {
	o.MaType = v
}

func (o GetTimeSeriesPercentB200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPercentB200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["sd"] = o.Sd
	toSerialize["ma_type"] = o.MaType
	return toSerialize, nil
}

func (o *GetTimeSeriesPercentB200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesPercentB200ResponseMetaIndicator := _GetTimeSeriesPercentB200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesPercentB200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPercentB200ResponseMetaIndicator(varGetTimeSeriesPercentB200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesPercentB200ResponseMetaIndicator struct {
	value *GetTimeSeriesPercentB200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesPercentB200ResponseMetaIndicator) Get() *GetTimeSeriesPercentB200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesPercentB200ResponseMetaIndicator) Set(val *GetTimeSeriesPercentB200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPercentB200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPercentB200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPercentB200ResponseMetaIndicator(val *GetTimeSeriesPercentB200ResponseMetaIndicator) *NullableGetTimeSeriesPercentB200ResponseMetaIndicator {
	return &NullableGetTimeSeriesPercentB200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPercentB200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPercentB200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
