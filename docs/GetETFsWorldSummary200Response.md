# GetETFsWorldSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etf** | [**GetETFsWorldSummary200ResponseEtf**](GetETFsWorldSummary200ResponseEtf.md) |  | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetETFsWorldSummary200Response

`func NewGetETFsWorldSummary200Response(etf GetETFsWorldSummary200ResponseEtf, status string, ) *GetETFsWorldSummary200Response`

NewGetETFsWorldSummary200Response instantiates a new GetETFsWorldSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetETFsWorldSummary200ResponseWithDefaults

`func NewGetETFsWorldSummary200ResponseWithDefaults() *GetETFsWorldSummary200Response`

NewGetETFsWorldSummary200ResponseWithDefaults instantiates a new GetETFsWorldSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtf

`func (o *GetETFsWorldSummary200Response) GetEtf() GetETFsWorldSummary200ResponseEtf`

GetEtf returns the Etf field if non-nil, zero value otherwise.

### GetEtfOk

`func (o *GetETFsWorldSummary200Response) GetEtfOk() (*GetETFsWorldSummary200ResponseEtf, bool)`

GetEtfOk returns a tuple with the Etf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtf

`func (o *GetETFsWorldSummary200Response) SetEtf(v GetETFsWorldSummary200ResponseEtf)`

SetEtf sets Etf field to given value.


### GetStatus

`func (o *GetETFsWorldSummary200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetETFsWorldSummary200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetETFsWorldSummary200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


