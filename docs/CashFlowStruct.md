# CashFlowStruct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalDate** | **string** | Date of fiscal period ending | 
**Quarter** | Pointer to **string** | Fiscal quarter. Visible when &#x60;&amp;period&#x3D;quarterly&#x60; | [optional] 
**Year** | Pointer to **int64** | Fiscal year | [optional] 
**OperatingActivities** | Pointer to [**CashFlowStructOperatingActivities**](CashFlowStructOperatingActivities.md) |  | [optional] 
**InvestingActivities** | Pointer to [**CashFlowStructInvestingActivities**](CashFlowStructInvestingActivities.md) |  | [optional] 
**FinancingActivities** | Pointer to [**CashFlowStructFinancingActivities**](CashFlowStructFinancingActivities.md) |  | [optional] 
**EndCashPosition** | Pointer to **float64** | Returns the amount of cash a company has when adding the change in cash and beginning cash balance for the current fiscal period | [optional] 
**IncomeTaxPaid** | Pointer to **float64** | Refers to supplemental data about income tax paid | [optional] 
**InterestPaid** | Pointer to **float64** | Refers to supplemental data about interest paid | [optional] 
**FreeCashFlow** | Pointer to **float64** | Represents the cash a company generates after accounting for cash outflows to support operations and maintain its capital assets | [optional] 

## Methods

### NewCashFlowStruct

`func NewCashFlowStruct(fiscalDate string, ) *CashFlowStruct`

NewCashFlowStruct instantiates a new CashFlowStruct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashFlowStructWithDefaults

`func NewCashFlowStructWithDefaults() *CashFlowStruct`

NewCashFlowStructWithDefaults instantiates a new CashFlowStruct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalDate

`func (o *CashFlowStruct) GetFiscalDate() string`

GetFiscalDate returns the FiscalDate field if non-nil, zero value otherwise.

### GetFiscalDateOk

`func (o *CashFlowStruct) GetFiscalDateOk() (*string, bool)`

GetFiscalDateOk returns a tuple with the FiscalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalDate

`func (o *CashFlowStruct) SetFiscalDate(v string)`

SetFiscalDate sets FiscalDate field to given value.


### GetQuarter

`func (o *CashFlowStruct) GetQuarter() string`

GetQuarter returns the Quarter field if non-nil, zero value otherwise.

### GetQuarterOk

`func (o *CashFlowStruct) GetQuarterOk() (*string, bool)`

GetQuarterOk returns a tuple with the Quarter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarter

`func (o *CashFlowStruct) SetQuarter(v string)`

SetQuarter sets Quarter field to given value.

### HasQuarter

`func (o *CashFlowStruct) HasQuarter() bool`

HasQuarter returns a boolean if a field has been set.

### GetYear

`func (o *CashFlowStruct) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *CashFlowStruct) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *CashFlowStruct) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *CashFlowStruct) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetOperatingActivities

`func (o *CashFlowStruct) GetOperatingActivities() CashFlowStructOperatingActivities`

GetOperatingActivities returns the OperatingActivities field if non-nil, zero value otherwise.

### GetOperatingActivitiesOk

`func (o *CashFlowStruct) GetOperatingActivitiesOk() (*CashFlowStructOperatingActivities, bool)`

GetOperatingActivitiesOk returns a tuple with the OperatingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingActivities

`func (o *CashFlowStruct) SetOperatingActivities(v CashFlowStructOperatingActivities)`

SetOperatingActivities sets OperatingActivities field to given value.

### HasOperatingActivities

`func (o *CashFlowStruct) HasOperatingActivities() bool`

HasOperatingActivities returns a boolean if a field has been set.

### GetInvestingActivities

`func (o *CashFlowStruct) GetInvestingActivities() CashFlowStructInvestingActivities`

GetInvestingActivities returns the InvestingActivities field if non-nil, zero value otherwise.

### GetInvestingActivitiesOk

`func (o *CashFlowStruct) GetInvestingActivitiesOk() (*CashFlowStructInvestingActivities, bool)`

GetInvestingActivitiesOk returns a tuple with the InvestingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestingActivities

`func (o *CashFlowStruct) SetInvestingActivities(v CashFlowStructInvestingActivities)`

SetInvestingActivities sets InvestingActivities field to given value.

### HasInvestingActivities

`func (o *CashFlowStruct) HasInvestingActivities() bool`

HasInvestingActivities returns a boolean if a field has been set.

### GetFinancingActivities

`func (o *CashFlowStruct) GetFinancingActivities() CashFlowStructFinancingActivities`

GetFinancingActivities returns the FinancingActivities field if non-nil, zero value otherwise.

### GetFinancingActivitiesOk

`func (o *CashFlowStruct) GetFinancingActivitiesOk() (*CashFlowStructFinancingActivities, bool)`

GetFinancingActivitiesOk returns a tuple with the FinancingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancingActivities

`func (o *CashFlowStruct) SetFinancingActivities(v CashFlowStructFinancingActivities)`

SetFinancingActivities sets FinancingActivities field to given value.

### HasFinancingActivities

`func (o *CashFlowStruct) HasFinancingActivities() bool`

HasFinancingActivities returns a boolean if a field has been set.

### GetEndCashPosition

`func (o *CashFlowStruct) GetEndCashPosition() float64`

GetEndCashPosition returns the EndCashPosition field if non-nil, zero value otherwise.

### GetEndCashPositionOk

`func (o *CashFlowStruct) GetEndCashPositionOk() (*float64, bool)`

GetEndCashPositionOk returns a tuple with the EndCashPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCashPosition

`func (o *CashFlowStruct) SetEndCashPosition(v float64)`

SetEndCashPosition sets EndCashPosition field to given value.

### HasEndCashPosition

`func (o *CashFlowStruct) HasEndCashPosition() bool`

HasEndCashPosition returns a boolean if a field has been set.

### GetIncomeTaxPaid

`func (o *CashFlowStruct) GetIncomeTaxPaid() float64`

GetIncomeTaxPaid returns the IncomeTaxPaid field if non-nil, zero value otherwise.

### GetIncomeTaxPaidOk

`func (o *CashFlowStruct) GetIncomeTaxPaidOk() (*float64, bool)`

GetIncomeTaxPaidOk returns a tuple with the IncomeTaxPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeTaxPaid

`func (o *CashFlowStruct) SetIncomeTaxPaid(v float64)`

SetIncomeTaxPaid sets IncomeTaxPaid field to given value.

### HasIncomeTaxPaid

`func (o *CashFlowStruct) HasIncomeTaxPaid() bool`

HasIncomeTaxPaid returns a boolean if a field has been set.

### GetInterestPaid

`func (o *CashFlowStruct) GetInterestPaid() float64`

GetInterestPaid returns the InterestPaid field if non-nil, zero value otherwise.

### GetInterestPaidOk

`func (o *CashFlowStruct) GetInterestPaidOk() (*float64, bool)`

GetInterestPaidOk returns a tuple with the InterestPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestPaid

`func (o *CashFlowStruct) SetInterestPaid(v float64)`

SetInterestPaid sets InterestPaid field to given value.

### HasInterestPaid

`func (o *CashFlowStruct) HasInterestPaid() bool`

HasInterestPaid returns a boolean if a field has been set.

### GetFreeCashFlow

`func (o *CashFlowStruct) GetFreeCashFlow() float64`

GetFreeCashFlow returns the FreeCashFlow field if non-nil, zero value otherwise.

### GetFreeCashFlowOk

`func (o *CashFlowStruct) GetFreeCashFlowOk() (*float64, bool)`

GetFreeCashFlowOk returns a tuple with the FreeCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCashFlow

`func (o *CashFlowStruct) SetFreeCashFlow(v float64)`

SetFreeCashFlow sets FreeCashFlow field to given value.

### HasFreeCashFlow

`func (o *CashFlowStruct) HasFreeCashFlow() bool`

HasFreeCashFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


