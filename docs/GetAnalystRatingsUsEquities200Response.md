# GetAnalystRatingsUsEquities200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetAnalystRatingsLight200ResponseMeta**](GetAnalystRatingsLight200ResponseMeta.md) |  | 
**Ratings** | Pointer to [**[]GetAnalystRatingsUsEquities200ResponseRatingsInner**](GetAnalystRatingsUsEquities200ResponseRatingsInner.md) | List of analyst ratings | [optional] 
**Status** | **string** | Response status | 

## Methods

### NewGetAnalystRatingsUsEquities200Response

`func NewGetAnalystRatingsUsEquities200Response(meta GetAnalystRatingsLight200ResponseMeta, status string, ) *GetAnalystRatingsUsEquities200Response`

NewGetAnalystRatingsUsEquities200Response instantiates a new GetAnalystRatingsUsEquities200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalystRatingsUsEquities200ResponseWithDefaults

`func NewGetAnalystRatingsUsEquities200ResponseWithDefaults() *GetAnalystRatingsUsEquities200Response`

NewGetAnalystRatingsUsEquities200ResponseWithDefaults instantiates a new GetAnalystRatingsUsEquities200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetAnalystRatingsUsEquities200Response) GetMeta() GetAnalystRatingsLight200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAnalystRatingsUsEquities200Response) GetMetaOk() (*GetAnalystRatingsLight200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAnalystRatingsUsEquities200Response) SetMeta(v GetAnalystRatingsLight200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetRatings

`func (o *GetAnalystRatingsUsEquities200Response) GetRatings() []GetAnalystRatingsUsEquities200ResponseRatingsInner`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *GetAnalystRatingsUsEquities200Response) GetRatingsOk() (*[]GetAnalystRatingsUsEquities200ResponseRatingsInner, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *GetAnalystRatingsUsEquities200Response) SetRatings(v []GetAnalystRatingsUsEquities200ResponseRatingsInner)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *GetAnalystRatingsUsEquities200Response) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetStatus

`func (o *GetAnalystRatingsUsEquities200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAnalystRatingsUsEquities200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAnalystRatingsUsEquities200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


