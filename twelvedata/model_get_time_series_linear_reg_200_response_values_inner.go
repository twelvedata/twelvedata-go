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

// checks if the GetTimeSeriesLinearReg200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLinearReg200ResponseValuesInner{}

// GetTimeSeriesLinearReg200ResponseValuesInner struct for GetTimeSeriesLinearReg200ResponseValuesInner
type GetTimeSeriesLinearReg200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// linearreg value
	Linearreg string `json:"linearreg"`
}

type _GetTimeSeriesLinearReg200ResponseValuesInner GetTimeSeriesLinearReg200ResponseValuesInner

// NewGetTimeSeriesLinearReg200ResponseValuesInner instantiates a new GetTimeSeriesLinearReg200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLinearReg200ResponseValuesInner(datetime string, linearreg string) *GetTimeSeriesLinearReg200ResponseValuesInner {
	this := GetTimeSeriesLinearReg200ResponseValuesInner{}
	this.Datetime = datetime
	this.Linearreg = linearreg
	return &this
}

// NewGetTimeSeriesLinearReg200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLinearReg200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLinearReg200ResponseValuesInnerWithDefaults() *GetTimeSeriesLinearReg200ResponseValuesInner {
	this := GetTimeSeriesLinearReg200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesLinearReg200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetLinearreg returns the Linearreg field value
func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetLinearreg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Linearreg
}

// GetLinearregOk returns a tuple with the Linearreg field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetLinearregOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linearreg, true
}

// SetLinearreg sets field value
func (o *GetTimeSeriesLinearReg200ResponseValuesInner) SetLinearreg(v string) {
	o.Linearreg = v
}

func (o GetTimeSeriesLinearReg200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLinearReg200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["linearreg"] = o.Linearreg
	return toSerialize, nil
}

func (o *GetTimeSeriesLinearReg200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"linearreg",
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

	varGetTimeSeriesLinearReg200ResponseValuesInner := _GetTimeSeriesLinearReg200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesLinearReg200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLinearReg200ResponseValuesInner(varGetTimeSeriesLinearReg200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesLinearReg200ResponseValuesInner struct {
	value *GetTimeSeriesLinearReg200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesLinearReg200ResponseValuesInner) Get() *GetTimeSeriesLinearReg200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesLinearReg200ResponseValuesInner) Set(val *GetTimeSeriesLinearReg200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLinearReg200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLinearReg200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLinearReg200ResponseValuesInner(val *GetTimeSeriesLinearReg200ResponseValuesInner) *NullableGetTimeSeriesLinearReg200ResponseValuesInner {
	return &NullableGetTimeSeriesLinearReg200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLinearReg200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLinearReg200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
