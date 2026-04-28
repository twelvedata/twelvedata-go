# GetTimeSeriesMfi200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Mfi** | **string** | MFI value | 

## Methods

### NewGetTimeSeriesMfi200ResponseValuesInner

`func NewGetTimeSeriesMfi200ResponseValuesInner(datetime string, mfi string, ) *GetTimeSeriesMfi200ResponseValuesInner`

NewGetTimeSeriesMfi200ResponseValuesInner instantiates a new GetTimeSeriesMfi200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMfi200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMfi200ResponseValuesInnerWithDefaults() *GetTimeSeriesMfi200ResponseValuesInner`

NewGetTimeSeriesMfi200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMfi200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMfi200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMfi200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMfi200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMfi

`func (o *GetTimeSeriesMfi200ResponseValuesInner) GetMfi() string`

GetMfi returns the Mfi field if non-nil, zero value otherwise.

### GetMfiOk

`func (o *GetTimeSeriesMfi200ResponseValuesInner) GetMfiOk() (*string, bool)`

GetMfiOk returns a tuple with the Mfi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfi

`func (o *GetTimeSeriesMfi200ResponseValuesInner) SetMfi(v string)`

SetMfi sets Mfi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


