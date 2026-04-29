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

// checks if the InlineObject14ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject14ValuesInner{}

// InlineObject14ValuesInner struct for InlineObject14ValuesInner
type InlineObject14ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SIN value
	Sin string `json:"sin"`
}

type _InlineObject14ValuesInner InlineObject14ValuesInner

// NewInlineObject14ValuesInner instantiates a new InlineObject14ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject14ValuesInner(datetime string, sin string) *InlineObject14ValuesInner {
	this := InlineObject14ValuesInner{}
	this.Datetime = datetime
	this.Sin = sin
	return &this
}

// NewInlineObject14ValuesInnerWithDefaults instantiates a new InlineObject14ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject14ValuesInnerWithDefaults() *InlineObject14ValuesInner {
	this := InlineObject14ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject14ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject14ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject14ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSin returns the Sin field value
func (o *InlineObject14ValuesInner) GetSin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sin
}

// GetSinOk returns a tuple with the Sin field value
// and a boolean to check if the value has been set.
func (o *InlineObject14ValuesInner) GetSinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sin, true
}

// SetSin sets field value
func (o *InlineObject14ValuesInner) SetSin(v string) {
	o.Sin = v
}

func (o InlineObject14ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject14ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sin"] = o.Sin
	return toSerialize, nil
}

func (o *InlineObject14ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sin",
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

	varInlineObject14ValuesInner := _InlineObject14ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject14ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject14ValuesInner(varInlineObject14ValuesInner)

	return err
}

type NullableInlineObject14ValuesInner struct {
	value *InlineObject14ValuesInner
	isSet bool
}

func (v NullableInlineObject14ValuesInner) Get() *InlineObject14ValuesInner {
	return v.value
}

func (v *NullableInlineObject14ValuesInner) Set(val *InlineObject14ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject14ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject14ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject14ValuesInner(val *InlineObject14ValuesInner) *NullableInlineObject14ValuesInner {
	return &NullableInlineObject14ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject14ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject14ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
