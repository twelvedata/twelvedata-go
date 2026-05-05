// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject12ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject12ValuesInner{}

// InlineObject12ValuesInner struct for InlineObject12ValuesInner
type InlineObject12ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// COSH value
	Cosh string `json:"cosh"`
}

type _InlineObject12ValuesInner InlineObject12ValuesInner

// NewInlineObject12ValuesInner instantiates a new InlineObject12ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject12ValuesInner(datetime string, cosh string) *InlineObject12ValuesInner {
	this := InlineObject12ValuesInner{}
	this.Datetime = datetime
	this.Cosh = cosh
	return &this
}

// NewInlineObject12ValuesInnerWithDefaults instantiates a new InlineObject12ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject12ValuesInnerWithDefaults() *InlineObject12ValuesInner {
	this := InlineObject12ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject12ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject12ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject12ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCosh returns the Cosh field value
func (o *InlineObject12ValuesInner) GetCosh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cosh
}

// GetCoshOk returns a tuple with the Cosh field value
// and a boolean to check if the value has been set.
func (o *InlineObject12ValuesInner) GetCoshOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cosh, true
}

// SetCosh sets field value
func (o *InlineObject12ValuesInner) SetCosh(v string) {
	o.Cosh = v
}

func (o InlineObject12ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject12ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["cosh"] = o.Cosh
	return toSerialize, nil
}

func (o *InlineObject12ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"cosh",
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

	varInlineObject12ValuesInner := _InlineObject12ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject12ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject12ValuesInner(varInlineObject12ValuesInner)

	return err
}

type NullableInlineObject12ValuesInner struct {
	value *InlineObject12ValuesInner
	isSet bool
}

func (v NullableInlineObject12ValuesInner) Get() *InlineObject12ValuesInner {
	return v.value
}

func (v *NullableInlineObject12ValuesInner) Set(val *InlineObject12ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject12ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject12ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject12ValuesInner(val *InlineObject12ValuesInner) *NullableInlineObject12ValuesInner {
	return &NullableInlineObject12ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject12ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject12ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
