// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the InlineObject11ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject11ValuesInner{}

// InlineObject11ValuesInner struct for InlineObject11ValuesInner
type InlineObject11ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// COS value
	Cos string `json:"cos"`
}

type _InlineObject11ValuesInner InlineObject11ValuesInner

// NewInlineObject11ValuesInner instantiates a new InlineObject11ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject11ValuesInner(datetime string, cos string) *InlineObject11ValuesInner {
	this := InlineObject11ValuesInner{}
	this.Datetime = datetime
	this.Cos = cos
	return &this
}

// NewInlineObject11ValuesInnerWithDefaults instantiates a new InlineObject11ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject11ValuesInnerWithDefaults() *InlineObject11ValuesInner {
	this := InlineObject11ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject11ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject11ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject11ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCos returns the Cos field value
func (o *InlineObject11ValuesInner) GetCos() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cos
}

// GetCosOk returns a tuple with the Cos field value
// and a boolean to check if the value has been set.
func (o *InlineObject11ValuesInner) GetCosOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cos, true
}

// SetCos sets field value
func (o *InlineObject11ValuesInner) SetCos(v string) {
	o.Cos = v
}

func (o InlineObject11ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject11ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["cos"] = o.Cos
	return toSerialize, nil
}

func (o *InlineObject11ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"cos",
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

	varInlineObject11ValuesInner := _InlineObject11ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInlineObject11ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject11ValuesInner(varInlineObject11ValuesInner)

	return err
}

type NullableInlineObject11ValuesInner struct {
	value *InlineObject11ValuesInner
	isSet bool
}

func (v NullableInlineObject11ValuesInner) Get() *InlineObject11ValuesInner {
	return v.value
}

func (v *NullableInlineObject11ValuesInner) Set(val *InlineObject11ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject11ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject11ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject11ValuesInner(val *InlineObject11ValuesInner) *NullableInlineObject11ValuesInner {
	return &NullableInlineObject11ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject11ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject11ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
