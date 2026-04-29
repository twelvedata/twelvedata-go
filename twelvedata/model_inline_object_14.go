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

// checks if the InlineObject14 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject14{}

// InlineObject14 struct for InlineObject14
type InlineObject14 struct {
	Meta InlineObject14Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject14ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject14 InlineObject14

// NewInlineObject14 instantiates a new InlineObject14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject14(meta InlineObject14Meta, values []InlineObject14ValuesInner, status string) *InlineObject14 {
	this := InlineObject14{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject14WithDefaults instantiates a new InlineObject14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject14WithDefaults() *InlineObject14 {
	this := InlineObject14{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject14) GetMeta() InlineObject14Meta {
	if o == nil {
		var ret InlineObject14Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetMetaOk() (*InlineObject14Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject14) SetMeta(v InlineObject14Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject14) GetValues() []InlineObject14ValuesInner {
	if o == nil {
		var ret []InlineObject14ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetValuesOk() ([]InlineObject14ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject14) SetValues(v []InlineObject14ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject14) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject14) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject14) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject14) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject14) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject14) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject14 := _InlineObject14{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject14)

	if err != nil {
		return err
	}

	*o = InlineObject14(varInlineObject14)

	return err
}

type NullableInlineObject14 struct {
	value *InlineObject14
	isSet bool
}

func (v NullableInlineObject14) Get() *InlineObject14 {
	return v.value
}

func (v *NullableInlineObject14) Set(val *InlineObject14) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject14) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject14(val *InlineObject14) *NullableInlineObject14 {
	return &NullableInlineObject14{value: val, isSet: true}
}

func (v NullableInlineObject14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
