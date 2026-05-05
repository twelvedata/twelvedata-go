// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesTypPrice200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesTypPrice200ResponseValuesInner{}

// GetTimeSeriesTypPrice200ResponseValuesInner struct for GetTimeSeriesTypPrice200ResponseValuesInner
type GetTimeSeriesTypPrice200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// typprice value
	Typprice string `json:"typprice"`
}

type _GetTimeSeriesTypPrice200ResponseValuesInner GetTimeSeriesTypPrice200ResponseValuesInner

// NewGetTimeSeriesTypPrice200ResponseValuesInner instantiates a new GetTimeSeriesTypPrice200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesTypPrice200ResponseValuesInner(datetime string, typprice string) *GetTimeSeriesTypPrice200ResponseValuesInner {
	this := GetTimeSeriesTypPrice200ResponseValuesInner{}
	this.Datetime = datetime
	this.Typprice = typprice
	return &this
}

// NewGetTimeSeriesTypPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTypPrice200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesTypPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesTypPrice200ResponseValuesInner {
	this := GetTimeSeriesTypPrice200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesTypPrice200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetTypprice returns the Typprice field value
func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetTypprice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Typprice
}

// GetTyppriceOk returns a tuple with the Typprice field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetTyppriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typprice, true
}

// SetTypprice sets field value
func (o *GetTimeSeriesTypPrice200ResponseValuesInner) SetTypprice(v string) {
	o.Typprice = v
}

func (o GetTimeSeriesTypPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesTypPrice200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["typprice"] = o.Typprice
	return toSerialize, nil
}

func (o *GetTimeSeriesTypPrice200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"typprice",
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

	varGetTimeSeriesTypPrice200ResponseValuesInner := _GetTimeSeriesTypPrice200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesTypPrice200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesTypPrice200ResponseValuesInner(varGetTimeSeriesTypPrice200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesTypPrice200ResponseValuesInner struct {
	value *GetTimeSeriesTypPrice200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesTypPrice200ResponseValuesInner) Get() *GetTimeSeriesTypPrice200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesTypPrice200ResponseValuesInner) Set(val *GetTimeSeriesTypPrice200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesTypPrice200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesTypPrice200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesTypPrice200ResponseValuesInner(val *GetTimeSeriesTypPrice200ResponseValuesInner) *NullableGetTimeSeriesTypPrice200ResponseValuesInner {
	return &NullableGetTimeSeriesTypPrice200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesTypPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesTypPrice200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
