# BalanceSheetConsolidatedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalDate** | **string** | Date of fiscal period ending | 
**Assets** | Pointer to [**AssetsInfo**](AssetsInfo.md) |  | [optional] 

## Methods

### NewBalanceSheetConsolidatedItem

`func NewBalanceSheetConsolidatedItem(fiscalDate string, ) *BalanceSheetConsolidatedItem`

NewBalanceSheetConsolidatedItem instantiates a new BalanceSheetConsolidatedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceSheetConsolidatedItemWithDefaults

`func NewBalanceSheetConsolidatedItemWithDefaults() *BalanceSheetConsolidatedItem`

NewBalanceSheetConsolidatedItemWithDefaults instantiates a new BalanceSheetConsolidatedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalDate

`func (o *BalanceSheetConsolidatedItem) GetFiscalDate() string`

GetFiscalDate returns the FiscalDate field if non-nil, zero value otherwise.

### GetFiscalDateOk

`func (o *BalanceSheetConsolidatedItem) GetFiscalDateOk() (*string, bool)`

GetFiscalDateOk returns a tuple with the FiscalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalDate

`func (o *BalanceSheetConsolidatedItem) SetFiscalDate(v string)`

SetFiscalDate sets FiscalDate field to given value.


### GetAssets

`func (o *BalanceSheetConsolidatedItem) GetAssets() AssetsInfo`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *BalanceSheetConsolidatedItem) GetAssetsOk() (*AssetsInfo, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *BalanceSheetConsolidatedItem) SetAssets(v AssetsInfo)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *BalanceSheetConsolidatedItem) HasAssets() bool`

HasAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


