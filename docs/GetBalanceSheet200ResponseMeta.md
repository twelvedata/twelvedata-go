# GetBalanceSheet200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Ticker symbol of instrument | 
**Name** | **string** | Name of the company | 
**Currency** | **string** | Currency of the balance sheet according to the ISO 4217 standard | 
**Exchange** | **string** | Exchange where instrument is traded | 
**MicCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 
**ExchangeTimezone** | **string** | Exchange timezone | 
**Period** | **string** | Values can be annual or quarterly | 

## Methods

### NewGetBalanceSheet200ResponseMeta

`func NewGetBalanceSheet200ResponseMeta(symbol string, name string, currency string, exchange string, micCode string, exchangeTimezone string, period string, ) *GetBalanceSheet200ResponseMeta`

NewGetBalanceSheet200ResponseMeta instantiates a new GetBalanceSheet200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceSheet200ResponseMetaWithDefaults

`func NewGetBalanceSheet200ResponseMetaWithDefaults() *GetBalanceSheet200ResponseMeta`

NewGetBalanceSheet200ResponseMetaWithDefaults instantiates a new GetBalanceSheet200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetBalanceSheet200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetBalanceSheet200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetBalanceSheet200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *GetBalanceSheet200ResponseMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBalanceSheet200ResponseMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBalanceSheet200ResponseMeta) SetName(v string)`

SetName sets Name field to given value.


### GetCurrency

`func (o *GetBalanceSheet200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBalanceSheet200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBalanceSheet200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchange

`func (o *GetBalanceSheet200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetBalanceSheet200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetBalanceSheet200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetBalanceSheet200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetBalanceSheet200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetBalanceSheet200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetExchangeTimezone

`func (o *GetBalanceSheet200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetBalanceSheet200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetBalanceSheet200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.


### GetPeriod

`func (o *GetBalanceSheet200ResponseMeta) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetBalanceSheet200ResponseMeta) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetBalanceSheet200ResponseMeta) SetPeriod(v string)`

SetPeriod sets Period field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


