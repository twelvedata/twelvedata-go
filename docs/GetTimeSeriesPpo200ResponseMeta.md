# GetTimeSeriesPpo200ResponseMeta

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
**Indicator** | [**GetTimeSeriesPpo200ResponseMetaIndicator**](GetTimeSeriesPpo200ResponseMetaIndicator.md) |  | 

## Methods

### NewGetTimeSeriesPpo200ResponseMeta

`func NewGetTimeSeriesPpo200ResponseMeta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string, indicator GetTimeSeriesPpo200ResponseMetaIndicator, ) *GetTimeSeriesPpo200ResponseMeta`

NewGetTimeSeriesPpo200ResponseMeta instantiates a new GetTimeSeriesPpo200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPpo200ResponseMetaWithDefaults

`func NewGetTimeSeriesPpo200ResponseMetaWithDefaults() *GetTimeSeriesPpo200ResponseMeta`

NewGetTimeSeriesPpo200ResponseMetaWithDefaults instantiates a new GetTimeSeriesPpo200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTimeSeriesPpo200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTimeSeriesPpo200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetInterval

`func (o *GetTimeSeriesPpo200ResponseMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetTimeSeriesPpo200ResponseMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetCurrency

`func (o *GetTimeSeriesPpo200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetTimeSeriesPpo200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchangeTimezone

`func (o *GetTimeSeriesPpo200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetTimeSeriesPpo200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.


### GetExchange

`func (o *GetTimeSeriesPpo200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTimeSeriesPpo200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetTimeSeriesPpo200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTimeSeriesPpo200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetType

`func (o *GetTimeSeriesPpo200ResponseMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTimeSeriesPpo200ResponseMeta) SetType(v string)`

SetType sets Type field to given value.


### GetIndicator

`func (o *GetTimeSeriesPpo200ResponseMeta) GetIndicator() GetTimeSeriesPpo200ResponseMetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *GetTimeSeriesPpo200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesPpo200ResponseMetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *GetTimeSeriesPpo200ResponseMeta) SetIndicator(v GetTimeSeriesPpo200ResponseMetaIndicator)`

SetIndicator sets Indicator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


