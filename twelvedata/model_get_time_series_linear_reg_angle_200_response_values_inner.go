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

// checks if the GetTimeSeriesLinearRegAngle200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesLinearRegAngle200ResponseValuesInner{}

// GetTimeSeriesLinearRegAngle200ResponseValuesInner struct for GetTimeSeriesLinearRegAngle200ResponseValuesInner
type GetTimeSeriesLinearRegAngle200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Linear regression angle value
	Linearregangle string `json:"linearregangle"`
}

type _GetTimeSeriesLinearRegAngle200ResponseValuesInner GetTimeSeriesLinearRegAngle200ResponseValuesInner

// NewGetTimeSeriesLinearRegAngle200ResponseValuesInner instantiates a new GetTimeSeriesLinearRegAngle200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesLinearRegAngle200ResponseValuesInner(datetime string, linearregangle string) *GetTimeSeriesLinearRegAngle200ResponseValuesInner {
	this := GetTimeSeriesLinearRegAngle200ResponseValuesInner{}
	this.Datetime = datetime
	this.Linearregangle = linearregangle
	return &this
}

// NewGetTimeSeriesLinearRegAngle200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLinearRegAngle200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesLinearRegAngle200ResponseValuesInnerWithDefaults() *GetTimeSeriesLinearRegAngle200ResponseValuesInner {
	this := GetTimeSeriesLinearRegAngle200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetLinearregangle returns the Linearregangle field value
func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetLinearregangle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Linearregangle
}

// GetLinearregangleOk returns a tuple with the Linearregangle field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetLinearregangleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linearregangle, true
}

// SetLinearregangle sets field value
func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) SetLinearregangle(v string) {
	o.Linearregangle = v
}

func (o GetTimeSeriesLinearRegAngle200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesLinearRegAngle200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["linearregangle"] = o.Linearregangle
	return toSerialize, nil
}

func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"linearregangle",
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

	varGetTimeSeriesLinearRegAngle200ResponseValuesInner := _GetTimeSeriesLinearRegAngle200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesLinearRegAngle200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesLinearRegAngle200ResponseValuesInner(varGetTimeSeriesLinearRegAngle200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner struct {
	value *GetTimeSeriesLinearRegAngle200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner) Get() *GetTimeSeriesLinearRegAngle200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner) Set(val *GetTimeSeriesLinearRegAngle200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesLinearRegAngle200ResponseValuesInner(val *GetTimeSeriesLinearRegAngle200ResponseValuesInner) *NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner {
	return &NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesLinearRegAngle200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
