# GetTimeSeriesTypPrice200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Typprice** | **string** | typprice value | 

## Methods

### NewGetTimeSeriesTypPrice200ResponseValuesInner

`func NewGetTimeSeriesTypPrice200ResponseValuesInner(datetime string, typprice string, ) *GetTimeSeriesTypPrice200ResponseValuesInner`

NewGetTimeSeriesTypPrice200ResponseValuesInner instantiates a new GetTimeSeriesTypPrice200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTypPrice200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesTypPrice200ResponseValuesInnerWithDefaults() *GetTimeSeriesTypPrice200ResponseValuesInner`

NewGetTimeSeriesTypPrice200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTypPrice200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesTypPrice200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTypprice

`func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetTypprice() string`

GetTypprice returns the Typprice field if non-nil, zero value otherwise.

### GetTyppriceOk

`func (o *GetTimeSeriesTypPrice200ResponseValuesInner) GetTyppriceOk() (*string, bool)`

GetTyppriceOk returns a tuple with the Typprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypprice

`func (o *GetTimeSeriesTypPrice200ResponseValuesInner) SetTypprice(v string)`

SetTypprice sets Typprice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


