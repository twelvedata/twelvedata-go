# GetETFsFamily200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **map[string][]string** | List of ETFs by country | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetETFsFamily200Response

`func NewGetETFsFamily200Response(result map[string][]string, status string, ) *GetETFsFamily200Response`

NewGetETFsFamily200Response instantiates a new GetETFsFamily200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsFamily200ResponseWithDefaults

`func NewGetETFsFamily200ResponseWithDefaults() *GetETFsFamily200Response`

NewGetETFsFamily200ResponseWithDefaults instantiates a new GetETFsFamily200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetETFsFamily200Response) GetResult() map[string][]string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetETFsFamily200Response) GetResultOk() (*map[string][]string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetETFsFamily200Response) SetResult(v map[string][]string)`

SetResult sets Result field to given value.


### GetStatus

`func (o *GetETFsFamily200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetETFsFamily200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetETFsFamily200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


