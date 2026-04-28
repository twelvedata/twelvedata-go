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

// checks if the InlineObject11 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject11{}

// InlineObject11 struct for InlineObject11
type InlineObject11 struct {
	Meta InlineObject11Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject11ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject11 InlineObject11

// NewInlineObject11 instantiates a new InlineObject11 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject11(meta InlineObject11Meta, values []InlineObject11ValuesInner, status string) *InlineObject11 {
	this := InlineObject11{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject11WithDefaults instantiates a new InlineObject11 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject11WithDefaults() *InlineObject11 {
	this := InlineObject11{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject11) GetMeta() InlineObject11Meta {
	if o == nil {
		var ret InlineObject11Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject11) GetMetaOk() (*InlineObject11Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject11) SetMeta(v InlineObject11Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject11) GetValues() []InlineObject11ValuesInner {
	if o == nil {
		var ret []InlineObject11ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject11) GetValuesOk() ([]InlineObject11ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject11) SetValues(v []InlineObject11ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject11) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject11) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject11) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject11) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject11) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject11) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject11 := _InlineObject11{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineObject11)

	if err != nil {
		return err
	}

	*o = InlineObject11(varInlineObject11)

	return err
}

type NullableInlineObject11 struct {
	value *InlineObject11
	isSet bool
}

func (v NullableInlineObject11) Get() *InlineObject11 {
	return v.value
}

func (v *NullableInlineObject11) Set(val *InlineObject11) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject11) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject11) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject11(val *InlineObject11) *NullableInlineObject11 {
	return &NullableInlineObject11{value: val, isSet: true}
}

func (v NullableInlineObject11) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject11) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
