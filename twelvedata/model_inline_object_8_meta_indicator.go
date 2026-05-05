// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject8MetaIndicator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject8MetaIndicator{}

// InlineObject8MetaIndicator Technical indicator information
type InlineObject8MetaIndicator struct {
	// Name of the technical indicator
	Name string `json:"name"`
	// Price type on which technical indicator is calculated
	SeriesType string `json:"series_type"`
}

type _InlineObject8MetaIndicator InlineObject8MetaIndicator

// NewInlineObject8MetaIndicator instantiates a new InlineObject8MetaIndicator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8MetaIndicator(name string, seriesType string) *InlineObject8MetaIndicator {
	this := InlineObject8MetaIndicator{}
	this.Name = name
	this.SeriesType = seriesType
	return &this
}

// NewInlineObject8MetaIndicatorWithDefaults instantiates a new InlineObject8MetaIndicator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8MetaIndicatorWithDefaults() *InlineObject8MetaIndicator {
	this := InlineObject8MetaIndicator{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject8MetaIndicator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject8MetaIndicator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject8MetaIndicator) SetName(v string) {
	o.Name = v
}

// GetSeriesType returns the SeriesType field value
func (o *InlineObject8MetaIndicator) GetSeriesType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeriesType
}

// GetSeriesTypeOk returns a tuple with the SeriesType field value
// and a boolean to check if the value has been set.
func (o *InlineObject8MetaIndicator) GetSeriesTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesType, true
}

// SetSeriesType sets field value
func (o *InlineObject8MetaIndicator) SetSeriesType(v string) {
	o.SeriesType = v
}

func (o InlineObject8MetaIndicator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject8MetaIndicator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["series_type"] = o.SeriesType
	return toSerialize, nil
}

func (o *InlineObject8MetaIndicator) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject8MetaIndicator := _InlineObject8MetaIndicator{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject8MetaIndicator)

	if err != nil {
		return err
	}

	*o = InlineObject8MetaIndicator(varInlineObject8MetaIndicator)

	return err
}

type NullableInlineObject8MetaIndicator struct {
	value *InlineObject8MetaIndicator
	isSet bool
}

func (v NullableInlineObject8MetaIndicator) Get() *InlineObject8MetaIndicator {
	return v.value
}

func (v *NullableInlineObject8MetaIndicator) Set(val *InlineObject8MetaIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8MetaIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8MetaIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8MetaIndicator(val *InlineObject8MetaIndicator) *NullableInlineObject8MetaIndicator {
	return &NullableInlineObject8MetaIndicator{value: val, isSet: true}
}

func (v NullableInlineObject8MetaIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8MetaIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
