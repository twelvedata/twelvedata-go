// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMa200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMa200ResponseMetaIndicator{}

// GetTimeSeriesMa200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMa200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
	// The type of moving average used
	MaType string `json:"ma_type"`
}

type _GetTimeSeriesMa200ResponseMetaIndicator GetTimeSeriesMa200ResponseMetaIndicator

// NewGetTimeSeriesMa200ResponseMetaIndicator instantiates a new GetTimeSeriesMa200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMa200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, maType string) *GetTimeSeriesMa200ResponseMetaIndicator {
	this := GetTimeSeriesMa200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	this.MaType = maType
	return &this
}

// NewGetTimeSeriesMa200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMa200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMa200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMa200ResponseMetaIndicator {
	this := GetTimeSeriesMa200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

// GetMaType returns the MaType field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetMaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaType
}

// GetMaTypeOk returns a tuple with the MaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetMaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaType, true
}

// SetMaType sets field value
func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetMaType(v string) {
	o.MaType = v
}

func (o GetTimeSeriesMa200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMa200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["ma_type"] = o.MaType
	return toSerialize, nil
}

func (o *GetTimeSeriesMa200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"time_period",
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

	varGetTimeSeriesMa200ResponseMetaIndicator := _GetTimeSeriesMa200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMa200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMa200ResponseMetaIndicator(varGetTimeSeriesMa200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMa200ResponseMetaIndicator struct {
	value *GetTimeSeriesMa200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMa200ResponseMetaIndicator) Get() *GetTimeSeriesMa200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMa200ResponseMetaIndicator) Set(val *GetTimeSeriesMa200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMa200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMa200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMa200ResponseMetaIndicator(val *GetTimeSeriesMa200ResponseMetaIndicator) *NullableGetTimeSeriesMa200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMa200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMa200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMa200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
