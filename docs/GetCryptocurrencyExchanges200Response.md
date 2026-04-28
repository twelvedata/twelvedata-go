# GetCryptocurrencyExchanges200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CryptocurrencyExchangesResponseItem**](CryptocurrencyExchangesResponseItem.md) | List of cryptocurrency exchanges | 
**Status** | **string** | Response status | 

## Methods

### NewGetCryptocurrencyExchanges200Response

`func NewGetCryptocurrencyExchanges200Response(data []CryptocurrencyExchangesResponseItem, status string, ) *GetCryptocurrencyExchanges200Response`

NewGetCryptocurrencyExchanges200Response instantiates a new GetCryptocurrencyExchanges200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCryptocurrencyExchanges200ResponseWithDefaults

`func NewGetCryptocurrencyExchanges200ResponseWithDefaults() *GetCryptocurrencyExchanges200Response`

NewGetCryptocurrencyExchanges200ResponseWithDefaults instantiates a new GetCryptocurrencyExchanges200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCryptocurrencyExchanges200Response) GetData() []CryptocurrencyExchangesResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCryptocurrencyExchanges200Response) GetDataOk() (*[]CryptocurrencyExchangesResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCryptocurrencyExchanges200Response) SetData(v []CryptocurrencyExchangesResponseItem)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetCryptocurrencyExchanges200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCryptocurrencyExchanges200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCryptocurrencyExchanges200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


