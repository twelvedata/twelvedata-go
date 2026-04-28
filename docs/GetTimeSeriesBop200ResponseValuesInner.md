# GetTimeSeriesBop200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Bop** | **string** | Bop value | 

## Methods

### NewGetTimeSeriesBop200ResponseValuesInner

`func NewGetTimeSeriesBop200ResponseValuesInner(datetime string, bop string, ) *GetTimeSeriesBop200ResponseValuesInner`

NewGetTimeSeriesBop200ResponseValuesInner instantiates a new GetTimeSeriesBop200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesBop200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesBop200ResponseValuesInnerWithDefaults() *GetTimeSeriesBop200ResponseValuesInner`

NewGetTimeSeriesBop200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesBop200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesBop200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesBop200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesBop200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetBop

`func (o *GetTimeSeriesBop200ResponseValuesInner) GetBop() string`

GetBop returns the Bop field if non-nil, zero value otherwise.

### GetBopOk

`func (o *GetTimeSeriesBop200ResponseValuesInner) GetBopOk() (*string, bool)`

GetBopOk returns a tuple with the Bop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBop

`func (o *GetTimeSeriesBop200ResponseValuesInner) SetBop(v string)`

SetBop sets Bop field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


