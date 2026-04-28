# GetBalanceSheetConsolidated200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceSheet** | [**[]BalanceSheetConsolidatedItem**](BalanceSheetConsolidatedItem.md) | Balance sheet data | 
**Status** | Pointer to **string** | Response status | [optional] 

## Methods

### NewGetBalanceSheetConsolidated200Response

`func NewGetBalanceSheetConsolidated200Response(balanceSheet []BalanceSheetConsolidatedItem, ) *GetBalanceSheetConsolidated200Response`

NewGetBalanceSheetConsolidated200Response instantiates a new GetBalanceSheetConsolidated200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceSheetConsolidated200ResponseWithDefaults

`func NewGetBalanceSheetConsolidated200ResponseWithDefaults() *GetBalanceSheetConsolidated200Response`

NewGetBalanceSheetConsolidated200ResponseWithDefaults instantiates a new GetBalanceSheetConsolidated200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceSheet

`func (o *GetBalanceSheetConsolidated200Response) GetBalanceSheet() []BalanceSheetConsolidatedItem`

GetBalanceSheet returns the BalanceSheet field if non-nil, zero value otherwise.

### GetBalanceSheetOk

`func (o *GetBalanceSheetConsolidated200Response) GetBalanceSheetOk() (*[]BalanceSheetConsolidatedItem, bool)`

GetBalanceSheetOk returns a tuple with the BalanceSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSheet

`func (o *GetBalanceSheetConsolidated200Response) SetBalanceSheet(v []BalanceSheetConsolidatedItem)`

SetBalanceSheet sets BalanceSheet field to given value.


### GetStatus

`func (o *GetBalanceSheetConsolidated200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBalanceSheetConsolidated200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBalanceSheetConsolidated200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBalanceSheetConsolidated200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


