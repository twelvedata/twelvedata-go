# CashFlowData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalDate** | **string** | Date of fiscal period ending | 
**Year** | Pointer to **int64** | Year of the cash flow statement | [optional] 
**CashFlowFromOperatingActivities** | Pointer to [**CashFlowDataCashFlowFromOperatingActivities**](CashFlowDataCashFlowFromOperatingActivities.md) |  | [optional] 
**CashFlowFromInvestingActivities** | Pointer to [**CashFlowDataCashFlowFromInvestingActivities**](CashFlowDataCashFlowFromInvestingActivities.md) |  | [optional] 
**CashFlowFromFinancingActivities** | Pointer to [**CashFlowDataCashFlowFromFinancingActivities**](CashFlowDataCashFlowFromFinancingActivities.md) |  | [optional] 
**SupplementalData** | Pointer to [**CashFlowDataSupplementalData**](CashFlowDataSupplementalData.md) |  | [optional] 
**ForeignAndDomesticSales** | Pointer to [**CashFlowDataForeignAndDomesticSales**](CashFlowDataForeignAndDomesticSales.md) |  | [optional] 
**CashPosition** | Pointer to [**CashFlowDataCashPosition**](CashFlowDataCashPosition.md) |  | [optional] 
**DirectMethodCashFlow** | Pointer to [**CashFlowDataDirectMethodCashFlow**](CashFlowDataDirectMethodCashFlow.md) |  | [optional] 

## Methods

### NewCashFlowData

`func NewCashFlowData(fiscalDate string, ) *CashFlowData`

NewCashFlowData instantiates a new CashFlowData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashFlowDataWithDefaults

`func NewCashFlowDataWithDefaults() *CashFlowData`

NewCashFlowDataWithDefaults instantiates a new CashFlowData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalDate

`func (o *CashFlowData) GetFiscalDate() string`

GetFiscalDate returns the FiscalDate field if non-nil, zero value otherwise.

### GetFiscalDateOk

`func (o *CashFlowData) GetFiscalDateOk() (*string, bool)`

GetFiscalDateOk returns a tuple with the FiscalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalDate

`func (o *CashFlowData) SetFiscalDate(v string)`

SetFiscalDate sets FiscalDate field to given value.


### GetYear

`func (o *CashFlowData) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *CashFlowData) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *CashFlowData) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *CashFlowData) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetCashFlowFromOperatingActivities

`func (o *CashFlowData) GetCashFlowFromOperatingActivities() CashFlowDataCashFlowFromOperatingActivities`

GetCashFlowFromOperatingActivities returns the CashFlowFromOperatingActivities field if non-nil, zero value otherwise.

### GetCashFlowFromOperatingActivitiesOk

`func (o *CashFlowData) GetCashFlowFromOperatingActivitiesOk() (*CashFlowDataCashFlowFromOperatingActivities, bool)`

GetCashFlowFromOperatingActivitiesOk returns a tuple with the CashFlowFromOperatingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowFromOperatingActivities

`func (o *CashFlowData) SetCashFlowFromOperatingActivities(v CashFlowDataCashFlowFromOperatingActivities)`

SetCashFlowFromOperatingActivities sets CashFlowFromOperatingActivities field to given value.

### HasCashFlowFromOperatingActivities

`func (o *CashFlowData) HasCashFlowFromOperatingActivities() bool`

HasCashFlowFromOperatingActivities returns a boolean if a field has been set.

### GetCashFlowFromInvestingActivities

`func (o *CashFlowData) GetCashFlowFromInvestingActivities() CashFlowDataCashFlowFromInvestingActivities`

GetCashFlowFromInvestingActivities returns the CashFlowFromInvestingActivities field if non-nil, zero value otherwise.

### GetCashFlowFromInvestingActivitiesOk

`func (o *CashFlowData) GetCashFlowFromInvestingActivitiesOk() (*CashFlowDataCashFlowFromInvestingActivities, bool)`

GetCashFlowFromInvestingActivitiesOk returns a tuple with the CashFlowFromInvestingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowFromInvestingActivities

`func (o *CashFlowData) SetCashFlowFromInvestingActivities(v CashFlowDataCashFlowFromInvestingActivities)`

SetCashFlowFromInvestingActivities sets CashFlowFromInvestingActivities field to given value.

### HasCashFlowFromInvestingActivities

`func (o *CashFlowData) HasCashFlowFromInvestingActivities() bool`

HasCashFlowFromInvestingActivities returns a boolean if a field has been set.

### GetCashFlowFromFinancingActivities

`func (o *CashFlowData) GetCashFlowFromFinancingActivities() CashFlowDataCashFlowFromFinancingActivities`

GetCashFlowFromFinancingActivities returns the CashFlowFromFinancingActivities field if non-nil, zero value otherwise.

### GetCashFlowFromFinancingActivitiesOk

`func (o *CashFlowData) GetCashFlowFromFinancingActivitiesOk() (*CashFlowDataCashFlowFromFinancingActivities, bool)`

GetCashFlowFromFinancingActivitiesOk returns a tuple with the CashFlowFromFinancingActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlowFromFinancingActivities

`func (o *CashFlowData) SetCashFlowFromFinancingActivities(v CashFlowDataCashFlowFromFinancingActivities)`

SetCashFlowFromFinancingActivities sets CashFlowFromFinancingActivities field to given value.

### HasCashFlowFromFinancingActivities

`func (o *CashFlowData) HasCashFlowFromFinancingActivities() bool`

HasCashFlowFromFinancingActivities returns a boolean if a field has been set.

### GetSupplementalData

`func (o *CashFlowData) GetSupplementalData() CashFlowDataSupplementalData`

GetSupplementalData returns the SupplementalData field if non-nil, zero value otherwise.

### GetSupplementalDataOk

`func (o *CashFlowData) GetSupplementalDataOk() (*CashFlowDataSupplementalData, bool)`

GetSupplementalDataOk returns a tuple with the SupplementalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalData

`func (o *CashFlowData) SetSupplementalData(v CashFlowDataSupplementalData)`

SetSupplementalData sets SupplementalData field to given value.

### HasSupplementalData

`func (o *CashFlowData) HasSupplementalData() bool`

HasSupplementalData returns a boolean if a field has been set.

### GetForeignAndDomesticSales

`func (o *CashFlowData) GetForeignAndDomesticSales() CashFlowDataForeignAndDomesticSales`

GetForeignAndDomesticSales returns the ForeignAndDomesticSales field if non-nil, zero value otherwise.

### GetForeignAndDomesticSalesOk

`func (o *CashFlowData) GetForeignAndDomesticSalesOk() (*CashFlowDataForeignAndDomesticSales, bool)`

GetForeignAndDomesticSalesOk returns a tuple with the ForeignAndDomesticSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAndDomesticSales

`func (o *CashFlowData) SetForeignAndDomesticSales(v CashFlowDataForeignAndDomesticSales)`

SetForeignAndDomesticSales sets ForeignAndDomesticSales field to given value.

### HasForeignAndDomesticSales

`func (o *CashFlowData) HasForeignAndDomesticSales() bool`

HasForeignAndDomesticSales returns a boolean if a field has been set.

### GetCashPosition

`func (o *CashFlowData) GetCashPosition() CashFlowDataCashPosition`

GetCashPosition returns the CashPosition field if non-nil, zero value otherwise.

### GetCashPositionOk

`func (o *CashFlowData) GetCashPositionOk() (*CashFlowDataCashPosition, bool)`

GetCashPositionOk returns a tuple with the CashPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashPosition

`func (o *CashFlowData) SetCashPosition(v CashFlowDataCashPosition)`

SetCashPosition sets CashPosition field to given value.

### HasCashPosition

`func (o *CashFlowData) HasCashPosition() bool`

HasCashPosition returns a boolean if a field has been set.

### GetDirectMethodCashFlow

`func (o *CashFlowData) GetDirectMethodCashFlow() CashFlowDataDirectMethodCashFlow`

GetDirectMethodCashFlow returns the DirectMethodCashFlow field if non-nil, zero value otherwise.

### GetDirectMethodCashFlowOk

`func (o *CashFlowData) GetDirectMethodCashFlowOk() (*CashFlowDataDirectMethodCashFlow, bool)`

GetDirectMethodCashFlowOk returns a tuple with the DirectMethodCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMethodCashFlow

`func (o *CashFlowData) SetDirectMethodCashFlow(v CashFlowDataDirectMethodCashFlow)`

SetDirectMethodCashFlow sets DirectMethodCashFlow field to given value.

### HasDirectMethodCashFlow

`func (o *CashFlowData) HasDirectMethodCashFlow() bool`

HasDirectMethodCashFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


