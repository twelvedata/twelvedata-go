/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoLiabilitiesNonCurrentLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoLiabilitiesNonCurrentLiabilities{}

// AssetsInfoLiabilitiesNonCurrentLiabilities Non-current liabilities information
type AssetsInfoLiabilitiesNonCurrentLiabilities struct {
	// Total non current liabilities net minority interest
	TotalNonCurrentLiabilitiesNetMinorityInterest *float64                                                                         `json:"total_non_current_liabilities_net_minority_interest,omitempty"`
	LongTermDebtAndCapitalLeaseObligation         *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation `json:"long_term_debt_and_capital_lease_obligation,omitempty"`
	// Long term provisions
	LongTermProvisions *float64 `json:"long_term_provisions,omitempty"`
	// Non current pension and other postretirement benefit plans
	NonCurrentPensionAndOtherPostretirementBenefitPlans *float64 `json:"non_current_pension_and_other_postretirement_benefit_plans,omitempty"`
	// Non current accrued expenses
	NonCurrentAccruedExpenses *float64 `json:"non_current_accrued_expenses,omitempty"`
	// Due to related parties non current
	DueToRelatedPartiesNonCurrent *float64 `json:"due_to_related_parties_non_current,omitempty"`
	// Trade and other payables non current
	TradeAndOtherPayablesNonCurrent *float64 `json:"trade_and_other_payables_non_current,omitempty"`
	// Non current deferred liabilities
	NonCurrentDeferredLiabilities *float64 `json:"non_current_deferred_liabilities,omitempty"`
	// Non current deferred revenue
	NonCurrentDeferredRevenue *float64 `json:"non_current_deferred_revenue,omitempty"`
	// Non current deferred taxes liabilities
	NonCurrentDeferredTaxesLiabilities *float64 `json:"non_current_deferred_taxes_liabilities,omitempty"`
	// Other non current liabilities
	OtherNonCurrentLiabilities *float64 `json:"other_non_current_liabilities,omitempty"`
	// Preferred securities outside stock equity
	PreferredSecuritiesOutsideStockEquity *float64 `json:"preferred_securities_outside_stock_equity,omitempty"`
	// Derivative product liabilities
	DerivativeProductLiabilities *float64 `json:"derivative_product_liabilities,omitempty"`
	// Capital lease obligations
	CapitalLeaseObligations *float64 `json:"capital_lease_obligations,omitempty"`
	// Restricted common stock
	RestrictedCommonStock *float64 `json:"restricted_common_stock,omitempty"`
}

// NewAssetsInfoLiabilitiesNonCurrentLiabilities instantiates a new AssetsInfoLiabilitiesNonCurrentLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoLiabilitiesNonCurrentLiabilities() *AssetsInfoLiabilitiesNonCurrentLiabilities {
	this := AssetsInfoLiabilitiesNonCurrentLiabilities{}
	return &this
}

// NewAssetsInfoLiabilitiesNonCurrentLiabilitiesWithDefaults instantiates a new AssetsInfoLiabilitiesNonCurrentLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoLiabilitiesNonCurrentLiabilitiesWithDefaults() *AssetsInfoLiabilitiesNonCurrentLiabilities {
	this := AssetsInfoLiabilitiesNonCurrentLiabilities{}
	return &this
}

// GetTotalNonCurrentLiabilitiesNetMinorityInterest returns the TotalNonCurrentLiabilitiesNetMinorityInterest field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilitiesNetMinorityInterest() float64 {
	if o == nil || IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentLiabilitiesNetMinorityInterest
}

// GetTotalNonCurrentLiabilitiesNetMinorityInterestOk returns a tuple with the TotalNonCurrentLiabilitiesNetMinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilitiesNetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		return nil, false
	}
	return o.TotalNonCurrentLiabilitiesNetMinorityInterest, true
}

// HasTotalNonCurrentLiabilitiesNetMinorityInterest returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasTotalNonCurrentLiabilitiesNetMinorityInterest() bool {
	if o != nil && !IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		return true
	}

	return false
}

// SetTotalNonCurrentLiabilitiesNetMinorityInterest gets a reference to the given float64 and assigns it to the TotalNonCurrentLiabilitiesNetMinorityInterest field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetTotalNonCurrentLiabilitiesNetMinorityInterest(v float64) {
	o.TotalNonCurrentLiabilitiesNetMinorityInterest = &v
}

// GetLongTermDebtAndCapitalLeaseObligation returns the LongTermDebtAndCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermDebtAndCapitalLeaseObligation() AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation {
	if o == nil || IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		var ret AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation
		return ret
	}
	return *o.LongTermDebtAndCapitalLeaseObligation
}

// GetLongTermDebtAndCapitalLeaseObligationOk returns a tuple with the LongTermDebtAndCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermDebtAndCapitalLeaseObligationOk() (*AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation, bool) {
	if o == nil || IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		return nil, false
	}
	return o.LongTermDebtAndCapitalLeaseObligation, true
}

// HasLongTermDebtAndCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasLongTermDebtAndCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetLongTermDebtAndCapitalLeaseObligation gets a reference to the given AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation and assigns it to the LongTermDebtAndCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetLongTermDebtAndCapitalLeaseObligation(v AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) {
	o.LongTermDebtAndCapitalLeaseObligation = &v
}

// GetLongTermProvisions returns the LongTermProvisions field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermProvisions() float64 {
	if o == nil || IsNil(o.LongTermProvisions) {
		var ret float64
		return ret
	}
	return *o.LongTermProvisions
}

// GetLongTermProvisionsOk returns a tuple with the LongTermProvisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetLongTermProvisionsOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermProvisions) {
		return nil, false
	}
	return o.LongTermProvisions, true
}

// HasLongTermProvisions returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasLongTermProvisions() bool {
	if o != nil && !IsNil(o.LongTermProvisions) {
		return true
	}

	return false
}

// SetLongTermProvisions gets a reference to the given float64 and assigns it to the LongTermProvisions field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetLongTermProvisions(v float64) {
	o.LongTermProvisions = &v
}

// GetNonCurrentPensionAndOtherPostretirementBenefitPlans returns the NonCurrentPensionAndOtherPostretirementBenefitPlans field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentPensionAndOtherPostretirementBenefitPlans() float64 {
	if o == nil || IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		var ret float64
		return ret
	}
	return *o.NonCurrentPensionAndOtherPostretirementBenefitPlans
}

// GetNonCurrentPensionAndOtherPostretirementBenefitPlansOk returns a tuple with the NonCurrentPensionAndOtherPostretirementBenefitPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentPensionAndOtherPostretirementBenefitPlansOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		return nil, false
	}
	return o.NonCurrentPensionAndOtherPostretirementBenefitPlans, true
}

// HasNonCurrentPensionAndOtherPostretirementBenefitPlans returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentPensionAndOtherPostretirementBenefitPlans() bool {
	if o != nil && !IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		return true
	}

	return false
}

// SetNonCurrentPensionAndOtherPostretirementBenefitPlans gets a reference to the given float64 and assigns it to the NonCurrentPensionAndOtherPostretirementBenefitPlans field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentPensionAndOtherPostretirementBenefitPlans(v float64) {
	o.NonCurrentPensionAndOtherPostretirementBenefitPlans = &v
}

// GetNonCurrentAccruedExpenses returns the NonCurrentAccruedExpenses field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentAccruedExpenses() float64 {
	if o == nil || IsNil(o.NonCurrentAccruedExpenses) {
		var ret float64
		return ret
	}
	return *o.NonCurrentAccruedExpenses
}

// GetNonCurrentAccruedExpensesOk returns a tuple with the NonCurrentAccruedExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentAccruedExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentAccruedExpenses) {
		return nil, false
	}
	return o.NonCurrentAccruedExpenses, true
}

// HasNonCurrentAccruedExpenses returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentAccruedExpenses() bool {
	if o != nil && !IsNil(o.NonCurrentAccruedExpenses) {
		return true
	}

	return false
}

// SetNonCurrentAccruedExpenses gets a reference to the given float64 and assigns it to the NonCurrentAccruedExpenses field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentAccruedExpenses(v float64) {
	o.NonCurrentAccruedExpenses = &v
}

// GetDueToRelatedPartiesNonCurrent returns the DueToRelatedPartiesNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDueToRelatedPartiesNonCurrent() float64 {
	if o == nil || IsNil(o.DueToRelatedPartiesNonCurrent) {
		var ret float64
		return ret
	}
	return *o.DueToRelatedPartiesNonCurrent
}

// GetDueToRelatedPartiesNonCurrentOk returns a tuple with the DueToRelatedPartiesNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDueToRelatedPartiesNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.DueToRelatedPartiesNonCurrent) {
		return nil, false
	}
	return o.DueToRelatedPartiesNonCurrent, true
}

// HasDueToRelatedPartiesNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasDueToRelatedPartiesNonCurrent() bool {
	if o != nil && !IsNil(o.DueToRelatedPartiesNonCurrent) {
		return true
	}

	return false
}

// SetDueToRelatedPartiesNonCurrent gets a reference to the given float64 and assigns it to the DueToRelatedPartiesNonCurrent field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetDueToRelatedPartiesNonCurrent(v float64) {
	o.DueToRelatedPartiesNonCurrent = &v
}

// GetTradeAndOtherPayablesNonCurrent returns the TradeAndOtherPayablesNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTradeAndOtherPayablesNonCurrent() float64 {
	if o == nil || IsNil(o.TradeAndOtherPayablesNonCurrent) {
		var ret float64
		return ret
	}
	return *o.TradeAndOtherPayablesNonCurrent
}

// GetTradeAndOtherPayablesNonCurrentOk returns a tuple with the TradeAndOtherPayablesNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetTradeAndOtherPayablesNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.TradeAndOtherPayablesNonCurrent) {
		return nil, false
	}
	return o.TradeAndOtherPayablesNonCurrent, true
}

// HasTradeAndOtherPayablesNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasTradeAndOtherPayablesNonCurrent() bool {
	if o != nil && !IsNil(o.TradeAndOtherPayablesNonCurrent) {
		return true
	}

	return false
}

// SetTradeAndOtherPayablesNonCurrent gets a reference to the given float64 and assigns it to the TradeAndOtherPayablesNonCurrent field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetTradeAndOtherPayablesNonCurrent(v float64) {
	o.TradeAndOtherPayablesNonCurrent = &v
}

// GetNonCurrentDeferredLiabilities returns the NonCurrentDeferredLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredLiabilities() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredLiabilities) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredLiabilities
}

// GetNonCurrentDeferredLiabilitiesOk returns a tuple with the NonCurrentDeferredLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredLiabilities) {
		return nil, false
	}
	return o.NonCurrentDeferredLiabilities, true
}

// HasNonCurrentDeferredLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentDeferredLiabilities() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredLiabilities) {
		return true
	}

	return false
}

// SetNonCurrentDeferredLiabilities gets a reference to the given float64 and assigns it to the NonCurrentDeferredLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentDeferredLiabilities(v float64) {
	o.NonCurrentDeferredLiabilities = &v
}

// GetNonCurrentDeferredRevenue returns the NonCurrentDeferredRevenue field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredRevenue() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredRevenue) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredRevenue
}

// GetNonCurrentDeferredRevenueOk returns a tuple with the NonCurrentDeferredRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredRevenue) {
		return nil, false
	}
	return o.NonCurrentDeferredRevenue, true
}

// HasNonCurrentDeferredRevenue returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentDeferredRevenue() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredRevenue) {
		return true
	}

	return false
}

// SetNonCurrentDeferredRevenue gets a reference to the given float64 and assigns it to the NonCurrentDeferredRevenue field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentDeferredRevenue(v float64) {
	o.NonCurrentDeferredRevenue = &v
}

// GetNonCurrentDeferredTaxesLiabilities returns the NonCurrentDeferredTaxesLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredTaxesLiabilities() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredTaxesLiabilities
}

// GetNonCurrentDeferredTaxesLiabilitiesOk returns a tuple with the NonCurrentDeferredTaxesLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetNonCurrentDeferredTaxesLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		return nil, false
	}
	return o.NonCurrentDeferredTaxesLiabilities, true
}

// HasNonCurrentDeferredTaxesLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasNonCurrentDeferredTaxesLiabilities() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		return true
	}

	return false
}

// SetNonCurrentDeferredTaxesLiabilities gets a reference to the given float64 and assigns it to the NonCurrentDeferredTaxesLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetNonCurrentDeferredTaxesLiabilities(v float64) {
	o.NonCurrentDeferredTaxesLiabilities = &v
}

// GetOtherNonCurrentLiabilities returns the OtherNonCurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilities() float64 {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentLiabilities
}

// GetOtherNonCurrentLiabilitiesOk returns a tuple with the OtherNonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		return nil, false
	}
	return o.OtherNonCurrentLiabilities, true
}

// HasOtherNonCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasOtherNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.OtherNonCurrentLiabilities) {
		return true
	}

	return false
}

// SetOtherNonCurrentLiabilities gets a reference to the given float64 and assigns it to the OtherNonCurrentLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetOtherNonCurrentLiabilities(v float64) {
	o.OtherNonCurrentLiabilities = &v
}

// GetPreferredSecuritiesOutsideStockEquity returns the PreferredSecuritiesOutsideStockEquity field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetPreferredSecuritiesOutsideStockEquity() float64 {
	if o == nil || IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		var ret float64
		return ret
	}
	return *o.PreferredSecuritiesOutsideStockEquity
}

// GetPreferredSecuritiesOutsideStockEquityOk returns a tuple with the PreferredSecuritiesOutsideStockEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetPreferredSecuritiesOutsideStockEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		return nil, false
	}
	return o.PreferredSecuritiesOutsideStockEquity, true
}

// HasPreferredSecuritiesOutsideStockEquity returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasPreferredSecuritiesOutsideStockEquity() bool {
	if o != nil && !IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		return true
	}

	return false
}

// SetPreferredSecuritiesOutsideStockEquity gets a reference to the given float64 and assigns it to the PreferredSecuritiesOutsideStockEquity field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetPreferredSecuritiesOutsideStockEquity(v float64) {
	o.PreferredSecuritiesOutsideStockEquity = &v
}

// GetDerivativeProductLiabilities returns the DerivativeProductLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilities() float64 {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		var ret float64
		return ret
	}
	return *o.DerivativeProductLiabilities
}

// GetDerivativeProductLiabilitiesOk returns a tuple with the DerivativeProductLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		return nil, false
	}
	return o.DerivativeProductLiabilities, true
}

// HasDerivativeProductLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasDerivativeProductLiabilities() bool {
	if o != nil && !IsNil(o.DerivativeProductLiabilities) {
		return true
	}

	return false
}

// SetDerivativeProductLiabilities gets a reference to the given float64 and assigns it to the DerivativeProductLiabilities field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetDerivativeProductLiabilities(v float64) {
	o.DerivativeProductLiabilities = &v
}

// GetCapitalLeaseObligations returns the CapitalLeaseObligations field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetCapitalLeaseObligations() float64 {
	if o == nil || IsNil(o.CapitalLeaseObligations) {
		var ret float64
		return ret
	}
	return *o.CapitalLeaseObligations
}

// GetCapitalLeaseObligationsOk returns a tuple with the CapitalLeaseObligations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetCapitalLeaseObligationsOk() (*float64, bool) {
	if o == nil || IsNil(o.CapitalLeaseObligations) {
		return nil, false
	}
	return o.CapitalLeaseObligations, true
}

// HasCapitalLeaseObligations returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasCapitalLeaseObligations() bool {
	if o != nil && !IsNil(o.CapitalLeaseObligations) {
		return true
	}

	return false
}

// SetCapitalLeaseObligations gets a reference to the given float64 and assigns it to the CapitalLeaseObligations field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetCapitalLeaseObligations(v float64) {
	o.CapitalLeaseObligations = &v
}

// GetRestrictedCommonStock returns the RestrictedCommonStock field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetRestrictedCommonStock() float64 {
	if o == nil || IsNil(o.RestrictedCommonStock) {
		var ret float64
		return ret
	}
	return *o.RestrictedCommonStock
}

// GetRestrictedCommonStockOk returns a tuple with the RestrictedCommonStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) GetRestrictedCommonStockOk() (*float64, bool) {
	if o == nil || IsNil(o.RestrictedCommonStock) {
		return nil, false
	}
	return o.RestrictedCommonStock, true
}

// HasRestrictedCommonStock returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) HasRestrictedCommonStock() bool {
	if o != nil && !IsNil(o.RestrictedCommonStock) {
		return true
	}

	return false
}

// SetRestrictedCommonStock gets a reference to the given float64 and assigns it to the RestrictedCommonStock field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilities) SetRestrictedCommonStock(v float64) {
	o.RestrictedCommonStock = &v
}

func (o AssetsInfoLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoLiabilitiesNonCurrentLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalNonCurrentLiabilitiesNetMinorityInterest) {
		toSerialize["total_non_current_liabilities_net_minority_interest"] = o.TotalNonCurrentLiabilitiesNetMinorityInterest
	}
	if !IsNil(o.LongTermDebtAndCapitalLeaseObligation) {
		toSerialize["long_term_debt_and_capital_lease_obligation"] = o.LongTermDebtAndCapitalLeaseObligation
	}
	if !IsNil(o.LongTermProvisions) {
		toSerialize["long_term_provisions"] = o.LongTermProvisions
	}
	if !IsNil(o.NonCurrentPensionAndOtherPostretirementBenefitPlans) {
		toSerialize["non_current_pension_and_other_postretirement_benefit_plans"] = o.NonCurrentPensionAndOtherPostretirementBenefitPlans
	}
	if !IsNil(o.NonCurrentAccruedExpenses) {
		toSerialize["non_current_accrued_expenses"] = o.NonCurrentAccruedExpenses
	}
	if !IsNil(o.DueToRelatedPartiesNonCurrent) {
		toSerialize["due_to_related_parties_non_current"] = o.DueToRelatedPartiesNonCurrent
	}
	if !IsNil(o.TradeAndOtherPayablesNonCurrent) {
		toSerialize["trade_and_other_payables_non_current"] = o.TradeAndOtherPayablesNonCurrent
	}
	if !IsNil(o.NonCurrentDeferredLiabilities) {
		toSerialize["non_current_deferred_liabilities"] = o.NonCurrentDeferredLiabilities
	}
	if !IsNil(o.NonCurrentDeferredRevenue) {
		toSerialize["non_current_deferred_revenue"] = o.NonCurrentDeferredRevenue
	}
	if !IsNil(o.NonCurrentDeferredTaxesLiabilities) {
		toSerialize["non_current_deferred_taxes_liabilities"] = o.NonCurrentDeferredTaxesLiabilities
	}
	if !IsNil(o.OtherNonCurrentLiabilities) {
		toSerialize["other_non_current_liabilities"] = o.OtherNonCurrentLiabilities
	}
	if !IsNil(o.PreferredSecuritiesOutsideStockEquity) {
		toSerialize["preferred_securities_outside_stock_equity"] = o.PreferredSecuritiesOutsideStockEquity
	}
	if !IsNil(o.DerivativeProductLiabilities) {
		toSerialize["derivative_product_liabilities"] = o.DerivativeProductLiabilities
	}
	if !IsNil(o.CapitalLeaseObligations) {
		toSerialize["capital_lease_obligations"] = o.CapitalLeaseObligations
	}
	if !IsNil(o.RestrictedCommonStock) {
		toSerialize["restricted_common_stock"] = o.RestrictedCommonStock
	}
	return toSerialize, nil
}

type NullableAssetsInfoLiabilitiesNonCurrentLiabilities struct {
	value *AssetsInfoLiabilitiesNonCurrentLiabilities
	isSet bool
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilities) Get() *AssetsInfoLiabilitiesNonCurrentLiabilities {
	return v.value
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilities) Set(val *AssetsInfoLiabilitiesNonCurrentLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoLiabilitiesNonCurrentLiabilities(val *AssetsInfoLiabilitiesNonCurrentLiabilities) *NullableAssetsInfoLiabilitiesNonCurrentLiabilities {
	return &NullableAssetsInfoLiabilitiesNonCurrentLiabilities{value: val, isSet: true}
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
