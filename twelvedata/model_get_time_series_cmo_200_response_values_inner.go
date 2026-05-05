// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesCmo200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCmo200ResponseValuesInner{}

// GetTimeSeriesCmo200ResponseValuesInner struct for GetTimeSeriesCmo200ResponseValuesInner
type GetTimeSeriesCmo200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// CMO value
	Cmo string `json:"cmo"`
}

type _GetTimeSeriesCmo200ResponseValuesInner GetTimeSeriesCmo200ResponseValuesInner

// NewGetTimeSeriesCmo200ResponseValuesInner instantiates a new GetTimeSeriesCmo200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCmo200ResponseValuesInner(datetime string, cmo string) *GetTimeSeriesCmo200ResponseValuesInner {
	this := GetTimeSeriesCmo200ResponseValuesInner{}
	this.Datetime = datetime
	this.Cmo = cmo
	return &this
}

// NewGetTimeSeriesCmo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCmo200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCmo200ResponseValuesInnerWithDefaults() *GetTimeSeriesCmo200ResponseValuesInner {
	this := GetTimeSeriesCmo200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesCmo200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCmo200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesCmo200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCmo returns the Cmo field value
func (o *GetTimeSeriesCmo200ResponseValuesInner) GetCmo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cmo
}

// GetCmoOk returns a tuple with the Cmo field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCmo200ResponseValuesInner) GetCmoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cmo, true
}

// SetCmo sets field value
func (o *GetTimeSeriesCmo200ResponseValuesInner) SetCmo(v string) {
	o.Cmo = v
}

func (o GetTimeSeriesCmo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCmo200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["cmo"] = o.Cmo
	return toSerialize, nil
}

func (o *GetTimeSeriesCmo200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"cmo",
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

	varGetTimeSeriesCmo200ResponseValuesInner := _GetTimeSeriesCmo200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCmo200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCmo200ResponseValuesInner(varGetTimeSeriesCmo200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesCmo200ResponseValuesInner struct {
	value *GetTimeSeriesCmo200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesCmo200ResponseValuesInner) Get() *GetTimeSeriesCmo200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesCmo200ResponseValuesInner) Set(val *GetTimeSeriesCmo200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCmo200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCmo200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCmo200ResponseValuesInner(val *GetTimeSeriesCmo200ResponseValuesInner) *NullableGetTimeSeriesCmo200ResponseValuesInner {
	return &NullableGetTimeSeriesCmo200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCmo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCmo200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
