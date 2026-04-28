# GetTimeSeriesMama200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Mama** | **string** | MAMA value | 
**Fama** | **string** | FAMA value | 

## Methods

### NewGetTimeSeriesMama200ResponseValuesInner

`func NewGetTimeSeriesMama200ResponseValuesInner(datetime string, mama string, fama string, ) *GetTimeSeriesMama200ResponseValuesInner`

NewGetTimeSeriesMama200ResponseValuesInner instantiates a new GetTimeSeriesMama200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMama200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMama200ResponseValuesInnerWithDefaults() *GetTimeSeriesMama200ResponseValuesInner`

NewGetTimeSeriesMama200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMama200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMama200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMama200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMama200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMama

`func (o *GetTimeSeriesMama200ResponseValuesInner) GetMama() string`

GetMama returns the Mama field if non-nil, zero value otherwise.

### GetMamaOk

`func (o *GetTimeSeriesMama200ResponseValuesInner) GetMamaOk() (*string, bool)`

GetMamaOk returns a tuple with the Mama field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMama

`func (o *GetTimeSeriesMama200ResponseValuesInner) SetMama(v string)`

SetMama sets Mama field to given value.


### GetFama

`func (o *GetTimeSeriesMama200ResponseValuesInner) GetFama() string`

GetFama returns the Fama field if non-nil, zero value otherwise.

### GetFamaOk

`func (o *GetTimeSeriesMama200ResponseValuesInner) GetFamaOk() (*string, bool)`

GetFamaOk returns a tuple with the Fama field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFama

`func (o *GetTimeSeriesMama200ResponseValuesInner) SetFama(v string)`

SetFama sets Fama field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


