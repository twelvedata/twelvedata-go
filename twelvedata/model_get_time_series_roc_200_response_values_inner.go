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

// checks if the GetTimeSeriesRoc200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRoc200ResponseValuesInner{}

// GetTimeSeriesRoc200ResponseValuesInner struct for GetTimeSeriesRoc200ResponseValuesInner
type GetTimeSeriesRoc200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// roc value
	Roc string `json:"roc"`
}

type _GetTimeSeriesRoc200ResponseValuesInner GetTimeSeriesRoc200ResponseValuesInner

// NewGetTimeSeriesRoc200ResponseValuesInner instantiates a new GetTimeSeriesRoc200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRoc200ResponseValuesInner(datetime string, roc string) *GetTimeSeriesRoc200ResponseValuesInner {
	this := GetTimeSeriesRoc200ResponseValuesInner{}
	this.Datetime = datetime
	this.Roc = roc
	return &this
}

// NewGetTimeSeriesRoc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRoc200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRoc200ResponseValuesInnerWithDefaults() *GetTimeSeriesRoc200ResponseValuesInner {
	this := GetTimeSeriesRoc200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesRoc200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRoc200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesRoc200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetRoc returns the Roc field value
func (o *GetTimeSeriesRoc200ResponseValuesInner) GetRoc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Roc
}

// GetRocOk returns a tuple with the Roc field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRoc200ResponseValuesInner) GetRocOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roc, true
}

// SetRoc sets field value
func (o *GetTimeSeriesRoc200ResponseValuesInner) SetRoc(v string) {
	o.Roc = v
}

func (o GetTimeSeriesRoc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRoc200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["roc"] = o.Roc
	return toSerialize, nil
}

func (o *GetTimeSeriesRoc200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"roc",
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

	varGetTimeSeriesRoc200ResponseValuesInner := _GetTimeSeriesRoc200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesRoc200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRoc200ResponseValuesInner(varGetTimeSeriesRoc200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesRoc200ResponseValuesInner struct {
	value *GetTimeSeriesRoc200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesRoc200ResponseValuesInner) Get() *GetTimeSeriesRoc200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesRoc200ResponseValuesInner) Set(val *GetTimeSeriesRoc200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRoc200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRoc200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRoc200ResponseValuesInner(val *GetTimeSeriesRoc200ResponseValuesInner) *NullableGetTimeSeriesRoc200ResponseValuesInner {
	return &NullableGetTimeSeriesRoc200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRoc200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRoc200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
