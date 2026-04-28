/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities{}

// GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities Current liabilities section
type GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities struct {
	// Refers to an account within the general ledger that represents an obligation to pay off a short-term debt to creditors or suppliers
	AccountsPayable *float64 `json:"accounts_payable,omitempty"`
	// Represents payments that a company is obligated to pay in the future for which goods and services have already been delivered
	AccruedExpenses *float64 `json:"accrued_expenses,omitempty"`
	// Represents current debt and capital lease obligations
	ShortTermDebt *float64 `json:"short_term_debt,omitempty"`
	// Represents advance payments a company receives for products or services that are to be delivered or performed in the future
	DeferredRevenue *float64 `json:"deferred_revenue,omitempty"`
	// Represents taxes due to the government within one year
	TaxPayable *float64 `json:"tax_payable,omitempty"`
	// Represents to pensions to be paid out
	Pensions *float64 `json:"pensions,omitempty"`
	// Represents other current liabilities
	OtherCurrentLiabilities *float64 `json:"other_current_liabilities,omitempty"`
	// Represents total current liabilities
	TotalCurrentLiabilities *float64 `json:"total_current_liabilities,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilitiesWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilitiesWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities{}
	return &this
}

// GetAccountsPayable returns the AccountsPayable field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetAccountsPayable() float64 {
	if o == nil || IsNil(o.AccountsPayable) {
		var ret float64
		return ret
	}
	return *o.AccountsPayable
}

// GetAccountsPayableOk returns a tuple with the AccountsPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetAccountsPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccountsPayable) {
		return nil, false
	}
	return o.AccountsPayable, true
}

// HasAccountsPayable returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasAccountsPayable() bool {
	if o != nil && !IsNil(o.AccountsPayable) {
		return true
	}

	return false
}

// SetAccountsPayable gets a reference to the given float64 and assigns it to the AccountsPayable field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetAccountsPayable(v float64) {
	o.AccountsPayable = &v
}

// GetAccruedExpenses returns the AccruedExpenses field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetAccruedExpenses() float64 {
	if o == nil || IsNil(o.AccruedExpenses) {
		var ret float64
		return ret
	}
	return *o.AccruedExpenses
}

// GetAccruedExpensesOk returns a tuple with the AccruedExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetAccruedExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.AccruedExpenses) {
		return nil, false
	}
	return o.AccruedExpenses, true
}

// HasAccruedExpenses returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasAccruedExpenses() bool {
	if o != nil && !IsNil(o.AccruedExpenses) {
		return true
	}

	return false
}

// SetAccruedExpenses gets a reference to the given float64 and assigns it to the AccruedExpenses field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetAccruedExpenses(v float64) {
	o.AccruedExpenses = &v
}

// GetShortTermDebt returns the ShortTermDebt field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetShortTermDebt() float64 {
	if o == nil || IsNil(o.ShortTermDebt) {
		var ret float64
		return ret
	}
	return *o.ShortTermDebt
}

// GetShortTermDebtOk returns a tuple with the ShortTermDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetShortTermDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortTermDebt) {
		return nil, false
	}
	return o.ShortTermDebt, true
}

// HasShortTermDebt returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasShortTermDebt() bool {
	if o != nil && !IsNil(o.ShortTermDebt) {
		return true
	}

	return false
}

// SetShortTermDebt gets a reference to the given float64 and assigns it to the ShortTermDebt field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetShortTermDebt(v float64) {
	o.ShortTermDebt = &v
}

// GetDeferredRevenue returns the DeferredRevenue field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetDeferredRevenue() float64 {
	if o == nil || IsNil(o.DeferredRevenue) {
		var ret float64
		return ret
	}
	return *o.DeferredRevenue
}

// GetDeferredRevenueOk returns a tuple with the DeferredRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetDeferredRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.DeferredRevenue) {
		return nil, false
	}
	return o.DeferredRevenue, true
}

// HasDeferredRevenue returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasDeferredRevenue() bool {
	if o != nil && !IsNil(o.DeferredRevenue) {
		return true
	}

	return false
}

// SetDeferredRevenue gets a reference to the given float64 and assigns it to the DeferredRevenue field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetDeferredRevenue(v float64) {
	o.DeferredRevenue = &v
}

// GetTaxPayable returns the TaxPayable field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetTaxPayable() float64 {
	if o == nil || IsNil(o.TaxPayable) {
		var ret float64
		return ret
	}
	return *o.TaxPayable
}

// GetTaxPayableOk returns a tuple with the TaxPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetTaxPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxPayable) {
		return nil, false
	}
	return o.TaxPayable, true
}

// HasTaxPayable returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasTaxPayable() bool {
	if o != nil && !IsNil(o.TaxPayable) {
		return true
	}

	return false
}

// SetTaxPayable gets a reference to the given float64 and assigns it to the TaxPayable field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetTaxPayable(v float64) {
	o.TaxPayable = &v
}

// GetPensions returns the Pensions field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetPensions() float64 {
	if o == nil || IsNil(o.Pensions) {
		var ret float64
		return ret
	}
	return *o.Pensions
}

// GetPensionsOk returns a tuple with the Pensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetPensionsOk() (*float64, bool) {
	if o == nil || IsNil(o.Pensions) {
		return nil, false
	}
	return o.Pensions, true
}

// HasPensions returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasPensions() bool {
	if o != nil && !IsNil(o.Pensions) {
		return true
	}

	return false
}

// SetPensions gets a reference to the given float64 and assigns it to the Pensions field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetPensions(v float64) {
	o.Pensions = &v
}

// GetOtherCurrentLiabilities returns the OtherCurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetOtherCurrentLiabilities() float64 {
	if o == nil || IsNil(o.OtherCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentLiabilities
}

// GetOtherCurrentLiabilitiesOk returns a tuple with the OtherCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetOtherCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentLiabilities) {
		return nil, false
	}
	return o.OtherCurrentLiabilities, true
}

// HasOtherCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasOtherCurrentLiabilities() bool {
	if o != nil && !IsNil(o.OtherCurrentLiabilities) {
		return true
	}

	return false
}

// SetOtherCurrentLiabilities gets a reference to the given float64 and assigns it to the OtherCurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetOtherCurrentLiabilities(v float64) {
	o.OtherCurrentLiabilities = &v
}

// GetTotalCurrentLiabilities returns the TotalCurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetTotalCurrentLiabilities() float64 {
	if o == nil || IsNil(o.TotalCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.TotalCurrentLiabilities
}

// GetTotalCurrentLiabilitiesOk returns a tuple with the TotalCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) GetTotalCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCurrentLiabilities) {
		return nil, false
	}
	return o.TotalCurrentLiabilities, true
}

// HasTotalCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) HasTotalCurrentLiabilities() bool {
	if o != nil && !IsNil(o.TotalCurrentLiabilities) {
		return true
	}

	return false
}

// SetTotalCurrentLiabilities gets a reference to the given float64 and assigns it to the TotalCurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) SetTotalCurrentLiabilities(v float64) {
	o.TotalCurrentLiabilities = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountsPayable) {
		toSerialize["accounts_payable"] = o.AccountsPayable
	}
	if !IsNil(o.AccruedExpenses) {
		toSerialize["accrued_expenses"] = o.AccruedExpenses
	}
	if !IsNil(o.ShortTermDebt) {
		toSerialize["short_term_debt"] = o.ShortTermDebt
	}
	if !IsNil(o.DeferredRevenue) {
		toSerialize["deferred_revenue"] = o.DeferredRevenue
	}
	if !IsNil(o.TaxPayable) {
		toSerialize["tax_payable"] = o.TaxPayable
	}
	if !IsNil(o.Pensions) {
		toSerialize["pensions"] = o.Pensions
	}
	if !IsNil(o.OtherCurrentLiabilities) {
		toSerialize["other_current_liabilities"] = o.OtherCurrentLiabilities
	}
	if !IsNil(o.TotalCurrentLiabilities) {
		toSerialize["total_current_liabilities"] = o.TotalCurrentLiabilities
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) Get() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesCurrentLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
