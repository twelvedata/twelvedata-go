# AssetsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAssets** | Pointer to **float64** | Total assets | [optional] 
**CurrentAssets** | Pointer to [**AssetsInfoCurrentAssets**](AssetsInfoCurrentAssets.md) |  | [optional] 
**NonCurrentAssets** | Pointer to [**AssetsInfoNonCurrentAssets**](AssetsInfoNonCurrentAssets.md) |  | [optional] 
**Liabilities** | Pointer to [**AssetsInfoLiabilities**](AssetsInfoLiabilities.md) |  | [optional] 

## Methods

### NewAssetsInfo

`func NewAssetsInfo() *AssetsInfo`

NewAssetsInfo instantiates a new AssetsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsInfoWithDefaults

`func NewAssetsInfoWithDefaults() *AssetsInfo`

NewAssetsInfoWithDefaults instantiates a new AssetsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAssets

`func (o *AssetsInfo) GetTotalAssets() float64`

GetTotalAssets returns the TotalAssets field if non-nil, zero value otherwise.

### GetTotalAssetsOk

`func (o *AssetsInfo) GetTotalAssetsOk() (*float64, bool)`

GetTotalAssetsOk returns a tuple with the TotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssets

`func (o *AssetsInfo) SetTotalAssets(v float64)`

SetTotalAssets sets TotalAssets field to given value.

### HasTotalAssets

`func (o *AssetsInfo) HasTotalAssets() bool`

HasTotalAssets returns a boolean if a field has been set.

### GetCurrentAssets

`func (o *AssetsInfo) GetCurrentAssets() AssetsInfoCurrentAssets`

GetCurrentAssets returns the CurrentAssets field if non-nil, zero value otherwise.

### GetCurrentAssetsOk

`func (o *AssetsInfo) GetCurrentAssetsOk() (*AssetsInfoCurrentAssets, bool)`

GetCurrentAssetsOk returns a tuple with the CurrentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAssets

`func (o *AssetsInfo) SetCurrentAssets(v AssetsInfoCurrentAssets)`

SetCurrentAssets sets CurrentAssets field to given value.

### HasCurrentAssets

`func (o *AssetsInfo) HasCurrentAssets() bool`

HasCurrentAssets returns a boolean if a field has been set.

### GetNonCurrentAssets

`func (o *AssetsInfo) GetNonCurrentAssets() AssetsInfoNonCurrentAssets`

GetNonCurrentAssets returns the NonCurrentAssets field if non-nil, zero value otherwise.

### GetNonCurrentAssetsOk

`func (o *AssetsInfo) GetNonCurrentAssetsOk() (*AssetsInfoNonCurrentAssets, bool)`

GetNonCurrentAssetsOk returns a tuple with the NonCurrentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonCurrentAssets

`func (o *AssetsInfo) SetNonCurrentAssets(v AssetsInfoNonCurrentAssets)`

SetNonCurrentAssets sets NonCurrentAssets field to given value.

### HasNonCurrentAssets

`func (o *AssetsInfo) HasNonCurrentAssets() bool`

HasNonCurrentAssets returns a boolean if a field has been set.

### GetLiabilities

`func (o *AssetsInfo) GetLiabilities() AssetsInfoLiabilities`

GetLiabilities returns the Liabilities field if non-nil, zero value otherwise.

### GetLiabilitiesOk

`func (o *AssetsInfo) GetLiabilitiesOk() (*AssetsInfoLiabilities, bool)`

GetLiabilitiesOk returns a tuple with the Liabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilities

`func (o *AssetsInfo) SetLiabilities(v AssetsInfoLiabilities)`

SetLiabilities sets Liabilities field to given value.

### HasLiabilities

`func (o *AssetsInfo) HasLiabilities() bool`

HasLiabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


