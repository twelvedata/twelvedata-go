# ApiNotFoundErrorResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** | Error code | 
**Message** | **string** | Error message | 
**Status** | **string** | Error status | 

## Methods

### NewApiNotFoundErrorResponseBody

`func NewApiNotFoundErrorResponseBody(code int64, message string, status string, ) *ApiNotFoundErrorResponseBody`

NewApiNotFoundErrorResponseBody instantiates a new ApiNotFoundErrorResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNotFoundErrorResponseBodyWithDefaults

`func NewApiNotFoundErrorResponseBodyWithDefaults() *ApiNotFoundErrorResponseBody`

NewApiNotFoundErrorResponseBodyWithDefaults instantiates a new ApiNotFoundErrorResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiNotFoundErrorResponseBody) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiNotFoundErrorResponseBody) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiNotFoundErrorResponseBody) SetCode(v int64)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ApiNotFoundErrorResponseBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiNotFoundErrorResponseBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiNotFoundErrorResponseBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *ApiNotFoundErrorResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiNotFoundErrorResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiNotFoundErrorResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


