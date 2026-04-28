# GetETFsWorldPerformance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etf** | [**GetETFsWorldPerformance200ResponseEtf**](GetETFsWorldPerformance200ResponseEtf.md) |  | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetETFsWorldPerformance200Response

`func NewGetETFsWorldPerformance200Response(etf GetETFsWorldPerformance200ResponseEtf, status string, ) *GetETFsWorldPerformance200Response`

NewGetETFsWorldPerformance200Response instantiates a new GetETFsWorldPerformance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsWorldPerformance200ResponseWithDefaults

`func NewGetETFsWorldPerformance200ResponseWithDefaults() *GetETFsWorldPerformance200Response`

NewGetETFsWorldPerformance200ResponseWithDefaults instantiates a new GetETFsWorldPerformance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtf

`func (o *GetETFsWorldPerformance200Response) GetEtf() GetETFsWorldPerformance200ResponseEtf`

GetEtf returns the Etf field if non-nil, zero value otherwise.

### GetEtfOk

`func (o *GetETFsWorldPerformance200Response) GetEtfOk() (*GetETFsWorldPerformance200ResponseEtf, bool)`

GetEtfOk returns a tuple with the Etf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtf

`func (o *GetETFsWorldPerformance200Response) SetEtf(v GetETFsWorldPerformance200ResponseEtf)`

SetEtf sets Etf field to given value.


### GetStatus

`func (o *GetETFsWorldPerformance200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetETFsWorldPerformance200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetETFsWorldPerformance200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


