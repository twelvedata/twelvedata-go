# GetBalanceSheet200ResponseBalanceSheetInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiscalDate** | **string** | Date of fiscal period ending | 
**Year** | Pointer to **int64** | Fiscal year | [optional] 
**Assets** | Pointer to [**GetBalanceSheet200ResponseBalanceSheetInnerAssets**](GetBalanceSheet200ResponseBalanceSheetInnerAssets.md) |  | [optional] 
**Liabilities** | Pointer to [**GetBalanceSheet200ResponseBalanceSheetInnerLiabilities**](GetBalanceSheet200ResponseBalanceSheetInnerLiabilities.md) |  | [optional] 
**ShareholdersEquity** | Pointer to [**GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity**](GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity.md) |  | [optional] 

## Methods

### NewGetBalanceSheet200ResponseBalanceSheetInner

`func NewGetBalanceSheet200ResponseBalanceSheetInner(fiscalDate string, ) *GetBalanceSheet200ResponseBalanceSheetInner`

NewGetBalanceSheet200ResponseBalanceSheetInner instantiates a new GetBalanceSheet200ResponseBalanceSheetInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceSheet200ResponseBalanceSheetInnerWithDefaults

`func NewGetBalanceSheet200ResponseBalanceSheetInnerWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInner`

NewGetBalanceSheet200ResponseBalanceSheetInnerWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiscalDate

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetFiscalDate() string`

GetFiscalDate returns the FiscalDate field if non-nil, zero value otherwise.

### GetFiscalDateOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetFiscalDateOk() (*string, bool)`

GetFiscalDateOk returns a tuple with the FiscalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalDate

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetFiscalDate(v string)`

SetFiscalDate sets FiscalDate field to given value.


### GetYear

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetAssets() GetBalanceSheet200ResponseBalanceSheetInnerAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetAssetsOk() (*GetBalanceSheet200ResponseBalanceSheetInnerAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetAssets(v GetBalanceSheet200ResponseBalanceSheetInnerAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetLiabilities

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetLiabilities() GetBalanceSheet200ResponseBalanceSheetInnerLiabilities`

GetLiabilities returns the Liabilities field if non-nil, zero value otherwise.

### GetLiabilitiesOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetLiabilitiesOk() (*GetBalanceSheet200ResponseBalanceSheetInnerLiabilities, bool)`

GetLiabilitiesOk returns a tuple with the Liabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilities

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetLiabilities(v GetBalanceSheet200ResponseBalanceSheetInnerLiabilities)`

SetLiabilities sets Liabilities field to given value.

### HasLiabilities

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasLiabilities() bool`

HasLiabilities returns a boolean if a field has been set.

### GetShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetShareholdersEquity() GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity`

GetShareholdersEquity returns the ShareholdersEquity field if non-nil, zero value otherwise.

### GetShareholdersEquityOk

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) GetShareholdersEquityOk() (*GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity, bool)`

GetShareholdersEquityOk returns a tuple with the ShareholdersEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) SetShareholdersEquity(v GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity)`

SetShareholdersEquity sets ShareholdersEquity field to given value.

### HasShareholdersEquity

`func (o *GetBalanceSheet200ResponseBalanceSheetInner) HasShareholdersEquity() bool`

HasShareholdersEquity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


