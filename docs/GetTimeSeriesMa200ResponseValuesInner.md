# GetTimeSeriesMa200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Ma** | **string** | MA value | 

## Methods

### NewGetTimeSeriesMa200ResponseValuesInner

`func NewGetTimeSeriesMa200ResponseValuesInner(datetime string, ma string, ) *GetTimeSeriesMa200ResponseValuesInner`

NewGetTimeSeriesMa200ResponseValuesInner instantiates a new GetTimeSeriesMa200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMa200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMa200ResponseValuesInnerWithDefaults() *GetTimeSeriesMa200ResponseValuesInner`

NewGetTimeSeriesMa200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMa200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMa200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMa200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMa200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMa

`func (o *GetTimeSeriesMa200ResponseValuesInner) GetMa() string`

GetMa returns the Ma field if non-nil, zero value otherwise.

### GetMaOk

`func (o *GetTimeSeriesMa200ResponseValuesInner) GetMaOk() (*string, bool)`

GetMaOk returns a tuple with the Ma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMa

`func (o *GetTimeSeriesMa200ResponseValuesInner) SetMa(v string)`

SetMa sets Ma field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


