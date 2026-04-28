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

// checks if the GetTimeSeriesHtDcPhase200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtDcPhase200ResponseValuesInner{}

// GetTimeSeriesHtDcPhase200ResponseValuesInner struct for GetTimeSeriesHtDcPhase200ResponseValuesInner
type GetTimeSeriesHtDcPhase200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// HT_DCPHASE value
	HtDcphase string `json:"ht_dcphase"`
}

type _GetTimeSeriesHtDcPhase200ResponseValuesInner GetTimeSeriesHtDcPhase200ResponseValuesInner

// NewGetTimeSeriesHtDcPhase200ResponseValuesInner instantiates a new GetTimeSeriesHtDcPhase200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtDcPhase200ResponseValuesInner(datetime string, htDcphase string) *GetTimeSeriesHtDcPhase200ResponseValuesInner {
	this := GetTimeSeriesHtDcPhase200ResponseValuesInner{}
	this.Datetime = datetime
	this.HtDcphase = htDcphase
	return &this
}

// NewGetTimeSeriesHtDcPhase200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtDcPhase200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtDcPhase200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtDcPhase200ResponseValuesInner {
	this := GetTimeSeriesHtDcPhase200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetHtDcphase returns the HtDcphase field value
func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetHtDcphase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtDcphase
}

// GetHtDcphaseOk returns a tuple with the HtDcphase field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetHtDcphaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtDcphase, true
}

// SetHtDcphase sets field value
func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) SetHtDcphase(v string) {
	o.HtDcphase = v
}

func (o GetTimeSeriesHtDcPhase200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtDcPhase200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ht_dcphase"] = o.HtDcphase
	return toSerialize, nil
}

func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ht_dcphase",
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

	varGetTimeSeriesHtDcPhase200ResponseValuesInner := _GetTimeSeriesHtDcPhase200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtDcPhase200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtDcPhase200ResponseValuesInner(varGetTimeSeriesHtDcPhase200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesHtDcPhase200ResponseValuesInner struct {
	value *GetTimeSeriesHtDcPhase200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesHtDcPhase200ResponseValuesInner) Get() *GetTimeSeriesHtDcPhase200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesHtDcPhase200ResponseValuesInner) Set(val *GetTimeSeriesHtDcPhase200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtDcPhase200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtDcPhase200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtDcPhase200ResponseValuesInner(val *GetTimeSeriesHtDcPhase200ResponseValuesInner) *NullableGetTimeSeriesHtDcPhase200ResponseValuesInner {
	return &NullableGetTimeSeriesHtDcPhase200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtDcPhase200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtDcPhase200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
