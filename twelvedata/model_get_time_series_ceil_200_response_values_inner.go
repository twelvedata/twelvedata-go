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

// checks if the GetTimeSeriesCeil200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCeil200ResponseValuesInner{}

// GetTimeSeriesCeil200ResponseValuesInner struct for GetTimeSeriesCeil200ResponseValuesInner
type GetTimeSeriesCeil200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Ceil value
	Ceil string `json:"ceil"`
}

type _GetTimeSeriesCeil200ResponseValuesInner GetTimeSeriesCeil200ResponseValuesInner

// NewGetTimeSeriesCeil200ResponseValuesInner instantiates a new GetTimeSeriesCeil200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCeil200ResponseValuesInner(datetime string, ceil string) *GetTimeSeriesCeil200ResponseValuesInner {
	this := GetTimeSeriesCeil200ResponseValuesInner{}
	this.Datetime = datetime
	this.Ceil = ceil
	return &this
}

// NewGetTimeSeriesCeil200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCeil200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCeil200ResponseValuesInnerWithDefaults() *GetTimeSeriesCeil200ResponseValuesInner {
	this := GetTimeSeriesCeil200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesCeil200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCeil200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesCeil200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCeil returns the Ceil field value
func (o *GetTimeSeriesCeil200ResponseValuesInner) GetCeil() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ceil
}

// GetCeilOk returns a tuple with the Ceil field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCeil200ResponseValuesInner) GetCeilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ceil, true
}

// SetCeil sets field value
func (o *GetTimeSeriesCeil200ResponseValuesInner) SetCeil(v string) {
	o.Ceil = v
}

func (o GetTimeSeriesCeil200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCeil200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["ceil"] = o.Ceil
	return toSerialize, nil
}

func (o *GetTimeSeriesCeil200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"ceil",
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

	varGetTimeSeriesCeil200ResponseValuesInner := _GetTimeSeriesCeil200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCeil200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCeil200ResponseValuesInner(varGetTimeSeriesCeil200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesCeil200ResponseValuesInner struct {
	value *GetTimeSeriesCeil200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesCeil200ResponseValuesInner) Get() *GetTimeSeriesCeil200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesCeil200ResponseValuesInner) Set(val *GetTimeSeriesCeil200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCeil200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCeil200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCeil200ResponseValuesInner(val *GetTimeSeriesCeil200ResponseValuesInner) *NullableGetTimeSeriesCeil200ResponseValuesInner {
	return &NullableGetTimeSeriesCeil200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCeil200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCeil200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
