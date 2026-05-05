// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAdOsc200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdOsc200ResponseValuesInner{}

// GetTimeSeriesAdOsc200ResponseValuesInner struct for GetTimeSeriesAdOsc200ResponseValuesInner
type GetTimeSeriesAdOsc200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Adosc value
	Adosc string `json:"adosc"`
}

type _GetTimeSeriesAdOsc200ResponseValuesInner GetTimeSeriesAdOsc200ResponseValuesInner

// NewGetTimeSeriesAdOsc200ResponseValuesInner instantiates a new GetTimeSeriesAdOsc200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdOsc200ResponseValuesInner(datetime string, adosc string) *GetTimeSeriesAdOsc200ResponseValuesInner {
	this := GetTimeSeriesAdOsc200ResponseValuesInner{}
	this.Datetime = datetime
	this.Adosc = adosc
	return &this
}

// NewGetTimeSeriesAdOsc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAdOsc200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdOsc200ResponseValuesInnerWithDefaults() *GetTimeSeriesAdOsc200ResponseValuesInner {
	this := GetTimeSeriesAdOsc200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAdOsc200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAdosc returns the Adosc field value
func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetAdosc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Adosc
}

// GetAdoscOk returns a tuple with the Adosc field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdOsc200ResponseValuesInner) GetAdoscOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Adosc, true
}

// SetAdosc sets field value
func (o *GetTimeSeriesAdOsc200ResponseValuesInner) SetAdosc(v string) {
	o.Adosc = v
}

func (o GetTimeSeriesAdOsc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdOsc200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["adosc"] = o.Adosc
	return toSerialize, nil
}

func (o *GetTimeSeriesAdOsc200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"adosc",
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

	varGetTimeSeriesAdOsc200ResponseValuesInner := _GetTimeSeriesAdOsc200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAdOsc200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdOsc200ResponseValuesInner(varGetTimeSeriesAdOsc200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAdOsc200ResponseValuesInner struct {
	value *GetTimeSeriesAdOsc200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAdOsc200ResponseValuesInner) Get() *GetTimeSeriesAdOsc200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAdOsc200ResponseValuesInner) Set(val *GetTimeSeriesAdOsc200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdOsc200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdOsc200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdOsc200ResponseValuesInner(val *GetTimeSeriesAdOsc200ResponseValuesInner) *NullableGetTimeSeriesAdOsc200ResponseValuesInner {
	return &NullableGetTimeSeriesAdOsc200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdOsc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdOsc200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
