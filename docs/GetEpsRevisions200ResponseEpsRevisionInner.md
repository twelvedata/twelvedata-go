# GetEpsRevisions200ResponseEpsRevisionInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date of the EPS estimate | 
**Period** | **string** | Period of estimation, can be &#x60;current_quarter&#x60;, &#x60;next_quarter&#x60;, &#x60;current_year&#x60;, or &#x60;next_year&#x60; | 
**UpLastWeek** | Pointer to **int64** | Number of up revisions over the last 7 days | [optional] 
**UpLastMonth** | Pointer to **int64** | Number of up revisions over the last 30 days | [optional] 
**DownLastWeek** | Pointer to **int64** | Number of down revisions over the last 7 days | [optional] 
**DownLastMonth** | Pointer to **int64** | Number of down revisions over the last 30 days | [optional] 

## Methods

### NewGetEpsRevisions200ResponseEpsRevisionInner

`func NewGetEpsRevisions200ResponseEpsRevisionInner(date string, period string, ) *GetEpsRevisions200ResponseEpsRevisionInner`

NewGetEpsRevisions200ResponseEpsRevisionInner instantiates a new GetEpsRevisions200ResponseEpsRevisionInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEpsRevisions200ResponseEpsRevisionInnerWithDefaults

`func NewGetEpsRevisions200ResponseEpsRevisionInnerWithDefaults() *GetEpsRevisions200ResponseEpsRevisionInner`

NewGetEpsRevisions200ResponseEpsRevisionInnerWithDefaults instantiates a new GetEpsRevisions200ResponseEpsRevisionInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetPeriod

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetUpLastWeek

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastWeek() int64`

GetUpLastWeek returns the UpLastWeek field if non-nil, zero value otherwise.

### GetUpLastWeekOk

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastWeekOk() (*int64, bool)`

GetUpLastWeekOk returns a tuple with the UpLastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLastWeek

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetUpLastWeek(v int64)`

SetUpLastWeek sets UpLastWeek field to given value.

### HasUpLastWeek

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasUpLastWeek() bool`

HasUpLastWeek returns a boolean if a field has been set.

### GetUpLastMonth

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastMonth() int64`

GetUpLastMonth returns the UpLastMonth field if non-nil, zero value otherwise.

### GetUpLastMonthOk

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastMonthOk() (*int64, bool)`

GetUpLastMonthOk returns a tuple with the UpLastMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLastMonth

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetUpLastMonth(v int64)`

SetUpLastMonth sets UpLastMonth field to given value.

### HasUpLastMonth

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasUpLastMonth() bool`

HasUpLastMonth returns a boolean if a field has been set.

### GetDownLastWeek

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastWeek() int64`

GetDownLastWeek returns the DownLastWeek field if non-nil, zero value otherwise.

### GetDownLastWeekOk

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastWeekOk() (*int64, bool)`

GetDownLastWeekOk returns a tuple with the DownLastWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLastWeek

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetDownLastWeek(v int64)`

SetDownLastWeek sets DownLastWeek field to given value.

### HasDownLastWeek

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasDownLastWeek() bool`

HasDownLastWeek returns a boolean if a field has been set.

### GetDownLastMonth

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastMonth() int64`

GetDownLastMonth returns the DownLastMonth field if non-nil, zero value otherwise.

### GetDownLastMonthOk

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastMonthOk() (*int64, bool)`

GetDownLastMonthOk returns a tuple with the DownLastMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLastMonth

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetDownLastMonth(v int64)`

SetDownLastMonth sets DownLastMonth field to given value.

### HasDownLastMonth

`func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasDownLastMonth() bool`

HasDownLastMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


