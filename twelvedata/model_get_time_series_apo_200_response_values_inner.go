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

// checks if the GetTimeSeriesApo200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesApo200ResponseValuesInner{}

// GetTimeSeriesApo200ResponseValuesInner struct for GetTimeSeriesApo200ResponseValuesInner
type GetTimeSeriesApo200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// APO value
	Apo string `json:"apo"`
}

type _GetTimeSeriesApo200ResponseValuesInner GetTimeSeriesApo200ResponseValuesInner

// NewGetTimeSeriesApo200ResponseValuesInner instantiates a new GetTimeSeriesApo200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesApo200ResponseValuesInner(datetime string, apo string) *GetTimeSeriesApo200ResponseValuesInner {
	this := GetTimeSeriesApo200ResponseValuesInner{}
	this.Datetime = datetime
	this.Apo = apo
	return &this
}

// NewGetTimeSeriesApo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesApo200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesApo200ResponseValuesInnerWithDefaults() *GetTimeSeriesApo200ResponseValuesInner {
	this := GetTimeSeriesApo200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesApo200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesApo200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesApo200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetApo returns the Apo field value
func (o *GetTimeSeriesApo200ResponseValuesInner) GetApo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apo
}

// GetApoOk returns a tuple with the Apo field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesApo200ResponseValuesInner) GetApoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apo, true
}

// SetApo sets field value
func (o *GetTimeSeriesApo200ResponseValuesInner) SetApo(v string) {
	o.Apo = v
}

func (o GetTimeSeriesApo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesApo200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["apo"] = o.Apo
	return toSerialize, nil
}

func (o *GetTimeSeriesApo200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"apo",
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

	varGetTimeSeriesApo200ResponseValuesInner := _GetTimeSeriesApo200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesApo200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesApo200ResponseValuesInner(varGetTimeSeriesApo200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesApo200ResponseValuesInner struct {
	value *GetTimeSeriesApo200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesApo200ResponseValuesInner) Get() *GetTimeSeriesApo200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesApo200ResponseValuesInner) Set(val *GetTimeSeriesApo200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesApo200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesApo200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesApo200ResponseValuesInner(val *GetTimeSeriesApo200ResponseValuesInner) *NullableGetTimeSeriesApo200ResponseValuesInner {
	return &NullableGetTimeSeriesApo200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesApo200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesApo200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
