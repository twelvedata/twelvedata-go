# GetTimeSeriesCoppock200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Coppock** | **string** | Coppock value | 

## Methods

### NewGetTimeSeriesCoppock200ResponseValuesInner

`func NewGetTimeSeriesCoppock200ResponseValuesInner(datetime string, coppock string, ) *GetTimeSeriesCoppock200ResponseValuesInner`

NewGetTimeSeriesCoppock200ResponseValuesInner instantiates a new GetTimeSeriesCoppock200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCoppock200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesCoppock200ResponseValuesInnerWithDefaults() *GetTimeSeriesCoppock200ResponseValuesInner`

NewGetTimeSeriesCoppock200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCoppock200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesCoppock200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCoppock

`func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetCoppock() string`

GetCoppock returns the Coppock field if non-nil, zero value otherwise.

### GetCoppockOk

`func (o *GetTimeSeriesCoppock200ResponseValuesInner) GetCoppockOk() (*string, bool)`

GetCoppockOk returns a tuple with the Coppock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoppock

`func (o *GetTimeSeriesCoppock200ResponseValuesInner) SetCoppock(v string)`

SetCoppock sets Coppock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


