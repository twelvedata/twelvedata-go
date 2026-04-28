# GetTimeSeriesLn200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Ln** | **string** | Natural logarithm value | 

## Methods

### NewGetTimeSeriesLn200ResponseValuesInner

`func NewGetTimeSeriesLn200ResponseValuesInner(datetime string, ln string, ) *GetTimeSeriesLn200ResponseValuesInner`

NewGetTimeSeriesLn200ResponseValuesInner instantiates a new GetTimeSeriesLn200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLn200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesLn200ResponseValuesInnerWithDefaults() *GetTimeSeriesLn200ResponseValuesInner`

NewGetTimeSeriesLn200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLn200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesLn200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesLn200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesLn200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetLn

`func (o *GetTimeSeriesLn200ResponseValuesInner) GetLn() string`

GetLn returns the Ln field if non-nil, zero value otherwise.

### GetLnOk

`func (o *GetTimeSeriesLn200ResponseValuesInner) GetLnOk() (*string, bool)`

GetLnOk returns a tuple with the Ln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLn

`func (o *GetTimeSeriesLn200ResponseValuesInner) SetLn(v string)`

SetLn sets Ln field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


