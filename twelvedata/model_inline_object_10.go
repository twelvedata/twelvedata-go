// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject10 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject10{}

// InlineObject10 struct for InlineObject10
type InlineObject10 struct {
	Meta InlineObject10Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject10ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject10 InlineObject10

// NewInlineObject10 instantiates a new InlineObject10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject10(meta InlineObject10Meta, values []InlineObject10ValuesInner, status string) *InlineObject10 {
	this := InlineObject10{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject10WithDefaults instantiates a new InlineObject10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject10WithDefaults() *InlineObject10 {
	this := InlineObject10{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject10) GetMeta() InlineObject10Meta {
	if o == nil {
		var ret InlineObject10Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetMetaOk() (*InlineObject10Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject10) SetMeta(v InlineObject10Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject10) GetValues() []InlineObject10ValuesInner {
	if o == nil {
		var ret []InlineObject10ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetValuesOk() ([]InlineObject10ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject10) SetValues(v []InlineObject10ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject10) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject10) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject10) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject10) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject10) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject10) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meta",
		"values",
		"status",
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

	varInlineObject10 := _InlineObject10{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject10)

	if err != nil {
		return err
	}

	*o = InlineObject10(varInlineObject10)

	return err
}

type NullableInlineObject10 struct {
	value *InlineObject10
	isSet bool
}

func (v NullableInlineObject10) Get() *InlineObject10 {
	return v.value
}

func (v *NullableInlineObject10) Set(val *InlineObject10) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject10) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject10(val *InlineObject10) *NullableInlineObject10 {
	return &NullableInlineObject10{value: val, isSet: true}
}

func (v NullableInlineObject10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
