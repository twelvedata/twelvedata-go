// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner{}

// GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner struct for GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner
type GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner struct {
	// Period of a load adjusted return
	Period *string `json:"period,omitempty"`
	// Actual return (%) an investor sees after accounting for fees and sales charges are deducted from a mutual fund's performance
	Return *float64 `json:"return,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner() *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) SetPeriod(v string) {
	o.Period = &v
}

// GetReturn returns the Return field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) GetReturn() float64 {
	if o == nil || IsNil(o.Return) {
		var ret float64
		return ret
	}
	return *o.Return
}

// GetReturnOk returns a tuple with the Return field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) GetReturnOk() (*float64, bool) {
	if o == nil || IsNil(o.Return) {
		return nil, false
	}
	return o.Return, true
}

// HasReturn returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) HasReturn() bool {
	if o != nil && !IsNil(o.Return) {
		return true
	}

	return false
}

// SetReturn gets a reference to the given float64 and assigns it to the Return field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) SetReturn(v float64) {
	o.Return = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Return) {
		toSerialize["return"] = o.Return
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) Get() *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) Set(val *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner(val *GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
