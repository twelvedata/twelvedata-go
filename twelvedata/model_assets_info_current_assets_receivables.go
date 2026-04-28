/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoCurrentAssetsReceivables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoCurrentAssetsReceivables{}

// AssetsInfoCurrentAssetsReceivables Receivables information
type AssetsInfoCurrentAssetsReceivables struct {
	// Total receivables
	TotalReceivables *float64 `json:"total_receivables,omitempty"`
	// Accounts receivable
	AccountsReceivable *float64 `json:"accounts_receivable,omitempty"`
	// Gross accounts receivable
	GrossAccountsReceivable *float64 `json:"gross_accounts_receivable,omitempty"`
	// Allowance for doubtful accounts receivable
	AllowanceForDoubtfulAccountsReceivable *float64 `json:"allowance_for_doubtful_accounts_receivable,omitempty"`
	// Receivables adjustments allowances
	ReceivablesAdjustmentsAllowances *float64 `json:"receivables_adjustments_allowances,omitempty"`
	// Other receivables
	OtherReceivables *float64 `json:"other_receivables,omitempty"`
	// Due from related parties current
	DueFromRelatedPartiesCurrent *float64 `json:"due_from_related_parties_current,omitempty"`
	// Taxes receivable
	TaxesReceivable *float64 `json:"taxes_receivable,omitempty"`
	// Accrued interest receivable
	AccruedInterestReceivable *float64 `json:"accrued_interest_receivable,omitempty"`
	// Notes receivable
	NotesReceivable *float64 `json:"notes_receivable,omitempty"`
	// Loans receivable
	LoansReceivable *float64 `json:"loans_receivable,omitempty"`
}

// NewAssetsInfoCurrentAssetsReceivables instantiates a new AssetsInfoCurrentAssetsReceivables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoCurrentAssetsReceivables() *AssetsInfoCurrentAssetsReceivables {
	this := AssetsInfoCurrentAssetsReceivables{}
	return &this
}

// NewAssetsInfoCurrentAssetsReceivablesWithDefaults instantiates a new AssetsInfoCurrentAssetsReceivables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoCurrentAssetsReceivablesWithDefaults() *AssetsInfoCurrentAssetsReceivables {
	this := AssetsInfoCurrentAssetsReceivables{}
	return &this
}

// GetTotalReceivables returns the TotalReceivables field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetTotalReceivables() float64 {
	if o == nil || IsNil(o.TotalReceivables) {
		var ret float64
		return ret
	}
	return *o.TotalReceivables
}

// GetTotalReceivablesOk returns a tuple with the TotalReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetTotalReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalReceivables) {
		return nil, false
	}
	return o.TotalReceivables, true
}

// HasTotalReceivables returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasTotalReceivables() bool {
	if o != nil && !IsNil(o.TotalReceivables) {
		return true
	}

	return false
}

// SetTotalReceivables gets a reference to the given float64 and assigns it to the TotalReceivables field.
func (o *AssetsInfoCurrentAssetsReceivables) SetTotalReceivables(v float64) {
	o.TotalReceivables = &v
}

// GetAccountsReceivable returns the AccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccountsReceivable() float64 {
	if o == nil || IsNil(o.AccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.AccountsReceivable
}

// GetAccountsReceivableOk returns a tuple with the AccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccountsReceivable) {
		return nil, false
	}
	return o.AccountsReceivable, true
}

// HasAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasAccountsReceivable() bool {
	if o != nil && !IsNil(o.AccountsReceivable) {
		return true
	}

	return false
}

// SetAccountsReceivable gets a reference to the given float64 and assigns it to the AccountsReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetAccountsReceivable(v float64) {
	o.AccountsReceivable = &v
}

// GetGrossAccountsReceivable returns the GrossAccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetGrossAccountsReceivable() float64 {
	if o == nil || IsNil(o.GrossAccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.GrossAccountsReceivable
}

// GetGrossAccountsReceivableOk returns a tuple with the GrossAccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetGrossAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossAccountsReceivable) {
		return nil, false
	}
	return o.GrossAccountsReceivable, true
}

// HasGrossAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasGrossAccountsReceivable() bool {
	if o != nil && !IsNil(o.GrossAccountsReceivable) {
		return true
	}

	return false
}

// SetGrossAccountsReceivable gets a reference to the given float64 and assigns it to the GrossAccountsReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetGrossAccountsReceivable(v float64) {
	o.GrossAccountsReceivable = &v
}

// GetAllowanceForDoubtfulAccountsReceivable returns the AllowanceForDoubtfulAccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetAllowanceForDoubtfulAccountsReceivable() float64 {
	if o == nil || IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.AllowanceForDoubtfulAccountsReceivable
}

// GetAllowanceForDoubtfulAccountsReceivableOk returns a tuple with the AllowanceForDoubtfulAccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetAllowanceForDoubtfulAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		return nil, false
	}
	return o.AllowanceForDoubtfulAccountsReceivable, true
}

// HasAllowanceForDoubtfulAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasAllowanceForDoubtfulAccountsReceivable() bool {
	if o != nil && !IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		return true
	}

	return false
}

// SetAllowanceForDoubtfulAccountsReceivable gets a reference to the given float64 and assigns it to the AllowanceForDoubtfulAccountsReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetAllowanceForDoubtfulAccountsReceivable(v float64) {
	o.AllowanceForDoubtfulAccountsReceivable = &v
}

// GetReceivablesAdjustmentsAllowances returns the ReceivablesAdjustmentsAllowances field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetReceivablesAdjustmentsAllowances() float64 {
	if o == nil || IsNil(o.ReceivablesAdjustmentsAllowances) {
		var ret float64
		return ret
	}
	return *o.ReceivablesAdjustmentsAllowances
}

// GetReceivablesAdjustmentsAllowancesOk returns a tuple with the ReceivablesAdjustmentsAllowances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetReceivablesAdjustmentsAllowancesOk() (*float64, bool) {
	if o == nil || IsNil(o.ReceivablesAdjustmentsAllowances) {
		return nil, false
	}
	return o.ReceivablesAdjustmentsAllowances, true
}

// HasReceivablesAdjustmentsAllowances returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasReceivablesAdjustmentsAllowances() bool {
	if o != nil && !IsNil(o.ReceivablesAdjustmentsAllowances) {
		return true
	}

	return false
}

// SetReceivablesAdjustmentsAllowances gets a reference to the given float64 and assigns it to the ReceivablesAdjustmentsAllowances field.
func (o *AssetsInfoCurrentAssetsReceivables) SetReceivablesAdjustmentsAllowances(v float64) {
	o.ReceivablesAdjustmentsAllowances = &v
}

// GetOtherReceivables returns the OtherReceivables field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetOtherReceivables() float64 {
	if o == nil || IsNil(o.OtherReceivables) {
		var ret float64
		return ret
	}
	return *o.OtherReceivables
}

// GetOtherReceivablesOk returns a tuple with the OtherReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetOtherReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherReceivables) {
		return nil, false
	}
	return o.OtherReceivables, true
}

// HasOtherReceivables returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasOtherReceivables() bool {
	if o != nil && !IsNil(o.OtherReceivables) {
		return true
	}

	return false
}

// SetOtherReceivables gets a reference to the given float64 and assigns it to the OtherReceivables field.
func (o *AssetsInfoCurrentAssetsReceivables) SetOtherReceivables(v float64) {
	o.OtherReceivables = &v
}

// GetDueFromRelatedPartiesCurrent returns the DueFromRelatedPartiesCurrent field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetDueFromRelatedPartiesCurrent() float64 {
	if o == nil || IsNil(o.DueFromRelatedPartiesCurrent) {
		var ret float64
		return ret
	}
	return *o.DueFromRelatedPartiesCurrent
}

// GetDueFromRelatedPartiesCurrentOk returns a tuple with the DueFromRelatedPartiesCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetDueFromRelatedPartiesCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.DueFromRelatedPartiesCurrent) {
		return nil, false
	}
	return o.DueFromRelatedPartiesCurrent, true
}

// HasDueFromRelatedPartiesCurrent returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasDueFromRelatedPartiesCurrent() bool {
	if o != nil && !IsNil(o.DueFromRelatedPartiesCurrent) {
		return true
	}

	return false
}

// SetDueFromRelatedPartiesCurrent gets a reference to the given float64 and assigns it to the DueFromRelatedPartiesCurrent field.
func (o *AssetsInfoCurrentAssetsReceivables) SetDueFromRelatedPartiesCurrent(v float64) {
	o.DueFromRelatedPartiesCurrent = &v
}

// GetTaxesReceivable returns the TaxesReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetTaxesReceivable() float64 {
	if o == nil || IsNil(o.TaxesReceivable) {
		var ret float64
		return ret
	}
	return *o.TaxesReceivable
}

// GetTaxesReceivableOk returns a tuple with the TaxesReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetTaxesReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxesReceivable) {
		return nil, false
	}
	return o.TaxesReceivable, true
}

// HasTaxesReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasTaxesReceivable() bool {
	if o != nil && !IsNil(o.TaxesReceivable) {
		return true
	}

	return false
}

// SetTaxesReceivable gets a reference to the given float64 and assigns it to the TaxesReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetTaxesReceivable(v float64) {
	o.TaxesReceivable = &v
}

// GetAccruedInterestReceivable returns the AccruedInterestReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccruedInterestReceivable() float64 {
	if o == nil || IsNil(o.AccruedInterestReceivable) {
		var ret float64
		return ret
	}
	return *o.AccruedInterestReceivable
}

// GetAccruedInterestReceivableOk returns a tuple with the AccruedInterestReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetAccruedInterestReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccruedInterestReceivable) {
		return nil, false
	}
	return o.AccruedInterestReceivable, true
}

// HasAccruedInterestReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasAccruedInterestReceivable() bool {
	if o != nil && !IsNil(o.AccruedInterestReceivable) {
		return true
	}

	return false
}

// SetAccruedInterestReceivable gets a reference to the given float64 and assigns it to the AccruedInterestReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetAccruedInterestReceivable(v float64) {
	o.AccruedInterestReceivable = &v
}

// GetNotesReceivable returns the NotesReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetNotesReceivable() float64 {
	if o == nil || IsNil(o.NotesReceivable) {
		var ret float64
		return ret
	}
	return *o.NotesReceivable
}

// GetNotesReceivableOk returns a tuple with the NotesReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetNotesReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.NotesReceivable) {
		return nil, false
	}
	return o.NotesReceivable, true
}

// HasNotesReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasNotesReceivable() bool {
	if o != nil && !IsNil(o.NotesReceivable) {
		return true
	}

	return false
}

// SetNotesReceivable gets a reference to the given float64 and assigns it to the NotesReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetNotesReceivable(v float64) {
	o.NotesReceivable = &v
}

// GetLoansReceivable returns the LoansReceivable field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssetsReceivables) GetLoansReceivable() float64 {
	if o == nil || IsNil(o.LoansReceivable) {
		var ret float64
		return ret
	}
	return *o.LoansReceivable
}

// GetLoansReceivableOk returns a tuple with the LoansReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssetsReceivables) GetLoansReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.LoansReceivable) {
		return nil, false
	}
	return o.LoansReceivable, true
}

// HasLoansReceivable returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssetsReceivables) HasLoansReceivable() bool {
	if o != nil && !IsNil(o.LoansReceivable) {
		return true
	}

	return false
}

// SetLoansReceivable gets a reference to the given float64 and assigns it to the LoansReceivable field.
func (o *AssetsInfoCurrentAssetsReceivables) SetLoansReceivable(v float64) {
	o.LoansReceivable = &v
}

func (o AssetsInfoCurrentAssetsReceivables) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoCurrentAssetsReceivables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalReceivables) {
		toSerialize["total_receivables"] = o.TotalReceivables
	}
	if !IsNil(o.AccountsReceivable) {
		toSerialize["accounts_receivable"] = o.AccountsReceivable
	}
	if !IsNil(o.GrossAccountsReceivable) {
		toSerialize["gross_accounts_receivable"] = o.GrossAccountsReceivable
	}
	if !IsNil(o.AllowanceForDoubtfulAccountsReceivable) {
		toSerialize["allowance_for_doubtful_accounts_receivable"] = o.AllowanceForDoubtfulAccountsReceivable
	}
	if !IsNil(o.ReceivablesAdjustmentsAllowances) {
		toSerialize["receivables_adjustments_allowances"] = o.ReceivablesAdjustmentsAllowances
	}
	if !IsNil(o.OtherReceivables) {
		toSerialize["other_receivables"] = o.OtherReceivables
	}
	if !IsNil(o.DueFromRelatedPartiesCurrent) {
		toSerialize["due_from_related_parties_current"] = o.DueFromRelatedPartiesCurrent
	}
	if !IsNil(o.TaxesReceivable) {
		toSerialize["taxes_receivable"] = o.TaxesReceivable
	}
	if !IsNil(o.AccruedInterestReceivable) {
		toSerialize["accrued_interest_receivable"] = o.AccruedInterestReceivable
	}
	if !IsNil(o.NotesReceivable) {
		toSerialize["notes_receivable"] = o.NotesReceivable
	}
	if !IsNil(o.LoansReceivable) {
		toSerialize["loans_receivable"] = o.LoansReceivable
	}
	return toSerialize, nil
}

type NullableAssetsInfoCurrentAssetsReceivables struct {
	value *AssetsInfoCurrentAssetsReceivables
	isSet bool
}

func (v NullableAssetsInfoCurrentAssetsReceivables) Get() *AssetsInfoCurrentAssetsReceivables {
	return v.value
}

func (v *NullableAssetsInfoCurrentAssetsReceivables) Set(val *AssetsInfoCurrentAssetsReceivables) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoCurrentAssetsReceivables) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoCurrentAssetsReceivables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoCurrentAssetsReceivables(val *AssetsInfoCurrentAssetsReceivables) *NullableAssetsInfoCurrentAssetsReceivables {
	return &NullableAssetsInfoCurrentAssetsReceivables{value: val, isSet: true}
}

func (v NullableAssetsInfoCurrentAssetsReceivables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoCurrentAssetsReceivables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
