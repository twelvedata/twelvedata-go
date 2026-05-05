// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject12 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject12{}

// InlineObject12 struct for InlineObject12
type InlineObject12 struct {
	Meta InlineObject12Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject12ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject12 InlineObject12

// NewInlineObject12 instantiates a new InlineObject12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject12(meta InlineObject12Meta, values []InlineObject12ValuesInner, status string) *InlineObject12 {
	this := InlineObject12{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject12WithDefaults instantiates a new InlineObject12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject12WithDefaults() *InlineObject12 {
	this := InlineObject12{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject12) GetMeta() InlineObject12Meta {
	if o == nil {
		var ret InlineObject12Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetMetaOk() (*InlineObject12Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject12) SetMeta(v InlineObject12Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject12) GetValues() []InlineObject12ValuesInner {
	if o == nil {
		var ret []InlineObject12ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetValuesOk() ([]InlineObject12ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject12) SetValues(v []InlineObject12ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject12) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject12) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject12) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject12) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject12) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject12) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject12 := _InlineObject12{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject12)

	if err != nil {
		return err
	}

	*o = InlineObject12(varInlineObject12)

	return err
}

type NullableInlineObject12 struct {
	value *InlineObject12
	isSet bool
}

func (v NullableInlineObject12) Get() *InlineObject12 {
	return v.value
}

func (v *NullableInlineObject12) Set(val *InlineObject12) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject12) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject12(val *InlineObject12) *NullableInlineObject12 {
	return &NullableInlineObject12{value: val, isSet: true}
}

func (v NullableInlineObject12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
