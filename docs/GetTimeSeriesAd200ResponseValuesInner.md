# GetTimeSeriesAd200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Ad** | **string** | AD value | 

## Methods

### NewGetTimeSeriesAd200ResponseValuesInner

`func NewGetTimeSeriesAd200ResponseValuesInner(datetime string, ad string, ) *GetTimeSeriesAd200ResponseValuesInner`

NewGetTimeSeriesAd200ResponseValuesInner instantiates a new GetTimeSeriesAd200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAd200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAd200ResponseValuesInnerWithDefaults() *GetTimeSeriesAd200ResponseValuesInner`

NewGetTimeSeriesAd200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAd200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAd200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAd200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAd200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAd

`func (o *GetTimeSeriesAd200ResponseValuesInner) GetAd() string`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *GetTimeSeriesAd200ResponseValuesInner) GetAdOk() (*string, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *GetTimeSeriesAd200ResponseValuesInner) SetAd(v string)`

SetAd sets Ad field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


