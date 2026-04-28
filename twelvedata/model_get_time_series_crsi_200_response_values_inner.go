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

// checks if the GetTimeSeriesCrsi200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCrsi200ResponseValuesInner{}

// GetTimeSeriesCrsi200ResponseValuesInner struct for GetTimeSeriesCrsi200ResponseValuesInner
type GetTimeSeriesCrsi200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// crsi value
	Crsi string `json:"crsi"`
}

type _GetTimeSeriesCrsi200ResponseValuesInner GetTimeSeriesCrsi200ResponseValuesInner

// NewGetTimeSeriesCrsi200ResponseValuesInner instantiates a new GetTimeSeriesCrsi200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCrsi200ResponseValuesInner(datetime string, crsi string) *GetTimeSeriesCrsi200ResponseValuesInner {
	this := GetTimeSeriesCrsi200ResponseValuesInner{}
	this.Datetime = datetime
	this.Crsi = crsi
	return &this
}

// NewGetTimeSeriesCrsi200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCrsi200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCrsi200ResponseValuesInnerWithDefaults() *GetTimeSeriesCrsi200ResponseValuesInner {
	this := GetTimeSeriesCrsi200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesCrsi200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCrsi returns the Crsi field value
func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetCrsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Crsi
}

// GetCrsiOk returns a tuple with the Crsi field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetCrsiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crsi, true
}

// SetCrsi sets field value
func (o *GetTimeSeriesCrsi200ResponseValuesInner) SetCrsi(v string) {
	o.Crsi = v
}

func (o GetTimeSeriesCrsi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCrsi200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["crsi"] = o.Crsi
	return toSerialize, nil
}

func (o *GetTimeSeriesCrsi200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"crsi",
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

	varGetTimeSeriesCrsi200ResponseValuesInner := _GetTimeSeriesCrsi200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesCrsi200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCrsi200ResponseValuesInner(varGetTimeSeriesCrsi200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesCrsi200ResponseValuesInner struct {
	value *GetTimeSeriesCrsi200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesCrsi200ResponseValuesInner) Get() *GetTimeSeriesCrsi200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesCrsi200ResponseValuesInner) Set(val *GetTimeSeriesCrsi200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCrsi200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCrsi200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCrsi200ResponseValuesInner(val *GetTimeSeriesCrsi200ResponseValuesInner) *NullableGetTimeSeriesCrsi200ResponseValuesInner {
	return &NullableGetTimeSeriesCrsi200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCrsi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCrsi200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
