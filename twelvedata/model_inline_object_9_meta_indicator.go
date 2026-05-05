// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject9MetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject9MetaIndicator{}

// InlineObject9MetaIndicator Technical indicator information
type InlineObject9MetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _InlineObject9MetaIndicator InlineObject9MetaIndicator

// NewInlineObject9MetaIndicator instantiates a new InlineObject9MetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject9MetaIndicator(name string, seriesType string) *InlineObject9MetaIndicator {
	this := InlineObject9MetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewInlineObject9MetaIndicatorWithDefaults instantiates a new InlineObject9MetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject9MetaIndicatorWithDefaults() *InlineObject9MetaIndicator {
	this := InlineObject9MetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject9MetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject9MetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject9MetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *InlineObject9MetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *InlineObject9MetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *InlineObject9MetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o InlineObject9MetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject9MetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *InlineObject9MetaIndicator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"series_type",
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

	varInlineObject9MetaIndicator := _InlineObject9MetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject9MetaIndicator)

	if err != nil {
		return err
	}

	*o = InlineObject9MetaIndicator(varInlineObject9MetaIndicator)

	return err
}

type NullableInlineObject9MetaIndicator struct {
	value *InlineObject9MetaIndicator
	isSet bool
}

func (v NullableInlineObject9MetaIndicator) Get() *InlineObject9MetaIndicator {
	return v.value
}

func (v *NullableInlineObject9MetaIndicator) Set(val *InlineObject9MetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject9MetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject9MetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject9MetaIndicator(val *InlineObject9MetaIndicator) *NullableInlineObject9MetaIndicator {
	return &NullableInlineObject9MetaIndicator{value: val, isSet: true}
}

func (v NullableInlineObject9MetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject9MetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
