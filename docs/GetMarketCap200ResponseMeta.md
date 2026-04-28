# GetMarketCap200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Ticker of the company | 
**Name** | **string** | Name of the company | 
**Currency** | **string** | Currency in which instrument is traded by ISO 4217 standard | 
**Exchange** | **string** | Exchange name where the company is listed | 
**MicCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
**ExchangeTimezone** | **string** | Exchange timezone | 

## Methods

### NewGetMarketCap200ResponseMeta

`func NewGetMarketCap200ResponseMeta(symbol string, name string, currency string, exchange string, micCode string, exchangeTimezone string, ) *GetMarketCap200ResponseMeta`

NewGetMarketCap200ResponseMeta instantiates a new GetMarketCap200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketCap200ResponseMetaWithDefaults

`func NewGetMarketCap200ResponseMetaWithDefaults() *GetMarketCap200ResponseMeta`

NewGetMarketCap200ResponseMetaWithDefaults instantiates a new GetMarketCap200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetMarketCap200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetMarketCap200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetMarketCap200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *GetMarketCap200ResponseMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMarketCap200ResponseMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMarketCap200ResponseMeta) SetName(v string)`

SetName sets Name field to given value.


### GetCurrency

`func (o *GetMarketCap200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetMarketCap200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetMarketCap200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchange

`func (o *GetMarketCap200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetMarketCap200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetMarketCap200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetMarketCap200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetMarketCap200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetMarketCap200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetExchangeTimezone

`func (o *GetMarketCap200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetMarketCap200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetMarketCap200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


