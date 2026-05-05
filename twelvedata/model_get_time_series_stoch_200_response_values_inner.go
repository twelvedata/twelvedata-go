// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesStoch200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStoch200ResponseValuesInner{}

// GetTimeSeriesStoch200ResponseValuesInner struct for GetTimeSeriesStoch200ResponseValuesInner
type GetTimeSeriesStoch200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// slow_k value
	SlowK string `json:"slow_k"`
	// slow_d value
	SlowD string `json:"slow_d"`
}

type _GetTimeSeriesStoch200ResponseValuesInner GetTimeSeriesStoch200ResponseValuesInner

// NewGetTimeSeriesStoch200ResponseValuesInner instantiates a new GetTimeSeriesStoch200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStoch200ResponseValuesInner(datetime string, slowK string, slowD string) *GetTimeSeriesStoch200ResponseValuesInner {
	this := GetTimeSeriesStoch200ResponseValuesInner{}
	this.Datetime = datetime
	this.SlowK = slowK
	this.SlowD = slowD
	return &this
}

// NewGetTimeSeriesStoch200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesStoch200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStoch200ResponseValuesInnerWithDefaults() *GetTimeSeriesStoch200ResponseValuesInner {
	this := GetTimeSeriesStoch200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesStoch200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesStoch200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetSlowK returns the SlowK field value
func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowK() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlowK
}

// GetSlowKOk returns a tuple with the SlowK field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowKOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowK, true
}

// SetSlowK sets field value
func (o *GetTimeSeriesStoch200ResponseValuesInner) SetSlowK(v string) {
	o.SlowK = v
}

// GetSlowD returns the SlowD field value
func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowD() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SlowD
}

// GetSlowDOk returns a tuple with the SlowD field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowD, true
}

// SetSlowD sets field value
func (o *GetTimeSeriesStoch200ResponseValuesInner) SetSlowD(v string) {
	o.SlowD = v
}

func (o GetTimeSeriesStoch200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStoch200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["slow_k"] = o.SlowK
	toSerialize["slow_d"] = o.SlowD
	return toSerialize, nil
}

func (o *GetTimeSeriesStoch200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"slow_k",
		"slow_d",
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

	varGetTimeSeriesStoch200ResponseValuesInner := _GetTimeSeriesStoch200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStoch200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStoch200ResponseValuesInner(varGetTimeSeriesStoch200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesStoch200ResponseValuesInner struct {
	value *GetTimeSeriesStoch200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesStoch200ResponseValuesInner) Get() *GetTimeSeriesStoch200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesStoch200ResponseValuesInner) Set(val *GetTimeSeriesStoch200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStoch200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStoch200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStoch200ResponseValuesInner(val *GetTimeSeriesStoch200ResponseValuesInner) *NullableGetTimeSeriesStoch200ResponseValuesInner {
	return &NullableGetTimeSeriesStoch200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStoch200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStoch200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
