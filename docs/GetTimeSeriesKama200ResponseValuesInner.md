# GetTimeSeriesKama200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Kama** | **string** | Kama value | 

## Methods

### NewGetTimeSeriesKama200ResponseValuesInner

`func NewGetTimeSeriesKama200ResponseValuesInner(datetime string, kama string, ) *GetTimeSeriesKama200ResponseValuesInner`

NewGetTimeSeriesKama200ResponseValuesInner instantiates a new GetTimeSeriesKama200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesKama200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesKama200ResponseValuesInnerWithDefaults() *GetTimeSeriesKama200ResponseValuesInner`

NewGetTimeSeriesKama200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesKama200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesKama200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesKama200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesKama200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetKama

`func (o *GetTimeSeriesKama200ResponseValuesInner) GetKama() string`

GetKama returns the Kama field if non-nil, zero value otherwise.

### GetKamaOk

`func (o *GetTimeSeriesKama200ResponseValuesInner) GetKamaOk() (*string, bool)`

GetKamaOk returns a tuple with the Kama field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKama

`func (o *GetTimeSeriesKama200ResponseValuesInner) SetKama(v string)`

SetKama sets Kama field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


