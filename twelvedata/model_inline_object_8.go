// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject8 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject8{}

// InlineObject8 struct for InlineObject8
type InlineObject8 struct {
	Meta InlineObject8Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject8ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject8 InlineObject8

// NewInlineObject8 instantiates a new InlineObject8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8(meta InlineObject8Meta, values []InlineObject8ValuesInner, status string) *InlineObject8 {
	this := InlineObject8{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject8WithDefaults instantiates a new InlineObject8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8WithDefaults() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject8) GetMeta() InlineObject8Meta {
	if o == nil {
		var ret InlineObject8Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetMetaOk() (*InlineObject8Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject8) SetMeta(v InlineObject8Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject8) GetValues() []InlineObject8ValuesInner {
	if o == nil {
		var ret []InlineObject8ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetValuesOk() ([]InlineObject8ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject8) SetValues(v []InlineObject8ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject8) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject8) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject8) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject8) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject8) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject8 := _InlineObject8{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject8)

	if err != nil {
		return err
	}

	*o = InlineObject8(varInlineObject8)

	return err
}

type NullableInlineObject8 struct {
	value *InlineObject8
	isSet bool
}

func (v NullableInlineObject8) Get() *InlineObject8 {
	return v.value
}

func (v *NullableInlineObject8) Set(val *InlineObject8) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8(val *InlineObject8) *NullableInlineObject8 {
	return &NullableInlineObject8{value: val, isSet: true}
}

func (v NullableInlineObject8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
