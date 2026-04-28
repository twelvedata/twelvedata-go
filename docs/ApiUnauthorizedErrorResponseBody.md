# ApiUnauthorizedErrorResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** | Error code | 
**Message** | **string** | Error message | 
**Status** | **string** | Error status | 

## Methods

### NewApiUnauthorizedErrorResponseBody

`func NewApiUnauthorizedErrorResponseBody(code int64, message string, status string, ) *ApiUnauthorizedErrorResponseBody`

NewApiUnauthorizedErrorResponseBody instantiates a new ApiUnauthorizedErrorResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUnauthorizedErrorResponseBodyWithDefaults

`func NewApiUnauthorizedErrorResponseBodyWithDefaults() *ApiUnauthorizedErrorResponseBody`

NewApiUnauthorizedErrorResponseBodyWithDefaults instantiates a new ApiUnauthorizedErrorResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiUnauthorizedErrorResponseBody) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiUnauthorizedErrorResponseBody) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiUnauthorizedErrorResponseBody) SetCode(v int64)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ApiUnauthorizedErrorResponseBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiUnauthorizedErrorResponseBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiUnauthorizedErrorResponseBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *ApiUnauthorizedErrorResponseBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiUnauthorizedErrorResponseBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiUnauthorizedErrorResponseBody) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


