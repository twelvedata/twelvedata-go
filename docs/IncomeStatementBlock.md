# IncomeStatementBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalDate** | **string** | Date of fiscal period ending | 
**Quarter** | Pointer to **int64** | Fiscal quarter. Visible when &#x60;&amp;period&#x3D;quarterly&#x60; | [optional] 
**Year** | Pointer to **int64** | Fiscal year | [optional] 
**Sales** | Pointer to **int64** | Refers to total reported revenue | [optional] 
**CostOfGoods** | Pointer to **int64** | Refers to cost of revenue | [optional] 
**GrossProfit** | Pointer to **int64** | Refers to net gross profit: &#x60;sales&#x60; - &#x60;cost_of_goods&#x60; | [optional] 
**OperatingExpense** | Pointer to [**IncomeStatementBlockOperatingExpense**](IncomeStatementBlockOperatingExpense.md) |  | [optional] 
**OperatingIncome** | Pointer to **int64** | Refers to net operating income: &#x60;gross_profit&#x60; - &#x60;research_and_development&#x60; - &#x60;selling_general_and_administrative&#x60; | [optional] 
**NonOperatingInterest** | Pointer to [**IncomeStatementBlockNonOperatingInterest**](IncomeStatementBlockNonOperatingInterest.md) |  | [optional] 
**OtherIncomeExpense** | Pointer to **int64** | Refers to other incomes or expenses | [optional] 
**PretaxIncome** | Pointer to **int64** | Refers to earnings before tax: &#x60;operating_income&#x60; + &#x60;net_non_operating_interest&#x60; - &#x60;other_income_expense&#x60; | [optional] 
**IncomeTax** | Pointer to **int64** | Refers to a tax provision | [optional] 
**NetIncome** | Pointer to **int64** | Refers to net income: &#x60;pretax_income&#x60; - &#x60;income_tax&#x60; | [optional] 
**EpsBasic** | Pointer to **float64** | Refers to earnings per share (EPS) | [optional] 
**EpsDiluted** | Pointer to **float64** | Refers to diluted earnings per share (EPS) | [optional] 
**BasicSharesOutstanding** | Pointer to **int64** | Refers for the shares outstanding held by all its shareholders | [optional] 
**DilutedSharesOutstanding** | Pointer to **int64** | Refers to the total number of shares a company would have if all dilutive securities were exercised and converted into shares | [optional] 
**Ebit** | Pointer to **int64** | Refers to earnings before interest and taxes (EBIT) measure | [optional] 
**Ebitda** | Pointer to **int64** | Refers to EBITDA (earnings before interest, taxes, depreciation, and amortization) measure | [optional] 
**NetIncomeContinuousOperations** | Pointer to **int64** | Refers to the after-tax earnings that a business has generated from its operational activities | [optional] 
**MinorityInterests** | Pointer to **int64** | Refers to amount of minority interests paid out | [optional] 
**PreferredStockDividends** | Pointer to **int64** | Refers to dividend that is allocated to and paid on a company&#39;s preferred shares | [optional] 

## Methods

### NewIncomeStatementBlock

`func NewIncomeStatementBlock(fiscalDate string, ) *IncomeStatementBlock`

NewIncomeStatementBlock instantiates a new IncomeStatementBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeStatementBlockWithDefaults

`func NewIncomeStatementBlockWithDefaults() *IncomeStatementBlock`

NewIncomeStatementBlockWithDefaults instantiates a new IncomeStatementBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalDate

`func (o *IncomeStatementBlock) GetFiscalDate() string`

GetFiscalDate returns the FiscalDate field if non-nil, zero value otherwise.

### GetFiscalDateOk

`func (o *IncomeStatementBlock) GetFiscalDateOk() (*string, bool)`

GetFiscalDateOk returns a tuple with the FiscalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalDate

`func (o *IncomeStatementBlock) SetFiscalDate(v string)`

SetFiscalDate sets FiscalDate field to given value.


### GetQuarter

`func (o *IncomeStatementBlock) GetQuarter() int64`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *IncomeStatementBlock) GetQuarterOk() (*int64, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *IncomeStatementBlock) SetQuarter(v int64)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *IncomeStatementBlock) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.

### GetYear

`func (o *IncomeStatementBlock) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *IncomeStatementBlock) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *IncomeStatementBlock) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *IncomeStatementBlock) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetSales

`func (o *IncomeStatementBlock) GetSales() int64`

GetSales returns the Sales field if non-nil, zero value otherwise.

### GetSalesOk

`func (o *IncomeStatementBlock) GetSalesOk() (*int64, bool)`

GetSalesOk returns a tuple with the Sales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSales

`func (o *IncomeStatementBlock) SetSales(v int64)`

SetSales sets Sales field to given value.

### HasSales

`func (o *IncomeStatementBlock) HasSales() bool`

HasSales returns a boolean if a field has been set.

### GetCostOfGoods

`func (o *IncomeStatementBlock) GetCostOfGoods() int64`

GetCostOfGoods returns the CostOfGoods field if non-nil, zero value otherwise.

### GetCostOfGoodsOk

`func (o *IncomeStatementBlock) GetCostOfGoodsOk() (*int64, bool)`

GetCostOfGoodsOk returns a tuple with the CostOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoods

`func (o *IncomeStatementBlock) SetCostOfGoods(v int64)`

SetCostOfGoods sets CostOfGoods field to given value.

### HasCostOfGoods

`func (o *IncomeStatementBlock) HasCostOfGoods() bool`

HasCostOfGoods returns a boolean if a field has been set.

### GetGrossProfit

`func (o *IncomeStatementBlock) GetGrossProfit() int64`

GetGrossProfit returns the GrossProfit field if non-nil, zero value otherwise.

### GetGrossProfitOk

`func (o *IncomeStatementBlock) GetGrossProfitOk() (*int64, bool)`

GetGrossProfitOk returns a tuple with the GrossProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossProfit

`func (o *IncomeStatementBlock) SetGrossProfit(v int64)`

SetGrossProfit sets GrossProfit field to given value.

### HasGrossProfit

`func (o *IncomeStatementBlock) HasGrossProfit() bool`

HasGrossProfit returns a boolean if a field has been set.

### GetOperatingExpense

`func (o *IncomeStatementBlock) GetOperatingExpense() IncomeStatementBlockOperatingExpense`

GetOperatingExpense returns the OperatingExpense field if non-nil, zero value otherwise.

### GetOperatingExpenseOk

`func (o *IncomeStatementBlock) GetOperatingExpenseOk() (*IncomeStatementBlockOperatingExpense, bool)`

GetOperatingExpenseOk returns a tuple with the OperatingExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpense

`func (o *IncomeStatementBlock) SetOperatingExpense(v IncomeStatementBlockOperatingExpense)`

SetOperatingExpense sets OperatingExpense field to given value.

### HasOperatingExpense

`func (o *IncomeStatementBlock) HasOperatingExpense() bool`

HasOperatingExpense returns a boolean if a field has been set.

### GetOperatingIncome

`func (o *IncomeStatementBlock) GetOperatingIncome() int64`

GetOperatingIncome returns the OperatingIncome field if non-nil, zero value otherwise.

### GetOperatingIncomeOk

`func (o *IncomeStatementBlock) GetOperatingIncomeOk() (*int64, bool)`

GetOperatingIncomeOk returns a tuple with the OperatingIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingIncome

`func (o *IncomeStatementBlock) SetOperatingIncome(v int64)`

SetOperatingIncome sets OperatingIncome field to given value.

### HasOperatingIncome

`func (o *IncomeStatementBlock) HasOperatingIncome() bool`

HasOperatingIncome returns a boolean if a field has been set.

### GetNonOperatingInterest

`func (o *IncomeStatementBlock) GetNonOperatingInterest() IncomeStatementBlockNonOperatingInterest`

GetNonOperatingInterest returns the NonOperatingInterest field if non-nil, zero value otherwise.

### GetNonOperatingInterestOk

`func (o *IncomeStatementBlock) GetNonOperatingInterestOk() (*IncomeStatementBlockNonOperatingInterest, bool)`

GetNonOperatingInterestOk returns a tuple with the NonOperatingInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonOperatingInterest

`func (o *IncomeStatementBlock) SetNonOperatingInterest(v IncomeStatementBlockNonOperatingInterest)`

SetNonOperatingInterest sets NonOperatingInterest field to given value.

### HasNonOperatingInterest

`func (o *IncomeStatementBlock) HasNonOperatingInterest() bool`

HasNonOperatingInterest returns a boolean if a field has been set.

### GetOtherIncomeExpense

`func (o *IncomeStatementBlock) GetOtherIncomeExpense() int64`

GetOtherIncomeExpense returns the OtherIncomeExpense field if non-nil, zero value otherwise.

### GetOtherIncomeExpenseOk

`func (o *IncomeStatementBlock) GetOtherIncomeExpenseOk() (*int64, bool)`

GetOtherIncomeExpenseOk returns a tuple with the OtherIncomeExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIncomeExpense

`func (o *IncomeStatementBlock) SetOtherIncomeExpense(v int64)`

SetOtherIncomeExpense sets OtherIncomeExpense field to given value.

### HasOtherIncomeExpense

`func (o *IncomeStatementBlock) HasOtherIncomeExpense() bool`

HasOtherIncomeExpense returns a boolean if a field has been set.

### GetPretaxIncome

`func (o *IncomeStatementBlock) GetPretaxIncome() int64`

GetPretaxIncome returns the PretaxIncome field if non-nil, zero value otherwise.

### GetPretaxIncomeOk

`func (o *IncomeStatementBlock) GetPretaxIncomeOk() (*int64, bool)`

GetPretaxIncomeOk returns a tuple with the PretaxIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPretaxIncome

`func (o *IncomeStatementBlock) SetPretaxIncome(v int64)`

SetPretaxIncome sets PretaxIncome field to given value.

### HasPretaxIncome

`func (o *IncomeStatementBlock) HasPretaxIncome() bool`

HasPretaxIncome returns a boolean if a field has been set.

### GetIncomeTax

`func (o *IncomeStatementBlock) GetIncomeTax() int64`

GetIncomeTax returns the IncomeTax field if non-nil, zero value otherwise.

### GetIncomeTaxOk

`func (o *IncomeStatementBlock) GetIncomeTaxOk() (*int64, bool)`

GetIncomeTaxOk returns a tuple with the IncomeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeTax

`func (o *IncomeStatementBlock) SetIncomeTax(v int64)`

SetIncomeTax sets IncomeTax field to given value.

### HasIncomeTax

`func (o *IncomeStatementBlock) HasIncomeTax() bool`

HasIncomeTax returns a boolean if a field has been set.

### GetNetIncome

`func (o *IncomeStatementBlock) GetNetIncome() int64`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *IncomeStatementBlock) GetNetIncomeOk() (*int64, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *IncomeStatementBlock) SetNetIncome(v int64)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *IncomeStatementBlock) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### GetEpsBasic

`func (o *IncomeStatementBlock) GetEpsBasic() float64`

GetEpsBasic returns the EpsBasic field if non-nil, zero value otherwise.

### GetEpsBasicOk

`func (o *IncomeStatementBlock) GetEpsBasicOk() (*float64, bool)`

GetEpsBasicOk returns a tuple with the EpsBasic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBasic

`func (o *IncomeStatementBlock) SetEpsBasic(v float64)`

SetEpsBasic sets EpsBasic field to given value.

### HasEpsBasic

`func (o *IncomeStatementBlock) HasEpsBasic() bool`

HasEpsBasic returns a boolean if a field has been set.

### GetEpsDiluted

`func (o *IncomeStatementBlock) GetEpsDiluted() float64`

GetEpsDiluted returns the EpsDiluted field if non-nil, zero value otherwise.

### GetEpsDilutedOk

`func (o *IncomeStatementBlock) GetEpsDilutedOk() (*float64, bool)`

GetEpsDilutedOk returns a tuple with the EpsDiluted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsDiluted

`func (o *IncomeStatementBlock) SetEpsDiluted(v float64)`

SetEpsDiluted sets EpsDiluted field to given value.

### HasEpsDiluted

`func (o *IncomeStatementBlock) HasEpsDiluted() bool`

HasEpsDiluted returns a boolean if a field has been set.

### GetBasicSharesOutstanding

`func (o *IncomeStatementBlock) GetBasicSharesOutstanding() int64`

GetBasicSharesOutstanding returns the BasicSharesOutstanding field if non-nil, zero value otherwise.

### GetBasicSharesOutstandingOk

`func (o *IncomeStatementBlock) GetBasicSharesOutstandingOk() (*int64, bool)`

GetBasicSharesOutstandingOk returns a tuple with the BasicSharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicSharesOutstanding

`func (o *IncomeStatementBlock) SetBasicSharesOutstanding(v int64)`

SetBasicSharesOutstanding sets BasicSharesOutstanding field to given value.

### HasBasicSharesOutstanding

`func (o *IncomeStatementBlock) HasBasicSharesOutstanding() bool`

HasBasicSharesOutstanding returns a boolean if a field has been set.

### GetDilutedSharesOutstanding

`func (o *IncomeStatementBlock) GetDilutedSharesOutstanding() int64`

GetDilutedSharesOutstanding returns the DilutedSharesOutstanding field if non-nil, zero value otherwise.

### GetDilutedSharesOutstandingOk

`func (o *IncomeStatementBlock) GetDilutedSharesOutstandingOk() (*int64, bool)`

GetDilutedSharesOutstandingOk returns a tuple with the DilutedSharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDilutedSharesOutstanding

`func (o *IncomeStatementBlock) SetDilutedSharesOutstanding(v int64)`

SetDilutedSharesOutstanding sets DilutedSharesOutstanding field to given value.

### HasDilutedSharesOutstanding

`func (o *IncomeStatementBlock) HasDilutedSharesOutstanding() bool`

HasDilutedSharesOutstanding returns a boolean if a field has been set.

### GetEbit

`func (o *IncomeStatementBlock) GetEbit() int64`

GetEbit returns the Ebit field if non-nil, zero value otherwise.

### GetEbitOk

`func (o *IncomeStatementBlock) GetEbitOk() (*int64, bool)`

GetEbitOk returns a tuple with the Ebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbit

`func (o *IncomeStatementBlock) SetEbit(v int64)`

SetEbit sets Ebit field to given value.

### HasEbit

`func (o *IncomeStatementBlock) HasEbit() bool`

HasEbit returns a boolean if a field has been set.

### GetEbitda

`func (o *IncomeStatementBlock) GetEbitda() int64`

GetEbitda returns the Ebitda field if non-nil, zero value otherwise.

### GetEbitdaOk

`func (o *IncomeStatementBlock) GetEbitdaOk() (*int64, bool)`

GetEbitdaOk returns a tuple with the Ebitda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbitda

`func (o *IncomeStatementBlock) SetEbitda(v int64)`

SetEbitda sets Ebitda field to given value.

### HasEbitda

`func (o *IncomeStatementBlock) HasEbitda() bool`

HasEbitda returns a boolean if a field has been set.

### GetNetIncomeContinuousOperations

`func (o *IncomeStatementBlock) GetNetIncomeContinuousOperations() int64`

GetNetIncomeContinuousOperations returns the NetIncomeContinuousOperations field if non-nil, zero value otherwise.

### GetNetIncomeContinuousOperationsOk

`func (o *IncomeStatementBlock) GetNetIncomeContinuousOperationsOk() (*int64, bool)`

GetNetIncomeContinuousOperationsOk returns a tuple with the NetIncomeContinuousOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncomeContinuousOperations

`func (o *IncomeStatementBlock) SetNetIncomeContinuousOperations(v int64)`

SetNetIncomeContinuousOperations sets NetIncomeContinuousOperations field to given value.

### HasNetIncomeContinuousOperations

`func (o *IncomeStatementBlock) HasNetIncomeContinuousOperations() bool`

HasNetIncomeContinuousOperations returns a boolean if a field has been set.

### GetMinorityInterests

`func (o *IncomeStatementBlock) GetMinorityInterests() int64`

GetMinorityInterests returns the MinorityInterests field if non-nil, zero value otherwise.

### GetMinorityInterestsOk

`func (o *IncomeStatementBlock) GetMinorityInterestsOk() (*int64, bool)`

GetMinorityInterestsOk returns a tuple with the MinorityInterests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorityInterests

`func (o *IncomeStatementBlock) SetMinorityInterests(v int64)`

SetMinorityInterests sets MinorityInterests field to given value.

### HasMinorityInterests

`func (o *IncomeStatementBlock) HasMinorityInterests() bool`

HasMinorityInterests returns a boolean if a field has been set.

### GetPreferredStockDividends

`func (o *IncomeStatementBlock) GetPreferredStockDividends() int64`

GetPreferredStockDividends returns the PreferredStockDividends field if non-nil, zero value otherwise.

### GetPreferredStockDividendsOk

`func (o *IncomeStatementBlock) GetPreferredStockDividendsOk() (*int64, bool)`

GetPreferredStockDividendsOk returns a tuple with the PreferredStockDividends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStockDividends

`func (o *IncomeStatementBlock) SetPreferredStockDividends(v int64)`

SetPreferredStockDividends sets PreferredStockDividends field to given value.

### HasPreferredStockDividends

`func (o *IncomeStatementBlock) HasPreferredStockDividends() bool`

HasPreferredStockDividends returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


