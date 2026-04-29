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

// checks if the GetTimeSeriesSuperTrend200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSuperTrend200ResponseValuesInner{}

// GetTimeSeriesSuperTrend200ResponseValuesInner struct for GetTimeSeriesSuperTrend200ResponseValuesInner
type GetTimeSeriesSuperTrend200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SuperTrend value
	Supertrend string `json:"supertrend"`
}

type _GetTimeSeriesSuperTrend200ResponseValuesInner GetTimeSeriesSuperTrend200ResponseValuesInner

// NewGetTimeSeriesSuperTrend200ResponseValuesInner instantiates a new GetTimeSeriesSuperTrend200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSuperTrend200ResponseValuesInner(datetime string, supertrend string) *GetTimeSeriesSuperTrend200ResponseValuesInner {
	this := GetTimeSeriesSuperTrend200ResponseValuesInner{}
	this.Datetime = datetime
	this.Supertrend = supertrend
	return &this
}

// NewGetTimeSeriesSuperTrend200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSuperTrend200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSuperTrend200ResponseValuesInnerWithDefaults() *GetTimeSeriesSuperTrend200ResponseValuesInner {
	this := GetTimeSeriesSuperTrend200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSupertrend returns the Supertrend field value
func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetSupertrend() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supertrend
}

// GetSupertrendOk returns a tuple with the Supertrend field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetSupertrendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supertrend, true
}

// SetSupertrend sets field value
func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) SetSupertrend(v string) {
	o.Supertrend = v
}

func (o GetTimeSeriesSuperTrend200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSuperTrend200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["supertrend"] = o.Supertrend
	return toSerialize, nil
}

func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"supertrend",
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

	varGetTimeSeriesSuperTrend200ResponseValuesInner := _GetTimeSeriesSuperTrend200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSuperTrend200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSuperTrend200ResponseValuesInner(varGetTimeSeriesSuperTrend200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSuperTrend200ResponseValuesInner struct {
	value *GetTimeSeriesSuperTrend200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSuperTrend200ResponseValuesInner) Get() *GetTimeSeriesSuperTrend200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSuperTrend200ResponseValuesInner) Set(val *GetTimeSeriesSuperTrend200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSuperTrend200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSuperTrend200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSuperTrend200ResponseValuesInner(val *GetTimeSeriesSuperTrend200ResponseValuesInner) *NullableGetTimeSeriesSuperTrend200ResponseValuesInner {
	return &NullableGetTimeSeriesSuperTrend200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSuperTrend200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSuperTrend200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
