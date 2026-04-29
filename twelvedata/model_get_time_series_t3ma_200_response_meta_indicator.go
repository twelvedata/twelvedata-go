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

// checks if the GetTimeSeriesT3ma200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesT3ma200ResponseMetaIndicator{}

// GetTimeSeriesT3ma200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesT3ma200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
	// The time period used for calculation in the indicator
	TimePeriod int64 `json:"time_period"`
	// The factor used to adjust the indicator's volatility
	VFactor float64 `json:"v_factor"`
}

type _GetTimeSeriesT3ma200ResponseMetaIndicator GetTimeSeriesT3ma200ResponseMetaIndicator

// NewGetTimeSeriesT3ma200ResponseMetaIndicator instantiates a new GetTimeSeriesT3ma200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesT3ma200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, vFactor float64) *GetTimeSeriesT3ma200ResponseMetaIndicator {
	this := GetTimeSeriesT3ma200ResponseMetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	this.TimePeriod = timePeriod
	this.VFactor = vFactor
	return &this
}

// NewGetTimeSeriesT3ma200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesT3ma200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesT3ma200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesT3ma200ResponseMetaIndicator {
	this := GetTimeSeriesT3ma200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

// GetTimePeriod returns the TimePeriod field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetTimePeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimePeriod, true
}

// SetTimePeriod sets field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetTimePeriod(v int64) {
	o.TimePeriod = v
}

// GetVFactor returns the VFactor field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetVFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.VFactor
}

// GetVFactorOk returns a tuple with the VFactor field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetVFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VFactor, true
}

// SetVFactor sets field value
func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetVFactor(v float64) {
	o.VFactor = v
}

func (o GetTimeSeriesT3ma200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesT3ma200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	toSerialize["time_period"] = o.TimePeriod
	toSerialize["v_factor"] = o.VFactor
	return toSerialize, nil
}

func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
		"time_period",
		"v_factor",
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

	varGetTimeSeriesT3ma200ResponseMetaIndicator := _GetTimeSeriesT3ma200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesT3ma200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesT3ma200ResponseMetaIndicator(varGetTimeSeriesT3ma200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesT3ma200ResponseMetaIndicator struct {
	value *GetTimeSeriesT3ma200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesT3ma200ResponseMetaIndicator) Get() *GetTimeSeriesT3ma200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesT3ma200ResponseMetaIndicator) Set(val *GetTimeSeriesT3ma200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesT3ma200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesT3ma200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesT3ma200ResponseMetaIndicator(val *GetTimeSeriesT3ma200ResponseMetaIndicator) *NullableGetTimeSeriesT3ma200ResponseMetaIndicator {
	return &NullableGetTimeSeriesT3ma200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesT3ma200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesT3ma200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
