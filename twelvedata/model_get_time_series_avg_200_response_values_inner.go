// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAvg200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAvg200ResponseValuesInner{}

// GetTimeSeriesAvg200ResponseValuesInner struct for GetTimeSeriesAvg200ResponseValuesInner
type GetTimeSeriesAvg200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Avg value
	Avg string `json:"avg"`
}

type _GetTimeSeriesAvg200ResponseValuesInner GetTimeSeriesAvg200ResponseValuesInner

// NewGetTimeSeriesAvg200ResponseValuesInner instantiates a new GetTimeSeriesAvg200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAvg200ResponseValuesInner(datetime string, avg string) *GetTimeSeriesAvg200ResponseValuesInner {
	this := GetTimeSeriesAvg200ResponseValuesInner{}
	this.Datetime = datetime
	this.Avg = avg
	return &this
}

// NewGetTimeSeriesAvg200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAvg200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAvg200ResponseValuesInnerWithDefaults() *GetTimeSeriesAvg200ResponseValuesInner {
	this := GetTimeSeriesAvg200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAvg200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvg200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAvg200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAvg returns the Avg field value
func (o *GetTimeSeriesAvg200ResponseValuesInner) GetAvg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avg
}

// GetAvgOk returns a tuple with the Avg field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvg200ResponseValuesInner) GetAvgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avg, true
}

// SetAvg sets field value
func (o *GetTimeSeriesAvg200ResponseValuesInner) SetAvg(v string) {
	o.Avg = v
}

func (o GetTimeSeriesAvg200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAvg200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["avg"] = o.Avg
	return toSerialize, nil
}

func (o *GetTimeSeriesAvg200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"avg",
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

	varGetTimeSeriesAvg200ResponseValuesInner := _GetTimeSeriesAvg200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAvg200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAvg200ResponseValuesInner(varGetTimeSeriesAvg200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAvg200ResponseValuesInner struct {
	value *GetTimeSeriesAvg200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAvg200ResponseValuesInner) Get() *GetTimeSeriesAvg200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAvg200ResponseValuesInner) Set(val *GetTimeSeriesAvg200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAvg200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAvg200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAvg200ResponseValuesInner(val *GetTimeSeriesAvg200ResponseValuesInner) *NullableGetTimeSeriesAvg200ResponseValuesInner {
	return &NullableGetTimeSeriesAvg200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAvg200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAvg200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
