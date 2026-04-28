# Advanced200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** | HTTP status code | [optional] 
**Status** | Pointer to **string** | Status of the request | [optional] 
**Data** | Pointer to **map[string]map[string]interface{}** | Response data containing individual request results | [optional] 

## Methods

### NewAdvanced200Response

`func NewAdvanced200Response() *Advanced200Response`

NewAdvanced200Response instantiates a new Advanced200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvanced200ResponseWithDefaults

`func NewAdvanced200ResponseWithDefaults() *Advanced200Response`

NewAdvanced200ResponseWithDefaults instantiates a new Advanced200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Advanced200Response) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Advanced200Response) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Advanced200Response) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *Advanced200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *Advanced200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Advanced200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Advanced200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Advanced200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *Advanced200Response) GetData() map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Advanced200Response) GetDataOk() (*map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Advanced200Response) SetData(v map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Advanced200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


