// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesPpo200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPpo200ResponseValuesInner{}

// GetTimeSeriesPpo200ResponseValuesInner struct for GetTimeSeriesPpo200ResponseValuesInner
type GetTimeSeriesPpo200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// PPO value
	Ppo string `json:"ppo"`
}

type _GetTimeSeriesPpo200ResponseValuesInner GetTimeSeriesPpo200ResponseValuesInner

// NewGetTimeSeriesPpo200ResponseValuesInner instantiates a new GetTimeSeriesPpo200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPpo200ResponseValuesInner(datetime string, ppo string) *GetTimeSeriesPpo200ResponseValuesInner {
	this := GetTimeSeriesPpo200ResponseValuesInner{}
	this.Datetime = datetime
	this.Ppo = ppo
	return &this
}

// NewGetTimeSeriesPpo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPpo200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPpo200ResponseValuesInnerWithDefaults() *GetTimeSeriesPpo200ResponseValuesInner {
	this := GetTimeSeriesPpo200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesPpo200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPpo200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesPpo200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetPpo returns the Ppo field value
func (o *GetTimeSeriesPpo200ResponseValuesInner) GetPpo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ppo
}

// GetPpoOk returns a tuple with the Ppo field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPpo200ResponseValuesInner) GetPpoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ppo, true
}

// SetPpo sets field value
func (o *GetTimeSeriesPpo200ResponseValuesInner) SetPpo(v string) {
	o.Ppo = v
}

func (o GetTimeSeriesPpo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPpo200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ppo"] = o.Ppo
	return toSerialize, nil
}

func (o *GetTimeSeriesPpo200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ppo",
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

	varGetTimeSeriesPpo200ResponseValuesInner := _GetTimeSeriesPpo200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesPpo200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPpo200ResponseValuesInner(varGetTimeSeriesPpo200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesPpo200ResponseValuesInner struct {
	value *GetTimeSeriesPpo200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesPpo200ResponseValuesInner) Get() *GetTimeSeriesPpo200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesPpo200ResponseValuesInner) Set(val *GetTimeSeriesPpo200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPpo200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPpo200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPpo200ResponseValuesInner(val *GetTimeSeriesPpo200ResponseValuesInner) *NullableGetTimeSeriesPpo200ResponseValuesInner {
	return &NullableGetTimeSeriesPpo200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPpo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPpo200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
