# GetAnalystRatingsUsEquities200ResponseRatingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date when the rating was released | 
**Firm** | **string** | Firm that issued the ranking | 
**AnalystName** | Pointer to **string** | Name of an analyst | [optional] 
**RatingChange** | Pointer to **string** | Defines the action of the firm to ranking, could be &#x60;Maintains&#x60;, &#x60;Upgrade&#x60;, &#x60;Downgrade&#x60;, &#x60;Initiates&#x60;, &#x60;Reiterates&#x60;, &#x60;Assumes&#x60;, or &#x60;Reinstates&#x60; | [optional] 
**RatingCurrent** | Pointer to **string** | Current firm&#39;s ranking of the instrument | [optional] 
**RatingPrior** | Pointer to **string** | Prior firm&#39;s ranking of the instrument | [optional] 
**Time** | Pointer to **string** | Time when the rating was released or updated | [optional] 
**ActionPriceTarget** | Pointer to **string** | Action that firm took towards target price | [optional] 
**PriceTargetCurrent** | Pointer to **float64** | Current firm&#39;s price target for the instrument | [optional] 
**PriceTargetPrior** | Pointer to **float64** | Prior firm&#39;s price target for the instrument | [optional] 

## Methods

### NewGetAnalystRatingsUsEquities200ResponseRatingsInner

`func NewGetAnalystRatingsUsEquities200ResponseRatingsInner(date string, firm string, ) *GetAnalystRatingsUsEquities200ResponseRatingsInner`

NewGetAnalystRatingsUsEquities200ResponseRatingsInner instantiates a new GetAnalystRatingsUsEquities200ResponseRatingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalystRatingsUsEquities200ResponseRatingsInnerWithDefaults

`func NewGetAnalystRatingsUsEquities200ResponseRatingsInnerWithDefaults() *GetAnalystRatingsUsEquities200ResponseRatingsInner`

NewGetAnalystRatingsUsEquities200ResponseRatingsInnerWithDefaults instantiates a new GetAnalystRatingsUsEquities200ResponseRatingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetFirm

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetFirm() string`

GetFirm returns the Firm field if non-nil, zero value otherwise.

### GetFirmOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetFirmOk() (*string, bool)`

GetFirmOk returns a tuple with the Firm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirm

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetFirm(v string)`

SetFirm sets Firm field to given value.


### GetAnalystName

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetAnalystName() string`

GetAnalystName returns the AnalystName field if non-nil, zero value otherwise.

### GetAnalystNameOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetAnalystNameOk() (*string, bool)`

GetAnalystNameOk returns a tuple with the AnalystName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalystName

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetAnalystName(v string)`

SetAnalystName sets AnalystName field to given value.

### HasAnalystName

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasAnalystName() bool`

HasAnalystName returns a boolean if a field has been set.

### GetRatingChange

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingChange() string`

GetRatingChange returns the RatingChange field if non-nil, zero value otherwise.

### GetRatingChangeOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingChangeOk() (*string, bool)`

GetRatingChangeOk returns a tuple with the RatingChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingChange

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetRatingChange(v string)`

SetRatingChange sets RatingChange field to given value.

### HasRatingChange

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasRatingChange() bool`

HasRatingChange returns a boolean if a field has been set.

### GetRatingCurrent

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingCurrent() string`

GetRatingCurrent returns the RatingCurrent field if non-nil, zero value otherwise.

### GetRatingCurrentOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingCurrentOk() (*string, bool)`

GetRatingCurrentOk returns a tuple with the RatingCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingCurrent

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetRatingCurrent(v string)`

SetRatingCurrent sets RatingCurrent field to given value.

### HasRatingCurrent

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasRatingCurrent() bool`

HasRatingCurrent returns a boolean if a field has been set.

### GetRatingPrior

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingPrior() string`

GetRatingPrior returns the RatingPrior field if non-nil, zero value otherwise.

### GetRatingPriorOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetRatingPriorOk() (*string, bool)`

GetRatingPriorOk returns a tuple with the RatingPrior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingPrior

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetRatingPrior(v string)`

SetRatingPrior sets RatingPrior field to given value.

### HasRatingPrior

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasRatingPrior() bool`

HasRatingPrior returns a boolean if a field has been set.

### GetTime

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetActionPriceTarget

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetActionPriceTarget() string`

GetActionPriceTarget returns the ActionPriceTarget field if non-nil, zero value otherwise.

### GetActionPriceTargetOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetActionPriceTargetOk() (*string, bool)`

GetActionPriceTargetOk returns a tuple with the ActionPriceTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPriceTarget

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetActionPriceTarget(v string)`

SetActionPriceTarget sets ActionPriceTarget field to given value.

### HasActionPriceTarget

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasActionPriceTarget() bool`

HasActionPriceTarget returns a boolean if a field has been set.

### GetPriceTargetCurrent

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetCurrent() float64`

GetPriceTargetCurrent returns the PriceTargetCurrent field if non-nil, zero value otherwise.

### GetPriceTargetCurrentOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetCurrentOk() (*float64, bool)`

GetPriceTargetCurrentOk returns a tuple with the PriceTargetCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTargetCurrent

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetPriceTargetCurrent(v float64)`

SetPriceTargetCurrent sets PriceTargetCurrent field to given value.

### HasPriceTargetCurrent

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasPriceTargetCurrent() bool`

HasPriceTargetCurrent returns a boolean if a field has been set.

### GetPriceTargetPrior

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetPrior() float64`

GetPriceTargetPrior returns the PriceTargetPrior field if non-nil, zero value otherwise.

### GetPriceTargetPriorOk

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) GetPriceTargetPriorOk() (*float64, bool)`

GetPriceTargetPriorOk returns a tuple with the PriceTargetPrior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTargetPrior

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) SetPriceTargetPrior(v float64)`

SetPriceTargetPrior sets PriceTargetPrior field to given value.

### HasPriceTargetPrior

`func (o *GetAnalystRatingsUsEquities200ResponseRatingsInner) HasPriceTargetPrior() bool`

HasPriceTargetPrior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


