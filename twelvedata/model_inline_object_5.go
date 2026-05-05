// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject5{}

// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	Calls []OptionSide `json:"calls,omitempty"`
	Puts  []OptionSide `json:"puts,omitempty"`
}

// NewInlineObject5 instantiates a new InlineObject5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject5() *InlineObject5 {
	this := InlineObject5{}
	return &this
}

// NewInlineObject5WithDefaults instantiates a new InlineObject5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject5WithDefaults() *InlineObject5 {
	this := InlineObject5{}
	return &this
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *InlineObject5) GetCalls() []OptionSide {
	if o == nil || IsNil(o.Calls) {
		var ret []OptionSide
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetCallsOk() ([]OptionSide, bool) {
	if o == nil || IsNil(o.Calls) {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *InlineObject5) HasCalls() bool {
	if o != nil && !IsNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given []OptionSide and assigns it to the Calls field.
func (o *InlineObject5) SetCalls(v []OptionSide) {
	o.Calls = v
}

// GetPuts returns the Puts field value if set, zero value otherwise.
func (o *InlineObject5) GetPuts() []OptionSide {
	if o == nil || IsNil(o.Puts) {
		var ret []OptionSide
		return ret
	}
	return o.Puts
}

// GetPutsOk returns a tuple with the Puts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetPutsOk() ([]OptionSide, bool) {
	if o == nil || IsNil(o.Puts) {
		return nil, false
	}
	return o.Puts, true
}

// HasPuts returns a boolean if a field has been set.
func (o *InlineObject5) HasPuts() bool {
	if o != nil && !IsNil(o.Puts) {
		return true
	}

	return false
}

// SetPuts gets a reference to the given []OptionSide and assigns it to the Puts field.
func (o *InlineObject5) SetPuts(v []OptionSide) {
	o.Puts = v
}

func (o InlineObject5) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Calls) {
		toSerialize["calls"] = o.Calls
	}
	if !IsNil(o.Puts) {
		toSerialize["puts"] = o.Puts
	}
	return toSerialize, nil
}

type NullableInlineObject5 struct {
	value *InlineObject5
	isSet bool
}

func (v NullableInlineObject5) Get() *InlineObject5 {
	return v.value
}

func (v *NullableInlineObject5) Set(val *InlineObject5) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject5) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject5(val *InlineObject5) *NullableInlineObject5 {
	return &NullableInlineObject5{value: val, isSet: true}
}

func (v NullableInlineObject5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
