# GetTimeSeriesTema200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Tema** | **string** | TEMA value | 

## Methods

### NewGetTimeSeriesTema200ResponseValuesInner

`func NewGetTimeSeriesTema200ResponseValuesInner(datetime string, tema string, ) *GetTimeSeriesTema200ResponseValuesInner`

NewGetTimeSeriesTema200ResponseValuesInner instantiates a new GetTimeSeriesTema200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTema200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesTema200ResponseValuesInnerWithDefaults() *GetTimeSeriesTema200ResponseValuesInner`

NewGetTimeSeriesTema200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesTema200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesTema200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesTema200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesTema200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetTema

`func (o *GetTimeSeriesTema200ResponseValuesInner) GetTema() string`

GetTema returns the Tema field if non-nil, zero value otherwise.

### GetTemaOk

`func (o *GetTimeSeriesTema200ResponseValuesInner) GetTemaOk() (*string, bool)`

GetTemaOk returns a tuple with the Tema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTema

`func (o *GetTimeSeriesTema200ResponseValuesInner) SetTema(v string)`

SetTema sets Tema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


