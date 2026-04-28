# GetTimeSeriesRocr100200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Rocr100** | **string** | rocr100 value | 

## Methods

### NewGetTimeSeriesRocr100200ResponseValuesInner

`func NewGetTimeSeriesRocr100200ResponseValuesInner(datetime string, rocr100 string, ) *GetTimeSeriesRocr100200ResponseValuesInner`

NewGetTimeSeriesRocr100200ResponseValuesInner instantiates a new GetTimeSeriesRocr100200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRocr100200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesRocr100200ResponseValuesInnerWithDefaults() *GetTimeSeriesRocr100200ResponseValuesInner`

NewGetTimeSeriesRocr100200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRocr100200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesRocr100200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetRocr100

`func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetRocr100() string`

GetRocr100 returns the Rocr100 field if non-nil, zero value otherwise.

### GetRocr100Ok

`func (o *GetTimeSeriesRocr100200ResponseValuesInner) GetRocr100Ok() (*string, bool)`

GetRocr100Ok returns a tuple with the Rocr100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocr100

`func (o *GetTimeSeriesRocr100200ResponseValuesInner) SetRocr100(v string)`

SetRocr100 sets Rocr100 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


