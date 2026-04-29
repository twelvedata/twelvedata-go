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

// checks if the GetTimeSeriesSuperTrend200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSuperTrend200ResponseMetaIndicator{}

// GetTimeSeriesSuperTrend200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesSuperTrend200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The period used for calculation in the indicator
	Period int64 `json:"period"`
	// The factor used to adjust the indicator's sensitivity
	Multiplier int64 `json:"multiplier"`
}

type _GetTimeSeriesSuperTrend200ResponseMetaIndicator GetTimeSeriesSuperTrend200ResponseMetaIndicator

// NewGetTimeSeriesSuperTrend200ResponseMetaIndicator instantiates a new GetTimeSeriesSuperTrend200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSuperTrend200ResponseMetaIndicator(name string, period int64, multiplier int64) *GetTimeSeriesSuperTrend200ResponseMetaIndicator {
	this := GetTimeSeriesSuperTrend200ResponseMetaIndicator{}
	this.Name = name
	this.Period = period
	this.Multiplier = multiplier
	return &this
}

// NewGetTimeSeriesSuperTrend200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesSuperTrend200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSuperTrend200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesSuperTrend200ResponseMetaIndicator {
	this := GetTimeSeriesSuperTrend200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetPeriod returns the Period field value
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) GetPeriod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) GetPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) SetPeriod(v int64) {
	o.Period = v
}

// GetMultiplier returns the Multiplier field value
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) GetMultiplier() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Multiplier
}

// GetMultiplierOk returns a tuple with the Multiplier field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) GetMultiplierOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Multiplier, true
}

// SetMultiplier sets field value
func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) SetMultiplier(v int64) {
	o.Multiplier = v
}

func (o GetTimeSeriesSuperTrend200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSuperTrend200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["period"] = o.Period
	toSerialize["multiplier"] = o.Multiplier
	return toSerialize, nil
}

func (o *GetTimeSeriesSuperTrend200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"period",
		"multiplier",
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

	varGetTimeSeriesSuperTrend200ResponseMetaIndicator := _GetTimeSeriesSuperTrend200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSuperTrend200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSuperTrend200ResponseMetaIndicator(varGetTimeSeriesSuperTrend200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator struct {
	value *GetTimeSeriesSuperTrend200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator) Get() *GetTimeSeriesSuperTrend200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator) Set(val *GetTimeSeriesSuperTrend200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSuperTrend200ResponseMetaIndicator(val *GetTimeSeriesSuperTrend200ResponseMetaIndicator) *NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator {
	return &NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSuperTrend200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
