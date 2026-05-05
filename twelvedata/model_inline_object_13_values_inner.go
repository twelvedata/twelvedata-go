// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject13ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject13ValuesInner{}

// InlineObject13ValuesInner struct for InlineObject13ValuesInner
type InlineObject13ValuesInner struct {
	Datetime *string `json:"datetime,omitempty"`
	Mavp     *string `json:"mavp,omitempty"`
}

// NewInlineObject13ValuesInner instantiates a new InlineObject13ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject13ValuesInner() *InlineObject13ValuesInner {
	this := InlineObject13ValuesInner{}
	return &this
}

// NewInlineObject13ValuesInnerWithDefaults instantiates a new InlineObject13ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject13ValuesInnerWithDefaults() *InlineObject13ValuesInner {
	this := InlineObject13ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value if set, zero value otherwise.
func (o *InlineObject13ValuesInner) GetDatetime() string {
	if o == nil || IsNil(o.Datetime) {
		var ret string
		return ret
	}
	return *o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil || IsNil(o.Datetime) {
		return nil, false
	}
	return o.Datetime, true
}

// HasDatetime returns a boolean if a field has been set.
func (o *InlineObject13ValuesInner) HasDatetime() bool {
	if o != nil && !IsNil(o.Datetime) {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given string and assigns it to the Datetime field.
func (o *InlineObject13ValuesInner) SetDatetime(v string) {
	o.Datetime = &v
}

// GetMavp returns the Mavp field value if set, zero value otherwise.
func (o *InlineObject13ValuesInner) GetMavp() string {
	if o == nil || IsNil(o.Mavp) {
		var ret string
		return ret
	}
	return *o.Mavp
}

// GetMavpOk returns a tuple with the Mavp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject13ValuesInner) GetMavpOk() (*string, bool) {
	if o == nil || IsNil(o.Mavp) {
		return nil, false
	}
	return o.Mavp, true
}

// HasMavp returns a boolean if a field has been set.
func (o *InlineObject13ValuesInner) HasMavp() bool {
	if o != nil && !IsNil(o.Mavp) {
		return true
	}

	return false
}

// SetMavp gets a reference to the given string and assigns it to the Mavp field.
func (o *InlineObject13ValuesInner) SetMavp(v string) {
	o.Mavp = &v
}

func (o InlineObject13ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject13ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Datetime) {
		toSerialize["datetime"] = o.Datetime
	}
	if !IsNil(o.Mavp) {
		toSerialize["mavp"] = o.Mavp
	}
	return toSerialize, nil
}

type NullableInlineObject13ValuesInner struct {
	value *InlineObject13ValuesInner
	isSet bool
}

func (v NullableInlineObject13ValuesInner) Get() *InlineObject13ValuesInner {
	return v.value
}

func (v *NullableInlineObject13ValuesInner) Set(val *InlineObject13ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject13ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject13ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject13ValuesInner(val *InlineObject13ValuesInner) *NullableInlineObject13ValuesInner {
	return &NullableInlineObject13ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject13ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject13ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
