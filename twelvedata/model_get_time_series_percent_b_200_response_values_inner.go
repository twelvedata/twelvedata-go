// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesPercentB200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPercentB200ResponseValuesInner{}

// GetTimeSeriesPercentB200ResponseValuesInner struct for GetTimeSeriesPercentB200ResponseValuesInner
type GetTimeSeriesPercentB200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Percent_b value
	PercentB string `json:"percent_b"`
}

type _GetTimeSeriesPercentB200ResponseValuesInner GetTimeSeriesPercentB200ResponseValuesInner

// NewGetTimeSeriesPercentB200ResponseValuesInner instantiates a new GetTimeSeriesPercentB200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPercentB200ResponseValuesInner(datetime string, percentB string) *GetTimeSeriesPercentB200ResponseValuesInner {
	this := GetTimeSeriesPercentB200ResponseValuesInner{}
	this.Datetime = datetime
	this.PercentB = percentB
	return &this
}

// NewGetTimeSeriesPercentB200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPercentB200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPercentB200ResponseValuesInnerWithDefaults() *GetTimeSeriesPercentB200ResponseValuesInner {
	this := GetTimeSeriesPercentB200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesPercentB200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetPercentB returns the PercentB field value
func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetPercentB() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PercentB
}

// GetPercentBOk returns a tuple with the PercentB field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetPercentBOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentB, true
}

// SetPercentB sets field value
func (o *GetTimeSeriesPercentB200ResponseValuesInner) SetPercentB(v string) {
	o.PercentB = v
}

func (o GetTimeSeriesPercentB200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPercentB200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["percent_b"] = o.PercentB
	return toSerialize, nil
}

func (o *GetTimeSeriesPercentB200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"percent_b",
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

	varGetTimeSeriesPercentB200ResponseValuesInner := _GetTimeSeriesPercentB200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesPercentB200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPercentB200ResponseValuesInner(varGetTimeSeriesPercentB200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesPercentB200ResponseValuesInner struct {
	value *GetTimeSeriesPercentB200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesPercentB200ResponseValuesInner) Get() *GetTimeSeriesPercentB200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesPercentB200ResponseValuesInner) Set(val *GetTimeSeriesPercentB200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPercentB200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPercentB200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPercentB200ResponseValuesInner(val *GetTimeSeriesPercentB200ResponseValuesInner) *NullableGetTimeSeriesPercentB200ResponseValuesInner {
	return &NullableGetTimeSeriesPercentB200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPercentB200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPercentB200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
