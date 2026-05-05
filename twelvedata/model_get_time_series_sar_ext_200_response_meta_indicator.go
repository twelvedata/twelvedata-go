// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesSarExt200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSarExt200ResponseMetaIndicator{}

// GetTimeSeriesSarExt200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesSarExt200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The initial value for the indicator calculation
	StartValue float64 `json:"start_value"`
	// The adjustment applied when the indicator's direction changes
	OffsetOnReverse float64 `json:"offset_on_reverse"`
	// The maximum acceleration value for long positions
	AccelerationLimitLong float64 `json:"acceleration_limit_long"`
	// The acceleration value for long positions
	AccelerationLong float64 `json:"acceleration_long"`
	// The highest allowed acceleration for long positions
	AccelerationMaxLong float64 `json:"acceleration_max_long"`
	// The maximum acceleration value for short positions
	AccelerationLimitShort float64 `json:"acceleration_limit_short"`
	// The acceleration value for short positions
	AccelerationShort float64 `json:"acceleration_short"`
	// The highest allowed acceleration for short positions
	AccelerationMaxShort float64 `json:"acceleration_max_short"`
}

type _GetTimeSeriesSarExt200ResponseMetaIndicator GetTimeSeriesSarExt200ResponseMetaIndicator

// NewGetTimeSeriesSarExt200ResponseMetaIndicator instantiates a new GetTimeSeriesSarExt200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSarExt200ResponseMetaIndicator(name string, startValue float64, offsetOnReverse float64, accelerationLimitLong float64, accelerationLong float64, accelerationMaxLong float64, accelerationLimitShort float64, accelerationShort float64, accelerationMaxShort float64) *GetTimeSeriesSarExt200ResponseMetaIndicator {
	this := GetTimeSeriesSarExt200ResponseMetaIndicator{}
	this.Name = name
	this.StartValue = startValue
	this.OffsetOnReverse = offsetOnReverse
	this.AccelerationLimitLong = accelerationLimitLong
	this.AccelerationLong = accelerationLong
	this.AccelerationMaxLong = accelerationMaxLong
	this.AccelerationLimitShort = accelerationLimitShort
	this.AccelerationShort = accelerationShort
	this.AccelerationMaxShort = accelerationMaxShort
	return &this
}

// NewGetTimeSeriesSarExt200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesSarExt200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSarExt200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesSarExt200ResponseMetaIndicator {
	this := GetTimeSeriesSarExt200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetStartValue returns the StartValue field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetStartValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.StartValue
}

// GetStartValueOk returns a tuple with the StartValue field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetStartValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartValue, true
}

// SetStartValue sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetStartValue(v float64) {
	o.StartValue = v
}

// GetOffsetOnReverse returns the OffsetOnReverse field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetOffsetOnReverse() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.OffsetOnReverse
}

// GetOffsetOnReverseOk returns a tuple with the OffsetOnReverse field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetOffsetOnReverseOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OffsetOnReverse, true
}

// SetOffsetOnReverse sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetOffsetOnReverse(v float64) {
	o.OffsetOnReverse = v
}

// GetAccelerationLimitLong returns the AccelerationLimitLong field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitLong() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AccelerationLimitLong
}

// GetAccelerationLimitLongOk returns a tuple with the AccelerationLimitLong field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitLongOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccelerationLimitLong, true
}

// SetAccelerationLimitLong sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationLimitLong(v float64) {
	o.AccelerationLimitLong = v
}

// GetAccelerationLong returns the AccelerationLong field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLong() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AccelerationLong
}

// GetAccelerationLongOk returns a tuple with the AccelerationLong field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLongOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccelerationLong, true
}

// SetAccelerationLong sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationLong(v float64) {
	o.AccelerationLong = v
}

// GetAccelerationMaxLong returns the AccelerationMaxLong field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxLong() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AccelerationMaxLong
}

// GetAccelerationMaxLongOk returns a tuple with the AccelerationMaxLong field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxLongOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccelerationMaxLong, true
}

// SetAccelerationMaxLong sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationMaxLong(v float64) {
	o.AccelerationMaxLong = v
}

// GetAccelerationLimitShort returns the AccelerationLimitShort field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitShort() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AccelerationLimitShort
}

// GetAccelerationLimitShortOk returns a tuple with the AccelerationLimitShort field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationLimitShortOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccelerationLimitShort, true
}

// SetAccelerationLimitShort sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationLimitShort(v float64) {
	o.AccelerationLimitShort = v
}

// GetAccelerationShort returns the AccelerationShort field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationShort() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AccelerationShort
}

// GetAccelerationShortOk returns a tuple with the AccelerationShort field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationShortOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccelerationShort, true
}

// SetAccelerationShort sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationShort(v float64) {
	o.AccelerationShort = v
}

// GetAccelerationMaxShort returns the AccelerationMaxShort field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxShort() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AccelerationMaxShort
}

// GetAccelerationMaxShortOk returns a tuple with the AccelerationMaxShort field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) GetAccelerationMaxShortOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccelerationMaxShort, true
}

// SetAccelerationMaxShort sets field value
func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) SetAccelerationMaxShort(v float64) {
	o.AccelerationMaxShort = v
}

func (o GetTimeSeriesSarExt200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSarExt200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["start_value"] = o.StartValue
	toSerialize["offset_on_reverse"] = o.OffsetOnReverse
	toSerialize["acceleration_limit_long"] = o.AccelerationLimitLong
	toSerialize["acceleration_long"] = o.AccelerationLong
	toSerialize["acceleration_max_long"] = o.AccelerationMaxLong
	toSerialize["acceleration_limit_short"] = o.AccelerationLimitShort
	toSerialize["acceleration_short"] = o.AccelerationShort
	toSerialize["acceleration_max_short"] = o.AccelerationMaxShort
	return toSerialize, nil
}

func (o *GetTimeSeriesSarExt200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"start_value",
		"offset_on_reverse",
		"acceleration_limit_long",
		"acceleration_long",
		"acceleration_max_long",
		"acceleration_limit_short",
		"acceleration_short",
		"acceleration_max_short",
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

	varGetTimeSeriesSarExt200ResponseMetaIndicator := _GetTimeSeriesSarExt200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSarExt200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSarExt200ResponseMetaIndicator(varGetTimeSeriesSarExt200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesSarExt200ResponseMetaIndicator struct {
	value *GetTimeSeriesSarExt200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesSarExt200ResponseMetaIndicator) Get() *GetTimeSeriesSarExt200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesSarExt200ResponseMetaIndicator) Set(val *GetTimeSeriesSarExt200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSarExt200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSarExt200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSarExt200ResponseMetaIndicator(val *GetTimeSeriesSarExt200ResponseMetaIndicator) *NullableGetTimeSeriesSarExt200ResponseMetaIndicator {
	return &NullableGetTimeSeriesSarExt200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSarExt200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSarExt200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
