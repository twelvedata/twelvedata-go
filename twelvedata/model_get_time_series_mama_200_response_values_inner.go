// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMama200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMama200ResponseValuesInner{}

// GetTimeSeriesMama200ResponseValuesInner struct for GetTimeSeriesMama200ResponseValuesInner
type GetTimeSeriesMama200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// MAMA value
	Mama string `json:"mama"`
	// FAMA value
	Fama string `json:"fama"`
}

type _GetTimeSeriesMama200ResponseValuesInner GetTimeSeriesMama200ResponseValuesInner

// NewGetTimeSeriesMama200ResponseValuesInner instantiates a new GetTimeSeriesMama200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMama200ResponseValuesInner(datetime string, mama string, fama string) *GetTimeSeriesMama200ResponseValuesInner {
	this := GetTimeSeriesMama200ResponseValuesInner{}
	this.Datetime = datetime
	this.Mama = mama
	this.Fama = fama
	return &this
}

// NewGetTimeSeriesMama200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMama200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMama200ResponseValuesInnerWithDefaults() *GetTimeSeriesMama200ResponseValuesInner {
	this := GetTimeSeriesMama200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMama200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMama200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMama200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMama returns the Mama field value
func (o *GetTimeSeriesMama200ResponseValuesInner) GetMama() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mama
}

// GetMamaOk returns a tuple with the Mama field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMama200ResponseValuesInner) GetMamaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mama, true
}

// SetMama sets field value
func (o *GetTimeSeriesMama200ResponseValuesInner) SetMama(v string) {
	o.Mama = v
}

// GetFama returns the Fama field value
func (o *GetTimeSeriesMama200ResponseValuesInner) GetFama() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fama
}

// GetFamaOk returns a tuple with the Fama field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMama200ResponseValuesInner) GetFamaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fama, true
}

// SetFama sets field value
func (o *GetTimeSeriesMama200ResponseValuesInner) SetFama(v string) {
	o.Fama = v
}

func (o GetTimeSeriesMama200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMama200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["mama"] = o.Mama
	toSerialize["fama"] = o.Fama
	return toSerialize, nil
}

func (o *GetTimeSeriesMama200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"mama",
		"fama",
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

	varGetTimeSeriesMama200ResponseValuesInner := _GetTimeSeriesMama200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMama200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMama200ResponseValuesInner(varGetTimeSeriesMama200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMama200ResponseValuesInner struct {
	value *GetTimeSeriesMama200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMama200ResponseValuesInner) Get() *GetTimeSeriesMama200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMama200ResponseValuesInner) Set(val *GetTimeSeriesMama200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMama200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMama200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMama200ResponseValuesInner(val *GetTimeSeriesMama200ResponseValuesInner) *NullableGetTimeSeriesMama200ResponseValuesInner {
	return &NullableGetTimeSeriesMama200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMama200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMama200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
