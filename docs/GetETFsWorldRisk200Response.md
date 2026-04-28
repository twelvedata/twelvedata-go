# GetETFsWorldRisk200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etf** | [**GetETFsWorldRisk200ResponseEtf**](GetETFsWorldRisk200ResponseEtf.md) |  | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetETFsWorldRisk200Response

`func NewGetETFsWorldRisk200Response(etf GetETFsWorldRisk200ResponseEtf, status string, ) *GetETFsWorldRisk200Response`

NewGetETFsWorldRisk200Response instantiates a new GetETFsWorldRisk200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsWorldRisk200ResponseWithDefaults

`func NewGetETFsWorldRisk200ResponseWithDefaults() *GetETFsWorldRisk200Response`

NewGetETFsWorldRisk200ResponseWithDefaults instantiates a new GetETFsWorldRisk200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtf

`func (o *GetETFsWorldRisk200Response) GetEtf() GetETFsWorldRisk200ResponseEtf`

GetEtf returns the Etf field if non-nil, zero value otherwise.

### GetEtfOk

`func (o *GetETFsWorldRisk200Response) GetEtfOk() (*GetETFsWorldRisk200ResponseEtf, bool)`

GetEtfOk returns a tuple with the Etf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtf

`func (o *GetETFsWorldRisk200Response) SetEtf(v GetETFsWorldRisk200ResponseEtf)`

SetEtf sets Etf field to given value.


### GetStatus

`func (o *GetETFsWorldRisk200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetETFsWorldRisk200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetETFsWorldRisk200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


