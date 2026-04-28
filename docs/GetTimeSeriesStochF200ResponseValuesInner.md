# GetTimeSeriesStochF200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**FastK** | **string** | fast_k value | 
**FastD** | **string** | fast_d value | 

## Methods

### NewGetTimeSeriesStochF200ResponseValuesInner

`func NewGetTimeSeriesStochF200ResponseValuesInner(datetime string, fastK string, fastD string, ) *GetTimeSeriesStochF200ResponseValuesInner`

NewGetTimeSeriesStochF200ResponseValuesInner instantiates a new GetTimeSeriesStochF200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStochF200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesStochF200ResponseValuesInnerWithDefaults() *GetTimeSeriesStochF200ResponseValuesInner`

NewGetTimeSeriesStochF200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesStochF200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesStochF200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesStochF200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesStochF200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetFastK

`func (o *GetTimeSeriesStochF200ResponseValuesInner) GetFastK() string`

GetFastK returns the FastK field if non-nil, zero value otherwise.

### GetFastKOk

`func (o *GetTimeSeriesStochF200ResponseValuesInner) GetFastKOk() (*string, bool)`

GetFastKOk returns a tuple with the FastK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastK

`func (o *GetTimeSeriesStochF200ResponseValuesInner) SetFastK(v string)`

SetFastK sets FastK field to given value.


### GetFastD

`func (o *GetTimeSeriesStochF200ResponseValuesInner) GetFastD() string`

GetFastD returns the FastD field if non-nil, zero value otherwise.

### GetFastDOk

`func (o *GetTimeSeriesStochF200ResponseValuesInner) GetFastDOk() (*string, bool)`

GetFastDOk returns a tuple with the FastD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastD

`func (o *GetTimeSeriesStochF200ResponseValuesInner) SetFastD(v string)`

SetFastD sets FastD field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


