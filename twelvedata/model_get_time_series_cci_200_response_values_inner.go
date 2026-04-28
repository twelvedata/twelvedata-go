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

// checks if the GetTimeSeriesCci200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCci200ResponseValuesInner{}

// GetTimeSeriesCci200ResponseValuesInner struct for GetTimeSeriesCci200ResponseValuesInner
type GetTimeSeriesCci200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// CCI value
	Cci string `json:"cci"`
}

type _GetTimeSeriesCci200ResponseValuesInner GetTimeSeriesCci200ResponseValuesInner

// NewGetTimeSeriesCci200ResponseValuesInner instantiates a new GetTimeSeriesCci200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCci200ResponseValuesInner(datetime string, cci string) *GetTimeSeriesCci200ResponseValuesInner {
	this := GetTimeSeriesCci200ResponseValuesInner{}
	this.Datetime = datetime
	this.Cci = cci
	return &this
}

// NewGetTimeSeriesCci200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCci200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCci200ResponseValuesInnerWithDefaults() *GetTimeSeriesCci200ResponseValuesInner {
	this := GetTimeSeriesCci200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesCci200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCci200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesCci200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCci returns the Cci field value
func (o *GetTimeSeriesCci200ResponseValuesInner) GetCci() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cci
}

// GetCciOk returns a tuple with the Cci field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCci200ResponseValuesInner) GetCciOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cci, true
}

// SetCci sets field value
func (o *GetTimeSeriesCci200ResponseValuesInner) SetCci(v string) {
	o.Cci = v
}

func (o GetTimeSeriesCci200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCci200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["cci"] = o.Cci
	return toSerialize, nil
}

func (o *GetTimeSeriesCci200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"cci",
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

	varGetTimeSeriesCci200ResponseValuesInner := _GetTimeSeriesCci200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesCci200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCci200ResponseValuesInner(varGetTimeSeriesCci200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesCci200ResponseValuesInner struct {
	value *GetTimeSeriesCci200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesCci200ResponseValuesInner) Get() *GetTimeSeriesCci200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesCci200ResponseValuesInner) Set(val *GetTimeSeriesCci200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCci200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCci200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCci200ResponseValuesInner(val *GetTimeSeriesCci200ResponseValuesInner) *NullableGetTimeSeriesCci200ResponseValuesInner {
	return &NullableGetTimeSeriesCci200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCci200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCci200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
