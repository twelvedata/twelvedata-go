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

// checks if the GetMarketCap200ResponseMarketCapInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarketCap200ResponseMarketCapInner{}

// GetMarketCap200ResponseMarketCapInner struct for GetMarketCap200ResponseMarketCapInner
type GetMarketCap200ResponseMarketCapInner struct {
	// Market capitalization date
	Date string `json:"date"`
	// Market capitalization value
	Value int64 `json:"value"`
}

type _GetMarketCap200ResponseMarketCapInner GetMarketCap200ResponseMarketCapInner

// NewGetMarketCap200ResponseMarketCapInner instantiates a new GetMarketCap200ResponseMarketCapInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarketCap200ResponseMarketCapInner(date string, value int64) *GetMarketCap200ResponseMarketCapInner {
	this := GetMarketCap200ResponseMarketCapInner{}
	this.Date = date
	this.Value = value
	return &this
}

// NewGetMarketCap200ResponseMarketCapInnerWithDefaults instantiates a new GetMarketCap200ResponseMarketCapInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarketCap200ResponseMarketCapInnerWithDefaults() *GetMarketCap200ResponseMarketCapInner {
	this := GetMarketCap200ResponseMarketCapInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetMarketCap200ResponseMarketCapInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetMarketCap200ResponseMarketCapInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetMarketCap200ResponseMarketCapInner) SetDate(v string) {
	o.Date = v
}

// GetValue returns the Value field value
func (o *GetMarketCap200ResponseMarketCapInner) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetMarketCap200ResponseMarketCapInner) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetMarketCap200ResponseMarketCapInner) SetValue(v int64) {
	o.Value = v
}

func (o GetMarketCap200ResponseMarketCapInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarketCap200ResponseMarketCapInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *GetMarketCap200ResponseMarketCapInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"value",
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

	varGetMarketCap200ResponseMarketCapInner := _GetMarketCap200ResponseMarketCapInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetMarketCap200ResponseMarketCapInner)

	if err != nil {
		return err
	}

	*o = GetMarketCap200ResponseMarketCapInner(varGetMarketCap200ResponseMarketCapInner)

	return err
}

type NullableGetMarketCap200ResponseMarketCapInner struct {
	value *GetMarketCap200ResponseMarketCapInner
	isSet bool
}

func (v NullableGetMarketCap200ResponseMarketCapInner) Get() *GetMarketCap200ResponseMarketCapInner {
	return v.value
}

func (v *NullableGetMarketCap200ResponseMarketCapInner) Set(val *GetMarketCap200ResponseMarketCapInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarketCap200ResponseMarketCapInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarketCap200ResponseMarketCapInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarketCap200ResponseMarketCapInner(val *GetMarketCap200ResponseMarketCapInner) *NullableGetMarketCap200ResponseMarketCapInner {
	return &NullableGetMarketCap200ResponseMarketCapInner{value: val, isSet: true}
}

func (v NullableGetMarketCap200ResponseMarketCapInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarketCap200ResponseMarketCapInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
