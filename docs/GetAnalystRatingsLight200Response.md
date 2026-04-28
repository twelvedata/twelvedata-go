# GetAnalystRatingsLight200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetAnalystRatingsLight200ResponseMeta**](GetAnalystRatingsLight200ResponseMeta.md) |  | 
**Ratings** | Pointer to [**[]GetAnalystRatingsLight200ResponseRatingsInner**](GetAnalystRatingsLight200ResponseRatingsInner.md) | List of analyst ratings | [optional] 
**Status** | **string** | Response status | 

## Methods

### NewGetAnalystRatingsLight200Response

`func NewGetAnalystRatingsLight200Response(meta GetAnalystRatingsLight200ResponseMeta, status string, ) *GetAnalystRatingsLight200Response`

NewGetAnalystRatingsLight200Response instantiates a new GetAnalystRatingsLight200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalystRatingsLight200ResponseWithDefaults

`func NewGetAnalystRatingsLight200ResponseWithDefaults() *GetAnalystRatingsLight200Response`

NewGetAnalystRatingsLight200ResponseWithDefaults instantiates a new GetAnalystRatingsLight200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetAnalystRatingsLight200Response) GetMeta() GetAnalystRatingsLight200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAnalystRatingsLight200Response) GetMetaOk() (*GetAnalystRatingsLight200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAnalystRatingsLight200Response) SetMeta(v GetAnalystRatingsLight200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetRatings

`func (o *GetAnalystRatingsLight200Response) GetRatings() []GetAnalystRatingsLight200ResponseRatingsInner`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *GetAnalystRatingsLight200Response) GetRatingsOk() (*[]GetAnalystRatingsLight200ResponseRatingsInner, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *GetAnalystRatingsLight200Response) SetRatings(v []GetAnalystRatingsLight200ResponseRatingsInner)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *GetAnalystRatingsLight200Response) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetStatus

`func (o *GetAnalystRatingsLight200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAnalystRatingsLight200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAnalystRatingsLight200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


