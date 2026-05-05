// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAvgPrice200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAvgPrice200ResponseValuesInner{}

// GetTimeSeriesAvgPrice200ResponseValuesInner struct for GetTimeSeriesAvgPrice200ResponseValuesInner
type GetTimeSeriesAvgPrice200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Avgprice value
	Avgprice string `json:"avgprice"`
}

type _GetTimeSeriesAvgPrice200ResponseValuesInner GetTimeSeriesAvgPrice200ResponseValuesInner

// NewGetTimeSeriesAvgPrice200ResponseValuesInner instantiates a new GetTimeSeriesAvgPrice200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAvgPrice200ResponseValuesInner(datetime string, avgprice string) *GetTimeSeriesAvgPrice200ResponseValuesInner {
	this := GetTimeSeriesAvgPrice200ResponseValuesInner{}
	this.Datetime = datetime
	this.Avgprice = avgprice
	return &this
}

// NewGetTimeSeriesAvgPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAvgPrice200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAvgPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesAvgPrice200ResponseValuesInner {
	this := GetTimeSeriesAvgPrice200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAvgprice returns the Avgprice field value
func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetAvgprice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avgprice
}

// GetAvgpriceOk returns a tuple with the Avgprice field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetAvgpriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avgprice, true
}

// SetAvgprice sets field value
func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) SetAvgprice(v string) {
	o.Avgprice = v
}

func (o GetTimeSeriesAvgPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAvgPrice200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["avgprice"] = o.Avgprice
	return toSerialize, nil
}

func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"avgprice",
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

	varGetTimeSeriesAvgPrice200ResponseValuesInner := _GetTimeSeriesAvgPrice200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAvgPrice200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAvgPrice200ResponseValuesInner(varGetTimeSeriesAvgPrice200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAvgPrice200ResponseValuesInner struct {
	value *GetTimeSeriesAvgPrice200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAvgPrice200ResponseValuesInner) Get() *GetTimeSeriesAvgPrice200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAvgPrice200ResponseValuesInner) Set(val *GetTimeSeriesAvgPrice200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAvgPrice200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAvgPrice200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAvgPrice200ResponseValuesInner(val *GetTimeSeriesAvgPrice200ResponseValuesInner) *NullableGetTimeSeriesAvgPrice200ResponseValuesInner {
	return &NullableGetTimeSeriesAvgPrice200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAvgPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAvgPrice200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
