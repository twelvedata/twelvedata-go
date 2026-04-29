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

// checks if the GetTimeSeriesBop200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesBop200ResponseValuesInner{}

// GetTimeSeriesBop200ResponseValuesInner struct for GetTimeSeriesBop200ResponseValuesInner
type GetTimeSeriesBop200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Bop value
	Bop string `json:"bop"`
}

type _GetTimeSeriesBop200ResponseValuesInner GetTimeSeriesBop200ResponseValuesInner

// NewGetTimeSeriesBop200ResponseValuesInner instantiates a new GetTimeSeriesBop200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesBop200ResponseValuesInner(datetime string, bop string) *GetTimeSeriesBop200ResponseValuesInner {
	this := GetTimeSeriesBop200ResponseValuesInner{}
	this.Datetime = datetime
	this.Bop = bop
	return &this
}

// NewGetTimeSeriesBop200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesBop200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesBop200ResponseValuesInnerWithDefaults() *GetTimeSeriesBop200ResponseValuesInner {
	this := GetTimeSeriesBop200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesBop200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBop200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesBop200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetBop returns the Bop field value
func (o *GetTimeSeriesBop200ResponseValuesInner) GetBop() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bop
}

// GetBopOk returns a tuple with the Bop field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBop200ResponseValuesInner) GetBopOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bop, true
}

// SetBop sets field value
func (o *GetTimeSeriesBop200ResponseValuesInner) SetBop(v string) {
	o.Bop = v
}

func (o GetTimeSeriesBop200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesBop200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["bop"] = o.Bop
	return toSerialize, nil
}

func (o *GetTimeSeriesBop200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"bop",
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

	varGetTimeSeriesBop200ResponseValuesInner := _GetTimeSeriesBop200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesBop200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesBop200ResponseValuesInner(varGetTimeSeriesBop200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesBop200ResponseValuesInner struct {
	value *GetTimeSeriesBop200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesBop200ResponseValuesInner) Get() *GetTimeSeriesBop200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesBop200ResponseValuesInner) Set(val *GetTimeSeriesBop200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesBop200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesBop200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesBop200ResponseValuesInner(val *GetTimeSeriesBop200ResponseValuesInner) *NullableGetTimeSeriesBop200ResponseValuesInner {
	return &NullableGetTimeSeriesBop200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesBop200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesBop200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
