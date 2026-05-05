// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesMaxIndex200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesMaxIndex200ResponseValuesInner{}

// GetTimeSeriesMaxIndex200ResponseValuesInner struct for GetTimeSeriesMaxIndex200ResponseValuesInner
type GetTimeSeriesMaxIndex200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// maxidx value
	Maxidx string `json:"maxidx"`
}

type _GetTimeSeriesMaxIndex200ResponseValuesInner GetTimeSeriesMaxIndex200ResponseValuesInner

// NewGetTimeSeriesMaxIndex200ResponseValuesInner instantiates a new GetTimeSeriesMaxIndex200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesMaxIndex200ResponseValuesInner(datetime string, maxidx string) *GetTimeSeriesMaxIndex200ResponseValuesInner {
	this := GetTimeSeriesMaxIndex200ResponseValuesInner{}
	this.Datetime = datetime
	this.Maxidx = maxidx
	return &this
}

// NewGetTimeSeriesMaxIndex200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMaxIndex200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesMaxIndex200ResponseValuesInnerWithDefaults() *GetTimeSeriesMaxIndex200ResponseValuesInner {
	this := GetTimeSeriesMaxIndex200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetMaxidx returns the Maxidx field value
func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetMaxidx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Maxidx
}

// GetMaxidxOk returns a tuple with the Maxidx field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetMaxidxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maxidx, true
}

// SetMaxidx sets field value
func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) SetMaxidx(v string) {
	o.Maxidx = v
}

func (o GetTimeSeriesMaxIndex200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesMaxIndex200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	toSerialize["maxidx"] = o.Maxidx
	return toSerialize, nil
}

func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"maxidx",
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

	varGetTimeSeriesMaxIndex200ResponseValuesInner := _GetTimeSeriesMaxIndex200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesMaxIndex200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesMaxIndex200ResponseValuesInner(varGetTimeSeriesMaxIndex200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesMaxIndex200ResponseValuesInner struct {
	value *GetTimeSeriesMaxIndex200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesMaxIndex200ResponseValuesInner) Get() *GetTimeSeriesMaxIndex200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesMaxIndex200ResponseValuesInner) Set(val *GetTimeSeriesMaxIndex200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesMaxIndex200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesMaxIndex200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesMaxIndex200ResponseValuesInner(val *GetTimeSeriesMaxIndex200ResponseValuesInner) *NullableGetTimeSeriesMaxIndex200ResponseValuesInner {
	return &NullableGetTimeSeriesMaxIndex200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesMaxIndex200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesMaxIndex200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
