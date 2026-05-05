// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesSum200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSum200ResponseValuesInner{}

// GetTimeSeriesSum200ResponseValuesInner struct for GetTimeSeriesSum200ResponseValuesInner
type GetTimeSeriesSum200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Sum value
	Sum string `json:"sum"`
}

type _GetTimeSeriesSum200ResponseValuesInner GetTimeSeriesSum200ResponseValuesInner

// NewGetTimeSeriesSum200ResponseValuesInner instantiates a new GetTimeSeriesSum200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSum200ResponseValuesInner(datetime string, sum string) *GetTimeSeriesSum200ResponseValuesInner {
	this := GetTimeSeriesSum200ResponseValuesInner{}
	this.Datetime = datetime
	this.Sum = sum
	return &this
}

// NewGetTimeSeriesSum200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSum200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSum200ResponseValuesInnerWithDefaults() *GetTimeSeriesSum200ResponseValuesInner {
	this := GetTimeSeriesSum200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSum200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSum200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSum200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSum returns the Sum field value
func (o *GetTimeSeriesSum200ResponseValuesInner) GetSum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sum
}

// GetSumOk returns a tuple with the Sum field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSum200ResponseValuesInner) GetSumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sum, true
}

// SetSum sets field value
func (o *GetTimeSeriesSum200ResponseValuesInner) SetSum(v string) {
	o.Sum = v
}

func (o GetTimeSeriesSum200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSum200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sum"] = o.Sum
	return toSerialize, nil
}

func (o *GetTimeSeriesSum200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sum",
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

	varGetTimeSeriesSum200ResponseValuesInner := _GetTimeSeriesSum200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSum200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSum200ResponseValuesInner(varGetTimeSeriesSum200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSum200ResponseValuesInner struct {
	value *GetTimeSeriesSum200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSum200ResponseValuesInner) Get() *GetTimeSeriesSum200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSum200ResponseValuesInner) Set(val *GetTimeSeriesSum200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSum200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSum200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSum200ResponseValuesInner(val *GetTimeSeriesSum200ResponseValuesInner) *NullableGetTimeSeriesSum200ResponseValuesInner {
	return &NullableGetTimeSeriesSum200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSum200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSum200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
