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

// checks if the GetTimeSeriesFloor200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesFloor200ResponseValuesInner{}

// GetTimeSeriesFloor200ResponseValuesInner struct for GetTimeSeriesFloor200ResponseValuesInner
type GetTimeSeriesFloor200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Floor value
	Floor string `json:"floor"`
}

type _GetTimeSeriesFloor200ResponseValuesInner GetTimeSeriesFloor200ResponseValuesInner

// NewGetTimeSeriesFloor200ResponseValuesInner instantiates a new GetTimeSeriesFloor200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesFloor200ResponseValuesInner(datetime string, floor string) *GetTimeSeriesFloor200ResponseValuesInner {
	this := GetTimeSeriesFloor200ResponseValuesInner{}
	this.Datetime = datetime
	this.Floor = floor
	return &this
}

// NewGetTimeSeriesFloor200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesFloor200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesFloor200ResponseValuesInnerWithDefaults() *GetTimeSeriesFloor200ResponseValuesInner {
	this := GetTimeSeriesFloor200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesFloor200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesFloor200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesFloor200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetFloor returns the Floor field value
func (o *GetTimeSeriesFloor200ResponseValuesInner) GetFloor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Floor
}

// GetFloorOk returns a tuple with the Floor field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesFloor200ResponseValuesInner) GetFloorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Floor, true
}

// SetFloor sets field value
func (o *GetTimeSeriesFloor200ResponseValuesInner) SetFloor(v string) {
	o.Floor = v
}

func (o GetTimeSeriesFloor200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesFloor200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["floor"] = o.Floor
	return toSerialize, nil
}

func (o *GetTimeSeriesFloor200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"floor",
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

	varGetTimeSeriesFloor200ResponseValuesInner := _GetTimeSeriesFloor200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesFloor200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesFloor200ResponseValuesInner(varGetTimeSeriesFloor200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesFloor200ResponseValuesInner struct {
	value *GetTimeSeriesFloor200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesFloor200ResponseValuesInner) Get() *GetTimeSeriesFloor200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesFloor200ResponseValuesInner) Set(val *GetTimeSeriesFloor200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesFloor200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesFloor200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesFloor200ResponseValuesInner(val *GetTimeSeriesFloor200ResponseValuesInner) *NullableGetTimeSeriesFloor200ResponseValuesInner {
	return &NullableGetTimeSeriesFloor200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesFloor200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesFloor200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
