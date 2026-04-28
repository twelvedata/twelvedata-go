/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoNonCurrentAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoNonCurrentAssets{}

// AssetsInfoNonCurrentAssets Non-current assets information
type AssetsInfoNonCurrentAssets struct {
	// Total non current assets
	TotalNonCurrentAssets *float64 `json:"total_non_current_assets,omitempty"`
	// Financial assets
	FinancialAssets *float64 `json:"financial_assets,omitempty"`
	// Investments and advances
	InvestmentsAndAdvances *float64 `json:"investments_and_advances,omitempty"`
	// Other investments
	OtherInvestments *float64 `json:"other_investments,omitempty"`
	// Investment in financial assets
	InvestmentInFinancialAssets *float64 `json:"investment_in_financial_assets,omitempty"`
	// Held to maturity securities
	HeldToMaturitySecurities *float64 `json:"held_to_maturity_securities,omitempty"`
	// Available for sale securities
	AvailableForSaleSecurities *float64 `json:"available_for_sale_securities,omitempty"`
	// Financial assets designated as fair value through profit or loss total
	FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal *float64 `json:"financial_assets_designated_as_fair_value_through_profit_or_loss_total,omitempty"`
	// Trading securities
	TradingSecurities *float64 `json:"trading_securities,omitempty"`
	// Long term equity investment
	LongTermEquityInvestment *float64 `json:"long_term_equity_investment,omitempty"`
	// Investments in joint ventures at cost
	InvestmentsInJointVenturesAtCost *float64 `json:"investments_in_joint_ventures_at_cost,omitempty"`
	// Investments in other ventures under equity method
	InvestmentsInOtherVenturesUnderEquityMethod *float64 `json:"investments_in_other_ventures_under_equity_method,omitempty"`
	// Investments in associates at cost
	InvestmentsInAssociatesAtCost *float64 `json:"investments_in_associates_at_cost,omitempty"`
	// Investments in subsidiaries at cost
	InvestmentsInSubsidiariesAtCost *float64 `json:"investments_in_subsidiaries_at_cost,omitempty"`
	// Investment properties
	InvestmentProperties             *float64                                                    `json:"investment_properties,omitempty"`
	GoodwillAndOtherIntangibleAssets *AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets `json:"goodwill_and_other_intangible_assets,omitempty"`
	// Net ppe
	NetPpe *float64 `json:"net_ppe,omitempty"`
	// Gross ppe
	GrossPpe *float64 `json:"gross_ppe,omitempty"`
	// Accumulated depreciation
	AccumulatedDepreciation *float64 `json:"accumulated_depreciation,omitempty"`
	// Leases
	Leases *float64 `json:"leases,omitempty"`
	// Construction in progress
	ConstructionInProgress *float64 `json:"construction_in_progress,omitempty"`
	// Other properties
	OtherProperties *float64 `json:"other_properties,omitempty"`
	// Machinery furniture equipment
	MachineryFurnitureEquipment *float64 `json:"machinery_furniture_equipment,omitempty"`
	// Buildings and improvements
	BuildingsAndImprovements *float64 `json:"buildings_and_improvements,omitempty"`
	// Land and improvements
	LandAndImprovements *float64 `json:"land_and_improvements,omitempty"`
	// Properties
	Properties *float64 `json:"properties,omitempty"`
	// Non current accounts receivable
	NonCurrentAccountsReceivable *float64 `json:"non_current_accounts_receivable,omitempty"`
	// Non current note receivables
	NonCurrentNoteReceivables *float64 `json:"non_current_note_receivables,omitempty"`
	// Due from related parties non current
	DueFromRelatedPartiesNonCurrent *float64 `json:"due_from_related_parties_non_current,omitempty"`
	// Non current prepaid assets
	NonCurrentPrepaidAssets *float64 `json:"non_current_prepaid_assets,omitempty"`
	// Non current deferred assets
	NonCurrentDeferredAssets *float64 `json:"non_current_deferred_assets,omitempty"`
	// Non current deferred taxes assets
	NonCurrentDeferredTaxesAssets *float64 `json:"non_current_deferred_taxes_assets,omitempty"`
	// Defined pension benefit
	DefinedPensionBenefit *float64 `json:"defined_pension_benefit,omitempty"`
	// Other non current assets
	OtherNonCurrentAssets *float64 `json:"other_non_current_assets,omitempty"`
}

// NewAssetsInfoNonCurrentAssets instantiates a new AssetsInfoNonCurrentAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoNonCurrentAssets() *AssetsInfoNonCurrentAssets {
	this := AssetsInfoNonCurrentAssets{}
	return &this
}

// NewAssetsInfoNonCurrentAssetsWithDefaults instantiates a new AssetsInfoNonCurrentAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoNonCurrentAssetsWithDefaults() *AssetsInfoNonCurrentAssets {
	this := AssetsInfoNonCurrentAssets{}
	return &this
}

// GetTotalNonCurrentAssets returns the TotalNonCurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetTotalNonCurrentAssets() float64 {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentAssets
}

// GetTotalNonCurrentAssetsOk returns a tuple with the TotalNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetTotalNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentAssets) {
		return nil, false
	}
	return o.TotalNonCurrentAssets, true
}

// HasTotalNonCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasTotalNonCurrentAssets() bool {
	if o != nil && !IsNil(o.TotalNonCurrentAssets) {
		return true
	}

	return false
}

// SetTotalNonCurrentAssets gets a reference to the given float64 and assigns it to the TotalNonCurrentAssets field.
func (o *AssetsInfoNonCurrentAssets) SetTotalNonCurrentAssets(v float64) {
	o.TotalNonCurrentAssets = &v
}

// GetFinancialAssets returns the FinancialAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssets() float64 {
	if o == nil || IsNil(o.FinancialAssets) {
		var ret float64
		return ret
	}
	return *o.FinancialAssets
}

// GetFinancialAssetsOk returns a tuple with the FinancialAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancialAssets) {
		return nil, false
	}
	return o.FinancialAssets, true
}

// HasFinancialAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasFinancialAssets() bool {
	if o != nil && !IsNil(o.FinancialAssets) {
		return true
	}

	return false
}

// SetFinancialAssets gets a reference to the given float64 and assigns it to the FinancialAssets field.
func (o *AssetsInfoNonCurrentAssets) SetFinancialAssets(v float64) {
	o.FinancialAssets = &v
}

// GetInvestmentsAndAdvances returns the InvestmentsAndAdvances field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsAndAdvances() float64 {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		var ret float64
		return ret
	}
	return *o.InvestmentsAndAdvances
}

// GetInvestmentsAndAdvancesOk returns a tuple with the InvestmentsAndAdvances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsAndAdvancesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsAndAdvances) {
		return nil, false
	}
	return o.InvestmentsAndAdvances, true
}

// HasInvestmentsAndAdvances returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsAndAdvances() bool {
	if o != nil && !IsNil(o.InvestmentsAndAdvances) {
		return true
	}

	return false
}

// SetInvestmentsAndAdvances gets a reference to the given float64 and assigns it to the InvestmentsAndAdvances field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsAndAdvances(v float64) {
	o.InvestmentsAndAdvances = &v
}

// GetOtherInvestments returns the OtherInvestments field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetOtherInvestments() float64 {
	if o == nil || IsNil(o.OtherInvestments) {
		var ret float64
		return ret
	}
	return *o.OtherInvestments
}

// GetOtherInvestmentsOk returns a tuple with the OtherInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetOtherInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherInvestments) {
		return nil, false
	}
	return o.OtherInvestments, true
}

// HasOtherInvestments returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasOtherInvestments() bool {
	if o != nil && !IsNil(o.OtherInvestments) {
		return true
	}

	return false
}

// SetOtherInvestments gets a reference to the given float64 and assigns it to the OtherInvestments field.
func (o *AssetsInfoNonCurrentAssets) SetOtherInvestments(v float64) {
	o.OtherInvestments = &v
}

// GetInvestmentInFinancialAssets returns the InvestmentInFinancialAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentInFinancialAssets() float64 {
	if o == nil || IsNil(o.InvestmentInFinancialAssets) {
		var ret float64
		return ret
	}
	return *o.InvestmentInFinancialAssets
}

// GetInvestmentInFinancialAssetsOk returns a tuple with the InvestmentInFinancialAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentInFinancialAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentInFinancialAssets) {
		return nil, false
	}
	return o.InvestmentInFinancialAssets, true
}

// HasInvestmentInFinancialAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentInFinancialAssets() bool {
	if o != nil && !IsNil(o.InvestmentInFinancialAssets) {
		return true
	}

	return false
}

// SetInvestmentInFinancialAssets gets a reference to the given float64 and assigns it to the InvestmentInFinancialAssets field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentInFinancialAssets(v float64) {
	o.InvestmentInFinancialAssets = &v
}

// GetHeldToMaturitySecurities returns the HeldToMaturitySecurities field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetHeldToMaturitySecurities() float64 {
	if o == nil || IsNil(o.HeldToMaturitySecurities) {
		var ret float64
		return ret
	}
	return *o.HeldToMaturitySecurities
}

// GetHeldToMaturitySecuritiesOk returns a tuple with the HeldToMaturitySecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetHeldToMaturitySecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.HeldToMaturitySecurities) {
		return nil, false
	}
	return o.HeldToMaturitySecurities, true
}

// HasHeldToMaturitySecurities returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasHeldToMaturitySecurities() bool {
	if o != nil && !IsNil(o.HeldToMaturitySecurities) {
		return true
	}

	return false
}

// SetHeldToMaturitySecurities gets a reference to the given float64 and assigns it to the HeldToMaturitySecurities field.
func (o *AssetsInfoNonCurrentAssets) SetHeldToMaturitySecurities(v float64) {
	o.HeldToMaturitySecurities = &v
}

// GetAvailableForSaleSecurities returns the AvailableForSaleSecurities field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetAvailableForSaleSecurities() float64 {
	if o == nil || IsNil(o.AvailableForSaleSecurities) {
		var ret float64
		return ret
	}
	return *o.AvailableForSaleSecurities
}

// GetAvailableForSaleSecuritiesOk returns a tuple with the AvailableForSaleSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetAvailableForSaleSecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.AvailableForSaleSecurities) {
		return nil, false
	}
	return o.AvailableForSaleSecurities, true
}

// HasAvailableForSaleSecurities returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasAvailableForSaleSecurities() bool {
	if o != nil && !IsNil(o.AvailableForSaleSecurities) {
		return true
	}

	return false
}

// SetAvailableForSaleSecurities gets a reference to the given float64 and assigns it to the AvailableForSaleSecurities field.
func (o *AssetsInfoNonCurrentAssets) SetAvailableForSaleSecurities(v float64) {
	o.AvailableForSaleSecurities = &v
}

// GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal returns the FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal() float64 {
	if o == nil || IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		var ret float64
		return ret
	}
	return *o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal
}

// GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotalOk returns a tuple with the FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		return nil, false
	}
	return o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal, true
}

// HasFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal() bool {
	if o != nil && !IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		return true
	}

	return false
}

// SetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal gets a reference to the given float64 and assigns it to the FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal field.
func (o *AssetsInfoNonCurrentAssets) SetFinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal(v float64) {
	o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal = &v
}

// GetTradingSecurities returns the TradingSecurities field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetTradingSecurities() float64 {
	if o == nil || IsNil(o.TradingSecurities) {
		var ret float64
		return ret
	}
	return *o.TradingSecurities
}

// GetTradingSecuritiesOk returns a tuple with the TradingSecurities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetTradingSecuritiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TradingSecurities) {
		return nil, false
	}
	return o.TradingSecurities, true
}

// HasTradingSecurities returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasTradingSecurities() bool {
	if o != nil && !IsNil(o.TradingSecurities) {
		return true
	}

	return false
}

// SetTradingSecurities gets a reference to the given float64 and assigns it to the TradingSecurities field.
func (o *AssetsInfoNonCurrentAssets) SetTradingSecurities(v float64) {
	o.TradingSecurities = &v
}

// GetLongTermEquityInvestment returns the LongTermEquityInvestment field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetLongTermEquityInvestment() float64 {
	if o == nil || IsNil(o.LongTermEquityInvestment) {
		var ret float64
		return ret
	}
	return *o.LongTermEquityInvestment
}

// GetLongTermEquityInvestmentOk returns a tuple with the LongTermEquityInvestment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetLongTermEquityInvestmentOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermEquityInvestment) {
		return nil, false
	}
	return o.LongTermEquityInvestment, true
}

// HasLongTermEquityInvestment returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasLongTermEquityInvestment() bool {
	if o != nil && !IsNil(o.LongTermEquityInvestment) {
		return true
	}

	return false
}

// SetLongTermEquityInvestment gets a reference to the given float64 and assigns it to the LongTermEquityInvestment field.
func (o *AssetsInfoNonCurrentAssets) SetLongTermEquityInvestment(v float64) {
	o.LongTermEquityInvestment = &v
}

// GetInvestmentsInJointVenturesAtCost returns the InvestmentsInJointVenturesAtCost field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInJointVenturesAtCost() float64 {
	if o == nil || IsNil(o.InvestmentsInJointVenturesAtCost) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInJointVenturesAtCost
}

// GetInvestmentsInJointVenturesAtCostOk returns a tuple with the InvestmentsInJointVenturesAtCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInJointVenturesAtCostOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInJointVenturesAtCost) {
		return nil, false
	}
	return o.InvestmentsInJointVenturesAtCost, true
}

// HasInvestmentsInJointVenturesAtCost returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInJointVenturesAtCost() bool {
	if o != nil && !IsNil(o.InvestmentsInJointVenturesAtCost) {
		return true
	}

	return false
}

// SetInvestmentsInJointVenturesAtCost gets a reference to the given float64 and assigns it to the InvestmentsInJointVenturesAtCost field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInJointVenturesAtCost(v float64) {
	o.InvestmentsInJointVenturesAtCost = &v
}

// GetInvestmentsInOtherVenturesUnderEquityMethod returns the InvestmentsInOtherVenturesUnderEquityMethod field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInOtherVenturesUnderEquityMethod() float64 {
	if o == nil || IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInOtherVenturesUnderEquityMethod
}

// GetInvestmentsInOtherVenturesUnderEquityMethodOk returns a tuple with the InvestmentsInOtherVenturesUnderEquityMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInOtherVenturesUnderEquityMethodOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		return nil, false
	}
	return o.InvestmentsInOtherVenturesUnderEquityMethod, true
}

// HasInvestmentsInOtherVenturesUnderEquityMethod returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInOtherVenturesUnderEquityMethod() bool {
	if o != nil && !IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		return true
	}

	return false
}

// SetInvestmentsInOtherVenturesUnderEquityMethod gets a reference to the given float64 and assigns it to the InvestmentsInOtherVenturesUnderEquityMethod field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInOtherVenturesUnderEquityMethod(v float64) {
	o.InvestmentsInOtherVenturesUnderEquityMethod = &v
}

// GetInvestmentsInAssociatesAtCost returns the InvestmentsInAssociatesAtCost field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInAssociatesAtCost() float64 {
	if o == nil || IsNil(o.InvestmentsInAssociatesAtCost) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInAssociatesAtCost
}

// GetInvestmentsInAssociatesAtCostOk returns a tuple with the InvestmentsInAssociatesAtCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInAssociatesAtCostOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInAssociatesAtCost) {
		return nil, false
	}
	return o.InvestmentsInAssociatesAtCost, true
}

// HasInvestmentsInAssociatesAtCost returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInAssociatesAtCost() bool {
	if o != nil && !IsNil(o.InvestmentsInAssociatesAtCost) {
		return true
	}

	return false
}

// SetInvestmentsInAssociatesAtCost gets a reference to the given float64 and assigns it to the InvestmentsInAssociatesAtCost field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInAssociatesAtCost(v float64) {
	o.InvestmentsInAssociatesAtCost = &v
}

// GetInvestmentsInSubsidiariesAtCost returns the InvestmentsInSubsidiariesAtCost field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInSubsidiariesAtCost() float64 {
	if o == nil || IsNil(o.InvestmentsInSubsidiariesAtCost) {
		var ret float64
		return ret
	}
	return *o.InvestmentsInSubsidiariesAtCost
}

// GetInvestmentsInSubsidiariesAtCostOk returns a tuple with the InvestmentsInSubsidiariesAtCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentsInSubsidiariesAtCostOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentsInSubsidiariesAtCost) {
		return nil, false
	}
	return o.InvestmentsInSubsidiariesAtCost, true
}

// HasInvestmentsInSubsidiariesAtCost returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentsInSubsidiariesAtCost() bool {
	if o != nil && !IsNil(o.InvestmentsInSubsidiariesAtCost) {
		return true
	}

	return false
}

// SetInvestmentsInSubsidiariesAtCost gets a reference to the given float64 and assigns it to the InvestmentsInSubsidiariesAtCost field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentsInSubsidiariesAtCost(v float64) {
	o.InvestmentsInSubsidiariesAtCost = &v
}

// GetInvestmentProperties returns the InvestmentProperties field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentProperties() float64 {
	if o == nil || IsNil(o.InvestmentProperties) {
		var ret float64
		return ret
	}
	return *o.InvestmentProperties
}

// GetInvestmentPropertiesOk returns a tuple with the InvestmentProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetInvestmentPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestmentProperties) {
		return nil, false
	}
	return o.InvestmentProperties, true
}

// HasInvestmentProperties returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasInvestmentProperties() bool {
	if o != nil && !IsNil(o.InvestmentProperties) {
		return true
	}

	return false
}

// SetInvestmentProperties gets a reference to the given float64 and assigns it to the InvestmentProperties field.
func (o *AssetsInfoNonCurrentAssets) SetInvestmentProperties(v float64) {
	o.InvestmentProperties = &v
}

// GetGoodwillAndOtherIntangibleAssets returns the GoodwillAndOtherIntangibleAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetGoodwillAndOtherIntangibleAssets() AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets {
	if o == nil || IsNil(o.GoodwillAndOtherIntangibleAssets) {
		var ret AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets
		return ret
	}
	return *o.GoodwillAndOtherIntangibleAssets
}

// GetGoodwillAndOtherIntangibleAssetsOk returns a tuple with the GoodwillAndOtherIntangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetGoodwillAndOtherIntangibleAssetsOk() (*AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets, bool) {
	if o == nil || IsNil(o.GoodwillAndOtherIntangibleAssets) {
		return nil, false
	}
	return o.GoodwillAndOtherIntangibleAssets, true
}

// HasGoodwillAndOtherIntangibleAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasGoodwillAndOtherIntangibleAssets() bool {
	if o != nil && !IsNil(o.GoodwillAndOtherIntangibleAssets) {
		return true
	}

	return false
}

// SetGoodwillAndOtherIntangibleAssets gets a reference to the given AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets and assigns it to the GoodwillAndOtherIntangibleAssets field.
func (o *AssetsInfoNonCurrentAssets) SetGoodwillAndOtherIntangibleAssets(v AssetsInfoNonCurrentAssetsGoodwillAndOtherIntangibleAssets) {
	o.GoodwillAndOtherIntangibleAssets = &v
}

// GetNetPpe returns the NetPpe field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNetPpe() float64 {
	if o == nil || IsNil(o.NetPpe) {
		var ret float64
		return ret
	}
	return *o.NetPpe
}

// GetNetPpeOk returns a tuple with the NetPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNetPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.NetPpe) {
		return nil, false
	}
	return o.NetPpe, true
}

// HasNetPpe returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNetPpe() bool {
	if o != nil && !IsNil(o.NetPpe) {
		return true
	}

	return false
}

// SetNetPpe gets a reference to the given float64 and assigns it to the NetPpe field.
func (o *AssetsInfoNonCurrentAssets) SetNetPpe(v float64) {
	o.NetPpe = &v
}

// GetGrossPpe returns the GrossPpe field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetGrossPpe() float64 {
	if o == nil || IsNil(o.GrossPpe) {
		var ret float64
		return ret
	}
	return *o.GrossPpe
}

// GetGrossPpeOk returns a tuple with the GrossPpe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetGrossPpeOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossPpe) {
		return nil, false
	}
	return o.GrossPpe, true
}

// HasGrossPpe returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasGrossPpe() bool {
	if o != nil && !IsNil(o.GrossPpe) {
		return true
	}

	return false
}

// SetGrossPpe gets a reference to the given float64 and assigns it to the GrossPpe field.
func (o *AssetsInfoNonCurrentAssets) SetGrossPpe(v float64) {
	o.GrossPpe = &v
}

// GetAccumulatedDepreciation returns the AccumulatedDepreciation field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetAccumulatedDepreciation() float64 {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		var ret float64
		return ret
	}
	return *o.AccumulatedDepreciation
}

// GetAccumulatedDepreciationOk returns a tuple with the AccumulatedDepreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetAccumulatedDepreciationOk() (*float64, bool) {
	if o == nil || IsNil(o.AccumulatedDepreciation) {
		return nil, false
	}
	return o.AccumulatedDepreciation, true
}

// HasAccumulatedDepreciation returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasAccumulatedDepreciation() bool {
	if o != nil && !IsNil(o.AccumulatedDepreciation) {
		return true
	}

	return false
}

// SetAccumulatedDepreciation gets a reference to the given float64 and assigns it to the AccumulatedDepreciation field.
func (o *AssetsInfoNonCurrentAssets) SetAccumulatedDepreciation(v float64) {
	o.AccumulatedDepreciation = &v
}

// GetLeases returns the Leases field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetLeases() float64 {
	if o == nil || IsNil(o.Leases) {
		var ret float64
		return ret
	}
	return *o.Leases
}

// GetLeasesOk returns a tuple with the Leases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetLeasesOk() (*float64, bool) {
	if o == nil || IsNil(o.Leases) {
		return nil, false
	}
	return o.Leases, true
}

// HasLeases returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasLeases() bool {
	if o != nil && !IsNil(o.Leases) {
		return true
	}

	return false
}

// SetLeases gets a reference to the given float64 and assigns it to the Leases field.
func (o *AssetsInfoNonCurrentAssets) SetLeases(v float64) {
	o.Leases = &v
}

// GetConstructionInProgress returns the ConstructionInProgress field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetConstructionInProgress() float64 {
	if o == nil || IsNil(o.ConstructionInProgress) {
		var ret float64
		return ret
	}
	return *o.ConstructionInProgress
}

// GetConstructionInProgressOk returns a tuple with the ConstructionInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetConstructionInProgressOk() (*float64, bool) {
	if o == nil || IsNil(o.ConstructionInProgress) {
		return nil, false
	}
	return o.ConstructionInProgress, true
}

// HasConstructionInProgress returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasConstructionInProgress() bool {
	if o != nil && !IsNil(o.ConstructionInProgress) {
		return true
	}

	return false
}

// SetConstructionInProgress gets a reference to the given float64 and assigns it to the ConstructionInProgress field.
func (o *AssetsInfoNonCurrentAssets) SetConstructionInProgress(v float64) {
	o.ConstructionInProgress = &v
}

// GetOtherProperties returns the OtherProperties field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetOtherProperties() float64 {
	if o == nil || IsNil(o.OtherProperties) {
		var ret float64
		return ret
	}
	return *o.OtherProperties
}

// GetOtherPropertiesOk returns a tuple with the OtherProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetOtherPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherProperties) {
		return nil, false
	}
	return o.OtherProperties, true
}

// HasOtherProperties returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasOtherProperties() bool {
	if o != nil && !IsNil(o.OtherProperties) {
		return true
	}

	return false
}

// SetOtherProperties gets a reference to the given float64 and assigns it to the OtherProperties field.
func (o *AssetsInfoNonCurrentAssets) SetOtherProperties(v float64) {
	o.OtherProperties = &v
}

// GetMachineryFurnitureEquipment returns the MachineryFurnitureEquipment field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetMachineryFurnitureEquipment() float64 {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		var ret float64
		return ret
	}
	return *o.MachineryFurnitureEquipment
}

// GetMachineryFurnitureEquipmentOk returns a tuple with the MachineryFurnitureEquipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetMachineryFurnitureEquipmentOk() (*float64, bool) {
	if o == nil || IsNil(o.MachineryFurnitureEquipment) {
		return nil, false
	}
	return o.MachineryFurnitureEquipment, true
}

// HasMachineryFurnitureEquipment returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasMachineryFurnitureEquipment() bool {
	if o != nil && !IsNil(o.MachineryFurnitureEquipment) {
		return true
	}

	return false
}

// SetMachineryFurnitureEquipment gets a reference to the given float64 and assigns it to the MachineryFurnitureEquipment field.
func (o *AssetsInfoNonCurrentAssets) SetMachineryFurnitureEquipment(v float64) {
	o.MachineryFurnitureEquipment = &v
}

// GetBuildingsAndImprovements returns the BuildingsAndImprovements field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetBuildingsAndImprovements() float64 {
	if o == nil || IsNil(o.BuildingsAndImprovements) {
		var ret float64
		return ret
	}
	return *o.BuildingsAndImprovements
}

// GetBuildingsAndImprovementsOk returns a tuple with the BuildingsAndImprovements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetBuildingsAndImprovementsOk() (*float64, bool) {
	if o == nil || IsNil(o.BuildingsAndImprovements) {
		return nil, false
	}
	return o.BuildingsAndImprovements, true
}

// HasBuildingsAndImprovements returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasBuildingsAndImprovements() bool {
	if o != nil && !IsNil(o.BuildingsAndImprovements) {
		return true
	}

	return false
}

// SetBuildingsAndImprovements gets a reference to the given float64 and assigns it to the BuildingsAndImprovements field.
func (o *AssetsInfoNonCurrentAssets) SetBuildingsAndImprovements(v float64) {
	o.BuildingsAndImprovements = &v
}

// GetLandAndImprovements returns the LandAndImprovements field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetLandAndImprovements() float64 {
	if o == nil || IsNil(o.LandAndImprovements) {
		var ret float64
		return ret
	}
	return *o.LandAndImprovements
}

// GetLandAndImprovementsOk returns a tuple with the LandAndImprovements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetLandAndImprovementsOk() (*float64, bool) {
	if o == nil || IsNil(o.LandAndImprovements) {
		return nil, false
	}
	return o.LandAndImprovements, true
}

// HasLandAndImprovements returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasLandAndImprovements() bool {
	if o != nil && !IsNil(o.LandAndImprovements) {
		return true
	}

	return false
}

// SetLandAndImprovements gets a reference to the given float64 and assigns it to the LandAndImprovements field.
func (o *AssetsInfoNonCurrentAssets) SetLandAndImprovements(v float64) {
	o.LandAndImprovements = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetProperties() float64 {
	if o == nil || IsNil(o.Properties) {
		var ret float64
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetPropertiesOk() (*float64, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given float64 and assigns it to the Properties field.
func (o *AssetsInfoNonCurrentAssets) SetProperties(v float64) {
	o.Properties = &v
}

// GetNonCurrentAccountsReceivable returns the NonCurrentAccountsReceivable field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentAccountsReceivable() float64 {
	if o == nil || IsNil(o.NonCurrentAccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.NonCurrentAccountsReceivable
}

// GetNonCurrentAccountsReceivableOk returns a tuple with the NonCurrentAccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentAccountsReceivable) {
		return nil, false
	}
	return o.NonCurrentAccountsReceivable, true
}

// HasNonCurrentAccountsReceivable returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentAccountsReceivable() bool {
	if o != nil && !IsNil(o.NonCurrentAccountsReceivable) {
		return true
	}

	return false
}

// SetNonCurrentAccountsReceivable gets a reference to the given float64 and assigns it to the NonCurrentAccountsReceivable field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentAccountsReceivable(v float64) {
	o.NonCurrentAccountsReceivable = &v
}

// GetNonCurrentNoteReceivables returns the NonCurrentNoteReceivables field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentNoteReceivables() float64 {
	if o == nil || IsNil(o.NonCurrentNoteReceivables) {
		var ret float64
		return ret
	}
	return *o.NonCurrentNoteReceivables
}

// GetNonCurrentNoteReceivablesOk returns a tuple with the NonCurrentNoteReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentNoteReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentNoteReceivables) {
		return nil, false
	}
	return o.NonCurrentNoteReceivables, true
}

// HasNonCurrentNoteReceivables returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentNoteReceivables() bool {
	if o != nil && !IsNil(o.NonCurrentNoteReceivables) {
		return true
	}

	return false
}

// SetNonCurrentNoteReceivables gets a reference to the given float64 and assigns it to the NonCurrentNoteReceivables field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentNoteReceivables(v float64) {
	o.NonCurrentNoteReceivables = &v
}

// GetDueFromRelatedPartiesNonCurrent returns the DueFromRelatedPartiesNonCurrent field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetDueFromRelatedPartiesNonCurrent() float64 {
	if o == nil || IsNil(o.DueFromRelatedPartiesNonCurrent) {
		var ret float64
		return ret
	}
	return *o.DueFromRelatedPartiesNonCurrent
}

// GetDueFromRelatedPartiesNonCurrentOk returns a tuple with the DueFromRelatedPartiesNonCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetDueFromRelatedPartiesNonCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.DueFromRelatedPartiesNonCurrent) {
		return nil, false
	}
	return o.DueFromRelatedPartiesNonCurrent, true
}

// HasDueFromRelatedPartiesNonCurrent returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasDueFromRelatedPartiesNonCurrent() bool {
	if o != nil && !IsNil(o.DueFromRelatedPartiesNonCurrent) {
		return true
	}

	return false
}

// SetDueFromRelatedPartiesNonCurrent gets a reference to the given float64 and assigns it to the DueFromRelatedPartiesNonCurrent field.
func (o *AssetsInfoNonCurrentAssets) SetDueFromRelatedPartiesNonCurrent(v float64) {
	o.DueFromRelatedPartiesNonCurrent = &v
}

// GetNonCurrentPrepaidAssets returns the NonCurrentPrepaidAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentPrepaidAssets() float64 {
	if o == nil || IsNil(o.NonCurrentPrepaidAssets) {
		var ret float64
		return ret
	}
	return *o.NonCurrentPrepaidAssets
}

// GetNonCurrentPrepaidAssetsOk returns a tuple with the NonCurrentPrepaidAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentPrepaidAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentPrepaidAssets) {
		return nil, false
	}
	return o.NonCurrentPrepaidAssets, true
}

// HasNonCurrentPrepaidAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentPrepaidAssets() bool {
	if o != nil && !IsNil(o.NonCurrentPrepaidAssets) {
		return true
	}

	return false
}

// SetNonCurrentPrepaidAssets gets a reference to the given float64 and assigns it to the NonCurrentPrepaidAssets field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentPrepaidAssets(v float64) {
	o.NonCurrentPrepaidAssets = &v
}

// GetNonCurrentDeferredAssets returns the NonCurrentDeferredAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredAssets() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredAssets) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredAssets
}

// GetNonCurrentDeferredAssetsOk returns a tuple with the NonCurrentDeferredAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredAssets) {
		return nil, false
	}
	return o.NonCurrentDeferredAssets, true
}

// HasNonCurrentDeferredAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentDeferredAssets() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredAssets) {
		return true
	}

	return false
}

// SetNonCurrentDeferredAssets gets a reference to the given float64 and assigns it to the NonCurrentDeferredAssets field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentDeferredAssets(v float64) {
	o.NonCurrentDeferredAssets = &v
}

// GetNonCurrentDeferredTaxesAssets returns the NonCurrentDeferredTaxesAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredTaxesAssets() float64 {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesAssets) {
		var ret float64
		return ret
	}
	return *o.NonCurrentDeferredTaxesAssets
}

// GetNonCurrentDeferredTaxesAssetsOk returns a tuple with the NonCurrentDeferredTaxesAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetNonCurrentDeferredTaxesAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NonCurrentDeferredTaxesAssets) {
		return nil, false
	}
	return o.NonCurrentDeferredTaxesAssets, true
}

// HasNonCurrentDeferredTaxesAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasNonCurrentDeferredTaxesAssets() bool {
	if o != nil && !IsNil(o.NonCurrentDeferredTaxesAssets) {
		return true
	}

	return false
}

// SetNonCurrentDeferredTaxesAssets gets a reference to the given float64 and assigns it to the NonCurrentDeferredTaxesAssets field.
func (o *AssetsInfoNonCurrentAssets) SetNonCurrentDeferredTaxesAssets(v float64) {
	o.NonCurrentDeferredTaxesAssets = &v
}

// GetDefinedPensionBenefit returns the DefinedPensionBenefit field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetDefinedPensionBenefit() float64 {
	if o == nil || IsNil(o.DefinedPensionBenefit) {
		var ret float64
		return ret
	}
	return *o.DefinedPensionBenefit
}

// GetDefinedPensionBenefitOk returns a tuple with the DefinedPensionBenefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetDefinedPensionBenefitOk() (*float64, bool) {
	if o == nil || IsNil(o.DefinedPensionBenefit) {
		return nil, false
	}
	return o.DefinedPensionBenefit, true
}

// HasDefinedPensionBenefit returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasDefinedPensionBenefit() bool {
	if o != nil && !IsNil(o.DefinedPensionBenefit) {
		return true
	}

	return false
}

// SetDefinedPensionBenefit gets a reference to the given float64 and assigns it to the DefinedPensionBenefit field.
func (o *AssetsInfoNonCurrentAssets) SetDefinedPensionBenefit(v float64) {
	o.DefinedPensionBenefit = &v
}

// GetOtherNonCurrentAssets returns the OtherNonCurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfoNonCurrentAssets) GetOtherNonCurrentAssets() float64 {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentAssets
}

// GetOtherNonCurrentAssetsOk returns a tuple with the OtherNonCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoNonCurrentAssets) GetOtherNonCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentAssets) {
		return nil, false
	}
	return o.OtherNonCurrentAssets, true
}

// HasOtherNonCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfoNonCurrentAssets) HasOtherNonCurrentAssets() bool {
	if o != nil && !IsNil(o.OtherNonCurrentAssets) {
		return true
	}

	return false
}

// SetOtherNonCurrentAssets gets a reference to the given float64 and assigns it to the OtherNonCurrentAssets field.
func (o *AssetsInfoNonCurrentAssets) SetOtherNonCurrentAssets(v float64) {
	o.OtherNonCurrentAssets = &v
}

func (o AssetsInfoNonCurrentAssets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoNonCurrentAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalNonCurrentAssets) {
		toSerialize["total_non_current_assets"] = o.TotalNonCurrentAssets
	}
	if !IsNil(o.FinancialAssets) {
		toSerialize["financial_assets"] = o.FinancialAssets
	}
	if !IsNil(o.InvestmentsAndAdvances) {
		toSerialize["investments_and_advances"] = o.InvestmentsAndAdvances
	}
	if !IsNil(o.OtherInvestments) {
		toSerialize["other_investments"] = o.OtherInvestments
	}
	if !IsNil(o.InvestmentInFinancialAssets) {
		toSerialize["investment_in_financial_assets"] = o.InvestmentInFinancialAssets
	}
	if !IsNil(o.HeldToMaturitySecurities) {
		toSerialize["held_to_maturity_securities"] = o.HeldToMaturitySecurities
	}
	if !IsNil(o.AvailableForSaleSecurities) {
		toSerialize["available_for_sale_securities"] = o.AvailableForSaleSecurities
	}
	if !IsNil(o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal) {
		toSerialize["financial_assets_designated_as_fair_value_through_profit_or_loss_total"] = o.FinancialAssetsDesignatedAsFairValueThroughProfitOrLossTotal
	}
	if !IsNil(o.TradingSecurities) {
		toSerialize["trading_securities"] = o.TradingSecurities
	}
	if !IsNil(o.LongTermEquityInvestment) {
		toSerialize["long_term_equity_investment"] = o.LongTermEquityInvestment
	}
	if !IsNil(o.InvestmentsInJointVenturesAtCost) {
		toSerialize["investments_in_joint_ventures_at_cost"] = o.InvestmentsInJointVenturesAtCost
	}
	if !IsNil(o.InvestmentsInOtherVenturesUnderEquityMethod) {
		toSerialize["investments_in_other_ventures_under_equity_method"] = o.InvestmentsInOtherVenturesUnderEquityMethod
	}
	if !IsNil(o.InvestmentsInAssociatesAtCost) {
		toSerialize["investments_in_associates_at_cost"] = o.InvestmentsInAssociatesAtCost
	}
	if !IsNil(o.InvestmentsInSubsidiariesAtCost) {
		toSerialize["investments_in_subsidiaries_at_cost"] = o.InvestmentsInSubsidiariesAtCost
	}
	if !IsNil(o.InvestmentProperties) {
		toSerialize["investment_properties"] = o.InvestmentProperties
	}
	if !IsNil(o.GoodwillAndOtherIntangibleAssets) {
		toSerialize["goodwill_and_other_intangible_assets"] = o.GoodwillAndOtherIntangibleAssets
	}
	if !IsNil(o.NetPpe) {
		toSerialize["net_ppe"] = o.NetPpe
	}
	if !IsNil(o.GrossPpe) {
		toSerialize["gross_ppe"] = o.GrossPpe
	}
	if !IsNil(o.AccumulatedDepreciation) {
		toSerialize["accumulated_depreciation"] = o.AccumulatedDepreciation
	}
	if !IsNil(o.Leases) {
		toSerialize["leases"] = o.Leases
	}
	if !IsNil(o.ConstructionInProgress) {
		toSerialize["construction_in_progress"] = o.ConstructionInProgress
	}
	if !IsNil(o.OtherProperties) {
		toSerialize["other_properties"] = o.OtherProperties
	}
	if !IsNil(o.MachineryFurnitureEquipment) {
		toSerialize["machinery_furniture_equipment"] = o.MachineryFurnitureEquipment
	}
	if !IsNil(o.BuildingsAndImprovements) {
		toSerialize["buildings_and_improvements"] = o.BuildingsAndImprovements
	}
	if !IsNil(o.LandAndImprovements) {
		toSerialize["land_and_improvements"] = o.LandAndImprovements
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.NonCurrentAccountsReceivable) {
		toSerialize["non_current_accounts_receivable"] = o.NonCurrentAccountsReceivable
	}
	if !IsNil(o.NonCurrentNoteReceivables) {
		toSerialize["non_current_note_receivables"] = o.NonCurrentNoteReceivables
	}
	if !IsNil(o.DueFromRelatedPartiesNonCurrent) {
		toSerialize["due_from_related_parties_non_current"] = o.DueFromRelatedPartiesNonCurrent
	}
	if !IsNil(o.NonCurrentPrepaidAssets) {
		toSerialize["non_current_prepaid_assets"] = o.NonCurrentPrepaidAssets
	}
	if !IsNil(o.NonCurrentDeferredAssets) {
		toSerialize["non_current_deferred_assets"] = o.NonCurrentDeferredAssets
	}
	if !IsNil(o.NonCurrentDeferredTaxesAssets) {
		toSerialize["non_current_deferred_taxes_assets"] = o.NonCurrentDeferredTaxesAssets
	}
	if !IsNil(o.DefinedPensionBenefit) {
		toSerialize["defined_pension_benefit"] = o.DefinedPensionBenefit
	}
	if !IsNil(o.OtherNonCurrentAssets) {
		toSerialize["other_non_current_assets"] = o.OtherNonCurrentAssets
	}
	return toSerialize, nil
}

type NullableAssetsInfoNonCurrentAssets struct {
	value *AssetsInfoNonCurrentAssets
	isSet bool
}

func (v NullableAssetsInfoNonCurrentAssets) Get() *AssetsInfoNonCurrentAssets {
	return v.value
}

func (v *NullableAssetsInfoNonCurrentAssets) Set(val *AssetsInfoNonCurrentAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoNonCurrentAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoNonCurrentAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoNonCurrentAssets(val *AssetsInfoNonCurrentAssets) *NullableAssetsInfoNonCurrentAssets {
	return &NullableAssetsInfoNonCurrentAssets{value: val, isSet: true}
}

func (v NullableAssetsInfoNonCurrentAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoNonCurrentAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
