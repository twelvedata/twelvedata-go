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

// checks if the GetTimeSeriesHtDcPeriod200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesHtDcPeriod200ResponseValuesInner{}

// GetTimeSeriesHtDcPeriod200ResponseValuesInner struct for GetTimeSeriesHtDcPeriod200ResponseValuesInner
type GetTimeSeriesHtDcPeriod200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ht_dcperiod value
	HtDcperiod string `json:"ht_dcperiod"`
}

type _GetTimeSeriesHtDcPeriod200ResponseValuesInner GetTimeSeriesHtDcPeriod200ResponseValuesInner

// NewGetTimeSeriesHtDcPeriod200ResponseValuesInner instantiates a new GetTimeSeriesHtDcPeriod200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesHtDcPeriod200ResponseValuesInner(datetime string, htDcperiod string) *GetTimeSeriesHtDcPeriod200ResponseValuesInner {
	this := GetTimeSeriesHtDcPeriod200ResponseValuesInner{}
	this.Datetime = datetime
	this.HtDcperiod = htDcperiod
	return &this
}

// NewGetTimeSeriesHtDcPeriod200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtDcPeriod200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesHtDcPeriod200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtDcPeriod200ResponseValuesInner {
	this := GetTimeSeriesHtDcPeriod200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetHtDcperiod returns the HtDcperiod field value
func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetHtDcperiod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtDcperiod
}

// GetHtDcperiodOk returns a tuple with the HtDcperiod field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetHtDcperiodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtDcperiod, true
}

// SetHtDcperiod sets field value
func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) SetHtDcperiod(v string) {
	o.HtDcperiod = v
}

func (o GetTimeSeriesHtDcPeriod200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesHtDcPeriod200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ht_dcperiod"] = o.HtDcperiod
	return toSerialize, nil
}

func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ht_dcperiod",
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

	varGetTimeSeriesHtDcPeriod200ResponseValuesInner := _GetTimeSeriesHtDcPeriod200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesHtDcPeriod200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesHtDcPeriod200ResponseValuesInner(varGetTimeSeriesHtDcPeriod200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner struct {
	value *GetTimeSeriesHtDcPeriod200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner) Get() *GetTimeSeriesHtDcPeriod200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner) Set(val *GetTimeSeriesHtDcPeriod200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesHtDcPeriod200ResponseValuesInner(val *GetTimeSeriesHtDcPeriod200ResponseValuesInner) *NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner {
	return &NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesHtDcPeriod200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
