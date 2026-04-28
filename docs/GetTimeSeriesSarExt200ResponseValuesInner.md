# GetTimeSeriesSarExt200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Sarext** | **string** | SAREXT value | 

## Methods

### NewGetTimeSeriesSarExt200ResponseValuesInner

`func NewGetTimeSeriesSarExt200ResponseValuesInner(datetime string, sarext string, ) *GetTimeSeriesSarExt200ResponseValuesInner`

NewGetTimeSeriesSarExt200ResponseValuesInner instantiates a new GetTimeSeriesSarExt200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSarExt200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesSarExt200ResponseValuesInnerWithDefaults() *GetTimeSeriesSarExt200ResponseValuesInner`

NewGetTimeSeriesSarExt200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSarExt200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesSarExt200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSarext

`func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetSarext() string`

GetSarext returns the Sarext field if non-nil, zero value otherwise.

### GetSarextOk

`func (o *GetTimeSeriesSarExt200ResponseValuesInner) GetSarextOk() (*string, bool)`

GetSarextOk returns a tuple with the Sarext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSarext

`func (o *GetTimeSeriesSarExt200ResponseValuesInner) SetSarext(v string)`

SetSarext sets Sarext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


