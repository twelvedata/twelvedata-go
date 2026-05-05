// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesTRange200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTRange200ResponseMetaIndicator{}

// GetTimeSeriesTRange200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesTRange200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
}

type _GetTimeSeriesTRange200ResponseMetaIndicator GetTimeSeriesTRange200ResponseMetaIndicator

// NewGetTimeSeriesTRange200ResponseMetaIndicator instantiates a new GetTimeSeriesTRange200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTRange200ResponseMetaIndicator(name string) *GetTimeSeriesTRange200ResponseMetaIndicator {
	this := GetTimeSeriesTRange200ResponseMetaIndicator{}
	this.Name = name
	return &this
}

// NewGetTimeSeriesTRange200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesTRange200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTRange200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesTRange200ResponseMetaIndicator {
	this := GetTimeSeriesTRange200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesTRange200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTRange200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesTRange200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

func (o GetTimeSeriesTRange200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTRange200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GetTimeSeriesTRange200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesTRange200ResponseMetaIndicator := _GetTimeSeriesTRange200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTRange200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTRange200ResponseMetaIndicator(varGetTimeSeriesTRange200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesTRange200ResponseMetaIndicator struct {
	value *GetTimeSeriesTRange200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesTRange200ResponseMetaIndicator) Get() *GetTimeSeriesTRange200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesTRange200ResponseMetaIndicator) Set(val *GetTimeSeriesTRange200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTRange200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTRange200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTRange200ResponseMetaIndicator(val *GetTimeSeriesTRange200ResponseMetaIndicator) *NullableGetTimeSeriesTRange200ResponseMetaIndicator {
	return &NullableGetTimeSeriesTRange200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTRange200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTRange200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
