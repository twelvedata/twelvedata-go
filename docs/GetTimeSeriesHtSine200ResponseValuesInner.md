# GetTimeSeriesHtSine200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**HtSine** | **string** | ht_sine value | 
**HtLeadsine** | **string** | ht_leadsine value | 

## Methods

### NewGetTimeSeriesHtSine200ResponseValuesInner

`func NewGetTimeSeriesHtSine200ResponseValuesInner(datetime string, htSine string, htLeadsine string, ) *GetTimeSeriesHtSine200ResponseValuesInner`

NewGetTimeSeriesHtSine200ResponseValuesInner instantiates a new GetTimeSeriesHtSine200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtSine200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesHtSine200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtSine200ResponseValuesInner`

NewGetTimeSeriesHtSine200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtSine200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetHtSine

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtSine() string`

GetHtSine returns the HtSine field if non-nil, zero value otherwise.

### GetHtSineOk

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtSineOk() (*string, bool)`

GetHtSineOk returns a tuple with the HtSine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtSine

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) SetHtSine(v string)`

SetHtSine sets HtSine field to given value.


### GetHtLeadsine

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtLeadsine() string`

GetHtLeadsine returns the HtLeadsine field if non-nil, zero value otherwise.

### GetHtLeadsineOk

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) GetHtLeadsineOk() (*string, bool)`

GetHtLeadsineOk returns a tuple with the HtLeadsine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtLeadsine

`func (o *GetTimeSeriesHtSine200ResponseValuesInner) SetHtLeadsine(v string)`

SetHtLeadsine sets HtLeadsine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


