# GetTimeSeriesMom200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Mom** | **string** | Mom value | 

## Methods

### NewGetTimeSeriesMom200ResponseValuesInner

`func NewGetTimeSeriesMom200ResponseValuesInner(datetime string, mom string, ) *GetTimeSeriesMom200ResponseValuesInner`

NewGetTimeSeriesMom200ResponseValuesInner instantiates a new GetTimeSeriesMom200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMom200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMom200ResponseValuesInnerWithDefaults() *GetTimeSeriesMom200ResponseValuesInner`

NewGetTimeSeriesMom200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMom200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMom200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMom200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMom200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMom

`func (o *GetTimeSeriesMom200ResponseValuesInner) GetMom() string`

GetMom returns the Mom field if non-nil, zero value otherwise.

### GetMomOk

`func (o *GetTimeSeriesMom200ResponseValuesInner) GetMomOk() (*string, bool)`

GetMomOk returns a tuple with the Mom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMom

`func (o *GetTimeSeriesMom200ResponseValuesInner) SetMom(v string)`

SetMom sets Mom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


