// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesStochF200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStochF200ResponseMetaIndicator{}

// GetTimeSeriesStochF200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesStochF200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The fast_k period used for calculation in the indicator
	FastKPeriod int64 `json:"fast_k_period"`
	// The fast_d period used for calculation in the indicator
	FastDPeriod int64 `json:"fast_d_period"`
	// The type of fast Displaced Moving Average used
	FastDmaType string `json:"fast_dma_type"`
}

type _GetTimeSeriesStochF200ResponseMetaIndicator GetTimeSeriesStochF200ResponseMetaIndicator

// NewGetTimeSeriesStochF200ResponseMetaIndicator instantiates a new GetTimeSeriesStochF200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStochF200ResponseMetaIndicator(name string, fastKPeriod int64, fastDPeriod int64, fastDmaType string) *GetTimeSeriesStochF200ResponseMetaIndicator {
	this := GetTimeSeriesStochF200ResponseMetaIndicator{}
	this.Name = name
	this.FastKPeriod = fastKPeriod
	this.FastDPeriod = fastDPeriod
	this.FastDmaType = fastDmaType
	return &this
}

// NewGetTimeSeriesStochF200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesStochF200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStochF200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesStochF200ResponseMetaIndicator {
	this := GetTimeSeriesStochF200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetFastKPeriod returns the FastKPeriod field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetFastKPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastKPeriod
}

// GetFastKPeriodOk returns a tuple with the FastKPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetFastKPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastKPeriod, true
}

// SetFastKPeriod sets field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) SetFastKPeriod(v int64) {
	o.FastKPeriod = v
}

// GetFastDPeriod returns the FastDPeriod field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetFastDPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FastDPeriod
}

// GetFastDPeriodOk returns a tuple with the FastDPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetFastDPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastDPeriod, true
}

// SetFastDPeriod sets field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) SetFastDPeriod(v int64) {
	o.FastDPeriod = v
}

// GetFastDmaType returns the FastDmaType field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetFastDmaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FastDmaType
}

// GetFastDmaTypeOk returns a tuple with the FastDmaType field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) GetFastDmaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FastDmaType, true
}

// SetFastDmaType sets field value
func (o *GetTimeSeriesStochF200ResponseMetaIndicator) SetFastDmaType(v string) {
	o.FastDmaType = v
}

func (o GetTimeSeriesStochF200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStochF200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["fast_k_period"] = o.FastKPeriod
	toSerialize["fast_d_period"] = o.FastDPeriod
	toSerialize["fast_dma_type"] = o.FastDmaType
	return toSerialize, nil
}

func (o *GetTimeSeriesStochF200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"fast_k_period",
		"fast_d_period",
		"fast_dma_type",
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

	varGetTimeSeriesStochF200ResponseMetaIndicator := _GetTimeSeriesStochF200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStochF200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStochF200ResponseMetaIndicator(varGetTimeSeriesStochF200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesStochF200ResponseMetaIndicator struct {
	value *GetTimeSeriesStochF200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesStochF200ResponseMetaIndicator) Get() *GetTimeSeriesStochF200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesStochF200ResponseMetaIndicator) Set(val *GetTimeSeriesStochF200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStochF200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStochF200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStochF200ResponseMetaIndicator(val *GetTimeSeriesStochF200ResponseMetaIndicator) *NullableGetTimeSeriesStochF200ResponseMetaIndicator {
	return &NullableGetTimeSeriesStochF200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStochF200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStochF200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
