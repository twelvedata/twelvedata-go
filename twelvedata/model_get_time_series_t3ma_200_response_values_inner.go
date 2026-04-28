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

// checks if the GetTimeSeriesT3ma200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesT3ma200ResponseValuesInner{}

// GetTimeSeriesT3ma200ResponseValuesInner struct for GetTimeSeriesT3ma200ResponseValuesInner
type GetTimeSeriesT3ma200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// T3MA value
	T3ma string `json:"t3ma"`
}

type _GetTimeSeriesT3ma200ResponseValuesInner GetTimeSeriesT3ma200ResponseValuesInner

// NewGetTimeSeriesT3ma200ResponseValuesInner instantiates a new GetTimeSeriesT3ma200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesT3ma200ResponseValuesInner(datetime string, t3ma string) *GetTimeSeriesT3ma200ResponseValuesInner {
	this := GetTimeSeriesT3ma200ResponseValuesInner{}
	this.Datetime = datetime
	this.T3ma = t3ma
	return &this
}

// NewGetTimeSeriesT3ma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesT3ma200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesT3ma200ResponseValuesInnerWithDefaults() *GetTimeSeriesT3ma200ResponseValuesInner {
	this := GetTimeSeriesT3ma200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesT3ma200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetT3ma returns the T3ma field value
func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetT3ma() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T3ma
}

// GetT3maOk returns a tuple with the T3ma field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesT3ma200ResponseValuesInner) GetT3maOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T3ma, true
}

// SetT3ma sets field value
func (o *GetTimeSeriesT3ma200ResponseValuesInner) SetT3ma(v string) {
	o.T3ma = v
}

func (o GetTimeSeriesT3ma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesT3ma200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["t3ma"] = o.T3ma
	return toSerialize, nil
}

func (o *GetTimeSeriesT3ma200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"t3ma",
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

	varGetTimeSeriesT3ma200ResponseValuesInner := _GetTimeSeriesT3ma200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesT3ma200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesT3ma200ResponseValuesInner(varGetTimeSeriesT3ma200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesT3ma200ResponseValuesInner struct {
	value *GetTimeSeriesT3ma200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesT3ma200ResponseValuesInner) Get() *GetTimeSeriesT3ma200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesT3ma200ResponseValuesInner) Set(val *GetTimeSeriesT3ma200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesT3ma200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesT3ma200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesT3ma200ResponseValuesInner(val *GetTimeSeriesT3ma200ResponseValuesInner) *NullableGetTimeSeriesT3ma200ResponseValuesInner {
	return &NullableGetTimeSeriesT3ma200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesT3ma200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesT3ma200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
