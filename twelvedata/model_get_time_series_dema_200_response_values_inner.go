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

// checks if the GetTimeSeriesDema200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesDema200ResponseValuesInner{}

// GetTimeSeriesDema200ResponseValuesInner struct for GetTimeSeriesDema200ResponseValuesInner
type GetTimeSeriesDema200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Dema value
	Dema string `json:"dema"`
}

type _GetTimeSeriesDema200ResponseValuesInner GetTimeSeriesDema200ResponseValuesInner

// NewGetTimeSeriesDema200ResponseValuesInner instantiates a new GetTimeSeriesDema200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesDema200ResponseValuesInner(datetime string, dema string) *GetTimeSeriesDema200ResponseValuesInner {
	this := GetTimeSeriesDema200ResponseValuesInner{}
	this.Datetime = datetime
	this.Dema = dema
	return &this
}

// NewGetTimeSeriesDema200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesDema200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesDema200ResponseValuesInnerWithDefaults() *GetTimeSeriesDema200ResponseValuesInner {
	this := GetTimeSeriesDema200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesDema200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDema200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesDema200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetDema returns the Dema field value
func (o *GetTimeSeriesDema200ResponseValuesInner) GetDema() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dema
}

// GetDemaOk returns a tuple with the Dema field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDema200ResponseValuesInner) GetDemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dema, true
}

// SetDema sets field value
func (o *GetTimeSeriesDema200ResponseValuesInner) SetDema(v string) {
	o.Dema = v
}

func (o GetTimeSeriesDema200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesDema200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["dema"] = o.Dema
	return toSerialize, nil
}

func (o *GetTimeSeriesDema200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"dema",
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

	varGetTimeSeriesDema200ResponseValuesInner := _GetTimeSeriesDema200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesDema200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesDema200ResponseValuesInner(varGetTimeSeriesDema200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesDema200ResponseValuesInner struct {
	value *GetTimeSeriesDema200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesDema200ResponseValuesInner) Get() *GetTimeSeriesDema200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesDema200ResponseValuesInner) Set(val *GetTimeSeriesDema200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesDema200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesDema200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesDema200ResponseValuesInner(val *GetTimeSeriesDema200ResponseValuesInner) *NullableGetTimeSeriesDema200ResponseValuesInner {
	return &NullableGetTimeSeriesDema200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesDema200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesDema200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
