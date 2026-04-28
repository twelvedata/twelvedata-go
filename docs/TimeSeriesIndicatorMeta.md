# TimeSeriesIndicatorMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | The ticker symbol of an instrument for which data was requested. | [optional] 
**Interval** | Pointer to **string** | The time gap between consecutive data points. | [optional] 
**Currency** | Pointer to **string** | The currency of a traded instrument. | [optional] 
**ExchangeTimezone** | Pointer to **string** | The timezone of the exchange where the instrument is traded. | [optional] 
**Exchange** | Pointer to **string** | The exchange name where the instrument is traded. | [optional] 
**MicCode** | Pointer to **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded. | [optional] 
**Type** | Pointer to **string** | The asset class to which the instrument belongs. | [optional] 

## Methods

### NewTimeSeriesIndicatorMeta

`func NewTimeSeriesIndicatorMeta() *TimeSeriesIndicatorMeta`

NewTimeSeriesIndicatorMeta instantiates a new TimeSeriesIndicatorMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesIndicatorMetaWithDefaults

`func NewTimeSeriesIndicatorMetaWithDefaults() *TimeSeriesIndicatorMeta`

NewTimeSeriesIndicatorMetaWithDefaults instantiates a new TimeSeriesIndicatorMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *TimeSeriesIndicatorMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TimeSeriesIndicatorMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TimeSeriesIndicatorMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TimeSeriesIndicatorMeta) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetInterval

`func (o *TimeSeriesIndicatorMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *TimeSeriesIndicatorMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *TimeSeriesIndicatorMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *TimeSeriesIndicatorMeta) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetCurrency

`func (o *TimeSeriesIndicatorMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TimeSeriesIndicatorMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TimeSeriesIndicatorMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TimeSeriesIndicatorMeta) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeTimezone

`func (o *TimeSeriesIndicatorMeta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *TimeSeriesIndicatorMeta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *TimeSeriesIndicatorMeta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.

### HasExchangeTimezone

`func (o *TimeSeriesIndicatorMeta) HasExchangeTimezone() bool`

HasExchangeTimezone returns a boolean if a field has been set.

### GetExchange

`func (o *TimeSeriesIndicatorMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *TimeSeriesIndicatorMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *TimeSeriesIndicatorMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *TimeSeriesIndicatorMeta) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMicCode

`func (o *TimeSeriesIndicatorMeta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *TimeSeriesIndicatorMeta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *TimeSeriesIndicatorMeta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *TimeSeriesIndicatorMeta) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetType

`func (o *TimeSeriesIndicatorMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeSeriesIndicatorMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeSeriesIndicatorMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimeSeriesIndicatorMeta) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


