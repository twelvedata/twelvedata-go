// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesPivotPointsHL200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesPivotPointsHL200ResponseValuesInner{}

// GetTimeSeriesPivotPointsHL200ResponseValuesInner struct for GetTimeSeriesPivotPointsHL200ResponseValuesInner
type GetTimeSeriesPivotPointsHL200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// `1` if it is a high pivot point, otherwise `0`
	PivotPointH int64 `json:"pivot_point_h"`
	// `1` if it is a low pivot point, otherwise `0`
	PivotPointL int64 `json:"pivot_point_l"`
}

type _GetTimeSeriesPivotPointsHL200ResponseValuesInner GetTimeSeriesPivotPointsHL200ResponseValuesInner

// NewGetTimeSeriesPivotPointsHL200ResponseValuesInner instantiates a new GetTimeSeriesPivotPointsHL200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesPivotPointsHL200ResponseValuesInner(datetime string, pivotPointH int64, pivotPointL int64) *GetTimeSeriesPivotPointsHL200ResponseValuesInner {
	this := GetTimeSeriesPivotPointsHL200ResponseValuesInner{}
	this.Datetime = datetime
	this.PivotPointH = pivotPointH
	this.PivotPointL = pivotPointL
	return &this
}

// NewGetTimeSeriesPivotPointsHL200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPivotPointsHL200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesPivotPointsHL200ResponseValuesInnerWithDefaults() *GetTimeSeriesPivotPointsHL200ResponseValuesInner {
	this := GetTimeSeriesPivotPointsHL200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetPivotPointH returns the PivotPointH field value
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointH() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PivotPointH
}

// GetPivotPointHOk returns a tuple with the PivotPointH field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointHOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PivotPointH, true
}

// SetPivotPointH sets field value
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) SetPivotPointH(v int64) {
	o.PivotPointH = v
}

// GetPivotPointL returns the PivotPointL field value
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointL() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PivotPointL
}

// GetPivotPointLOk returns a tuple with the PivotPointL field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointLOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PivotPointL, true
}

// SetPivotPointL sets field value
func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) SetPivotPointL(v int64) {
	o.PivotPointL = v
}

func (o GetTimeSeriesPivotPointsHL200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesPivotPointsHL200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["pivot_point_h"] = o.PivotPointH
	toSerialize["pivot_point_l"] = o.PivotPointL
	return toSerialize, nil
}

func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"pivot_point_h",
		"pivot_point_l",
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

	varGetTimeSeriesPivotPointsHL200ResponseValuesInner := _GetTimeSeriesPivotPointsHL200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesPivotPointsHL200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesPivotPointsHL200ResponseValuesInner(varGetTimeSeriesPivotPointsHL200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner struct {
	value *GetTimeSeriesPivotPointsHL200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner) Get() *GetTimeSeriesPivotPointsHL200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner) Set(val *GetTimeSeriesPivotPointsHL200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesPivotPointsHL200ResponseValuesInner(val *GetTimeSeriesPivotPointsHL200ResponseValuesInner) *NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner {
	return &NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesPivotPointsHL200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
