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

// checks if the GetTimeSeriesObv200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesObv200ResponseValuesInner{}

// GetTimeSeriesObv200ResponseValuesInner struct for GetTimeSeriesObv200ResponseValuesInner
type GetTimeSeriesObv200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// obv value
	Obv string `json:"obv"`
}

type _GetTimeSeriesObv200ResponseValuesInner GetTimeSeriesObv200ResponseValuesInner

// NewGetTimeSeriesObv200ResponseValuesInner instantiates a new GetTimeSeriesObv200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesObv200ResponseValuesInner(datetime string, obv string) *GetTimeSeriesObv200ResponseValuesInner {
	this := GetTimeSeriesObv200ResponseValuesInner{}
	this.Datetime = datetime
	this.Obv = obv
	return &this
}

// NewGetTimeSeriesObv200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesObv200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesObv200ResponseValuesInnerWithDefaults() *GetTimeSeriesObv200ResponseValuesInner {
	this := GetTimeSeriesObv200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesObv200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesObv200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesObv200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetObv returns the Obv field value
func (o *GetTimeSeriesObv200ResponseValuesInner) GetObv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Obv
}

// GetObvOk returns a tuple with the Obv field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesObv200ResponseValuesInner) GetObvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Obv, true
}

// SetObv sets field value
func (o *GetTimeSeriesObv200ResponseValuesInner) SetObv(v string) {
	o.Obv = v
}

func (o GetTimeSeriesObv200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesObv200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["obv"] = o.Obv
	return toSerialize, nil
}

func (o *GetTimeSeriesObv200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"obv",
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

	varGetTimeSeriesObv200ResponseValuesInner := _GetTimeSeriesObv200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesObv200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesObv200ResponseValuesInner(varGetTimeSeriesObv200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesObv200ResponseValuesInner struct {
	value *GetTimeSeriesObv200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesObv200ResponseValuesInner) Get() *GetTimeSeriesObv200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesObv200ResponseValuesInner) Set(val *GetTimeSeriesObv200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesObv200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesObv200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesObv200ResponseValuesInner(val *GetTimeSeriesObv200ResponseValuesInner) *NullableGetTimeSeriesObv200ResponseValuesInner {
	return &NullableGetTimeSeriesObv200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesObv200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesObv200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
