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

// checks if the InlineObject17ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject17ValuesInner{}

// InlineObject17ValuesInner struct for InlineObject17ValuesInner
type InlineObject17ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// TANH value
	Tanh string `json:"tanh"`
}

type _InlineObject17ValuesInner InlineObject17ValuesInner

// NewInlineObject17ValuesInner instantiates a new InlineObject17ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject17ValuesInner(datetime string, tanh string) *InlineObject17ValuesInner {
	this := InlineObject17ValuesInner{}
	this.Datetime = datetime
	this.Tanh = tanh
	return &this
}

// NewInlineObject17ValuesInnerWithDefaults instantiates a new InlineObject17ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject17ValuesInnerWithDefaults() *InlineObject17ValuesInner {
	this := InlineObject17ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject17ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject17ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject17ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTanh returns the Tanh field value
func (o *InlineObject17ValuesInner) GetTanh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tanh
}

// GetTanhOk returns a tuple with the Tanh field value
// and a boolean to check if the value has been set.
func (o *InlineObject17ValuesInner) GetTanhOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tanh, true
}

// SetTanh sets field value
func (o *InlineObject17ValuesInner) SetTanh(v string) {
	o.Tanh = v
}

func (o InlineObject17ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject17ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["tanh"] = o.Tanh
	return toSerialize, nil
}

func (o *InlineObject17ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"tanh",
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

	varInlineObject17ValuesInner := _InlineObject17ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject17ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject17ValuesInner(varInlineObject17ValuesInner)

	return err
}

type NullableInlineObject17ValuesInner struct {
	value *InlineObject17ValuesInner
	isSet bool
}

func (v NullableInlineObject17ValuesInner) Get() *InlineObject17ValuesInner {
	return v.value
}

func (v *NullableInlineObject17ValuesInner) Set(val *InlineObject17ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject17ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject17ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject17ValuesInner(val *InlineObject17ValuesInner) *NullableInlineObject17ValuesInner {
	return &NullableInlineObject17ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject17ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject17ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
