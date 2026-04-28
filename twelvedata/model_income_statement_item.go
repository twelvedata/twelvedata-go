/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IncomeStatementItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItem{}

// IncomeStatementItem IncomeStatementItem represents a single income statement record
type IncomeStatementItem struct {
	// Date of fiscal period ending
	FiscalDate string `json:"fiscal_date"`
	// Fiscal year
	Year                        int64                                           `json:"year"`
	Revenue                     *IncomeStatementItemRevenue                     `json:"revenue,omitempty"`
	GrossProfit                 *IncomeStatementItemGrossProfit                 `json:"gross_profit,omitempty"`
	OperatingIncome             *IncomeStatementItemOperatingIncome             `json:"operating_income,omitempty"`
	NetIncome                   *IncomeStatementItemNetIncome                   `json:"net_income,omitempty"`
	EarningsPerShare            *IncomeStatementItemEarningsPerShare            `json:"earnings_per_share,omitempty"`
	Expenses                    *IncomeStatementItemExpenses                    `json:"expenses,omitempty"`
	InterestIncomeAndExpense    *IncomeStatementItemInterestIncomeAndExpense    `json:"interest_income_and_expense,omitempty"`
	OtherIncomeAndExpenses      *IncomeStatementItemOtherIncomeAndExpenses      `json:"other_income_and_expenses,omitempty"`
	Taxes                       *IncomeStatementItemTaxes                       `json:"taxes,omitempty"`
	DepreciationAndAmortization *IncomeStatementItemDepreciationAndAmortization `json:"depreciation_and_amortization,omitempty"`
	Ebitda                      *IncomeStatementItemEbitda                      `json:"ebitda,omitempty"`
	DividendsAndShares          *IncomeStatementItemDividendsAndShares          `json:"dividends_and_shares,omitempty"`
	UnusualItems                *IncomeStatementItemUnusualItems                `json:"unusual_items,omitempty"`
	Depreciation                *IncomeStatementItemDepreciation                `json:"depreciation,omitempty"`
	PretaxIncome                *IncomeStatementItemPretaxIncome                `json:"pretax_income,omitempty"`
	SpecialIncomeCharges        *IncomeStatementItemSpecialIncomeCharges        `json:"special_income_charges,omitempty"`
}

type _IncomeStatementItem IncomeStatementItem

// NewIncomeStatementItem instantiates a new IncomeStatementItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItem(fiscalDate string, year int64) *IncomeStatementItem {
	this := IncomeStatementItem{}
	this.FiscalDate = fiscalDate
	this.Year = year
	return &this
}

// NewIncomeStatementItemWithDefaults instantiates a new IncomeStatementItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemWithDefaults() *IncomeStatementItem {
	this := IncomeStatementItem{}
	return &this
}

// GetFiscalDate returns the FiscalDate field value
func (o *IncomeStatementItem) GetFiscalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiscalDate
}

// GetFiscalDateOk returns a tuple with the FiscalDate field value
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetFiscalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiscalDate, true
}

// SetFiscalDate sets field value
func (o *IncomeStatementItem) SetFiscalDate(v string) {
	o.FiscalDate = v
}

// GetYear returns the Year field value
func (o *IncomeStatementItem) GetYear() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetYearOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *IncomeStatementItem) SetYear(v int64) {
	o.Year = v
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetRevenue() IncomeStatementItemRevenue {
	if o == nil || IsNil(o.Revenue) {
		var ret IncomeStatementItemRevenue
		return ret
	}
	return *o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetRevenueOk() (*IncomeStatementItemRevenue, bool) {
	if o == nil || IsNil(o.Revenue) {
		return nil, false
	}
	return o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasRevenue() bool {
	if o != nil && !IsNil(o.Revenue) {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given IncomeStatementItemRevenue and assigns it to the Revenue field.
func (o *IncomeStatementItem) SetRevenue(v IncomeStatementItemRevenue) {
	o.Revenue = &v
}

// GetGrossProfit returns the GrossProfit field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetGrossProfit() IncomeStatementItemGrossProfit {
	if o == nil || IsNil(o.GrossProfit) {
		var ret IncomeStatementItemGrossProfit
		return ret
	}
	return *o.GrossProfit
}

// GetGrossProfitOk returns a tuple with the GrossProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetGrossProfitOk() (*IncomeStatementItemGrossProfit, bool) {
	if o == nil || IsNil(o.GrossProfit) {
		return nil, false
	}
	return o.GrossProfit, true
}

// HasGrossProfit returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasGrossProfit() bool {
	if o != nil && !IsNil(o.GrossProfit) {
		return true
	}

	return false
}

// SetGrossProfit gets a reference to the given IncomeStatementItemGrossProfit and assigns it to the GrossProfit field.
func (o *IncomeStatementItem) SetGrossProfit(v IncomeStatementItemGrossProfit) {
	o.GrossProfit = &v
}

// GetOperatingIncome returns the OperatingIncome field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetOperatingIncome() IncomeStatementItemOperatingIncome {
	if o == nil || IsNil(o.OperatingIncome) {
		var ret IncomeStatementItemOperatingIncome
		return ret
	}
	return *o.OperatingIncome
}

// GetOperatingIncomeOk returns a tuple with the OperatingIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetOperatingIncomeOk() (*IncomeStatementItemOperatingIncome, bool) {
	if o == nil || IsNil(o.OperatingIncome) {
		return nil, false
	}
	return o.OperatingIncome, true
}

// HasOperatingIncome returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasOperatingIncome() bool {
	if o != nil && !IsNil(o.OperatingIncome) {
		return true
	}

	return false
}

// SetOperatingIncome gets a reference to the given IncomeStatementItemOperatingIncome and assigns it to the OperatingIncome field.
func (o *IncomeStatementItem) SetOperatingIncome(v IncomeStatementItemOperatingIncome) {
	o.OperatingIncome = &v
}

// GetNetIncome returns the NetIncome field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetNetIncome() IncomeStatementItemNetIncome {
	if o == nil || IsNil(o.NetIncome) {
		var ret IncomeStatementItemNetIncome
		return ret
	}
	return *o.NetIncome
}

// GetNetIncomeOk returns a tuple with the NetIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetNetIncomeOk() (*IncomeStatementItemNetIncome, bool) {
	if o == nil || IsNil(o.NetIncome) {
		return nil, false
	}
	return o.NetIncome, true
}

// HasNetIncome returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasNetIncome() bool {
	if o != nil && !IsNil(o.NetIncome) {
		return true
	}

	return false
}

// SetNetIncome gets a reference to the given IncomeStatementItemNetIncome and assigns it to the NetIncome field.
func (o *IncomeStatementItem) SetNetIncome(v IncomeStatementItemNetIncome) {
	o.NetIncome = &v
}

// GetEarningsPerShare returns the EarningsPerShare field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetEarningsPerShare() IncomeStatementItemEarningsPerShare {
	if o == nil || IsNil(o.EarningsPerShare) {
		var ret IncomeStatementItemEarningsPerShare
		return ret
	}
	return *o.EarningsPerShare
}

// GetEarningsPerShareOk returns a tuple with the EarningsPerShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetEarningsPerShareOk() (*IncomeStatementItemEarningsPerShare, bool) {
	if o == nil || IsNil(o.EarningsPerShare) {
		return nil, false
	}
	return o.EarningsPerShare, true
}

// HasEarningsPerShare returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasEarningsPerShare() bool {
	if o != nil && !IsNil(o.EarningsPerShare) {
		return true
	}

	return false
}

// SetEarningsPerShare gets a reference to the given IncomeStatementItemEarningsPerShare and assigns it to the EarningsPerShare field.
func (o *IncomeStatementItem) SetEarningsPerShare(v IncomeStatementItemEarningsPerShare) {
	o.EarningsPerShare = &v
}

// GetExpenses returns the Expenses field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetExpenses() IncomeStatementItemExpenses {
	if o == nil || IsNil(o.Expenses) {
		var ret IncomeStatementItemExpenses
		return ret
	}
	return *o.Expenses
}

// GetExpensesOk returns a tuple with the Expenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetExpensesOk() (*IncomeStatementItemExpenses, bool) {
	if o == nil || IsNil(o.Expenses) {
		return nil, false
	}
	return o.Expenses, true
}

// HasExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasExpenses() bool {
	if o != nil && !IsNil(o.Expenses) {
		return true
	}

	return false
}

// SetExpenses gets a reference to the given IncomeStatementItemExpenses and assigns it to the Expenses field.
func (o *IncomeStatementItem) SetExpenses(v IncomeStatementItemExpenses) {
	o.Expenses = &v
}

// GetInterestIncomeAndExpense returns the InterestIncomeAndExpense field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetInterestIncomeAndExpense() IncomeStatementItemInterestIncomeAndExpense {
	if o == nil || IsNil(o.InterestIncomeAndExpense) {
		var ret IncomeStatementItemInterestIncomeAndExpense
		return ret
	}
	return *o.InterestIncomeAndExpense
}

// GetInterestIncomeAndExpenseOk returns a tuple with the InterestIncomeAndExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetInterestIncomeAndExpenseOk() (*IncomeStatementItemInterestIncomeAndExpense, bool) {
	if o == nil || IsNil(o.InterestIncomeAndExpense) {
		return nil, false
	}
	return o.InterestIncomeAndExpense, true
}

// HasInterestIncomeAndExpense returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasInterestIncomeAndExpense() bool {
	if o != nil && !IsNil(o.InterestIncomeAndExpense) {
		return true
	}

	return false
}

// SetInterestIncomeAndExpense gets a reference to the given IncomeStatementItemInterestIncomeAndExpense and assigns it to the InterestIncomeAndExpense field.
func (o *IncomeStatementItem) SetInterestIncomeAndExpense(v IncomeStatementItemInterestIncomeAndExpense) {
	o.InterestIncomeAndExpense = &v
}

// GetOtherIncomeAndExpenses returns the OtherIncomeAndExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetOtherIncomeAndExpenses() IncomeStatementItemOtherIncomeAndExpenses {
	if o == nil || IsNil(o.OtherIncomeAndExpenses) {
		var ret IncomeStatementItemOtherIncomeAndExpenses
		return ret
	}
	return *o.OtherIncomeAndExpenses
}

// GetOtherIncomeAndExpensesOk returns a tuple with the OtherIncomeAndExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetOtherIncomeAndExpensesOk() (*IncomeStatementItemOtherIncomeAndExpenses, bool) {
	if o == nil || IsNil(o.OtherIncomeAndExpenses) {
		return nil, false
	}
	return o.OtherIncomeAndExpenses, true
}

// HasOtherIncomeAndExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasOtherIncomeAndExpenses() bool {
	if o != nil && !IsNil(o.OtherIncomeAndExpenses) {
		return true
	}

	return false
}

// SetOtherIncomeAndExpenses gets a reference to the given IncomeStatementItemOtherIncomeAndExpenses and assigns it to the OtherIncomeAndExpenses field.
func (o *IncomeStatementItem) SetOtherIncomeAndExpenses(v IncomeStatementItemOtherIncomeAndExpenses) {
	o.OtherIncomeAndExpenses = &v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetTaxes() IncomeStatementItemTaxes {
	if o == nil || IsNil(o.Taxes) {
		var ret IncomeStatementItemTaxes
		return ret
	}
	return *o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetTaxesOk() (*IncomeStatementItemTaxes, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given IncomeStatementItemTaxes and assigns it to the Taxes field.
func (o *IncomeStatementItem) SetTaxes(v IncomeStatementItemTaxes) {
	o.Taxes = &v
}

// GetDepreciationAndAmortization returns the DepreciationAndAmortization field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetDepreciationAndAmortization() IncomeStatementItemDepreciationAndAmortization {
	if o == nil || IsNil(o.DepreciationAndAmortization) {
		var ret IncomeStatementItemDepreciationAndAmortization
		return ret
	}
	return *o.DepreciationAndAmortization
}

// GetDepreciationAndAmortizationOk returns a tuple with the DepreciationAndAmortization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetDepreciationAndAmortizationOk() (*IncomeStatementItemDepreciationAndAmortization, bool) {
	if o == nil || IsNil(o.DepreciationAndAmortization) {
		return nil, false
	}
	return o.DepreciationAndAmortization, true
}

// HasDepreciationAndAmortization returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasDepreciationAndAmortization() bool {
	if o != nil && !IsNil(o.DepreciationAndAmortization) {
		return true
	}

	return false
}

// SetDepreciationAndAmortization gets a reference to the given IncomeStatementItemDepreciationAndAmortization and assigns it to the DepreciationAndAmortization field.
func (o *IncomeStatementItem) SetDepreciationAndAmortization(v IncomeStatementItemDepreciationAndAmortization) {
	o.DepreciationAndAmortization = &v
}

// GetEbitda returns the Ebitda field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetEbitda() IncomeStatementItemEbitda {
	if o == nil || IsNil(o.Ebitda) {
		var ret IncomeStatementItemEbitda
		return ret
	}
	return *o.Ebitda
}

// GetEbitdaOk returns a tuple with the Ebitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetEbitdaOk() (*IncomeStatementItemEbitda, bool) {
	if o == nil || IsNil(o.Ebitda) {
		return nil, false
	}
	return o.Ebitda, true
}

// HasEbitda returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasEbitda() bool {
	if o != nil && !IsNil(o.Ebitda) {
		return true
	}

	return false
}

// SetEbitda gets a reference to the given IncomeStatementItemEbitda and assigns it to the Ebitda field.
func (o *IncomeStatementItem) SetEbitda(v IncomeStatementItemEbitda) {
	o.Ebitda = &v
}

// GetDividendsAndShares returns the DividendsAndShares field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetDividendsAndShares() IncomeStatementItemDividendsAndShares {
	if o == nil || IsNil(o.DividendsAndShares) {
		var ret IncomeStatementItemDividendsAndShares
		return ret
	}
	return *o.DividendsAndShares
}

// GetDividendsAndSharesOk returns a tuple with the DividendsAndShares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetDividendsAndSharesOk() (*IncomeStatementItemDividendsAndShares, bool) {
	if o == nil || IsNil(o.DividendsAndShares) {
		return nil, false
	}
	return o.DividendsAndShares, true
}

// HasDividendsAndShares returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasDividendsAndShares() bool {
	if o != nil && !IsNil(o.DividendsAndShares) {
		return true
	}

	return false
}

// SetDividendsAndShares gets a reference to the given IncomeStatementItemDividendsAndShares and assigns it to the DividendsAndShares field.
func (o *IncomeStatementItem) SetDividendsAndShares(v IncomeStatementItemDividendsAndShares) {
	o.DividendsAndShares = &v
}

// GetUnusualItems returns the UnusualItems field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetUnusualItems() IncomeStatementItemUnusualItems {
	if o == nil || IsNil(o.UnusualItems) {
		var ret IncomeStatementItemUnusualItems
		return ret
	}
	return *o.UnusualItems
}

// GetUnusualItemsOk returns a tuple with the UnusualItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetUnusualItemsOk() (*IncomeStatementItemUnusualItems, bool) {
	if o == nil || IsNil(o.UnusualItems) {
		return nil, false
	}
	return o.UnusualItems, true
}

// HasUnusualItems returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasUnusualItems() bool {
	if o != nil && !IsNil(o.UnusualItems) {
		return true
	}

	return false
}

// SetUnusualItems gets a reference to the given IncomeStatementItemUnusualItems and assigns it to the UnusualItems field.
func (o *IncomeStatementItem) SetUnusualItems(v IncomeStatementItemUnusualItems) {
	o.UnusualItems = &v
}

// GetDepreciation returns the Depreciation field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetDepreciation() IncomeStatementItemDepreciation {
	if o == nil || IsNil(o.Depreciation) {
		var ret IncomeStatementItemDepreciation
		return ret
	}
	return *o.Depreciation
}

// GetDepreciationOk returns a tuple with the Depreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetDepreciationOk() (*IncomeStatementItemDepreciation, bool) {
	if o == nil || IsNil(o.Depreciation) {
		return nil, false
	}
	return o.Depreciation, true
}

// HasDepreciation returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasDepreciation() bool {
	if o != nil && !IsNil(o.Depreciation) {
		return true
	}

	return false
}

// SetDepreciation gets a reference to the given IncomeStatementItemDepreciation and assigns it to the Depreciation field.
func (o *IncomeStatementItem) SetDepreciation(v IncomeStatementItemDepreciation) {
	o.Depreciation = &v
}

// GetPretaxIncome returns the PretaxIncome field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetPretaxIncome() IncomeStatementItemPretaxIncome {
	if o == nil || IsNil(o.PretaxIncome) {
		var ret IncomeStatementItemPretaxIncome
		return ret
	}
	return *o.PretaxIncome
}

// GetPretaxIncomeOk returns a tuple with the PretaxIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetPretaxIncomeOk() (*IncomeStatementItemPretaxIncome, bool) {
	if o == nil || IsNil(o.PretaxIncome) {
		return nil, false
	}
	return o.PretaxIncome, true
}

// HasPretaxIncome returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasPretaxIncome() bool {
	if o != nil && !IsNil(o.PretaxIncome) {
		return true
	}

	return false
}

// SetPretaxIncome gets a reference to the given IncomeStatementItemPretaxIncome and assigns it to the PretaxIncome field.
func (o *IncomeStatementItem) SetPretaxIncome(v IncomeStatementItemPretaxIncome) {
	o.PretaxIncome = &v
}

// GetSpecialIncomeCharges returns the SpecialIncomeCharges field value if set, zero value otherwise.
func (o *IncomeStatementItem) GetSpecialIncomeCharges() IncomeStatementItemSpecialIncomeCharges {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		var ret IncomeStatementItemSpecialIncomeCharges
		return ret
	}
	return *o.SpecialIncomeCharges
}

// GetSpecialIncomeChargesOk returns a tuple with the SpecialIncomeCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItem) GetSpecialIncomeChargesOk() (*IncomeStatementItemSpecialIncomeCharges, bool) {
	if o == nil || IsNil(o.SpecialIncomeCharges) {
		return nil, false
	}
	return o.SpecialIncomeCharges, true
}

// HasSpecialIncomeCharges returns a boolean if a field has been set.
func (o *IncomeStatementItem) HasSpecialIncomeCharges() bool {
	if o != nil && !IsNil(o.SpecialIncomeCharges) {
		return true
	}

	return false
}

// SetSpecialIncomeCharges gets a reference to the given IncomeStatementItemSpecialIncomeCharges and assigns it to the SpecialIncomeCharges field.
func (o *IncomeStatementItem) SetSpecialIncomeCharges(v IncomeStatementItemSpecialIncomeCharges) {
	o.SpecialIncomeCharges = &v
}

func (o IncomeStatementItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fiscal_date"] = o.FiscalDate
	toSerialize["year"] = o.Year
	if !IsNil(o.Revenue) {
		toSerialize["revenue"] = o.Revenue
	}
	if !IsNil(o.GrossProfit) {
		toSerialize["gross_profit"] = o.GrossProfit
	}
	if !IsNil(o.OperatingIncome) {
		toSerialize["operating_income"] = o.OperatingIncome
	}
	if !IsNil(o.NetIncome) {
		toSerialize["net_income"] = o.NetIncome
	}
	if !IsNil(o.EarningsPerShare) {
		toSerialize["earnings_per_share"] = o.EarningsPerShare
	}
	if !IsNil(o.Expenses) {
		toSerialize["expenses"] = o.Expenses
	}
	if !IsNil(o.InterestIncomeAndExpense) {
		toSerialize["interest_income_and_expense"] = o.InterestIncomeAndExpense
	}
	if !IsNil(o.OtherIncomeAndExpenses) {
		toSerialize["other_income_and_expenses"] = o.OtherIncomeAndExpenses
	}
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	if !IsNil(o.DepreciationAndAmortization) {
		toSerialize["depreciation_and_amortization"] = o.DepreciationAndAmortization
	}
	if !IsNil(o.Ebitda) {
		toSerialize["ebitda"] = o.Ebitda
	}
	if !IsNil(o.DividendsAndShares) {
		toSerialize["dividends_and_shares"] = o.DividendsAndShares
	}
	if !IsNil(o.UnusualItems) {
		toSerialize["unusual_items"] = o.UnusualItems
	}
	if !IsNil(o.Depreciation) {
		toSerialize["depreciation"] = o.Depreciation
	}
	if !IsNil(o.PretaxIncome) {
		toSerialize["pretax_income"] = o.PretaxIncome
	}
	if !IsNil(o.SpecialIncomeCharges) {
		toSerialize["special_income_charges"] = o.SpecialIncomeCharges
	}
	return toSerialize, nil
}

func (o *IncomeStatementItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fiscal_date",
		"year",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIncomeStatementItem := _IncomeStatementItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIncomeStatementItem)

	if err != nil {
		return err
	}

	*o = IncomeStatementItem(varIncomeStatementItem)

	return err
}

type NullableIncomeStatementItem struct {
	value *IncomeStatementItem
	isSet bool
}

func (v NullableIncomeStatementItem) Get() *IncomeStatementItem {
	return v.value
}

func (v *NullableIncomeStatementItem) Set(val *IncomeStatementItem) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItem) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItem(val *IncomeStatementItem) *NullableIncomeStatementItem {
	return &NullableIncomeStatementItem{value: val, isSet: true}
}

func (v NullableIncomeStatementItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
