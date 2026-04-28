# GetTimeSeriesMinMaxIndex200ResponseMeta

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
**Indicator** | [**GetTimeSeriesMinMaxIndex200ResponseMetaIndicator**](GetTimeSeriesMinMaxIndex200ResponseMetaIndicator.md) |  | 

## Methods

### NewGetTimeSeriesMinMaxIndex200ResponseMeta

`func NewGetTimeSeriesMinMaxIndex200ResponseMeta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string, indicator GetTimeSeriesMinMaxIndex200ResponseMetaIndicator, ) *GetTimeSeriesMinMaxIndex200ResponseMeta`

NewGetTimeSeriesMinMaxIndex200ResponseMeta instantiates a new GetTimeSeriesMinMaxIndex200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinMaxIndex200ResponseMetaWithDefaults

`func NewGetTimeSeriesMinMaxIndex200ResponseMetaWithDefaults() *GetTimeSeriesMinMaxIndex200ResponseMeta`

NewGetTimeSeriesMinMaxIndex200ResponseMetaWithDefaults instantiates a new GetTimeSeriesMinMaxIndex200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetInterval

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetCurrency

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchangeTimezone

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.


### GetExchange

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetType

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetType(v string)`

SetType sets Type field to given value.


### GetIndicator

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetIndicator() GetTimeSeriesMinMaxIndex200ResponseMetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) GetIndicatorOk() (*GetTimeSeriesMinMaxIndex200ResponseMetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *GetTimeSeriesMinMaxIndex200ResponseMeta) SetIndicator(v GetTimeSeriesMinMaxIndex200ResponseMetaIndicator)`

SetIndicator sets Indicator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


