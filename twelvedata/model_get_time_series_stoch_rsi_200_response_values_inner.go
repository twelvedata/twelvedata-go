// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesStochRsi200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesStochRsi200ResponseValuesInner{}

// GetTimeSeriesStochRsi200ResponseValuesInner struct for GetTimeSeriesStochRsi200ResponseValuesInner
type GetTimeSeriesStochRsi200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// K value
	K string `json:"k"`
	// D value
	D string `json:"d"`
}

type _GetTimeSeriesStochRsi200ResponseValuesInner GetTimeSeriesStochRsi200ResponseValuesInner

// NewGetTimeSeriesStochRsi200ResponseValuesInner instantiates a new GetTimeSeriesStochRsi200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesStochRsi200ResponseValuesInner(datetime string, k string, d string) *GetTimeSeriesStochRsi200ResponseValuesInner {
	this := GetTimeSeriesStochRsi200ResponseValuesInner{}
	this.Datetime = datetime
	this.K = k
	this.D = d
	return &this
}

// NewGetTimeSeriesStochRsi200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesStochRsi200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesStochRsi200ResponseValuesInnerWithDefaults() *GetTimeSeriesStochRsi200ResponseValuesInner {
	this := GetTimeSeriesStochRsi200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetK returns the K field value
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) GetK() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.K
}

// GetKOk returns a tuple with the K field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) GetKOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.K, true
}

// SetK sets field value
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) SetK(v string) {
	o.K = v
}

// GetD returns the D field value
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) GetD() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.D
}

// GetDOk returns a tuple with the D field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) GetDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.D, true
}

// SetD sets field value
func (o *GetTimeSeriesStochRsi200ResponseValuesInner) SetD(v string) {
	o.D = v
}

func (o GetTimeSeriesStochRsi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesStochRsi200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["k"] = o.K
	toSerialize["d"] = o.D
	return toSerialize, nil
}

func (o *GetTimeSeriesStochRsi200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"k",
		"d",
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

	varGetTimeSeriesStochRsi200ResponseValuesInner := _GetTimeSeriesStochRsi200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesStochRsi200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesStochRsi200ResponseValuesInner(varGetTimeSeriesStochRsi200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesStochRsi200ResponseValuesInner struct {
	value *GetTimeSeriesStochRsi200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesStochRsi200ResponseValuesInner) Get() *GetTimeSeriesStochRsi200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesStochRsi200ResponseValuesInner) Set(val *GetTimeSeriesStochRsi200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesStochRsi200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesStochRsi200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesStochRsi200ResponseValuesInner(val *GetTimeSeriesStochRsi200ResponseValuesInner) *NullableGetTimeSeriesStochRsi200ResponseValuesInner {
	return &NullableGetTimeSeriesStochRsi200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesStochRsi200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesStochRsi200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
