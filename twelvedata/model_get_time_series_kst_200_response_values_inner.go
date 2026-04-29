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

// checks if the GetTimeSeriesKst200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesKst200ResponseValuesInner{}

// GetTimeSeriesKst200ResponseValuesInner struct for GetTimeSeriesKst200ResponseValuesInner
type GetTimeSeriesKst200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// KST value
	Kst string `json:"kst"`
	// KST signal value
	KstSignal string `json:"kst_signal"`
}

type _GetTimeSeriesKst200ResponseValuesInner GetTimeSeriesKst200ResponseValuesInner

// NewGetTimeSeriesKst200ResponseValuesInner instantiates a new GetTimeSeriesKst200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesKst200ResponseValuesInner(datetime string, kst string, kstSignal string) *GetTimeSeriesKst200ResponseValuesInner {
	this := GetTimeSeriesKst200ResponseValuesInner{}
	this.Datetime = datetime
	this.Kst = kst
	this.KstSignal = kstSignal
	return &this
}

// NewGetTimeSeriesKst200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesKst200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesKst200ResponseValuesInnerWithDefaults() *GetTimeSeriesKst200ResponseValuesInner {
	this := GetTimeSeriesKst200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesKst200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesKst200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetKst returns the Kst field value
func (o *GetTimeSeriesKst200ResponseValuesInner) GetKst() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kst
}

// GetKstOk returns a tuple with the Kst field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseValuesInner) GetKstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kst, true
}

// SetKst sets field value
func (o *GetTimeSeriesKst200ResponseValuesInner) SetKst(v string) {
	o.Kst = v
}

// GetKstSignal returns the KstSignal field value
func (o *GetTimeSeriesKst200ResponseValuesInner) GetKstSignal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KstSignal
}

// GetKstSignalOk returns a tuple with the KstSignal field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesKst200ResponseValuesInner) GetKstSignalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KstSignal, true
}

// SetKstSignal sets field value
func (o *GetTimeSeriesKst200ResponseValuesInner) SetKstSignal(v string) {
	o.KstSignal = v
}

func (o GetTimeSeriesKst200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesKst200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["kst"] = o.Kst
	toSerialize["kst_signal"] = o.KstSignal
	return toSerialize, nil
}

func (o *GetTimeSeriesKst200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"kst",
		"kst_signal",
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

	varGetTimeSeriesKst200ResponseValuesInner := _GetTimeSeriesKst200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesKst200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesKst200ResponseValuesInner(varGetTimeSeriesKst200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesKst200ResponseValuesInner struct {
	value *GetTimeSeriesKst200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesKst200ResponseValuesInner) Get() *GetTimeSeriesKst200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesKst200ResponseValuesInner) Set(val *GetTimeSeriesKst200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesKst200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesKst200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesKst200ResponseValuesInner(val *GetTimeSeriesKst200ResponseValuesInner) *NullableGetTimeSeriesKst200ResponseValuesInner {
	return &NullableGetTimeSeriesKst200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesKst200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesKst200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
