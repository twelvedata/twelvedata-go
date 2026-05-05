// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesRvol200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRvol200ResponseValuesInner{}

// GetTimeSeriesRvol200ResponseValuesInner struct for GetTimeSeriesRvol200ResponseValuesInner
type GetTimeSeriesRvol200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// RVOL value
	Rvol string `json:"rvol"`
}

type _GetTimeSeriesRvol200ResponseValuesInner GetTimeSeriesRvol200ResponseValuesInner

// NewGetTimeSeriesRvol200ResponseValuesInner instantiates a new GetTimeSeriesRvol200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRvol200ResponseValuesInner(datetime string, rvol string) *GetTimeSeriesRvol200ResponseValuesInner {
	this := GetTimeSeriesRvol200ResponseValuesInner{}
	this.Datetime = datetime
	this.Rvol = rvol
	return &this
}

// NewGetTimeSeriesRvol200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRvol200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRvol200ResponseValuesInnerWithDefaults() *GetTimeSeriesRvol200ResponseValuesInner {
	this := GetTimeSeriesRvol200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesRvol200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRvol200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesRvol200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetRvol returns the Rvol field value
func (o *GetTimeSeriesRvol200ResponseValuesInner) GetRvol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rvol
}

// GetRvolOk returns a tuple with the Rvol field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRvol200ResponseValuesInner) GetRvolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rvol, true
}

// SetRvol sets field value
func (o *GetTimeSeriesRvol200ResponseValuesInner) SetRvol(v string) {
	o.Rvol = v
}

func (o GetTimeSeriesRvol200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRvol200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["rvol"] = o.Rvol
	return toSerialize, nil
}

func (o *GetTimeSeriesRvol200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"rvol",
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

	varGetTimeSeriesRvol200ResponseValuesInner := _GetTimeSeriesRvol200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesRvol200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRvol200ResponseValuesInner(varGetTimeSeriesRvol200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesRvol200ResponseValuesInner struct {
	value *GetTimeSeriesRvol200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesRvol200ResponseValuesInner) Get() *GetTimeSeriesRvol200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesRvol200ResponseValuesInner) Set(val *GetTimeSeriesRvol200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRvol200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRvol200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRvol200ResponseValuesInner(val *GetTimeSeriesRvol200ResponseValuesInner) *NullableGetTimeSeriesRvol200ResponseValuesInner {
	return &NullableGetTimeSeriesRvol200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRvol200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRvol200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
