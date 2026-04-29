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

// checks if the InlineObject16ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject16ValuesInner{}

// InlineObject16ValuesInner struct for InlineObject16ValuesInner
type InlineObject16ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// TAN value
	Tan string `json:"tan"`
}

type _InlineObject16ValuesInner InlineObject16ValuesInner

// NewInlineObject16ValuesInner instantiates a new InlineObject16ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16ValuesInner(datetime string, tan string) *InlineObject16ValuesInner {
	this := InlineObject16ValuesInner{}
	this.Datetime = datetime
	this.Tan = tan
	return &this
}

// NewInlineObject16ValuesInnerWithDefaults instantiates a new InlineObject16ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16ValuesInnerWithDefaults() *InlineObject16ValuesInner {
	this := InlineObject16ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject16ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject16ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject16ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTan returns the Tan field value
func (o *InlineObject16ValuesInner) GetTan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tan
}

// GetTanOk returns a tuple with the Tan field value
// and a boolean to check if the value has been set.
func (o *InlineObject16ValuesInner) GetTanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tan, true
}

// SetTan sets field value
func (o *InlineObject16ValuesInner) SetTan(v string) {
	o.Tan = v
}

func (o InlineObject16ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject16ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["tan"] = o.Tan
	return toSerialize, nil
}

func (o *InlineObject16ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"tan",
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

	varInlineObject16ValuesInner := _InlineObject16ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject16ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject16ValuesInner(varInlineObject16ValuesInner)

	return err
}

type NullableInlineObject16ValuesInner struct {
	value *InlineObject16ValuesInner
	isSet bool
}

func (v NullableInlineObject16ValuesInner) Get() *InlineObject16ValuesInner {
	return v.value
}

func (v *NullableInlineObject16ValuesInner) Set(val *InlineObject16ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16ValuesInner(val *InlineObject16ValuesInner) *NullableInlineObject16ValuesInner {
	return &NullableInlineObject16ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject16ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
