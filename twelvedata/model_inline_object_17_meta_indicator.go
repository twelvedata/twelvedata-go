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

// checks if the InlineObject17MetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject17MetaIndicator{}

// InlineObject17MetaIndicator Technical indicator information
type InlineObject17MetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _InlineObject17MetaIndicator InlineObject17MetaIndicator

// NewInlineObject17MetaIndicator instantiates a new InlineObject17MetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject17MetaIndicator(name string, seriesType string) *InlineObject17MetaIndicator {
	this := InlineObject17MetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewInlineObject17MetaIndicatorWithDefaults instantiates a new InlineObject17MetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject17MetaIndicatorWithDefaults() *InlineObject17MetaIndicator {
	this := InlineObject17MetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject17MetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject17MetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject17MetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *InlineObject17MetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *InlineObject17MetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *InlineObject17MetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o InlineObject17MetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject17MetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *InlineObject17MetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject17MetaIndicator := _InlineObject17MetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject17MetaIndicator)

	if err != nil {
		return err
	}

	*o = InlineObject17MetaIndicator(varInlineObject17MetaIndicator)

	return err
}

type NullableInlineObject17MetaIndicator struct {
	value *InlineObject17MetaIndicator
	isSet bool
}

func (v NullableInlineObject17MetaIndicator) Get() *InlineObject17MetaIndicator {
	return v.value
}

func (v *NullableInlineObject17MetaIndicator) Set(val *InlineObject17MetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject17MetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject17MetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject17MetaIndicator(val *InlineObject17MetaIndicator) *NullableInlineObject17MetaIndicator {
	return &NullableInlineObject17MetaIndicator{value: val, isSet: true}
}

func (v NullableInlineObject17MetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject17MetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
