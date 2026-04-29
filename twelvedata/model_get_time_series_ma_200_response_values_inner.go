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

// checks if the GetTimeSeriesMa200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMa200ResponseValuesInner{}

// GetTimeSeriesMa200ResponseValuesInner struct for GetTimeSeriesMa200ResponseValuesInner
type GetTimeSeriesMa200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// MA value
	Ma string `json:"ma"`
}

type _GetTimeSeriesMa200ResponseValuesInner GetTimeSeriesMa200ResponseValuesInner

// NewGetTimeSeriesMa200ResponseValuesInner instantiates a new GetTimeSeriesMa200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMa200ResponseValuesInner(datetime string, ma string) *GetTimeSeriesMa200ResponseValuesInner {
	this := GetTimeSeriesMa200ResponseValuesInner{}
	this.Datetime = datetime
	this.Ma = ma
	return &this
}

// NewGetTimeSeriesMa200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMa200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMa200ResponseValuesInnerWithDefaults() *GetTimeSeriesMa200ResponseValuesInner {
	this := GetTimeSeriesMa200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMa200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMa200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMa returns the Ma field value
func (o *GetTimeSeriesMa200ResponseValuesInner) GetMa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ma
}

// GetMaOk returns a tuple with the Ma field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMa200ResponseValuesInner) GetMaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ma, true
}

// SetMa sets field value
func (o *GetTimeSeriesMa200ResponseValuesInner) SetMa(v string) {
	o.Ma = v
}

func (o GetTimeSeriesMa200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMa200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ma"] = o.Ma
	return toSerialize, nil
}

func (o *GetTimeSeriesMa200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ma",
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

	varGetTimeSeriesMa200ResponseValuesInner := _GetTimeSeriesMa200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMa200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMa200ResponseValuesInner(varGetTimeSeriesMa200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMa200ResponseValuesInner struct {
	value *GetTimeSeriesMa200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMa200ResponseValuesInner) Get() *GetTimeSeriesMa200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMa200ResponseValuesInner) Set(val *GetTimeSeriesMa200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMa200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMa200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMa200ResponseValuesInner(val *GetTimeSeriesMa200ResponseValuesInner) *NullableGetTimeSeriesMa200ResponseValuesInner {
	return &NullableGetTimeSeriesMa200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMa200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMa200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
