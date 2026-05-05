// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMidPrice200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMidPrice200ResponseValuesInner{}

// GetTimeSeriesMidPrice200ResponseValuesInner struct for GetTimeSeriesMidPrice200ResponseValuesInner
type GetTimeSeriesMidPrice200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Midprice value
	Midprice string `json:"midprice"`
}

type _GetTimeSeriesMidPrice200ResponseValuesInner GetTimeSeriesMidPrice200ResponseValuesInner

// NewGetTimeSeriesMidPrice200ResponseValuesInner instantiates a new GetTimeSeriesMidPrice200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMidPrice200ResponseValuesInner(datetime string, midprice string) *GetTimeSeriesMidPrice200ResponseValuesInner {
	this := GetTimeSeriesMidPrice200ResponseValuesInner{}
	this.Datetime = datetime
	this.Midprice = midprice
	return &this
}

// NewGetTimeSeriesMidPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMidPrice200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMidPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesMidPrice200ResponseValuesInner {
	this := GetTimeSeriesMidPrice200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMidPrice200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMidprice returns the Midprice field value
func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetMidprice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Midprice
}

// GetMidpriceOk returns a tuple with the Midprice field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMidPrice200ResponseValuesInner) GetMidpriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Midprice, true
}

// SetMidprice sets field value
func (o *GetTimeSeriesMidPrice200ResponseValuesInner) SetMidprice(v string) {
	o.Midprice = v
}

func (o GetTimeSeriesMidPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMidPrice200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["midprice"] = o.Midprice
	return toSerialize, nil
}

func (o *GetTimeSeriesMidPrice200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"midprice",
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

	varGetTimeSeriesMidPrice200ResponseValuesInner := _GetTimeSeriesMidPrice200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMidPrice200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMidPrice200ResponseValuesInner(varGetTimeSeriesMidPrice200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMidPrice200ResponseValuesInner struct {
	value *GetTimeSeriesMidPrice200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMidPrice200ResponseValuesInner) Get() *GetTimeSeriesMidPrice200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMidPrice200ResponseValuesInner) Set(val *GetTimeSeriesMidPrice200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMidPrice200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMidPrice200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMidPrice200ResponseValuesInner(val *GetTimeSeriesMidPrice200ResponseValuesInner) *NullableGetTimeSeriesMidPrice200ResponseValuesInner {
	return &NullableGetTimeSeriesMidPrice200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMidPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMidPrice200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
