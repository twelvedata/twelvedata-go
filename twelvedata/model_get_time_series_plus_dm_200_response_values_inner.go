// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesPlusDM200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPlusDM200ResponseValuesInner{}

// GetTimeSeriesPlusDM200ResponseValuesInner struct for GetTimeSeriesPlusDM200ResponseValuesInner
type GetTimeSeriesPlusDM200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// plus_dm value
	PlusDm string `json:"plus_dm"`
}

type _GetTimeSeriesPlusDM200ResponseValuesInner GetTimeSeriesPlusDM200ResponseValuesInner

// NewGetTimeSeriesPlusDM200ResponseValuesInner instantiates a new GetTimeSeriesPlusDM200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPlusDM200ResponseValuesInner(datetime string, plusDm string) *GetTimeSeriesPlusDM200ResponseValuesInner {
	this := GetTimeSeriesPlusDM200ResponseValuesInner{}
	this.Datetime = datetime
	this.PlusDm = plusDm
	return &this
}

// NewGetTimeSeriesPlusDM200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPlusDM200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPlusDM200ResponseValuesInnerWithDefaults() *GetTimeSeriesPlusDM200ResponseValuesInner {
	this := GetTimeSeriesPlusDM200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesPlusDM200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetPlusDm returns the PlusDm field value
func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetPlusDm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlusDm
}

// GetPlusDmOk returns a tuple with the PlusDm field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetPlusDmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlusDm, true
}

// SetPlusDm sets field value
func (o *GetTimeSeriesPlusDM200ResponseValuesInner) SetPlusDm(v string) {
	o.PlusDm = v
}

func (o GetTimeSeriesPlusDM200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPlusDM200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["plus_dm"] = o.PlusDm
	return toSerialize, nil
}

func (o *GetTimeSeriesPlusDM200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"plus_dm",
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

	varGetTimeSeriesPlusDM200ResponseValuesInner := _GetTimeSeriesPlusDM200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesPlusDM200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPlusDM200ResponseValuesInner(varGetTimeSeriesPlusDM200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesPlusDM200ResponseValuesInner struct {
	value *GetTimeSeriesPlusDM200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesPlusDM200ResponseValuesInner) Get() *GetTimeSeriesPlusDM200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesPlusDM200ResponseValuesInner) Set(val *GetTimeSeriesPlusDM200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPlusDM200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPlusDM200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPlusDM200ResponseValuesInner(val *GetTimeSeriesPlusDM200ResponseValuesInner) *NullableGetTimeSeriesPlusDM200ResponseValuesInner {
	return &NullableGetTimeSeriesPlusDM200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPlusDM200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPlusDM200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
