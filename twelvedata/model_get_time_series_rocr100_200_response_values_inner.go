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

// checks if the GetTimeSeriesRocr100200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRocr100200ResponseValuesInner{}

// GetTimeSeriesRocr100200ResponseValuesInner struct for GetTimeSeriesRocr100200ResponseValuesInner
type GetTimeSeriesRocr100200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// rocr100 value
	Rocr100 string `json:"rocr100"`
}

type _GetTimeSeriesRocr100200ResponseValuesInner GetTimeSeriesRocr100200ResponseValuesInner

// NewGetTimeSeriesRocr100200ResponseValuesInner instantiates a new GetTimeSeriesRocr100200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRocr100200ResponseValuesInner(datetime string, rocr100 string) *GetTimeSeriesRocr100200ResponseValuesInner {
	this := GetTimeSeriesRocr100200ResponseValuesInner{}
	this.Datetime = datetime
	this.Rocr100 = rocr100
	return &this
}

// NewGetTimeSeriesRocr100200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRocr100200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRocr100200ResponseValuesInnerWithDefaults() *GetTimeSeriesRocr100200ResponseValuesInner {
	this := GetTimeSeriesRocr100200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesRocr100200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetRocr100 returns the Rocr100 field value
func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetRocr100() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rocr100
}

// GetRocr100Ok returns a tuple with the Rocr100 field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetRocr100Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rocr100, true
}

// SetRocr100 sets field value
func (o *GetTimeSeriesRocr100200ResponseValuesInner) SetRocr100(v string) {
	o.Rocr100 = v
}

func (o GetTimeSeriesRocr100200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRocr100200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["rocr100"] = o.Rocr100
	return toSerialize, nil
}

func (o *GetTimeSeriesRocr100200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"rocr100",
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

	varGetTimeSeriesRocr100200ResponseValuesInner := _GetTimeSeriesRocr100200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesRocr100200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRocr100200ResponseValuesInner(varGetTimeSeriesRocr100200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesRocr100200ResponseValuesInner struct {
	value *GetTimeSeriesRocr100200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesRocr100200ResponseValuesInner) Get() *GetTimeSeriesRocr100200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesRocr100200ResponseValuesInner) Set(val *GetTimeSeriesRocr100200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRocr100200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRocr100200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRocr100200ResponseValuesInner(val *GetTimeSeriesRocr100200ResponseValuesInner) *NullableGetTimeSeriesRocr100200ResponseValuesInner {
	return &NullableGetTimeSeriesRocr100200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRocr100200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRocr100200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
