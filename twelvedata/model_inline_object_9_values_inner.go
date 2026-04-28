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

// checks if the InlineObject9ValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject9ValuesInner{}

// InlineObject9ValuesInner struct for InlineObject9ValuesInner
type InlineObject9ValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ASIN value
	Asin string `json:"asin"`
}

type _InlineObject9ValuesInner InlineObject9ValuesInner

// NewInlineObject9ValuesInner instantiates a new InlineObject9ValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject9ValuesInner(datetime string, asin string) *InlineObject9ValuesInner {
	this := InlineObject9ValuesInner{}
	this.Datetime = datetime
	this.Asin = asin
	return &this
}

// NewInlineObject9ValuesInnerWithDefaults instantiates a new InlineObject9ValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject9ValuesInnerWithDefaults() *InlineObject9ValuesInner {
	this := InlineObject9ValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *InlineObject9ValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *InlineObject9ValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *InlineObject9ValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAsin returns the Asin field value
func (o *InlineObject9ValuesInner) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *InlineObject9ValuesInner) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *InlineObject9ValuesInner) SetAsin(v string) {
	o.Asin = v
}

func (o InlineObject9ValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject9ValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["asin"] = o.Asin
	return toSerialize, nil
}

func (o *InlineObject9ValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"asin",
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

	varInlineObject9ValuesInner := _InlineObject9ValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineObject9ValuesInner)

	if err != nil {
		return err
	}

	*o = InlineObject9ValuesInner(varInlineObject9ValuesInner)

	return err
}

type NullableInlineObject9ValuesInner struct {
	value *InlineObject9ValuesInner
	isSet bool
}

func (v NullableInlineObject9ValuesInner) Get() *InlineObject9ValuesInner {
	return v.value
}

func (v *NullableInlineObject9ValuesInner) Set(val *InlineObject9ValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject9ValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject9ValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject9ValuesInner(val *InlineObject9ValuesInner) *NullableInlineObject9ValuesInner {
	return &NullableInlineObject9ValuesInner{value: val, isSet: true}
}

func (v NullableInlineObject9ValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject9ValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
