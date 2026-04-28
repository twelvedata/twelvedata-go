# GetBonds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**GetBonds200ResponseResult**](GetBonds200ResponseResult.md) |  | [optional] 
**Status** | **string** | Response status | 

## Methods

### NewGetBonds200Response

`func NewGetBonds200Response(status string, ) *GetBonds200Response`

NewGetBonds200Response instantiates a new GetBonds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBonds200ResponseWithDefaults

`func NewGetBonds200ResponseWithDefaults() *GetBonds200Response`

NewGetBonds200ResponseWithDefaults instantiates a new GetBonds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetBonds200Response) GetResult() GetBonds200ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetBonds200Response) GetResultOk() (*GetBonds200ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetBonds200Response) SetResult(v GetBonds200ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetBonds200Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *GetBonds200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBonds200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBonds200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


