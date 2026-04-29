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

// checks if the InlineObject9 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject9{}

// InlineObject9 struct for InlineObject9
type InlineObject9 struct {
	Meta InlineObject9Meta `json:"meta"`
	// Array of time series data points
	Values []InlineObject9ValuesInner `json:"values"`
	// Response status
	Status string `json:"status"`
}

type _InlineObject9 InlineObject9

// NewInlineObject9 instantiates a new InlineObject9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject9(meta InlineObject9Meta, values []InlineObject9ValuesInner, status string) *InlineObject9 {
	this := InlineObject9{}
	this.Meta = meta
	this.Values = values
	this.Status = status
	return &this
}

// NewInlineObject9WithDefaults instantiates a new InlineObject9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject9WithDefaults() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// GetMeta returns the Meta field value
func (o *InlineObject9) GetMeta() InlineObject9Meta {
	if o == nil {
		var ret InlineObject9Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetMetaOk() (*InlineObject9Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineObject9) SetMeta(v InlineObject9Meta) {
	o.Meta = v
}

// GetValues returns the Values field value
func (o *InlineObject9) GetValues() []InlineObject9ValuesInner {
	if o == nil {
		var ret []InlineObject9ValuesInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetValuesOk() ([]InlineObject9ValuesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *InlineObject9) SetValues(v []InlineObject9ValuesInner) {
	o.Values = v
}

// GetStatus returns the Status field value
func (o *InlineObject9) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineObject9) SetStatus(v string) {
	o.Status = v
}

func (o InlineObject9) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject9) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["values"] = o.Values
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *InlineObject9) UnmarshalJSON(data []byte) (err error) {
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

	varInlineObject9 := _InlineObject9{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject9)

	if err != nil {
		return err
	}

	*o = InlineObject9(varInlineObject9)

	return err
}

type NullableInlineObject9 struct {
	value *InlineObject9
	isSet bool
}

func (v NullableInlineObject9) Get() *InlineObject9 {
	return v.value
}

func (v *NullableInlineObject9) Set(val *InlineObject9) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject9) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject9(val *InlineObject9) *NullableInlineObject9 {
	return &NullableInlineObject9{value: val, isSet: true}
}

func (v NullableInlineObject9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
