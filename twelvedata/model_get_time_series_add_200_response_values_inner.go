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

// checks if the GetTimeSeriesAdd200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdd200ResponseValuesInner{}

// GetTimeSeriesAdd200ResponseValuesInner struct for GetTimeSeriesAdd200ResponseValuesInner
type GetTimeSeriesAdd200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Add value
	Add string `json:"add"`
}

type _GetTimeSeriesAdd200ResponseValuesInner GetTimeSeriesAdd200ResponseValuesInner

// NewGetTimeSeriesAdd200ResponseValuesInner instantiates a new GetTimeSeriesAdd200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdd200ResponseValuesInner(datetime string, add string) *GetTimeSeriesAdd200ResponseValuesInner {
	this := GetTimeSeriesAdd200ResponseValuesInner{}
	this.Datetime = datetime
	this.Add = add
	return &this
}

// NewGetTimeSeriesAdd200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAdd200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdd200ResponseValuesInnerWithDefaults() *GetTimeSeriesAdd200ResponseValuesInner {
	this := GetTimeSeriesAdd200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAdd200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAdd200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAdd returns the Add field value
func (o *GetTimeSeriesAdd200ResponseValuesInner) GetAdd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Add
}

// GetAddOk returns a tuple with the Add field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdd200ResponseValuesInner) GetAddOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Add, true
}

// SetAdd sets field value
func (o *GetTimeSeriesAdd200ResponseValuesInner) SetAdd(v string) {
	o.Add = v
}

func (o GetTimeSeriesAdd200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdd200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["add"] = o.Add
	return toSerialize, nil
}

func (o *GetTimeSeriesAdd200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"add",
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

	varGetTimeSeriesAdd200ResponseValuesInner := _GetTimeSeriesAdd200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesAdd200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdd200ResponseValuesInner(varGetTimeSeriesAdd200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAdd200ResponseValuesInner struct {
	value *GetTimeSeriesAdd200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAdd200ResponseValuesInner) Get() *GetTimeSeriesAdd200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAdd200ResponseValuesInner) Set(val *GetTimeSeriesAdd200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdd200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdd200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdd200ResponseValuesInner(val *GetTimeSeriesAdd200ResponseValuesInner) *NullableGetTimeSeriesAdd200ResponseValuesInner {
	return &NullableGetTimeSeriesAdd200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdd200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdd200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
