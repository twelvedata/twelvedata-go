// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesDx200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesDx200ResponseValuesInner{}

// GetTimeSeriesDx200ResponseValuesInner struct for GetTimeSeriesDx200ResponseValuesInner
type GetTimeSeriesDx200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// dx value
	Dx string `json:"dx"`
}

type _GetTimeSeriesDx200ResponseValuesInner GetTimeSeriesDx200ResponseValuesInner

// NewGetTimeSeriesDx200ResponseValuesInner instantiates a new GetTimeSeriesDx200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesDx200ResponseValuesInner(datetime string, dx string) *GetTimeSeriesDx200ResponseValuesInner {
	this := GetTimeSeriesDx200ResponseValuesInner{}
	this.Datetime = datetime
	this.Dx = dx
	return &this
}

// NewGetTimeSeriesDx200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesDx200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesDx200ResponseValuesInnerWithDefaults() *GetTimeSeriesDx200ResponseValuesInner {
	this := GetTimeSeriesDx200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesDx200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDx200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesDx200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetDx returns the Dx field value
func (o *GetTimeSeriesDx200ResponseValuesInner) GetDx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dx
}

// GetDxOk returns a tuple with the Dx field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesDx200ResponseValuesInner) GetDxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dx, true
}

// SetDx sets field value
func (o *GetTimeSeriesDx200ResponseValuesInner) SetDx(v string) {
	o.Dx = v
}

func (o GetTimeSeriesDx200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesDx200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["dx"] = o.Dx
	return toSerialize, nil
}

func (o *GetTimeSeriesDx200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"dx",
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

	varGetTimeSeriesDx200ResponseValuesInner := _GetTimeSeriesDx200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesDx200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesDx200ResponseValuesInner(varGetTimeSeriesDx200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesDx200ResponseValuesInner struct {
	value *GetTimeSeriesDx200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesDx200ResponseValuesInner) Get() *GetTimeSeriesDx200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesDx200ResponseValuesInner) Set(val *GetTimeSeriesDx200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesDx200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesDx200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesDx200ResponseValuesInner(val *GetTimeSeriesDx200ResponseValuesInner) *NullableGetTimeSeriesDx200ResponseValuesInner {
	return &NullableGetTimeSeriesDx200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesDx200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesDx200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
