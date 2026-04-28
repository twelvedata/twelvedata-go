# GetTimeSeriesAdxr200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Adxr** | **string** | Adxr value | 

## Methods

### NewGetTimeSeriesAdxr200ResponseValuesInner

`func NewGetTimeSeriesAdxr200ResponseValuesInner(datetime string, adxr string, ) *GetTimeSeriesAdxr200ResponseValuesInner`

NewGetTimeSeriesAdxr200ResponseValuesInner instantiates a new GetTimeSeriesAdxr200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAdxr200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAdxr200ResponseValuesInnerWithDefaults() *GetTimeSeriesAdxr200ResponseValuesInner`

NewGetTimeSeriesAdxr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAdxr200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAdxr200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAdxr

`func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetAdxr() string`

GetAdxr returns the Adxr field if non-nil, zero value otherwise.

### GetAdxrOk

`func (o *GetTimeSeriesAdxr200ResponseValuesInner) GetAdxrOk() (*string, bool)`

GetAdxrOk returns a tuple with the Adxr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdxr

`func (o *GetTimeSeriesAdxr200ResponseValuesInner) SetAdxr(v string)`

SetAdxr sets Adxr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


