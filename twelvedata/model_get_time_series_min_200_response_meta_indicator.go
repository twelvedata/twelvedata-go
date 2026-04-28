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

// checks if the GetTimeSeriesMin200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMin200ResponseMetaIndicator{}

// GetTimeSeriesMin200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMin200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesMin200ResponseMetaIndicator GetTimeSeriesMin200ResponseMetaIndicator

// NewGetTimeSeriesMin200ResponseMetaIndicator instantiates a new GetTimeSeriesMin200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMin200ResponseMetaIndicator(name string, seriesType string, timePeriod int64) *GetTimeSeriesMin200ResponseMetaIndicator {
	this := GetTimeSeriesMin200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesMin200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMin200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMin200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMin200ResponseMetaIndicator {
	this := GetTimeSeriesMin200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMin200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMin200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesMin200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesMin200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesMin200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesMin200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesMin200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMin200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesMin200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
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

	varGetTimeSeriesMin200ResponseMetaIndicator := _GetTimeSeriesMin200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMin200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMin200ResponseMetaIndicator(varGetTimeSeriesMin200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMin200ResponseMetaIndicator struct {
	value *GetTimeSeriesMin200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMin200ResponseMetaIndicator) Get() *GetTimeSeriesMin200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMin200ResponseMetaIndicator) Set(val *GetTimeSeriesMin200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMin200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMin200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMin200ResponseMetaIndicator(val *GetTimeSeriesMin200ResponseMetaIndicator) *NullableGetTimeSeriesMin200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMin200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMin200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMin200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
