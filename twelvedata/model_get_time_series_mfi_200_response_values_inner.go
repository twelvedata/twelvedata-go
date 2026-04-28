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

// checks if the GetTimeSeriesMfi200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMfi200ResponseValuesInner{}

// GetTimeSeriesMfi200ResponseValuesInner struct for GetTimeSeriesMfi200ResponseValuesInner
type GetTimeSeriesMfi200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// MFI value
	Mfi string `json:"mfi"`
}

type _GetTimeSeriesMfi200ResponseValuesInner GetTimeSeriesMfi200ResponseValuesInner

// NewGetTimeSeriesMfi200ResponseValuesInner instantiates a new GetTimeSeriesMfi200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMfi200ResponseValuesInner(datetime string, mfi string) *GetTimeSeriesMfi200ResponseValuesInner {
	this := GetTimeSeriesMfi200ResponseValuesInner{}
	this.Datetime = datetime
	this.Mfi = mfi
	return &this
}

// NewGetTimeSeriesMfi200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMfi200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMfi200ResponseValuesInnerWithDefaults() *GetTimeSeriesMfi200ResponseValuesInner {
	this := GetTimeSeriesMfi200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMfi200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMfi200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMfi200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMfi returns the Mfi field value
func (o *GetTimeSeriesMfi200ResponseValuesInner) GetMfi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mfi
}

// GetMfiOk returns a tuple with the Mfi field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMfi200ResponseValuesInner) GetMfiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mfi, true
}

// SetMfi sets field value
func (o *GetTimeSeriesMfi200ResponseValuesInner) SetMfi(v string) {
	o.Mfi = v
}

func (o GetTimeSeriesMfi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMfi200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["mfi"] = o.Mfi
	return toSerialize, nil
}

func (o *GetTimeSeriesMfi200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"mfi",
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

	varGetTimeSeriesMfi200ResponseValuesInner := _GetTimeSeriesMfi200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesMfi200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMfi200ResponseValuesInner(varGetTimeSeriesMfi200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMfi200ResponseValuesInner struct {
	value *GetTimeSeriesMfi200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMfi200ResponseValuesInner) Get() *GetTimeSeriesMfi200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMfi200ResponseValuesInner) Set(val *GetTimeSeriesMfi200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMfi200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMfi200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMfi200ResponseValuesInner(val *GetTimeSeriesMfi200ResponseValuesInner) *NullableGetTimeSeriesMfi200ResponseValuesInner {
	return &NullableGetTimeSeriesMfi200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMfi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMfi200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
