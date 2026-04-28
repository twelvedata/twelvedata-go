# StockExchange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewStockExchange

`func NewStockExchange() *StockExchange`

NewStockExchange instantiates a new StockExchange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockExchangeWithDefaults

`func NewStockExchangeWithDefaults() *StockExchange`

NewStockExchangeWithDefaults instantiates a new StockExchange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *StockExchange) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StockExchange) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StockExchange) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StockExchange) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetName

`func (o *StockExchange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StockExchange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StockExchange) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StockExchange) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


