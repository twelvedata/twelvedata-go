// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject17 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject17{}

// InlineObject17 struct for InlineObject17
type InlineObject17 struct {
	Meta InlineObject17Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject17ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject17 InlineObject17

// NewInlineObject17 instantiates a new InlineObject17 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject17(meta InlineObject17Meta, values []InlineObject17ValuesInner, status string) *InlineObject17 {
	this := InlineObject17{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject17WithDefaults instantiates a new InlineObject17 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject17WithDefaults() *InlineObject17 {
	this := InlineObject17{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject17) GetMeta() InlineObject17Meta {
	if o == nil {
		var ret InlineObject17Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject17) GetMetaOk() (*InlineObject17Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject17) SetMeta(v InlineObject17Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject17) GetValues() []InlineObject17ValuesInner {
	if o == nil {
		var ret []InlineObject17ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject17) GetValuesOk() ([]InlineObject17ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject17) SetValues(v []InlineObject17ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject17) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject17) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject17) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject17) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject17) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject17) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject17 := _InlineObject17{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject17)

	if err != nil {
		return err
	}

	*o = InlineObject17(varInlineObject17)

	return err
}

type NullableInlineObject17 struct {
	value *InlineObject17
	isSet bool
}

func (v NullableInlineObject17) Get() *InlineObject17 {
	return v.value
}

func (v *NullableInlineObject17) Set(val *InlineObject17) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject17) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject17) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject17(val *InlineObject17) *NullableInlineObject17 {
	return &NullableInlineObject17{value: val, isSet: true}
}

func (v NullableInlineObject17) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject17) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
