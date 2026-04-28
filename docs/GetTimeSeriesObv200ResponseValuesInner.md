# GetTimeSeriesObv200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Obv** | **string** | obv value | 

## Methods

### NewGetTimeSeriesObv200ResponseValuesInner

`func NewGetTimeSeriesObv200ResponseValuesInner(datetime string, obv string, ) *GetTimeSeriesObv200ResponseValuesInner`

NewGetTimeSeriesObv200ResponseValuesInner instantiates a new GetTimeSeriesObv200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesObv200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesObv200ResponseValuesInnerWithDefaults() *GetTimeSeriesObv200ResponseValuesInner`

NewGetTimeSeriesObv200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesObv200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesObv200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesObv200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesObv200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetObv

`func (o *GetTimeSeriesObv200ResponseValuesInner) GetObv() string`

GetObv returns the Obv field if non-nil, zero value otherwise.

### GetObvOk

`func (o *GetTimeSeriesObv200ResponseValuesInner) GetObvOk() (*string, bool)`

GetObvOk returns a tuple with the Obv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObv

`func (o *GetTimeSeriesObv200ResponseValuesInner) SetObv(v string)`

SetObv sets Obv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


