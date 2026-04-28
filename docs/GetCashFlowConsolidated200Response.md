# GetCashFlowConsolidated200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashFlow** | [**[]CashFlowData**](CashFlowData.md) | Cash flow data | 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetCashFlowConsolidated200Response

`func NewGetCashFlowConsolidated200Response(cashFlow []CashFlowData, ) *GetCashFlowConsolidated200Response`

NewGetCashFlowConsolidated200Response instantiates a new GetCashFlowConsolidated200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCashFlowConsolidated200ResponseWithDefaults

`func NewGetCashFlowConsolidated200ResponseWithDefaults() *GetCashFlowConsolidated200Response`

NewGetCashFlowConsolidated200ResponseWithDefaults instantiates a new GetCashFlowConsolidated200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashFlow

`func (o *GetCashFlowConsolidated200Response) GetCashFlow() []CashFlowData`

GetCashFlow returns the CashFlow field if non-nil, zero value otherwise.

### GetCashFlowOk

`func (o *GetCashFlowConsolidated200Response) GetCashFlowOk() (*[]CashFlowData, bool)`

GetCashFlowOk returns a tuple with the CashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlow

`func (o *GetCashFlowConsolidated200Response) SetCashFlow(v []CashFlowData)`

SetCashFlow sets CashFlow field to given value.


### GetStatus

`func (o *GetCashFlowConsolidated200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCashFlowConsolidated200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCashFlowConsolidated200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCashFlowConsolidated200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


