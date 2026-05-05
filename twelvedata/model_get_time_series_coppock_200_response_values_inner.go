// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesCoppock200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCoppock200ResponseValuesInner{}

// GetTimeSeriesCoppock200ResponseValuesInner struct for GetTimeSeriesCoppock200ResponseValuesInner
type GetTimeSeriesCoppock200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Coppock value
	Coppock string `json:"coppock"`
}

type _GetTimeSeriesCoppock200ResponseValuesInner GetTimeSeriesCoppock200ResponseValuesInner

// NewGetTimeSeriesCoppock200ResponseValuesInner instantiates a new GetTimeSeriesCoppock200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCoppock200ResponseValuesInner(datetime string, coppock string) *GetTimeSeriesCoppock200ResponseValuesInner {
	this := GetTimeSeriesCoppock200ResponseValuesInner{}
	this.Datetime = datetime
	this.Coppock = coppock
	return &this
}

// NewGetTimeSeriesCoppock200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCoppock200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCoppock200ResponseValuesInnerWithDefaults() *GetTimeSeriesCoppock200ResponseValuesInner {
	this := GetTimeSeriesCoppock200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesCoppock200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCoppock returns the Coppock field value
func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetCoppock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Coppock
}

// GetCoppockOk returns a tuple with the Coppock field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetCoppockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coppock, true
}

// SetCoppock sets field value
func (o *GetTimeSeriesCoppock200ResponseValuesInner) SetCoppock(v string) {
	o.Coppock = v
}

func (o GetTimeSeriesCoppock200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCoppock200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["coppock"] = o.Coppock
	return toSerialize, nil
}

func (o *GetTimeSeriesCoppock200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"coppock",
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

	varGetTimeSeriesCoppock200ResponseValuesInner := _GetTimeSeriesCoppock200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCoppock200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCoppock200ResponseValuesInner(varGetTimeSeriesCoppock200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesCoppock200ResponseValuesInner struct {
	value *GetTimeSeriesCoppock200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesCoppock200ResponseValuesInner) Get() *GetTimeSeriesCoppock200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesCoppock200ResponseValuesInner) Set(val *GetTimeSeriesCoppock200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCoppock200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCoppock200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCoppock200ResponseValuesInner(val *GetTimeSeriesCoppock200ResponseValuesInner) *NullableGetTimeSeriesCoppock200ResponseValuesInner {
	return &NullableGetTimeSeriesCoppock200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCoppock200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCoppock200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
