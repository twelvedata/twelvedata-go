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

// checks if the GetTimeSeriesSub200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSub200ResponseValuesInner{}

// GetTimeSeriesSub200ResponseValuesInner struct for GetTimeSeriesSub200ResponseValuesInner
type GetTimeSeriesSub200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SUB value
	Sub string `json:"sub"`
}

type _GetTimeSeriesSub200ResponseValuesInner GetTimeSeriesSub200ResponseValuesInner

// NewGetTimeSeriesSub200ResponseValuesInner instantiates a new GetTimeSeriesSub200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSub200ResponseValuesInner(datetime string, sub string) *GetTimeSeriesSub200ResponseValuesInner {
	this := GetTimeSeriesSub200ResponseValuesInner{}
	this.Datetime = datetime
	this.Sub = sub
	return &this
}

// NewGetTimeSeriesSub200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSub200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSub200ResponseValuesInnerWithDefaults() *GetTimeSeriesSub200ResponseValuesInner {
	this := GetTimeSeriesSub200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSub200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSub200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSub200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSub returns the Sub field value
func (o *GetTimeSeriesSub200ResponseValuesInner) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSub200ResponseValuesInner) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *GetTimeSeriesSub200ResponseValuesInner) SetSub(v string) {
	o.Sub = v
}

func (o GetTimeSeriesSub200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSub200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sub"] = o.Sub
	return toSerialize, nil
}

func (o *GetTimeSeriesSub200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sub",
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

	varGetTimeSeriesSub200ResponseValuesInner := _GetTimeSeriesSub200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSub200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSub200ResponseValuesInner(varGetTimeSeriesSub200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSub200ResponseValuesInner struct {
	value *GetTimeSeriesSub200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSub200ResponseValuesInner) Get() *GetTimeSeriesSub200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSub200ResponseValuesInner) Set(val *GetTimeSeriesSub200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSub200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSub200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSub200ResponseValuesInner(val *GetTimeSeriesSub200ResponseValuesInner) *NullableGetTimeSeriesSub200ResponseValuesInner {
	return &NullableGetTimeSeriesSub200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSub200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSub200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
