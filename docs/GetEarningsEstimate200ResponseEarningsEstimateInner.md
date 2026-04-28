# GetEarningsEstimate200ResponseEarningsEstimateInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date of the earnings estimate | 
**Period** | **string** | Period of estimation, can be &#x60;current_quarter&#x60;, &#x60;next_quarter&#x60;, &#x60;current_year&#x60;, or &#x60;next_year&#x60; | 
**NumberOfAnalysts** | Pointer to **int64** | Number of analysts that made the estimation | [optional] 
**AvgEstimate** | Pointer to **float64** | Average estimation across analysts | [optional] 
**LowEstimate** | Pointer to **float64** | Lowest estimation given by an analyst | [optional] 
**HighEstimate** | Pointer to **float64** | Highest estimation given by an analyst | [optional] 
**YearAgoEps** | Pointer to **float64** | Average estimation of this period&#39;s earnings given a year ago | [optional] 

## Methods

### NewGetEarningsEstimate200ResponseEarningsEstimateInner

`func NewGetEarningsEstimate200ResponseEarningsEstimateInner(date string, period string, ) *GetEarningsEstimate200ResponseEarningsEstimateInner`

NewGetEarningsEstimate200ResponseEarningsEstimateInner instantiates a new GetEarningsEstimate200ResponseEarningsEstimateInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEarningsEstimate200ResponseEarningsEstimateInnerWithDefaults

`func NewGetEarningsEstimate200ResponseEarningsEstimateInnerWithDefaults() *GetEarningsEstimate200ResponseEarningsEstimateInner`

NewGetEarningsEstimate200ResponseEarningsEstimateInnerWithDefaults instantiates a new GetEarningsEstimate200ResponseEarningsEstimateInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetPeriod

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetNumberOfAnalysts

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetNumberOfAnalysts() int64`

GetNumberOfAnalysts returns the NumberOfAnalysts field if non-nil, zero value otherwise.

### GetNumberOfAnalystsOk

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetNumberOfAnalystsOk() (*int64, bool)`

GetNumberOfAnalystsOk returns a tuple with the NumberOfAnalysts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAnalysts

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetNumberOfAnalysts(v int64)`

SetNumberOfAnalysts sets NumberOfAnalysts field to given value.

### HasNumberOfAnalysts

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasNumberOfAnalysts() bool`

HasNumberOfAnalysts returns a boolean if a field has been set.

### GetAvgEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetAvgEstimate() float64`

GetAvgEstimate returns the AvgEstimate field if non-nil, zero value otherwise.

### GetAvgEstimateOk

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetAvgEstimateOk() (*float64, bool)`

GetAvgEstimateOk returns a tuple with the AvgEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetAvgEstimate(v float64)`

SetAvgEstimate sets AvgEstimate field to given value.

### HasAvgEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasAvgEstimate() bool`

HasAvgEstimate returns a boolean if a field has been set.

### GetLowEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetLowEstimate() float64`

GetLowEstimate returns the LowEstimate field if non-nil, zero value otherwise.

### GetLowEstimateOk

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetLowEstimateOk() (*float64, bool)`

GetLowEstimateOk returns a tuple with the LowEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetLowEstimate(v float64)`

SetLowEstimate sets LowEstimate field to given value.

### HasLowEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasLowEstimate() bool`

HasLowEstimate returns a boolean if a field has been set.

### GetHighEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetHighEstimate() float64`

GetHighEstimate returns the HighEstimate field if non-nil, zero value otherwise.

### GetHighEstimateOk

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetHighEstimateOk() (*float64, bool)`

GetHighEstimateOk returns a tuple with the HighEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetHighEstimate(v float64)`

SetHighEstimate sets HighEstimate field to given value.

### HasHighEstimate

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasHighEstimate() bool`

HasHighEstimate returns a boolean if a field has been set.

### GetYearAgoEps

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetYearAgoEps() float64`

GetYearAgoEps returns the YearAgoEps field if non-nil, zero value otherwise.

### GetYearAgoEpsOk

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetYearAgoEpsOk() (*float64, bool)`

GetYearAgoEpsOk returns a tuple with the YearAgoEps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearAgoEps

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetYearAgoEps(v float64)`

SetYearAgoEps sets YearAgoEps field to given value.

### HasYearAgoEps

`func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasYearAgoEps() bool`

HasYearAgoEps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


