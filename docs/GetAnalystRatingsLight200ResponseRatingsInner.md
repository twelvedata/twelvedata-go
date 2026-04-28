# GetAnalystRatingsLight200ResponseRatingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date when the rating was released | 
**Firm** | **string** | Firm that issued the ranking | 
**RatingChange** | Pointer to **string** | Defines the action of the firm to ranking, could be &#x60;Maintains&#x60;, &#x60;Upgrade&#x60;, &#x60;Downgrade&#x60;, &#x60;Initiates&#x60; or &#x60;Reiterates&#x60; | [optional] 
**RatingCurrent** | Pointer to **string** | Current firm&#39;s ranking of the instrument | [optional] 
**RatingPrior** | Pointer to **string** | Prior firm&#39;s ranking of the instrument | [optional] 

## Methods

### NewGetAnalystRatingsLight200ResponseRatingsInner

`func NewGetAnalystRatingsLight200ResponseRatingsInner(date string, firm string, ) *GetAnalystRatingsLight200ResponseRatingsInner`

NewGetAnalystRatingsLight200ResponseRatingsInner instantiates a new GetAnalystRatingsLight200ResponseRatingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalystRatingsLight200ResponseRatingsInnerWithDefaults

`func NewGetAnalystRatingsLight200ResponseRatingsInnerWithDefaults() *GetAnalystRatingsLight200ResponseRatingsInner`

NewGetAnalystRatingsLight200ResponseRatingsInnerWithDefaults instantiates a new GetAnalystRatingsLight200ResponseRatingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetFirm

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetFirm() string`

GetFirm returns the Firm field if non-nil, zero value otherwise.

### GetFirmOk

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetFirmOk() (*string, bool)`

GetFirmOk returns a tuple with the Firm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirm

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetFirm(v string)`

SetFirm sets Firm field to given value.


### GetRatingChange

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingChange() string`

GetRatingChange returns the RatingChange field if non-nil, zero value otherwise.

### GetRatingChangeOk

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingChangeOk() (*string, bool)`

GetRatingChangeOk returns a tuple with the RatingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingChange

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetRatingChange(v string)`

SetRatingChange sets RatingChange field to given value.

### HasRatingChange

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) HasRatingChange() bool`

HasRatingChange returns a boolean if a field has been set.

### GetRatingCurrent

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingCurrent() string`

GetRatingCurrent returns the RatingCurrent field if non-nil, zero value otherwise.

### GetRatingCurrentOk

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingCurrentOk() (*string, bool)`

GetRatingCurrentOk returns a tuple with the RatingCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingCurrent

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetRatingCurrent(v string)`

SetRatingCurrent sets RatingCurrent field to given value.

### HasRatingCurrent

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) HasRatingCurrent() bool`

HasRatingCurrent returns a boolean if a field has been set.

### GetRatingPrior

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingPrior() string`

GetRatingPrior returns the RatingPrior field if non-nil, zero value otherwise.

### GetRatingPriorOk

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) GetRatingPriorOk() (*string, bool)`

GetRatingPriorOk returns a tuple with the RatingPrior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingPrior

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) SetRatingPrior(v string)`

SetRatingPrior sets RatingPrior field to given value.

### HasRatingPrior

`func (o *GetAnalystRatingsLight200ResponseRatingsInner) HasRatingPrior() bool`

HasRatingPrior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


