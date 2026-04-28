# EarningsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date of earning release | 
**Time** | Pointer to **string** | Time of earning release, can be either of the following values: &#x60;Pre Market&#x60;, &#x60;After Hours&#x60;, &#x60;Time Not Supplied&#x60; | [optional] 
**EpsEstimate** | Pointer to **float64** | Analyst estimate of the future company earning | [optional] 
**EpsActual** | Pointer to **float64** | Actual value of reported earning | [optional] 
**Difference** | Pointer to **float64** | Delta between &#x60;eps_actual&#x60; and &#x60;eps_estimate&#x60; | [optional] 
**SurprisePrc** | Pointer to **float64** | Surprise in the percentage of the &#x60;eps_actual&#x60; related to &#x60;eps_estimate&#x60; | [optional] 

## Methods

### NewEarningsItem

`func NewEarningsItem(date string, ) *EarningsItem`

NewEarningsItem instantiates a new EarningsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningsItemWithDefaults

`func NewEarningsItemWithDefaults() *EarningsItem`

NewEarningsItemWithDefaults instantiates a new EarningsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EarningsItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EarningsItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EarningsItem) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *EarningsItem) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EarningsItem) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EarningsItem) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *EarningsItem) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetEpsEstimate

`func (o *EarningsItem) GetEpsEstimate() float64`

GetEpsEstimate returns the EpsEstimate field if non-nil, zero value otherwise.

### GetEpsEstimateOk

`func (o *EarningsItem) GetEpsEstimateOk() (*float64, bool)`

GetEpsEstimateOk returns a tuple with the EpsEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsEstimate

`func (o *EarningsItem) SetEpsEstimate(v float64)`

SetEpsEstimate sets EpsEstimate field to given value.

### HasEpsEstimate

`func (o *EarningsItem) HasEpsEstimate() bool`

HasEpsEstimate returns a boolean if a field has been set.

### GetEpsActual

`func (o *EarningsItem) GetEpsActual() float64`

GetEpsActual returns the EpsActual field if non-nil, zero value otherwise.

### GetEpsActualOk

`func (o *EarningsItem) GetEpsActualOk() (*float64, bool)`

GetEpsActualOk returns a tuple with the EpsActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsActual

`func (o *EarningsItem) SetEpsActual(v float64)`

SetEpsActual sets EpsActual field to given value.

### HasEpsActual

`func (o *EarningsItem) HasEpsActual() bool`

HasEpsActual returns a boolean if a field has been set.

### GetDifference

`func (o *EarningsItem) GetDifference() float64`

GetDifference returns the Difference field if non-nil, zero value otherwise.

### GetDifferenceOk

`func (o *EarningsItem) GetDifferenceOk() (*float64, bool)`

GetDifferenceOk returns a tuple with the Difference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifference

`func (o *EarningsItem) SetDifference(v float64)`

SetDifference sets Difference field to given value.

### HasDifference

`func (o *EarningsItem) HasDifference() bool`

HasDifference returns a boolean if a field has been set.

### GetSurprisePrc

`func (o *EarningsItem) GetSurprisePrc() float64`

GetSurprisePrc returns the SurprisePrc field if non-nil, zero value otherwise.

### GetSurprisePrcOk

`func (o *EarningsItem) GetSurprisePrcOk() (*float64, bool)`

GetSurprisePrcOk returns a tuple with the SurprisePrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurprisePrc

`func (o *EarningsItem) SetSurprisePrc(v float64)`

SetSurprisePrc sets SurprisePrc field to given value.

### HasSurprisePrc

`func (o *EarningsItem) HasSurprisePrc() bool`

HasSurprisePrc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


