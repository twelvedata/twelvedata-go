# GetStatistics200ResponseStatisticsFinancials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalYearEnds** | Pointer to **string** | Refers to the completion of the last 12-month accounting period | [optional] 
**MostRecentQuarter** | Pointer to **string** | The most recent quarter (MRQ) refers to the fiscal quarter that most recently ended | [optional] 
**GrossMargin** | Pointer to **float64** | The portion of a company&#39;s revenue left over after direct costs are subtracted | [optional] 
**ProfitMargin** | Pointer to **float64** | Refers to gross profit margin. Calculated by dividing net income by sales revenue | [optional] 
**OperatingMargin** | Pointer to **float64** | Operating margin is calculated by dividing operating income by net sales | [optional] 
**ReturnOnAssetsTtm** | Pointer to **float64** | Return on assets (ROA) is calculated by dividing net income by total assets | [optional] 
**ReturnOnEquityTtm** | Pointer to **float64** | Return on equity (ROE) is calculated by dividing net income by average shareholders&#39; equity | [optional] 
**IncomeStatement** | Pointer to [**GetStatistics200ResponseStatisticsFinancialsIncomeStatement**](GetStatistics200ResponseStatisticsFinancialsIncomeStatement.md) |  | [optional] 
**BalanceSheet** | Pointer to [**GetStatistics200ResponseStatisticsFinancialsBalanceSheet**](GetStatistics200ResponseStatisticsFinancialsBalanceSheet.md) |  | [optional] 
**CashFlow** | Pointer to [**GetStatistics200ResponseStatisticsFinancialsCashFlow**](GetStatistics200ResponseStatisticsFinancialsCashFlow.md) |  | [optional] 

## Methods

### NewGetStatistics200ResponseStatisticsFinancials

`func NewGetStatistics200ResponseStatisticsFinancials() *GetStatistics200ResponseStatisticsFinancials`

NewGetStatistics200ResponseStatisticsFinancials instantiates a new GetStatistics200ResponseStatisticsFinancials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatistics200ResponseStatisticsFinancialsWithDefaults

`func NewGetStatistics200ResponseStatisticsFinancialsWithDefaults() *GetStatistics200ResponseStatisticsFinancials`

NewGetStatistics200ResponseStatisticsFinancialsWithDefaults instantiates a new GetStatistics200ResponseStatisticsFinancials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalYearEnds

`func (o *GetStatistics200ResponseStatisticsFinancials) GetFiscalYearEnds() string`

GetFiscalYearEnds returns the FiscalYearEnds field if non-nil, zero value otherwise.

### GetFiscalYearEndsOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetFiscalYearEndsOk() (*string, bool)`

GetFiscalYearEndsOk returns a tuple with the FiscalYearEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearEnds

`func (o *GetStatistics200ResponseStatisticsFinancials) SetFiscalYearEnds(v string)`

SetFiscalYearEnds sets FiscalYearEnds field to given value.

### HasFiscalYearEnds

`func (o *GetStatistics200ResponseStatisticsFinancials) HasFiscalYearEnds() bool`

HasFiscalYearEnds returns a boolean if a field has been set.

### GetMostRecentQuarter

`func (o *GetStatistics200ResponseStatisticsFinancials) GetMostRecentQuarter() string`

GetMostRecentQuarter returns the MostRecentQuarter field if non-nil, zero value otherwise.

### GetMostRecentQuarterOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetMostRecentQuarterOk() (*string, bool)`

GetMostRecentQuarterOk returns a tuple with the MostRecentQuarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentQuarter

`func (o *GetStatistics200ResponseStatisticsFinancials) SetMostRecentQuarter(v string)`

SetMostRecentQuarter sets MostRecentQuarter field to given value.

### HasMostRecentQuarter

`func (o *GetStatistics200ResponseStatisticsFinancials) HasMostRecentQuarter() bool`

HasMostRecentQuarter returns a boolean if a field has been set.

### GetGrossMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) GetGrossMargin() float64`

GetGrossMargin returns the GrossMargin field if non-nil, zero value otherwise.

### GetGrossMarginOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetGrossMarginOk() (*float64, bool)`

GetGrossMarginOk returns a tuple with the GrossMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) SetGrossMargin(v float64)`

SetGrossMargin sets GrossMargin field to given value.

### HasGrossMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) HasGrossMargin() bool`

HasGrossMargin returns a boolean if a field has been set.

### GetProfitMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) GetProfitMargin() float64`

GetProfitMargin returns the ProfitMargin field if non-nil, zero value otherwise.

### GetProfitMarginOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetProfitMarginOk() (*float64, bool)`

GetProfitMarginOk returns a tuple with the ProfitMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) SetProfitMargin(v float64)`

SetProfitMargin sets ProfitMargin field to given value.

### HasProfitMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) HasProfitMargin() bool`

HasProfitMargin returns a boolean if a field has been set.

### GetOperatingMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) GetOperatingMargin() float64`

GetOperatingMargin returns the OperatingMargin field if non-nil, zero value otherwise.

### GetOperatingMarginOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetOperatingMarginOk() (*float64, bool)`

GetOperatingMarginOk returns a tuple with the OperatingMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) SetOperatingMargin(v float64)`

SetOperatingMargin sets OperatingMargin field to given value.

### HasOperatingMargin

`func (o *GetStatistics200ResponseStatisticsFinancials) HasOperatingMargin() bool`

HasOperatingMargin returns a boolean if a field has been set.

### GetReturnOnAssetsTtm

`func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnAssetsTtm() float64`

GetReturnOnAssetsTtm returns the ReturnOnAssetsTtm field if non-nil, zero value otherwise.

### GetReturnOnAssetsTtmOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnAssetsTtmOk() (*float64, bool)`

GetReturnOnAssetsTtmOk returns a tuple with the ReturnOnAssetsTtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnAssetsTtm

`func (o *GetStatistics200ResponseStatisticsFinancials) SetReturnOnAssetsTtm(v float64)`

SetReturnOnAssetsTtm sets ReturnOnAssetsTtm field to given value.

### HasReturnOnAssetsTtm

`func (o *GetStatistics200ResponseStatisticsFinancials) HasReturnOnAssetsTtm() bool`

HasReturnOnAssetsTtm returns a boolean if a field has been set.

### GetReturnOnEquityTtm

`func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnEquityTtm() float64`

GetReturnOnEquityTtm returns the ReturnOnEquityTtm field if non-nil, zero value otherwise.

### GetReturnOnEquityTtmOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetReturnOnEquityTtmOk() (*float64, bool)`

GetReturnOnEquityTtmOk returns a tuple with the ReturnOnEquityTtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnEquityTtm

`func (o *GetStatistics200ResponseStatisticsFinancials) SetReturnOnEquityTtm(v float64)`

SetReturnOnEquityTtm sets ReturnOnEquityTtm field to given value.

### HasReturnOnEquityTtm

`func (o *GetStatistics200ResponseStatisticsFinancials) HasReturnOnEquityTtm() bool`

HasReturnOnEquityTtm returns a boolean if a field has been set.

### GetIncomeStatement

`func (o *GetStatistics200ResponseStatisticsFinancials) GetIncomeStatement() GetStatistics200ResponseStatisticsFinancialsIncomeStatement`

GetIncomeStatement returns the IncomeStatement field if non-nil, zero value otherwise.

### GetIncomeStatementOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetIncomeStatementOk() (*GetStatistics200ResponseStatisticsFinancialsIncomeStatement, bool)`

GetIncomeStatementOk returns a tuple with the IncomeStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeStatement

`func (o *GetStatistics200ResponseStatisticsFinancials) SetIncomeStatement(v GetStatistics200ResponseStatisticsFinancialsIncomeStatement)`

SetIncomeStatement sets IncomeStatement field to given value.

### HasIncomeStatement

`func (o *GetStatistics200ResponseStatisticsFinancials) HasIncomeStatement() bool`

HasIncomeStatement returns a boolean if a field has been set.

### GetBalanceSheet

`func (o *GetStatistics200ResponseStatisticsFinancials) GetBalanceSheet() GetStatistics200ResponseStatisticsFinancialsBalanceSheet`

GetBalanceSheet returns the BalanceSheet field if non-nil, zero value otherwise.

### GetBalanceSheetOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetBalanceSheetOk() (*GetStatistics200ResponseStatisticsFinancialsBalanceSheet, bool)`

GetBalanceSheetOk returns a tuple with the BalanceSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSheet

`func (o *GetStatistics200ResponseStatisticsFinancials) SetBalanceSheet(v GetStatistics200ResponseStatisticsFinancialsBalanceSheet)`

SetBalanceSheet sets BalanceSheet field to given value.

### HasBalanceSheet

`func (o *GetStatistics200ResponseStatisticsFinancials) HasBalanceSheet() bool`

HasBalanceSheet returns a boolean if a field has been set.

### GetCashFlow

`func (o *GetStatistics200ResponseStatisticsFinancials) GetCashFlow() GetStatistics200ResponseStatisticsFinancialsCashFlow`

GetCashFlow returns the CashFlow field if non-nil, zero value otherwise.

### GetCashFlowOk

`func (o *GetStatistics200ResponseStatisticsFinancials) GetCashFlowOk() (*GetStatistics200ResponseStatisticsFinancialsCashFlow, bool)`

GetCashFlowOk returns a tuple with the CashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlow

`func (o *GetStatistics200ResponseStatisticsFinancials) SetCashFlow(v GetStatistics200ResponseStatisticsFinancialsCashFlow)`

SetCashFlow sets CashFlow field to given value.

### HasCashFlow

`func (o *GetStatistics200ResponseStatisticsFinancials) HasCashFlow() bool`

HasCashFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


