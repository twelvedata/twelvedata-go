# GetEpsTrend200ResponseEpsTrendInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date of the estimation | 
**Period** | **string** | Period of estimation, can be &#x60;current_quarter&#x60;, &#x60;next_quarter&#x60;, &#x60;current_year&#x60;, or &#x60;next_year&#x60; | 
**CurrentEstimate** | Pointer to **float64** | Actual EPS estimation for the period | [optional] 
**Var7DaysAgo** | Pointer to **float64** | EPS estimation value 7 days ago | [optional] 
**Var30DaysAgo** | Pointer to **float64** | EPS estimation value 30 days ago | [optional] 
**Var60DaysAgo** | Pointer to **float64** | EPS estimation value 60 days ago | [optional] 
**Var90DaysAgo** | Pointer to **float64** | EPS estimation value 90 days ago | [optional] 

## Methods

### NewGetEpsTrend200ResponseEpsTrendInner

`func NewGetEpsTrend200ResponseEpsTrendInner(date string, period string, ) *GetEpsTrend200ResponseEpsTrendInner`

NewGetEpsTrend200ResponseEpsTrendInner instantiates a new GetEpsTrend200ResponseEpsTrendInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEpsTrend200ResponseEpsTrendInnerWithDefaults

`func NewGetEpsTrend200ResponseEpsTrendInnerWithDefaults() *GetEpsTrend200ResponseEpsTrendInner`

NewGetEpsTrend200ResponseEpsTrendInnerWithDefaults instantiates a new GetEpsTrend200ResponseEpsTrendInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetEpsTrend200ResponseEpsTrendInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetPeriod

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetEpsTrend200ResponseEpsTrendInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetCurrentEstimate

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetCurrentEstimate() float64`

GetCurrentEstimate returns the CurrentEstimate field if non-nil, zero value otherwise.

### GetCurrentEstimateOk

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetCurrentEstimateOk() (*float64, bool)`

GetCurrentEstimateOk returns a tuple with the CurrentEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEstimate

`func (o *GetEpsTrend200ResponseEpsTrendInner) SetCurrentEstimate(v float64)`

SetCurrentEstimate sets CurrentEstimate field to given value.

### HasCurrentEstimate

`func (o *GetEpsTrend200ResponseEpsTrendInner) HasCurrentEstimate() bool`

HasCurrentEstimate returns a boolean if a field has been set.

### GetVar7DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar7DaysAgo() float64`

GetVar7DaysAgo returns the Var7DaysAgo field if non-nil, zero value otherwise.

### GetVar7DaysAgoOk

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar7DaysAgoOk() (*float64, bool)`

GetVar7DaysAgoOk returns a tuple with the Var7DaysAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar7DaysAgo(v float64)`

SetVar7DaysAgo sets Var7DaysAgo field to given value.

### HasVar7DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar7DaysAgo() bool`

HasVar7DaysAgo returns a boolean if a field has been set.

### GetVar30DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar30DaysAgo() float64`

GetVar30DaysAgo returns the Var30DaysAgo field if non-nil, zero value otherwise.

### GetVar30DaysAgoOk

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar30DaysAgoOk() (*float64, bool)`

GetVar30DaysAgoOk returns a tuple with the Var30DaysAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar30DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar30DaysAgo(v float64)`

SetVar30DaysAgo sets Var30DaysAgo field to given value.

### HasVar30DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar30DaysAgo() bool`

HasVar30DaysAgo returns a boolean if a field has been set.

### GetVar60DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar60DaysAgo() float64`

GetVar60DaysAgo returns the Var60DaysAgo field if non-nil, zero value otherwise.

### GetVar60DaysAgoOk

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar60DaysAgoOk() (*float64, bool)`

GetVar60DaysAgoOk returns a tuple with the Var60DaysAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar60DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar60DaysAgo(v float64)`

SetVar60DaysAgo sets Var60DaysAgo field to given value.

### HasVar60DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar60DaysAgo() bool`

HasVar60DaysAgo returns a boolean if a field has been set.

### GetVar90DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar90DaysAgo() float64`

GetVar90DaysAgo returns the Var90DaysAgo field if non-nil, zero value otherwise.

### GetVar90DaysAgoOk

`func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar90DaysAgoOk() (*float64, bool)`

GetVar90DaysAgoOk returns a tuple with the Var90DaysAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar90DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar90DaysAgo(v float64)`

SetVar90DaysAgo sets Var90DaysAgo field to given value.

### HasVar90DaysAgo

`func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar90DaysAgo() bool`

HasVar90DaysAgo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


