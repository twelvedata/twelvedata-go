# GetEod200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol passed | 
**Exchange** | **string** | Exchange where instrument is traded | 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Currency** | Pointer to **string** | Currency in which instrument is denominated | [optional] 
**Datetime** | **string** | Datetime in defined timezone referring to when the bar with specified interval was opened | 
**Close** | **string** | The most recent end of day close price | 

## Methods

### NewGetEod200Response

`func NewGetEod200Response(symbol string, exchange string, datetime string, close string, ) *GetEod200Response`

NewGetEod200Response instantiates a new GetEod200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEod200ResponseWithDefaults

`func NewGetEod200ResponseWithDefaults() *GetEod200Response`

NewGetEod200ResponseWithDefaults instantiates a new GetEod200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetEod200Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetEod200Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetEod200Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetExchange

`func (o *GetEod200Response) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetEod200Response) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetEod200Response) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetEod200Response) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetEod200Response) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetEod200Response) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *GetEod200Response) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetCurrency

`func (o *GetEod200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetEod200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetEod200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetEod200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDatetime

`func (o *GetEod200Response) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetEod200Response) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetEod200Response) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetClose

`func (o *GetEod200Response) GetClose() string`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *GetEod200Response) GetCloseOk() (*string, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *GetEod200Response) SetClose(v string)`

SetClose sets Close field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


