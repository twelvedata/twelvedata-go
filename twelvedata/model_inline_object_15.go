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

// checks if the InlineObject15 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject15{}

// InlineObject15 struct for InlineObject15
type InlineObject15 struct {
	Meta InlineObject15Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject15ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject15 InlineObject15

// NewInlineObject15 instantiates a new InlineObject15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject15(meta InlineObject15Meta, values []InlineObject15ValuesInner, status string) *InlineObject15 {
	this := InlineObject15{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject15WithDefaults instantiates a new InlineObject15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject15WithDefaults() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject15) GetMeta() InlineObject15Meta {
	if o == nil {
		var ret InlineObject15Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetMetaOk() (*InlineObject15Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject15) SetMeta(v InlineObject15Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject15) GetValues() []InlineObject15ValuesInner {
	if o == nil {
		var ret []InlineObject15ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetValuesOk() ([]InlineObject15ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject15) SetValues(v []InlineObject15ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject15) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject15) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject15) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject15) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject15) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject15 := _InlineObject15{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject15)

	if err != nil {
		return err
	}

	*o = InlineObject15(varInlineObject15)

	return err
}

type NullableInlineObject15 struct {
	value *InlineObject15
	isSet bool
}

func (v NullableInlineObject15) Get() *InlineObject15 {
	return v.value
}

func (v *NullableInlineObject15) Set(val *InlineObject15) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject15) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject15(val *InlineObject15) *NullableInlineObject15 {
	return &NullableInlineObject15{value: val, isSet: true}
}

func (v NullableInlineObject15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
