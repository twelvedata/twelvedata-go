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

// checks if the GetTimeSeriesExp200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesExp200ResponseValuesInner{}

// GetTimeSeriesExp200ResponseValuesInner struct for GetTimeSeriesExp200ResponseValuesInner
type GetTimeSeriesExp200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Exp value
	Exp string `json:"exp"`
}

type _GetTimeSeriesExp200ResponseValuesInner GetTimeSeriesExp200ResponseValuesInner

// NewGetTimeSeriesExp200ResponseValuesInner instantiates a new GetTimeSeriesExp200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesExp200ResponseValuesInner(datetime string, exp string) *GetTimeSeriesExp200ResponseValuesInner {
	this := GetTimeSeriesExp200ResponseValuesInner{}
	this.Datetime = datetime
	this.Exp = exp
	return &this
}

// NewGetTimeSeriesExp200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesExp200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesExp200ResponseValuesInnerWithDefaults() *GetTimeSeriesExp200ResponseValuesInner {
	this := GetTimeSeriesExp200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesExp200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesExp200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesExp200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetExp returns the Exp field value
func (o *GetTimeSeriesExp200ResponseValuesInner) GetExp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exp
}

// GetExpOk returns a tuple with the Exp field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesExp200ResponseValuesInner) GetExpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exp, true
}

// SetExp sets field value
func (o *GetTimeSeriesExp200ResponseValuesInner) SetExp(v string) {
	o.Exp = v
}

func (o GetTimeSeriesExp200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesExp200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["exp"] = o.Exp
	return toSerialize, nil
}

func (o *GetTimeSeriesExp200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"exp",
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

	varGetTimeSeriesExp200ResponseValuesInner := _GetTimeSeriesExp200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesExp200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesExp200ResponseValuesInner(varGetTimeSeriesExp200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesExp200ResponseValuesInner struct {
	value *GetTimeSeriesExp200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesExp200ResponseValuesInner) Get() *GetTimeSeriesExp200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesExp200ResponseValuesInner) Set(val *GetTimeSeriesExp200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesExp200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesExp200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesExp200ResponseValuesInner(val *GetTimeSeriesExp200ResponseValuesInner) *NullableGetTimeSeriesExp200ResponseValuesInner {
	return &NullableGetTimeSeriesExp200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesExp200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesExp200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
