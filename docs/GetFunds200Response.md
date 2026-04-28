# GetFunds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**GetFunds200ResponseResult**](GetFunds200ResponseResult.md) |  | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetFunds200Response

`func NewGetFunds200Response(result GetFunds200ResponseResult, status string, ) *GetFunds200Response`

NewGetFunds200Response instantiates a new GetFunds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFunds200ResponseWithDefaults

`func NewGetFunds200ResponseWithDefaults() *GetFunds200Response`

NewGetFunds200ResponseWithDefaults instantiates a new GetFunds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetFunds200Response) GetResult() GetFunds200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFunds200Response) GetResultOk() (*GetFunds200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFunds200Response) SetResult(v GetFunds200ResponseResult)`

SetResult sets Result field to given value.


### GetStatus

`func (o *GetFunds200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFunds200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFunds200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


