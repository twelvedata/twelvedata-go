# EdgarFilingValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cik** | **int64** | CIK code | 
**FiledAt** | **int64** | Filing date in UNIX timestamp | 
**Files** | Pointer to [**[]EdgarFilingFile**](EdgarFilingFile.md) | Filing files | [optional] 
**FilingUrl** | **string** | Full URL of the filing | 
**FormType** | **string** | Filing form type | 
**Ticker** | **[]string** | Ticker symbols associated with the filing | 

## Methods

### NewEdgarFilingValue

`func NewEdgarFilingValue(cik int64, filedAt int64, filingUrl string, formType string, ticker []string, ) *EdgarFilingValue`

NewEdgarFilingValue instantiates a new EdgarFilingValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgarFilingValueWithDefaults

`func NewEdgarFilingValueWithDefaults() *EdgarFilingValue`

NewEdgarFilingValueWithDefaults instantiates a new EdgarFilingValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCik

`func (o *EdgarFilingValue) GetCik() int64`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *EdgarFilingValue) GetCikOk() (*int64, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *EdgarFilingValue) SetCik(v int64)`

SetCik sets Cik field to given value.


### GetFiledAt

`func (o *EdgarFilingValue) GetFiledAt() int64`

GetFiledAt returns the FiledAt field if non-nil, zero value otherwise.

### GetFiledAtOk

`func (o *EdgarFilingValue) GetFiledAtOk() (*int64, bool)`

GetFiledAtOk returns a tuple with the FiledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiledAt

`func (o *EdgarFilingValue) SetFiledAt(v int64)`

SetFiledAt sets FiledAt field to given value.


### GetFiles

`func (o *EdgarFilingValue) GetFiles() []EdgarFilingFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *EdgarFilingValue) GetFilesOk() (*[]EdgarFilingFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *EdgarFilingValue) SetFiles(v []EdgarFilingFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *EdgarFilingValue) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFilingUrl

`func (o *EdgarFilingValue) GetFilingUrl() string`

GetFilingUrl returns the FilingUrl field if non-nil, zero value otherwise.

### GetFilingUrlOk

`func (o *EdgarFilingValue) GetFilingUrlOk() (*string, bool)`

GetFilingUrlOk returns a tuple with the FilingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingUrl

`func (o *EdgarFilingValue) SetFilingUrl(v string)`

SetFilingUrl sets FilingUrl field to given value.


### GetFormType

`func (o *EdgarFilingValue) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *EdgarFilingValue) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *EdgarFilingValue) SetFormType(v string)`

SetFormType sets FormType field to given value.


### GetTicker

`func (o *EdgarFilingValue) GetTicker() []string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *EdgarFilingValue) GetTickerOk() (*[]string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *EdgarFilingValue) SetTicker(v []string)`

SetTicker sets Ticker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


