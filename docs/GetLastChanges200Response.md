# GetLastChanges200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**GetLastChanges200ResponsePagination**](GetLastChanges200ResponsePagination.md) |  | 
**Data** | [**[]LastChangeResponseItem**](LastChangeResponseItem.md) | Data contains the list of last changes | 

## Methods

### NewGetLastChanges200Response

`func NewGetLastChanges200Response(pagination GetLastChanges200ResponsePagination, data []LastChangeResponseItem, ) *GetLastChanges200Response`

NewGetLastChanges200Response instantiates a new GetLastChanges200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLastChanges200ResponseWithDefaults

`func NewGetLastChanges200ResponseWithDefaults() *GetLastChanges200Response`

NewGetLastChanges200ResponseWithDefaults instantiates a new GetLastChanges200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetLastChanges200Response) GetPagination() GetLastChanges200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetLastChanges200Response) GetPaginationOk() (*GetLastChanges200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetLastChanges200Response) SetPagination(v GetLastChanges200ResponsePagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *GetLastChanges200Response) GetData() []LastChangeResponseItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLastChanges200Response) GetDataOk() (*[]LastChangeResponseItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLastChanges200Response) SetData(v []LastChangeResponseItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


