# CashFlowStructInvestingActivities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapitalExpenditures** | Pointer to **float64** | Capital expenditures (CapEx) are funds used by a company to acquire, upgrade, and maintain physical assets (PPE) | [optional] 
**NetIntangibles** | Pointer to **float64** | Represents purchase of a not physical asset | [optional] 
**NetAcquisitions** | Pointer to **float64** | Refers to net amount of business purchase and sale | [optional] 
**PurchaseOfInvestments** | Pointer to **float64** | Represents how much money has been used in making investments, including purchases of physical assets, investments in securities | [optional] 
**SaleOfInvestments** | Pointer to **float64** | Represents how much money has been generated from the sale of securities or assets | [optional] 
**OtherInvestingActivity** | Pointer to **float64** | Represents other investing activity | [optional] 
**InvestingCashFlow** | Pointer to **float64** | Returns total amount of cash flow used in investments | [optional] 

## Methods

### NewCashFlowStructInvestingActivities

`func NewCashFlowStructInvestingActivities() *CashFlowStructInvestingActivities`

NewCashFlowStructInvestingActivities instantiates a new CashFlowStructInvestingActivities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashFlowStructInvestingActivitiesWithDefaults

`func NewCashFlowStructInvestingActivitiesWithDefaults() *CashFlowStructInvestingActivities`

NewCashFlowStructInvestingActivitiesWithDefaults instantiates a new CashFlowStructInvestingActivities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapitalExpenditures

`func (o *CashFlowStructInvestingActivities) GetCapitalExpenditures() float64`

GetCapitalExpenditures returns the CapitalExpenditures field if non-nil, zero value otherwise.

### GetCapitalExpendituresOk

`func (o *CashFlowStructInvestingActivities) GetCapitalExpendituresOk() (*float64, bool)`

GetCapitalExpendituresOk returns a tuple with the CapitalExpenditures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalExpenditures

`func (o *CashFlowStructInvestingActivities) SetCapitalExpenditures(v float64)`

SetCapitalExpenditures sets CapitalExpenditures field to given value.

### HasCapitalExpenditures

`func (o *CashFlowStructInvestingActivities) HasCapitalExpenditures() bool`

HasCapitalExpenditures returns a boolean if a field has been set.

### GetNetIntangibles

`func (o *CashFlowStructInvestingActivities) GetNetIntangibles() float64`

GetNetIntangibles returns the NetIntangibles field if non-nil, zero value otherwise.

### GetNetIntangiblesOk

`func (o *CashFlowStructInvestingActivities) GetNetIntangiblesOk() (*float64, bool)`

GetNetIntangiblesOk returns a tuple with the NetIntangibles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIntangibles

`func (o *CashFlowStructInvestingActivities) SetNetIntangibles(v float64)`

SetNetIntangibles sets NetIntangibles field to given value.

### HasNetIntangibles

`func (o *CashFlowStructInvestingActivities) HasNetIntangibles() bool`

HasNetIntangibles returns a boolean if a field has been set.

### GetNetAcquisitions

`func (o *CashFlowStructInvestingActivities) GetNetAcquisitions() float64`

GetNetAcquisitions returns the NetAcquisitions field if non-nil, zero value otherwise.

### GetNetAcquisitionsOk

`func (o *CashFlowStructInvestingActivities) GetNetAcquisitionsOk() (*float64, bool)`

GetNetAcquisitionsOk returns a tuple with the NetAcquisitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAcquisitions

`func (o *CashFlowStructInvestingActivities) SetNetAcquisitions(v float64)`

SetNetAcquisitions sets NetAcquisitions field to given value.

### HasNetAcquisitions

`func (o *CashFlowStructInvestingActivities) HasNetAcquisitions() bool`

HasNetAcquisitions returns a boolean if a field has been set.

### GetPurchaseOfInvestments

`func (o *CashFlowStructInvestingActivities) GetPurchaseOfInvestments() float64`

GetPurchaseOfInvestments returns the PurchaseOfInvestments field if non-nil, zero value otherwise.

### GetPurchaseOfInvestmentsOk

`func (o *CashFlowStructInvestingActivities) GetPurchaseOfInvestmentsOk() (*float64, bool)`

GetPurchaseOfInvestmentsOk returns a tuple with the PurchaseOfInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOfInvestments

`func (o *CashFlowStructInvestingActivities) SetPurchaseOfInvestments(v float64)`

SetPurchaseOfInvestments sets PurchaseOfInvestments field to given value.

### HasPurchaseOfInvestments

`func (o *CashFlowStructInvestingActivities) HasPurchaseOfInvestments() bool`

HasPurchaseOfInvestments returns a boolean if a field has been set.

### GetSaleOfInvestments

`func (o *CashFlowStructInvestingActivities) GetSaleOfInvestments() float64`

GetSaleOfInvestments returns the SaleOfInvestments field if non-nil, zero value otherwise.

### GetSaleOfInvestmentsOk

`func (o *CashFlowStructInvestingActivities) GetSaleOfInvestmentsOk() (*float64, bool)`

GetSaleOfInvestmentsOk returns a tuple with the SaleOfInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleOfInvestments

`func (o *CashFlowStructInvestingActivities) SetSaleOfInvestments(v float64)`

SetSaleOfInvestments sets SaleOfInvestments field to given value.

### HasSaleOfInvestments

`func (o *CashFlowStructInvestingActivities) HasSaleOfInvestments() bool`

HasSaleOfInvestments returns a boolean if a field has been set.

### GetOtherInvestingActivity

`func (o *CashFlowStructInvestingActivities) GetOtherInvestingActivity() float64`

GetOtherInvestingActivity returns the OtherInvestingActivity field if non-nil, zero value otherwise.

### GetOtherInvestingActivityOk

`func (o *CashFlowStructInvestingActivities) GetOtherInvestingActivityOk() (*float64, bool)`

GetOtherInvestingActivityOk returns a tuple with the OtherInvestingActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherInvestingActivity

`func (o *CashFlowStructInvestingActivities) SetOtherInvestingActivity(v float64)`

SetOtherInvestingActivity sets OtherInvestingActivity field to given value.

### HasOtherInvestingActivity

`func (o *CashFlowStructInvestingActivities) HasOtherInvestingActivity() bool`

HasOtherInvestingActivity returns a boolean if a field has been set.

### GetInvestingCashFlow

`func (o *CashFlowStructInvestingActivities) GetInvestingCashFlow() float64`

GetInvestingCashFlow returns the InvestingCashFlow field if non-nil, zero value otherwise.

### GetInvestingCashFlowOk

`func (o *CashFlowStructInvestingActivities) GetInvestingCashFlowOk() (*float64, bool)`

GetInvestingCashFlowOk returns a tuple with the InvestingCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestingCashFlow

`func (o *CashFlowStructInvestingActivities) SetInvestingCashFlow(v float64)`

SetInvestingCashFlow sets InvestingCashFlow field to given value.

### HasInvestingCashFlow

`func (o *CashFlowStructInvestingActivities) HasInvestingCashFlow() bool`

HasInvestingCashFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


