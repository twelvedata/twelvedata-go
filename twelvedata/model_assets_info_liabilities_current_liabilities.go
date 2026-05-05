// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoLiabilitiesCurrentLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoLiabilitiesCurrentLiabilities{}

// AssetsInfoLiabilitiesCurrentLiabilities Current liabilities information
type AssetsInfoLiabilitiesCurrentLiabilities struct {
	// Total current liabilities
	TotalCurrentLiabilities *float64 `json:"total_current_liabilities,omitempty"`
	// Current debt and capital lease obligation
	CurrentDebtAndCapitalLeaseObligation *float64 `json:"current_debt_and_capital_lease_obligation,omitempty"`
	// Current debt
	CurrentDebt *float64 `json:"current_debt,omitempty"`
	// Current capital lease obligation
	CurrentCapitalLeaseObligation *float64 `json:"current_capital_lease_obligation,omitempty"`
	// Other current borrowings
	OtherCurrentBorrowings *float64 `json:"other_current_borrowings,omitempty"`
	// Line of credit
	LineOfCredit *float64 `json:"line_of_credit,omitempty"`
	// Commercial paper
	CommercialPaper *float64 `json:"commercial_paper,omitempty"`
	// Current notes payable
	CurrentNotesPayable *float64 `json:"current_notes_payable,omitempty"`
	// Current provisions
	CurrentProvisions          *float64                                                           `json:"current_provisions,omitempty"`
	PayablesAndAccruedExpenses *AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses `json:"payables_and_accrued_expenses,omitempty"`
	// Pension and other post retirement benefit plans current
	PensionAndOtherPostRetirementBenefitPlansCurrent *float64 `json:"pension_and_other_post_retirement_benefit_plans_current,omitempty"`
	// Employee benefits
	EmployeeBenefits *float64 `json:"employee_benefits,omitempty"`
	// Current deferred liabilities
	CurrentDeferredLiabilities *float64 `json:"current_deferred_liabilities,omitempty"`
	// Current deferred revenue
	CurrentDeferredRevenue *float64 `json:"current_deferred_revenue,omitempty"`
	// Current deferred taxes liabilities
	CurrentDeferredTaxesLiabilities *float64 `json:"current_deferred_taxes_liabilities,omitempty"`
	// Other current liabilities
	OtherCurrentLiabilities *float64 `json:"other_current_liabilities,omitempty"`
	// Liabilities held for sale non current
	LiabilitiesHeldForSaleNonCurrent *float64 `json:"liabilities_held_for_sale_non_current,omitempty"`
}

// NewAssetsInfoLiabilitiesCurrentLiabilities instantiates a new AssetsInfoLiabilitiesCurrentLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoLiabilitiesCurrentLiabilities() *AssetsInfoLiabilitiesCurrentLiabilities {
	this := AssetsInfoLiabilitiesCurrentLiabilities{}
	return &this
}

// NewAssetsInfoLiabilitiesCurrentLiabilitiesWithDefaults instantiates a new AssetsInfoLiabilitiesCurrentLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoLiabilitiesCurrentLiabilitiesWithDefaults() *AssetsInfoLiabilitiesCurrentLiabilities {
	this := AssetsInfoLiabilitiesCurrentLiabilities{}
	return &this
}

// GetTotalCurrentLiabilities returns the TotalCurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetTotalCurrentLiabilities() float64 {
	if o == nil || IsNil(o.TotalCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.TotalCurrentLiabilities
}

// GetTotalCurrentLiabilitiesOk returns a tuple with the TotalCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetTotalCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCurrentLiabilities) {
		return nil, false
	}
	return o.TotalCurrentLiabilities, true
}

// HasTotalCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasTotalCurrentLiabilities() bool {
	if o != nil && !IsNil(o.TotalCurrentLiabilities) {
		return true
	}

	return false
}

// SetTotalCurrentLiabilities gets a reference to the given float64 and assigns it to the TotalCurrentLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetTotalCurrentLiabilities(v float64) {
	o.TotalCurrentLiabilities = &v
}

// GetCurrentDebtAndCapitalLeaseObligation returns the CurrentDebtAndCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebtAndCapitalLeaseObligation() float64 {
	if o == nil || IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		var ret float64
		return ret
	}
	return *o.CurrentDebtAndCapitalLeaseObligation
}

// GetCurrentDebtAndCapitalLeaseObligationOk returns a tuple with the CurrentDebtAndCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebtAndCapitalLeaseObligationOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		return nil, false
	}
	return o.CurrentDebtAndCapitalLeaseObligation, true
}

// HasCurrentDebtAndCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDebtAndCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetCurrentDebtAndCapitalLeaseObligation gets a reference to the given float64 and assigns it to the CurrentDebtAndCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDebtAndCapitalLeaseObligation(v float64) {
	o.CurrentDebtAndCapitalLeaseObligation = &v
}

// GetCurrentDebt returns the CurrentDebt field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebt() float64 {
	if o == nil || IsNil(o.CurrentDebt) {
		var ret float64
		return ret
	}
	return *o.CurrentDebt
}

// GetCurrentDebtOk returns a tuple with the CurrentDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDebt) {
		return nil, false
	}
	return o.CurrentDebt, true
}

// HasCurrentDebt returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDebt() bool {
	if o != nil && !IsNil(o.CurrentDebt) {
		return true
	}

	return false
}

// SetCurrentDebt gets a reference to the given float64 and assigns it to the CurrentDebt field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDebt(v float64) {
	o.CurrentDebt = &v
}

// GetCurrentCapitalLeaseObligation returns the CurrentCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentCapitalLeaseObligation() float64 {
	if o == nil || IsNil(o.CurrentCapitalLeaseObligation) {
		var ret float64
		return ret
	}
	return *o.CurrentCapitalLeaseObligation
}

// GetCurrentCapitalLeaseObligationOk returns a tuple with the CurrentCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentCapitalLeaseObligationOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentCapitalLeaseObligation) {
		return nil, false
	}
	return o.CurrentCapitalLeaseObligation, true
}

// HasCurrentCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.CurrentCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetCurrentCapitalLeaseObligation gets a reference to the given float64 and assigns it to the CurrentCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentCapitalLeaseObligation(v float64) {
	o.CurrentCapitalLeaseObligation = &v
}

// GetOtherCurrentBorrowings returns the OtherCurrentBorrowings field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentBorrowings() float64 {
	if o == nil || IsNil(o.OtherCurrentBorrowings) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentBorrowings
}

// GetOtherCurrentBorrowingsOk returns a tuple with the OtherCurrentBorrowings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentBorrowingsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentBorrowings) {
		return nil, false
	}
	return o.OtherCurrentBorrowings, true
}

// HasOtherCurrentBorrowings returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasOtherCurrentBorrowings() bool {
	if o != nil && !IsNil(o.OtherCurrentBorrowings) {
		return true
	}

	return false
}

// SetOtherCurrentBorrowings gets a reference to the given float64 and assigns it to the OtherCurrentBorrowings field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetOtherCurrentBorrowings(v float64) {
	o.OtherCurrentBorrowings = &v
}

// GetLineOfCredit returns the LineOfCredit field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLineOfCredit() float64 {
	if o == nil || IsNil(o.LineOfCredit) {
		var ret float64
		return ret
	}
	return *o.LineOfCredit
}

// GetLineOfCreditOk returns a tuple with the LineOfCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLineOfCreditOk() (*float64, bool) {
	if o == nil || IsNil(o.LineOfCredit) {
		return nil, false
	}
	return o.LineOfCredit, true
}

// HasLineOfCredit returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasLineOfCredit() bool {
	if o != nil && !IsNil(o.LineOfCredit) {
		return true
	}

	return false
}

// SetLineOfCredit gets a reference to the given float64 and assigns it to the LineOfCredit field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetLineOfCredit(v float64) {
	o.LineOfCredit = &v
}

// GetCommercialPaper returns the CommercialPaper field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCommercialPaper() float64 {
	if o == nil || IsNil(o.CommercialPaper) {
		var ret float64
		return ret
	}
	return *o.CommercialPaper
}

// GetCommercialPaperOk returns a tuple with the CommercialPaper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCommercialPaperOk() (*float64, bool) {
	if o == nil || IsNil(o.CommercialPaper) {
		return nil, false
	}
	return o.CommercialPaper, true
}

// HasCommercialPaper returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCommercialPaper() bool {
	if o != nil && !IsNil(o.CommercialPaper) {
		return true
	}

	return false
}

// SetCommercialPaper gets a reference to the given float64 and assigns it to the CommercialPaper field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCommercialPaper(v float64) {
	o.CommercialPaper = &v
}

// GetCurrentNotesPayable returns the CurrentNotesPayable field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentNotesPayable() float64 {
	if o == nil || IsNil(o.CurrentNotesPayable) {
		var ret float64
		return ret
	}
	return *o.CurrentNotesPayable
}

// GetCurrentNotesPayableOk returns a tuple with the CurrentNotesPayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentNotesPayableOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentNotesPayable) {
		return nil, false
	}
	return o.CurrentNotesPayable, true
}

// HasCurrentNotesPayable returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentNotesPayable() bool {
	if o != nil && !IsNil(o.CurrentNotesPayable) {
		return true
	}

	return false
}

// SetCurrentNotesPayable gets a reference to the given float64 and assigns it to the CurrentNotesPayable field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentNotesPayable(v float64) {
	o.CurrentNotesPayable = &v
}

// GetCurrentProvisions returns the CurrentProvisions field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentProvisions() float64 {
	if o == nil || IsNil(o.CurrentProvisions) {
		var ret float64
		return ret
	}
	return *o.CurrentProvisions
}

// GetCurrentProvisionsOk returns a tuple with the CurrentProvisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentProvisionsOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentProvisions) {
		return nil, false
	}
	return o.CurrentProvisions, true
}

// HasCurrentProvisions returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentProvisions() bool {
	if o != nil && !IsNil(o.CurrentProvisions) {
		return true
	}

	return false
}

// SetCurrentProvisions gets a reference to the given float64 and assigns it to the CurrentProvisions field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentProvisions(v float64) {
	o.CurrentProvisions = &v
}

// GetPayablesAndAccruedExpenses returns the PayablesAndAccruedExpenses field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPayablesAndAccruedExpenses() AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses {
	if o == nil || IsNil(o.PayablesAndAccruedExpenses) {
		var ret AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses
		return ret
	}
	return *o.PayablesAndAccruedExpenses
}

// GetPayablesAndAccruedExpensesOk returns a tuple with the PayablesAndAccruedExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPayablesAndAccruedExpensesOk() (*AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses, bool) {
	if o == nil || IsNil(o.PayablesAndAccruedExpenses) {
		return nil, false
	}
	return o.PayablesAndAccruedExpenses, true
}

// HasPayablesAndAccruedExpenses returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasPayablesAndAccruedExpenses() bool {
	if o != nil && !IsNil(o.PayablesAndAccruedExpenses) {
		return true
	}

	return false
}

// SetPayablesAndAccruedExpenses gets a reference to the given AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses and assigns it to the PayablesAndAccruedExpenses field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetPayablesAndAccruedExpenses(v AssetsInfoLiabilitiesCurrentLiabilitiesPayablesAndAccruedExpenses) {
	o.PayablesAndAccruedExpenses = &v
}

// GetPensionAndOtherPostRetirementBenefitPlansCurrent returns the PensionAndOtherPostRetirementBenefitPlansCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPensionAndOtherPostRetirementBenefitPlansCurrent() float64 {
	if o == nil || IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		var ret float64
		return ret
	}
	return *o.PensionAndOtherPostRetirementBenefitPlansCurrent
}

// GetPensionAndOtherPostRetirementBenefitPlansCurrentOk returns a tuple with the PensionAndOtherPostRetirementBenefitPlansCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetPensionAndOtherPostRetirementBenefitPlansCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		return nil, false
	}
	return o.PensionAndOtherPostRetirementBenefitPlansCurrent, true
}

// HasPensionAndOtherPostRetirementBenefitPlansCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasPensionAndOtherPostRetirementBenefitPlansCurrent() bool {
	if o != nil && !IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		return true
	}

	return false
}

// SetPensionAndOtherPostRetirementBenefitPlansCurrent gets a reference to the given float64 and assigns it to the PensionAndOtherPostRetirementBenefitPlansCurrent field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetPensionAndOtherPostRetirementBenefitPlansCurrent(v float64) {
	o.PensionAndOtherPostRetirementBenefitPlansCurrent = &v
}

// GetEmployeeBenefits returns the EmployeeBenefits field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetEmployeeBenefits() float64 {
	if o == nil || IsNil(o.EmployeeBenefits) {
		var ret float64
		return ret
	}
	return *o.EmployeeBenefits
}

// GetEmployeeBenefitsOk returns a tuple with the EmployeeBenefits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetEmployeeBenefitsOk() (*float64, bool) {
	if o == nil || IsNil(o.EmployeeBenefits) {
		return nil, false
	}
	return o.EmployeeBenefits, true
}

// HasEmployeeBenefits returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasEmployeeBenefits() bool {
	if o != nil && !IsNil(o.EmployeeBenefits) {
		return true
	}

	return false
}

// SetEmployeeBenefits gets a reference to the given float64 and assigns it to the EmployeeBenefits field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetEmployeeBenefits(v float64) {
	o.EmployeeBenefits = &v
}

// GetCurrentDeferredLiabilities returns the CurrentDeferredLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredLiabilities() float64 {
	if o == nil || IsNil(o.CurrentDeferredLiabilities) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredLiabilities
}

// GetCurrentDeferredLiabilitiesOk returns a tuple with the CurrentDeferredLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredLiabilities) {
		return nil, false
	}
	return o.CurrentDeferredLiabilities, true
}

// HasCurrentDeferredLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDeferredLiabilities() bool {
	if o != nil && !IsNil(o.CurrentDeferredLiabilities) {
		return true
	}

	return false
}

// SetCurrentDeferredLiabilities gets a reference to the given float64 and assigns it to the CurrentDeferredLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDeferredLiabilities(v float64) {
	o.CurrentDeferredLiabilities = &v
}

// GetCurrentDeferredRevenue returns the CurrentDeferredRevenue field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredRevenue() float64 {
	if o == nil || IsNil(o.CurrentDeferredRevenue) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredRevenue
}

// GetCurrentDeferredRevenueOk returns a tuple with the CurrentDeferredRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredRevenue) {
		return nil, false
	}
	return o.CurrentDeferredRevenue, true
}

// HasCurrentDeferredRevenue returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDeferredRevenue() bool {
	if o != nil && !IsNil(o.CurrentDeferredRevenue) {
		return true
	}

	return false
}

// SetCurrentDeferredRevenue gets a reference to the given float64 and assigns it to the CurrentDeferredRevenue field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDeferredRevenue(v float64) {
	o.CurrentDeferredRevenue = &v
}

// GetCurrentDeferredTaxesLiabilities returns the CurrentDeferredTaxesLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredTaxesLiabilities() float64 {
	if o == nil || IsNil(o.CurrentDeferredTaxesLiabilities) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredTaxesLiabilities
}

// GetCurrentDeferredTaxesLiabilitiesOk returns a tuple with the CurrentDeferredTaxesLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetCurrentDeferredTaxesLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredTaxesLiabilities) {
		return nil, false
	}
	return o.CurrentDeferredTaxesLiabilities, true
}

// HasCurrentDeferredTaxesLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasCurrentDeferredTaxesLiabilities() bool {
	if o != nil && !IsNil(o.CurrentDeferredTaxesLiabilities) {
		return true
	}

	return false
}

// SetCurrentDeferredTaxesLiabilities gets a reference to the given float64 and assigns it to the CurrentDeferredTaxesLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetCurrentDeferredTaxesLiabilities(v float64) {
	o.CurrentDeferredTaxesLiabilities = &v
}

// GetOtherCurrentLiabilities returns the OtherCurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentLiabilities() float64 {
	if o == nil || IsNil(o.OtherCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentLiabilities
}

// GetOtherCurrentLiabilitiesOk returns a tuple with the OtherCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetOtherCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentLiabilities) {
		return nil, false
	}
	return o.OtherCurrentLiabilities, true
}

// HasOtherCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasOtherCurrentLiabilities() bool {
	if o != nil && !IsNil(o.OtherCurrentLiabilities) {
		return true
	}

	return false
}

// SetOtherCurrentLiabilities gets a reference to the given float64 and assigns it to the OtherCurrentLiabilities field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetOtherCurrentLiabilities(v float64) {
	o.OtherCurrentLiabilities = &v
}

// GetLiabilitiesHeldForSaleNonCurrent returns the LiabilitiesHeldForSaleNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLiabilitiesHeldForSaleNonCurrent() float64 {
	if o == nil || IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		var ret float64
		return ret
	}
	return *o.LiabilitiesHeldForSaleNonCurrent
}

// GetLiabilitiesHeldForSaleNonCurrentOk returns a tuple with the LiabilitiesHeldForSaleNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) GetLiabilitiesHeldForSaleNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		return nil, false
	}
	return o.LiabilitiesHeldForSaleNonCurrent, true
}

// HasLiabilitiesHeldForSaleNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) HasLiabilitiesHeldForSaleNonCurrent() bool {
	if o != nil && !IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		return true
	}

	return false
}

// SetLiabilitiesHeldForSaleNonCurrent gets a reference to the given float64 and assigns it to the LiabilitiesHeldForSaleNonCurrent field.
func (o *AssetsInfoLiabilitiesCurrentLiabilities) SetLiabilitiesHeldForSaleNonCurrent(v float64) {
	o.LiabilitiesHeldForSaleNonCurrent = &v
}

func (o AssetsInfoLiabilitiesCurrentLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoLiabilitiesCurrentLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCurrentLiabilities) {
		toSerialize["total_current_liabilities"] = o.TotalCurrentLiabilities
	}
	if !IsNil(o.CurrentDebtAndCapitalLeaseObligation) {
		toSerialize["current_debt_and_capital_lease_obligation"] = o.CurrentDebtAndCapitalLeaseObligation
	}
	if !IsNil(o.CurrentDebt) {
		toSerialize["current_debt"] = o.CurrentDebt
	}
	if !IsNil(o.CurrentCapitalLeaseObligation) {
		toSerialize["current_capital_lease_obligation"] = o.CurrentCapitalLeaseObligation
	}
	if !IsNil(o.OtherCurrentBorrowings) {
		toSerialize["other_current_borrowings"] = o.OtherCurrentBorrowings
	}
	if !IsNil(o.LineOfCredit) {
		toSerialize["line_of_credit"] = o.LineOfCredit
	}
	if !IsNil(o.CommercialPaper) {
		toSerialize["commercial_paper"] = o.CommercialPaper
	}
	if !IsNil(o.CurrentNotesPayable) {
		toSerialize["current_notes_payable"] = o.CurrentNotesPayable
	}
	if !IsNil(o.CurrentProvisions) {
		toSerialize["current_provisions"] = o.CurrentProvisions
	}
	if !IsNil(o.PayablesAndAccruedExpenses) {
		toSerialize["payables_and_accrued_expenses"] = o.PayablesAndAccruedExpenses
	}
	if !IsNil(o.PensionAndOtherPostRetirementBenefitPlansCurrent) {
		toSerialize["pension_and_other_post_retirement_benefit_plans_current"] = o.PensionAndOtherPostRetirementBenefitPlansCurrent
	}
	if !IsNil(o.EmployeeBenefits) {
		toSerialize["employee_benefits"] = o.EmployeeBenefits
	}
	if !IsNil(o.CurrentDeferredLiabilities) {
		toSerialize["current_deferred_liabilities"] = o.CurrentDeferredLiabilities
	}
	if !IsNil(o.CurrentDeferredRevenue) {
		toSerialize["current_deferred_revenue"] = o.CurrentDeferredRevenue
	}
	if !IsNil(o.CurrentDeferredTaxesLiabilities) {
		toSerialize["current_deferred_taxes_liabilities"] = o.CurrentDeferredTaxesLiabilities
	}
	if !IsNil(o.OtherCurrentLiabilities) {
		toSerialize["other_current_liabilities"] = o.OtherCurrentLiabilities
	}
	if !IsNil(o.LiabilitiesHeldForSaleNonCurrent) {
		toSerialize["liabilities_held_for_sale_non_current"] = o.LiabilitiesHeldForSaleNonCurrent
	}
	return toSerialize, nil
}

type NullableAssetsInfoLiabilitiesCurrentLiabilities struct {
	value *AssetsInfoLiabilitiesCurrentLiabilities
	isSet bool
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilities) Get() *AssetsInfoLiabilitiesCurrentLiabilities {
	return v.value
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilities) Set(val *AssetsInfoLiabilitiesCurrentLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoLiabilitiesCurrentLiabilities(val *AssetsInfoLiabilitiesCurrentLiabilities) *NullableAssetsInfoLiabilitiesCurrentLiabilities {
	return &NullableAssetsInfoLiabilitiesCurrentLiabilities{value: val, isSet: true}
}

func (v NullableAssetsInfoLiabilitiesCurrentLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoLiabilitiesCurrentLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
