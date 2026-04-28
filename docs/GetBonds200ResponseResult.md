# GetBonds200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total number of matching instruments | 
**List** | Pointer to [**[]BondResponseItem**](BondResponseItem.md) |  | [optional] 

## Methods

### NewGetBonds200ResponseResult

`func NewGetBonds200ResponseResult(count int64, ) *GetBonds200ResponseResult`

NewGetBonds200ResponseResult instantiates a new GetBonds200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBonds200ResponseResultWithDefaults

`func NewGetBonds200ResponseResultWithDefaults() *GetBonds200ResponseResult`

NewGetBonds200ResponseResultWithDefaults instantiates a new GetBonds200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetBonds200ResponseResult) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetBonds200ResponseResult) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetBonds200ResponseResult) SetCount(v int64)`

SetCount sets Count field to given value.


### GetList

`func (o *GetBonds200ResponseResult) GetList() []BondResponseItem`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetBonds200ResponseResult) GetListOk() (*[]BondResponseItem, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetBonds200ResponseResult) SetList(v []BondResponseItem)`

SetList sets List field to given value.

### HasList

`func (o *GetBonds200ResponseResult) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


