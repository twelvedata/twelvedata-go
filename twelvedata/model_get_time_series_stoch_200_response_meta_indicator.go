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

// checks if the GetTimeSeriesStoch200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStoch200ResponseMetaIndicator{}

// GetTimeSeriesStoch200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesStoch200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The time period for the fast %K line in the Stochastic Oscillator
	FastKPeriod int64 `json:"fast_k_period"`
	// The time period for the slow %K line in the Stochastic Oscillator
	SlowKPeriod int64 `json:"slow_k_period"`
	// The time period for the slow %D line in the Stochastic Oscillator
	SlowDPeriod int64 `json:"slow_d_period"`
	// The type of slow %K Moving Average used
	SlowKmaType string `json:"slow_kma_type"`
	// The type of slow Displaced Moving Average used
	SlowDmaType string `json:"slow_dma_type"`
}

type _GetTimeSeriesStoch200ResponseMetaIndicator GetTimeSeriesStoch200ResponseMetaIndicator

// NewGetTimeSeriesStoch200ResponseMetaIndicator instantiates a new GetTimeSeriesStoch200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStoch200ResponseMetaIndicator(name string, fastKPeriod int64, slowKPeriod int64, slowDPeriod int64, slowKmaType string, slowDmaType string) *GetTimeSeriesStoch200ResponseMetaIndicator {
	this := GetTimeSeriesStoch200ResponseMetaIndicator{}
	this.Name = name
	this.FastKPeriod = fastKPeriod
	this.SlowKPeriod = slowKPeriod
	this.SlowDPeriod = slowDPeriod
	this.SlowKmaType = slowKmaType
	this.SlowDmaType = slowDmaType
	return &this
}

// NewGetTimeSeriesStoch200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesStoch200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStoch200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesStoch200ResponseMetaIndicator {
	this := GetTimeSeriesStoch200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetFastKPeriod returns the FastKPeriod field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetFastKPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastKPeriod
}

// GetFastKPeriodOk returns a tuple with the FastKPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetFastKPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastKPeriod, true
}

// SetFastKPeriod sets field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) SetFastKPeriod(v int64) {
	o.FastKPeriod = v
}

// GetSlowKPeriod returns the SlowKPeriod field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowKPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowKPeriod
}

// GetSlowKPeriodOk returns a tuple with the SlowKPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowKPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowKPeriod, true
}

// SetSlowKPeriod sets field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) SetSlowKPeriod(v int64) {
	o.SlowKPeriod = v
}

// GetSlowDPeriod returns the SlowDPeriod field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowDPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SlowDPeriod
}

// GetSlowDPeriodOk returns a tuple with the SlowDPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowDPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowDPeriod, true
}

// SetSlowDPeriod sets field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) SetSlowDPeriod(v int64) {
	o.SlowDPeriod = v
}

// GetSlowKmaType returns the SlowKmaType field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowKmaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlowKmaType
}

// GetSlowKmaTypeOk returns a tuple with the SlowKmaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowKmaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowKmaType, true
}

// SetSlowKmaType sets field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) SetSlowKmaType(v string) {
	o.SlowKmaType = v
}

// GetSlowDmaType returns the SlowDmaType field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowDmaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlowDmaType
}

// GetSlowDmaTypeOk returns a tuple with the SlowDmaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) GetSlowDmaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowDmaType, true
}

// SetSlowDmaType sets field value
func (o *GetTimeSeriesStoch200ResponseMetaIndicator) SetSlowDmaType(v string) {
	o.SlowDmaType = v
}

func (o GetTimeSeriesStoch200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStoch200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fast_k_period"] = o.FastKPeriod
	toSerialize["slow_k_period"] = o.SlowKPeriod
	toSerialize["slow_d_period"] = o.SlowDPeriod
	toSerialize["slow_kma_type"] = o.SlowKmaType
	toSerialize["slow_dma_type"] = o.SlowDmaType
	return toSerialize, nil
}

func (o *GetTimeSeriesStoch200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fast_k_period",
		"slow_k_period",
		"slow_d_period",
		"slow_kma_type",
		"slow_dma_type",
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

	varGetTimeSeriesStoch200ResponseMetaIndicator := _GetTimeSeriesStoch200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStoch200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStoch200ResponseMetaIndicator(varGetTimeSeriesStoch200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesStoch200ResponseMetaIndicator struct {
	value *GetTimeSeriesStoch200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesStoch200ResponseMetaIndicator) Get() *GetTimeSeriesStoch200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesStoch200ResponseMetaIndicator) Set(val *GetTimeSeriesStoch200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStoch200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStoch200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStoch200ResponseMetaIndicator(val *GetTimeSeriesStoch200ResponseMetaIndicator) *NullableGetTimeSeriesStoch200ResponseMetaIndicator {
	return &NullableGetTimeSeriesStoch200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStoch200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStoch200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
