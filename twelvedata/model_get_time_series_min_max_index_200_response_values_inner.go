// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMinMaxIndex200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMinMaxIndex200ResponseValuesInner{}

// GetTimeSeriesMinMaxIndex200ResponseValuesInner struct for GetTimeSeriesMinMaxIndex200ResponseValuesInner
type GetTimeSeriesMinMaxIndex200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Index of the lowest value over the specified period
	Minidx string `json:"minidx"`
	// Index of the highest value over the specified period
	Maxidx string `json:"maxidx"`
}

type _GetTimeSeriesMinMaxIndex200ResponseValuesInner GetTimeSeriesMinMaxIndex200ResponseValuesInner

// NewGetTimeSeriesMinMaxIndex200ResponseValuesInner instantiates a new GetTimeSeriesMinMaxIndex200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMinMaxIndex200ResponseValuesInner(datetime string, minidx string, maxidx string) *GetTimeSeriesMinMaxIndex200ResponseValuesInner {
	this := GetTimeSeriesMinMaxIndex200ResponseValuesInner{}
	this.Datetime = datetime
	this.Minidx = minidx
	this.Maxidx = maxidx
	return &this
}

// NewGetTimeSeriesMinMaxIndex200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinMaxIndex200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMinMaxIndex200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinMaxIndex200ResponseValuesInner {
	this := GetTimeSeriesMinMaxIndex200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMinidx returns the Minidx field value
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMinidx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Minidx
}

// GetMinidxOk returns a tuple with the Minidx field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMinidxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minidx, true
}

// SetMinidx sets field value
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) SetMinidx(v string) {
	o.Minidx = v
}

// GetMaxidx returns the Maxidx field value
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMaxidx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Maxidx
}

// GetMaxidxOk returns a tuple with the Maxidx field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMaxidxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maxidx, true
}

// SetMaxidx sets field value
func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) SetMaxidx(v string) {
	o.Maxidx = v
}

func (o GetTimeSeriesMinMaxIndex200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMinMaxIndex200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["minidx"] = o.Minidx
	toSerialize["maxidx"] = o.Maxidx
	return toSerialize, nil
}

func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"minidx",
		"maxidx",
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

	varGetTimeSeriesMinMaxIndex200ResponseValuesInner := _GetTimeSeriesMinMaxIndex200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMinMaxIndex200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMinMaxIndex200ResponseValuesInner(varGetTimeSeriesMinMaxIndex200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner struct {
	value *GetTimeSeriesMinMaxIndex200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner) Get() *GetTimeSeriesMinMaxIndex200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner) Set(val *GetTimeSeriesMinMaxIndex200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMinMaxIndex200ResponseValuesInner(val *GetTimeSeriesMinMaxIndex200ResponseValuesInner) *NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner {
	return &NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMinMaxIndex200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
