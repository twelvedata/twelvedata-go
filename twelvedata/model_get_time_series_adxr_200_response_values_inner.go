// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesAdxr200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesAdxr200ResponseValuesInner{}

// GetTimeSeriesAdxr200ResponseValuesInner struct for GetTimeSeriesAdxr200ResponseValuesInner
type GetTimeSeriesAdxr200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Adxr value
	Adxr string `json:"adxr"`
}

type _GetTimeSeriesAdxr200ResponseValuesInner GetTimeSeriesAdxr200ResponseValuesInner

// NewGetTimeSeriesAdxr200ResponseValuesInner instantiates a new GetTimeSeriesAdxr200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesAdxr200ResponseValuesInner(datetime string, adxr string) *GetTimeSeriesAdxr200ResponseValuesInner {
	this := GetTimeSeriesAdxr200ResponseValuesInner{}
	this.Datetime = datetime
	this.Adxr = adxr
	return &this
}

// NewGetTimeSeriesAdxr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAdxr200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesAdxr200ResponseValuesInnerWithDefaults() *GetTimeSeriesAdxr200ResponseValuesInner {
	this := GetTimeSeriesAdxr200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesAdxr200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetAdxr returns the Adxr field value
func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetAdxr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Adxr
}

// GetAdxrOk returns a tuple with the Adxr field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetAdxrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Adxr, true
}

// SetAdxr sets field value
func (o *GetTimeSeriesAdxr200ResponseValuesInner) SetAdxr(v string) {
	o.Adxr = v
}

func (o GetTimeSeriesAdxr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesAdxr200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["adxr"] = o.Adxr
	return toSerialize, nil
}

func (o *GetTimeSeriesAdxr200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"adxr",
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

	varGetTimeSeriesAdxr200ResponseValuesInner := _GetTimeSeriesAdxr200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesAdxr200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesAdxr200ResponseValuesInner(varGetTimeSeriesAdxr200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesAdxr200ResponseValuesInner struct {
	value *GetTimeSeriesAdxr200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesAdxr200ResponseValuesInner) Get() *GetTimeSeriesAdxr200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesAdxr200ResponseValuesInner) Set(val *GetTimeSeriesAdxr200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesAdxr200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesAdxr200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesAdxr200ResponseValuesInner(val *GetTimeSeriesAdxr200ResponseValuesInner) *NullableGetTimeSeriesAdxr200ResponseValuesInner {
	return &NullableGetTimeSeriesAdxr200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesAdxr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesAdxr200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
