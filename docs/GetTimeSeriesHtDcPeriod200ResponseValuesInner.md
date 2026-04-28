# GetTimeSeriesHtDcPeriod200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**HtDcperiod** | **string** | ht_dcperiod value | 

## Methods

### NewGetTimeSeriesHtDcPeriod200ResponseValuesInner

`func NewGetTimeSeriesHtDcPeriod200ResponseValuesInner(datetime string, htDcperiod string, ) *GetTimeSeriesHtDcPeriod200ResponseValuesInner`

NewGetTimeSeriesHtDcPeriod200ResponseValuesInner instantiates a new GetTimeSeriesHtDcPeriod200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtDcPeriod200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesHtDcPeriod200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtDcPeriod200ResponseValuesInner`

NewGetTimeSeriesHtDcPeriod200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtDcPeriod200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetHtDcperiod

`func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetHtDcperiod() string`

GetHtDcperiod returns the HtDcperiod field if non-nil, zero value otherwise.

### GetHtDcperiodOk

`func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) GetHtDcperiodOk() (*string, bool)`

GetHtDcperiodOk returns a tuple with the HtDcperiod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtDcperiod

`func (o *GetTimeSeriesHtDcPeriod200ResponseValuesInner) SetHtDcperiod(v string)`

SetHtDcperiod sets HtDcperiod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


