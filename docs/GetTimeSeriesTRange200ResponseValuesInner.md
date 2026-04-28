# GetTimeSeriesTRange200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Trange** | **string** | trange value | 

## Methods

### NewGetTimeSeriesTRange200ResponseValuesInner

`func NewGetTimeSeriesTRange200ResponseValuesInner(datetime string, trange string, ) *GetTimeSeriesTRange200ResponseValuesInner`

NewGetTimeSeriesTRange200ResponseValuesInner instantiates a new GetTimeSeriesTRange200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTRange200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesTRange200ResponseValuesInnerWithDefaults() *GetTimeSeriesTRange200ResponseValuesInner`

NewGetTimeSeriesTRange200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTRange200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesTRange200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesTRange200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesTRange200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTrange

`func (o *GetTimeSeriesTRange200ResponseValuesInner) GetTrange() string`

GetTrange returns the Trange field if non-nil, zero value otherwise.

### GetTrangeOk

`func (o *GetTimeSeriesTRange200ResponseValuesInner) GetTrangeOk() (*string, bool)`

GetTrangeOk returns a tuple with the Trange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrange

`func (o *GetTimeSeriesTRange200ResponseValuesInner) SetTrange(v string)`

SetTrange sets Trange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


