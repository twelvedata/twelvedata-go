// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemExpenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemExpenses{}

// IncomeStatementItemExpenses Expenses information
type IncomeStatementItemExpenses struct {
	// Total expenses
	TotalExpenses *float64 `json:"total_expenses,omitempty"`
	// Selling general and administration expense
	SellingGeneralAndAdministrationExpense *float64 `json:"selling_general_and_administration_expense,omitempty"`
	// Selling and marketing expense
	SellingAndMarketingExpense *float64 `json:"selling_and_marketing_expense,omitempty"`
	// General and administrative expense
	GeneralAndAdministrativeExpense *float64 `json:"general_and_administrative_expense,omitempty"`
	// Other general and administrative expense
	OtherGeneralAndAdministrativeExpense *float64 `json:"other_general_and_administrative_expense,omitempty"`
	// Depreciation amortization depletion income statement
	DepreciationAmortizationDepletionIncomeStatement *float64 `json:"depreciation_amortization_depletion_income_statement,omitempty"`
	// Research and development expense
	ResearchAndDevelopmentExpense *float64 `json:"research_and_development_expense,omitempty"`
	// Insurance and claims expense
	InsuranceAndClaimsExpense *float64 `json:"insurance_and_claims_expense,omitempty"`
	// Rent and landing fees
	RentAndLandingFees *float64 `json:"rent_and_landing_fees,omitempty"`
	// Salaries and wages expense
	SalariesAndWagesExpense *float64 `json:"salaries_and_wages_expense,omitempty"`
	// Rent expense supplemental
	RentExpenseSupplemental *float64 `json:"rent_expense_supplemental,omitempty"`
	// Provision for doubtful accounts
	ProvisionForDoubtfulAccounts *float64 `json:"provision_for_doubtful_accounts,omitempty"`
}

// NewIncomeStatementItemExpenses instantiates a new IncomeStatementItemExpenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemExpenses() *IncomeStatementItemExpenses {
	this := IncomeStatementItemExpenses{}
	return &this
}

// NewIncomeStatementItemExpensesWithDefaults instantiates a new IncomeStatementItemExpenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemExpensesWithDefaults() *IncomeStatementItemExpenses {
	this := IncomeStatementItemExpenses{}
	return &this
}

// GetTotalExpenses returns the TotalExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetTotalExpenses() float64 {
	if o == nil || IsNil(o.TotalExpenses) {
		var ret float64
		return ret
	}
	return *o.TotalExpenses
}

// GetTotalExpensesOk returns a tuple with the TotalExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetTotalExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalExpenses) {
		return nil, false
	}
	return o.TotalExpenses, true
}

// HasTotalExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasTotalExpenses() bool {
	if o != nil && !IsNil(o.TotalExpenses) {
		return true
	}

	return false
}

// SetTotalExpenses gets a reference to the given float64 and assigns it to the TotalExpenses field.
func (o *IncomeStatementItemExpenses) SetTotalExpenses(v float64) {
	o.TotalExpenses = &v
}

// GetSellingGeneralAndAdministrationExpense returns the SellingGeneralAndAdministrationExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetSellingGeneralAndAdministrationExpense() float64 {
	if o == nil || IsNil(o.SellingGeneralAndAdministrationExpense) {
		var ret float64
		return ret
	}
	return *o.SellingGeneralAndAdministrationExpense
}

// GetSellingGeneralAndAdministrationExpenseOk returns a tuple with the SellingGeneralAndAdministrationExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetSellingGeneralAndAdministrationExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.SellingGeneralAndAdministrationExpense) {
		return nil, false
	}
	return o.SellingGeneralAndAdministrationExpense, true
}

// HasSellingGeneralAndAdministrationExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasSellingGeneralAndAdministrationExpense() bool {
	if o != nil && !IsNil(o.SellingGeneralAndAdministrationExpense) {
		return true
	}

	return false
}

// SetSellingGeneralAndAdministrationExpense gets a reference to the given float64 and assigns it to the SellingGeneralAndAdministrationExpense field.
func (o *IncomeStatementItemExpenses) SetSellingGeneralAndAdministrationExpense(v float64) {
	o.SellingGeneralAndAdministrationExpense = &v
}

// GetSellingAndMarketingExpense returns the SellingAndMarketingExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetSellingAndMarketingExpense() float64 {
	if o == nil || IsNil(o.SellingAndMarketingExpense) {
		var ret float64
		return ret
	}
	return *o.SellingAndMarketingExpense
}

// GetSellingAndMarketingExpenseOk returns a tuple with the SellingAndMarketingExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetSellingAndMarketingExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.SellingAndMarketingExpense) {
		return nil, false
	}
	return o.SellingAndMarketingExpense, true
}

// HasSellingAndMarketingExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasSellingAndMarketingExpense() bool {
	if o != nil && !IsNil(o.SellingAndMarketingExpense) {
		return true
	}

	return false
}

// SetSellingAndMarketingExpense gets a reference to the given float64 and assigns it to the SellingAndMarketingExpense field.
func (o *IncomeStatementItemExpenses) SetSellingAndMarketingExpense(v float64) {
	o.SellingAndMarketingExpense = &v
}

// GetGeneralAndAdministrativeExpense returns the GeneralAndAdministrativeExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetGeneralAndAdministrativeExpense() float64 {
	if o == nil || IsNil(o.GeneralAndAdministrativeExpense) {
		var ret float64
		return ret
	}
	return *o.GeneralAndAdministrativeExpense
}

// GetGeneralAndAdministrativeExpenseOk returns a tuple with the GeneralAndAdministrativeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetGeneralAndAdministrativeExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.GeneralAndAdministrativeExpense) {
		return nil, false
	}
	return o.GeneralAndAdministrativeExpense, true
}

// HasGeneralAndAdministrativeExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasGeneralAndAdministrativeExpense() bool {
	if o != nil && !IsNil(o.GeneralAndAdministrativeExpense) {
		return true
	}

	return false
}

// SetGeneralAndAdministrativeExpense gets a reference to the given float64 and assigns it to the GeneralAndAdministrativeExpense field.
func (o *IncomeStatementItemExpenses) SetGeneralAndAdministrativeExpense(v float64) {
	o.GeneralAndAdministrativeExpense = &v
}

// GetOtherGeneralAndAdministrativeExpense returns the OtherGeneralAndAdministrativeExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetOtherGeneralAndAdministrativeExpense() float64 {
	if o == nil || IsNil(o.OtherGeneralAndAdministrativeExpense) {
		var ret float64
		return ret
	}
	return *o.OtherGeneralAndAdministrativeExpense
}

// GetOtherGeneralAndAdministrativeExpenseOk returns a tuple with the OtherGeneralAndAdministrativeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetOtherGeneralAndAdministrativeExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherGeneralAndAdministrativeExpense) {
		return nil, false
	}
	return o.OtherGeneralAndAdministrativeExpense, true
}

// HasOtherGeneralAndAdministrativeExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasOtherGeneralAndAdministrativeExpense() bool {
	if o != nil && !IsNil(o.OtherGeneralAndAdministrativeExpense) {
		return true
	}

	return false
}

// SetOtherGeneralAndAdministrativeExpense gets a reference to the given float64 and assigns it to the OtherGeneralAndAdministrativeExpense field.
func (o *IncomeStatementItemExpenses) SetOtherGeneralAndAdministrativeExpense(v float64) {
	o.OtherGeneralAndAdministrativeExpense = &v
}

// GetDepreciationAmortizationDepletionIncomeStatement returns the DepreciationAmortizationDepletionIncomeStatement field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetDepreciationAmortizationDepletionIncomeStatement() float64 {
	if o == nil || IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		var ret float64
		return ret
	}
	return *o.DepreciationAmortizationDepletionIncomeStatement
}

// GetDepreciationAmortizationDepletionIncomeStatementOk returns a tuple with the DepreciationAmortizationDepletionIncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetDepreciationAmortizationDepletionIncomeStatementOk() (*float64, bool) {
	if o == nil || IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		return nil, false
	}
	return o.DepreciationAmortizationDepletionIncomeStatement, true
}

// HasDepreciationAmortizationDepletionIncomeStatement returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasDepreciationAmortizationDepletionIncomeStatement() bool {
	if o != nil && !IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		return true
	}

	return false
}

// SetDepreciationAmortizationDepletionIncomeStatement gets a reference to the given float64 and assigns it to the DepreciationAmortizationDepletionIncomeStatement field.
func (o *IncomeStatementItemExpenses) SetDepreciationAmortizationDepletionIncomeStatement(v float64) {
	o.DepreciationAmortizationDepletionIncomeStatement = &v
}

// GetResearchAndDevelopmentExpense returns the ResearchAndDevelopmentExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetResearchAndDevelopmentExpense() float64 {
	if o == nil || IsNil(o.ResearchAndDevelopmentExpense) {
		var ret float64
		return ret
	}
	return *o.ResearchAndDevelopmentExpense
}

// GetResearchAndDevelopmentExpenseOk returns a tuple with the ResearchAndDevelopmentExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetResearchAndDevelopmentExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.ResearchAndDevelopmentExpense) {
		return nil, false
	}
	return o.ResearchAndDevelopmentExpense, true
}

// HasResearchAndDevelopmentExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasResearchAndDevelopmentExpense() bool {
	if o != nil && !IsNil(o.ResearchAndDevelopmentExpense) {
		return true
	}

	return false
}

// SetResearchAndDevelopmentExpense gets a reference to the given float64 and assigns it to the ResearchAndDevelopmentExpense field.
func (o *IncomeStatementItemExpenses) SetResearchAndDevelopmentExpense(v float64) {
	o.ResearchAndDevelopmentExpense = &v
}

// GetInsuranceAndClaimsExpense returns the InsuranceAndClaimsExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetInsuranceAndClaimsExpense() float64 {
	if o == nil || IsNil(o.InsuranceAndClaimsExpense) {
		var ret float64
		return ret
	}
	return *o.InsuranceAndClaimsExpense
}

// GetInsuranceAndClaimsExpenseOk returns a tuple with the InsuranceAndClaimsExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetInsuranceAndClaimsExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.InsuranceAndClaimsExpense) {
		return nil, false
	}
	return o.InsuranceAndClaimsExpense, true
}

// HasInsuranceAndClaimsExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasInsuranceAndClaimsExpense() bool {
	if o != nil && !IsNil(o.InsuranceAndClaimsExpense) {
		return true
	}

	return false
}

// SetInsuranceAndClaimsExpense gets a reference to the given float64 and assigns it to the InsuranceAndClaimsExpense field.
func (o *IncomeStatementItemExpenses) SetInsuranceAndClaimsExpense(v float64) {
	o.InsuranceAndClaimsExpense = &v
}

// GetRentAndLandingFees returns the RentAndLandingFees field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetRentAndLandingFees() float64 {
	if o == nil || IsNil(o.RentAndLandingFees) {
		var ret float64
		return ret
	}
	return *o.RentAndLandingFees
}

// GetRentAndLandingFeesOk returns a tuple with the RentAndLandingFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetRentAndLandingFeesOk() (*float64, bool) {
	if o == nil || IsNil(o.RentAndLandingFees) {
		return nil, false
	}
	return o.RentAndLandingFees, true
}

// HasRentAndLandingFees returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasRentAndLandingFees() bool {
	if o != nil && !IsNil(o.RentAndLandingFees) {
		return true
	}

	return false
}

// SetRentAndLandingFees gets a reference to the given float64 and assigns it to the RentAndLandingFees field.
func (o *IncomeStatementItemExpenses) SetRentAndLandingFees(v float64) {
	o.RentAndLandingFees = &v
}

// GetSalariesAndWagesExpense returns the SalariesAndWagesExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetSalariesAndWagesExpense() float64 {
	if o == nil || IsNil(o.SalariesAndWagesExpense) {
		var ret float64
		return ret
	}
	return *o.SalariesAndWagesExpense
}

// GetSalariesAndWagesExpenseOk returns a tuple with the SalariesAndWagesExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetSalariesAndWagesExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.SalariesAndWagesExpense) {
		return nil, false
	}
	return o.SalariesAndWagesExpense, true
}

// HasSalariesAndWagesExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasSalariesAndWagesExpense() bool {
	if o != nil && !IsNil(o.SalariesAndWagesExpense) {
		return true
	}

	return false
}

// SetSalariesAndWagesExpense gets a reference to the given float64 and assigns it to the SalariesAndWagesExpense field.
func (o *IncomeStatementItemExpenses) SetSalariesAndWagesExpense(v float64) {
	o.SalariesAndWagesExpense = &v
}

// GetRentExpenseSupplemental returns the RentExpenseSupplemental field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetRentExpenseSupplemental() float64 {
	if o == nil || IsNil(o.RentExpenseSupplemental) {
		var ret float64
		return ret
	}
	return *o.RentExpenseSupplemental
}

// GetRentExpenseSupplementalOk returns a tuple with the RentExpenseSupplemental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetRentExpenseSupplementalOk() (*float64, bool) {
	if o == nil || IsNil(o.RentExpenseSupplemental) {
		return nil, false
	}
	return o.RentExpenseSupplemental, true
}

// HasRentExpenseSupplemental returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasRentExpenseSupplemental() bool {
	if o != nil && !IsNil(o.RentExpenseSupplemental) {
		return true
	}

	return false
}

// SetRentExpenseSupplemental gets a reference to the given float64 and assigns it to the RentExpenseSupplemental field.
func (o *IncomeStatementItemExpenses) SetRentExpenseSupplemental(v float64) {
	o.RentExpenseSupplemental = &v
}

// GetProvisionForDoubtfulAccounts returns the ProvisionForDoubtfulAccounts field value if set, zero value otherwise.
func (o *IncomeStatementItemExpenses) GetProvisionForDoubtfulAccounts() float64 {
	if o == nil || IsNil(o.ProvisionForDoubtfulAccounts) {
		var ret float64
		return ret
	}
	return *o.ProvisionForDoubtfulAccounts
}

// GetProvisionForDoubtfulAccountsOk returns a tuple with the ProvisionForDoubtfulAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemExpenses) GetProvisionForDoubtfulAccountsOk() (*float64, bool) {
	if o == nil || IsNil(o.ProvisionForDoubtfulAccounts) {
		return nil, false
	}
	return o.ProvisionForDoubtfulAccounts, true
}

// HasProvisionForDoubtfulAccounts returns a boolean if a field has been set.
func (o *IncomeStatementItemExpenses) HasProvisionForDoubtfulAccounts() bool {
	if o != nil && !IsNil(o.ProvisionForDoubtfulAccounts) {
		return true
	}

	return false
}

// SetProvisionForDoubtfulAccounts gets a reference to the given float64 and assigns it to the ProvisionForDoubtfulAccounts field.
func (o *IncomeStatementItemExpenses) SetProvisionForDoubtfulAccounts(v float64) {
	o.ProvisionForDoubtfulAccounts = &v
}

func (o IncomeStatementItemExpenses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemExpenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalExpenses) {
		toSerialize["total_expenses"] = o.TotalExpenses
	}
	if !IsNil(o.SellingGeneralAndAdministrationExpense) {
		toSerialize["selling_general_and_administration_expense"] = o.SellingGeneralAndAdministrationExpense
	}
	if !IsNil(o.SellingAndMarketingExpense) {
		toSerialize["selling_and_marketing_expense"] = o.SellingAndMarketingExpense
	}
	if !IsNil(o.GeneralAndAdministrativeExpense) {
		toSerialize["general_and_administrative_expense"] = o.GeneralAndAdministrativeExpense
	}
	if !IsNil(o.OtherGeneralAndAdministrativeExpense) {
		toSerialize["other_general_and_administrative_expense"] = o.OtherGeneralAndAdministrativeExpense
	}
	if !IsNil(o.DepreciationAmortizationDepletionIncomeStatement) {
		toSerialize["depreciation_amortization_depletion_income_statement"] = o.DepreciationAmortizationDepletionIncomeStatement
	}
	if !IsNil(o.ResearchAndDevelopmentExpense) {
		toSerialize["research_and_development_expense"] = o.ResearchAndDevelopmentExpense
	}
	if !IsNil(o.InsuranceAndClaimsExpense) {
		toSerialize["insurance_and_claims_expense"] = o.InsuranceAndClaimsExpense
	}
	if !IsNil(o.RentAndLandingFees) {
		toSerialize["rent_and_landing_fees"] = o.RentAndLandingFees
	}
	if !IsNil(o.SalariesAndWagesExpense) {
		toSerialize["salaries_and_wages_expense"] = o.SalariesAndWagesExpense
	}
	if !IsNil(o.RentExpenseSupplemental) {
		toSerialize["rent_expense_supplemental"] = o.RentExpenseSupplemental
	}
	if !IsNil(o.ProvisionForDoubtfulAccounts) {
		toSerialize["provision_for_doubtful_accounts"] = o.ProvisionForDoubtfulAccounts
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemExpenses struct {
	value *IncomeStatementItemExpenses
	isSet bool
}

func (v NullableIncomeStatementItemExpenses) Get() *IncomeStatementItemExpenses {
	return v.value
}

func (v *NullableIncomeStatementItemExpenses) Set(val *IncomeStatementItemExpenses) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemExpenses) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemExpenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemExpenses(val *IncomeStatementItemExpenses) *NullableIncomeStatementItemExpenses {
	return &NullableIncomeStatementItemExpenses{value: val, isSet: true}
}

func (v NullableIncomeStatementItemExpenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemExpenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
