// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMinIndex200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinIndex200ResponseValuesInner{}

// GetTimeSeriesMinIndex200ResponseValuesInner struct for GetTimeSeriesMinIndex200ResponseValuesInner
type GetTimeSeriesMinIndex200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Index of lowest value over period
	Minidx string `json:"minidx"`
}

type _GetTimeSeriesMinIndex200ResponseValuesInner GetTimeSeriesMinIndex200ResponseValuesInner

// NewGetTimeSeriesMinIndex200ResponseValuesInner instantiates a new GetTimeSeriesMinIndex200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinIndex200ResponseValuesInner(datetime string, minidx string) *GetTimeSeriesMinIndex200ResponseValuesInner {
	this := GetTimeSeriesMinIndex200ResponseValuesInner{}
	this.Datetime = datetime
	this.Minidx = minidx
	return &this
}

// NewGetTimeSeriesMinIndex200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinIndex200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinIndex200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinIndex200ResponseValuesInner {
	this := GetTimeSeriesMinIndex200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMinIndex200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMinidx returns the Minidx field value
func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetMinidx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Minidx
}

// GetMinidxOk returns a tuple with the Minidx field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetMinidxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minidx, true
}

// SetMinidx sets field value
func (o *GetTimeSeriesMinIndex200ResponseValuesInner) SetMinidx(v string) {
	o.Minidx = v
}

func (o GetTimeSeriesMinIndex200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinIndex200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["minidx"] = o.Minidx
	return toSerialize, nil
}

func (o *GetTimeSeriesMinIndex200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"minidx",
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

	varGetTimeSeriesMinIndex200ResponseValuesInner := _GetTimeSeriesMinIndex200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinIndex200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinIndex200ResponseValuesInner(varGetTimeSeriesMinIndex200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMinIndex200ResponseValuesInner struct {
	value *GetTimeSeriesMinIndex200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMinIndex200ResponseValuesInner) Get() *GetTimeSeriesMinIndex200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMinIndex200ResponseValuesInner) Set(val *GetTimeSeriesMinIndex200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinIndex200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinIndex200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinIndex200ResponseValuesInner(val *GetTimeSeriesMinIndex200ResponseValuesInner) *NullableGetTimeSeriesMinIndex200ResponseValuesInner {
	return &NullableGetTimeSeriesMinIndex200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinIndex200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinIndex200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
