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

// checks if the GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator{}

// GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
}

type _GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator

// NewGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator instantiates a new GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator(name string) *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator {
	this := GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator{}
	this.Name = name
	return &this
}

// NewGetTimeSeriesHeikinashiCandles200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHeikinashiCandles200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator {
	this := GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

func (o GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator := _GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator(varGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator struct {
	value *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) Get() *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) Set(val *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator(val *GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) *NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator {
	return &NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHeikinashiCandles200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
