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

// checks if the GetTimeSeriesWillR200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesWillR200ResponseValuesInner{}

// GetTimeSeriesWillR200ResponseValuesInner struct for GetTimeSeriesWillR200ResponseValuesInner
type GetTimeSeriesWillR200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Williams %R value
	Willr string `json:"willr"`
}

type _GetTimeSeriesWillR200ResponseValuesInner GetTimeSeriesWillR200ResponseValuesInner

// NewGetTimeSeriesWillR200ResponseValuesInner instantiates a new GetTimeSeriesWillR200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesWillR200ResponseValuesInner(datetime string, willr string) *GetTimeSeriesWillR200ResponseValuesInner {
	this := GetTimeSeriesWillR200ResponseValuesInner{}
	this.Datetime = datetime
	this.Willr = willr
	return &this
}

// NewGetTimeSeriesWillR200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesWillR200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesWillR200ResponseValuesInnerWithDefaults() *GetTimeSeriesWillR200ResponseValuesInner {
	this := GetTimeSeriesWillR200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesWillR200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWillR200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesWillR200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetWillr returns the Willr field value
func (o *GetTimeSeriesWillR200ResponseValuesInner) GetWillr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Willr
}

// GetWillrOk returns a tuple with the Willr field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesWillR200ResponseValuesInner) GetWillrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Willr, true
}

// SetWillr sets field value
func (o *GetTimeSeriesWillR200ResponseValuesInner) SetWillr(v string) {
	o.Willr = v
}

func (o GetTimeSeriesWillR200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesWillR200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["willr"] = o.Willr
	return toSerialize, nil
}

func (o *GetTimeSeriesWillR200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"willr",
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

	varGetTimeSeriesWillR200ResponseValuesInner := _GetTimeSeriesWillR200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesWillR200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesWillR200ResponseValuesInner(varGetTimeSeriesWillR200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesWillR200ResponseValuesInner struct {
	value *GetTimeSeriesWillR200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesWillR200ResponseValuesInner) Get() *GetTimeSeriesWillR200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesWillR200ResponseValuesInner) Set(val *GetTimeSeriesWillR200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesWillR200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesWillR200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesWillR200ResponseValuesInner(val *GetTimeSeriesWillR200ResponseValuesInner) *NullableGetTimeSeriesWillR200ResponseValuesInner {
	return &NullableGetTimeSeriesWillR200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesWillR200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesWillR200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
