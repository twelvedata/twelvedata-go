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

// checks if the GetTimeSeriesRocp200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRocp200ResponseValuesInner{}

// GetTimeSeriesRocp200ResponseValuesInner struct for GetTimeSeriesRocp200ResponseValuesInner
type GetTimeSeriesRocp200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ROCP value
	Rocp string `json:"rocp"`
}

type _GetTimeSeriesRocp200ResponseValuesInner GetTimeSeriesRocp200ResponseValuesInner

// NewGetTimeSeriesRocp200ResponseValuesInner instantiates a new GetTimeSeriesRocp200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRocp200ResponseValuesInner(datetime string, rocp string) *GetTimeSeriesRocp200ResponseValuesInner {
	this := GetTimeSeriesRocp200ResponseValuesInner{}
	this.Datetime = datetime
	this.Rocp = rocp
	return &this
}

// NewGetTimeSeriesRocp200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRocp200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRocp200ResponseValuesInnerWithDefaults() *GetTimeSeriesRocp200ResponseValuesInner {
	this := GetTimeSeriesRocp200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesRocp200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocp200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesRocp200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetRocp returns the Rocp field value
func (o *GetTimeSeriesRocp200ResponseValuesInner) GetRocp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rocp
}

// GetRocpOk returns a tuple with the Rocp field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocp200ResponseValuesInner) GetRocpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rocp, true
}

// SetRocp sets field value
func (o *GetTimeSeriesRocp200ResponseValuesInner) SetRocp(v string) {
	o.Rocp = v
}

func (o GetTimeSeriesRocp200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRocp200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["rocp"] = o.Rocp
	return toSerialize, nil
}

func (o *GetTimeSeriesRocp200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"rocp",
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

	varGetTimeSeriesRocp200ResponseValuesInner := _GetTimeSeriesRocp200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesRocp200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRocp200ResponseValuesInner(varGetTimeSeriesRocp200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesRocp200ResponseValuesInner struct {
	value *GetTimeSeriesRocp200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesRocp200ResponseValuesInner) Get() *GetTimeSeriesRocp200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesRocp200ResponseValuesInner) Set(val *GetTimeSeriesRocp200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRocp200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRocp200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRocp200ResponseValuesInner(val *GetTimeSeriesRocp200ResponseValuesInner) *NullableGetTimeSeriesRocp200ResponseValuesInner {
	return &NullableGetTimeSeriesRocp200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRocp200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRocp200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
