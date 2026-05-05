// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesTRange200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTRange200ResponseValuesInner{}

// GetTimeSeriesTRange200ResponseValuesInner struct for GetTimeSeriesTRange200ResponseValuesInner
type GetTimeSeriesTRange200ResponseValuesInner struct {
	// datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// trange value
	Trange string `json:"trange"`
}

type _GetTimeSeriesTRange200ResponseValuesInner GetTimeSeriesTRange200ResponseValuesInner

// NewGetTimeSeriesTRange200ResponseValuesInner instantiates a new GetTimeSeriesTRange200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTRange200ResponseValuesInner(datetime string, trange string) *GetTimeSeriesTRange200ResponseValuesInner {
	this := GetTimeSeriesTRange200ResponseValuesInner{}
	this.Datetime = datetime
	this.Trange = trange
	return &this
}

// NewGetTimeSeriesTRange200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTRange200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTRange200ResponseValuesInnerWithDefaults() *GetTimeSeriesTRange200ResponseValuesInner {
	this := GetTimeSeriesTRange200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesTRange200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTRange200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesTRange200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTrange returns the Trange field value
func (o *GetTimeSeriesTRange200ResponseValuesInner) GetTrange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trange
}

// GetTrangeOk returns a tuple with the Trange field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTRange200ResponseValuesInner) GetTrangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trange, true
}

// SetTrange sets field value
func (o *GetTimeSeriesTRange200ResponseValuesInner) SetTrange(v string) {
	o.Trange = v
}

func (o GetTimeSeriesTRange200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTRange200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["trange"] = o.Trange
	return toSerialize, nil
}

func (o *GetTimeSeriesTRange200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"trange",
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

	varGetTimeSeriesTRange200ResponseValuesInner := _GetTimeSeriesTRange200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTRange200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTRange200ResponseValuesInner(varGetTimeSeriesTRange200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesTRange200ResponseValuesInner struct {
	value *GetTimeSeriesTRange200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesTRange200ResponseValuesInner) Get() *GetTimeSeriesTRange200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesTRange200ResponseValuesInner) Set(val *GetTimeSeriesTRange200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTRange200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTRange200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTRange200ResponseValuesInner(val *GetTimeSeriesTRange200ResponseValuesInner) *NullableGetTimeSeriesTRange200ResponseValuesInner {
	return &NullableGetTimeSeriesTRange200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTRange200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTRange200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
