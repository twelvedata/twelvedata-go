# GetSymbolSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SymbolSearchResponseItem**](SymbolSearchResponseItem.md) | List of symbols matching the search criteria | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetSymbolSearch200Response

`func NewGetSymbolSearch200Response(data []SymbolSearchResponseItem, status string, ) *GetSymbolSearch200Response`

NewGetSymbolSearch200Response instantiates a new GetSymbolSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSymbolSearch200ResponseWithDefaults

`func NewGetSymbolSearch200ResponseWithDefaults() *GetSymbolSearch200Response`

NewGetSymbolSearch200ResponseWithDefaults instantiates a new GetSymbolSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSymbolSearch200Response) GetData() []SymbolSearchResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSymbolSearch200Response) GetDataOk() (*[]SymbolSearchResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSymbolSearch200Response) SetData(v []SymbolSearchResponseItem)`

SetData sets Data field to given value.


### GetStatus

`func (o *GetSymbolSearch200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSymbolSearch200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSymbolSearch200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


