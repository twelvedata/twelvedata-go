# GetTimeSeriesWclPrice200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Wclprice** | **string** | wclprice value | 

## Methods

### NewGetTimeSeriesWclPrice200ResponseValuesInner

`func NewGetTimeSeriesWclPrice200ResponseValuesInner(datetime string, wclprice string, ) *GetTimeSeriesWclPrice200ResponseValuesInner`

NewGetTimeSeriesWclPrice200ResponseValuesInner instantiates a new GetTimeSeriesWclPrice200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesWclPrice200ResponseValuesInner`

NewGetTimeSeriesWclPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesWclPrice200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetWclprice

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetWclprice() string`

GetWclprice returns the Wclprice field if non-nil, zero value otherwise.

### GetWclpriceOk

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) GetWclpriceOk() (*string, bool)`

GetWclpriceOk returns a tuple with the Wclprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWclprice

`func (o *GetTimeSeriesWclPrice200ResponseValuesInner) SetWclprice(v string)`

SetWclprice sets Wclprice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


