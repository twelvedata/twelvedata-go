// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesCorrel200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesCorrel200ResponseValuesInner{}

// GetTimeSeriesCorrel200ResponseValuesInner struct for GetTimeSeriesCorrel200ResponseValuesInner
type GetTimeSeriesCorrel200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Correl value
	Correl string `json:"correl"`
}

type _GetTimeSeriesCorrel200ResponseValuesInner GetTimeSeriesCorrel200ResponseValuesInner

// NewGetTimeSeriesCorrel200ResponseValuesInner instantiates a new GetTimeSeriesCorrel200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesCorrel200ResponseValuesInner(datetime string, correl string) *GetTimeSeriesCorrel200ResponseValuesInner {
	this := GetTimeSeriesCorrel200ResponseValuesInner{}
	this.Datetime = datetime
	this.Correl = correl
	return &this
}

// NewGetTimeSeriesCorrel200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCorrel200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesCorrel200ResponseValuesInnerWithDefaults() *GetTimeSeriesCorrel200ResponseValuesInner {
	this := GetTimeSeriesCorrel200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesCorrel200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetCorrel returns the Correl field value
func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetCorrel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Correl
}

// GetCorrelOk returns a tuple with the Correl field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetCorrelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Correl, true
}

// SetCorrel sets field value
func (o *GetTimeSeriesCorrel200ResponseValuesInner) SetCorrel(v string) {
	o.Correl = v
}

func (o GetTimeSeriesCorrel200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesCorrel200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["correl"] = o.Correl
	return toSerialize, nil
}

func (o *GetTimeSeriesCorrel200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"correl",
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

	varGetTimeSeriesCorrel200ResponseValuesInner := _GetTimeSeriesCorrel200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesCorrel200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesCorrel200ResponseValuesInner(varGetTimeSeriesCorrel200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesCorrel200ResponseValuesInner struct {
	value *GetTimeSeriesCorrel200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesCorrel200ResponseValuesInner) Get() *GetTimeSeriesCorrel200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesCorrel200ResponseValuesInner) Set(val *GetTimeSeriesCorrel200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesCorrel200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesCorrel200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesCorrel200ResponseValuesInner(val *GetTimeSeriesCorrel200ResponseValuesInner) *NullableGetTimeSeriesCorrel200ResponseValuesInner {
	return &NullableGetTimeSeriesCorrel200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesCorrel200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesCorrel200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
