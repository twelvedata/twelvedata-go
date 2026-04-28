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

// checks if the GetTimeSeriesKama200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKama200ResponseValuesInner{}

// GetTimeSeriesKama200ResponseValuesInner struct for GetTimeSeriesKama200ResponseValuesInner
type GetTimeSeriesKama200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Kama value
	Kama string `json:"kama"`
}

type _GetTimeSeriesKama200ResponseValuesInner GetTimeSeriesKama200ResponseValuesInner

// NewGetTimeSeriesKama200ResponseValuesInner instantiates a new GetTimeSeriesKama200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKama200ResponseValuesInner(datetime string, kama string) *GetTimeSeriesKama200ResponseValuesInner {
	this := GetTimeSeriesKama200ResponseValuesInner{}
	this.Datetime = datetime
	this.Kama = kama
	return &this
}

// NewGetTimeSeriesKama200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesKama200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKama200ResponseValuesInnerWithDefaults() *GetTimeSeriesKama200ResponseValuesInner {
	this := GetTimeSeriesKama200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesKama200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesKama200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetKama returns the Kama field value
func (o *GetTimeSeriesKama200ResponseValuesInner) GetKama() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kama
}

// GetKamaOk returns a tuple with the Kama field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKama200ResponseValuesInner) GetKamaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kama, true
}

// SetKama sets field value
func (o *GetTimeSeriesKama200ResponseValuesInner) SetKama(v string) {
	o.Kama = v
}

func (o GetTimeSeriesKama200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKama200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["kama"] = o.Kama
	return toSerialize, nil
}

func (o *GetTimeSeriesKama200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"kama",
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

	varGetTimeSeriesKama200ResponseValuesInner := _GetTimeSeriesKama200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesKama200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKama200ResponseValuesInner(varGetTimeSeriesKama200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesKama200ResponseValuesInner struct {
	value *GetTimeSeriesKama200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesKama200ResponseValuesInner) Get() *GetTimeSeriesKama200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesKama200ResponseValuesInner) Set(val *GetTimeSeriesKama200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKama200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKama200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKama200ResponseValuesInner(val *GetTimeSeriesKama200ResponseValuesInner) *NullableGetTimeSeriesKama200ResponseValuesInner {
	return &NullableGetTimeSeriesKama200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKama200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKama200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
