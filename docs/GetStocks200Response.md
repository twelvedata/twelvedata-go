# GetStocks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]StocksResponseItem**](StocksResponseItem.md) | List of stock instruments | 
**Status** | **string** | Response status | 

## Methods

### NewGetStocks200Response

`func NewGetStocks200Response(data []StocksResponseItem, status string, ) *GetStocks200Response`

NewGetStocks200Response instantiates a new GetStocks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStocks200ResponseWithDefaults

`func NewGetStocks200ResponseWithDefaults() *GetStocks200Response`

NewGetStocks200ResponseWithDefaults instantiates a new GetStocks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetStocks200Response) GetData() []StocksResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetStocks200Response) GetDataOk() (*[]StocksResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetStocks200Response) SetData(v []StocksResponseItem)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetStocks200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStocks200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStocks200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


