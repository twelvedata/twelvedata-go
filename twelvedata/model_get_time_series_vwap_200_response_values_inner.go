// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetTimeSeriesVwap200ResponseValuesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTimeSeriesVwap200ResponseValuesInner{}

// GetTimeSeriesVwap200ResponseValuesInner struct for GetTimeSeriesVwap200ResponseValuesInner
type GetTimeSeriesVwap200ResponseValuesInner struct {
	// Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened
	Datetime string `json:"datetime"`
	// VWAP lower value
	VwapLower *string `json:"vwap_lower,omitempty"`
	// VWAP value
	Vwap string `json:"vwap"`
	// VWAP upper value
	VwapUpper *string `json:"vwap_upper,omitempty"`
}

type _GetTimeSeriesVwap200ResponseValuesInner GetTimeSeriesVwap200ResponseValuesInner

// NewGetTimeSeriesVwap200ResponseValuesInner instantiates a new GetTimeSeriesVwap200ResponseValuesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTimeSeriesVwap200ResponseValuesInner(datetime string, vwap string) *GetTimeSeriesVwap200ResponseValuesInner {
	this := GetTimeSeriesVwap200ResponseValuesInner{}
	this.Datetime = datetime
	this.Vwap = vwap
	return &this
}

// NewGetTimeSeriesVwap200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesVwap200ResponseValuesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTimeSeriesVwap200ResponseValuesInnerWithDefaults() *GetTimeSeriesVwap200ResponseValuesInner {
	this := GetTimeSeriesVwap200ResponseValuesInner{}
	return &this
}

// GetDatetime returns the Datetime field value
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datetime, true
}

// SetDatetime sets field value
func (o *GetTimeSeriesVwap200ResponseValuesInner) SetDatetime(v string) {
	o.Datetime = v
}

// GetVwapLower returns the VwapLower field value if set, zero value otherwise.
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapLower() string {
	if o == nil || IsNil(o.VwapLower) {
		var ret string
		return ret
	}
	return *o.VwapLower
}

// GetVwapLowerOk returns a tuple with the VwapLower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapLowerOk() (*string, bool) {
	if o == nil || IsNil(o.VwapLower) {
		return nil, false
	}
	return o.VwapLower, true
}

// HasVwapLower returns a boolean if a field has been set.
func (o *GetTimeSeriesVwap200ResponseValuesInner) HasVwapLower() bool {
	if o != nil && !IsNil(o.VwapLower) {
		return true
	}

	return false
}

// SetVwapLower gets a reference to the given string and assigns it to the VwapLower field.
func (o *GetTimeSeriesVwap200ResponseValuesInner) SetVwapLower(v string) {
	o.VwapLower = &v
}

// GetVwap returns the Vwap field value
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwap() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vwap
}

// GetVwapOk returns a tuple with the Vwap field value
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vwap, true
}

// SetVwap sets field value
func (o *GetTimeSeriesVwap200ResponseValuesInner) SetVwap(v string) {
	o.Vwap = v
}

// GetVwapUpper returns the VwapUpper field value if set, zero value otherwise.
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapUpper() string {
	if o == nil || IsNil(o.VwapUpper) {
		var ret string
		return ret
	}
	return *o.VwapUpper
}

// GetVwapUpperOk returns a tuple with the VwapUpper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTimeSeriesVwap200ResponseValuesInner) GetVwapUpperOk() (*string, bool) {
	if o == nil || IsNil(o.VwapUpper) {
		return nil, false
	}
	return o.VwapUpper, true
}

// HasVwapUpper returns a boolean if a field has been set.
func (o *GetTimeSeriesVwap200ResponseValuesInner) HasVwapUpper() bool {
	if o != nil && !IsNil(o.VwapUpper) {
		return true
	}

	return false
}

// SetVwapUpper gets a reference to the given string and assigns it to the VwapUpper field.
func (o *GetTimeSeriesVwap200ResponseValuesInner) SetVwapUpper(v string) {
	o.VwapUpper = &v
}

func (o GetTimeSeriesVwap200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTimeSeriesVwap200ResponseValuesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datetime"] = o.Datetime
	if !IsNil(o.VwapLower) {
		toSerialize["vwap_lower"] = o.VwapLower
	}
	toSerialize["vwap"] = o.Vwap
	if !IsNil(o.VwapUpper) {
		toSerialize["vwap_upper"] = o.VwapUpper
	}
	return toSerialize, nil
}

func (o *GetTimeSeriesVwap200ResponseValuesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"datetime",
		"vwap",
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

	varGetTimeSeriesVwap200ResponseValuesInner := _GetTimeSeriesVwap200ResponseValuesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetTimeSeriesVwap200ResponseValuesInner)

	if err != nil {
		return err
	}

	*o = GetTimeSeriesVwap200ResponseValuesInner(varGetTimeSeriesVwap200ResponseValuesInner)

	return err
}

type NullableGetTimeSeriesVwap200ResponseValuesInner struct {
	value *GetTimeSeriesVwap200ResponseValuesInner
	isSet bool
}

func (v NullableGetTimeSeriesVwap200ResponseValuesInner) Get() *GetTimeSeriesVwap200ResponseValuesInner {
	return v.value
}

func (v *NullableGetTimeSeriesVwap200ResponseValuesInner) Set(val *GetTimeSeriesVwap200ResponseValuesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeSeriesVwap200ResponseValuesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeSeriesVwap200ResponseValuesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeSeriesVwap200ResponseValuesInner(val *GetTimeSeriesVwap200ResponseValuesInner) *NullableGetTimeSeriesVwap200ResponseValuesInner {
	return &NullableGetTimeSeriesVwap200ResponseValuesInner{value: val, isSet: true}
}

func (v NullableGetTimeSeriesVwap200ResponseValuesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeSeriesVwap200ResponseValuesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
