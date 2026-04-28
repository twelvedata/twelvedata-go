# GetETFsWorldComposition200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etf** | [**GetETFsWorldComposition200ResponseEtf**](GetETFsWorldComposition200ResponseEtf.md) |  | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetETFsWorldComposition200Response

`func NewGetETFsWorldComposition200Response(etf GetETFsWorldComposition200ResponseEtf, status string, ) *GetETFsWorldComposition200Response`

NewGetETFsWorldComposition200Response instantiates a new GetETFsWorldComposition200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsWorldComposition200ResponseWithDefaults

`func NewGetETFsWorldComposition200ResponseWithDefaults() *GetETFsWorldComposition200Response`

NewGetETFsWorldComposition200ResponseWithDefaults instantiates a new GetETFsWorldComposition200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtf

`func (o *GetETFsWorldComposition200Response) GetEtf() GetETFsWorldComposition200ResponseEtf`

GetEtf returns the Etf field if non-nil, zero value otherwise.

### GetEtfOk

`func (o *GetETFsWorldComposition200Response) GetEtfOk() (*GetETFsWorldComposition200ResponseEtf, bool)`

GetEtfOk returns a tuple with the Etf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtf

`func (o *GetETFsWorldComposition200Response) SetEtf(v GetETFsWorldComposition200ResponseEtf)`

SetEtf sets Etf field to given value.


### GetStatus

`func (o *GetETFsWorldComposition200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetETFsWorldComposition200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetETFsWorldComposition200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


