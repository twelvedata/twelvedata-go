# GetExchangeRate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Requested currency symbol | 
**Rate** | **float64** | Real-time exchange rate for the corresponding symbol | 
**Timestamp** | **int64** | Unix timestamp of the rate | 

## Methods

### NewGetExchangeRate200Response

`func NewGetExchangeRate200Response(symbol string, rate float64, timestamp int64, ) *GetExchangeRate200Response`

NewGetExchangeRate200Response instantiates a new GetExchangeRate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeRate200ResponseWithDefaults

`func NewGetExchangeRate200ResponseWithDefaults() *GetExchangeRate200Response`

NewGetExchangeRate200ResponseWithDefaults instantiates a new GetExchangeRate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetExchangeRate200Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetExchangeRate200Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetExchangeRate200Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetRate

`func (o *GetExchangeRate200Response) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetExchangeRate200Response) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetExchangeRate200Response) SetRate(v float64)`

SetRate sets Rate field to given value.


### GetTimestamp

`func (o *GetExchangeRate200Response) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetExchangeRate200Response) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetExchangeRate200Response) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


