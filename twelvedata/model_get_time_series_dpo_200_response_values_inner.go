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

// checks if the GetTimeSeriesDpo200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesDpo200ResponseValuesInner{}

// GetTimeSeriesDpo200ResponseValuesInner struct for GetTimeSeriesDpo200ResponseValuesInner
type GetTimeSeriesDpo200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// DPO value
	Dpo string `json:"dpo"`
}

type _GetTimeSeriesDpo200ResponseValuesInner GetTimeSeriesDpo200ResponseValuesInner

// NewGetTimeSeriesDpo200ResponseValuesInner instantiates a new GetTimeSeriesDpo200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesDpo200ResponseValuesInner(datetime string, dpo string) *GetTimeSeriesDpo200ResponseValuesInner {
	this := GetTimeSeriesDpo200ResponseValuesInner{}
	this.Datetime = datetime
	this.Dpo = dpo
	return &this
}

// NewGetTimeSeriesDpo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesDpo200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesDpo200ResponseValuesInnerWithDefaults() *GetTimeSeriesDpo200ResponseValuesInner {
	this := GetTimeSeriesDpo200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesDpo200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetDpo returns the Dpo field value
func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDpo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dpo
}

// GetDpoOk returns a tuple with the Dpo field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDpoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dpo, true
}

// SetDpo sets field value
func (o *GetTimeSeriesDpo200ResponseValuesInner) SetDpo(v string) {
	o.Dpo = v
}

func (o GetTimeSeriesDpo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesDpo200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["dpo"] = o.Dpo
	return toSerialize, nil
}

func (o *GetTimeSeriesDpo200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"dpo",
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

	varGetTimeSeriesDpo200ResponseValuesInner := _GetTimeSeriesDpo200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesDpo200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesDpo200ResponseValuesInner(varGetTimeSeriesDpo200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesDpo200ResponseValuesInner struct {
	value *GetTimeSeriesDpo200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesDpo200ResponseValuesInner) Get() *GetTimeSeriesDpo200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesDpo200ResponseValuesInner) Set(val *GetTimeSeriesDpo200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesDpo200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesDpo200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesDpo200ResponseValuesInner(val *GetTimeSeriesDpo200ResponseValuesInner) *NullableGetTimeSeriesDpo200ResponseValuesInner {
	return &NullableGetTimeSeriesDpo200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesDpo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesDpo200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
