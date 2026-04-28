# GetFunds200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Total number of matching instruments | 
**List** | [**[]FundResponseItem**](FundResponseItem.md) | List of funds | 

## Methods

### NewGetFunds200ResponseResult

`func NewGetFunds200ResponseResult(count int64, list []FundResponseItem, ) *GetFunds200ResponseResult`

NewGetFunds200ResponseResult instantiates a new GetFunds200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFunds200ResponseResultWithDefaults

`func NewGetFunds200ResponseResultWithDefaults() *GetFunds200ResponseResult`

NewGetFunds200ResponseResultWithDefaults instantiates a new GetFunds200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetFunds200ResponseResult) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetFunds200ResponseResult) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetFunds200ResponseResult) SetCount(v int64)`

SetCount sets Count field to given value.


### GetList

`func (o *GetFunds200ResponseResult) GetList() []FundResponseItem`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetFunds200ResponseResult) GetListOk() (*[]FundResponseItem, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetFunds200ResponseResult) SetList(v []FundResponseItem)`

SetList sets List field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


