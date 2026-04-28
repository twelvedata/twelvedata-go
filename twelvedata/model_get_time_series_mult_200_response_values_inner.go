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

// checks if the GetTimeSeriesMult200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMult200ResponseValuesInner{}

// GetTimeSeriesMult200ResponseValuesInner struct for GetTimeSeriesMult200ResponseValuesInner
type GetTimeSeriesMult200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Mult value
	Mult string `json:"mult"`
}

type _GetTimeSeriesMult200ResponseValuesInner GetTimeSeriesMult200ResponseValuesInner

// NewGetTimeSeriesMult200ResponseValuesInner instantiates a new GetTimeSeriesMult200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMult200ResponseValuesInner(datetime string, mult string) *GetTimeSeriesMult200ResponseValuesInner {
	this := GetTimeSeriesMult200ResponseValuesInner{}
	this.Datetime = datetime
	this.Mult = mult
	return &this
}

// NewGetTimeSeriesMult200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMult200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMult200ResponseValuesInnerWithDefaults() *GetTimeSeriesMult200ResponseValuesInner {
	this := GetTimeSeriesMult200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMult200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMult200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMult returns the Mult field value
func (o *GetTimeSeriesMult200ResponseValuesInner) GetMult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mult
}

// GetMultOk returns a tuple with the Mult field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMult200ResponseValuesInner) GetMultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mult, true
}

// SetMult sets field value
func (o *GetTimeSeriesMult200ResponseValuesInner) SetMult(v string) {
	o.Mult = v
}

func (o GetTimeSeriesMult200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMult200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["mult"] = o.Mult
	return toSerialize, nil
}

func (o *GetTimeSeriesMult200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"mult",
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

	varGetTimeSeriesMult200ResponseValuesInner := _GetTimeSeriesMult200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMult200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMult200ResponseValuesInner(varGetTimeSeriesMult200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMult200ResponseValuesInner struct {
	value *GetTimeSeriesMult200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMult200ResponseValuesInner) Get() *GetTimeSeriesMult200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMult200ResponseValuesInner) Set(val *GetTimeSeriesMult200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMult200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMult200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMult200ResponseValuesInner(val *GetTimeSeriesMult200ResponseValuesInner) *NullableGetTimeSeriesMult200ResponseValuesInner {
	return &NullableGetTimeSeriesMult200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMult200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMult200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
