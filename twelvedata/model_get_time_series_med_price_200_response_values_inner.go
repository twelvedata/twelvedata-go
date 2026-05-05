// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMedPrice200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMedPrice200ResponseValuesInner{}

// GetTimeSeriesMedPrice200ResponseValuesInner struct for GetTimeSeriesMedPrice200ResponseValuesInner
type GetTimeSeriesMedPrice200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Medprice value
	Medprice string `json:"medprice"`
}

type _GetTimeSeriesMedPrice200ResponseValuesInner GetTimeSeriesMedPrice200ResponseValuesInner

// NewGetTimeSeriesMedPrice200ResponseValuesInner instantiates a new GetTimeSeriesMedPrice200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMedPrice200ResponseValuesInner(datetime string, medprice string) *GetTimeSeriesMedPrice200ResponseValuesInner {
	this := GetTimeSeriesMedPrice200ResponseValuesInner{}
	this.Datetime = datetime
	this.Medprice = medprice
	return &this
}

// NewGetTimeSeriesMedPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMedPrice200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMedPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesMedPrice200ResponseValuesInner {
	this := GetTimeSeriesMedPrice200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMedPrice200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMedprice returns the Medprice field value
func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetMedprice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Medprice
}

// GetMedpriceOk returns a tuple with the Medprice field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMedPrice200ResponseValuesInner) GetMedpriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Medprice, true
}

// SetMedprice sets field value
func (o *GetTimeSeriesMedPrice200ResponseValuesInner) SetMedprice(v string) {
	o.Medprice = v
}

func (o GetTimeSeriesMedPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMedPrice200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["medprice"] = o.Medprice
	return toSerialize, nil
}

func (o *GetTimeSeriesMedPrice200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"medprice",
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

	varGetTimeSeriesMedPrice200ResponseValuesInner := _GetTimeSeriesMedPrice200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMedPrice200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMedPrice200ResponseValuesInner(varGetTimeSeriesMedPrice200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMedPrice200ResponseValuesInner struct {
	value *GetTimeSeriesMedPrice200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMedPrice200ResponseValuesInner) Get() *GetTimeSeriesMedPrice200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMedPrice200ResponseValuesInner) Set(val *GetTimeSeriesMedPrice200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMedPrice200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMedPrice200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMedPrice200ResponseValuesInner(val *GetTimeSeriesMedPrice200ResponseValuesInner) *NullableGetTimeSeriesMedPrice200ResponseValuesInner {
	return &NullableGetTimeSeriesMedPrice200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMedPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMedPrice200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
