// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesHtTrendMode200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtTrendMode200ResponseValuesInner{}

// GetTimeSeriesHtTrendMode200ResponseValuesInner struct for GetTimeSeriesHtTrendMode200ResponseValuesInner
type GetTimeSeriesHtTrendMode200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ht_trendmode value
	HtTrendmode string `json:"ht_trendmode"`
}

type _GetTimeSeriesHtTrendMode200ResponseValuesInner GetTimeSeriesHtTrendMode200ResponseValuesInner

// NewGetTimeSeriesHtTrendMode200ResponseValuesInner instantiates a new GetTimeSeriesHtTrendMode200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtTrendMode200ResponseValuesInner(datetime string, htTrendmode string) *GetTimeSeriesHtTrendMode200ResponseValuesInner {
	this := GetTimeSeriesHtTrendMode200ResponseValuesInner{}
	this.Datetime = datetime
	this.HtTrendmode = htTrendmode
	return &this
}

// NewGetTimeSeriesHtTrendMode200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtTrendMode200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtTrendMode200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtTrendMode200ResponseValuesInner {
	this := GetTimeSeriesHtTrendMode200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetHtTrendmode returns the HtTrendmode field value
func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetHtTrendmode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtTrendmode
}

// GetHtTrendmodeOk returns a tuple with the HtTrendmode field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetHtTrendmodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtTrendmode, true
}

// SetHtTrendmode sets field value
func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) SetHtTrendmode(v string) {
	o.HtTrendmode = v
}

func (o GetTimeSeriesHtTrendMode200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtTrendMode200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ht_trendmode"] = o.HtTrendmode
	return toSerialize, nil
}

func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ht_trendmode",
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

	varGetTimeSeriesHtTrendMode200ResponseValuesInner := _GetTimeSeriesHtTrendMode200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesHtTrendMode200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtTrendMode200ResponseValuesInner(varGetTimeSeriesHtTrendMode200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesHtTrendMode200ResponseValuesInner struct {
	value *GetTimeSeriesHtTrendMode200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesHtTrendMode200ResponseValuesInner) Get() *GetTimeSeriesHtTrendMode200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesHtTrendMode200ResponseValuesInner) Set(val *GetTimeSeriesHtTrendMode200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtTrendMode200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtTrendMode200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtTrendMode200ResponseValuesInner(val *GetTimeSeriesHtTrendMode200ResponseValuesInner) *NullableGetTimeSeriesHtTrendMode200ResponseValuesInner {
	return &NullableGetTimeSeriesHtTrendMode200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtTrendMode200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtTrendMode200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
