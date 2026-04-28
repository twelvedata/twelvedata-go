# IncomeStatementItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalDate** | **string** | Date of fiscal period ending | 
**Year** | **int64** | Fiscal year | 
**Revenue** | Pointer to [**IncomeStatementItemRevenue**](IncomeStatementItemRevenue.md) |  | [optional] 
**GrossProfit** | Pointer to [**IncomeStatementItemGrossProfit**](IncomeStatementItemGrossProfit.md) |  | [optional] 
**OperatingIncome** | Pointer to [**IncomeStatementItemOperatingIncome**](IncomeStatementItemOperatingIncome.md) |  | [optional] 
**NetIncome** | Pointer to [**IncomeStatementItemNetIncome**](IncomeStatementItemNetIncome.md) |  | [optional] 
**EarningsPerShare** | Pointer to [**IncomeStatementItemEarningsPerShare**](IncomeStatementItemEarningsPerShare.md) |  | [optional] 
**Expenses** | Pointer to [**IncomeStatementItemExpenses**](IncomeStatementItemExpenses.md) |  | [optional] 
**InterestIncomeAndExpense** | Pointer to [**IncomeStatementItemInterestIncomeAndExpense**](IncomeStatementItemInterestIncomeAndExpense.md) |  | [optional] 
**OtherIncomeAndExpenses** | Pointer to [**IncomeStatementItemOtherIncomeAndExpenses**](IncomeStatementItemOtherIncomeAndExpenses.md) |  | [optional] 
**Taxes** | Pointer to [**IncomeStatementItemTaxes**](IncomeStatementItemTaxes.md) |  | [optional] 
**DepreciationAndAmortization** | Pointer to [**IncomeStatementItemDepreciationAndAmortization**](IncomeStatementItemDepreciationAndAmortization.md) |  | [optional] 
**Ebitda** | Pointer to [**IncomeStatementItemEbitda**](IncomeStatementItemEbitda.md) |  | [optional] 
**DividendsAndShares** | Pointer to [**IncomeStatementItemDividendsAndShares**](IncomeStatementItemDividendsAndShares.md) |  | [optional] 
**UnusualItems** | Pointer to [**IncomeStatementItemUnusualItems**](IncomeStatementItemUnusualItems.md) |  | [optional] 
**Depreciation** | Pointer to [**IncomeStatementItemDepreciation**](IncomeStatementItemDepreciation.md) |  | [optional] 
**PretaxIncome** | Pointer to [**IncomeStatementItemPretaxIncome**](IncomeStatementItemPretaxIncome.md) |  | [optional] 
**SpecialIncomeCharges** | Pointer to [**IncomeStatementItemSpecialIncomeCharges**](IncomeStatementItemSpecialIncomeCharges.md) |  | [optional] 

## Methods

### NewIncomeStatementItem

`func NewIncomeStatementItem(fiscalDate string, year int64, ) *IncomeStatementItem`

NewIncomeStatementItem instantiates a new IncomeStatementItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeStatementItemWithDefaults

`func NewIncomeStatementItemWithDefaults() *IncomeStatementItem`

NewIncomeStatementItemWithDefaults instantiates a new IncomeStatementItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalDate

`func (o *IncomeStatementItem) GetFiscalDate() string`

GetFiscalDate returns the FiscalDate field if non-nil, zero value otherwise.

### GetFiscalDateOk

`func (o *IncomeStatementItem) GetFiscalDateOk() (*string, bool)`

GetFiscalDateOk returns a tuple with the FiscalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalDate

`func (o *IncomeStatementItem) SetFiscalDate(v string)`

SetFiscalDate sets FiscalDate field to given value.


### GetYear

`func (o *IncomeStatementItem) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *IncomeStatementItem) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *IncomeStatementItem) SetYear(v int64)`

SetYear sets Year field to given value.


### GetRevenue

`func (o *IncomeStatementItem) GetRevenue() IncomeStatementItemRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *IncomeStatementItem) GetRevenueOk() (*IncomeStatementItemRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *IncomeStatementItem) SetRevenue(v IncomeStatementItemRevenue)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *IncomeStatementItem) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetGrossProfit

`func (o *IncomeStatementItem) GetGrossProfit() IncomeStatementItemGrossProfit`

GetGrossProfit returns the GrossProfit field if non-nil, zero value otherwise.

### GetGrossProfitOk

`func (o *IncomeStatementItem) GetGrossProfitOk() (*IncomeStatementItemGrossProfit, bool)`

GetGrossProfitOk returns a tuple with the GrossProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossProfit

`func (o *IncomeStatementItem) SetGrossProfit(v IncomeStatementItemGrossProfit)`

SetGrossProfit sets GrossProfit field to given value.

### HasGrossProfit

`func (o *IncomeStatementItem) HasGrossProfit() bool`

HasGrossProfit returns a boolean if a field has been set.

### GetOperatingIncome

`func (o *IncomeStatementItem) GetOperatingIncome() IncomeStatementItemOperatingIncome`

GetOperatingIncome returns the OperatingIncome field if non-nil, zero value otherwise.

### GetOperatingIncomeOk

`func (o *IncomeStatementItem) GetOperatingIncomeOk() (*IncomeStatementItemOperatingIncome, bool)`

GetOperatingIncomeOk returns a tuple with the OperatingIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingIncome

`func (o *IncomeStatementItem) SetOperatingIncome(v IncomeStatementItemOperatingIncome)`

SetOperatingIncome sets OperatingIncome field to given value.

### HasOperatingIncome

`func (o *IncomeStatementItem) HasOperatingIncome() bool`

HasOperatingIncome returns a boolean if a field has been set.

### GetNetIncome

`func (o *IncomeStatementItem) GetNetIncome() IncomeStatementItemNetIncome`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *IncomeStatementItem) GetNetIncomeOk() (*IncomeStatementItemNetIncome, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *IncomeStatementItem) SetNetIncome(v IncomeStatementItemNetIncome)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *IncomeStatementItem) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### GetEarningsPerShare

`func (o *IncomeStatementItem) GetEarningsPerShare() IncomeStatementItemEarningsPerShare`

GetEarningsPerShare returns the EarningsPerShare field if non-nil, zero value otherwise.

### GetEarningsPerShareOk

`func (o *IncomeStatementItem) GetEarningsPerShareOk() (*IncomeStatementItemEarningsPerShare, bool)`

GetEarningsPerShareOk returns a tuple with the EarningsPerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningsPerShare

`func (o *IncomeStatementItem) SetEarningsPerShare(v IncomeStatementItemEarningsPerShare)`

SetEarningsPerShare sets EarningsPerShare field to given value.

### HasEarningsPerShare

`func (o *IncomeStatementItem) HasEarningsPerShare() bool`

HasEarningsPerShare returns a boolean if a field has been set.

### GetExpenses

`func (o *IncomeStatementItem) GetExpenses() IncomeStatementItemExpenses`

GetExpenses returns the Expenses field if non-nil, zero value otherwise.

### GetExpensesOk

`func (o *IncomeStatementItem) GetExpensesOk() (*IncomeStatementItemExpenses, bool)`

GetExpensesOk returns a tuple with the Expenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenses

`func (o *IncomeStatementItem) SetExpenses(v IncomeStatementItemExpenses)`

SetExpenses sets Expenses field to given value.

### HasExpenses

`func (o *IncomeStatementItem) HasExpenses() bool`

HasExpenses returns a boolean if a field has been set.

### GetInterestIncomeAndExpense

`func (o *IncomeStatementItem) GetInterestIncomeAndExpense() IncomeStatementItemInterestIncomeAndExpense`

GetInterestIncomeAndExpense returns the InterestIncomeAndExpense field if non-nil, zero value otherwise.

### GetInterestIncomeAndExpenseOk

`func (o *IncomeStatementItem) GetInterestIncomeAndExpenseOk() (*IncomeStatementItemInterestIncomeAndExpense, bool)`

GetInterestIncomeAndExpenseOk returns a tuple with the InterestIncomeAndExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestIncomeAndExpense

`func (o *IncomeStatementItem) SetInterestIncomeAndExpense(v IncomeStatementItemInterestIncomeAndExpense)`

SetInterestIncomeAndExpense sets InterestIncomeAndExpense field to given value.

### HasInterestIncomeAndExpense

`func (o *IncomeStatementItem) HasInterestIncomeAndExpense() bool`

HasInterestIncomeAndExpense returns a boolean if a field has been set.

### GetOtherIncomeAndExpenses

`func (o *IncomeStatementItem) GetOtherIncomeAndExpenses() IncomeStatementItemOtherIncomeAndExpenses`

GetOtherIncomeAndExpenses returns the OtherIncomeAndExpenses field if non-nil, zero value otherwise.

### GetOtherIncomeAndExpensesOk

`func (o *IncomeStatementItem) GetOtherIncomeAndExpensesOk() (*IncomeStatementItemOtherIncomeAndExpenses, bool)`

GetOtherIncomeAndExpensesOk returns a tuple with the OtherIncomeAndExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIncomeAndExpenses

`func (o *IncomeStatementItem) SetOtherIncomeAndExpenses(v IncomeStatementItemOtherIncomeAndExpenses)`

SetOtherIncomeAndExpenses sets OtherIncomeAndExpenses field to given value.

### HasOtherIncomeAndExpenses

`func (o *IncomeStatementItem) HasOtherIncomeAndExpenses() bool`

HasOtherIncomeAndExpenses returns a boolean if a field has been set.

### GetTaxes

`func (o *IncomeStatementItem) GetTaxes() IncomeStatementItemTaxes`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *IncomeStatementItem) GetTaxesOk() (*IncomeStatementItemTaxes, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *IncomeStatementItem) SetTaxes(v IncomeStatementItemTaxes)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *IncomeStatementItem) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetDepreciationAndAmortization

`func (o *IncomeStatementItem) GetDepreciationAndAmortization() IncomeStatementItemDepreciationAndAmortization`

GetDepreciationAndAmortization returns the DepreciationAndAmortization field if non-nil, zero value otherwise.

### GetDepreciationAndAmortizationOk

`func (o *IncomeStatementItem) GetDepreciationAndAmortizationOk() (*IncomeStatementItemDepreciationAndAmortization, bool)`

GetDepreciationAndAmortizationOk returns a tuple with the DepreciationAndAmortization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationAndAmortization

`func (o *IncomeStatementItem) SetDepreciationAndAmortization(v IncomeStatementItemDepreciationAndAmortization)`

SetDepreciationAndAmortization sets DepreciationAndAmortization field to given value.

### HasDepreciationAndAmortization

`func (o *IncomeStatementItem) HasDepreciationAndAmortization() bool`

HasDepreciationAndAmortization returns a boolean if a field has been set.

### GetEbitda

`func (o *IncomeStatementItem) GetEbitda() IncomeStatementItemEbitda`

GetEbitda returns the Ebitda field if non-nil, zero value otherwise.

### GetEbitdaOk

`func (o *IncomeStatementItem) GetEbitdaOk() (*IncomeStatementItemEbitda, bool)`

GetEbitdaOk returns a tuple with the Ebitda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbitda

`func (o *IncomeStatementItem) SetEbitda(v IncomeStatementItemEbitda)`

SetEbitda sets Ebitda field to given value.

### HasEbitda

`func (o *IncomeStatementItem) HasEbitda() bool`

HasEbitda returns a boolean if a field has been set.

### GetDividendsAndShares

`func (o *IncomeStatementItem) GetDividendsAndShares() IncomeStatementItemDividendsAndShares`

GetDividendsAndShares returns the DividendsAndShares field if non-nil, zero value otherwise.

### GetDividendsAndSharesOk

`func (o *IncomeStatementItem) GetDividendsAndSharesOk() (*IncomeStatementItemDividendsAndShares, bool)`

GetDividendsAndSharesOk returns a tuple with the DividendsAndShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendsAndShares

`func (o *IncomeStatementItem) SetDividendsAndShares(v IncomeStatementItemDividendsAndShares)`

SetDividendsAndShares sets DividendsAndShares field to given value.

### HasDividendsAndShares

`func (o *IncomeStatementItem) HasDividendsAndShares() bool`

HasDividendsAndShares returns a boolean if a field has been set.

### GetUnusualItems

`func (o *IncomeStatementItem) GetUnusualItems() IncomeStatementItemUnusualItems`

GetUnusualItems returns the UnusualItems field if non-nil, zero value otherwise.

### GetUnusualItemsOk

`func (o *IncomeStatementItem) GetUnusualItemsOk() (*IncomeStatementItemUnusualItems, bool)`

GetUnusualItemsOk returns a tuple with the UnusualItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnusualItems

`func (o *IncomeStatementItem) SetUnusualItems(v IncomeStatementItemUnusualItems)`

SetUnusualItems sets UnusualItems field to given value.

### HasUnusualItems

`func (o *IncomeStatementItem) HasUnusualItems() bool`

HasUnusualItems returns a boolean if a field has been set.

### GetDepreciation

`func (o *IncomeStatementItem) GetDepreciation() IncomeStatementItemDepreciation`

GetDepreciation returns the Depreciation field if non-nil, zero value otherwise.

### GetDepreciationOk

`func (o *IncomeStatementItem) GetDepreciationOk() (*IncomeStatementItemDepreciation, bool)`

GetDepreciationOk returns a tuple with the Depreciation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciation

`func (o *IncomeStatementItem) SetDepreciation(v IncomeStatementItemDepreciation)`

SetDepreciation sets Depreciation field to given value.

### HasDepreciation

`func (o *IncomeStatementItem) HasDepreciation() bool`

HasDepreciation returns a boolean if a field has been set.

### GetPretaxIncome

`func (o *IncomeStatementItem) GetPretaxIncome() IncomeStatementItemPretaxIncome`

GetPretaxIncome returns the PretaxIncome field if non-nil, zero value otherwise.

### GetPretaxIncomeOk

`func (o *IncomeStatementItem) GetPretaxIncomeOk() (*IncomeStatementItemPretaxIncome, bool)`

GetPretaxIncomeOk returns a tuple with the PretaxIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPretaxIncome

`func (o *IncomeStatementItem) SetPretaxIncome(v IncomeStatementItemPretaxIncome)`

SetPretaxIncome sets PretaxIncome field to given value.

### HasPretaxIncome

`func (o *IncomeStatementItem) HasPretaxIncome() bool`

HasPretaxIncome returns a boolean if a field has been set.

### GetSpecialIncomeCharges

`func (o *IncomeStatementItem) GetSpecialIncomeCharges() IncomeStatementItemSpecialIncomeCharges`

GetSpecialIncomeCharges returns the SpecialIncomeCharges field if non-nil, zero value otherwise.

### GetSpecialIncomeChargesOk

`func (o *IncomeStatementItem) GetSpecialIncomeChargesOk() (*IncomeStatementItemSpecialIncomeCharges, bool)`

GetSpecialIncomeChargesOk returns a tuple with the SpecialIncomeCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialIncomeCharges

`func (o *IncomeStatementItem) SetSpecialIncomeCharges(v IncomeStatementItemSpecialIncomeCharges)`

SetSpecialIncomeCharges sets SpecialIncomeCharges field to given value.

### HasSpecialIncomeCharges

`func (o *IncomeStatementItem) HasSpecialIncomeCharges() bool`

HasSpecialIncomeCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


