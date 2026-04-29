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

// checks if the InlineObject15ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject15ValuesInner{}

// InlineObject15ValuesInner struct for InlineObject15ValuesInner
type InlineObject15ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SINH value
	Sinh string `json:"sinh"`
}

type _InlineObject15ValuesInner InlineObject15ValuesInner

// NewInlineObject15ValuesInner instantiates a new InlineObject15ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject15ValuesInner(datetime string, sinh string) *InlineObject15ValuesInner {
	this := InlineObject15ValuesInner{}
	this.Datetime = datetime
	this.Sinh = sinh
	return &this
}

// NewInlineObject15ValuesInnerWithDefaults instantiates a new InlineObject15ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject15ValuesInnerWithDefaults() *InlineObject15ValuesInner {
	this := InlineObject15ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject15ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject15ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject15ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSinh returns the Sinh field value
func (o *InlineObject15ValuesInner) GetSinh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sinh
}

// GetSinhOk returns a tuple with the Sinh field value
// and a boolean to check if the value has been set.
func (o *InlineObject15ValuesInner) GetSinhOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sinh, true
}

// SetSinh sets field value
func (o *InlineObject15ValuesInner) SetSinh(v string) {
	o.Sinh = v
}

func (o InlineObject15ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject15ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sinh"] = o.Sinh
	return toSerialize, nil
}

func (o *InlineObject15ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sinh",
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

	varInlineObject15ValuesInner := _InlineObject15ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject15ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject15ValuesInner(varInlineObject15ValuesInner)

	return err
}

type NullableInlineObject15ValuesInner struct {
	value *InlineObject15ValuesInner
	isSet bool
}

func (v NullableInlineObject15ValuesInner) Get() *InlineObject15ValuesInner {
	return v.value
}

func (v *NullableInlineObject15ValuesInner) Set(val *InlineObject15ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject15ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject15ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject15ValuesInner(val *InlineObject15ValuesInner) *NullableInlineObject15ValuesInner {
	return &NullableInlineObject15ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject15ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject15ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
