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

// checks if the GetTimeSeriesBeta200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesBeta200ResponseValuesInner{}

// GetTimeSeriesBeta200ResponseValuesInner struct for GetTimeSeriesBeta200ResponseValuesInner
type GetTimeSeriesBeta200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// Beta value
	Beta string `json:"beta"`
}

type _GetTimeSeriesBeta200ResponseValuesInner GetTimeSeriesBeta200ResponseValuesInner

// NewGetTimeSeriesBeta200ResponseValuesInner instantiates a new GetTimeSeriesBeta200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesBeta200ResponseValuesInner(datetime string, beta string) *GetTimeSeriesBeta200ResponseValuesInner {
	this := GetTimeSeriesBeta200ResponseValuesInner{}
	this.Datetime = datetime
	this.Beta = beta
	return &this
}

// NewGetTimeSeriesBeta200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesBeta200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesBeta200ResponseValuesInnerWithDefaults() *GetTimeSeriesBeta200ResponseValuesInner {
	this := GetTimeSeriesBeta200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesBeta200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBeta200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesBeta200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetBeta returns the Beta field value
func (o *GetTimeSeriesBeta200ResponseValuesInner) GetBeta() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Beta
}

// GetBetaOk returns a tuple with the Beta field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesBeta200ResponseValuesInner) GetBetaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Beta, true
}

// SetBeta sets field value
func (o *GetTimeSeriesBeta200ResponseValuesInner) SetBeta(v string) {
	o.Beta = v
}

func (o GetTimeSeriesBeta200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesBeta200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["beta"] = o.Beta
	return toSerialize, nil
}

func (o *GetTimeSeriesBeta200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"beta",
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

	varGetTimeSeriesBeta200ResponseValuesInner := _GetTimeSeriesBeta200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTimeSeriesBeta200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesBeta200ResponseValuesInner(varGetTimeSeriesBeta200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesBeta200ResponseValuesInner struct {
	value *GetTimeSeriesBeta200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesBeta200ResponseValuesInner) Get() *GetTimeSeriesBeta200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesBeta200ResponseValuesInner) Set(val *GetTimeSeriesBeta200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesBeta200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesBeta200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesBeta200ResponseValuesInner(val *GetTimeSeriesBeta200ResponseValuesInner) *NullableGetTimeSeriesBeta200ResponseValuesInner {
	return &NullableGetTimeSeriesBeta200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesBeta200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesBeta200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
