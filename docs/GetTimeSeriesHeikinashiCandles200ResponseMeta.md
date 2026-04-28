# GetTimeSeriesHeikinashiCandles200ResponseMeta

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
**Indicator** | [**GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator**](GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator.md) |  | 

## Methods

### NewGetTimeSeriesHeikinashiCandles200ResponseMeta

`func NewGetTimeSeriesHeikinashiCandles200ResponseMeta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string, indicator GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator, ) *GetTimeSeriesHeikinashiCandles200ResponseMeta`

NewGetTimeSeriesHeikinashiCandles200ResponseMeta instantiates a new GetTimeSeriesHeikinashiCandles200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHeikinashiCandles200ResponseMetaWithDefaults

`func NewGetTimeSeriesHeikinashiCandles200ResponseMetaWithDefaults() *GetTimeSeriesHeikinashiCandles200ResponseMeta`

NewGetTimeSeriesHeikinashiCandles200ResponseMetaWithDefaults instantiates a new GetTimeSeriesHeikinashiCandles200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetInterval

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetCurrency

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchangeTimezone

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.


### GetExchange

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetType

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetType(v string)`

SetType sets Type field to given value.


### GetIndicator

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetIndicator() GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *GetTimeSeriesHeikinashiCandles200ResponseMeta) SetIndicator(v GetTimeSeriesHeikinashiCandles200ResponseMetaIndicator)`

SetIndicator sets Indicator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


