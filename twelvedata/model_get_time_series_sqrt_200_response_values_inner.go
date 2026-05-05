// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesSqrt200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSqrt200ResponseValuesInner{}

// GetTimeSeriesSqrt200ResponseValuesInner struct for GetTimeSeriesSqrt200ResponseValuesInner
type GetTimeSeriesSqrt200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SQRT value
	Sqrt string `json:"sqrt"`
}

type _GetTimeSeriesSqrt200ResponseValuesInner GetTimeSeriesSqrt200ResponseValuesInner

// NewGetTimeSeriesSqrt200ResponseValuesInner instantiates a new GetTimeSeriesSqrt200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSqrt200ResponseValuesInner(datetime string, sqrt string) *GetTimeSeriesSqrt200ResponseValuesInner {
	this := GetTimeSeriesSqrt200ResponseValuesInner{}
	this.Datetime = datetime
	this.Sqrt = sqrt
	return &this
}

// NewGetTimeSeriesSqrt200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSqrt200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSqrt200ResponseValuesInnerWithDefaults() *GetTimeSeriesSqrt200ResponseValuesInner {
	this := GetTimeSeriesSqrt200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSqrt200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSqrt returns the Sqrt field value
func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetSqrt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sqrt
}

// GetSqrtOk returns a tuple with the Sqrt field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetSqrtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sqrt, true
}

// SetSqrt sets field value
func (o *GetTimeSeriesSqrt200ResponseValuesInner) SetSqrt(v string) {
	o.Sqrt = v
}

func (o GetTimeSeriesSqrt200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSqrt200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sqrt"] = o.Sqrt
	return toSerialize, nil
}

func (o *GetTimeSeriesSqrt200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sqrt",
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

	varGetTimeSeriesSqrt200ResponseValuesInner := _GetTimeSeriesSqrt200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSqrt200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSqrt200ResponseValuesInner(varGetTimeSeriesSqrt200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSqrt200ResponseValuesInner struct {
	value *GetTimeSeriesSqrt200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSqrt200ResponseValuesInner) Get() *GetTimeSeriesSqrt200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSqrt200ResponseValuesInner) Set(val *GetTimeSeriesSqrt200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSqrt200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSqrt200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSqrt200ResponseValuesInner(val *GetTimeSeriesSqrt200ResponseValuesInner) *NullableGetTimeSeriesSqrt200ResponseValuesInner {
	return &NullableGetTimeSeriesSqrt200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSqrt200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSqrt200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
