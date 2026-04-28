# GetBalanceSheet200ResponseBalanceSheetInnerAssets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAssets** | Pointer to [**GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets**](GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets.md) |  | [optional] 
**NonCurrentAssets** | Pointer to [**GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets**](GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets.md) |  | [optional] 
**TotalAssets** | Pointer to **float64** | The sum of total_current_assets + total_non_current_assets | [optional] 

## Methods

### NewGetBalanceSheet200ResponseBalanceSheetInnerAssets

`func NewGetBalanceSheet200ResponseBalanceSheetInnerAssets() *GetBalanceSheet200ResponseBalanceSheetInnerAssets`

NewGetBalanceSheet200ResponseBalanceSheetInnerAssets instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsWithDefaults

`func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerAssets`

NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetCurrentAssets() GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets`

GetCurrentAssets returns the CurrentAssets field if non-nil, zero value otherwise.

### GetCurrentAssetsOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetCurrentAssetsOk() (*GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets, bool)`

GetCurrentAssetsOk returns a tuple with the CurrentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) SetCurrentAssets(v GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets)`

SetCurrentAssets sets CurrentAssets field to given value.

### HasCurrentAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) HasCurrentAssets() bool`

HasCurrentAssets returns a boolean if a field has been set.

### GetNonCurrentAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetNonCurrentAssets() GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets`

GetNonCurrentAssets returns the NonCurrentAssets field if non-nil, zero value otherwise.

### GetNonCurrentAssetsOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetNonCurrentAssetsOk() (*GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets, bool)`

GetNonCurrentAssetsOk returns a tuple with the NonCurrentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrentAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) SetNonCurrentAssets(v GetBalanceSheet200ResponseBalanceSheetInnerAssetsNonCurrentAssets)`

SetNonCurrentAssets sets NonCurrentAssets field to given value.

### HasNonCurrentAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) HasNonCurrentAssets() bool`

HasNonCurrentAssets returns a boolean if a field has been set.

### GetTotalAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetTotalAssets() float64`

GetTotalAssets returns the TotalAssets field if non-nil, zero value otherwise.

### GetTotalAssetsOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) GetTotalAssetsOk() (*float64, bool)`

GetTotalAssetsOk returns a tuple with the TotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) SetTotalAssets(v float64)`

SetTotalAssets sets TotalAssets field to given value.

### HasTotalAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssets) HasTotalAssets() bool`

HasTotalAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


