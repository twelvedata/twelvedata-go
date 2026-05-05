// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesHlc3200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHlc3200ResponseMetaIndicator{}

// GetTimeSeriesHlc3200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesHlc3200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
}

type _GetTimeSeriesHlc3200ResponseMetaIndicator GetTimeSeriesHlc3200ResponseMetaIndicator

// NewGetTimeSeriesHlc3200ResponseMetaIndicator instantiates a new GetTimeSeriesHlc3200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHlc3200ResponseMetaIndicator(name string) *GetTimeSeriesHlc3200ResponseMetaIndicator {
	this := GetTimeSeriesHlc3200ResponseMetaIndicator{}
	this.Name = name
	return &this
}

// NewGetTimeSeriesHlc3200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesHlc3200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHlc3200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesHlc3200ResponseMetaIndicator {
	this := GetTimeSeriesHlc3200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesHlc3200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHlc3200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesHlc3200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

func (o GetTimeSeriesHlc3200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHlc3200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GetTimeSeriesHlc3200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesHlc3200ResponseMetaIndicator := _GetTimeSeriesHlc3200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesHlc3200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHlc3200ResponseMetaIndicator(varGetTimeSeriesHlc3200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesHlc3200ResponseMetaIndicator struct {
	value *GetTimeSeriesHlc3200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesHlc3200ResponseMetaIndicator) Get() *GetTimeSeriesHlc3200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesHlc3200ResponseMetaIndicator) Set(val *GetTimeSeriesHlc3200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHlc3200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHlc3200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHlc3200ResponseMetaIndicator(val *GetTimeSeriesHlc3200ResponseMetaIndicator) *NullableGetTimeSeriesHlc3200ResponseMetaIndicator {
	return &NullableGetTimeSeriesHlc3200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHlc3200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHlc3200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
