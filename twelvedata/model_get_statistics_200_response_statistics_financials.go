/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsFinancials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsFinancials{}

// GetStatistics200ResponseStatisticsFinancials Financial information of the company
type GetStatistics200ResponseStatisticsFinancials struct {
	// Refers to the completion of the last 12-month accounting period
	FiscalYearEnds *string `json:"fiscal_year_ends,omitempty"`
	// The most recent quarter (MRQ) refers to the fiscal quarter that most recently ended
	MostRecentQuarter *string `json:"most_recent_quarter,omitempty"`
	// The portion of a company's revenue left over after direct costs are subtracted
	GrossMargin *float64 `json:"gross_margin,omitempty"`
	// Refers to gross profit margin. Calculated by dividing net income by sales revenue
	ProfitMargin *float64 `json:"profit_margin,omitempty"`
	// Operating margin is calculated by dividing operating income by net sales
	OperatingMargin *float64 `json:"operating_margin,omitempty"`
	// Return on assets (ROA) is calculated by dividing net income by total assets
	ReturnOnAssetsTtm *float64 `json:"return_on_assets_ttm,omitempty"`
	// Return on equity (ROE) is calculated by dividing net income by average shareholders' equity
	ReturnOnEquityTtm *float64                                                     `json:"return_on_equity_ttm,omitempty"`
	IncomeStatement   *GetStatistics200ResponseStatisticsFinancialsIncomeStatement `json:"income_statement,omitempty"`
	BalanceSheet      *GetStatistics200ResponseStatisticsFinancialsBalanceSheet    `json:"balance_sheet,omitempty"`
	CashFlow          *GetStatistics200ResponseStatisticsFinancialsCashFlow        `json:"cash_flow,omitempty"`
}

// NewGetStatistics200ResponseStatisticsFinancials instantiates a new GetStatistics200ResponseStatisticsFinancials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsFinancials() *GetStatistics200ResponseStatisticsFinancials {
	this := GetStatistics200ResponseStatisticsFinancials{}
	return &this
}

// NewGetStatistics200ResponseStatisticsFinancialsWithDefaults instantiates a new GetStatistics200ResponseStatisticsFinancials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsFinancialsWithDefaults() *GetStatistics200ResponseStatisticsFinancials {
	this := GetStatistics200ResponseStatisticsFinancials{}
	return &this
}

// GetFiscalYearEnds returns the FiscalYearEnds field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetFiscalYearEnds() string {
	if o == nil || IsNil(o.FiscalYearEnds) {
		var ret string
		return ret
	}
	return *o.FiscalYearEnds
}

// GetFiscalYearEndsOk returns a tuple with the FiscalYearEnds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetFiscalYearEndsOk() (*string, bool) {
	if o == nil || IsNil(o.FiscalYearEnds) {
		return nil, false
	}
	return o.FiscalYearEnds, true
}

// HasFiscalYearEnds returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasFiscalYearEnds() bool {
	if o != nil && !IsNil(o.FiscalYearEnds) {
		return true
	}

	return false
}

// SetFiscalYearEnds gets a reference to the given string and assigns it to the FiscalYearEnds field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetFiscalYearEnds(v string) {
	o.FiscalYearEnds = &v
}

// GetMostRecentQuarter returns the MostRecentQuarter field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetMostRecentQuarter() string {
	if o == nil || IsNil(o.MostRecentQuarter) {
		var ret string
		return ret
	}
	return *o.MostRecentQuarter
}

// GetMostRecentQuarterOk returns a tuple with the MostRecentQuarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetMostRecentQuarterOk() (*string, bool) {
	if o == nil || IsNil(o.MostRecentQuarter) {
		return nil, false
	}
	return o.MostRecentQuarter, true
}

// HasMostRecentQuarter returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasMostRecentQuarter() bool {
	if o != nil && !IsNil(o.MostRecentQuarter) {
		return true
	}

	return false
}

// SetMostRecentQuarter gets a reference to the given string and assigns it to the MostRecentQuarter field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetMostRecentQuarter(v string) {
	o.MostRecentQuarter = &v
}

// GetGrossMargin returns the GrossMargin field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetGrossMargin() float64 {
	if o == nil || IsNil(o.GrossMargin) {
		var ret float64
		return ret
	}
	return *o.GrossMargin
}

// GetGrossMarginOk returns a tuple with the GrossMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetGrossMarginOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossMargin) {
		return nil, false
	}
	return o.GrossMargin, true
}

// HasGrossMargin returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasGrossMargin() bool {
	if o != nil && !IsNil(o.GrossMargin) {
		return true
	}

	return false
}

// SetGrossMargin gets a reference to the given float64 and assigns it to the GrossMargin field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetGrossMargin(v float64) {
	o.GrossMargin = &v
}

// GetProfitMargin returns the ProfitMargin field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetProfitMargin() float64 {
	if o == nil || IsNil(o.ProfitMargin) {
		var ret float64
		return ret
	}
	return *o.ProfitMargin
}

// GetProfitMarginOk returns a tuple with the ProfitMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetProfitMarginOk() (*float64, bool) {
	if o == nil || IsNil(o.ProfitMargin) {
		return nil, false
	}
	return o.ProfitMargin, true
}

// HasProfitMargin returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasProfitMargin() bool {
	if o != nil && !IsNil(o.ProfitMargin) {
		return true
	}

	return false
}

// SetProfitMargin gets a reference to the given float64 and assigns it to the ProfitMargin field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetProfitMargin(v float64) {
	o.ProfitMargin = &v
}

// GetOperatingMargin returns the OperatingMargin field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetOperatingMargin() float64 {
	if o == nil || IsNil(o.OperatingMargin) {
		var ret float64
		return ret
	}
	return *o.OperatingMargin
}

// GetOperatingMarginOk returns a tuple with the OperatingMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetOperatingMarginOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingMargin) {
		return nil, false
	}
	return o.OperatingMargin, true
}

// HasOperatingMargin returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasOperatingMargin() bool {
	if o != nil && !IsNil(o.OperatingMargin) {
		return true
	}

	return false
}

// SetOperatingMargin gets a reference to the given float64 and assigns it to the OperatingMargin field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetOperatingMargin(v float64) {
	o.OperatingMargin = &v
}

// GetReturnOnAssetsTtm returns the ReturnOnAssetsTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnAssetsTtm() float64 {
	if o == nil || IsNil(o.ReturnOnAssetsTtm) {
		var ret float64
		return ret
	}
	return *o.ReturnOnAssetsTtm
}

// GetReturnOnAssetsTtmOk returns a tuple with the ReturnOnAssetsTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnAssetsTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.ReturnOnAssetsTtm) {
		return nil, false
	}
	return o.ReturnOnAssetsTtm, true
}

// HasReturnOnAssetsTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasReturnOnAssetsTtm() bool {
	if o != nil && !IsNil(o.ReturnOnAssetsTtm) {
		return true
	}

	return false
}

// SetReturnOnAssetsTtm gets a reference to the given float64 and assigns it to the ReturnOnAssetsTtm field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetReturnOnAssetsTtm(v float64) {
	o.ReturnOnAssetsTtm = &v
}

// GetReturnOnEquityTtm returns the ReturnOnEquityTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnEquityTtm() float64 {
	if o == nil || IsNil(o.ReturnOnEquityTtm) {
		var ret float64
		return ret
	}
	return *o.ReturnOnEquityTtm
}

// GetReturnOnEquityTtmOk returns a tuple with the ReturnOnEquityTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnEquityTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.ReturnOnEquityTtm) {
		return nil, false
	}
	return o.ReturnOnEquityTtm, true
}

// HasReturnOnEquityTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasReturnOnEquityTtm() bool {
	if o != nil && !IsNil(o.ReturnOnEquityTtm) {
		return true
	}

	return false
}

// SetReturnOnEquityTtm gets a reference to the given float64 and assigns it to the ReturnOnEquityTtm field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetReturnOnEquityTtm(v float64) {
	o.ReturnOnEquityTtm = &v
}

// GetIncomeStatement returns the IncomeStatement field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetIncomeStatement() GetStatistics200ResponseStatisticsFinancialsIncomeStatement {
	if o == nil || IsNil(o.IncomeStatement) {
		var ret GetStatistics200ResponseStatisticsFinancialsIncomeStatement
		return ret
	}
	return *o.IncomeStatement
}

// GetIncomeStatementOk returns a tuple with the IncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetIncomeStatementOk() (*GetStatistics200ResponseStatisticsFinancialsIncomeStatement, bool) {
	if o == nil || IsNil(o.IncomeStatement) {
		return nil, false
	}
	return o.IncomeStatement, true
}

// HasIncomeStatement returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasIncomeStatement() bool {
	if o != nil && !IsNil(o.IncomeStatement) {
		return true
	}

	return false
}

// SetIncomeStatement gets a reference to the given GetStatistics200ResponseStatisticsFinancialsIncomeStatement and assigns it to the IncomeStatement field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetIncomeStatement(v GetStatistics200ResponseStatisticsFinancialsIncomeStatement) {
	o.IncomeStatement = &v
}

// GetBalanceSheet returns the BalanceSheet field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetBalanceSheet() GetStatistics200ResponseStatisticsFinancialsBalanceSheet {
	if o == nil || IsNil(o.BalanceSheet) {
		var ret GetStatistics200ResponseStatisticsFinancialsBalanceSheet
		return ret
	}
	return *o.BalanceSheet
}

// GetBalanceSheetOk returns a tuple with the BalanceSheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetBalanceSheetOk() (*GetStatistics200ResponseStatisticsFinancialsBalanceSheet, bool) {
	if o == nil || IsNil(o.BalanceSheet) {
		return nil, false
	}
	return o.BalanceSheet, true
}

// HasBalanceSheet returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasBalanceSheet() bool {
	if o != nil && !IsNil(o.BalanceSheet) {
		return true
	}

	return false
}

// SetBalanceSheet gets a reference to the given GetStatistics200ResponseStatisticsFinancialsBalanceSheet and assigns it to the BalanceSheet field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetBalanceSheet(v GetStatistics200ResponseStatisticsFinancialsBalanceSheet) {
	o.BalanceSheet = &v
}

// GetCashFlow returns the CashFlow field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancials) GetCashFlow() GetStatistics200ResponseStatisticsFinancialsCashFlow {
	if o == nil || IsNil(o.CashFlow) {
		var ret GetStatistics200ResponseStatisticsFinancialsCashFlow
		return ret
	}
	return *o.CashFlow
}

// GetCashFlowOk returns a tuple with the CashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) GetCashFlowOk() (*GetStatistics200ResponseStatisticsFinancialsCashFlow, bool) {
	if o == nil || IsNil(o.CashFlow) {
		return nil, false
	}
	return o.CashFlow, true
}

// HasCashFlow returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancials) HasCashFlow() bool {
	if o != nil && !IsNil(o.CashFlow) {
		return true
	}

	return false
}

// SetCashFlow gets a reference to the given GetStatistics200ResponseStatisticsFinancialsCashFlow and assigns it to the CashFlow field.
func (o *GetStatistics200ResponseStatisticsFinancials) SetCashFlow(v GetStatistics200ResponseStatisticsFinancialsCashFlow) {
	o.CashFlow = &v
}

func (o GetStatistics200ResponseStatisticsFinancials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsFinancials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiscalYearEnds) {
		toSerialize["fiscal_year_ends"] = o.FiscalYearEnds
	}
	if !IsNil(o.MostRecentQuarter) {
		toSerialize["most_recent_quarter"] = o.MostRecentQuarter
	}
	if !IsNil(o.GrossMargin) {
		toSerialize["gross_margin"] = o.GrossMargin
	}
	if !IsNil(o.ProfitMargin) {
		toSerialize["profit_margin"] = o.ProfitMargin
	}
	if !IsNil(o.OperatingMargin) {
		toSerialize["operating_margin"] = o.OperatingMargin
	}
	if !IsNil(o.ReturnOnAssetsTtm) {
		toSerialize["return_on_assets_ttm"] = o.ReturnOnAssetsTtm
	}
	if !IsNil(o.ReturnOnEquityTtm) {
		toSerialize["return_on_equity_ttm"] = o.ReturnOnEquityTtm
	}
	if !IsNil(o.IncomeStatement) {
		toSerialize["income_statement"] = o.IncomeStatement
	}
	if !IsNil(o.BalanceSheet) {
		toSerialize["balance_sheet"] = o.BalanceSheet
	}
	if !IsNil(o.CashFlow) {
		toSerialize["cash_flow"] = o.CashFlow
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsFinancials struct {
	value *GetStatistics200ResponseStatisticsFinancials
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsFinancials) Get() *GetStatistics200ResponseStatisticsFinancials {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsFinancials) Set(val *GetStatistics200ResponseStatisticsFinancials) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsFinancials) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsFinancials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsFinancials(val *GetStatistics200ResponseStatisticsFinancials) *NullableGetStatistics200ResponseStatisticsFinancials {
	return &NullableGetStatistics200ResponseStatisticsFinancials{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsFinancials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsFinancials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
