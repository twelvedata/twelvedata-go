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

// checks if the InlineObject16 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject16{}

// InlineObject16 struct for InlineObject16
type InlineObject16 struct {
	Meta InlineObject16Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject16ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject16 InlineObject16

// NewInlineObject16 instantiates a new InlineObject16 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject16(meta InlineObject16Meta, values []InlineObject16ValuesInner, status string) *InlineObject16 {
	this := InlineObject16{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject16WithDefaults instantiates a new InlineObject16 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject16WithDefaults() *InlineObject16 {
	this := InlineObject16{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject16) GetMeta() InlineObject16Meta {
	if o == nil {
		var ret InlineObject16Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetMetaOk() (*InlineObject16Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject16) SetMeta(v InlineObject16Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject16) GetValues() []InlineObject16ValuesInner {
	if o == nil {
		var ret []InlineObject16ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetValuesOk() ([]InlineObject16ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject16) SetValues(v []InlineObject16ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject16) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject16) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject16) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject16) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject16) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject16) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject16 := _InlineObject16{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineObject16)

	if err != nil {
		return err
	}

	*o = InlineObject16(varInlineObject16)

	return err
}

type NullableInlineObject16 struct {
	value *InlineObject16
	isSet bool
}

func (v NullableInlineObject16) Get() *InlineObject16 {
	return v.value
}

func (v *NullableInlineObject16) Set(val *InlineObject16) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject16) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject16) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject16(val *InlineObject16) *NullableInlineObject16 {
	return &NullableInlineObject16{value: val, isSet: true}
}

func (v NullableInlineObject16) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject16) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
