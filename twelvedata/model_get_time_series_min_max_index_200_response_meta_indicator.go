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

// checks if the GetTimeSeriesMinMaxIndex200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinMaxIndex200ResponseMetaIndicator{}

// GetTimeSeriesMinMaxIndex200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesMinMaxIndex200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesMinMaxIndex200ResponseMetaIndicator GetTimeSeriesMinMaxIndex200ResponseMetaIndicator

// NewGetTimeSeriesMinMaxIndex200ResponseMetaIndicator instantiates a new GetTimeSeriesMinMaxIndex200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinMaxIndex200ResponseMetaIndicator(name string, seriesType string, timePeriod int64) *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator {
	this := GetTimeSeriesMinMaxIndex200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesMinMaxIndex200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMinMaxIndex200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinMaxIndex200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator {
	this := GetTimeSeriesMinMaxIndex200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesMinMaxIndex200ResponseMetaIndicator := _GetTimeSeriesMinMaxIndex200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinMaxIndex200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinMaxIndex200ResponseMetaIndicator(varGetTimeSeriesMinMaxIndex200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator struct {
	value *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator) Get() *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator) Set(val *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator(val *GetTimeSeriesMinMaxIndex200ResponseMetaIndicator) *NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator {
	return &NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinMaxIndex200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
