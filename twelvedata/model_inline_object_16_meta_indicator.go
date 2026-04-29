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

// checks if the InlineObject16MetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject16MetaIndicator{}

// InlineObject16MetaIndicator Technical indicator information
type InlineObject16MetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _InlineObject16MetaIndicator InlineObject16MetaIndicator

// NewInlineObject16MetaIndicator instantiates a new InlineObject16MetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16MetaIndicator(name string, seriesType string) *InlineObject16MetaIndicator {
	this := InlineObject16MetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewInlineObject16MetaIndicatorWithDefaults instantiates a new InlineObject16MetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16MetaIndicatorWithDefaults() *InlineObject16MetaIndicator {
	this := InlineObject16MetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject16MetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject16MetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject16MetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *InlineObject16MetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *InlineObject16MetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *InlineObject16MetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o InlineObject16MetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject16MetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *InlineObject16MetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject16MetaIndicator := _InlineObject16MetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject16MetaIndicator)

	if err != nil {
		return err
	}

	*o = InlineObject16MetaIndicator(varInlineObject16MetaIndicator)

	return err
}

type NullableInlineObject16MetaIndicator struct {
	value *InlineObject16MetaIndicator
	isSet bool
}

func (v NullableInlineObject16MetaIndicator) Get() *InlineObject16MetaIndicator {
	return v.value
}

func (v *NullableInlineObject16MetaIndicator) Set(val *InlineObject16MetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16MetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16MetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16MetaIndicator(val *InlineObject16MetaIndicator) *NullableInlineObject16MetaIndicator {
	return &NullableInlineObject16MetaIndicator{value: val, isSet: true}
}

func (v NullableInlineObject16MetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16MetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
