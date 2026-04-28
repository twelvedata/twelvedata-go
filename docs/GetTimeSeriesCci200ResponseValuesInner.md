# GetTimeSeriesCci200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Cci** | **string** | CCI value | 

## Methods

### NewGetTimeSeriesCci200ResponseValuesInner

`func NewGetTimeSeriesCci200ResponseValuesInner(datetime string, cci string, ) *GetTimeSeriesCci200ResponseValuesInner`

NewGetTimeSeriesCci200ResponseValuesInner instantiates a new GetTimeSeriesCci200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCci200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesCci200ResponseValuesInnerWithDefaults() *GetTimeSeriesCci200ResponseValuesInner`

NewGetTimeSeriesCci200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCci200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesCci200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesCci200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesCci200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCci

`func (o *GetTimeSeriesCci200ResponseValuesInner) GetCci() string`

GetCci returns the Cci field if non-nil, zero value otherwise.

### GetCciOk

`func (o *GetTimeSeriesCci200ResponseValuesInner) GetCciOk() (*string, bool)`

GetCciOk returns a tuple with the Cci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCci

`func (o *GetTimeSeriesCci200ResponseValuesInner) SetCci(v string)`

SetCci sets Cci field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


