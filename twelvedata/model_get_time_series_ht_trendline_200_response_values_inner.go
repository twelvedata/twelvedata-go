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

// checks if the GetTimeSeriesHtTrendline200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtTrendline200ResponseValuesInner{}

// GetTimeSeriesHtTrendline200ResponseValuesInner struct for GetTimeSeriesHtTrendline200ResponseValuesInner
type GetTimeSeriesHtTrendline200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// HT_TRENDLINE value
	HtTrendline string `json:"ht_trendline"`
}

type _GetTimeSeriesHtTrendline200ResponseValuesInner GetTimeSeriesHtTrendline200ResponseValuesInner

// NewGetTimeSeriesHtTrendline200ResponseValuesInner instantiates a new GetTimeSeriesHtTrendline200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtTrendline200ResponseValuesInner(datetime string, htTrendline string) *GetTimeSeriesHtTrendline200ResponseValuesInner {
	this := GetTimeSeriesHtTrendline200ResponseValuesInner{}
	this.Datetime = datetime
	this.HtTrendline = htTrendline
	return &this
}

// NewGetTimeSeriesHtTrendline200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtTrendline200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtTrendline200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtTrendline200ResponseValuesInner {
	this := GetTimeSeriesHtTrendline200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetHtTrendline returns the HtTrendline field value
func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetHtTrendline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtTrendline
}

// GetHtTrendlineOk returns a tuple with the HtTrendline field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) GetHtTrendlineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtTrendline, true
}

// SetHtTrendline sets field value
func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) SetHtTrendline(v string) {
	o.HtTrendline = v
}

func (o GetTimeSeriesHtTrendline200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtTrendline200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ht_trendline"] = o.HtTrendline
	return toSerialize, nil
}

func (o *GetTimeSeriesHtTrendline200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ht_trendline",
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

	varGetTimeSeriesHtTrendline200ResponseValuesInner := _GetTimeSeriesHtTrendline200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtTrendline200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtTrendline200ResponseValuesInner(varGetTimeSeriesHtTrendline200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesHtTrendline200ResponseValuesInner struct {
	value *GetTimeSeriesHtTrendline200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesHtTrendline200ResponseValuesInner) Get() *GetTimeSeriesHtTrendline200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesHtTrendline200ResponseValuesInner) Set(val *GetTimeSeriesHtTrendline200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtTrendline200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtTrendline200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtTrendline200ResponseValuesInner(val *GetTimeSeriesHtTrendline200ResponseValuesInner) *NullableGetTimeSeriesHtTrendline200ResponseValuesInner {
	return &NullableGetTimeSeriesHtTrendline200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtTrendline200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtTrendline200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
