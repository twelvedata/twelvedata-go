# GetTimeSeriesMaxIndex200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Maxidx** | **string** | maxidx value | 

## Methods

### NewGetTimeSeriesMaxIndex200ResponseValuesInner

`func NewGetTimeSeriesMaxIndex200ResponseValuesInner(datetime string, maxidx string, ) *GetTimeSeriesMaxIndex200ResponseValuesInner`

NewGetTimeSeriesMaxIndex200ResponseValuesInner instantiates a new GetTimeSeriesMaxIndex200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMaxIndex200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMaxIndex200ResponseValuesInnerWithDefaults() *GetTimeSeriesMaxIndex200ResponseValuesInner`

NewGetTimeSeriesMaxIndex200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMaxIndex200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMaxidx

`func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetMaxidx() string`

GetMaxidx returns the Maxidx field if non-nil, zero value otherwise.

### GetMaxidxOk

`func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) GetMaxidxOk() (*string, bool)`

GetMaxidxOk returns a tuple with the Maxidx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxidx

`func (o *GetTimeSeriesMaxIndex200ResponseValuesInner) SetMaxidx(v string)`

SetMaxidx sets Maxidx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


