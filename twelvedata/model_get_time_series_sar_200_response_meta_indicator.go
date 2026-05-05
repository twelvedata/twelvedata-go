// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesSar200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSar200ResponseMetaIndicator{}

// GetTimeSeriesSar200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesSar200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// The acceleration factor used in the indicator calculation
	Acceleration float64 `json:"acceleration"`
	// The maximum value considered for the indicator calculation
	Maximum float64 `json:"maximum"`
}

type _GetTimeSeriesSar200ResponseMetaIndicator GetTimeSeriesSar200ResponseMetaIndicator

// NewGetTimeSeriesSar200ResponseMetaIndicator instantiates a new GetTimeSeriesSar200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSar200ResponseMetaIndicator(name string, acceleration float64, maximum float64) *GetTimeSeriesSar200ResponseMetaIndicator {
	this := GetTimeSeriesSar200ResponseMetaIndicator{}
	this.Name = name
	this.Acceleration = acceleration
	this.Maximum = maximum
	return &this
}

// NewGetTimeSeriesSar200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesSar200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSar200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesSar200ResponseMetaIndicator {
	this := GetTimeSeriesSar200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesSar200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesSar200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

// GetAcceleration returns the Acceleration field value
func (o *GetTimeSeriesSar200ResponseMetaIndicator) GetAcceleration() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Acceleration
}

// GetAccelerationOk returns a tuple with the Acceleration field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200ResponseMetaIndicator) GetAccelerationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Acceleration, true
}

// SetAcceleration sets field value
func (o *GetTimeSeriesSar200ResponseMetaIndicator) SetAcceleration(v float64) {
	o.Acceleration = v
}

// GetMaximum returns the Maximum field value
func (o *GetTimeSeriesSar200ResponseMetaIndicator) GetMaximum() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200ResponseMetaIndicator) GetMaximumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maximum, true
}

// SetMaximum sets field value
func (o *GetTimeSeriesSar200ResponseMetaIndicator) SetMaximum(v float64) {
	o.Maximum = v
}

func (o GetTimeSeriesSar200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSar200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["acceleration"] = o.Acceleration
	toSerialize["maximum"] = o.Maximum
	return toSerialize, nil
}

func (o *GetTimeSeriesSar200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"acceleration",
		"maximum",
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

	varGetTimeSeriesSar200ResponseMetaIndicator := _GetTimeSeriesSar200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSar200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSar200ResponseMetaIndicator(varGetTimeSeriesSar200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesSar200ResponseMetaIndicator struct {
	value *GetTimeSeriesSar200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesSar200ResponseMetaIndicator) Get() *GetTimeSeriesSar200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesSar200ResponseMetaIndicator) Set(val *GetTimeSeriesSar200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSar200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSar200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSar200ResponseMetaIndicator(val *GetTimeSeriesSar200ResponseMetaIndicator) *NullableGetTimeSeriesSar200ResponseMetaIndicator {
	return &NullableGetTimeSeriesSar200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSar200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSar200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
