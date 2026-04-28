/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{}

// GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing Pricing information for the mutual fund
type GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing struct {
	// Net Asset Value: fund value minus liabilities
	Nav *float64 `json:"nav,omitempty"`
	// Lowest price of the fund over the last year
	Var12MonthLow *float64 `json:"12_month_low,omitempty"`
	// Highest price of the fund over the last year
	Var12MonthHigh *float64 `json:"12_month_high,omitempty"`
	// Fund price at the end of the last month
	LastMonth *float64 `json:"last_month,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricingWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricingWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{}
	return &this
}

// GetNav returns the Nav field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetNav() float64 {
	if o == nil || IsNil(o.Nav) {
		var ret float64
		return ret
	}
	return *o.Nav
}

// GetNavOk returns a tuple with the Nav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetNavOk() (*float64, bool) {
	if o == nil || IsNil(o.Nav) {
		return nil, false
	}
	return o.Nav, true
}

// HasNav returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasNav() bool {
	if o != nil && !IsNil(o.Nav) {
		return true
	}

	return false
}

// SetNav gets a reference to the given float64 and assigns it to the Nav field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetNav(v float64) {
	o.Nav = &v
}

// GetVar12MonthLow returns the Var12MonthLow field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthLow() float64 {
	if o == nil || IsNil(o.Var12MonthLow) {
		var ret float64
		return ret
	}
	return *o.Var12MonthLow
}

// GetVar12MonthLowOk returns a tuple with the Var12MonthLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthLowOk() (*float64, bool) {
	if o == nil || IsNil(o.Var12MonthLow) {
		return nil, false
	}
	return o.Var12MonthLow, true
}

// HasVar12MonthLow returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasVar12MonthLow() bool {
	if o != nil && !IsNil(o.Var12MonthLow) {
		return true
	}

	return false
}

// SetVar12MonthLow gets a reference to the given float64 and assigns it to the Var12MonthLow field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetVar12MonthLow(v float64) {
	o.Var12MonthLow = &v
}

// GetVar12MonthHigh returns the Var12MonthHigh field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthHigh() float64 {
	if o == nil || IsNil(o.Var12MonthHigh) {
		var ret float64
		return ret
	}
	return *o.Var12MonthHigh
}

// GetVar12MonthHighOk returns a tuple with the Var12MonthHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetVar12MonthHighOk() (*float64, bool) {
	if o == nil || IsNil(o.Var12MonthHigh) {
		return nil, false
	}
	return o.Var12MonthHigh, true
}

// HasVar12MonthHigh returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasVar12MonthHigh() bool {
	if o != nil && !IsNil(o.Var12MonthHigh) {
		return true
	}

	return false
}

// SetVar12MonthHigh gets a reference to the given float64 and assigns it to the Var12MonthHigh field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetVar12MonthHigh(v float64) {
	o.Var12MonthHigh = &v
}

// GetLastMonth returns the LastMonth field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetLastMonth() float64 {
	if o == nil || IsNil(o.LastMonth) {
		var ret float64
		return ret
	}
	return *o.LastMonth
}

// GetLastMonthOk returns a tuple with the LastMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) GetLastMonthOk() (*float64, bool) {
	if o == nil || IsNil(o.LastMonth) {
		return nil, false
	}
	return o.LastMonth, true
}

// HasLastMonth returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) HasLastMonth() bool {
	if o != nil && !IsNil(o.LastMonth) {
		return true
	}

	return false
}

// SetLastMonth gets a reference to the given float64 and assigns it to the LastMonth field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) SetLastMonth(v float64) {
	o.LastMonth = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nav) {
		toSerialize["nav"] = o.Nav
	}
	if !IsNil(o.Var12MonthLow) {
		toSerialize["12_month_low"] = o.Var12MonthLow
	}
	if !IsNil(o.Var12MonthHigh) {
		toSerialize["12_month_high"] = o.Var12MonthHigh
	}
	if !IsNil(o.LastMonth) {
		toSerialize["last_month"] = o.LastMonth
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing struct {
	value *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) Get() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) Set(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
