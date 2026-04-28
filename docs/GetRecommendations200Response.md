# GetRecommendations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetRecommendations200ResponseMeta**](GetRecommendations200ResponseMeta.md) |  | 
**Trends** | [**GetRecommendations200ResponseTrends**](GetRecommendations200ResponseTrends.md) |  | 
**Rating** | Pointer to **float64** | Rating from 0 to 10 represents overall analysts&#39; recommendation. 0 to 2 - strong sell, 2 to 4 - sell, 4 to 6 - hold, 6 to 8 - buy, 8 to 10 - strong buy. | [optional] 
**Status** | **string** | Response status | 

## Methods

### NewGetRecommendations200Response

`func NewGetRecommendations200Response(meta GetRecommendations200ResponseMeta, trends GetRecommendations200ResponseTrends, status string, ) *GetRecommendations200Response`

NewGetRecommendations200Response instantiates a new GetRecommendations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecommendations200ResponseWithDefaults

`func NewGetRecommendations200ResponseWithDefaults() *GetRecommendations200Response`

NewGetRecommendations200ResponseWithDefaults instantiates a new GetRecommendations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetRecommendations200Response) GetMeta() GetRecommendations200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRecommendations200Response) GetMetaOk() (*GetRecommendations200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRecommendations200Response) SetMeta(v GetRecommendations200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetTrends

`func (o *GetRecommendations200Response) GetTrends() GetRecommendations200ResponseTrends`

GetTrends returns the Trends field if non-nil, zero value otherwise.

### GetTrendsOk

`func (o *GetRecommendations200Response) GetTrendsOk() (*GetRecommendations200ResponseTrends, bool)`

GetTrendsOk returns a tuple with the Trends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrends

`func (o *GetRecommendations200Response) SetTrends(v GetRecommendations200ResponseTrends)`

SetTrends sets Trends field to given value.


### GetRating

`func (o *GetRecommendations200Response) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GetRecommendations200Response) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GetRecommendations200Response) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GetRecommendations200Response) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetStatus

`func (o *GetRecommendations200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRecommendations200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRecommendations200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


