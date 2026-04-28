# CashFlowStructFinancingActivities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LongTermDebtIssuance** | Pointer to **float64** | Refers to the issuance of any financial obligations that extend beyond a 12 months period | [optional] 
**LongTermDebtPayments** | Pointer to **float64** | Refers to the payments of any financial obligations that extend beyond a 12 months period | [optional] 
**ShortTermDebtIssuance** | Pointer to **float64** | Refers to the issuance of any financial obligations that are expected to be paid off within a year | [optional] 
**CommonStockIssuance** | Pointer to **float64** | Represents a transaction whereby a company issues its own shares to the marketplace | [optional] 
**CommonStockRepurchase** | Pointer to **float64** | Represents a transaction whereby a company buys back its own shares from the marketplace | [optional] 
**CommonDividends** | Pointer to **float64** | Returns value of payment doled out by a company to its stockholders in the form of periodic distributions of cash | [optional] 
**OtherFinancingCharges** | Pointer to **float64** | Represents other financing charges | [optional] 
**FinancingCashFlow** | Pointer to **float64** | Returns cash flow from financing activities (CFF), which shows the net flows of cash that are used to fund the company | [optional] 

## Methods

### NewCashFlowStructFinancingActivities

`func NewCashFlowStructFinancingActivities() *CashFlowStructFinancingActivities`

NewCashFlowStructFinancingActivities instantiates a new CashFlowStructFinancingActivities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashFlowStructFinancingActivitiesWithDefaults

`func NewCashFlowStructFinancingActivitiesWithDefaults() *CashFlowStructFinancingActivities`

NewCashFlowStructFinancingActivitiesWithDefaults instantiates a new CashFlowStructFinancingActivities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLongTermDebtIssuance

`func (o *CashFlowStructFinancingActivities) GetLongTermDebtIssuance() float64`

GetLongTermDebtIssuance returns the LongTermDebtIssuance field if non-nil, zero value otherwise.

### GetLongTermDebtIssuanceOk

`func (o *CashFlowStructFinancingActivities) GetLongTermDebtIssuanceOk() (*float64, bool)`

GetLongTermDebtIssuanceOk returns a tuple with the LongTermDebtIssuance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermDebtIssuance

`func (o *CashFlowStructFinancingActivities) SetLongTermDebtIssuance(v float64)`

SetLongTermDebtIssuance sets LongTermDebtIssuance field to given value.

### HasLongTermDebtIssuance

`func (o *CashFlowStructFinancingActivities) HasLongTermDebtIssuance() bool`

HasLongTermDebtIssuance returns a boolean if a field has been set.

### GetLongTermDebtPayments

`func (o *CashFlowStructFinancingActivities) GetLongTermDebtPayments() float64`

GetLongTermDebtPayments returns the LongTermDebtPayments field if non-nil, zero value otherwise.

### GetLongTermDebtPaymentsOk

`func (o *CashFlowStructFinancingActivities) GetLongTermDebtPaymentsOk() (*float64, bool)`

GetLongTermDebtPaymentsOk returns a tuple with the LongTermDebtPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermDebtPayments

`func (o *CashFlowStructFinancingActivities) SetLongTermDebtPayments(v float64)`

SetLongTermDebtPayments sets LongTermDebtPayments field to given value.

### HasLongTermDebtPayments

`func (o *CashFlowStructFinancingActivities) HasLongTermDebtPayments() bool`

HasLongTermDebtPayments returns a boolean if a field has been set.

### GetShortTermDebtIssuance

`func (o *CashFlowStructFinancingActivities) GetShortTermDebtIssuance() float64`

GetShortTermDebtIssuance returns the ShortTermDebtIssuance field if non-nil, zero value otherwise.

### GetShortTermDebtIssuanceOk

`func (o *CashFlowStructFinancingActivities) GetShortTermDebtIssuanceOk() (*float64, bool)`

GetShortTermDebtIssuanceOk returns a tuple with the ShortTermDebtIssuance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTermDebtIssuance

`func (o *CashFlowStructFinancingActivities) SetShortTermDebtIssuance(v float64)`

SetShortTermDebtIssuance sets ShortTermDebtIssuance field to given value.

### HasShortTermDebtIssuance

`func (o *CashFlowStructFinancingActivities) HasShortTermDebtIssuance() bool`

HasShortTermDebtIssuance returns a boolean if a field has been set.

### GetCommonStockIssuance

`func (o *CashFlowStructFinancingActivities) GetCommonStockIssuance() float64`

GetCommonStockIssuance returns the CommonStockIssuance field if non-nil, zero value otherwise.

### GetCommonStockIssuanceOk

`func (o *CashFlowStructFinancingActivities) GetCommonStockIssuanceOk() (*float64, bool)`

GetCommonStockIssuanceOk returns a tuple with the CommonStockIssuance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStockIssuance

`func (o *CashFlowStructFinancingActivities) SetCommonStockIssuance(v float64)`

SetCommonStockIssuance sets CommonStockIssuance field to given value.

### HasCommonStockIssuance

`func (o *CashFlowStructFinancingActivities) HasCommonStockIssuance() bool`

HasCommonStockIssuance returns a boolean if a field has been set.

### GetCommonStockRepurchase

`func (o *CashFlowStructFinancingActivities) GetCommonStockRepurchase() float64`

GetCommonStockRepurchase returns the CommonStockRepurchase field if non-nil, zero value otherwise.

### GetCommonStockRepurchaseOk

`func (o *CashFlowStructFinancingActivities) GetCommonStockRepurchaseOk() (*float64, bool)`

GetCommonStockRepurchaseOk returns a tuple with the CommonStockRepurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonStockRepurchase

`func (o *CashFlowStructFinancingActivities) SetCommonStockRepurchase(v float64)`

SetCommonStockRepurchase sets CommonStockRepurchase field to given value.

### HasCommonStockRepurchase

`func (o *CashFlowStructFinancingActivities) HasCommonStockRepurchase() bool`

HasCommonStockRepurchase returns a boolean if a field has been set.

### GetCommonDividends

`func (o *CashFlowStructFinancingActivities) GetCommonDividends() float64`

GetCommonDividends returns the CommonDividends field if non-nil, zero value otherwise.

### GetCommonDividendsOk

`func (o *CashFlowStructFinancingActivities) GetCommonDividendsOk() (*float64, bool)`

GetCommonDividendsOk returns a tuple with the CommonDividends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonDividends

`func (o *CashFlowStructFinancingActivities) SetCommonDividends(v float64)`

SetCommonDividends sets CommonDividends field to given value.

### HasCommonDividends

`func (o *CashFlowStructFinancingActivities) HasCommonDividends() bool`

HasCommonDividends returns a boolean if a field has been set.

### GetOtherFinancingCharges

`func (o *CashFlowStructFinancingActivities) GetOtherFinancingCharges() float64`

GetOtherFinancingCharges returns the OtherFinancingCharges field if non-nil, zero value otherwise.

### GetOtherFinancingChargesOk

`func (o *CashFlowStructFinancingActivities) GetOtherFinancingChargesOk() (*float64, bool)`

GetOtherFinancingChargesOk returns a tuple with the OtherFinancingCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFinancingCharges

`func (o *CashFlowStructFinancingActivities) SetOtherFinancingCharges(v float64)`

SetOtherFinancingCharges sets OtherFinancingCharges field to given value.

### HasOtherFinancingCharges

`func (o *CashFlowStructFinancingActivities) HasOtherFinancingCharges() bool`

HasOtherFinancingCharges returns a boolean if a field has been set.

### GetFinancingCashFlow

`func (o *CashFlowStructFinancingActivities) GetFinancingCashFlow() float64`

GetFinancingCashFlow returns the FinancingCashFlow field if non-nil, zero value otherwise.

### GetFinancingCashFlowOk

`func (o *CashFlowStructFinancingActivities) GetFinancingCashFlowOk() (*float64, bool)`

GetFinancingCashFlowOk returns a tuple with the FinancingCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinancingCashFlow

`func (o *CashFlowStructFinancingActivities) SetFinancingCashFlow(v float64)`

SetFinancingCashFlow sets FinancingCashFlow field to given value.

### HasFinancingCashFlow

`func (o *CashFlowStructFinancingActivities) HasFinancingCashFlow() bool`

HasFinancingCashFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


