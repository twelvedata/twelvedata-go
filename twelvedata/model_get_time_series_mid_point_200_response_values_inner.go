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

// checks if the GetTimeSeriesMidPoint200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMidPoint200ResponseValuesInner{}

// GetTimeSeriesMidPoint200ResponseValuesInner struct for GetTimeSeriesMidPoint200ResponseValuesInner
type GetTimeSeriesMidPoint200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Midpoint value
	Midpoint string `json:"midpoint"`
}

type _GetTimeSeriesMidPoint200ResponseValuesInner GetTimeSeriesMidPoint200ResponseValuesInner

// NewGetTimeSeriesMidPoint200ResponseValuesInner instantiates a new GetTimeSeriesMidPoint200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMidPoint200ResponseValuesInner(datetime string, midpoint string) *GetTimeSeriesMidPoint200ResponseValuesInner {
	this := GetTimeSeriesMidPoint200ResponseValuesInner{}
	this.Datetime = datetime
	this.Midpoint = midpoint
	return &this
}

// NewGetTimeSeriesMidPoint200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMidPoint200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMidPoint200ResponseValuesInnerWithDefaults() *GetTimeSeriesMidPoint200ResponseValuesInner {
	this := GetTimeSeriesMidPoint200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMidPoint200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMidpoint returns the Midpoint field value
func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetMidpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Midpoint
}

// GetMidpointOk returns a tuple with the Midpoint field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetMidpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Midpoint, true
}

// SetMidpoint sets field value
func (o *GetTimeSeriesMidPoint200ResponseValuesInner) SetMidpoint(v string) {
	o.Midpoint = v
}

func (o GetTimeSeriesMidPoint200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMidPoint200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["midpoint"] = o.Midpoint
	return toSerialize, nil
}

func (o *GetTimeSeriesMidPoint200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"midpoint",
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

	varGetTimeSeriesMidPoint200ResponseValuesInner := _GetTimeSeriesMidPoint200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMidPoint200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMidPoint200ResponseValuesInner(varGetTimeSeriesMidPoint200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMidPoint200ResponseValuesInner struct {
	value *GetTimeSeriesMidPoint200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMidPoint200ResponseValuesInner) Get() *GetTimeSeriesMidPoint200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMidPoint200ResponseValuesInner) Set(val *GetTimeSeriesMidPoint200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMidPoint200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMidPoint200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMidPoint200ResponseValuesInner(val *GetTimeSeriesMidPoint200ResponseValuesInner) *NullableGetTimeSeriesMidPoint200ResponseValuesInner {
	return &NullableGetTimeSeriesMidPoint200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMidPoint200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMidPoint200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
