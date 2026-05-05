// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses{}

// GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses Costs associated with investing in the mutual fund, including gross and net expense ratios
type GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses struct {
	// Cost of investing in a mutual fund
	ExpenseRatioGross *float64 `json:"expense_ratio_gross,omitempty"`
	// Percentage of mutual fund assets steered toward a fund's operating expenses and fund management fees
	ExpenseRatioNet *float64 `json:"expense_ratio_net,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpensesWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpensesWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses {
	this := GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses{}
	return &this
}

// GetExpenseRatioGross returns the ExpenseRatioGross field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) GetExpenseRatioGross() float64 {
	if o == nil || IsNil(o.ExpenseRatioGross) {
		var ret float64
		return ret
	}
	return *o.ExpenseRatioGross
}

// GetExpenseRatioGrossOk returns a tuple with the ExpenseRatioGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) GetExpenseRatioGrossOk() (*float64, bool) {
	if o == nil || IsNil(o.ExpenseRatioGross) {
		return nil, false
	}
	return o.ExpenseRatioGross, true
}

// HasExpenseRatioGross returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) HasExpenseRatioGross() bool {
	if o != nil && !IsNil(o.ExpenseRatioGross) {
		return true
	}

	return false
}

// SetExpenseRatioGross gets a reference to the given float64 and assigns it to the ExpenseRatioGross field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) SetExpenseRatioGross(v float64) {
	o.ExpenseRatioGross = &v
}

// GetExpenseRatioNet returns the ExpenseRatioNet field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) GetExpenseRatioNet() float64 {
	if o == nil || IsNil(o.ExpenseRatioNet) {
		var ret float64
		return ret
	}
	return *o.ExpenseRatioNet
}

// GetExpenseRatioNetOk returns a tuple with the ExpenseRatioNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) GetExpenseRatioNetOk() (*float64, bool) {
	if o == nil || IsNil(o.ExpenseRatioNet) {
		return nil, false
	}
	return o.ExpenseRatioNet, true
}

// HasExpenseRatioNet returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) HasExpenseRatioNet() bool {
	if o != nil && !IsNil(o.ExpenseRatioNet) {
		return true
	}

	return false
}

// SetExpenseRatioNet gets a reference to the given float64 and assigns it to the ExpenseRatioNet field.
func (o *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) SetExpenseRatioNet(v float64) {
	o.ExpenseRatioNet = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpenseRatioGross) {
		toSerialize["expense_ratio_gross"] = o.ExpenseRatioGross
	}
	if !IsNil(o.ExpenseRatioNet) {
		toSerialize["expense_ratio_net"] = o.ExpenseRatioNet
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses struct {
	value *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) Get() *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) Set(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses(val *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
