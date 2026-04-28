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

// checks if the GetTimeSeriesAvgPrice200ResponseMetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAvgPrice200ResponseMetaIndicator{}

// GetTimeSeriesAvgPrice200ResponseMetaIndicator Technical indicator information
type GetTimeSeriesAvgPrice200ResponseMetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
}

type _GetTimeSeriesAvgPrice200ResponseMetaIndicator GetTimeSeriesAvgPrice200ResponseMetaIndicator

// NewGetTimeSeriesAvgPrice200ResponseMetaIndicator instantiates a new GetTimeSeriesAvgPrice200ResponseMetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAvgPrice200ResponseMetaIndicator(name string) *GetTimeSeriesAvgPrice200ResponseMetaIndicator {
	this := GetTimeSeriesAvgPrice200ResponseMetaIndicator{}
	this.Name = name
	return &this
}

// NewGetTimeSeriesAvgPrice200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesAvgPrice200ResponseMetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAvgPrice200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesAvgPrice200ResponseMetaIndicator {
	this := GetTimeSeriesAvgPrice200ResponseMetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *GetTimeSeriesAvgPrice200ResponseMetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvgPrice200ResponseMetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetTimeSeriesAvgPrice200ResponseMetaIndicator) SetName(v string) {
	o.Name = v
}

func (o GetTimeSeriesAvgPrice200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAvgPrice200ResponseMetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *GetTimeSeriesAvgPrice200ResponseMetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varGetTimeSeriesAvgPrice200ResponseMetaIndicator := _GetTimeSeriesAvgPrice200ResponseMetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAvgPrice200ResponseMetaIndicator)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAvgPrice200ResponseMetaIndicator(varGetTimeSeriesAvgPrice200ResponseMetaIndicator)

	return err
}

type NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator struct {
	value *GetTimeSeriesAvgPrice200ResponseMetaIndicator
	isSet bool
}

func (v NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator) Get() *GetTimeSeriesAvgPrice200ResponseMetaIndicator {
	return v.value
}

func (v *NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator) Set(val *GetTimeSeriesAvgPrice200ResponseMetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAvgPrice200ResponseMetaIndicator(val *GetTimeSeriesAvgPrice200ResponseMetaIndicator) *NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator {
	return &NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAvgPrice200ResponseMetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
