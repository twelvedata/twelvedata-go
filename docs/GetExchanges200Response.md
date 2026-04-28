# GetExchanges200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ExchangesResponseItem**](ExchangesResponseItem.md) | List of exchanges | 
**Status** | **string** | Response status | 

## Methods

### NewGetExchanges200Response

`func NewGetExchanges200Response(data []ExchangesResponseItem, status string, ) *GetExchanges200Response`

NewGetExchanges200Response instantiates a new GetExchanges200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchanges200ResponseWithDefaults

`func NewGetExchanges200ResponseWithDefaults() *GetExchanges200Response`

NewGetExchanges200ResponseWithDefaults instantiates a new GetExchanges200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetExchanges200Response) GetData() []ExchangesResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetExchanges200Response) GetDataOk() (*[]ExchangesResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetExchanges200Response) SetData(v []ExchangesResponseItem)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetExchanges200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetExchanges200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetExchanges200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


