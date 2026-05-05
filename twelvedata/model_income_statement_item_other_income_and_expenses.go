// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemOtherIncomeAndExpenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemOtherIncomeAndExpenses{}

// IncomeStatementItemOtherIncomeAndExpenses Other income and expenses information
type IncomeStatementItemOtherIncomeAndExpenses struct {
	// Other income expense
	OtherIncomeExpense *float64 `json:"other_income_expense,omitempty"`
	// Other non operating income expenses
	OtherNonOperatingIncomeExpenses *float64 `json:"other_non_operating_income_expenses,omitempty"`
	// Special income charges
	SpecialIncomeCharges *float64 `json:"special_income_charges,omitempty"`
	// Gain on sale of PPE
	GainOnSaleOfPpe *float64 `json:"gain_on_sale_of_ppe,omitempty"`
	// Gain on sale of business
	GainOnSaleOfBusiness *float64 `json:"gain_on_sale_of_business,omitempty"`
	// Gain on sale of security
	GainOnSaleOfSecurity *float64 `json:"gain_on_sale_of_security,omitempty"`
	// Other special charges
	OtherSpecialCharges *float64 `json:"other_special_charges,omitempty"`
	// Write off
	WriteOff *float64 `json:"write_off,omitempty"`
	// Impairment of capital assets
	ImpairmentOfCapitalAssets *float64 `json:"impairment_of_capital_assets,omitempty"`
	// Restructuring and merger acquisition
	RestructuringAndMergerAcquisition *float64 `json:"restructuring_and_merger_acquisition,omitempty"`
	// Securities amortization
	SecuritiesAmortization *float64 `json:"securities_amortization,omitempty"`
	// Earnings from equity interest
	EarningsFromEquityInterest *float64 `json:"earnings_from_equity_interest,omitempty"`
	// Earnings from equity interest net of tax
	EarningsFromEquityInterestNetOfTax *float64 `json:"earnings_from_equity_interest_net_of_tax,omitempty"`
	// Total other finance cost
	TotalOtherFinanceCost *float64 `json:"total_other_finance_cost,omitempty"`
}

// NewIncomeStatementItemOtherIncomeAndExpenses instantiates a new IncomeStatementItemOtherIncomeAndExpenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemOtherIncomeAndExpenses() *IncomeStatementItemOtherIncomeAndExpenses {
	this := IncomeStatementItemOtherIncomeAndExpenses{}
	return &this
}

// NewIncomeStatementItemOtherIncomeAndExpensesWithDefaults instantiates a new IncomeStatementItemOtherIncomeAndExpenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemOtherIncomeAndExpensesWithDefaults() *IncomeStatementItemOtherIncomeAndExpenses {
	this := IncomeStatementItemOtherIncomeAndExpenses{}
	return &this
}

// GetOtherIncomeExpense returns the OtherIncomeExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherIncomeExpense() float64 {
	if o == nil || IsNil(o.OtherIncomeExpense) {
		var ret float64
		return ret
	}
	return *o.OtherIncomeExpense
}

// GetOtherIncomeExpenseOk returns a tuple with the OtherIncomeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherIncomeExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherIncomeExpense) {
		return nil, false
	}
	return o.OtherIncomeExpense, true
}

// HasOtherIncomeExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasOtherIncomeExpense() bool {
	if o != nil && !IsNil(o.OtherIncomeExpense) {
		return true
	}

	return false
}

// SetOtherIncomeExpense gets a reference to the given float64 and assigns it to the OtherIncomeExpense field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetOtherIncomeExpense(v float64) {
	o.OtherIncomeExpense = &v
}

// GetOtherNonOperatingIncomeExpenses returns the OtherNonOperatingIncomeExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherNonOperatingIncomeExpenses() float64 {
	if o == nil || IsNil(o.OtherNonOperatingIncomeExpenses) {
		var ret float64
		return ret
	}
	return *o.OtherNonOperatingIncomeExpenses
}

// GetOtherNonOperatingIncomeExpensesOk returns a tuple with the OtherNonOperatingIncomeExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherNonOperatingIncomeExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonOperatingIncomeExpenses) {
		return nil, false
	}
	return o.OtherNonOperatingIncomeExpenses, true
}

// HasOtherNonOperatingIncomeExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasOtherNonOperatingIncomeExpenses() bool {
	if o != nil && !IsNil(o.OtherNonOperatingIncomeExpenses) {
		return true
	}

	return false
}

// SetOtherNonOperatingIncomeExpenses gets a reference to the given float64 and assigns it to the OtherNonOperatingIncomeExpenses field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetOtherNonOperatingIncomeExpenses(v float64) {
	o.OtherNonOperatingIncomeExpenses = &v
}

// GetSpecialIncomeCharges returns the SpecialIncomeCharges field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSpecialIncomeCharges() float64 {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		var ret float64
		return ret
	}
	return *o.SpecialIncomeCharges
}

// GetSpecialIncomeChargesOk returns a tuple with the SpecialIncomeCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSpecialIncomeChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		return nil, false
	}
	return o.SpecialIncomeCharges, true
}

// HasSpecialIncomeCharges returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasSpecialIncomeCharges() bool {
	if o != nil && !IsNil(o.SpecialIncomeCharges) {
		return true
	}

	return false
}

// SetSpecialIncomeCharges gets a reference to the given float64 and assigns it to the SpecialIncomeCharges field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetSpecialIncomeCharges(v float64) {
	o.SpecialIncomeCharges = &v
}

// GetGainOnSaleOfPpe returns the GainOnSaleOfPpe field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfPpe() float64 {
	if o == nil || IsNil(o.GainOnSaleOfPpe) {
		var ret float64
		return ret
	}
	return *o.GainOnSaleOfPpe
}

// GetGainOnSaleOfPpeOk returns a tuple with the GainOnSaleOfPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.GainOnSaleOfPpe) {
		return nil, false
	}
	return o.GainOnSaleOfPpe, true
}

// HasGainOnSaleOfPpe returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasGainOnSaleOfPpe() bool {
	if o != nil && !IsNil(o.GainOnSaleOfPpe) {
		return true
	}

	return false
}

// SetGainOnSaleOfPpe gets a reference to the given float64 and assigns it to the GainOnSaleOfPpe field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetGainOnSaleOfPpe(v float64) {
	o.GainOnSaleOfPpe = &v
}

// GetGainOnSaleOfBusiness returns the GainOnSaleOfBusiness field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfBusiness() float64 {
	if o == nil || IsNil(o.GainOnSaleOfBusiness) {
		var ret float64
		return ret
	}
	return *o.GainOnSaleOfBusiness
}

// GetGainOnSaleOfBusinessOk returns a tuple with the GainOnSaleOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfBusinessOk() (*float64, bool) {
	if o == nil || IsNil(o.GainOnSaleOfBusiness) {
		return nil, false
	}
	return o.GainOnSaleOfBusiness, true
}

// HasGainOnSaleOfBusiness returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasGainOnSaleOfBusiness() bool {
	if o != nil && !IsNil(o.GainOnSaleOfBusiness) {
		return true
	}

	return false
}

// SetGainOnSaleOfBusiness gets a reference to the given float64 and assigns it to the GainOnSaleOfBusiness field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetGainOnSaleOfBusiness(v float64) {
	o.GainOnSaleOfBusiness = &v
}

// GetGainOnSaleOfSecurity returns the GainOnSaleOfSecurity field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfSecurity() float64 {
	if o == nil || IsNil(o.GainOnSaleOfSecurity) {
		var ret float64
		return ret
	}
	return *o.GainOnSaleOfSecurity
}

// GetGainOnSaleOfSecurityOk returns a tuple with the GainOnSaleOfSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetGainOnSaleOfSecurityOk() (*float64, bool) {
	if o == nil || IsNil(o.GainOnSaleOfSecurity) {
		return nil, false
	}
	return o.GainOnSaleOfSecurity, true
}

// HasGainOnSaleOfSecurity returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasGainOnSaleOfSecurity() bool {
	if o != nil && !IsNil(o.GainOnSaleOfSecurity) {
		return true
	}

	return false
}

// SetGainOnSaleOfSecurity gets a reference to the given float64 and assigns it to the GainOnSaleOfSecurity field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetGainOnSaleOfSecurity(v float64) {
	o.GainOnSaleOfSecurity = &v
}

// GetOtherSpecialCharges returns the OtherSpecialCharges field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherSpecialCharges() float64 {
	if o == nil || IsNil(o.OtherSpecialCharges) {
		var ret float64
		return ret
	}
	return *o.OtherSpecialCharges
}

// GetOtherSpecialChargesOk returns a tuple with the OtherSpecialCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetOtherSpecialChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherSpecialCharges) {
		return nil, false
	}
	return o.OtherSpecialCharges, true
}

// HasOtherSpecialCharges returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasOtherSpecialCharges() bool {
	if o != nil && !IsNil(o.OtherSpecialCharges) {
		return true
	}

	return false
}

// SetOtherSpecialCharges gets a reference to the given float64 and assigns it to the OtherSpecialCharges field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetOtherSpecialCharges(v float64) {
	o.OtherSpecialCharges = &v
}

// GetWriteOff returns the WriteOff field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetWriteOff() float64 {
	if o == nil || IsNil(o.WriteOff) {
		var ret float64
		return ret
	}
	return *o.WriteOff
}

// GetWriteOffOk returns a tuple with the WriteOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetWriteOffOk() (*float64, bool) {
	if o == nil || IsNil(o.WriteOff) {
		return nil, false
	}
	return o.WriteOff, true
}

// HasWriteOff returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasWriteOff() bool {
	if o != nil && !IsNil(o.WriteOff) {
		return true
	}

	return false
}

// SetWriteOff gets a reference to the given float64 and assigns it to the WriteOff field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetWriteOff(v float64) {
	o.WriteOff = &v
}

// GetImpairmentOfCapitalAssets returns the ImpairmentOfCapitalAssets field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetImpairmentOfCapitalAssets() float64 {
	if o == nil || IsNil(o.ImpairmentOfCapitalAssets) {
		var ret float64
		return ret
	}
	return *o.ImpairmentOfCapitalAssets
}

// GetImpairmentOfCapitalAssetsOk returns a tuple with the ImpairmentOfCapitalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetImpairmentOfCapitalAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.ImpairmentOfCapitalAssets) {
		return nil, false
	}
	return o.ImpairmentOfCapitalAssets, true
}

// HasImpairmentOfCapitalAssets returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasImpairmentOfCapitalAssets() bool {
	if o != nil && !IsNil(o.ImpairmentOfCapitalAssets) {
		return true
	}

	return false
}

// SetImpairmentOfCapitalAssets gets a reference to the given float64 and assigns it to the ImpairmentOfCapitalAssets field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetImpairmentOfCapitalAssets(v float64) {
	o.ImpairmentOfCapitalAssets = &v
}

// GetRestructuringAndMergerAcquisition returns the RestructuringAndMergerAcquisition field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetRestructuringAndMergerAcquisition() float64 {
	if o == nil || IsNil(o.RestructuringAndMergerAcquisition) {
		var ret float64
		return ret
	}
	return *o.RestructuringAndMergerAcquisition
}

// GetRestructuringAndMergerAcquisitionOk returns a tuple with the RestructuringAndMergerAcquisition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetRestructuringAndMergerAcquisitionOk() (*float64, bool) {
	if o == nil || IsNil(o.RestructuringAndMergerAcquisition) {
		return nil, false
	}
	return o.RestructuringAndMergerAcquisition, true
}

// HasRestructuringAndMergerAcquisition returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasRestructuringAndMergerAcquisition() bool {
	if o != nil && !IsNil(o.RestructuringAndMergerAcquisition) {
		return true
	}

	return false
}

// SetRestructuringAndMergerAcquisition gets a reference to the given float64 and assigns it to the RestructuringAndMergerAcquisition field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetRestructuringAndMergerAcquisition(v float64) {
	o.RestructuringAndMergerAcquisition = &v
}

// GetSecuritiesAmortization returns the SecuritiesAmortization field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSecuritiesAmortization() float64 {
	if o == nil || IsNil(o.SecuritiesAmortization) {
		var ret float64
		return ret
	}
	return *o.SecuritiesAmortization
}

// GetSecuritiesAmortizationOk returns a tuple with the SecuritiesAmortization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetSecuritiesAmortizationOk() (*float64, bool) {
	if o == nil || IsNil(o.SecuritiesAmortization) {
		return nil, false
	}
	return o.SecuritiesAmortization, true
}

// HasSecuritiesAmortization returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasSecuritiesAmortization() bool {
	if o != nil && !IsNil(o.SecuritiesAmortization) {
		return true
	}

	return false
}

// SetSecuritiesAmortization gets a reference to the given float64 and assigns it to the SecuritiesAmortization field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetSecuritiesAmortization(v float64) {
	o.SecuritiesAmortization = &v
}

// GetEarningsFromEquityInterest returns the EarningsFromEquityInterest field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterest() float64 {
	if o == nil || IsNil(o.EarningsFromEquityInterest) {
		var ret float64
		return ret
	}
	return *o.EarningsFromEquityInterest
}

// GetEarningsFromEquityInterestOk returns a tuple with the EarningsFromEquityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.EarningsFromEquityInterest) {
		return nil, false
	}
	return o.EarningsFromEquityInterest, true
}

// HasEarningsFromEquityInterest returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasEarningsFromEquityInterest() bool {
	if o != nil && !IsNil(o.EarningsFromEquityInterest) {
		return true
	}

	return false
}

// SetEarningsFromEquityInterest gets a reference to the given float64 and assigns it to the EarningsFromEquityInterest field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetEarningsFromEquityInterest(v float64) {
	o.EarningsFromEquityInterest = &v
}

// GetEarningsFromEquityInterestNetOfTax returns the EarningsFromEquityInterestNetOfTax field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterestNetOfTax() float64 {
	if o == nil || IsNil(o.EarningsFromEquityInterestNetOfTax) {
		var ret float64
		return ret
	}
	return *o.EarningsFromEquityInterestNetOfTax
}

// GetEarningsFromEquityInterestNetOfTaxOk returns a tuple with the EarningsFromEquityInterestNetOfTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetEarningsFromEquityInterestNetOfTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.EarningsFromEquityInterestNetOfTax) {
		return nil, false
	}
	return o.EarningsFromEquityInterestNetOfTax, true
}

// HasEarningsFromEquityInterestNetOfTax returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasEarningsFromEquityInterestNetOfTax() bool {
	if o != nil && !IsNil(o.EarningsFromEquityInterestNetOfTax) {
		return true
	}

	return false
}

// SetEarningsFromEquityInterestNetOfTax gets a reference to the given float64 and assigns it to the EarningsFromEquityInterestNetOfTax field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetEarningsFromEquityInterestNetOfTax(v float64) {
	o.EarningsFromEquityInterestNetOfTax = &v
}

// GetTotalOtherFinanceCost returns the TotalOtherFinanceCost field value if set, zero value otherwise.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetTotalOtherFinanceCost() float64 {
	if o == nil || IsNil(o.TotalOtherFinanceCost) {
		var ret float64
		return ret
	}
	return *o.TotalOtherFinanceCost
}

// GetTotalOtherFinanceCostOk returns a tuple with the TotalOtherFinanceCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) GetTotalOtherFinanceCostOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalOtherFinanceCost) {
		return nil, false
	}
	return o.TotalOtherFinanceCost, true
}

// HasTotalOtherFinanceCost returns a boolean if a field has been set.
func (o *IncomeStatementItemOtherIncomeAndExpenses) HasTotalOtherFinanceCost() bool {
	if o != nil && !IsNil(o.TotalOtherFinanceCost) {
		return true
	}

	return false
}

// SetTotalOtherFinanceCost gets a reference to the given float64 and assigns it to the TotalOtherFinanceCost field.
func (o *IncomeStatementItemOtherIncomeAndExpenses) SetTotalOtherFinanceCost(v float64) {
	o.TotalOtherFinanceCost = &v
}

func (o IncomeStatementItemOtherIncomeAndExpenses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemOtherIncomeAndExpenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OtherIncomeExpense) {
		toSerialize["other_income_expense"] = o.OtherIncomeExpense
	}
	if !IsNil(o.OtherNonOperatingIncomeExpenses) {
		toSerialize["other_non_operating_income_expenses"] = o.OtherNonOperatingIncomeExpenses
	}
	if !IsNil(o.SpecialIncomeCharges) {
		toSerialize["special_income_charges"] = o.SpecialIncomeCharges
	}
	if !IsNil(o.GainOnSaleOfPpe) {
		toSerialize["gain_on_sale_of_ppe"] = o.GainOnSaleOfPpe
	}
	if !IsNil(o.GainOnSaleOfBusiness) {
		toSerialize["gain_on_sale_of_business"] = o.GainOnSaleOfBusiness
	}
	if !IsNil(o.GainOnSaleOfSecurity) {
		toSerialize["gain_on_sale_of_security"] = o.GainOnSaleOfSecurity
	}
	if !IsNil(o.OtherSpecialCharges) {
		toSerialize["other_special_charges"] = o.OtherSpecialCharges
	}
	if !IsNil(o.WriteOff) {
		toSerialize["write_off"] = o.WriteOff
	}
	if !IsNil(o.ImpairmentOfCapitalAssets) {
		toSerialize["impairment_of_capital_assets"] = o.ImpairmentOfCapitalAssets
	}
	if !IsNil(o.RestructuringAndMergerAcquisition) {
		toSerialize["restructuring_and_merger_acquisition"] = o.RestructuringAndMergerAcquisition
	}
	if !IsNil(o.SecuritiesAmortization) {
		toSerialize["securities_amortization"] = o.SecuritiesAmortization
	}
	if !IsNil(o.EarningsFromEquityInterest) {
		toSerialize["earnings_from_equity_interest"] = o.EarningsFromEquityInterest
	}
	if !IsNil(o.EarningsFromEquityInterestNetOfTax) {
		toSerialize["earnings_from_equity_interest_net_of_tax"] = o.EarningsFromEquityInterestNetOfTax
	}
	if !IsNil(o.TotalOtherFinanceCost) {
		toSerialize["total_other_finance_cost"] = o.TotalOtherFinanceCost
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemOtherIncomeAndExpenses struct {
	value *IncomeStatementItemOtherIncomeAndExpenses
	isSet bool
}

func (v NullableIncomeStatementItemOtherIncomeAndExpenses) Get() *IncomeStatementItemOtherIncomeAndExpenses {
	return v.value
}

func (v *NullableIncomeStatementItemOtherIncomeAndExpenses) Set(val *IncomeStatementItemOtherIncomeAndExpenses) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemOtherIncomeAndExpenses) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemOtherIncomeAndExpenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemOtherIncomeAndExpenses(val *IncomeStatementItemOtherIncomeAndExpenses) *NullableIncomeStatementItemOtherIncomeAndExpenses {
	return &NullableIncomeStatementItemOtherIncomeAndExpenses{value: val, isSet: true}
}

func (v NullableIncomeStatementItemOtherIncomeAndExpenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemOtherIncomeAndExpenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
