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

// checks if the GetTimeSeriesSarExt200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSarExt200ResponseValuesInner{}

// GetTimeSeriesSarExt200ResponseValuesInner struct for GetTimeSeriesSarExt200ResponseValuesInner
type GetTimeSeriesSarExt200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SAREXT value
	Sarext string `json:"sarext"`
}

type _GetTimeSeriesSarExt200ResponseValuesInner GetTimeSeriesSarExt200ResponseValuesInner

// NewGetTimeSeriesSarExt200ResponseValuesInner instantiates a new GetTimeSeriesSarExt200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSarExt200ResponseValuesInner(datetime string, sarext string) *GetTimeSeriesSarExt200ResponseValuesInner {
	this := GetTimeSeriesSarExt200ResponseValuesInner{}
	this.Datetime = datetime
	this.Sarext = sarext
	return &this
}

// NewGetTimeSeriesSarExt200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSarExt200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSarExt200ResponseValuesInnerWithDefaults() *GetTimeSeriesSarExt200ResponseValuesInner {
	this := GetTimeSeriesSarExt200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSarExt200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSarext returns the Sarext field value
func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetSarext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sarext
}

// GetSarextOk returns a tuple with the Sarext field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetSarextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sarext, true
}

// SetSarext sets field value
func (o *GetTimeSeriesSarExt200ResponseValuesInner) SetSarext(v string) {
	o.Sarext = v
}

func (o GetTimeSeriesSarExt200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSarExt200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sarext"] = o.Sarext
	return toSerialize, nil
}

func (o *GetTimeSeriesSarExt200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sarext",
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

	varGetTimeSeriesSarExt200ResponseValuesInner := _GetTimeSeriesSarExt200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesSarExt200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSarExt200ResponseValuesInner(varGetTimeSeriesSarExt200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSarExt200ResponseValuesInner struct {
	value *GetTimeSeriesSarExt200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSarExt200ResponseValuesInner) Get() *GetTimeSeriesSarExt200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSarExt200ResponseValuesInner) Set(val *GetTimeSeriesSarExt200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSarExt200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSarExt200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSarExt200ResponseValuesInner(val *GetTimeSeriesSarExt200ResponseValuesInner) *NullableGetTimeSeriesSarExt200ResponseValuesInner {
	return &NullableGetTimeSeriesSarExt200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSarExt200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSarExt200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
