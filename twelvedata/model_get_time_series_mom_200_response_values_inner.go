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

// checks if the GetTimeSeriesMom200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMom200ResponseValuesInner{}

// GetTimeSeriesMom200ResponseValuesInner struct for GetTimeSeriesMom200ResponseValuesInner
type GetTimeSeriesMom200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Mom value
	Mom string `json:"mom"`
}

type _GetTimeSeriesMom200ResponseValuesInner GetTimeSeriesMom200ResponseValuesInner

// NewGetTimeSeriesMom200ResponseValuesInner instantiates a new GetTimeSeriesMom200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMom200ResponseValuesInner(datetime string, mom string) *GetTimeSeriesMom200ResponseValuesInner {
	this := GetTimeSeriesMom200ResponseValuesInner{}
	this.Datetime = datetime
	this.Mom = mom
	return &this
}

// NewGetTimeSeriesMom200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMom200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMom200ResponseValuesInnerWithDefaults() *GetTimeSeriesMom200ResponseValuesInner {
	this := GetTimeSeriesMom200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMom200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMom200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMom200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMom returns the Mom field value
func (o *GetTimeSeriesMom200ResponseValuesInner) GetMom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mom
}

// GetMomOk returns a tuple with the Mom field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMom200ResponseValuesInner) GetMomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mom, true
}

// SetMom sets field value
func (o *GetTimeSeriesMom200ResponseValuesInner) SetMom(v string) {
	o.Mom = v
}

func (o GetTimeSeriesMom200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMom200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["mom"] = o.Mom
	return toSerialize, nil
}

func (o *GetTimeSeriesMom200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"mom",
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

	varGetTimeSeriesMom200ResponseValuesInner := _GetTimeSeriesMom200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMom200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMom200ResponseValuesInner(varGetTimeSeriesMom200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMom200ResponseValuesInner struct {
	value *GetTimeSeriesMom200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMom200ResponseValuesInner) Get() *GetTimeSeriesMom200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMom200ResponseValuesInner) Set(val *GetTimeSeriesMom200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMom200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMom200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMom200ResponseValuesInner(val *GetTimeSeriesMom200ResponseValuesInner) *NullableGetTimeSeriesMom200ResponseValuesInner {
	return &NullableGetTimeSeriesMom200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMom200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMom200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
