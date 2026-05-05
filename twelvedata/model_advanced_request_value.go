// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the AdvancedRequestValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedRequestValue{}

// AdvancedRequestValue struct for AdvancedRequestValue
type AdvancedRequestValue struct {
	// Requested url
	Url *string `json:"url,omitempty"`
}

// NewAdvancedRequestValue instantiates a new AdvancedRequestValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedRequestValue() *AdvancedRequestValue {
	this := AdvancedRequestValue{}
	return &this
}

// NewAdvancedRequestValueWithDefaults instantiates a new AdvancedRequestValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedRequestValueWithDefaults() *AdvancedRequestValue {
	this := AdvancedRequestValue{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AdvancedRequestValue) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedRequestValue) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AdvancedRequestValue) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AdvancedRequestValue) SetUrl(v string) {
	o.Url = &v
}

func (o AdvancedRequestValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedRequestValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAdvancedRequestValue struct {
	value *AdvancedRequestValue
	isSet bool
}

func (v NullableAdvancedRequestValue) Get() *AdvancedRequestValue {
	return v.value
}

func (v *NullableAdvancedRequestValue) Set(val *AdvancedRequestValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedRequestValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedRequestValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedRequestValue(val *AdvancedRequestValue) *NullableAdvancedRequestValue {
	return &NullableAdvancedRequestValue{value: val, isSet: true}
}

func (v NullableAdvancedRequestValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedRequestValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
