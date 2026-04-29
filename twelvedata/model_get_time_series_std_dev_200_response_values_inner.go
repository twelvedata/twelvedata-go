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

// checks if the GetTimeSeriesStdDev200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStdDev200ResponseValuesInner{}

// GetTimeSeriesStdDev200ResponseValuesInner struct for GetTimeSeriesStdDev200ResponseValuesInner
type GetTimeSeriesStdDev200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Standard Deviation value
	Stddev string `json:"stddev"`
}

type _GetTimeSeriesStdDev200ResponseValuesInner GetTimeSeriesStdDev200ResponseValuesInner

// NewGetTimeSeriesStdDev200ResponseValuesInner instantiates a new GetTimeSeriesStdDev200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStdDev200ResponseValuesInner(datetime string, stddev string) *GetTimeSeriesStdDev200ResponseValuesInner {
	this := GetTimeSeriesStdDev200ResponseValuesInner{}
	this.Datetime = datetime
	this.Stddev = stddev
	return &this
}

// NewGetTimeSeriesStdDev200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesStdDev200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStdDev200ResponseValuesInnerWithDefaults() *GetTimeSeriesStdDev200ResponseValuesInner {
	this := GetTimeSeriesStdDev200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesStdDev200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetStddev returns the Stddev field value
func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetStddev() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stddev
}

// GetStddevOk returns a tuple with the Stddev field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetStddevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stddev, true
}

// SetStddev sets field value
func (o *GetTimeSeriesStdDev200ResponseValuesInner) SetStddev(v string) {
	o.Stddev = v
}

func (o GetTimeSeriesStdDev200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStdDev200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["stddev"] = o.Stddev
	return toSerialize, nil
}

func (o *GetTimeSeriesStdDev200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"stddev",
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

	varGetTimeSeriesStdDev200ResponseValuesInner := _GetTimeSeriesStdDev200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStdDev200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStdDev200ResponseValuesInner(varGetTimeSeriesStdDev200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesStdDev200ResponseValuesInner struct {
	value *GetTimeSeriesStdDev200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesStdDev200ResponseValuesInner) Get() *GetTimeSeriesStdDev200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesStdDev200ResponseValuesInner) Set(val *GetTimeSeriesStdDev200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStdDev200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStdDev200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStdDev200ResponseValuesInner(val *GetTimeSeriesStdDev200ResponseValuesInner) *NullableGetTimeSeriesStdDev200ResponseValuesInner {
	return &NullableGetTimeSeriesStdDev200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStdDev200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStdDev200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
