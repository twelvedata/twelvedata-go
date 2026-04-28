# GetTimeSeriesSar200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The ticker symbol of an instrument for which data was requested. | 
**Interval** | **string** | The time gap between consecutive data points. | 
**Currency** | **string** | The currency of a traded instrument. | 
**ExchangeTimezone** | **string** | The timezone of the exchange where the instrument is traded. | 
**Exchange** | **string** | The exchange name where the instrument is traded. | 
**MicCode** | **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded. | 
**Type** | **string** | The asset class to which the instrument belongs. | 
**Indicator** | [**GetTimeSeriesSar200ResponseMetaIndicator**](GetTimeSeriesSar200ResponseMetaIndicator.md) |  | 

## Methods

### NewGetTimeSeriesSar200ResponseMeta

`func NewGetTimeSeriesSar200ResponseMeta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string, indicator GetTimeSeriesSar200ResponseMetaIndicator, ) *GetTimeSeriesSar200ResponseMeta`

NewGetTimeSeriesSar200ResponseMeta instantiates a new GetTimeSeriesSar200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSar200ResponseMetaWithDefaults

`func NewGetTimeSeriesSar200ResponseMetaWithDefaults() *GetTimeSeriesSar200ResponseMeta`

NewGetTimeSeriesSar200ResponseMetaWithDefaults instantiates a new GetTimeSeriesSar200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTimeSeriesSar200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTimeSeriesSar200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetInterval

`func (o *GetTimeSeriesSar200ResponseMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetTimeSeriesSar200ResponseMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetCurrency

`func (o *GetTimeSeriesSar200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetTimeSeriesSar200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchangeTimezone

`func (o *GetTimeSeriesSar200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetTimeSeriesSar200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.


### GetExchange

`func (o *GetTimeSeriesSar200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTimeSeriesSar200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetTimeSeriesSar200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTimeSeriesSar200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetType

`func (o *GetTimeSeriesSar200ResponseMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTimeSeriesSar200ResponseMeta) SetType(v string)`

SetType sets Type field to given value.


### GetIndicator

`func (o *GetTimeSeriesSar200ResponseMeta) GetIndicator() GetTimeSeriesSar200ResponseMetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *GetTimeSeriesSar200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesSar200ResponseMetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *GetTimeSeriesSar200ResponseMeta) SetIndicator(v GetTimeSeriesSar200ResponseMetaIndicator)`

SetIndicator sets Indicator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


