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

// checks if the GetTimeSeriesKst200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKst200ResponseMetaIndicator{}

// GetTimeSeriesKst200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesKst200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The time period for the first Rate of Change calculation
	RocPeriod1 int64 `json:"roc_period_1"`
	// The time period for the second Rate of Change calculation
	RocPeriod2 int64 `json:"roc_period_2"`
	// The time period for the third Rate of Change calculation
	RocPeriod3 int64 `json:"roc_period_3"`
	// The time period for the forth Rate of Change calculation
	RocPeriod4 int64 `json:"roc_period_4"`
	// The time period for the first Simple Moving Average
	SmaPeriod1 int64 `json:"sma_period_1"`
	// The time period for the second Simple Moving Average
	SmaPeriod2 int64 `json:"sma_period_2"`
	// The time period for the third Simple Moving Average
	SmaPeriod3 int64 `json:"sma_period_3"`
	// The time period for the forth Simple Moving Average
	SmaPeriod4 int64 `json:"sma_period_4"`
	// The time period used for generating the signal line
	SignalPeriod int64 `json:"signal_period"`
}

type _GetTimeSeriesKst200ResponseMetaIndicator GetTimeSeriesKst200ResponseMetaIndicator

// NewGetTimeSeriesKst200ResponseMetaIndicator instantiates a new GetTimeSeriesKst200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKst200ResponseMetaIndicator(name string, rocPeriod1 int64, rocPeriod2 int64, rocPeriod3 int64, rocPeriod4 int64, smaPeriod1 int64, smaPeriod2 int64, smaPeriod3 int64, smaPeriod4 int64, signalPeriod int64) *GetTimeSeriesKst200ResponseMetaIndicator {
	this := GetTimeSeriesKst200ResponseMetaIndicator{}
	this.Name = name
	this.RocPeriod1 = rocPeriod1
	this.RocPeriod2 = rocPeriod2
	this.RocPeriod3 = rocPeriod3
	this.RocPeriod4 = rocPeriod4
	this.SmaPeriod1 = smaPeriod1
	this.SmaPeriod2 = smaPeriod2
	this.SmaPeriod3 = smaPeriod3
	this.SmaPeriod4 = smaPeriod4
	this.SignalPeriod = signalPeriod
	return &this
}

// NewGetTimeSeriesKst200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesKst200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKst200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesKst200ResponseMetaIndicator {
	this := GetTimeSeriesKst200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetRocPeriod1 returns the RocPeriod1 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod1() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RocPeriod1
}

// GetRocPeriod1Ok returns a tuple with the RocPeriod1 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod1Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RocPeriod1, true
}

// SetRocPeriod1 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod1(v int64) {
	o.RocPeriod1 = v
}

// GetRocPeriod2 returns the RocPeriod2 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod2() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RocPeriod2
}

// GetRocPeriod2Ok returns a tuple with the RocPeriod2 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod2Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RocPeriod2, true
}

// SetRocPeriod2 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod2(v int64) {
	o.RocPeriod2 = v
}

// GetRocPeriod3 returns the RocPeriod3 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod3() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RocPeriod3
}

// GetRocPeriod3Ok returns a tuple with the RocPeriod3 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod3Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RocPeriod3, true
}

// SetRocPeriod3 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod3(v int64) {
	o.RocPeriod3 = v
}

// GetRocPeriod4 returns the RocPeriod4 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod4() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RocPeriod4
}

// GetRocPeriod4Ok returns a tuple with the RocPeriod4 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetRocPeriod4Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RocPeriod4, true
}

// SetRocPeriod4 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetRocPeriod4(v int64) {
	o.RocPeriod4 = v
}

// GetSmaPeriod1 returns the SmaPeriod1 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod1() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SmaPeriod1
}

// GetSmaPeriod1Ok returns a tuple with the SmaPeriod1 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod1Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmaPeriod1, true
}

// SetSmaPeriod1 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod1(v int64) {
	o.SmaPeriod1 = v
}

// GetSmaPeriod2 returns the SmaPeriod2 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod2() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SmaPeriod2
}

// GetSmaPeriod2Ok returns a tuple with the SmaPeriod2 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod2Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmaPeriod2, true
}

// SetSmaPeriod2 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod2(v int64) {
	o.SmaPeriod2 = v
}

// GetSmaPeriod3 returns the SmaPeriod3 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod3() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SmaPeriod3
}

// GetSmaPeriod3Ok returns a tuple with the SmaPeriod3 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod3Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmaPeriod3, true
}

// SetSmaPeriod3 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod3(v int64) {
	o.SmaPeriod3 = v
}

// GetSmaPeriod4 returns the SmaPeriod4 field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod4() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SmaPeriod4
}

// GetSmaPeriod4Ok returns a tuple with the SmaPeriod4 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSmaPeriod4Ok() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmaPeriod4, true
}

// SetSmaPeriod4 sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSmaPeriod4(v int64) {
	o.SmaPeriod4 = v
}

// GetSignalPeriod returns the SignalPeriod field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSignalPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SignalPeriod
}

// GetSignalPeriodOk returns a tuple with the SignalPeriod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseMetaIndicator) GetSignalPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalPeriod, true
}

// SetSignalPeriod sets field value
func (o *GetTimeSeriesKst200ResponseMetaIndicator) SetSignalPeriod(v int64) {
	o.SignalPeriod = v
}

func (o GetTimeSeriesKst200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKst200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["roc_period_1"] = o.RocPeriod1
	toSerialize["roc_period_2"] = o.RocPeriod2
	toSerialize["roc_period_3"] = o.RocPeriod3
	toSerialize["roc_period_4"] = o.RocPeriod4
	toSerialize["sma_period_1"] = o.SmaPeriod1
	toSerialize["sma_period_2"] = o.SmaPeriod2
	toSerialize["sma_period_3"] = o.SmaPeriod3
	toSerialize["sma_period_4"] = o.SmaPeriod4
	toSerialize["signal_period"] = o.SignalPeriod
	return toSerialize, nil
}

func (o *GetTimeSeriesKst200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"roc_period_1",
		"roc_period_2",
		"roc_period_3",
		"roc_period_4",
		"sma_period_1",
		"sma_period_2",
		"sma_period_3",
		"sma_period_4",
		"signal_period",
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

	varGetTimeSeriesKst200ResponseMetaIndicator := _GetTimeSeriesKst200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesKst200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKst200ResponseMetaIndicator(varGetTimeSeriesKst200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesKst200ResponseMetaIndicator struct {
	value *GetTimeSeriesKst200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesKst200ResponseMetaIndicator) Get() *GetTimeSeriesKst200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesKst200ResponseMetaIndicator) Set(val *GetTimeSeriesKst200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKst200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKst200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKst200ResponseMetaIndicator(val *GetTimeSeriesKst200ResponseMetaIndicator) *NullableGetTimeSeriesKst200ResponseMetaIndicator {
	return &NullableGetTimeSeriesKst200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKst200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKst200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
