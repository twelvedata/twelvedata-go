// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesRocr200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesRocr200ResponseValuesInner{}

// GetTimeSeriesRocr200ResponseValuesInner struct for GetTimeSeriesRocr200ResponseValuesInner
type GetTimeSeriesRocr200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// ROCR value
	Rocr string `json:"rocr"`
}

type _GetTimeSeriesRocr200ResponseValuesInner GetTimeSeriesRocr200ResponseValuesInner

// NewGetTimeSeriesRocr200ResponseValuesInner instantiates a new GetTimeSeriesRocr200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesRocr200ResponseValuesInner(datetime string, rocr string) *GetTimeSeriesRocr200ResponseValuesInner {
	this := GetTimeSeriesRocr200ResponseValuesInner{}
	this.Datetime = datetime
	this.Rocr = rocr
	return &this
}

// NewGetTimeSeriesRocr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRocr200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesRocr200ResponseValuesInnerWithDefaults() *GetTimeSeriesRocr200ResponseValuesInner {
	this := GetTimeSeriesRocr200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesRocr200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocr200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesRocr200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetRocr returns the Rocr field value
func (o *GetTimeSeriesRocr200ResponseValuesInner) GetRocr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rocr
}

// GetRocrOk returns a tuple with the Rocr field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesRocr200ResponseValuesInner) GetRocrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rocr, true
}

// SetRocr sets field value
func (o *GetTimeSeriesRocr200ResponseValuesInner) SetRocr(v string) {
	o.Rocr = v
}

func (o GetTimeSeriesRocr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesRocr200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["rocr"] = o.Rocr
	return toSerialize, nil
}

func (o *GetTimeSeriesRocr200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"rocr",
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

	varGetTimeSeriesRocr200ResponseValuesInner := _GetTimeSeriesRocr200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesRocr200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesRocr200ResponseValuesInner(varGetTimeSeriesRocr200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesRocr200ResponseValuesInner struct {
	value *GetTimeSeriesRocr200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesRocr200ResponseValuesInner) Get() *GetTimeSeriesRocr200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesRocr200ResponseValuesInner) Set(val *GetTimeSeriesRocr200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesRocr200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesRocr200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesRocr200ResponseValuesInner(val *GetTimeSeriesRocr200ResponseValuesInner) *NullableGetTimeSeriesRocr200ResponseValuesInner {
	return &NullableGetTimeSeriesRocr200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesRocr200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesRocr200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
