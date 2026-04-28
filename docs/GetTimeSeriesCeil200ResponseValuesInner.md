# GetTimeSeriesCeil200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Ceil** | **string** | Ceil value | 

## Methods

### NewGetTimeSeriesCeil200ResponseValuesInner

`func NewGetTimeSeriesCeil200ResponseValuesInner(datetime string, ceil string, ) *GetTimeSeriesCeil200ResponseValuesInner`

NewGetTimeSeriesCeil200ResponseValuesInner instantiates a new GetTimeSeriesCeil200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCeil200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesCeil200ResponseValuesInnerWithDefaults() *GetTimeSeriesCeil200ResponseValuesInner`

NewGetTimeSeriesCeil200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCeil200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesCeil200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesCeil200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesCeil200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCeil

`func (o *GetTimeSeriesCeil200ResponseValuesInner) GetCeil() string`

GetCeil returns the Ceil field if non-nil, zero value otherwise.

### GetCeilOk

`func (o *GetTimeSeriesCeil200ResponseValuesInner) GetCeilOk() (*string, bool)`

GetCeilOk returns a tuple with the Ceil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeil

`func (o *GetTimeSeriesCeil200ResponseValuesInner) SetCeil(v string)`

SetCeil sets Ceil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


