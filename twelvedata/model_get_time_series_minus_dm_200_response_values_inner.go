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

// checks if the GetTimeSeriesMinusDM200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinusDM200ResponseValuesInner{}

// GetTimeSeriesMinusDM200ResponseValuesInner struct for GetTimeSeriesMinusDM200ResponseValuesInner
type GetTimeSeriesMinusDM200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Minus Directional Movement value
	MinusDm string `json:"minus_dm"`
}

type _GetTimeSeriesMinusDM200ResponseValuesInner GetTimeSeriesMinusDM200ResponseValuesInner

// NewGetTimeSeriesMinusDM200ResponseValuesInner instantiates a new GetTimeSeriesMinusDM200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinusDM200ResponseValuesInner(datetime string, minusDm string) *GetTimeSeriesMinusDM200ResponseValuesInner {
	this := GetTimeSeriesMinusDM200ResponseValuesInner{}
	this.Datetime = datetime
	this.MinusDm = minusDm
	return &this
}

// NewGetTimeSeriesMinusDM200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinusDM200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinusDM200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinusDM200ResponseValuesInner {
	this := GetTimeSeriesMinusDM200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMinusDM200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMinusDm returns the MinusDm field value
func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetMinusDm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinusDm
}

// GetMinusDmOk returns a tuple with the MinusDm field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetMinusDmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinusDm, true
}

// SetMinusDm sets field value
func (o *GetTimeSeriesMinusDM200ResponseValuesInner) SetMinusDm(v string) {
	o.MinusDm = v
}

func (o GetTimeSeriesMinusDM200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinusDM200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["minus_dm"] = o.MinusDm
	return toSerialize, nil
}

func (o *GetTimeSeriesMinusDM200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"minus_dm",
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

	varGetTimeSeriesMinusDM200ResponseValuesInner := _GetTimeSeriesMinusDM200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinusDM200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinusDM200ResponseValuesInner(varGetTimeSeriesMinusDM200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMinusDM200ResponseValuesInner struct {
	value *GetTimeSeriesMinusDM200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMinusDM200ResponseValuesInner) Get() *GetTimeSeriesMinusDM200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMinusDM200ResponseValuesInner) Set(val *GetTimeSeriesMinusDM200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinusDM200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinusDM200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinusDM200ResponseValuesInner(val *GetTimeSeriesMinusDM200ResponseValuesInner) *NullableGetTimeSeriesMinusDM200ResponseValuesInner {
	return &NullableGetTimeSeriesMinusDM200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinusDM200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinusDM200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
