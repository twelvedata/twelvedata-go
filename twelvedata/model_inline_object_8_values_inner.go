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

// checks if the InlineObject8ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject8ValuesInner{}

// InlineObject8ValuesInner struct for InlineObject8ValuesInner
type InlineObject8ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ACOS value
	Acos string `json:"acos"`
}

type _InlineObject8ValuesInner InlineObject8ValuesInner

// NewInlineObject8ValuesInner instantiates a new InlineObject8ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8ValuesInner(datetime string, acos string) *InlineObject8ValuesInner {
	this := InlineObject8ValuesInner{}
	this.Datetime = datetime
	this.Acos = acos
	return &this
}

// NewInlineObject8ValuesInnerWithDefaults instantiates a new InlineObject8ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8ValuesInnerWithDefaults() *InlineObject8ValuesInner {
	this := InlineObject8ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject8ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject8ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject8ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAcos returns the Acos field value
func (o *InlineObject8ValuesInner) GetAcos() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Acos
}

// GetAcosOk returns a tuple with the Acos field value
// and a boolean to check if the value has been set.
func (o *InlineObject8ValuesInner) GetAcosOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Acos, true
}

// SetAcos sets field value
func (o *InlineObject8ValuesInner) SetAcos(v string) {
	o.Acos = v
}

func (o InlineObject8ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject8ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["acos"] = o.Acos
	return toSerialize, nil
}

func (o *InlineObject8ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"acos",
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

	varInlineObject8ValuesInner := _InlineObject8ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineObject8ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject8ValuesInner(varInlineObject8ValuesInner)

	return err
}

type NullableInlineObject8ValuesInner struct {
	value *InlineObject8ValuesInner
	isSet bool
}

func (v NullableInlineObject8ValuesInner) Get() *InlineObject8ValuesInner {
	return v.value
}

func (v *NullableInlineObject8ValuesInner) Set(val *InlineObject8ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8ValuesInner(val *InlineObject8ValuesInner) *NullableInlineObject8ValuesInner {
	return &NullableInlineObject8ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject8ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
