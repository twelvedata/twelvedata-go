# GetTimeSeriesDx200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Dx** | **string** | dx value | 

## Methods

### NewGetTimeSeriesDx200ResponseValuesInner

`func NewGetTimeSeriesDx200ResponseValuesInner(datetime string, dx string, ) *GetTimeSeriesDx200ResponseValuesInner`

NewGetTimeSeriesDx200ResponseValuesInner instantiates a new GetTimeSeriesDx200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDx200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesDx200ResponseValuesInnerWithDefaults() *GetTimeSeriesDx200ResponseValuesInner`

NewGetTimeSeriesDx200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesDx200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesDx200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesDx200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesDx200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetDx

`func (o *GetTimeSeriesDx200ResponseValuesInner) GetDx() string`

GetDx returns the Dx field if non-nil, zero value otherwise.

### GetDxOk

`func (o *GetTimeSeriesDx200ResponseValuesInner) GetDxOk() (*string, bool)`

GetDxOk returns a tuple with the Dx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDx

`func (o *GetTimeSeriesDx200ResponseValuesInner) SetDx(v string)`

SetDx sets Dx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


