# GetCurrencyConversion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Requested currency symbol | 
**Rate** | **float64** | Real-time exchange rate for the corresponding symbol | 
**Amount** | Pointer to **float64** | Amount of converted currency | [optional] 
**Timestamp** | **int64** | Unix timestamp of the rate | 

## Methods

### NewGetCurrencyConversion200Response

`func NewGetCurrencyConversion200Response(symbol string, rate float64, timestamp int64, ) *GetCurrencyConversion200Response`

NewGetCurrencyConversion200Response instantiates a new GetCurrencyConversion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrencyConversion200ResponseWithDefaults

`func NewGetCurrencyConversion200ResponseWithDefaults() *GetCurrencyConversion200Response`

NewGetCurrencyConversion200ResponseWithDefaults instantiates a new GetCurrencyConversion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetCurrencyConversion200Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetCurrencyConversion200Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetCurrencyConversion200Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetRate

`func (o *GetCurrencyConversion200Response) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetCurrencyConversion200Response) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetCurrencyConversion200Response) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetAmount

`func (o *GetCurrencyConversion200Response) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetCurrencyConversion200Response) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetCurrencyConversion200Response) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetCurrencyConversion200Response) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetCurrencyConversion200Response) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetCurrencyConversion200Response) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetCurrencyConversion200Response) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


