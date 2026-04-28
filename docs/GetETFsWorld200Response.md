# GetETFsWorld200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etf** | [**GetETFsWorld200ResponseEtf**](GetETFsWorld200ResponseEtf.md) |  | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetETFsWorld200Response

`func NewGetETFsWorld200Response(etf GetETFsWorld200ResponseEtf, status string, ) *GetETFsWorld200Response`

NewGetETFsWorld200Response instantiates a new GetETFsWorld200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsWorld200ResponseWithDefaults

`func NewGetETFsWorld200ResponseWithDefaults() *GetETFsWorld200Response`

NewGetETFsWorld200ResponseWithDefaults instantiates a new GetETFsWorld200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtf

`func (o *GetETFsWorld200Response) GetEtf() GetETFsWorld200ResponseEtf`

GetEtf returns the Etf field if non-nil, zero value otherwise.

### GetEtfOk

`func (o *GetETFsWorld200Response) GetEtfOk() (*GetETFsWorld200ResponseEtf, bool)`

GetEtfOk returns a tuple with the Etf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtf

`func (o *GetETFsWorld200Response) SetEtf(v GetETFsWorld200ResponseEtf)`

SetEtf sets Etf field to given value.


### GetStatus

`func (o *GetETFsWorld200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetETFsWorld200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetETFsWorld200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


