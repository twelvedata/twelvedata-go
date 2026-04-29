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

// checks if the GetTimeSeriesLinearRegSlope200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLinearRegSlope200ResponseMetaIndicator{}

// GetTimeSeriesLinearRegSlope200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesLinearRegSlope200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// Number of periods to average over
	TimePeriod int64 `json:"time_period"`
}

type _GetTimeSeriesLinearRegSlope200ResponseMetaIndicator GetTimeSeriesLinearRegSlope200ResponseMetaIndicator

// NewGetTimeSeriesLinearRegSlope200ResponseMetaIndicator instantiates a new GetTimeSeriesLinearRegSlope200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLinearRegSlope200ResponseMetaIndicator(name string, seriesType string, timePeriod int64) *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator {
	this := GetTimeSeriesLinearRegSlope200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	return &this
}

// NewGetTimeSeriesLinearRegSlope200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesLinearRegSlope200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLinearRegSlope200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator {
	this := GetTimeSeriesLinearRegSlope200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

func (o GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesLinearRegSlope200ResponseMetaIndicator := _GetTimeSeriesLinearRegSlope200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesLinearRegSlope200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLinearRegSlope200ResponseMetaIndicator(varGetTimeSeriesLinearRegSlope200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator struct {
	value *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator) Get() *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator) Set(val *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator(val *GetTimeSeriesLinearRegSlope200ResponseMetaIndicator) *NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator {
	return &NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLinearRegSlope200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
