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

// checks if the IncomeStatementBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementBlock{}

// IncomeStatementBlock struct for IncomeStatementBlock
type IncomeStatementBlock struct {
	// Date of fiscal period ending
	FiscalDate string `json:"fiscal_date"`
	// Fiscal quarter. Visible when `&period=quarterly`
	Quarter *int64 `json:"quarter,omitempty"`
	// Fiscal year
	Year *int64 `json:"year,omitempty"`
	// Refers to total reported revenue
	Sales *int64 `json:"sales,omitempty"`
	// Refers to cost of revenue
	CostOfGoods *int64 `json:"cost_of_goods,omitempty"`
	// Refers to net gross profit: `sales` - `cost_of_goods`
	GrossProfit      *int64                                `json:"gross_profit,omitempty"`
	OperatingExpense *IncomeStatementBlockOperatingExpense `json:"operating_expense,omitempty"`
	// Refers to net operating income: `gross_profit` - `research_and_development` - `selling_general_and_administrative`
	OperatingIncome      *int64                                    `json:"operating_income,omitempty"`
	NonOperatingInterest *IncomeStatementBlockNonOperatingInterest `json:"non_operating_interest,omitempty"`
	// Refers to other incomes or expenses
	OtherIncomeExpense *int64 `json:"other_income_expense,omitempty"`
	// Refers to earnings before tax: `operating_income` + `net_non_operating_interest` - `other_income_expense`
	PretaxIncome *int64 `json:"pretax_income,omitempty"`
	// Refers to a tax provision
	IncomeTax *int64 `json:"income_tax,omitempty"`
	// Refers to net income: `pretax_income` - `income_tax`
	NetIncome *int64 `json:"net_income,omitempty"`
	// Refers to earnings per share (EPS)
	EpsBasic *float64 `json:"eps_basic,omitempty"`
	// Refers to diluted earnings per share (EPS)
	EpsDiluted *float64 `json:"eps_diluted,omitempty"`
	// Refers for the shares outstanding held by all its shareholders
	BasicSharesOutstanding *int64 `json:"basic_shares_outstanding,omitempty"`
	// Refers to the total number of shares a company would have if all dilutive securities were exercised and converted into shares
	DilutedSharesOutstanding *int64 `json:"diluted_shares_outstanding,omitempty"`
	// Refers to earnings before interest and taxes (EBIT) measure
	Ebit *int64 `json:"ebit,omitempty"`
	// Refers to EBITDA (earnings before interest, taxes, depreciation, and amortization) measure
	Ebitda *int64 `json:"ebitda,omitempty"`
	// Refers to the after-tax earnings that a business has generated from its operational activities
	NetIncomeContinuousOperations *int64 `json:"net_income_continuous_operations,omitempty"`
	// Refers to amount of minority interests paid out
	MinorityInterests *int64 `json:"minority_interests,omitempty"`
	// Refers to dividend that is allocated to and paid on a company's preferred shares
	PreferredStockDividends *int64 `json:"preferred_stock_dividends,omitempty"`
}

type _IncomeStatementBlock IncomeStatementBlock

// NewIncomeStatementBlock instantiates a new IncomeStatementBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementBlock(fiscalDate string) *IncomeStatementBlock {
	this := IncomeStatementBlock{}
	this.FiscalDate = fiscalDate
	return &this
}

// NewIncomeStatementBlockWithDefaults instantiates a new IncomeStatementBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementBlockWithDefaults() *IncomeStatementBlock {
	this := IncomeStatementBlock{}
	return &this
}

// GetFiscalDate returns the FiscalDate field value
func (o *IncomeStatementBlock) GetFiscalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiscalDate
}

// GetFiscalDateOk returns a tuple with the FiscalDate field value
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetFiscalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiscalDate, true
}

// SetFiscalDate sets field value
func (o *IncomeStatementBlock) SetFiscalDate(v string) {
	o.FiscalDate = v
}

// GetQuarter returns the Quarter field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetQuarter() int64 {
	if o == nil || IsNil(o.Quarter) {
		var ret int64
		return ret
	}
	return *o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetQuarterOk() (*int64, bool) {
	if o == nil || IsNil(o.Quarter) {
		return nil, false
	}
	return o.Quarter, true
}

// HasQuarter returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasQuarter() bool {
	if o != nil && !IsNil(o.Quarter) {
		return true
	}

	return false
}

// SetQuarter gets a reference to the given int64 and assigns it to the Quarter field.
func (o *IncomeStatementBlock) SetQuarter(v int64) {
	o.Quarter = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *IncomeStatementBlock) SetYear(v int64) {
	o.Year = &v
}

// GetSales returns the Sales field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetSales() int64 {
	if o == nil || IsNil(o.Sales) {
		var ret int64
		return ret
	}
	return *o.Sales
}

// GetSalesOk returns a tuple with the Sales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetSalesOk() (*int64, bool) {
	if o == nil || IsNil(o.Sales) {
		return nil, false
	}
	return o.Sales, true
}

// HasSales returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasSales() bool {
	if o != nil && !IsNil(o.Sales) {
		return true
	}

	return false
}

// SetSales gets a reference to the given int64 and assigns it to the Sales field.
func (o *IncomeStatementBlock) SetSales(v int64) {
	o.Sales = &v
}

// GetCostOfGoods returns the CostOfGoods field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetCostOfGoods() int64 {
	if o == nil || IsNil(o.CostOfGoods) {
		var ret int64
		return ret
	}
	return *o.CostOfGoods
}

// GetCostOfGoodsOk returns a tuple with the CostOfGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetCostOfGoodsOk() (*int64, bool) {
	if o == nil || IsNil(o.CostOfGoods) {
		return nil, false
	}
	return o.CostOfGoods, true
}

// HasCostOfGoods returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasCostOfGoods() bool {
	if o != nil && !IsNil(o.CostOfGoods) {
		return true
	}

	return false
}

// SetCostOfGoods gets a reference to the given int64 and assigns it to the CostOfGoods field.
func (o *IncomeStatementBlock) SetCostOfGoods(v int64) {
	o.CostOfGoods = &v
}

// GetGrossProfit returns the GrossProfit field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetGrossProfit() int64 {
	if o == nil || IsNil(o.GrossProfit) {
		var ret int64
		return ret
	}
	return *o.GrossProfit
}

// GetGrossProfitOk returns a tuple with the GrossProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetGrossProfitOk() (*int64, bool) {
	if o == nil || IsNil(o.GrossProfit) {
		return nil, false
	}
	return o.GrossProfit, true
}

// HasGrossProfit returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasGrossProfit() bool {
	if o != nil && !IsNil(o.GrossProfit) {
		return true
	}

	return false
}

// SetGrossProfit gets a reference to the given int64 and assigns it to the GrossProfit field.
func (o *IncomeStatementBlock) SetGrossProfit(v int64) {
	o.GrossProfit = &v
}

// GetOperatingExpense returns the OperatingExpense field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetOperatingExpense() IncomeStatementBlockOperatingExpense {
	if o == nil || IsNil(o.OperatingExpense) {
		var ret IncomeStatementBlockOperatingExpense
		return ret
	}
	return *o.OperatingExpense
}

// GetOperatingExpenseOk returns a tuple with the OperatingExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetOperatingExpenseOk() (*IncomeStatementBlockOperatingExpense, bool) {
	if o == nil || IsNil(o.OperatingExpense) {
		return nil, false
	}
	return o.OperatingExpense, true
}

// HasOperatingExpense returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasOperatingExpense() bool {
	if o != nil && !IsNil(o.OperatingExpense) {
		return true
	}

	return false
}

// SetOperatingExpense gets a reference to the given IncomeStatementBlockOperatingExpense and assigns it to the OperatingExpense field.
func (o *IncomeStatementBlock) SetOperatingExpense(v IncomeStatementBlockOperatingExpense) {
	o.OperatingExpense = &v
}

// GetOperatingIncome returns the OperatingIncome field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetOperatingIncome() int64 {
	if o == nil || IsNil(o.OperatingIncome) {
		var ret int64
		return ret
	}
	return *o.OperatingIncome
}

// GetOperatingIncomeOk returns a tuple with the OperatingIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetOperatingIncomeOk() (*int64, bool) {
	if o == nil || IsNil(o.OperatingIncome) {
		return nil, false
	}
	return o.OperatingIncome, true
}

// HasOperatingIncome returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasOperatingIncome() bool {
	if o != nil && !IsNil(o.OperatingIncome) {
		return true
	}

	return false
}

// SetOperatingIncome gets a reference to the given int64 and assigns it to the OperatingIncome field.
func (o *IncomeStatementBlock) SetOperatingIncome(v int64) {
	o.OperatingIncome = &v
}

// GetNonOperatingInterest returns the NonOperatingInterest field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetNonOperatingInterest() IncomeStatementBlockNonOperatingInterest {
	if o == nil || IsNil(o.NonOperatingInterest) {
		var ret IncomeStatementBlockNonOperatingInterest
		return ret
	}
	return *o.NonOperatingInterest
}

// GetNonOperatingInterestOk returns a tuple with the NonOperatingInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetNonOperatingInterestOk() (*IncomeStatementBlockNonOperatingInterest, bool) {
	if o == nil || IsNil(o.NonOperatingInterest) {
		return nil, false
	}
	return o.NonOperatingInterest, true
}

// HasNonOperatingInterest returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasNonOperatingInterest() bool {
	if o != nil && !IsNil(o.NonOperatingInterest) {
		return true
	}

	return false
}

// SetNonOperatingInterest gets a reference to the given IncomeStatementBlockNonOperatingInterest and assigns it to the NonOperatingInterest field.
func (o *IncomeStatementBlock) SetNonOperatingInterest(v IncomeStatementBlockNonOperatingInterest) {
	o.NonOperatingInterest = &v
}

// GetOtherIncomeExpense returns the OtherIncomeExpense field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetOtherIncomeExpense() int64 {
	if o == nil || IsNil(o.OtherIncomeExpense) {
		var ret int64
		return ret
	}
	return *o.OtherIncomeExpense
}

// GetOtherIncomeExpenseOk returns a tuple with the OtherIncomeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetOtherIncomeExpenseOk() (*int64, bool) {
	if o == nil || IsNil(o.OtherIncomeExpense) {
		return nil, false
	}
	return o.OtherIncomeExpense, true
}

// HasOtherIncomeExpense returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasOtherIncomeExpense() bool {
	if o != nil && !IsNil(o.OtherIncomeExpense) {
		return true
	}

	return false
}

// SetOtherIncomeExpense gets a reference to the given int64 and assigns it to the OtherIncomeExpense field.
func (o *IncomeStatementBlock) SetOtherIncomeExpense(v int64) {
	o.OtherIncomeExpense = &v
}

// GetPretaxIncome returns the PretaxIncome field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetPretaxIncome() int64 {
	if o == nil || IsNil(o.PretaxIncome) {
		var ret int64
		return ret
	}
	return *o.PretaxIncome
}

// GetPretaxIncomeOk returns a tuple with the PretaxIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetPretaxIncomeOk() (*int64, bool) {
	if o == nil || IsNil(o.PretaxIncome) {
		return nil, false
	}
	return o.PretaxIncome, true
}

// HasPretaxIncome returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasPretaxIncome() bool {
	if o != nil && !IsNil(o.PretaxIncome) {
		return true
	}

	return false
}

// SetPretaxIncome gets a reference to the given int64 and assigns it to the PretaxIncome field.
func (o *IncomeStatementBlock) SetPretaxIncome(v int64) {
	o.PretaxIncome = &v
}

// GetIncomeTax returns the IncomeTax field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetIncomeTax() int64 {
	if o == nil || IsNil(o.IncomeTax) {
		var ret int64
		return ret
	}
	return *o.IncomeTax
}

// GetIncomeTaxOk returns a tuple with the IncomeTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetIncomeTaxOk() (*int64, bool) {
	if o == nil || IsNil(o.IncomeTax) {
		return nil, false
	}
	return o.IncomeTax, true
}

// HasIncomeTax returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasIncomeTax() bool {
	if o != nil && !IsNil(o.IncomeTax) {
		return true
	}

	return false
}

// SetIncomeTax gets a reference to the given int64 and assigns it to the IncomeTax field.
func (o *IncomeStatementBlock) SetIncomeTax(v int64) {
	o.IncomeTax = &v
}

// GetNetIncome returns the NetIncome field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetNetIncome() int64 {
	if o == nil || IsNil(o.NetIncome) {
		var ret int64
		return ret
	}
	return *o.NetIncome
}

// GetNetIncomeOk returns a tuple with the NetIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetNetIncomeOk() (*int64, bool) {
	if o == nil || IsNil(o.NetIncome) {
		return nil, false
	}
	return o.NetIncome, true
}

// HasNetIncome returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasNetIncome() bool {
	if o != nil && !IsNil(o.NetIncome) {
		return true
	}

	return false
}

// SetNetIncome gets a reference to the given int64 and assigns it to the NetIncome field.
func (o *IncomeStatementBlock) SetNetIncome(v int64) {
	o.NetIncome = &v
}

// GetEpsBasic returns the EpsBasic field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetEpsBasic() float64 {
	if o == nil || IsNil(o.EpsBasic) {
		var ret float64
		return ret
	}
	return *o.EpsBasic
}

// GetEpsBasicOk returns a tuple with the EpsBasic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetEpsBasicOk() (*float64, bool) {
	if o == nil || IsNil(o.EpsBasic) {
		return nil, false
	}
	return o.EpsBasic, true
}

// HasEpsBasic returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasEpsBasic() bool {
	if o != nil && !IsNil(o.EpsBasic) {
		return true
	}

	return false
}

// SetEpsBasic gets a reference to the given float64 and assigns it to the EpsBasic field.
func (o *IncomeStatementBlock) SetEpsBasic(v float64) {
	o.EpsBasic = &v
}

// GetEpsDiluted returns the EpsDiluted field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetEpsDiluted() float64 {
	if o == nil || IsNil(o.EpsDiluted) {
		var ret float64
		return ret
	}
	return *o.EpsDiluted
}

// GetEpsDilutedOk returns a tuple with the EpsDiluted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetEpsDilutedOk() (*float64, bool) {
	if o == nil || IsNil(o.EpsDiluted) {
		return nil, false
	}
	return o.EpsDiluted, true
}

// HasEpsDiluted returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasEpsDiluted() bool {
	if o != nil && !IsNil(o.EpsDiluted) {
		return true
	}

	return false
}

// SetEpsDiluted gets a reference to the given float64 and assigns it to the EpsDiluted field.
func (o *IncomeStatementBlock) SetEpsDiluted(v float64) {
	o.EpsDiluted = &v
}

// GetBasicSharesOutstanding returns the BasicSharesOutstanding field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetBasicSharesOutstanding() int64 {
	if o == nil || IsNil(o.BasicSharesOutstanding) {
		var ret int64
		return ret
	}
	return *o.BasicSharesOutstanding
}

// GetBasicSharesOutstandingOk returns a tuple with the BasicSharesOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetBasicSharesOutstandingOk() (*int64, bool) {
	if o == nil || IsNil(o.BasicSharesOutstanding) {
		return nil, false
	}
	return o.BasicSharesOutstanding, true
}

// HasBasicSharesOutstanding returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasBasicSharesOutstanding() bool {
	if o != nil && !IsNil(o.BasicSharesOutstanding) {
		return true
	}

	return false
}

// SetBasicSharesOutstanding gets a reference to the given int64 and assigns it to the BasicSharesOutstanding field.
func (o *IncomeStatementBlock) SetBasicSharesOutstanding(v int64) {
	o.BasicSharesOutstanding = &v
}

// GetDilutedSharesOutstanding returns the DilutedSharesOutstanding field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetDilutedSharesOutstanding() int64 {
	if o == nil || IsNil(o.DilutedSharesOutstanding) {
		var ret int64
		return ret
	}
	return *o.DilutedSharesOutstanding
}

// GetDilutedSharesOutstandingOk returns a tuple with the DilutedSharesOutstanding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetDilutedSharesOutstandingOk() (*int64, bool) {
	if o == nil || IsNil(o.DilutedSharesOutstanding) {
		return nil, false
	}
	return o.DilutedSharesOutstanding, true
}

// HasDilutedSharesOutstanding returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasDilutedSharesOutstanding() bool {
	if o != nil && !IsNil(o.DilutedSharesOutstanding) {
		return true
	}

	return false
}

// SetDilutedSharesOutstanding gets a reference to the given int64 and assigns it to the DilutedSharesOutstanding field.
func (o *IncomeStatementBlock) SetDilutedSharesOutstanding(v int64) {
	o.DilutedSharesOutstanding = &v
}

// GetEbit returns the Ebit field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetEbit() int64 {
	if o == nil || IsNil(o.Ebit) {
		var ret int64
		return ret
	}
	return *o.Ebit
}

// GetEbitOk returns a tuple with the Ebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetEbitOk() (*int64, bool) {
	if o == nil || IsNil(o.Ebit) {
		return nil, false
	}
	return o.Ebit, true
}

// HasEbit returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasEbit() bool {
	if o != nil && !IsNil(o.Ebit) {
		return true
	}

	return false
}

// SetEbit gets a reference to the given int64 and assigns it to the Ebit field.
func (o *IncomeStatementBlock) SetEbit(v int64) {
	o.Ebit = &v
}

// GetEbitda returns the Ebitda field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetEbitda() int64 {
	if o == nil || IsNil(o.Ebitda) {
		var ret int64
		return ret
	}
	return *o.Ebitda
}

// GetEbitdaOk returns a tuple with the Ebitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetEbitdaOk() (*int64, bool) {
	if o == nil || IsNil(o.Ebitda) {
		return nil, false
	}
	return o.Ebitda, true
}

// HasEbitda returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasEbitda() bool {
	if o != nil && !IsNil(o.Ebitda) {
		return true
	}

	return false
}

// SetEbitda gets a reference to the given int64 and assigns it to the Ebitda field.
func (o *IncomeStatementBlock) SetEbitda(v int64) {
	o.Ebitda = &v
}

// GetNetIncomeContinuousOperations returns the NetIncomeContinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetNetIncomeContinuousOperations() int64 {
	if o == nil || IsNil(o.NetIncomeContinuousOperations) {
		var ret int64
		return ret
	}
	return *o.NetIncomeContinuousOperations
}

// GetNetIncomeContinuousOperationsOk returns a tuple with the NetIncomeContinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetNetIncomeContinuousOperationsOk() (*int64, bool) {
	if o == nil || IsNil(o.NetIncomeContinuousOperations) {
		return nil, false
	}
	return o.NetIncomeContinuousOperations, true
}

// HasNetIncomeContinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasNetIncomeContinuousOperations() bool {
	if o != nil && !IsNil(o.NetIncomeContinuousOperations) {
		return true
	}

	return false
}

// SetNetIncomeContinuousOperations gets a reference to the given int64 and assigns it to the NetIncomeContinuousOperations field.
func (o *IncomeStatementBlock) SetNetIncomeContinuousOperations(v int64) {
	o.NetIncomeContinuousOperations = &v
}

// GetMinorityInterests returns the MinorityInterests field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetMinorityInterests() int64 {
	if o == nil || IsNil(o.MinorityInterests) {
		var ret int64
		return ret
	}
	return *o.MinorityInterests
}

// GetMinorityInterestsOk returns a tuple with the MinorityInterests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetMinorityInterestsOk() (*int64, bool) {
	if o == nil || IsNil(o.MinorityInterests) {
		return nil, false
	}
	return o.MinorityInterests, true
}

// HasMinorityInterests returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasMinorityInterests() bool {
	if o != nil && !IsNil(o.MinorityInterests) {
		return true
	}

	return false
}

// SetMinorityInterests gets a reference to the given int64 and assigns it to the MinorityInterests field.
func (o *IncomeStatementBlock) SetMinorityInterests(v int64) {
	o.MinorityInterests = &v
}

// GetPreferredStockDividends returns the PreferredStockDividends field value if set, zero value otherwise.
func (o *IncomeStatementBlock) GetPreferredStockDividends() int64 {
	if o == nil || IsNil(o.PreferredStockDividends) {
		var ret int64
		return ret
	}
	return *o.PreferredStockDividends
}

// GetPreferredStockDividendsOk returns a tuple with the PreferredStockDividends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlock) GetPreferredStockDividendsOk() (*int64, bool) {
	if o == nil || IsNil(o.PreferredStockDividends) {
		return nil, false
	}
	return o.PreferredStockDividends, true
}

// HasPreferredStockDividends returns a boolean if a field has been set.
func (o *IncomeStatementBlock) HasPreferredStockDividends() bool {
	if o != nil && !IsNil(o.PreferredStockDividends) {
		return true
	}

	return false
}

// SetPreferredStockDividends gets a reference to the given int64 and assigns it to the PreferredStockDividends field.
func (o *IncomeStatementBlock) SetPreferredStockDividends(v int64) {
	o.PreferredStockDividends = &v
}

func (o IncomeStatementBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fiscal_date"] = o.FiscalDate
	if !IsNil(o.Quarter) {
		toSerialize["quarter"] = o.Quarter
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Sales) {
		toSerialize["sales"] = o.Sales
	}
	if !IsNil(o.CostOfGoods) {
		toSerialize["cost_of_goods"] = o.CostOfGoods
	}
	if !IsNil(o.GrossProfit) {
		toSerialize["gross_profit"] = o.GrossProfit
	}
	if !IsNil(o.OperatingExpense) {
		toSerialize["operating_expense"] = o.OperatingExpense
	}
	if !IsNil(o.OperatingIncome) {
		toSerialize["operating_income"] = o.OperatingIncome
	}
	if !IsNil(o.NonOperatingInterest) {
		toSerialize["non_operating_interest"] = o.NonOperatingInterest
	}
	if !IsNil(o.OtherIncomeExpense) {
		toSerialize["other_income_expense"] = o.OtherIncomeExpense
	}
	if !IsNil(o.PretaxIncome) {
		toSerialize["pretax_income"] = o.PretaxIncome
	}
	if !IsNil(o.IncomeTax) {
		toSerialize["income_tax"] = o.IncomeTax
	}
	if !IsNil(o.NetIncome) {
		toSerialize["net_income"] = o.NetIncome
	}
	if !IsNil(o.EpsBasic) {
		toSerialize["eps_basic"] = o.EpsBasic
	}
	if !IsNil(o.EpsDiluted) {
		toSerialize["eps_diluted"] = o.EpsDiluted
	}
	if !IsNil(o.BasicSharesOutstanding) {
		toSerialize["basic_shares_outstanding"] = o.BasicSharesOutstanding
	}
	if !IsNil(o.DilutedSharesOutstanding) {
		toSerialize["diluted_shares_outstanding"] = o.DilutedSharesOutstanding
	}
	if !IsNil(o.Ebit) {
		toSerialize["ebit"] = o.Ebit
	}
	if !IsNil(o.Ebitda) {
		toSerialize["ebitda"] = o.Ebitda
	}
	if !IsNil(o.NetIncomeContinuousOperations) {
		toSerialize["net_income_continuous_operations"] = o.NetIncomeContinuousOperations
	}
	if !IsNil(o.MinorityInterests) {
		toSerialize["minority_interests"] = o.MinorityInterests
	}
	if !IsNil(o.PreferredStockDividends) {
		toSerialize["preferred_stock_dividends"] = o.PreferredStockDividends
	}
	return toSerialize, nil
}

func (o *IncomeStatementBlock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fiscal_date",
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

	varIncomeStatementBlock := _IncomeStatementBlock{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIncomeStatementBlock)

	if err != nil {
		return err
	}

	*o = IncomeStatementBlock(varIncomeStatementBlock)

	return err
}

type NullableIncomeStatementBlock struct {
	value *IncomeStatementBlock
	isSet bool
}

func (v NullableIncomeStatementBlock) Get() *IncomeStatementBlock {
	return v.value
}

func (v *NullableIncomeStatementBlock) Set(val *IncomeStatementBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementBlock(val *IncomeStatementBlock) *NullableIncomeStatementBlock {
	return &NullableIncomeStatementBlock{value: val, isSet: true}
}

func (v NullableIncomeStatementBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
