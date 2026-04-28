/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject13 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject13{}

// InlineObject13 struct for InlineObject13
type InlineObject13 struct {
	Meta   *InlineObject13Meta         `json:"meta,omitempty"`
	Values []InlineObject13ValuesInner `json:"values,omitempty"`
}

// NewInlineObject13 instantiates a new InlineObject13 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject13() *InlineObject13 {
	this := InlineObject13{}
	return &this
}

// NewInlineObject13WithDefaults instantiates a new InlineObject13 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject13WithDefaults() *InlineObject13 {
	this := InlineObject13{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineObject13) GetMeta() InlineObject13Meta {
	if o == nil || IsNil(o.Meta) {
		var ret InlineObject13Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13) GetMetaOk() (*InlineObject13Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineObject13) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineObject13Meta and assigns it to the Meta field.
func (o *InlineObject13) SetMeta(v InlineObject13Meta) {
	o.Meta = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *InlineObject13) GetValues() []InlineObject13ValuesInner {
	if o == nil || IsNil(o.Values) {
		var ret []InlineObject13ValuesInner
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13) GetValuesOk() ([]InlineObject13ValuesInner, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *InlineObject13) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []InlineObject13ValuesInner and assigns it to the Values field.
func (o *InlineObject13) SetValues(v []InlineObject13ValuesInner) {
	o.Values = v
}

func (o InlineObject13) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject13) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableInlineObject13 struct {
	value *InlineObject13
	isSet bool
}

func (v NullableInlineObject13) Get() *InlineObject13 {
	return v.value
}

func (v *NullableInlineObject13) Set(val *InlineObject13) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject13) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject13) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject13(val *InlineObject13) *NullableInlineObject13 {
	return &NullableInlineObject13{value: val, isSet: true}
}

func (v NullableInlineObject13) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject13) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
