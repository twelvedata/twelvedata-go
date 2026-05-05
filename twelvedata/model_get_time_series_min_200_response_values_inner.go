// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMin200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMin200ResponseValuesInner{}

// GetTimeSeriesMin200ResponseValuesInner struct for GetTimeSeriesMin200ResponseValuesInner
type GetTimeSeriesMin200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Min value
	Min string `json:"min"`
}

type _GetTimeSeriesMin200ResponseValuesInner GetTimeSeriesMin200ResponseValuesInner

// NewGetTimeSeriesMin200ResponseValuesInner instantiates a new GetTimeSeriesMin200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMin200ResponseValuesInner(datetime string, min string) *GetTimeSeriesMin200ResponseValuesInner {
	this := GetTimeSeriesMin200ResponseValuesInner{}
	this.Datetime = datetime
	this.Min = min
	return &this
}

// NewGetTimeSeriesMin200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMin200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMin200ResponseValuesInnerWithDefaults() *GetTimeSeriesMin200ResponseValuesInner {
	this := GetTimeSeriesMin200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMin200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMin200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMin returns the Min field value
func (o *GetTimeSeriesMin200ResponseValuesInner) GetMin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMin200ResponseValuesInner) GetMinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value
func (o *GetTimeSeriesMin200ResponseValuesInner) SetMin(v string) {
	o.Min = v
}

func (o GetTimeSeriesMin200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMin200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["min"] = o.Min
	return toSerialize, nil
}

func (o *GetTimeSeriesMin200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"min",
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

	varGetTimeSeriesMin200ResponseValuesInner := _GetTimeSeriesMin200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMin200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMin200ResponseValuesInner(varGetTimeSeriesMin200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMin200ResponseValuesInner struct {
	value *GetTimeSeriesMin200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMin200ResponseValuesInner) Get() *GetTimeSeriesMin200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMin200ResponseValuesInner) Set(val *GetTimeSeriesMin200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMin200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMin200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMin200ResponseValuesInner(val *GetTimeSeriesMin200ResponseValuesInner) *NullableGetTimeSeriesMin200ResponseValuesInner {
	return &NullableGetTimeSeriesMin200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMin200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMin200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
