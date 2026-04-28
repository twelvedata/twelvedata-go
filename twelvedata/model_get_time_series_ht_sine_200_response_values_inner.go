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

// checks if the GetTimeSeriesHtSine200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtSine200ResponseValuesInner{}

// GetTimeSeriesHtSine200ResponseValuesInner struct for GetTimeSeriesHtSine200ResponseValuesInner
type GetTimeSeriesHtSine200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ht_sine value
	HtSine string `json:"ht_sine"`
	// ht_leadsine value
	HtLeadsine string `json:"ht_leadsine"`
}

type _GetTimeSeriesHtSine200ResponseValuesInner GetTimeSeriesHtSine200ResponseValuesInner

// NewGetTimeSeriesHtSine200ResponseValuesInner instantiates a new GetTimeSeriesHtSine200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtSine200ResponseValuesInner(datetime string, htSine string, htLeadsine string) *GetTimeSeriesHtSine200ResponseValuesInner {
	this := GetTimeSeriesHtSine200ResponseValuesInner{}
	this.Datetime = datetime
	this.HtSine = htSine
	this.HtLeadsine = htLeadsine
	return &this
}

// NewGetTimeSeriesHtSine200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtSine200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtSine200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtSine200ResponseValuesInner {
	this := GetTimeSeriesHtSine200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesHtSine200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetHtSine returns the HtSine field value
func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtSine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtSine
}

// GetHtSineOk returns a tuple with the HtSine field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtSineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtSine, true
}

// SetHtSine sets field value
func (o *GetTimeSeriesHtSine200ResponseValuesInner) SetHtSine(v string) {
	o.HtSine = v
}

// GetHtLeadsine returns the HtLeadsine field value
func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtLeadsine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtLeadsine
}

// GetHtLeadsineOk returns a tuple with the HtLeadsine field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtLeadsineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtLeadsine, true
}

// SetHtLeadsine sets field value
func (o *GetTimeSeriesHtSine200ResponseValuesInner) SetHtLeadsine(v string) {
	o.HtLeadsine = v
}

func (o GetTimeSeriesHtSine200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtSine200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ht_sine"] = o.HtSine
	toSerialize["ht_leadsine"] = o.HtLeadsine
	return toSerialize, nil
}

func (o *GetTimeSeriesHtSine200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ht_sine",
		"ht_leadsine",
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

	varGetTimeSeriesHtSine200ResponseValuesInner := _GetTimeSeriesHtSine200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtSine200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtSine200ResponseValuesInner(varGetTimeSeriesHtSine200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesHtSine200ResponseValuesInner struct {
	value *GetTimeSeriesHtSine200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesHtSine200ResponseValuesInner) Get() *GetTimeSeriesHtSine200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesHtSine200ResponseValuesInner) Set(val *GetTimeSeriesHtSine200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtSine200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtSine200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtSine200ResponseValuesInner(val *GetTimeSeriesHtSine200ResponseValuesInner) *NullableGetTimeSeriesHtSine200ResponseValuesInner {
	return &NullableGetTimeSeriesHtSine200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtSine200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtSine200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
