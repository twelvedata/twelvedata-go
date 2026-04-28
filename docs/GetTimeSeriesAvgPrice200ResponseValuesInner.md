# GetTimeSeriesAvgPrice200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Avgprice** | **string** | Avgprice value | 

## Methods

### NewGetTimeSeriesAvgPrice200ResponseValuesInner

`func NewGetTimeSeriesAvgPrice200ResponseValuesInner(datetime string, avgprice string, ) *GetTimeSeriesAvgPrice200ResponseValuesInner`

NewGetTimeSeriesAvgPrice200ResponseValuesInner instantiates a new GetTimeSeriesAvgPrice200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAvgPrice200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAvgPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesAvgPrice200ResponseValuesInner`

NewGetTimeSeriesAvgPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAvgPrice200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAvgprice

`func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetAvgprice() string`

GetAvgprice returns the Avgprice field if non-nil, zero value otherwise.

### GetAvgpriceOk

`func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) GetAvgpriceOk() (*string, bool)`

GetAvgpriceOk returns a tuple with the Avgprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgprice

`func (o *GetTimeSeriesAvgPrice200ResponseValuesInner) SetAvgprice(v string)`

SetAvgprice sets Avgprice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


