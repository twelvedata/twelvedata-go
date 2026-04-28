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

// checks if the InlineObject10ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject10ValuesInner{}

// InlineObject10ValuesInner struct for InlineObject10ValuesInner
type InlineObject10ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ATAN value
	Atan string `json:"atan"`
}

type _InlineObject10ValuesInner InlineObject10ValuesInner

// NewInlineObject10ValuesInner instantiates a new InlineObject10ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10ValuesInner(datetime string, atan string) *InlineObject10ValuesInner {
	this := InlineObject10ValuesInner{}
	this.Datetime = datetime
	this.Atan = atan
	return &this
}

// NewInlineObject10ValuesInnerWithDefaults instantiates a new InlineObject10ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10ValuesInnerWithDefaults() *InlineObject10ValuesInner {
	this := InlineObject10ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject10ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject10ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject10ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAtan returns the Atan field value
func (o *InlineObject10ValuesInner) GetAtan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Atan
}

// GetAtanOk returns a tuple with the Atan field value
// and a boolean to check if the value has been set.
func (o *InlineObject10ValuesInner) GetAtanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Atan, true
}

// SetAtan sets field value
func (o *InlineObject10ValuesInner) SetAtan(v string) {
	o.Atan = v
}

func (o InlineObject10ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject10ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["atan"] = o.Atan
	return toSerialize, nil
}

func (o *InlineObject10ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"atan",
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

	varInlineObject10ValuesInner := _InlineObject10ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineObject10ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject10ValuesInner(varInlineObject10ValuesInner)

	return err
}

type NullableInlineObject10ValuesInner struct {
	value *InlineObject10ValuesInner
	isSet bool
}

func (v NullableInlineObject10ValuesInner) Get() *InlineObject10ValuesInner {
	return v.value
}

func (v *NullableInlineObject10ValuesInner) Set(val *InlineObject10ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10ValuesInner(val *InlineObject10ValuesInner) *NullableInlineObject10ValuesInner {
	return &NullableInlineObject10ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject10ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
