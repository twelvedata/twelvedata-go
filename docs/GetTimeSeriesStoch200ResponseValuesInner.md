# GetTimeSeriesStoch200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**SlowK** | **string** | slow_k value | 
**SlowD** | **string** | slow_d value | 

## Methods

### NewGetTimeSeriesStoch200ResponseValuesInner

`func NewGetTimeSeriesStoch200ResponseValuesInner(datetime string, slowK string, slowD string, ) *GetTimeSeriesStoch200ResponseValuesInner`

NewGetTimeSeriesStoch200ResponseValuesInner instantiates a new GetTimeSeriesStoch200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStoch200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesStoch200ResponseValuesInnerWithDefaults() *GetTimeSeriesStoch200ResponseValuesInner`

NewGetTimeSeriesStoch200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesStoch200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesStoch200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesStoch200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesStoch200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSlowK

`func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowK() string`

GetSlowK returns the SlowK field if non-nil, zero value otherwise.

### GetSlowKOk

`func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowKOk() (*string, bool)`

GetSlowKOk returns a tuple with the SlowK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowK

`func (o *GetTimeSeriesStoch200ResponseValuesInner) SetSlowK(v string)`

SetSlowK sets SlowK field to given value.


### GetSlowD

`func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowD() string`

GetSlowD returns the SlowD field if non-nil, zero value otherwise.

### GetSlowDOk

`func (o *GetTimeSeriesStoch200ResponseValuesInner) GetSlowDOk() (*string, bool)`

GetSlowDOk returns a tuple with the SlowD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowD

`func (o *GetTimeSeriesStoch200ResponseValuesInner) SetSlowD(v string)`

SetSlowD sets SlowD field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


