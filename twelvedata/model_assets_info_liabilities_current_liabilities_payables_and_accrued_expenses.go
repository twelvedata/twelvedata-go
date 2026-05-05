// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses{}

// AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses Payables and accrued expenses information
type AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses struct {
	// Total payables and accrued expenses
	TotalPayablesAndAccruedExpenses *float64 `json:"total_payables_and_accrued_expenses,omitempty"`
	// Accounts payable
	AccountsPayable *float64 `json:"accounts_payable,omitempty"`
	// Current accrued expenses
	CurrentAccruedExpenses *float64 `json:"current_accrued_expenses,omitempty"`
	// Interest payable
	InterestPayable *float64 `json:"interest_payable,omitempty"`
	// Payables
	Payables *float64 `json:"payables,omitempty"`
	// Other payable
	OtherPayable *float64 `json:"other_payable,omitempty"`
	// Due to related parties current
	DueToRelatedPartiesCurrent *float64 `json:"due_to_related_parties_current,omitempty"`
	// Dividends payable
	DividendsPayable *float64 `json:"dividends_payable,omitempty"`
	// Total tax payable
	TotalTaxPayable *float64 `json:"total_tax_payable,omitempty"`
	// Income tax payable
	IncomeTaxPayable *float64 `json:"income_tax_payable,omitempty"`
}

// NewAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses instantiates a new AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses() *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses {
	this := AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses{}
	return &this
}

// NewAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpensesWithDefaults instantiates a new AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpensesWithDefaults() *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses {
	this := AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses{}
	return &this
}

// GetTotalPayablesAndAccruedExpenses returns the TotalPayablesAndAccruedExpenses field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetTotalPayablesAndAccruedExpenses() float64 {
	if o == nil || IsNil(o.TotalPayablesAndAccruedExpenses) {
		var ret float64
		return ret
	}
	return *o.TotalPayablesAndAccruedExpenses
}

// GetTotalPayablesAndAccruedExpensesOk returns a tuple with the TotalPayablesAndAccruedExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetTotalPayablesAndAccruedExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalPayablesAndAccruedExpenses) {
		return nil, false
	}
	return o.TotalPayablesAndAccruedExpenses, true
}

// HasTotalPayablesAndAccruedExpenses returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasTotalPayablesAndAccruedExpenses() bool {
	if o != nil && !IsNil(o.TotalPayablesAndAccruedExpenses) {
		return true
	}

	return false
}

// SetTotalPayablesAndAccruedExpenses gets a reference to the given float64 and assigns it to the TotalPayablesAndAccruedExpenses field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetTotalPayablesAndAccruedExpenses(v float64) {
	o.TotalPayablesAndAccruedExpenses = &v
}

// GetAccountsPayable returns the AccountsPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetAccountsPayable() float64 {
	if o == nil || IsNil(o.AccountsPayable) {
		var ret float64
		return ret
	}
	return *o.AccountsPayable
}

// GetAccountsPayableOk returns a tuple with the AccountsPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetAccountsPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccountsPayable) {
		return nil, false
	}
	return o.AccountsPayable, true
}

// HasAccountsPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasAccountsPayable() bool {
	if o != nil && !IsNil(o.AccountsPayable) {
		return true
	}

	return false
}

// SetAccountsPayable gets a reference to the given float64 and assigns it to the AccountsPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetAccountsPayable(v float64) {
	o.AccountsPayable = &v
}

// GetCurrentAccruedExpenses returns the CurrentAccruedExpenses field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetCurrentAccruedExpenses() float64 {
	if o == nil || IsNil(o.CurrentAccruedExpenses) {
		var ret float64
		return ret
	}
	return *o.CurrentAccruedExpenses
}

// GetCurrentAccruedExpensesOk returns a tuple with the CurrentAccruedExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetCurrentAccruedExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentAccruedExpenses) {
		return nil, false
	}
	return o.CurrentAccruedExpenses, true
}

// HasCurrentAccruedExpenses returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasCurrentAccruedExpenses() bool {
	if o != nil && !IsNil(o.CurrentAccruedExpenses) {
		return true
	}

	return false
}

// SetCurrentAccruedExpenses gets a reference to the given float64 and assigns it to the CurrentAccruedExpenses field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetCurrentAccruedExpenses(v float64) {
	o.CurrentAccruedExpenses = &v
}

// GetInterestPayable returns the InterestPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetInterestPayable() float64 {
	if o == nil || IsNil(o.InterestPayable) {
		var ret float64
		return ret
	}
	return *o.InterestPayable
}

// GetInterestPayableOk returns a tuple with the InterestPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetInterestPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestPayable) {
		return nil, false
	}
	return o.InterestPayable, true
}

// HasInterestPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasInterestPayable() bool {
	if o != nil && !IsNil(o.InterestPayable) {
		return true
	}

	return false
}

// SetInterestPayable gets a reference to the given float64 and assigns it to the InterestPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetInterestPayable(v float64) {
	o.InterestPayable = &v
}

// GetPayables returns the Payables field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetPayables() float64 {
	if o == nil || IsNil(o.Payables) {
		var ret float64
		return ret
	}
	return *o.Payables
}

// GetPayablesOk returns a tuple with the Payables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetPayablesOk() (*float64, bool) {
	if o == nil || IsNil(o.Payables) {
		return nil, false
	}
	return o.Payables, true
}

// HasPayables returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasPayables() bool {
	if o != nil && !IsNil(o.Payables) {
		return true
	}

	return false
}

// SetPayables gets a reference to the given float64 and assigns it to the Payables field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetPayables(v float64) {
	o.Payables = &v
}

// GetOtherPayable returns the OtherPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetOtherPayable() float64 {
	if o == nil || IsNil(o.OtherPayable) {
		var ret float64
		return ret
	}
	return *o.OtherPayable
}

// GetOtherPayableOk returns a tuple with the OtherPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetOtherPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherPayable) {
		return nil, false
	}
	return o.OtherPayable, true
}

// HasOtherPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasOtherPayable() bool {
	if o != nil && !IsNil(o.OtherPayable) {
		return true
	}

	return false
}

// SetOtherPayable gets a reference to the given float64 and assigns it to the OtherPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetOtherPayable(v float64) {
	o.OtherPayable = &v
}

// GetDueToRelatedPartiesCurrent returns the DueToRelatedPartiesCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetDueToRelatedPartiesCurrent() float64 {
	if o == nil || IsNil(o.DueToRelatedPartiesCurrent) {
		var ret float64
		return ret
	}
	return *o.DueToRelatedPartiesCurrent
}

// GetDueToRelatedPartiesCurrentOk returns a tuple with the DueToRelatedPartiesCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetDueToRelatedPartiesCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.DueToRelatedPartiesCurrent) {
		return nil, false
	}
	return o.DueToRelatedPartiesCurrent, true
}

// HasDueToRelatedPartiesCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasDueToRelatedPartiesCurrent() bool {
	if o != nil && !IsNil(o.DueToRelatedPartiesCurrent) {
		return true
	}

	return false
}

// SetDueToRelatedPartiesCurrent gets a reference to the given float64 and assigns it to the DueToRelatedPartiesCurrent field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetDueToRelatedPartiesCurrent(v float64) {
	o.DueToRelatedPartiesCurrent = &v
}

// GetDividendsPayable returns the DividendsPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetDividendsPayable() float64 {
	if o == nil || IsNil(o.DividendsPayable) {
		var ret float64
		return ret
	}
	return *o.DividendsPayable
}

// GetDividendsPayableOk returns a tuple with the DividendsPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetDividendsPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.DividendsPayable) {
		return nil, false
	}
	return o.DividendsPayable, true
}

// HasDividendsPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasDividendsPayable() bool {
	if o != nil && !IsNil(o.DividendsPayable) {
		return true
	}

	return false
}

// SetDividendsPayable gets a reference to the given float64 and assigns it to the DividendsPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetDividendsPayable(v float64) {
	o.DividendsPayable = &v
}

// GetTotalTaxPayable returns the TotalTaxPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetTotalTaxPayable() float64 {
	if o == nil || IsNil(o.TotalTaxPayable) {
		var ret float64
		return ret
	}
	return *o.TotalTaxPayable
}

// GetTotalTaxPayableOk returns a tuple with the TotalTaxPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetTotalTaxPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTaxPayable) {
		return nil, false
	}
	return o.TotalTaxPayable, true
}

// HasTotalTaxPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasTotalTaxPayable() bool {
	if o != nil && !IsNil(o.TotalTaxPayable) {
		return true
	}

	return false
}

// SetTotalTaxPayable gets a reference to the given float64 and assigns it to the TotalTaxPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetTotalTaxPayable(v float64) {
	o.TotalTaxPayable = &v
}

// GetIncomeTaxPayable returns the IncomeTaxPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetIncomeTaxPayable() float64 {
	if o == nil || IsNil(o.IncomeTaxPayable) {
		var ret float64
		return ret
	}
	return *o.IncomeTaxPayable
}

// GetIncomeTaxPayableOk returns a tuple with the IncomeTaxPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) GetIncomeTaxPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.IncomeTaxPayable) {
		return nil, false
	}
	return o.IncomeTaxPayable, true
}

// HasIncomeTaxPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) HasIncomeTaxPayable() bool {
	if o != nil && !IsNil(o.IncomeTaxPayable) {
		return true
	}

	return false
}

// SetIncomeTaxPayable gets a reference to the given float64 and assigns it to the IncomeTaxPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) SetIncomeTaxPayable(v float64) {
	o.IncomeTaxPayable = &v
}

func (o AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalPayablesAndAccruedExpenses) {
		toSerialize["total_payables_and_accrued_expenses"] = o.TotalPayablesAndAccruedExpenses
	}
	if !IsNil(o.AccountsPayable) {
		toSerialize["accounts_payable"] = o.AccountsPayable
	}
	if !IsNil(o.CurrentAccruedExpenses) {
		toSerialize["current_accrued_expenses"] = o.CurrentAccruedExpenses
	}
	if !IsNil(o.InterestPayable) {
		toSerialize["interest_payable"] = o.InterestPayable
	}
	if !IsNil(o.Payables) {
		toSerialize["payables"] = o.Payables
	}
	if !IsNil(o.OtherPayable) {
		toSerialize["other_payable"] = o.OtherPayable
	}
	if !IsNil(o.DueToRelatedPartiesCurrent) {
		toSerialize["due_to_related_parties_current"] = o.DueToRelatedPartiesCurrent
	}
	if !IsNil(o.DividendsPayable) {
		toSerialize["dividends_payable"] = o.DividendsPayable
	}
	if !IsNil(o.TotalTaxPayable) {
		toSerialize["total_tax_payable"] = o.TotalTaxPayable
	}
	if !IsNil(o.IncomeTaxPayable) {
		toSerialize["income_tax_payable"] = o.IncomeTaxPayable
	}
	return toSerialize, nil
}

type NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses struct {
	value *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses
	isSet bool
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) Get() *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses {
	return v.value
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) Set(val *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses(val *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) *NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses {
	return &NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses{value: val, isSet: true}
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
