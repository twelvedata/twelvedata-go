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

// checks if the GetTimeSeriesWclPrice200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesWclPrice200ResponseMetaIndicator{}

// GetTimeSeriesWclPrice200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesWclPrice200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
}

type _GetTimeSeriesWclPrice200ResponseMetaIndicator GetTimeSeriesWclPrice200ResponseMetaIndicator

// NewGetTimeSeriesWclPrice200ResponseMetaIndicator instantiates a new GetTimeSeriesWclPrice200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesWclPrice200ResponseMetaIndicator(name string) *GetTimeSeriesWclPrice200ResponseMetaIndicator {
	this := GetTimeSeriesWclPrice200ResponseMetaIndicator{}
	this.Name = name
	return &this
}

// NewGetTimeSeriesWclPrice200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesWclPrice200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesWclPrice200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesWclPrice200ResponseMetaIndicator {
	this := GetTimeSeriesWclPrice200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesWclPrice200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWclPrice200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesWclPrice200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

func (o GetTimeSeriesWclPrice200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesWclPrice200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GetTimeSeriesWclPrice200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesWclPrice200ResponseMetaIndicator := _GetTimeSeriesWclPrice200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesWclPrice200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesWclPrice200ResponseMetaIndicator(varGetTimeSeriesWclPrice200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesWclPrice200ResponseMetaIndicator struct {
	value *GetTimeSeriesWclPrice200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesWclPrice200ResponseMetaIndicator) Get() *GetTimeSeriesWclPrice200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesWclPrice200ResponseMetaIndicator) Set(val *GetTimeSeriesWclPrice200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesWclPrice200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesWclPrice200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesWclPrice200ResponseMetaIndicator(val *GetTimeSeriesWclPrice200ResponseMetaIndicator) *NullableGetTimeSeriesWclPrice200ResponseMetaIndicator {
	return &NullableGetTimeSeriesWclPrice200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesWclPrice200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesWclPrice200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
