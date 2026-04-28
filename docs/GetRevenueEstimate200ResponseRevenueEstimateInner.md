# GetRevenueEstimate200ResponseRevenueEstimateInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Date of the estimate | [optional] 
**Period** | Pointer to **string** | Period of estimation, can be &#x60;current_quarter&#x60;, &#x60;next_quarter&#x60;, &#x60;current_year&#x60;, or &#x60;next_year&#x60; | [optional] 
**NumberOfAnalysts** | Pointer to **int64** | Number of analysts that made the estimation | [optional] 
**AvgEstimate** | Pointer to **float64** | Average estimation across analysts | [optional] 
**LowEstimate** | Pointer to **float64** | Lowest estimation given by an analyst | [optional] 
**HighEstimate** | Pointer to **float64** | Highest estimation given by an analyst | [optional] 
**YearAgoSales** | Pointer to **float64** | Total revenue received a year ago relative to period | [optional] 
**SalesGrowth** | Pointer to **float64** | Estimated sales growth of the period in relation to year-ago sales in prc (%) | [optional] 

## Methods

### NewGetRevenueEstimate200ResponseRevenueEstimateInner

`func NewGetRevenueEstimate200ResponseRevenueEstimateInner() *GetRevenueEstimate200ResponseRevenueEstimateInner`

NewGetRevenueEstimate200ResponseRevenueEstimateInner instantiates a new GetRevenueEstimate200ResponseRevenueEstimateInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevenueEstimate200ResponseRevenueEstimateInnerWithDefaults

`func NewGetRevenueEstimate200ResponseRevenueEstimateInnerWithDefaults() *GetRevenueEstimate200ResponseRevenueEstimateInner`

NewGetRevenueEstimate200ResponseRevenueEstimateInnerWithDefaults instantiates a new GetRevenueEstimate200ResponseRevenueEstimateInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPeriod

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetNumberOfAnalysts

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetNumberOfAnalysts() int64`

GetNumberOfAnalysts returns the NumberOfAnalysts field if non-nil, zero value otherwise.

### GetNumberOfAnalystsOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetNumberOfAnalystsOk() (*int64, bool)`

GetNumberOfAnalystsOk returns a tuple with the NumberOfAnalysts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAnalysts

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetNumberOfAnalysts(v int64)`

SetNumberOfAnalysts sets NumberOfAnalysts field to given value.

### HasNumberOfAnalysts

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasNumberOfAnalysts() bool`

HasNumberOfAnalysts returns a boolean if a field has been set.

### GetAvgEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetAvgEstimate() float64`

GetAvgEstimate returns the AvgEstimate field if non-nil, zero value otherwise.

### GetAvgEstimateOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetAvgEstimateOk() (*float64, bool)`

GetAvgEstimateOk returns a tuple with the AvgEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetAvgEstimate(v float64)`

SetAvgEstimate sets AvgEstimate field to given value.

### HasAvgEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasAvgEstimate() bool`

HasAvgEstimate returns a boolean if a field has been set.

### GetLowEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetLowEstimate() float64`

GetLowEstimate returns the LowEstimate field if non-nil, zero value otherwise.

### GetLowEstimateOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetLowEstimateOk() (*float64, bool)`

GetLowEstimateOk returns a tuple with the LowEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetLowEstimate(v float64)`

SetLowEstimate sets LowEstimate field to given value.

### HasLowEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasLowEstimate() bool`

HasLowEstimate returns a boolean if a field has been set.

### GetHighEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetHighEstimate() float64`

GetHighEstimate returns the HighEstimate field if non-nil, zero value otherwise.

### GetHighEstimateOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetHighEstimateOk() (*float64, bool)`

GetHighEstimateOk returns a tuple with the HighEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetHighEstimate(v float64)`

SetHighEstimate sets HighEstimate field to given value.

### HasHighEstimate

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasHighEstimate() bool`

HasHighEstimate returns a boolean if a field has been set.

### GetYearAgoSales

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetYearAgoSales() float64`

GetYearAgoSales returns the YearAgoSales field if non-nil, zero value otherwise.

### GetYearAgoSalesOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetYearAgoSalesOk() (*float64, bool)`

GetYearAgoSalesOk returns a tuple with the YearAgoSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearAgoSales

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetYearAgoSales(v float64)`

SetYearAgoSales sets YearAgoSales field to given value.

### HasYearAgoSales

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasYearAgoSales() bool`

HasYearAgoSales returns a boolean if a field has been set.

### GetSalesGrowth

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetSalesGrowth() float64`

GetSalesGrowth returns the SalesGrowth field if non-nil, zero value otherwise.

### GetSalesGrowthOk

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetSalesGrowthOk() (*float64, bool)`

GetSalesGrowthOk returns a tuple with the SalesGrowth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesGrowth

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetSalesGrowth(v float64)`

SetSalesGrowth sets SalesGrowth field to given value.

### HasSalesGrowth

`func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasSalesGrowth() bool`

HasSalesGrowth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


