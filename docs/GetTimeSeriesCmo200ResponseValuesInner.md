# GetTimeSeriesCmo200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Cmo** | **string** | CMO value | 

## Methods

### NewGetTimeSeriesCmo200ResponseValuesInner

`func NewGetTimeSeriesCmo200ResponseValuesInner(datetime string, cmo string, ) *GetTimeSeriesCmo200ResponseValuesInner`

NewGetTimeSeriesCmo200ResponseValuesInner instantiates a new GetTimeSeriesCmo200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCmo200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesCmo200ResponseValuesInnerWithDefaults() *GetTimeSeriesCmo200ResponseValuesInner`

NewGetTimeSeriesCmo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCmo200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesCmo200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesCmo200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesCmo200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCmo

`func (o *GetTimeSeriesCmo200ResponseValuesInner) GetCmo() string`

GetCmo returns the Cmo field if non-nil, zero value otherwise.

### GetCmoOk

`func (o *GetTimeSeriesCmo200ResponseValuesInner) GetCmoOk() (*string, bool)`

GetCmoOk returns a tuple with the Cmo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmo

`func (o *GetTimeSeriesCmo200ResponseValuesInner) SetCmo(v string)`

SetCmo sets Cmo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


