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

// checks if the GetTimeSeriesMcGinleyDynamic200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMcGinleyDynamic200ResponseValuesInner{}

// GetTimeSeriesMcGinleyDynamic200ResponseValuesInner struct for GetTimeSeriesMcGinleyDynamic200ResponseValuesInner
type GetTimeSeriesMcGinleyDynamic200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// McGinley Dynamic value
	McginleyDynamic string `json:"mcginley_dynamic"`
}

type _GetTimeSeriesMcGinleyDynamic200ResponseValuesInner GetTimeSeriesMcGinleyDynamic200ResponseValuesInner

// NewGetTimeSeriesMcGinleyDynamic200ResponseValuesInner instantiates a new GetTimeSeriesMcGinleyDynamic200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMcGinleyDynamic200ResponseValuesInner(datetime string, mcginleyDynamic string) *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner {
	this := GetTimeSeriesMcGinleyDynamic200ResponseValuesInner{}
	this.Datetime = datetime
	this.McginleyDynamic = mcginleyDynamic
	return &this
}

// NewGetTimeSeriesMcGinleyDynamic200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMcGinleyDynamic200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMcGinleyDynamic200ResponseValuesInnerWithDefaults() *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner {
	this := GetTimeSeriesMcGinleyDynamic200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMcginleyDynamic returns the McginleyDynamic field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) GetMcginleyDynamic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.McginleyDynamic
}

// GetMcginleyDynamicOk returns a tuple with the McginleyDynamic field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) GetMcginleyDynamicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.McginleyDynamic, true
}

// SetMcginleyDynamic sets field value
func (o *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) SetMcginleyDynamic(v string) {
	o.McginleyDynamic = v
}

func (o GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["mcginley_dynamic"] = o.McginleyDynamic
	return toSerialize, nil
}

func (o *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"mcginley_dynamic",
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

	varGetTimeSeriesMcGinleyDynamic200ResponseValuesInner := _GetTimeSeriesMcGinleyDynamic200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMcGinleyDynamic200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMcGinleyDynamic200ResponseValuesInner(varGetTimeSeriesMcGinleyDynamic200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner struct {
	value *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner) Get() *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner) Set(val *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner(val *GetTimeSeriesMcGinleyDynamic200ResponseValuesInner) *NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner {
	return &NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMcGinleyDynamic200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
