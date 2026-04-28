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

// checks if the GetTimeSeriesTema200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTema200ResponseValuesInner{}

// GetTimeSeriesTema200ResponseValuesInner struct for GetTimeSeriesTema200ResponseValuesInner
type GetTimeSeriesTema200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// TEMA value
	Tema string `json:"tema"`
}

type _GetTimeSeriesTema200ResponseValuesInner GetTimeSeriesTema200ResponseValuesInner

// NewGetTimeSeriesTema200ResponseValuesInner instantiates a new GetTimeSeriesTema200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTema200ResponseValuesInner(datetime string, tema string) *GetTimeSeriesTema200ResponseValuesInner {
	this := GetTimeSeriesTema200ResponseValuesInner{}
	this.Datetime = datetime
	this.Tema = tema
	return &this
}

// NewGetTimeSeriesTema200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTema200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTema200ResponseValuesInnerWithDefaults() *GetTimeSeriesTema200ResponseValuesInner {
	this := GetTimeSeriesTema200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesTema200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTema200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesTema200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTema returns the Tema field value
func (o *GetTimeSeriesTema200ResponseValuesInner) GetTema() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tema
}

// GetTemaOk returns a tuple with the Tema field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTema200ResponseValuesInner) GetTemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tema, true
}

// SetTema sets field value
func (o *GetTimeSeriesTema200ResponseValuesInner) SetTema(v string) {
	o.Tema = v
}

func (o GetTimeSeriesTema200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTema200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["tema"] = o.Tema
	return toSerialize, nil
}

func (o *GetTimeSeriesTema200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"tema",
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

	varGetTimeSeriesTema200ResponseValuesInner := _GetTimeSeriesTema200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesTema200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTema200ResponseValuesInner(varGetTimeSeriesTema200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesTema200ResponseValuesInner struct {
	value *GetTimeSeriesTema200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesTema200ResponseValuesInner) Get() *GetTimeSeriesTema200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesTema200ResponseValuesInner) Set(val *GetTimeSeriesTema200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTema200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTema200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTema200ResponseValuesInner(val *GetTimeSeriesTema200ResponseValuesInner) *NullableGetTimeSeriesTema200ResponseValuesInner {
	return &NullableGetTimeSeriesTema200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTema200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTema200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
