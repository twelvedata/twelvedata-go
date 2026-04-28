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

// checks if the GetTimeSeriesSar200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesSar200ResponseValuesInner{}

// GetTimeSeriesSar200ResponseValuesInner struct for GetTimeSeriesSar200ResponseValuesInner
type GetTimeSeriesSar200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// SAR value
	Sar string `json:"sar"`
}

type _GetTimeSeriesSar200ResponseValuesInner GetTimeSeriesSar200ResponseValuesInner

// NewGetTimeSeriesSar200ResponseValuesInner instantiates a new GetTimeSeriesSar200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesSar200ResponseValuesInner(datetime string, sar string) *GetTimeSeriesSar200ResponseValuesInner {
	this := GetTimeSeriesSar200ResponseValuesInner{}
	this.Datetime = datetime
	this.Sar = sar
	return &this
}

// NewGetTimeSeriesSar200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSar200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesSar200ResponseValuesInnerWithDefaults() *GetTimeSeriesSar200ResponseValuesInner {
	this := GetTimeSeriesSar200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesSar200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesSar200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSar returns the Sar field value
func (o *GetTimeSeriesSar200ResponseValuesInner) GetSar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sar
}

// GetSarOk returns a tuple with the Sar field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesSar200ResponseValuesInner) GetSarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sar, true
}

// SetSar sets field value
func (o *GetTimeSeriesSar200ResponseValuesInner) SetSar(v string) {
	o.Sar = v
}

func (o GetTimeSeriesSar200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesSar200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["sar"] = o.Sar
	return toSerialize, nil
}

func (o *GetTimeSeriesSar200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"sar",
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

	varGetTimeSeriesSar200ResponseValuesInner := _GetTimeSeriesSar200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesSar200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesSar200ResponseValuesInner(varGetTimeSeriesSar200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesSar200ResponseValuesInner struct {
	value *GetTimeSeriesSar200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesSar200ResponseValuesInner) Get() *GetTimeSeriesSar200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesSar200ResponseValuesInner) Set(val *GetTimeSeriesSar200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesSar200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesSar200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesSar200ResponseValuesInner(val *GetTimeSeriesSar200ResponseValuesInner) *NullableGetTimeSeriesSar200ResponseValuesInner {
	return &NullableGetTimeSeriesSar200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesSar200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesSar200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
