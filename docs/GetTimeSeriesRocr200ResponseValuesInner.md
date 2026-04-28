# GetTimeSeriesRocr200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Rocr** | **string** | ROCR value | 

## Methods

### NewGetTimeSeriesRocr200ResponseValuesInner

`func NewGetTimeSeriesRocr200ResponseValuesInner(datetime string, rocr string, ) *GetTimeSeriesRocr200ResponseValuesInner`

NewGetTimeSeriesRocr200ResponseValuesInner instantiates a new GetTimeSeriesRocr200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRocr200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesRocr200ResponseValuesInnerWithDefaults() *GetTimeSeriesRocr200ResponseValuesInner`

NewGetTimeSeriesRocr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRocr200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesRocr200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesRocr200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesRocr200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetRocr

`func (o *GetTimeSeriesRocr200ResponseValuesInner) GetRocr() string`

GetRocr returns the Rocr field if non-nil, zero value otherwise.

### GetRocrOk

`func (o *GetTimeSeriesRocr200ResponseValuesInner) GetRocrOk() (*string, bool)`

GetRocrOk returns a tuple with the Rocr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocr

`func (o *GetTimeSeriesRocr200ResponseValuesInner) SetRocr(v string)`

SetRocr sets Rocr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


