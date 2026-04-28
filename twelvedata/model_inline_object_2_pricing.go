/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject2Pricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject2Pricing{}

// InlineObject2Pricing struct for InlineObject2Pricing
type InlineObject2Pricing struct {
	Var12MonthHigh *float64 `json:"12_month_high,omitempty"`
	Var12MonthLow  *float64 `json:"12_month_low,omitempty"`
	LastMonth      *float64 `json:"last_month,omitempty"`
	Nav            *float64 `json:"nav,omitempty"`
}

// NewInlineObject2Pricing instantiates a new InlineObject2Pricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2Pricing() *InlineObject2Pricing {
	this := InlineObject2Pricing{}
	return &this
}

// NewInlineObject2PricingWithDefaults instantiates a new InlineObject2Pricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2PricingWithDefaults() *InlineObject2Pricing {
	this := InlineObject2Pricing{}
	return &this
}

// GetVar12MonthHigh returns the Var12MonthHigh field value if set, zero value otherwise.
func (o *InlineObject2Pricing) GetVar12MonthHigh() float64 {
	if o == nil || IsNil(o.Var12MonthHigh) {
		var ret float64
		return ret
	}
	return *o.Var12MonthHigh
}

// GetVar12MonthHighOk returns a tuple with the Var12MonthHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Pricing) GetVar12MonthHighOk() (*float64, bool) {
	if o == nil || IsNil(o.Var12MonthHigh) {
		return nil, false
	}
	return o.Var12MonthHigh, true
}

// HasVar12MonthHigh returns a boolean if a field has been set.
func (o *InlineObject2Pricing) HasVar12MonthHigh() bool {
	if o != nil && !IsNil(o.Var12MonthHigh) {
		return true
	}

	return false
}

// SetVar12MonthHigh gets a reference to the given float64 and assigns it to the Var12MonthHigh field.
func (o *InlineObject2Pricing) SetVar12MonthHigh(v float64) {
	o.Var12MonthHigh = &v
}

// GetVar12MonthLow returns the Var12MonthLow field value if set, zero value otherwise.
func (o *InlineObject2Pricing) GetVar12MonthLow() float64 {
	if o == nil || IsNil(o.Var12MonthLow) {
		var ret float64
		return ret
	}
	return *o.Var12MonthLow
}

// GetVar12MonthLowOk returns a tuple with the Var12MonthLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Pricing) GetVar12MonthLowOk() (*float64, bool) {
	if o == nil || IsNil(o.Var12MonthLow) {
		return nil, false
	}
	return o.Var12MonthLow, true
}

// HasVar12MonthLow returns a boolean if a field has been set.
func (o *InlineObject2Pricing) HasVar12MonthLow() bool {
	if o != nil && !IsNil(o.Var12MonthLow) {
		return true
	}

	return false
}

// SetVar12MonthLow gets a reference to the given float64 and assigns it to the Var12MonthLow field.
func (o *InlineObject2Pricing) SetVar12MonthLow(v float64) {
	o.Var12MonthLow = &v
}

// GetLastMonth returns the LastMonth field value if set, zero value otherwise.
func (o *InlineObject2Pricing) GetLastMonth() float64 {
	if o == nil || IsNil(o.LastMonth) {
		var ret float64
		return ret
	}
	return *o.LastMonth
}

// GetLastMonthOk returns a tuple with the LastMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Pricing) GetLastMonthOk() (*float64, bool) {
	if o == nil || IsNil(o.LastMonth) {
		return nil, false
	}
	return o.LastMonth, true
}

// HasLastMonth returns a boolean if a field has been set.
func (o *InlineObject2Pricing) HasLastMonth() bool {
	if o != nil && !IsNil(o.LastMonth) {
		return true
	}

	return false
}

// SetLastMonth gets a reference to the given float64 and assigns it to the LastMonth field.
func (o *InlineObject2Pricing) SetLastMonth(v float64) {
	o.LastMonth = &v
}

// GetNav returns the Nav field value if set, zero value otherwise.
func (o *InlineObject2Pricing) GetNav() float64 {
	if o == nil || IsNil(o.Nav) {
		var ret float64
		return ret
	}
	return *o.Nav
}

// GetNavOk returns a tuple with the Nav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Pricing) GetNavOk() (*float64, bool) {
	if o == nil || IsNil(o.Nav) {
		return nil, false
	}
	return o.Nav, true
}

// HasNav returns a boolean if a field has been set.
func (o *InlineObject2Pricing) HasNav() bool {
	if o != nil && !IsNil(o.Nav) {
		return true
	}

	return false
}

// SetNav gets a reference to the given float64 and assigns it to the Nav field.
func (o *InlineObject2Pricing) SetNav(v float64) {
	o.Nav = &v
}

func (o InlineObject2Pricing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject2Pricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var12MonthHigh) {
		toSerialize["12_month_high"] = o.Var12MonthHigh
	}
	if !IsNil(o.Var12MonthLow) {
		toSerialize["12_month_low"] = o.Var12MonthLow
	}
	if !IsNil(o.LastMonth) {
		toSerialize["last_month"] = o.LastMonth
	}
	if !IsNil(o.Nav) {
		toSerialize["nav"] = o.Nav
	}
	return toSerialize, nil
}

type NullableInlineObject2Pricing struct {
	value *InlineObject2Pricing
	isSet bool
}

func (v NullableInlineObject2Pricing) Get() *InlineObject2Pricing {
	return v.value
}

func (v *NullableInlineObject2Pricing) Set(val *InlineObject2Pricing) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2Pricing) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2Pricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2Pricing(val *InlineObject2Pricing) *NullableInlineObject2Pricing {
	return &NullableInlineObject2Pricing{value: val, isSet: true}
}

func (v NullableInlineObject2Pricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2Pricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
