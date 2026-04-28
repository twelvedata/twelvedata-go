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

// checks if the GetTimeSeriesWclPrice200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesWclPrice200ResponseValuesInner{}

// GetTimeSeriesWclPrice200ResponseValuesInner struct for GetTimeSeriesWclPrice200ResponseValuesInner
type GetTimeSeriesWclPrice200ResponseValuesInner struct {
	// datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// wclprice value
	Wclprice string `json:"wclprice"`
}

type _GetTimeSeriesWclPrice200ResponseValuesInner GetTimeSeriesWclPrice200ResponseValuesInner

// NewGetTimeSeriesWclPrice200ResponseValuesInner instantiates a new GetTimeSeriesWclPrice200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesWclPrice200ResponseValuesInner(datetime string, wclprice string) *GetTimeSeriesWclPrice200ResponseValuesInner {
	this := GetTimeSeriesWclPrice200ResponseValuesInner{}
	this.Datetime = datetime
	this.Wclprice = wclprice
	return &this
}

// NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesWclPrice200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesWclPrice200ResponseValuesInner {
	this := GetTimeSeriesWclPrice200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesWclPrice200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetWclprice returns the Wclprice field value
func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetWclprice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Wclprice
}

// GetWclpriceOk returns a tuple with the Wclprice field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetWclpriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wclprice, true
}

// SetWclprice sets field value
func (o *GetTimeSeriesWclPrice200ResponseValuesInner) SetWclprice(v string) {
	o.Wclprice = v
}

func (o GetTimeSeriesWclPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesWclPrice200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["wclprice"] = o.Wclprice
	return toSerialize, nil
}

func (o *GetTimeSeriesWclPrice200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"wclprice",
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

	varGetTimeSeriesWclPrice200ResponseValuesInner := _GetTimeSeriesWclPrice200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesWclPrice200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesWclPrice200ResponseValuesInner(varGetTimeSeriesWclPrice200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesWclPrice200ResponseValuesInner struct {
	value *GetTimeSeriesWclPrice200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesWclPrice200ResponseValuesInner) Get() *GetTimeSeriesWclPrice200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesWclPrice200ResponseValuesInner) Set(val *GetTimeSeriesWclPrice200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesWclPrice200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesWclPrice200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesWclPrice200ResponseValuesInner(val *GetTimeSeriesWclPrice200ResponseValuesInner) *NullableGetTimeSeriesWclPrice200ResponseValuesInner {
	return &NullableGetTimeSeriesWclPrice200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesWclPrice200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesWclPrice200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
