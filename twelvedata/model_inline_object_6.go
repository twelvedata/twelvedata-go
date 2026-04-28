/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject6{}

// InlineObject6 struct for InlineObject6
type InlineObject6 struct {
	Dates []string `json:"dates,omitempty"`
}

// NewInlineObject6 instantiates a new InlineObject6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject6() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// NewInlineObject6WithDefaults instantiates a new InlineObject6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject6WithDefaults() *InlineObject6 {
	this := InlineObject6{}
	return &this
}

// GetDates returns the Dates field value if set, zero value otherwise.
func (o *InlineObject6) GetDates() []string {
	if o == nil || IsNil(o.Dates) {
		var ret []string
		return ret
	}
	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject6) GetDatesOk() ([]string, bool) {
	if o == nil || IsNil(o.Dates) {
		return nil, false
	}
	return o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *InlineObject6) HasDates() bool {
	if o != nil && !IsNil(o.Dates) {
		return true
	}

	return false
}

// SetDates gets a reference to the given []string and assigns it to the Dates field.
func (o *InlineObject6) SetDates(v []string) {
	o.Dates = v
}

func (o InlineObject6) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dates) {
		toSerialize["dates"] = o.Dates
	}
	return toSerialize, nil
}

type NullableInlineObject6 struct {
	value *InlineObject6
	isSet bool
}

func (v NullableInlineObject6) Get() *InlineObject6 {
	return v.value
}

func (v *NullableInlineObject6) Set(val *InlineObject6) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject6) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject6(val *InlineObject6) *NullableInlineObject6 {
	return &NullableInlineObject6{value: val, isSet: true}
}

func (v NullableInlineObject6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
