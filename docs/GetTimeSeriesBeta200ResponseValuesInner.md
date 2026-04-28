# GetTimeSeriesBeta200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Beta** | **string** | Beta value | 

## Methods

### NewGetTimeSeriesBeta200ResponseValuesInner

`func NewGetTimeSeriesBeta200ResponseValuesInner(datetime string, beta string, ) *GetTimeSeriesBeta200ResponseValuesInner`

NewGetTimeSeriesBeta200ResponseValuesInner instantiates a new GetTimeSeriesBeta200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesBeta200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesBeta200ResponseValuesInnerWithDefaults() *GetTimeSeriesBeta200ResponseValuesInner`

NewGetTimeSeriesBeta200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesBeta200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesBeta200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesBeta200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesBeta200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetBeta

`func (o *GetTimeSeriesBeta200ResponseValuesInner) GetBeta() string`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *GetTimeSeriesBeta200ResponseValuesInner) GetBetaOk() (*string, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *GetTimeSeriesBeta200ResponseValuesInner) SetBeta(v string)`

SetBeta sets Beta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


