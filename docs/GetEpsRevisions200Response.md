# GetEpsRevisions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetEarningsEstimate200ResponseMeta**](GetEarningsEstimate200ResponseMeta.md) |  | 
**EpsRevision** | [**[]GetEpsRevisions200ResponseEpsRevisionInner**](GetEpsRevisions200ResponseEpsRevisionInner.md) | EPS revision data | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetEpsRevisions200Response

`func NewGetEpsRevisions200Response(meta GetEarningsEstimate200ResponseMeta, epsRevision []GetEpsRevisions200ResponseEpsRevisionInner, status string, ) *GetEpsRevisions200Response`

NewGetEpsRevisions200Response instantiates a new GetEpsRevisions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEpsRevisions200ResponseWithDefaults

`func NewGetEpsRevisions200ResponseWithDefaults() *GetEpsRevisions200Response`

NewGetEpsRevisions200ResponseWithDefaults instantiates a new GetEpsRevisions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetEpsRevisions200Response) GetMeta() GetEarningsEstimate200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetEpsRevisions200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetEpsRevisions200Response) SetMeta(v GetEarningsEstimate200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetEpsRevision

`func (o *GetEpsRevisions200Response) GetEpsRevision() []GetEpsRevisions200ResponseEpsRevisionInner`

GetEpsRevision returns the EpsRevision field if non-nil, zero value otherwise.

### GetEpsRevisionOk

`func (o *GetEpsRevisions200Response) GetEpsRevisionOk() (*[]GetEpsRevisions200ResponseEpsRevisionInner, bool)`

GetEpsRevisionOk returns a tuple with the EpsRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsRevision

`func (o *GetEpsRevisions200Response) SetEpsRevision(v []GetEpsRevisions200ResponseEpsRevisionInner)`

SetEpsRevision sets EpsRevision field to given value.


### GetStatus

`func (o *GetEpsRevisions200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEpsRevisions200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEpsRevisions200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


